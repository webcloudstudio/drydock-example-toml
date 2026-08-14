<!-- Compacted from stage_test.sh sha256=93254a9d7ac72f19345b48fc36faffb1dc79dcc66f55cf18fee48c2cdce988cd on 2026-08-14 by drydock build agent -->

Runs `go build -o toml-decoder ./cmd/toml-decoder`, then executes `sources/run_conformance.sh -run "$1"` against TOML_TEST_VERSION v2.2.0. Requires exactly one test-pattern argument and uses the command exit status as the verdict.
