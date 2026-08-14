# FEATURE: Standard Tables

| Field       | Value |
|-------------|-------|
| Version     | 20260814 V1 |
| Description | Parses standard TOML table headers, implicit tables, nested tables, empty tables, and root-table content. |
| Depends On  | FEATURE-Key-Semantics.md, FEATURE-Inline-Table-Closure.md |
| Provides    | TOML standard tables |
| Consumes    | TOML key paths, key-definition validation |

## Behavior

Standard table headers select the table receiving subsequent key/value pairs. The parser creates permitted implicit parent tables, supports nested and empty tables, and treats the nameless root table as the document's initial table.

## Programmatic Acceptance

Requires: executable=sh; scope=test

=== AC standard-tables-valid ===
Intent: The implementation passes the authoritative valid TOML standard-table conformance slice.
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
=== END AC standard-tables-valid ===

=== AC standard-tables-invalid ===
Intent: The implementation rejects invalid standard-table header and hierarchy cases in the authoritative suite.
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
=== END AC standard-tables-invalid ===

## User Acceptance

- None.

## Guardrails

- A standard table cannot be defined more than once.
- Implicit parent tables are created only when permitted by TOML semantics.
- Header key syntax follows the TOML key grammar.
