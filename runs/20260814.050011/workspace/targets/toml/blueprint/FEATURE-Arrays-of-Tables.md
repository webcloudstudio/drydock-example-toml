# FEATURE: Arrays of Tables

| Field       | Value |
|-------------|-------|
| Version     | 20260814 V1 |
| Description | Defines ordered TOML arrays of tables and nested table elements. |
| Depends On  | FEATURE-Standard-Tables.md, FEATURE-Key-Forms.md, FEATURE-Arrays.md |
| Provides    | TOML arrays of tables |
| Consumes    | TOML standard tables, TOML key paths, TOML arrays |

## Purpose

The parser recognizes `[[table]]` headers and creates ordered table elements. Each subsequent key/value pair applies to the newest element. Nested standard tables and nested arrays of tables target the current enclosing element.

## Behavior

- The first array-of-table header creates an array and its first table element.
- Repeated headers append elements in source order.
- Empty elements remain represented as empty objects.
- Nested tables and nested arrays of tables attach to the current element.
- Array elements retain their document order.
- Parsed values use the tagged JSON representation defined by `ARCHITECTURE.md`.

## Programmatic Acceptance

Requires: executable=go; scope=test
Requires: executable=sh; scope=test

=== AC arrays-of-tables-conformance ===
Intent: The implementation passes the authoritative TOML table conformance slice owned by this story.
Suite: scoped

import os
import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/table/*"],
    capture_output=True,
    text=True,
    env={**os.environ, "TOML_TEST_VERSION": "v2.2.0"},
)
print(result.stdout)
print(result.stderr, file=__import__("sys").stderr)
assert result.returncode == 0
=== END AC arrays-of-tables-conformance ===

=== AC arrays-of-tables-build ===
Intent: The decoder builds successfully after array-of-table support is implemented.

import subprocess

result = subprocess.run(["go", "build", "./cmd/toml-decoder"], capture_output=True, text=True)
print(result.stdout)
print(result.stderr, file=__import__("sys").stderr)
assert result.returncode == 0
=== END AC arrays-of-tables-build ===

## User Acceptance

- None.

## Guardrails

- The parser preserves array-of-table element order.
- The parser does not replace an array element with a later element.
