//go:build integration

package tests

import (
	"bytes"
	"testing"

	"josephcosentino.me/qr-code-tool/internal/cli"
)

func TestGen(t *testing.T) {
	t.Skip("Skip until encode is implemented")

	qrCode := "https://josephcosentino.me"
	qrCodeFile := "/tmp/qr-code.png"

	var stdout, stderr bytes.Buffer
	err := cli.CmdGen(
		[]string{"qr-gen", qrCode, qrCodeFile},
		&stdout,
		&stderr,
	)
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	actualQrCode, err := ReadQRCode(qrCodeFile)
	if err != nil {
		t.Fatalf("Unable to read QR code from %s: %v", qrCodeFile, err)
	}

	if actualQrCode != qrCode {
		t.Fatalf("QR Code does not match. expected: %s, actual: %s", qrCode, actualQrCode)
	}
}
