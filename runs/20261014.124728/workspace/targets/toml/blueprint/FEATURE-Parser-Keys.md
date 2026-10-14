# FEATURE: Parser Keys

| Field       | Value |
|-------------|-------|
| Version     | 20261014 V1 |
| Description | Bare, quoted, and dotted key forms, and the key-definition rules that reject redefinition. |
| Depends On  | ARCHITECTURE.md, FEATURE-Parser-Strings.md, FEATURE-Parser-Lexical-Scalars.md |
| Provides    | bare, quoted, and dotted key parsing with redefinition checks |
| Consumes    | None |

## Purpose

Parse bare keys (`A-Za-z0-9_-`), quoted keys (basic or literal strings), and dotted keys that create intermediate tables. Reject defining the same key twice. A bare key and a quoted key with the same name are equivalent. A key that holds an integer cannot later be used as a table.

## Programmatic Acceptance

=== AC keys-bare-round-trip ===
Intent: A bare key decodes as the named key with a tagged value.
Requires: executable=go; scope=test

import json
import subprocess
import sys

build = subprocess.run(
    ["go", "build", "-o", "toml-decoder", "./cmd/toml-decoder"],
    capture_output=True, text=True,
)
print(build.stderr, file=sys.stderr)
assert build.returncode == 0

key = "bare-key_1"
value = "5"
source = f"{key} = {value}\n"
result = subprocess.run(["./toml-decoder"], input=source, capture_output=True, text=True)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
decoded = json.loads(result.stdout)
assert decoded == {key: {"type": "integer", "value": value}}
=== END AC keys-bare-round-trip ===

=== AC keys-dotted-round-trip ===
Intent: A dotted key creates the intermediate table and stores the value at its leaf.
Requires: executable=go; scope=test

import json
import subprocess
import sys

build = subprocess.run(
    ["go", "build", "-o", "toml-decoder", "./cmd/toml-decoder"],
    capture_output=True, text=True,
)
print(build.stderr, file=sys.stderr)
assert build.returncode == 0

outer = "physical"
inner = "color"
payload = "orange"
source = f'{outer}.{inner} = "{payload}"\n'
result = subprocess.run(["./toml-decoder"], input=source, capture_output=True, text=True)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
decoded = json.loads(result.stdout)
assert decoded == {outer: {inner: {"type": "string", "value": payload}}}
=== END AC keys-dotted-round-trip ===

=== AC keys-redefinition-rejected ===
Intent: Defining the same key twice in one table is rejected with a non-zero exit.
Requires: executable=go; scope=test

import subprocess
import sys

build = subprocess.run(
    ["go", "build", "-o", "toml-decoder", "./cmd/toml-decoder"],
    capture_output=True, text=True,
)
print(build.stderr, file=sys.stderr)
assert build.returncode == 0

key = "name"
source = f'{key} = "Tom"\n{key} = "Pradyun"\n'
result = subprocess.run(["./toml-decoder"], input=source, capture_output=True, text=True)
print(result.stderr, file=sys.stderr)
assert result.returncode != 0
assert result.stdout == ""
=== END AC keys-redefinition-rejected ===

=== AC keys-suite-slice ===
Intent: The toml-test key slices execute and every key case passes.
Suite: scoped
Requires: executable=sh; scope=test
Requires: executable=go; scope=test
Requires: executable=toml-test; scope=test

import subprocess
import sys

pattern = "valid/key/**|invalid/key/**"
result = subprocess.run(
    ["sh", "sources/stage_test.sh", pattern],
    capture_output=True, text=True,
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC keys-suite-slice ===

## User Acceptance

- None.

## Guardrails

- A key defined twice is an error, whether bare or quoted.
- Whitespace around dot-separated key parts is ignored.
