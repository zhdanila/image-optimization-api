package imageproc

import "github.com/h2non/bimg"

func CompressFile(imageData []byte, compressionQuality int) ([]byte, error) {
	img := bimg.NewImage(imageData)

	options := bimg.Options{
		Quality: 1,
	}

	compressedImage, err := img.Process(options)
	if err != nil {
		return nil, err
	}

	return compressedImage, nil
}
