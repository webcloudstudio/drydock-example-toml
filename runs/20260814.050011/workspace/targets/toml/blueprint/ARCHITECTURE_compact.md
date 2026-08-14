<!-- Compacted from ARCHITECTURE.md sha256=87d3e8f16f04cd340948533f7dcb38d62908346d19b3ce7ae5e181b29422e39e on 2026-08-14 by drydock build agent -->

Go module `github.com/owner/toml-decoder` uses standard library only. `cmd/toml-decoder` reads TOML from stdin, emits tagged JSON to stdout, reports invalid input on stderr, and exits non-zero. `internal/toml` owns parsing and explicit value kinds while preserving integer lexemes.
