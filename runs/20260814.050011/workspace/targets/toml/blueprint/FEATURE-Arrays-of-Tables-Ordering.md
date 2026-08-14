# FEATURE: Arrays of Tables Ordering

| Field       | Value |
|-------------|-------|
| Version     | 20260814 V1 |
| Description | Rejects invalid ordering and structural conflicts involving TOML arrays of tables. |
| Depends On  | FEATURE-Arrays-of-Tables.md, FEATURE-Table-Redefinition.md |
| Provides    | array-of-table ordering validation |
| Consumes    | TOML arrays of tables, table-definition validation |

## Purpose

The parser enforces TOML 1.0.0 ordering and conflict rules for arrays of tables.

## Behavior

- A child table or child array cannot precede creation of its parent array element.
- A statically defined array cannot later be extended with an array-of-table header.
- An array of tables cannot be redefined as a normal table.
- A normal table cannot be redefined as an array of tables.
- Conflicting nested table and array-of-table declarations are rejected.
- Invalid documents produce a non-zero decoder exit status.

## Programmatic Acceptance

Requires: executable=go; scope=test
Requires: executable=sh; scope=test

=== AC arrays-of-tables-ordering-conformance ===
Intent: The implementation rejects the authoritative invalid table-ordering and conflict cases.
Suite: scoped

import os
import subprocess
import sys

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "invalid/table/*"],
    capture_output=True,
    text=True,
    env={**os.environ, "TOML_TEST_VERSION": "v2.2.0"},
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC arrays-of-tables-ordering-conformance ===

=== AC arrays-of-tables-ordering-exit ===
Intent: An invalid array-of-table ordering document exits unsuccessfully.

source = "[fruit.physical]\ncolor = \"red\"\n[[fruit]]\nname = \"apple\"\n"
result = subprocess.run(
    ["go", "run", "./cmd/toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
assert result.returncode != 0
=== END AC arrays-of-tables-ordering-exit ===

## User Acceptance

- None.

## Guardrails

- The parser rejects invalid ordering rather than assigning ambiguous parent elements.
- The parser does not silently reinterpret conflicting table kinds.
