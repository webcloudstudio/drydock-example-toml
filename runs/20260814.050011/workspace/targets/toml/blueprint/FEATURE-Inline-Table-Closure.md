# FEATURE: Inline-Table Closure

| Field       | Value |
|-------------|-------|
| Version     | 20260814 V1 |
| Description | Prevents keys and subtables from being added after an inline table is defined. |
| Depends On  | FEATURE-Inline-Tables.md, FEATURE-Key-Semantics.md |
| Provides    | inline-table immutability validation |
| Consumes    | TOML inline tables, TOML key paths |

## Behavior

An inline table and every table established within it become closed when the inline table is parsed. Later dotted keys and table headers cannot add or redefine content beneath that closed structure.

## Programmatic Acceptance

Requires: executable=sh; scope=test

=== AC inline-closure-invalid ===
Intent: The implementation rejects authoritative invalid cases that extend or redefine inline tables.
Suite: scoped

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "invalid/inline-table/*"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC inline-closure-invalid ===

=== AC inline-closure-boundary ===
Intent: The implementation passes the authoritative valid inline-table cases while preserving inline-table closure semantics.
Suite: scoped

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/inline-table/*"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC inline-closure-boundary ===

## User Acceptance

- None.

## Guardrails

- A closed inline table cannot receive later keys.
- A closed inline table cannot receive later subtables.
- Closure validation must not reject valid inline-table documents.
