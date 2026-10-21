<!-- Compacted from ARCHITECTURE.md sha256=00843c32c693ec143737b680714da79857b5ad055ffd9679a6cf73a43ebfc843 on 2026-10-20 by drydock build agent -->

- Go 1.22+ standard-library-only module.
- `cmd/toml-decoder`: stdin/stdout process boundary, diagnostics on stderr, non-zero on invalid TOML.
- `internal/toml`: TOML lexing, parsing, validation, document structure, and tagged values; no file, environment, or service access.
- No third-party dependencies or persistence.
- Required artifacts: `go.mod`, `cmd/toml-decoder/`, `internal/toml/`.
