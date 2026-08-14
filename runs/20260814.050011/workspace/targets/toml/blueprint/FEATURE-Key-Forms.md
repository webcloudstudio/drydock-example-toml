# FEATURE: Key Forms

| Field       | Value |
|-------------|-------|
| Version     | 20260814 V1 |
| Description | Parses bare, quoted, and dotted TOML keys. |
| Depends On  | ARCHITECTURE.md, FEATURE-Lexical-Strings.md, FEATURE-Lexical-Whitespace-Comments.md |
| Provides    | TOML key paths |
| Consumes    | TOML string values, parser module boundary |

## Behavior

The parser recognizes non-empty bare keys, quoted keys using basic or literal string syntax, and dotted keys containing bare or quoted components. Whitespace around dotted components is ignored. Key paths create nested tables while preserving the exact decoded key names.

## Programmatic Acceptance

Requires: executable=go; scope=test  
Requires: executable=sh; scope=test

=== AC key-forms-valid ===
Intent: The implementation passes the authoritative valid TOML key conformance slice.
Suite: scoped

import os
import subprocess
import sys

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/key/*"],
    capture_output=True,
    text=True,
    env={**os.environ, "TOML_TEST_VERSION": "v2.2.0"},
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC key-forms-valid ===

=== AC key-forms-invalid ===
Intent: The implementation rejects malformed key forms in the authoritative suite.
Suite: scoped

import os
import subprocess
import sys

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "invalid/key/*"],
    capture_output=True,
    text=True,
    env={**os.environ, "TOML_TEST_VERSION": "v2.2.0"},
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC key-forms-invalid ===

## User Acceptance

- None.

## Guardrails

- Bare keys accept only ASCII letters, digits, underscores, and dashes.
- Numeric-looking bare keys remain string key names.
