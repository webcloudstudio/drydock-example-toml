package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/owner/toml-decoder/internal/toml"
)

func main() {
	if err := run(os.Stdin, os.Stdout, os.Stderr); err != nil {
		os.Exit(1)
	}
}

func run(input io.Reader, output, diagnostics io.Writer) error {
	source, err := io.ReadAll(input)
	if err != nil {
		_, _ = fmt.Fprintln(diagnostics, err)
		return err
	}
	tagged, err := toml.Decode(string(source))
	if err != nil {
		_, _ = fmt.Fprintln(diagnostics, err)
		return err
	}
	if err := json.NewEncoder(output).Encode(tagged); err != nil {
		_, _ = fmt.Fprintln(diagnostics, err)
		return err
	}
	return nil
}
