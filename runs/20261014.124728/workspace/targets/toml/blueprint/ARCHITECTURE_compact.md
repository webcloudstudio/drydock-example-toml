<!-- Compacted from ARCHITECTURE.md sha256=97687a3f941d71ba070274e7fe1318c1951bc871777f1e85bac25312dc5ed387 on 2026-10-14 by drydock build agent -->

Module: github.com/owner/toml-decoder (go 1.22, no require directives, stdlib only).
Packages:
- cmd/toml-decoder (package main): reads stdin, calls toml.Decode, JSON-encodes result to stdout; errors to stderr, exit 1. No args, no config, no side effects.
- internal/toml: Decode(r io.Reader) (map[string]any, error); ErrNotImplemented sentinel (stub in this block).
Constraints: no third-party TOML modules; malformed input exits non-zero; scoring assets under sources/ are read-only.
