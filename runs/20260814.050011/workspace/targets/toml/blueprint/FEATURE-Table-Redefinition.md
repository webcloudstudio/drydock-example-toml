# FEATURE: Table Redefinition

| Field       | Value |
|-------------|-------|
| Version     | 20260814 V1 |
| Description | Rejects conflicting redefinitions among TOML values, standard tables, inline tables, and arrays. |
| Depends On  | FEATURE-Standard-Tables.md, FEATURE-Key-Semantics.md |
| Provides    | table-definition validation |
| Consumes    | TOML standard tables, key-definition validation |

## Behavior

The parser tracks each table and value definition. It rejects duplicate standard-table headers and conflicts that attempt to reinterpret a scalar, inline table, standard table, or array as another structure.

## Programmatic Acceptance

Requires: executable=sh; scope=test

=== AC table-redefinition-invalid ===
Intent: The implementation rejects every authoritative invalid table-redefinition case.
Suite: scoped

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "invalid/table/*"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC table-redefinition-invalid ===

=== AC table-redefinition-valid ===
Intent: The implementation preserves all valid standard-table definitions while enforcing redefinition rules.
Suite: scoped

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/table/*"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC table-redefinition-valid ===

## User Acceptance

- None.

## Guardrails

- Duplicate table definitions are rejected.
- Scalar values cannot be converted into tables.
- Structure conflicts produce a non-successful parse.
