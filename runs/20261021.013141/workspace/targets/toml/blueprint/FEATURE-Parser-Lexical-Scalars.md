# FEATURE: Parser Lexical Scalars

| Field       | Value |
|-------------|-------|
| Version     | 20261021 V1 |
| Description | Parses TOML booleans, integers, and floating-point scalar values. |
| Depends On  | FEATURE-Decoder-Contract.md |
| Provides    | TOML boolean, integer, and float parsing |
| Consumes    | stdin TOML decoder |

## Scope

Implement lowercase booleans, signed decimal integers, hexadecimal/octal/binary integers, underscore rules, signed 64-bit bounds, decimal fractions, exponents, special float values, and their lexical validation.

Tagged JSON must preserve each scalar's required TOML type and string value representation. Invalid forms, out-of-range integers, malformed underscores, malformed fractions, and unsupported boolean spellings must be rejected.

## Programmatic Acceptance

Requires: executable=sh; scope=test
Requires: executable=go; scope=test
Requires: executable=toml-test; scope=test

=== AC parser-lexical-scalars-conformance ===
Intent: The decoder passes the authoritative boolean, integer, and float cases selected by `valid/bool/**|valid/integer/**|valid/float/**|invalid/bool/**|invalid/integer/**|invalid/float/**`; the selector is expected to match all three scalar corpora and the scoped run must execute successfully.

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
    env={**os.environ, "TOML_TEST_VERSION": "v2.2.0"},
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC parser-lexical-scalars-conformance ===

=== AC parser-lexical-scalars-boundary ===
Intent: The scalar parser preserves the decoder's executable boundary after scalar support is added.

from pathlib import Path

assert Path("toml-decoder").is_file()
=== END AC parser-lexical-scalars-boundary ===

## User Acceptance

- None.

## Guardrails

- Do not silently round or reinterpret integers that cannot be represented losslessly.
- Reject malformed numeric syntax rather than accepting a permissive variant.
- Preserve the required tagged type for booleans, integers, and floats.
