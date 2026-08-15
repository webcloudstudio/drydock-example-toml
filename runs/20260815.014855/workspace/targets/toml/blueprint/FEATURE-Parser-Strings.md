# FEATURE: Parser Strings

| Field       | Value |
|-------------|-------|
| Version     | 20260815 V1 |
| Description | Defines TOML basic, literal, multiline-basic, and multiline-literal string parsing. |
| Depends On  | FEATURE-Decoder-Contract.md |
| Provides    | TOML basic, literal, multiline-basic, and multiline-literal string values |
| Consumes    | stdin TOML decoder, tagged JSON output |

## Behavior

The parser accepts all TOML v1.0.0 string forms and preserves their specified contents. Basic strings decode the permitted escape sequences and reject invalid escapes or invalid Unicode scalar values. Literal strings preserve backslashes without escape processing. Multiline strings support newline trimming, line-ending backslashes, embedded quotation rules, and the control-character restrictions defined by TOML v1.0.0.

## Programmatic Acceptance

=== AC parser-strings-conformance ===
Intent: The authoritative TOML conformance suite passes every valid and invalid string case owned by this story.
Suite: scoped
Requires: executable=sh; scope=test
Requires: executable=go; scope=test

import os
import subprocess

result = subprocess.run(
    [
        "sh",
        "sources/stage_test.sh",
        "valid/string/**|invalid/string/**",
    ],
    capture_output=True,
    text=True,
    env={**os.environ},
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC parser-strings-conformance ===

## User Acceptance

- None.

## Guardrails

- Preserve literal-string backslashes exactly.
- Reject unsupported escapes, unterminated delimiters, invalid control characters, and invalid Unicode scalar values.
- Do not accept permissive string syntax outside TOML v1.0.0.
