package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"toml/internal/toml"
)

func main() { os.Exit(run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr)) }
func run(args []string, in io.Reader, out, errout io.Writer) int {
	if len(args) != 0 {
		fmt.Fprintln(errout, "toml-decoder: arguments are not supported")
		return 2
	}
	b, e := io.ReadAll(in)
	if e != nil {
		fmt.Fprintf(errout, "read TOML: %v\n", e)
		return 1
	}
	v, e := toml.Decode(b)
	if e != nil {
		fmt.Fprintf(errout, "decode TOML: %v\n", e)
		return 1
	}
	if e = json.NewEncoder(out).Encode(v); e != nil {
		fmt.Fprintf(errout, "encode JSON: %v\n", e)
		return 1
	}
	return 0
}
