# FEATURE: TOML Inline Tables

| Field       | Value |
|-------------|-------|
| Version     | 20260814 V1 |
| Description | Parses TOML inline tables, nested inline tables, and dotted keys within inline tables. |
| Depends On  | FEATURE-Lexical-Strings.md, FEATURE-Lexical-Numbers.md, FEATURE-Lexical-Datetime.md, FEATURE-Key-Forms.md |
| Provides    | TOML inline tables |
| Consumes    | TOML scalar values, TOML key paths |

## Behavior

Inline tables are self-contained table values enclosed in braces. They support comma-separated key/value pairs, nested inline tables, dotted keys, and all TOML value types. They remain single-line structures and do not permit trailing commas or prohibited newlines.

## Programmatic Acceptance

Requires: executable=sh; scope=test

=== AC inline-tables-valid ===
Intent: The implementation passes the authoritative valid TOML inline-table conformance slice.
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
=== END AC inline-tables-valid ===

=== AC inline-tables-invalid ===
Intent: The implementation rejects every invalid TOML inline-table case in the authoritative suite.
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
=== END AC inline-tables-invalid ===

## User Acceptance

- None.

## Guardrails

- Inline tables cannot contain trailing commas.
- Inline tables cannot be extended across physical lines except through valid multiline values.
- Every inline-table key is defined within the table's braces.
