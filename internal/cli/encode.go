package cli

import (
	"fmt"
	"io"

	"josephcosentino.me/qr-tools/internal/encoder"
)

func CmdEncode(args []string, stdout, stderr io.Writer) error {
	url := args[1]
	path := args[2]

	err := encoder.Encode(url, path)
	if err != nil {
		return fmt.Errorf("failed to encode qr code: %w", err)
	}

	return nil
}
