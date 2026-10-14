# FEATURE: Decoder Contract

| Field       | Value |
|-------------|-------|
| Version     | 20261014 V1 |
| Description | Stdin-to-tagged-JSON contract, diagnostic stream, and exit status of the decoder command. |
| Depends On  | ARCHITECTURE.md, FEATURE-Parser-Keys.md, FEATURE-Parser-Lexical-Scalars.md, FEATURE-Parser-Arrays-Inline.md |
| Provides    | toml-decoder stdin/stdout filter contract |
| Consumes    | toml.Decode |

## Purpose

The decoder reads TOML from stdin and writes tagged JSON to stdout with exit status 0. On invalid TOML it writes a diagnostic to stderr, writes nothing to stdout, and exits non-zero. It takes no arguments and reads no configuration.

## Tagged JSON Encoding

- Tables encode as JSON objects. Empty tables encode as `{}`.
- Arrays encode as JSON arrays. Empty arrays encode as `[]`.
- Every value encodes as `{"type": "<TYPE>", "value": "<VALUE>"}`, with `value` always a JSON string.
- `type` is one of `string`, `integer`, `float`, `bool`, `datetime`, `datetime-local`, `date-local`, `time-local`.

## Programmatic Acceptance

=== AC decoder-stage-gate ===
Intent: The governed decoder gate passes: the decoder decodes a scalar and rejects an unterminated array.
Requires: executable=sh; scope=test
Requires: executable=go; scope=test
Requires: executable=python; scope=test

import subprocess
import sys

result = subprocess.run(
    ["sh", "sources/stage_contract.sh", "decoder"],
    capture_output=True, text=True,
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC decoder-stage-gate ===

=== AC decoder-scalar-round-trip ===
Intent: A scalar integer is written to stdout as a tagged value and exit status is zero.
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

key = "answer"
value = "42"
source = f"{key} = {value}\n"
result = subprocess.run(["./toml-decoder"], input=source, capture_output=True, text=True)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
decoded = json.loads(result.stdout)
assert decoded == {key: {"type": "integer", "value": value}}
=== END AC decoder-scalar-round-trip ===

=== AC decoder-table-round-trip ===
Intent: A standard table encodes as a nested JSON object with tagged values.
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

table = "tbl"
key = "a"
value = "7"
source = f"[{table}]\n{key} = {value}\n"
result = subprocess.run(["./toml-decoder"], input=source, capture_output=True, text=True)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
decoded = json.loads(result.stdout)
assert decoded == {table: {key: {"type": "integer", "value": value}}}
=== END AC decoder-table-round-trip ===

=== AC decoder-invalid-exit ===
Intent: Malformed TOML exits non-zero, writes nothing to stdout, and writes a diagnostic to stderr.
Requires: executable=go; scope=test

import subprocess
import sys

build = subprocess.run(
    ["go", "build", "-o", "toml-decoder", "./cmd/toml-decoder"],
    capture_output=True, text=True,
)
print(build.stderr, file=sys.stderr)
assert build.returncode == 0

source = "answer = [\n"
result = subprocess.run(["./toml-decoder"], input=source, capture_output=True, text=True)
print(result.stderr, file=sys.stderr)
assert result.returncode != 0
assert result.stdout == ""
assert result.stderr != ""
=== END AC decoder-invalid-exit ===

## User Acceptance

- None.

## Guardrails

- The decoder accepts no command-line arguments and reads no configuration file.
- Diagnostics go to stderr only. Stdout carries only tagged JSON.
