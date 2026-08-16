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
		fmt.Fprintln(os.Stderr, "toml-decoder takes no arguments")
		os.Exit(1)
	}

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	tagged, err := toml.Decode(string(input))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := json.NewEncoder(os.Stdout).Encode(tagged); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
