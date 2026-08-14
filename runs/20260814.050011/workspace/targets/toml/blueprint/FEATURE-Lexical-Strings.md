# FEATURE: Lexical Strings

| Field       | Value |
|-------------|-------|
| Version     | 20260814 V1 |
| Description | Defines TOML 1.0.0 basic, multiline, literal, and multiline literal string parsing. |
| Depends On  | ARCHITECTURE.md |
| Provides    | TOML string values |
| Consumes    | parser module boundary |

## String Forms

The parser supports basic strings, multiline basic strings, literal strings, and multiline literal strings. Basic strings implement the TOML escape sequences for control characters, quotes, backslashes, and Unicode scalar values. Literal strings preserve content without escape processing.

Multiline strings trim the first newline after the opening delimiter. Multiline basic strings support line-ending backslash continuation and valid TOML escapes. Multiline literal strings preserve content without escaping and reject prohibited control characters.

All string forms require valid UTF-8 and reject prohibited control characters. String values may be used as keys as well as values.

## Conformance Scope

The authoritative acceptance scope is the supplied `toml-test` string suite. The stage runner selects the `valid/string/*` cases and the complete invalid string cases without asserting on runner output.

## Guardrails

- Invalid escape sequences are rejected.
- Unterminated strings are rejected.
- Literal strings do not interpret backslashes as escapes.
- Prohibited control characters are rejected.
- String conformance is measured by the supplied suite.

## Programmatic Acceptance

=== AC strings-conformance ===
Intent: The implementation passes the authoritative TOML 1.0.0 string conformance slice.
Suite: scoped
import os
import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/string/*"],
    capture_output=True,
    text=True,
    env={**os.environ},
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC strings-conformance ===

=== AC strings-invalid-conformance ===
Intent: The implementation rejects the authoritative invalid string cases.
Suite: scoped
import os
import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "invalid/string/*"],
    capture_output=True,
    text=True,
    env={**os.environ},
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC strings-invalid-conformance ===

=== AC strings-process ===
Intent: A basic string is accepted through the decoder process boundary.
import json
import subprocess

value = "line\nvalue"
source = 'text = "line\\nvalue"\n'
result = subprocess.run(
    ["./toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
decoded = json.loads(result.stdout)
assert result.returncode == 0
assert decoded["text"]["type"] == "string"
assert decoded["text"]["value"] == value
=== END AC strings-process ===

=== AC strings-literal ===
Intent: A literal string preserves its backslashes as supplied input.
import json
import subprocess

value = r"C:\Users\nodejs"
source = f"text = '{value}'\n"
result = subprocess.run(
    ["./toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
decoded = json.loads(result.stdout)
assert result.returncode == 0
assert decoded["text"]["type"] == "string"
assert decoded["text"]["value"] == value
=== END AC strings-literal ===

## User Acceptance

- None.

## Guardrails

- None.
