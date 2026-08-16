# FEATURE: Parser Strings

| Field       | Value |
|-------------|-------|
| Version     | 20260816 V1 |
| Description | Defines TOML 1.0.0 basic, literal, and multiline string parsing. |
| Depends On  | ARCHITECTURE.md, FEATURE-Decoder-Contract.md |
| Provides    | TOML string parsing |
| Consumes    | internal/toml.Decode |

## Scope

The parser supports basic strings, literal strings, multiline basic strings, and multiline literal strings. It applies TOML escape rules, Unicode scalar escapes, delimiter rules, newline trimming, line-ending backslashes, and control-character restrictions.

Strings preserve their required TOML value semantics while the decoder emits them as tagged `string` values.

## Programmatic Acceptance

=== AC parser-strings-valid ===
Intent: The authoritative TOML suite passes every valid string case.
Suite: scoped
Requires: executable=sh; scope=test
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/string/**"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC parser-strings-valid ===

=== AC parser-strings-invalid ===
Intent: The authoritative TOML suite rejects every invalid string case.
Suite: scoped
Requires: executable=sh; scope=test
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "invalid/string/**"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC parser-strings-invalid ===

=== AC parser-strings-contract ===
Intent: String parsing remains reachable through the stdin decoder boundary.
Requires: executable=go; scope=test

import json
import subprocess

subprocess.run(
    ["go", "build", "-o", "toml-decoder", "./cmd/toml-decoder"],
    check=True,
)
source = 'message = "hello"\n'
result = subprocess.run(
    ["./toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
assert result.returncode == 0
decoded = json.loads(result.stdout)
assert decoded["message"]["type"] == "string"
assert decoded["message"]["value"] == "hello"
=== END AC parser-strings-contract ===

## User Acceptance

- None.

## Guardrails

- Malformed escapes, unterminated delimiters, invalid control characters, and invalid Unicode scalar escapes are rejected.
- Literal strings do not interpret escape sequences.
- Multiline delimiter and line-continuation rules follow TOML 1.0.0.
