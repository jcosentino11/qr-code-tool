package main

import (
	"fmt"
	"os"

	"josephcosentino.me/qr-tools/internal/cli"
)

func main() {
	if err := cli.CmdEncode(os.Args, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
