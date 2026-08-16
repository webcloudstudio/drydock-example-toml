# FEATURE: Parser Lexical Boundaries

| Field       | Value |
|-------------|-------|
| Version     | 20260816 V1 |
| Description | Defines TOML whitespace, comment, control-character, and UTF-8 validation. |
| Depends On  | ARCHITECTURE.md, FEATURE-Decoder-Contract.md |
| Provides    | TOML whitespace, comment, control-character, and UTF-8 validation |
| Consumes    | internal/toml.Decode |

## Scope

The parser recognizes spaces and tabs as whitespace, supports LF and CRLF line endings, treats hashes as comments outside strings, and rejects forbidden control characters and invalid UTF-8. Line boundaries must prevent adjacent key/value pairs from being accepted as one statement.

## Programmatic Acceptance

=== AC parser-boundaries-valid ===
Intent: The authoritative TOML suite passes every valid comment and boundary case.
Suite: scoped
Requires: executable=sh; scope=test
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/comment/**"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC parser-boundaries-valid ===

=== AC parser-boundaries-invalid-control ===
Intent: The authoritative TOML suite rejects every invalid control-character case.
Suite: scoped
Requires: executable=sh; scope=test
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "invalid/control/**"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC parser-boundaries-invalid-control ===

=== AC parser-boundaries-invalid-encoding ===
Intent: The authoritative TOML suite rejects every invalid UTF-8 encoding case.
Suite: scoped
Requires: executable=sh; scope=test
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "invalid/encoding/**"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC parser-boundaries-invalid-encoding ===

## User Acceptance

- None.

## Guardrails

- Invalid UTF-8 is never accepted.
- Forbidden control characters are rejected in comments, keys, and values according to TOML 1.0.0.
- Comment markers inside strings are preserved as string content.
- Adjacent key/value pairs without a valid line boundary are rejected.
