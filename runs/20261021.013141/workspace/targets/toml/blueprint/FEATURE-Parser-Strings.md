# FEATURE: Parser Strings

| Field       | Value |
|-------------|-------|
| Version     | 20261021 V1 |
| Description | Parses TOML basic, literal, and multiline string forms according to TOML 1.0.0. |
| Depends On  | FEATURE-Decoder-Contract.md |
| Provides    | TOML string parsing |
| Consumes    | stdin TOML decoder |

## Scope

Implement basic strings, multiline basic strings, literal strings, and multiline literal strings.

The parser must support valid escapes, Unicode scalar escapes, multiline newline trimming, line-ending backslashes, literal preservation, delimiter rules, and the control-character restrictions defined by TOML 1.0.0. String values must retain the required decoded content in tagged JSON.

## Programmatic Acceptance

Requires: executable=sh; scope=test
Requires: executable=go; scope=test
Requires: executable=toml-test; scope=test

=== AC parser-strings-conformance ===
Intent: The decoder passes the authoritative TOML string cases selected by `valid/string/**|invalid/string/**`; the selector is expected to match the string corpus and the scoped run must execute successfully.

import os
import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/string/**|invalid/string/**"],
    capture_output=True,
    text=True,
    env={**os.environ, "TOML_TEST_VERSION": "v2.2.0"},
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC parser-strings-conformance ===

## User Acceptance

- None.

## Guardrails

- Reject invalid escape sequences and invalid Unicode scalar values.
- Reject prohibited control characters.
- Preserve literal-string backslashes and content without interpreting escapes.
- Do not permit unterminated or malformed string delimiters.
