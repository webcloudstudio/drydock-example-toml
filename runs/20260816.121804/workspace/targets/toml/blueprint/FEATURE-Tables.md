# FEATURE: TOML Tables

| Field       | Value |
|-------------|-------|
| Version     | 20260816 V1 |
| Description | Parse standard tables and arrays of tables with TOML lifecycle and conflict rules. |
| Depends On  | ARCHITECTURE.md, FEATURE-Keys.md |
| Provides    | TOML standard tables and arrays-of-tables semantics |
| Consumes    | internal/toml |

## Standard Tables

Support table headers, dotted table paths, implicit parent tables, empty tables, and valid out-of-order definitions. Track table ownership so explicit headers cannot redefine an already-defined table or scalar.

## Arrays of Tables

Support array-of-table headers, creation of elements in encounter order, targeting of the most recently defined element, nested tables, and nested arrays of tables. Reject reversed parent ordering, static-array conflicts, table/array conflicts, and invalid redefinitions.

## Programmatic Acceptance

=== AC tables-valid ===
Intent: The decoder passes all authoritative valid TOML table and array-of-table cases owned by this feature.
Suite: scoped
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/table/**"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC tables-valid ===

=== AC tables-invalid ===
Intent: The decoder rejects all authoritative invalid TOML table and array-of-table cases owned by this feature.
Suite: scoped
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "invalid/table/**"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC tables-invalid ===

## User Acceptance

- None.

## Guardrails

- Arrays of tables must preserve document order.
- A child of an array-of-tables element must resolve to the most recently defined applicable element.
- Do not alter supplied scoring scripts.
