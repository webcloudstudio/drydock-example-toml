# FEATURE: Parser Tables

| Field       | Value |
|-------------|-------|
| Version     | 20261021 V1 |
| Description | Parse TOML tables and arrays of tables with redefinition semantics. |
| Depends On  | FEATURE-Parser-Keys.md |
| Provides    | TOML standard tables and arrays-of-tables |
| Consumes    | TOML key parsing, TOML scalar parsing |

## Purpose

Implement standard tables, implicit parent tables, arrays of tables, nested array elements, and TOML table conflict rules.

## Behavior

- Table headers create the selected table and implicit parents.
- A table cannot be defined more than once.
- Arrays of tables append elements in source order.
- Nested tables and arrays resolve against the most recently defined array element.
- Static arrays and incompatible table forms cannot be converted into tables or arrays of tables.
- Parent arrays must exist before child tables or child arrays are declared.

## Programmatic Acceptance

Requires: executable=go; scope=test
Requires: executable=sh; scope=test
Requires: executable=toml-test; scope=test

=== AC parser-tables-conformance ===
Intent: Execute every valid/table, invalid/table, and applicable spec-1.0.0 case in the authoritative TOML suite.
Suite: scoped
Probe: --list
Probe Count: lines

import os
import subprocess

selector = "valid/table/**|invalid/table/**|valid/spec-1.0.0/**|invalid/spec-1.0.0/**"
result = subprocess.run(
    ["sh", "sources/stage_test.sh", selector],
    capture_output=True,
    text=True,
    env={**os.environ, "TOML_TEST_VERSION": "v2.2.0"},
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC parser-tables-conformance ===

=== AC parser-tables-input-boundaries ===
Intent: Verify that an array of tables preserves element order and that table redefinition is rejected.
import json
import subprocess

valid_source = '[[products]]\nname = "Hammer"\n[[products]]\nname = "Nail"\n'
invalid_source = '[fruit]\napple = "red"\n[fruit]\norange = "orange"\n'

valid = subprocess.run(
    ["./toml-decoder"],
    input=valid_source,
    capture_output=True,
    text=True,
)
assert valid.returncode == 0
decoded = json.loads(valid.stdout)
assert len(decoded["products"]) == 2
assert decoded["products"][0]["name"]["value"] == "Hammer"
assert decoded["products"][1]["name"]["value"] == "Nail"

invalid = subprocess.run(
    ["./toml-decoder"],
    input=invalid_source,
    capture_output=True,
    text=True,
)
assert invalid.returncode != 0
assert invalid.stdout == ""
=== END AC parser-tables-input-boundaries ===

## User Acceptance

- None.

## Guardrails

- Reject table and array-of-table redefinitions and incompatible structural conflicts.
- Preserve array-of-table source order.
- Do not weaken invalid-input rejection to increase valid acceptance.
