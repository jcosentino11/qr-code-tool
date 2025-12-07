package util

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func LoadImage(path string) (image.Image, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open image %s: %w", path, err)
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, fmt.Errorf("unable to decode image %s: %w", path, err)
	}

	return img, nil
}
