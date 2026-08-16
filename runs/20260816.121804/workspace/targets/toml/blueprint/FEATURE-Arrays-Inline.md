# FEATURE: TOML Arrays and Inline Tables

| Field       | Value |
|-------------|-------|
| Version     | 20260816 V1 |
| Description | Parse TOML arrays and inline tables, including nesting and closure conflicts. |
| Depends On  | ARCHITECTURE.md, FEATURE-Keys.md, FEATURE-Tables.md |
| Provides    | TOML arrays and inline-table parsing |
| Consumes    | internal/toml |

## Arrays

Support empty, mixed-type, nested, multiline, and commented arrays. Permit trailing commas and arbitrary permitted whitespace or comments around elements and delimiters. Arrays may contain all TOML value types, including inline tables.

## Inline Tables

Support empty and nested inline tables with dotted keys. Inline tables must be fully defined within their braces, must not contain prohibited newlines or trailing commas, and cannot be extended after closure.

Reject conflicts between arrays, tables, inline tables, and later definitions.

## Programmatic Acceptance

=== AC arrays-valid ===
Intent: The decoder passes all authoritative valid TOML array cases owned by this feature.
Suite: scoped
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/array/**"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC arrays-valid ===

=== AC inline-valid ===
Intent: The decoder passes all authoritative valid TOML inline-table cases owned by this feature.
Suite: scoped
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/inline-table/**"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC inline-valid ===

=== AC arrays-invalid ===
Intent: The decoder rejects all authoritative invalid TOML array cases owned by this feature.
Suite: scoped
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "invalid/array/**"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC arrays-invalid ===

=== AC inline-invalid ===
Intent: The decoder rejects all authoritative invalid TOML inline-table cases owned by this feature.
Suite: scoped
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "invalid/inline-table/**"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC inline-invalid ===

## User Acceptance

- None.

## Guardrails

- Do not permit trailing commas or prohibited newlines in inline tables.
- Do not permit later extension of a closed inline table.
- Preserve heterogeneous array values and nested structure.
