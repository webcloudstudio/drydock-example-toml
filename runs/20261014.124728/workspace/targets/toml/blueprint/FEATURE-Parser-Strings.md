# FEATURE: Parser Strings

| Field       | Value |
|-------------|-------|
| Version     | 20261014 V1 |
| Description | Decoding of the four TOML string forms: basic, multi-line basic, literal, and multi-line literal. |
| Depends On  | ARCHITECTURE.md |
| Provides    | string value decoding in internal/toml |
| Consumes    | None |

## Purpose

Decode basic strings with the listed escapes only, multi-line basic strings with newline trimming and the line-ending backslash rule, literal strings with no escaping, and multi-line literal strings. Reject any other backslash escape and any control character not permitted by the specification.

## Programmatic Acceptance

=== AC strings-basic-round-trip ===
Intent: A basic string value is decoded and written as a tagged string.
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

key = "s"
payload = "text"
source = f'{key} = "{payload}"\n'
result = subprocess.run(["./toml-decoder"], input=source, capture_output=True, text=True)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
decoded = json.loads(result.stdout)
assert decoded == {key: {"type": "string", "value": payload}}
=== END AC strings-basic-round-trip ===

=== AC strings-bad-escape ===
Intent: A basic string containing an unlisted backslash escape is rejected.
Requires: executable=go; scope=test

import subprocess
import sys

build = subprocess.run(
    ["go", "build", "-o", "toml-decoder", "./cmd/toml-decoder"],
    capture_output=True, text=True,
)
print(build.stderr, file=sys.stderr)
assert build.returncode == 0

escape = "\\q"
source = f'bad = "{escape}"\n'
result = subprocess.run(["./toml-decoder"], input=source, capture_output=True, text=True)
print(result.stderr, file=sys.stderr)
assert result.returncode != 0
assert result.stdout == ""
=== END AC strings-bad-escape ===

=== AC strings-suite-slice ===
Intent: The toml-test string slice executes and every string case passes.
Suite: scoped
Requires: executable=sh; scope=test
Requires: executable=go; scope=test
Requires: executable=toml-test; scope=test

import subprocess
import sys

pattern = "valid/string/**|invalid/string/**"
result = subprocess.run(
    ["sh", "sources/stage_test.sh", pattern],
    capture_output=True, text=True,
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC strings-suite-slice ===

## User Acceptance

- None.

## Guardrails

- Only the escapes listed in the specification are accepted in basic strings.
- Multi-line basic strings trim a newline immediately after the opening delimiter.
