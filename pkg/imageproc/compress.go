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
	CompressionQualityHigh   CompressionQuality = 75
	CompressionQualityMedium CompressionQuality = 50
	CompressionQualityLow    CompressionQuality = 25
)

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
			return compressedImages, nil
		}
		compressedImages = append(compressedImages, bind.UploadedFile{
			FileName:    file.FileName,
			ContentType: file.ContentType,
			Size:        file.Size,
			Src:         file.Src,
			Tag:         file.Tag,
		})

		qualities := map[string]CompressionQuality{
			"high":   CompressionQualityHigh,
			"medium": CompressionQualityMedium,
			"low":    CompressionQualityLow,
		}

		for prefix, quality := range qualities {
			compressedSrc, err := compressFile(file.Src, quality)
			if err != nil {
				return nil, err
			}

			compressedImages = append(compressedImages, bind.UploadedFile{
				FileName:    fmt.Sprintf("%s_%s", prefix, file.FileName),
				ContentType: file.ContentType,
				Size:        int64(len(compressedSrc)),
				Src:         compressedSrc,
				Tag:         file.Tag,
			})
		}
	}

	return compressedImages, nil
}
