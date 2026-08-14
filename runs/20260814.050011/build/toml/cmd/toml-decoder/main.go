package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/owner/toml-decoder/internal/toml"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Fprintln(os.Stderr, "toml-decoder accepts no arguments")
		os.Exit(2)
	}
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	parsed, err := toml.Decode(string(input))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	tagged, err := toml.Tagged(parsed)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := json.NewEncoder(os.Stdout).Encode(tagged); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
