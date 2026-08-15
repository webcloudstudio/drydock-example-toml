# FEATURE: Parser Tables

| Field       | Value |
|-------------|-------|
| Version     | 20260815 V1 |
| Description | Parses standard tables and arrays of tables with TOML definition and targeting rules. |
| Depends On  | FEATURE-Parser-Keys.md |
| Provides    | standard table headers, array-of-table headers, table targeting and redefinition validation |
| Consumes    | stdin TOML decoder, dotted key document paths |

## Purpose

Support table headers and arrays of tables while preserving TOML ordering and redefinition semantics.

## Behavior

- Parse standard table headers with bare, quoted, and dotted names.
- Create implicit parent tables where permitted.
- Reject duplicate table definitions and table/scalar conflicts.
- Parse arrays of tables and append elements in source order.
- Target the most recently defined array element for nested tables and arrays.
- Reject child definitions whose array parent has not yet been established.
- Reject attempts to append to ordinary arrays or redefine tables as arrays.

## Programmatic Acceptance

=== AC parser-tables-conformance ===
Intent: The implementation passes the complete authoritative table conformance slice.
Suite: scoped
Requires: executable=go; scope=test
Requires: executable=toml-test; scope=test

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/table/**|invalid/table/**"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC parser-tables-conformance ===

## User Acceptance

- None.

## Guardrails

- Preserve array-of-table element order.
- Reject table redefinitions and invalid table/array conflicts.
