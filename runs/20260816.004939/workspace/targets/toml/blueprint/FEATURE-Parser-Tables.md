# FEATURE: Parser Tables

| Field       | Value |
|-------------|-------|
| Version     | 20260816 V1 |
| Description | Implements TOML standard tables and arrays of tables with ordering and conflict rules. |
| Depends On  | ARCHITECTURE.md, FEATURE-Parser-Keys.md |
| Provides    | TOML standard-table and array-of-tables state machines |
| Consumes    | internal/toml.Decode |

## Purpose

The parser implements standard table headers, implicit parents, table redefinition rules, arrays of tables, nested arrays of tables, and most-recent-element targeting.

## Workflow

1. Parse a standard table or array-of-tables header.
2. Resolve its path through existing tables and array elements.
3. Create implicit parents where TOML permits them.
4. Select or append the correct array element.
5. Reject redefinitions, reversed parent ordering, and table/array conflicts.
6. Parse subsequent assignments into the selected table.

## Programmatic Acceptance

Requires: executable=go; scope=test
Requires: executable=sh; scope=test
Requires: executable=toml-test; scope=test

=== AC parser-tables-conformance ===
Intent: The implementation passes the authoritative TOML table and array-of-table cases owned by this story.
Suite: scoped

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

=== AC parser-tables-redefinition ===
Intent: Redefining a standard table is rejected.
import subprocess

source = "[fruit]\napple = \"red\"\n[fruit]\norange = \"orange\"\n"
result = subprocess.run(
    ["./toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode != 0
=== END AC parser-tables-redefinition ===

=== AC parser-tables-array-conflict ===
Intent: Defining an array of tables where a static table already exists is rejected.
import subprocess

source = "[fruits]\nname = \"apple\"\n[[fruits]]\nname = \"orange\"\n"
result = subprocess.run(
    ["./toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode != 0
=== END AC parser-tables-array-conflict ===

## User Acceptance

- None.

## Guardrails

- A standard table cannot be defined more than once.
- Arrays of tables append elements in source order.
- Nested table headers target the most recently defined enclosing array element.
- A table and an array of tables cannot define the same path incompatibly.
