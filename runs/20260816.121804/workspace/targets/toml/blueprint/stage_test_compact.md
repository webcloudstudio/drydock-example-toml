<!-- Compacted from stage_test.sh sha256=93254a9d7ac72f19345b48fc36faffb1dc79dcc66f55cf18fee48c2cdce988cd on 2026-08-16 by drydock build agent -->

Runs one scoped TOML conformance slice. Requires exactly one `-run` pattern argument, builds `./cmd/toml-decoder`, invokes `sources/run_conformance.sh`, and uses the command exit status as the verdict. Defaults `TOML_TEST_VERSION` to `v2.2.0`.
