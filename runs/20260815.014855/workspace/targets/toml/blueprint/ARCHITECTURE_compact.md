<!-- Compacted from ARCHITECTURE.md sha256=e557add3314a661c6073caf7ff8c08f61479d72e82da73b2ff17c4636a2d45e6 on 2026-08-14 by drydock build agent -->

Go module `github.com/owner/toml-decoder`, Go 1.22+, standard library only. `cmd/toml-decoder` reads stdin, calls `internal/toml.Decode`, emits tagged JSON to stdout, and reports parse errors on stderr with non-zero exit. Tables/arrays map to JSON containers; scalar values use `{type, value}` string fields. No third-party dependencies.
