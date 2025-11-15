//go:build integration

package tests

import (
	"bytes"
	"testing"

	"josephcosentino.me/qr-code-tool/internal/cli"
)

func TestRead(t *testing.T) {
	var stdout, stderr bytes.Buffer
	args := []string{"qr-read", "test.png"}

	err := cli.CmdRead(args, &stdout, &stderr)
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}
}
