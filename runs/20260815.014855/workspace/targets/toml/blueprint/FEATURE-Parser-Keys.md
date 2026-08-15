# FEATURE: Parser Keys

| Field       | Value |
|-------------|-------|
| Version     | 20260815 V1 |
| Description | Parses TOML bare, quoted, and dotted keys with correct definition semantics. |
| Depends On  | FEATURE-Parser-Lexical-Boundaries.md |
| Provides    | bare keys, quoted keys, dotted keys, nested document paths, duplicate-key rejection |
| Consumes    | stdin TOML decoder, TOML lexical parsers |

## Purpose

Support TOML key forms and map key/value assignments into nested JSON objects.

## Behavior

- Accept bare keys containing ASCII letters, digits, underscores, and dashes.
- Accept basic-string and literal-string quoted keys, including empty quoted keys.
- Accept dotted keys with whitespace around components.
- Create implicit tables for dotted paths.
- Treat quoted and bare spellings of the same key as equivalent.
- Reject duplicate keys and attempts to extend scalar or inline-table values as tables.
- Preserve numeric-looking bare keys as string keys.

## Programmatic Acceptance

=== AC parser-keys-conformance ===
Intent: The implementation passes the complete authoritative key conformance slice.
Suite: scoped
Requires: executable=go; scope=test
Requires: executable=toml-test; scope=test

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/key/**|invalid/key/**"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC parser-keys-conformance ===

## User Acceptance

- None.

## Guardrails

- Do not accept duplicate or conflicting key definitions.
- Do not reinterpret numeric-looking bare keys as numeric values.
