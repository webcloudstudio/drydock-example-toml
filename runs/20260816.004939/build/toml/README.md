# toml

A standard-library-only Go command-line parser that converts TOML 1.0.0 input into tagged JSON.

## Intent

This project delivers a Go command-line filter that parses TOML 1.0.0 documents and emits the required tagged JSON representation. It serves automated consumers and conformance verification, with correctness defined by the supplied valid and invalid TOML corpus.

## What It Does

The terminal verification runs the supplied unfiltered `sources/full_test.sh` entry point against the installed `toml-test` v2.2.0 suite. The harness exit status is the complete release verdict.

The parser supports empty, nested, mixed, multiline, commented, and trailing-comma arrays, together with nested inline tables and dotted keys inside inline tables. Inline tables remain closed after their terminating brace.

The parser recognizes offset datetimes, local datetimes, local dates, and local times, including permitted separators, offsets, and fractional precision. It encodes values using the required tagged JSON representation.

The parser accepts bare, quoted, and dotted keys, including whitespace around dotted segments and implicit table creation. It rejects duplicate definitions and attempts to extend scalar values as tables.

The parser implements standard table headers, implicit parents, table redefinition rules, arrays of tables, nested arrays of tables, and most-recent-element targeting.

## Next Steps

- Complete the first-run checks above.
