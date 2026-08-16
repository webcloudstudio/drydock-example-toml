<!-- Compacted from ARCHITECTURE.md sha256=196feb56a8e9381d5fafe9f6400ca9200bdf6ea9cf0c3e1f2f937f542af805cb on 2026-08-15 by drydock build agent -->

Go module `github.com/owner/toml-decoder` using Go 1.22+ and no dependencies. `cmd/toml-decoder` reads stdin, calls `internal/toml.Decode`, emits tagged JSON to stdout, and reports errors to stderr with nonzero exit. `internal/toml` owns TOML parsing and tagged values: scalar values use string `type`/`value` fields; tables are objects and arrays are arrays.
