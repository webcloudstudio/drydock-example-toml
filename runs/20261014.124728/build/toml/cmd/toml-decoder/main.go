// Command toml-decoder reads a TOML document on stdin and writes its
// tagged-JSON description to stdout. Diagnostics go to stderr and malformed
// input exits non-zero.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/owner/toml-decoder/internal/toml"
)

func main() {
	if err := checkNoArgs(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "toml-decoder:", err)
		os.Exit(2)
	}
	if err := run(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "toml-decoder:", err)
		os.Exit(1)
	}
}

// checkNoArgs rejects any command-line argument. The decoder is a filter that
// takes its whole input on stdin, so an argument is a caller error, not an option to ignore.
func checkNoArgs(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("takes no arguments, got %q", args)
	}
	return nil
}

// run decodes the document on in and encodes the result as JSON to out.
func run(in io.Reader, out io.Writer) error {
	doc, err := toml.Decode(in)
	if err != nil {
		return fmt.Errorf("decode: %w", err)
	}
	return json.NewEncoder(out).Encode(doc)
}
