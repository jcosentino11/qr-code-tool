//go:build integration

package tests

import (
	"bytes"
	"testing"

	"josephcosentino.me/qr-code-tool/internal/cli"
)

func TestGen(t *testing.T) {
	var stdout, stderr bytes.Buffer
	args := []string{"qr-write", "https://josephcosentino.me"}

	err := cli.CmdGen(args, &stdout, &stderr)
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}
}
