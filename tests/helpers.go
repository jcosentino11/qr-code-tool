package tests

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"os"
	"path/filepath"
	"testing"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func SampleImagePath(t *testing.T, filename string) string {
	t.Helper()

	path := filepath.Join("..", "samples", filename)

	if _, err := os.Stat(path); err != nil {
		t.Fatalf("sample image not found: %s (error: %v)", filename, err)
	}

	return path
}

func ReadQRCode(path string) (string, error) {
	img, err := LoadImage(path)
	if err != nil {
		return "", err
	}

	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return "", err
	}

	reader := qrcode.NewQRCodeReader()

	result, err := reader.Decode(bmp, nil)
	if err != nil {
		return "", err
	}

	return result.GetText(), nil
}

func LoadImage(path string) (image.Image, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	return img, err
}
