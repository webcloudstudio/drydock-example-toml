# FEATURE: Parser Lexical Scalars

| Field       | Value |
|-------------|-------|
| Version     | 20260815 V1 |
| Description | Defines TOML boolean, integer, and floating-point scalar parsing. |
| Depends On  | FEATURE-Parser-Strings.md |
| Provides    | TOML boolean, integer, and float values |
| Consumes    | stdin TOML decoder, tagged JSON output, TOML string lexer |

## Behavior

The parser accepts lowercase booleans, signed decimal integers, hexadecimal, octal, and binary integers, valid underscore placement, signed 64-bit range boundaries, decimal fractions, exponents, special floats, and IEEE-754-compatible signed zero, infinity, and NaN forms. It rejects malformed digits, leading-zero decimal integers, invalid underscores, malformed fractions or exponents, unsupported special-float spellings, and out-of-range integers.

## Programmatic Acceptance

=== AC parser-lexical-scalars-conformance ===
Intent: The authoritative TOML conformance suite passes every valid and invalid boolean, integer, and float case owned by this story.
Suite: scoped
Requires: executable=sh; scope=test
Requires: executable=go; scope=test

import os
import subprocess

result = subprocess.run(
    [
        "sh",
        "sources/stage_test.sh",
        "valid/bool/**|valid/integer/**|valid/float/**|invalid/bool/**|invalid/integer/**|invalid/float/**",
    ],
    capture_output=True,
    text=True,
    env={**os.environ},
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC parser-lexical-scalars-conformance ===

## User Acceptance

- None.

## Guardrails

- Preserve scalar values as tagged JSON objects whose `value` fields are strings.
- Enforce TOML lexical and range rules instead of relying on permissive numeric parsing.
- Do not reinterpret invalid scalar forms as strings.
