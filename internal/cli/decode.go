package cli

import (
	"fmt"
	"io"

	"josephcosentino.me/qr-tools/internal/decoder"
)

func CmdDecode(args []string, stdout, stderr io.Writer) error {
	path := args[1]

	url, err := decoder.Decode(path)
	if err != nil {
		return err
	}

	_, err = stdout.Write([]byte(url + "\n"))
	if err != nil {
		return fmt.Errorf("write to stdout failed %w", err)
	}

	return nil
}
