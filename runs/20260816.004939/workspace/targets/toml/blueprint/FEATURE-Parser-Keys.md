# FEATURE: Parser Keys

| Field       | Value |
|-------------|-------|
| Version     | 20260816 V1 |
| Description | Parses TOML bare, quoted, and dotted keys while enforcing key-definition semantics. |
| Depends On  | ARCHITECTURE.md, FEATURE-Decoder-Contract.md, FEATURE-Parser-Strings.md, FEATURE-Parser-Lexical-Boundaries.md |
| Provides    | TOML key parsing and key-definition semantics |
| Consumes    | internal/toml.Decode |

## Purpose

The parser accepts bare, quoted, and dotted keys, including whitespace around dotted segments and implicit table creation. It rejects duplicate definitions and attempts to extend scalar values as tables.

## Workflow

1. Read a key/value assignment from the current TOML table context.
2. Parse each bare or quoted key segment.
3. Resolve dotted segments into nested tables.
4. Reject duplicate keys and scalar/table conflicts.
5. Store the parsed value in the resulting document tree.

## Programmatic Acceptance

Requires: executable=go; scope=test
Requires: executable=sh; scope=test
Requires: executable=toml-test; scope=test

=== AC parser-keys-conformance ===
Intent: The implementation passes the authoritative TOML key cases owned by this story.
Suite: scoped

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

=== AC parser-keys-invalid-extension ===
Intent: A scalar key cannot later be extended as a dotted table path.
import subprocess

source = "fruit.apple = 1\nfruit.apple.smooth = true\n"
result = subprocess.run(
    ["./toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode != 0
=== END AC parser-keys-invalid-extension ===

=== AC parser-keys-duplicate ===
Intent: Repeating an already-defined key is rejected with a nonzero exit status.
import subprocess

source = "name = \"Tom\"\nname = \"Pradyun\"\n"
result = subprocess.run(
    ["./toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode != 0
=== END AC parser-keys-duplicate ===

## User Acceptance

- None.

## Guardrails

- Bare keys contain only ASCII letters, digits, underscores, and dashes.
- Quoted keys follow TOML string syntax.
- A key must not be defined more than once.
- A scalar value must not be extended as a table.
