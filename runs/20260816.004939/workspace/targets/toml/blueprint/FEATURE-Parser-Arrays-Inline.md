# FEATURE: Parser Arrays and Inline Tables

| Field       | Value |
|-------------|-------|
| Version     | 20260816 V1 |
| Description | Parses TOML arrays and inline tables, including nesting, comments, and closure rules. |
| Depends On  | ARCHITECTURE.md, FEATURE-Parser-Keys.md, FEATURE-Parser-Tables.md |
| Provides    | TOML arrays and inline-table parsing |
| Consumes    | internal/toml.Decode |

## Purpose

The parser supports empty, nested, mixed, multiline, commented, and trailing-comma arrays, together with nested inline tables and dotted keys inside inline tables. Inline tables remain closed after their terminating brace.

## Workflow

1. Detect an array or inline-table value.
2. Parse nested values using the TOML value grammar.
3. Permit multiline whitespace, comments, and trailing commas in arrays.
4. Parse comma-separated inline-table entries without a trailing comma.
5. Close inline tables permanently after `}`.
6. Reject malformed delimiters and later extensions of inline tables.

## Programmatic Acceptance

Requires: executable=go; scope=test
Requires: executable=sh; scope=test
Requires: executable=toml-test; scope=test

=== AC parser-arrays-inline-conformance ===
Intent: The implementation passes the authoritative array and inline-table cases owned by this story.
Suite: scoped

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

=== AC parser-arrays-trailing-comma ===
Intent: Arrays permit a trailing comma and preserve their parsed elements.
import subprocess
import json

source = "values = [1, 2,]\n"
result = subprocess.run(
    ["./toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
decoded = json.loads(result.stdout)
assert len(decoded["values"]) == 2
assert decoded["values"][0]["type"] == "integer"
assert decoded["values"][1]["type"] == "integer"
=== END AC parser-arrays-trailing-comma ===

=== AC parser-inline-closure ===
Intent: An inline table cannot be extended by a later assignment.
import subprocess

source = "[product]\ntype = { name = \"Nail\" }\ntype.edible = false\n"
result = subprocess.run(
    ["./toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode != 0
=== END AC parser-inline-closure ===

## User Acceptance

- None.

## Guardrails

- Arrays may be nested and mixed-type.
- Arrays may contain comments, newlines, and a trailing comma.
- Inline tables do not allow a trailing comma or unpermitted newlines.
- Closed inline tables cannot be extended later.
