<!-- Compacted from ARCHITECTURE.md sha256=6646ccb2ae8ad921dc31d34da5e69944e93580c2c27202b96b0aca8e9fa9b284 on 2026-08-16 by drydock build agent -->

Go module `github.com/owner/toml-decoder` using Go 1.22+ and no dependencies. `cmd/toml-decoder` reads stdin, invokes `internal/toml`, emits tagged JSON on stdout, and reports errors on stderr with non-zero exit. `internal/toml` owns TOML parsing. Scalars use `{type, value}` string fields; tables and arrays map to JSON objects and arrays.
