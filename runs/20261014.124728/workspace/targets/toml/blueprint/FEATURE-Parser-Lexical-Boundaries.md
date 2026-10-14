# FEATURE: Parser Lexical Boundaries

| Field       | Value |
|-------------|-------|
| Version     | 20261014 V1 |
| Description | Comment syntax, forbidden control characters, and UTF-8 validity of TOML input. |
| Depends On  | ARCHITECTURE.md, FEATURE-Parser-Strings.md, FEATURE-Parser-Lexical-Scalars.md |
| Provides    | comment handling and control-character and UTF-8 validation |
| Consumes    | None |

## Purpose

A hash begins a comment that runs to end of line, except inside a string. Control characters other than tab are rejected in comments and literal strings. Input that is not valid UTF-8 is rejected.

## Programmatic Acceptance

=== AC boundaries-comment-round-trip ===
Intent: A trailing comment is ignored and the value before it decodes normally.
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
value = "1"
note = "note"
source = f"{key} = {value} # {note}\n"
result = subprocess.run(["./toml-decoder"], input=source, capture_output=True, text=True)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
decoded = json.loads(result.stdout)
assert decoded == {key: {"type": "integer", "value": value}}
=== END AC boundaries-comment-round-trip ===

=== AC boundaries-control-in-comment ===
Intent: A control character in a comment is rejected with a non-zero exit.
Requires: executable=go; scope=test

import subprocess
import sys

build = subprocess.run(
    ["go", "build", "-o", "toml-decoder", "./cmd/toml-decoder"],
    capture_output=True, text=True,
)
print(build.stderr, file=sys.stderr)
assert build.returncode == 0

control = chr(1)
source = f"# {control}\nn = 1\n"
result = subprocess.run(["./toml-decoder"], input=source, capture_output=True, text=True)
print(result.stderr, file=sys.stderr)
assert result.returncode != 0
assert result.stdout == ""
=== END AC boundaries-control-in-comment ===

=== AC boundaries-invalid-utf8 ===
Intent: Input that is not valid UTF-8 is rejected with a non-zero exit.
Requires: executable=go; scope=test

import subprocess
import sys

build = subprocess.run(
    ["go", "build", "-o", "toml-decoder", "./cmd/toml-decoder"],
    capture_output=True, text=True,
)
print(build.stderr, file=sys.stderr)
assert build.returncode == 0

result = subprocess.run(["./toml-decoder"], input=b"\xff", capture_output=True)
print(result.stderr, file=sys.stderr)
assert result.returncode != 0
=== END AC boundaries-invalid-utf8 ===

=== AC boundaries-suite-slice ===
Intent: The toml-test comment, control, and encoding slices execute and every case passes.
Suite: scoped
Requires: executable=sh; scope=test
Requires: executable=go; scope=test
Requires: executable=toml-test; scope=test

import subprocess
import sys

pattern = "valid/comment/**|invalid/control/**|invalid/encoding/**"
result = subprocess.run(
    ["sh", "sources/stage_test.sh", pattern],
    capture_output=True, text=True,
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC boundaries-suite-slice ===

## User Acceptance

- None.

## Guardrails

- Control characters other than tab are rejected in comments and literal strings.
- Invalid UTF-8 input is never decoded partially.
