<!-- Compacted from stage_test.sh sha256=93254a9d7ac72f19345b48fc36faffb1dc79dcc66f55cf18fee48c2cdce988cd on 2026-08-14 by drydock build agent -->

Runs one governed conformance slice: `sh sources/stage_test.sh '<toml-test -run pattern>'`. Requires exactly one pattern argument, builds `./cmd/toml-decoder`, sets `TOML_TEST_VERSION` (default `v2.2.0`), and uses the decoder’s exit status as the verdict.
