# FEATURE: Parser Arrays Inline

| Field       | Value |
|-------------|-------|
| Version     | 20261021 V1 |
| Description | Parse TOML arrays and inline tables with nesting and closure rules. |
| Depends On  | FEATURE-Parser-Keys.md, FEATURE-Parser-Tables.md |
| Provides    | TOML arrays and inline tables |
| Consumes    | TOML key parsing, TOML scalar parsing, TOML string parsing |

## Purpose

Implement TOML arrays and inline tables, including nesting, mixed values, multiline arrays, trailing commas, and inline-table closure semantics.

## Behavior

- Arrays may be empty, nested, mixed-type, multiline, and trailing-comma terminated.
- Array comments and whitespace are accepted where TOML permits them.
- Inline tables may contain nested inline tables and dotted keys.
- Inline tables do not permit a trailing comma or invalid multiline layout.
- Inline tables are closed to later extension.
- Duplicate keys and structural conflicts are rejected.

## Programmatic Acceptance

Requires: executable=go; scope=test
Requires: executable=sh; scope=test
Requires: executable=toml-test; scope=test

=== AC parser-arrays-inline-conformance ===
Intent: Execute every valid array and inline-table case and every corresponding invalid case in the authoritative TOML suite.
Suite: scoped
Probe: --list
Probe Count: lines

import os
import subprocess

selector = "valid/array/**|valid/inline-table/**|invalid/array/**|invalid/inline-table/**"
result = subprocess.run(
    ["sh", "sources/stage_test.sh", selector],
    capture_output=True,
    text=True,
    env={**os.environ, "TOML_TEST_VERSION": "v2.2.0"},
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC parser-arrays-inline-conformance ===

=== AC parser-arrays-inline-input-boundaries ===
Intent: Verify that a nested mixed array is decoded and that extending an inline table is rejected.
import json
import subprocess

valid_source = 'values = [[1, 2], ["a", "b"]]\n'
invalid_source = '[product]\ntype = { name = "Nail" }\ntype.edible = false\n'

valid = subprocess.run(
    ["./toml-decoder"],
    input=valid_source,
    capture_output=True,
    text=True,
)
assert valid.returncode == 0
decoded = json.loads(valid.stdout)
assert len(decoded["values"]) == 2
assert len(decoded["values"][0]) == 2
assert decoded["values"][0][0]["type"] == "integer"
assert decoded["values"][1][0]["type"] == "string"

invalid = subprocess.run(
    ["./toml-decoder"],
    input=invalid_source,
    capture_output=True,
    text=True,
)
assert invalid.returncode != 0
assert invalid.stdout == ""
=== END AC parser-arrays-inline-input-boundaries ===

## User Acceptance

- None.

## Guardrails

- Do not permit inline-table extension after closure.
- Do not accept trailing commas in inline tables.
- Do not reinterpret malformed arrays as valid values.
