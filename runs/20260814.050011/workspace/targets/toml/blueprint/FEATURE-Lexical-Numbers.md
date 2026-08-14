# FEATURE: Lexical Numbers

| Field       | Value |
|-------------|-------|
| Version     | 20260814 V1 |
| Description | Parses TOML integers, floating-point values, special floats, and booleans. |
| Depends On  | ARCHITECTURE.md |
| Provides    | TOML integer values, TOML float values, TOML boolean values |
| Consumes    | parser module boundary |

## Behavior

The parser recognizes TOML 1.0.0 decimal, hexadecimal, octal, binary, floating-point, exponent, infinity, NaN, and boolean values. Underscore placement, signs, leading-zero rules, and lexical boundaries follow the specification. Signed integers are represented losslessly across the supported 64-bit range. Unsupported or malformed numeric and boolean tokens are rejected.

## Programmatic Acceptance

Requires: executable=go; scope=test  
Requires: executable=sh; scope=test

=== AC lexical-numbers-integer ===
Intent: The implementation passes the authoritative TOML integer conformance slice.
Suite: scoped

import os
import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/integer/*"],
    capture_output=True,
    text=True,
    env={**os.environ, "TOML_TEST_VERSION": "v2.2.0"},
)
print(result.stdout)
print(result.stderr, file=__import__("sys").stderr)
assert result.returncode == 0
=== END AC lexical-numbers-integer ===

=== AC lexical-numbers-float ===
Intent: The implementation passes the authoritative TOML float conformance slice.
Suite: scoped

import os
import subprocess
import sys

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/float/*"],
    capture_output=True,
    text=True,
    env={**os.environ, "TOML_TEST_VERSION": "v2.2.0"},
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC lexical-numbers-float ===

=== AC lexical-numbers-bool ===
Intent: The implementation passes the authoritative TOML boolean conformance slice.
Suite: scoped

import os
import subprocess
import sys

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/bool/*"],
    capture_output=True,
    text=True,
    env={**os.environ, "TOML_TEST_VERSION": "v2.2.0"},
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC lexical-numbers-bool ===

## User Acceptance

- None.

## Guardrails

- The implementation uses only the Go standard library.
- Integer values are not converted through floating-point representations.
