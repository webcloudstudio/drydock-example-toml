# FEATURE: Parser Keys

| Field       | Value |
|-------------|-------|
| Version     | 20261021 V1 |
| Description | Parse TOML key forms and enforce key-definition uniqueness. |
| Depends On  | FEATURE-Decoder-Contract.md, FEATURE-Parser-Strings.md, FEATURE-Parser-Lexical-Scalars.md, FEATURE-Parser-Lexical-Boundaries.md |
| Provides    | TOML bare, quoted, dotted, and numeric-looking key parsing |
| Consumes    | stdin TOML decoder, TOML string parsing, TOML scalar parsing |

## Purpose

Implement TOML v1.0.0 bare, quoted, dotted, empty quoted, and numeric-looking keys. Construct nested tables for dotted keys while preserving quoted key spelling and rejecting duplicate or conflicting definitions.

## Behavior

- Bare keys contain only ASCII letters, digits, underscores, and dashes.
- Quoted keys use basic or literal string syntax and may contain broader Unicode content.
- Dotted keys create nested tables and ignore whitespace around separators.
- Numeric-looking bare keys remain string keys.
- Equivalent key spellings cannot define the same key more than once.
- A value cannot later be reinterpreted as a table.

## Programmatic Acceptance

Requires: executable=go; scope=test
Requires: executable=sh; scope=test
Requires: executable=toml-test; scope=test

=== AC parser-keys-conformance ===
Intent: Execute every valid/key and invalid/key case in the authoritative TOML suite.
Suite: scoped
Probe: --list
Probe Count: lines

import os
import subprocess

selector = "valid/key/**|invalid/key/**"
result = subprocess.run(
    ["sh", "sources/stage_test.sh", selector],
    capture_output=True,
    text=True,
    env={**os.environ, "TOML_TEST_VERSION": "v2.2.0"},
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC parser-keys-conformance ===

=== AC parser-keys-input-boundaries ===
Intent: Verify that a dotted key supplied as input creates nested object state and that a duplicate definition is rejected.
import json
import subprocess

dotted_source = 'fruit.apple = "red"\n'
duplicate_source = 'name = "Tom"\nname = "Pradyun"\n'

dotted = subprocess.run(
    ["./toml-decoder"],
    input=dotted_source,
    capture_output=True,
    text=True,
)
assert dotted.returncode == 0
decoded = json.loads(dotted.stdout)
assert decoded["fruit"]["apple"]["type"] == "string"
assert decoded["fruit"]["apple"]["value"] == "red"

duplicate = subprocess.run(
    ["./toml-decoder"],
    input=duplicate_source,
    capture_output=True,
    text=True,
)
assert duplicate.returncode != 0
assert duplicate.stdout == ""
=== END AC parser-keys-input-boundaries ===

## User Acceptance

- None.

## Guardrails

- Do not permit duplicate definitions through alternate bare and quoted spellings.
- Do not reinterpret scalar values as tables.
- Do not modify the supplied conformance harness.
