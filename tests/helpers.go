package tests

import (
	"bytes"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"path/filepath"
	"strings"

	"os"
	"testing"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"josephcosentino.me/qr-code-tool/internal/cli"
)

func Encode(t *testing.T, qrCode, qrCodeFile string) (bytes.Buffer, bytes.Buffer) {
	t.Helper()

	var stdout, stderr bytes.Buffer
	err := cli.CmdEncode(
		[]string{"qr-encode", qrCode, qrCodeFile},
		&stdout,
		&stderr,
	)
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	return stdout, stderr
}

func Decode(t *testing.T, qrCodeFile string) string {
	t.Helper()

	var stdout, stderr bytes.Buffer
	err := cli.CmdDecode(
		[]string{"qr-decode", qrCodeFile},
		&stdout,
		&stderr,
	)
	if err != nil {
		t.Fatalf("Decoding failed: %v", err)
	}

	res := stdout.String()

	// trim whitespace added by cli
	res = strings.TrimSpace(res)

	return res
}

func DecodeSampleImage(t *testing.T, qrCodeFile string) string {
	t.Helper()

	path := filepath.Join("..", "samples", qrCodeFile)

	if _, err := os.Stat(path); err != nil {
		t.Fatalf("sample image not found: %s (error: %v)", qrCodeFile, err)
	}

	return Decode(t, path)
}

func DecodeWithExternalLib(t *testing.T, path string) string {
	t.Helper()

	img := LoadImage(t, path)

	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		t.Fatalf("unable to create image bitmap: %v", err)
	}

	reader := qrcode.NewQRCodeReader()

	result, err := reader.Decode(bmp, nil)
	if err != nil {
		t.Fatalf("unable to decode image bitmap: %v", err)
	}

	return result.GetText()
}

func LoadImage(t *testing.T, path string) image.Image {
	t.Helper()

	reader, err := os.Open(path)
	if err != nil {
		t.Fatalf("unable to open image path %s: %v", path, err)
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		t.Fatalf("unable to decode image: %v", err)
	}

	return img
}
