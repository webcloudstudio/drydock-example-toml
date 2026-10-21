<!-- Compacted from stage_test.sh sha256=93254a9d7ac72f19345b48fc36faffb1dc79dcc66f55cf18fee48c2cdce988cd on 2026-10-20 by drydock build agent -->

- Usage: `sh sources/stage_test.sh <toml-test -run pattern>`
- Builds `./cmd/toml-decoder`, sets `DECODER` to the resulting binary, and invokes `sources/run_conformance.sh`.
- Uses `TOML_TEST_VERSION` (default `v2.2.0`).
- Exit status is the conformance verdict.
