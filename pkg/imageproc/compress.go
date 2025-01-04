package imageproc

import (
	"fmt"
	"github.com/h2non/bimg"
	"image-optimization-api/pkg/bind"
)

type CompressionQuality int

func (cq CompressionQuality) Int() int {
	return int(cq)
}

const (
	CompressionQualityOriginal CompressionQuality = 100
	CompressionQualityHigh     CompressionQuality = 75
	CompressionQualityMedium   CompressionQuality = 50
	CompressionQualityLow      CompressionQuality = 25
)

func IntToCompressionQuality(value int) CompressionQuality {
	switch value {
	case 100:
		return CompressionQualityOriginal
	case 75:
		return CompressionQualityHigh
	case 50:
		return CompressionQualityMedium
	case 25:
		return CompressionQualityLow
	default:
		return CompressionQualityHigh
	}
}

func GetAllCompressionQualities() []CompressionQuality {
	return []CompressionQuality{CompressionQualityHigh, CompressionQualityMedium, CompressionQualityLow, CompressionQualityOriginal}
}

func compressFile(imageData []byte, compressionQuality CompressionQuality) ([]byte, error) {
	img := bimg.NewImage(imageData)

	options := bimg.Options{
		Quality: compressionQuality.Int(),
	}

	compressedImage, err := img.Process(options)
	if err != nil {
		return nil, err
	}

	return compressedImage, nil
}

func GetCompressedImages(images []bind.UploadedFile) ([]bind.UploadedFile, error) {
	var compressedImages []bind.UploadedFile

	for _, file := range images {
		if len(file.Src) == 0 {
			continue
		}

		qualities := GetAllCompressionQualities()
		for _, quality := range qualities {
			compressedSrc, err := compressFile(file.Src, quality)
			if err != nil {
				return nil, err
			}

			compressedImages = append(compressedImages, bind.UploadedFile{
				FileName:    GenerateImageID(file.FileName, quality),
				ContentType: file.ContentType,
				Size:        int64(len(compressedSrc)),
				Src:         compressedSrc,
				Tag:         file.Tag,
			})
		}
	}

	return compressedImages, nil
}

func GenerateImageID(fileName string, quality CompressionQuality) string {
	qualityStr := map[CompressionQuality]string{
		CompressionQualityOriginal: "",
		CompressionQualityHigh:     "high_",
		CompressionQualityMedium:   "medium_",
		CompressionQualityLow:      "low_",
	}[quality]

	return fmt.Sprintf("%s%s", qualityStr, fileName)
}
