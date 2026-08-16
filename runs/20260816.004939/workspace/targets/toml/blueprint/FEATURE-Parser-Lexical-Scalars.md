# FEATURE: Parser Lexical Scalars

| Field       | Value |
|-------------|-------|
| Version     | 20260816 V1 |
| Description | Defines TOML 1.0.0 boolean, integer, and floating-point scalar parsing. |
| Depends On  | ARCHITECTURE.md, FEATURE-Decoder-Contract.md |
| Provides    | TOML boolean, integer, and float parsing |
| Consumes    | internal/toml.Decode |

## Scope

The parser supports lowercase booleans, signed decimal integers, hexadecimal, octal and binary integers, valid underscore placement, the signed 64-bit range, decimal and exponent floats, special values, and signed zero. Invalid numeric forms are rejected rather than coerced.

Integer values are preserved losslessly and emitted as tagged values with string-valued representations.

## Programmatic Acceptance

=== AC parser-scalars-valid ===
Intent: The authoritative TOML suite passes every valid boolean, integer, and float case.
Suite: scoped
Requires: executable=sh; scope=test
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/bool/**|valid/integer/**|valid/float/**"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC parser-scalars-valid ===

=== AC parser-scalars-invalid ===
Intent: The authoritative TOML suite rejects every invalid boolean, integer, and float case.
Suite: scoped
Requires: executable=sh; scope=test
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "invalid/bool/**|invalid/integer/**|invalid/float/**"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC parser-scalars-invalid ===

=== AC parser-scalars-integer-contract ===
Intent: A signed 64-bit integer is emitted without numeric precision loss.
Requires: executable=go; scope=test

import json
import subprocess

subprocess.run(
    ["go", "build", "-o", "toml-decoder", "./cmd/toml-decoder"],
    check=True,
)
integer_value = "9223372036854775807"
source = f"value = {integer_value}\n"
result = subprocess.run(
    ["./toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
assert result.returncode == 0
decoded = json.loads(result.stdout)
assert decoded["value"]["type"] == "integer"
assert decoded["value"]["value"] == integer_value
=== END AC parser-scalars-integer-contract ===

## User Acceptance

- None.

## Guardrails

- Integer overflow and underflow are rejected.
- Invalid leading zeros, signs, bases, underscores, decimal points, and exponents are rejected.
- No scalar is silently converted to a different TOML type.
