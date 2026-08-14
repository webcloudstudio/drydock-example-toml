# FEATURE: TOML Arrays

| Field       | Value |
|-------------|-------|
| Version     | 20260814 V1 |
| Description | Parses TOML arrays, including nested, mixed-type, multiline, commented, and trailing-comma arrays. |
| Depends On  | FEATURE-Lexical-Strings.md, FEATURE-Lexical-Numbers.md, FEATURE-Lexical-Datetime.md, FEATURE-Key-Forms.md |
| Provides    | TOML arrays |
| Consumes    | TOML scalar values, TOML key paths |

## Behavior

Arrays are ordered collections enclosed in square brackets. They may be empty, nested, multiline, mixed-type, commented, and terminated by a trailing comma. Each element uses the TOML value grammar and preserves its tagged representation.

## Programmatic Acceptance

Requires: executable=sh; scope=test

=== AC arrays-valid ===
Intent: The implementation passes the authoritative valid TOML array conformance slice.
Suite: scoped

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/array/*"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC arrays-valid ===

=== AC arrays-invalid ===
Intent: The implementation rejects every invalid TOML array case in the authoritative suite.
Suite: scoped

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "invalid/array/*"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC arrays-invalid ===

## User Acceptance

- None.

## Guardrails

- Array order is preserved.
- Trailing commas are accepted only where TOML 1.0.0 permits them.
- Invalid array syntax is rejected.
