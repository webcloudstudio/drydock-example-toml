# FEATURE: TOML Keys

| Field       | Value |
|-------------|-------|
| Version     | 20260816 V1 |
| Description | Parse TOML bare, quoted, dotted, and duplicate-key definitions. |
| Depends On  | ARCHITECTURE.md, FEATURE-Strings.md |
| Provides    | TOML bare, quoted, dotted, and duplicate-key semantics |
| Consumes    | internal/toml |

## Key Parsing

Support bare keys containing ASCII letters, digits, underscores, and dashes; quoted basic and literal keys; and dotted paths combining these forms. Whitespace around dotted components is permitted.

Assignments create nested tables as needed. Numeric-looking bare keys remain string keys. Quoted and bare spellings that resolve to the same key are equivalent.

## Definition Semantics

Reject empty bare keys, duplicate definitions, assignments through scalar values, and attempts to redefine tables or values. Preserve insertion-independent nested object semantics.

## Programmatic Acceptance

=== AC keys-valid ===
Intent: The decoder passes all authoritative valid TOML key cases owned by this feature.
Suite: scoped
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/key/**"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC keys-valid ===

=== AC keys-invalid ===
Intent: The decoder rejects all authoritative invalid TOML key cases owned by this feature.
Suite: scoped
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "invalid/key/**"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC keys-invalid ===

## User Acceptance

- None.

## Guardrails

- Use only the Go standard library.
- Do not reinterpret or modify the supplied conformance harness.
