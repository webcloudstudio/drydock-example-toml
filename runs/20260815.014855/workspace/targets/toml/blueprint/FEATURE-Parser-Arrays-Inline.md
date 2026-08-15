# FEATURE: Parser Arrays and Inline Tables

| Field       | Value |
|-------------|-------|
| Version     | 20260815 V1 |
| Description | Parses TOML arrays and closed inline tables, including nested and mixed values. |
| Depends On  | FEATURE-Parser-Keys.md |
| Provides    | TOML arrays, nested arrays, mixed values, inline tables, inline-table closure |
| Consumes    | stdin TOML decoder, TOML key parser, TOML scalar and string parsers |

## Purpose

Implement TOML array and inline-table values with their nesting, whitespace, comment, and closure rules.

## Behavior

- Parse empty, single-line, multiline, nested, and mixed-type arrays.
- Permit comments and trailing commas in arrays.
- Parse inline tables containing all supported TOML value types.
- Support nested inline tables and dotted keys inside inline tables.
- Reject trailing commas and newlines that are not valid within inline-table values.
- Reject attempts to extend an inline table after its closing brace.
- Preserve array element order and tagged value types.

## Programmatic Acceptance

=== AC parser-arrays-inline-conformance ===
Intent: The implementation passes the complete authoritative array and inline-table conformance slice.
Suite: scoped
Requires: executable=go; scope=test
Requires: executable=toml-test; scope=test

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/array/**|valid/inline-table/**|invalid/array/**|invalid/inline-table/**"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC parser-arrays-inline-conformance ===

## User Acceptance

- None.

## Guardrails

- Do not allow inline tables to be extended after closure.
- Do not accept trailing commas in inline tables.
