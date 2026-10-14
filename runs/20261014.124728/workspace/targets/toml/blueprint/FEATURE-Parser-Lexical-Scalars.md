# FEATURE: Parser Lexical Scalars

| Field       | Value |
|-------------|-------|
| Version     | 20261014 V1 |
| Description | Decoding of booleans, integers, and floats, with range and form rules from the specification. |
| Depends On  | ARCHITECTURE.md |
| Provides    | boolean, integer, and float decoding in internal/toml |
| Consumes    | None |

## Purpose

Decode booleans, decimal and prefixed integers (`0x`, `0o`, `0b`) with underscores, and floats with fractional and exponent parts, plus the special floats `inf`, `nan`, and their signed forms. Integers must fit signed 64-bit and are rejected otherwise.

## Programmatic Acceptance

=== AC scalars-integer-round-trip ===
Intent: The largest signed 64-bit integer decodes losslessly as a tagged integer.
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

key = "n"
number = str(2**63 - 1)
source = f"{key} = {number}\n"
result = subprocess.run(["./toml-decoder"], input=source, capture_output=True, text=True)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
decoded = json.loads(result.stdout)
assert decoded == {key: {"type": "integer", "value": number}}
=== END AC scalars-integer-round-trip ===

=== AC scalars-integer-overflow ===
Intent: An integer one past the signed 64-bit maximum is rejected.
Requires: executable=go; scope=test

import subprocess
import sys

build = subprocess.run(
    ["go", "build", "-o", "toml-decoder", "./cmd/toml-decoder"],
    capture_output=True, text=True,
)
print(build.stderr, file=sys.stderr)
assert build.returncode == 0

key = "n"
over = str(2**63)
source = f"{key} = {over}\n"
result = subprocess.run(["./toml-decoder"], input=source, capture_output=True, text=True)
print(result.stderr, file=sys.stderr)
assert result.returncode != 0
assert result.stdout == ""
=== END AC scalars-integer-overflow ===

=== AC scalars-suite-slice ===
Intent: The toml-test slices for booleans, integers, and floats execute and every case passes.
Suite: scoped
Requires: executable=sh; scope=test
Requires: executable=go; scope=test
Requires: executable=toml-test; scope=test

import subprocess
import sys

pattern = "valid/bool/**|valid/integer/**|valid/float/**|invalid/bool/**|invalid/integer/**|invalid/float/**"
result = subprocess.run(
    ["sh", "sources/stage_test.sh", pattern],
    capture_output=True, text=True,
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC scalars-suite-slice ===

## User Acceptance

- None.

## Guardrails

- Integers are never silently rounded or truncated.
- Leading `+` is not allowed on prefixed integers.
