package tests

import (
	"bytes"
	"image"
	"path/filepath"
	"strings"

	"os"
	"testing"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"josephcosentino.me/qr-tools/internal/cli"
	"josephcosentino.me/qr-tools/internal/config"
	"josephcosentino.me/qr-tools/internal/util"
)

func Encode(t *testing.T, qrCode, qrCodeFile string) (bytes.Buffer, bytes.Buffer) {
	t.Helper()

	var stdout, stderr bytes.Buffer
	err := cli.CmdEncode(
		[]string{config.EncodeBinaryName, qrCode, qrCodeFile},
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
		[]string{config.DecodeBinaryName, qrCodeFile},
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

func DecodeSample(t *testing.T, sampleFilename string) string {
	t.Helper()

	path := filepath.Join("..", SamplesDir, sampleFilename)

	if _, err := os.Stat(path); err != nil {
		t.Fatalf("sample image not found: %s (error: %v)", sampleFilename, err)
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

	img, err := util.LoadImage(path)
	if err != nil {
		t.Fatalf("unable to load image %s: %v", path, err)
	}

	return img
}
