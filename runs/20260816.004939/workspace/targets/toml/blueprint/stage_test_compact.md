<!-- Compacted from stage_test.sh sha256=93254a9d7ac72f19345b48fc36faffb1dc79dcc66f55cf18fee48c2cdce988cd on 2026-08-15 by drydock build agent -->

- Runs one authoritative TOML conformance scope.
- Requires exactly one `-run` pattern argument.
- Builds `./cmd/toml-decoder`, then invokes `sources/run_conformance.sh -run "$1"`.
- Exit status is the complete verdict; no output parsing.
