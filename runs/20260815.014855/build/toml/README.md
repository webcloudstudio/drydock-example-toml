# toml

A Go command-line TOML v1.0.0 parser that converts stdin into tagged JSON and rejects invalid documents.

## Intent

This project delivers a command-line TOML v1.0.0 parser for automated callers. It reads one TOML document from standard input and emits a tagged JSON representation, while rejecting malformed documents with a diagnostic and non-zero exit. Correctness is defined by the supplied upstream conformance suite.

## What It Does

Provide the terminal, unfiltered verification of the complete TOML parser implementation.

Implement TOML array and inline-table values with their nesting, whitespace, comment, and closure rules.

Recognize TOML v1.0.0 date and time forms and serialize them using the decoder's required tagged JSON representation.

Support TOML key forms and map key/value assignments into nested JSON objects.

Support table headers and arrays of tables while preserving TOML ordering and redefinition semantics.

## Next Steps

- Complete the first-run checks above.
