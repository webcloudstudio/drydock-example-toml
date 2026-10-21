# toml

A standard-library-only Go parser that decodes TOML 1.0.0 from stdin into tagged JSON.

## Intent

Build a Go command-line filter that parses TOML v1.0.0 and emits the required tagged JSON representation. It serves the supplied conformance harness and must correctly accept valid documents while rejecting invalid ones.

## What It Does

Run the pinned upstream `toml-test` v2.2.0 suite without filtering or reinterpretation and use its exit status as the completed decoder verdict.

Provide a deterministic command-line filter for TOML input. The command delegates parsing to `internal/toml` and serializes the resulting tagged representation as JSON.

Implement TOML arrays and inline tables, including nesting, mixed values, multiline arrays, trailing commas, and inline-table closure semantics.

Implement TOML v1.0.0 offset datetime, local datetime, local date, and local time values with tagged JSON representations and required precision handling.

Implement TOML v1.0.0 bare, quoted, dotted, empty quoted, and numeric-looking keys. Construct nested tables for dotted keys while preserving quoted key spelling and rejecting duplicate or conflicting definitions.

Implement standard tables, implicit parent tables, arrays of tables, nested array elements, and TOML table conflict rules.

## Stack

Go

## Next Steps

- Complete the first-run checks above.
