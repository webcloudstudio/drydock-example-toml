# FEATURE: Parser Arrays Inline

| Field       | Value |
|-------------|-------|
| Version     | 20261014 V1 |
| Description | Arrays, inline tables, and the closure rule that forbids extending an inline table. |
| Depends On  | ARCHITECTURE.md, FEATURE-Parser-Strings.md, FEATURE-Parser-Lexical-Scalars.md, FEATURE-Parser-Keys.md |
| Provides    | array and inline-table parsing in internal/toml |
| Consumes    | None |

## Purpose

Parse arrays of mixed value types, with nesting, optional trailing comma, and multi-line layout with comments. Parse inline tables on a single line with no trailing comma. An inline table is closed: keys cannot be added to it through dotted keys or table headers.

## Programmatic Acceptance

=== AC arrays-round-trip ===
Intent: An array of two integers encodes as a JSON array of tagged values.
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

key = "xs"
first = "1"
second = "2"
source = f"{key} = [{first}, {second}]\n"
result = subprocess.run(["./toml-decoder"], input=source, capture_output=True, text=True)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
decoded = json.loads(result.stdout)
assert decoded == {key: [
    {"type": "integer", "value": first},
    {"type": "integer", "value": second},
]}
=== END AC arrays-round-trip ===

=== AC inline-table-round-trip ===
Intent: An inline table on one line encodes as a nested JSON object.
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

key = "point"
inner = "x"
value = "1"
source = f"{key} = {{ {inner} = {value} }}\n"
result = subprocess.run(["./toml-decoder"], input=source, capture_output=True, text=True)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
decoded = json.loads(result.stdout)
assert decoded == {key: {inner: {"type": "integer", "value": value}}}
=== END AC inline-table-round-trip ===

=== AC inline-table-closed ===
Intent: Adding a key to an inline table through a later table header is rejected.
Requires: executable=go; scope=test

import subprocess
import sys

build = subprocess.run(
    ["go", "build", "-o", "toml-decoder", "./cmd/toml-decoder"],
    capture_output=True, text=True,
)
print(build.stderr, file=sys.stderr)
assert build.returncode == 0

key = "t"
source = f"{key} = {{ a = 1 }}\n[{key}]\nb = 2\n"
result = subprocess.run(["./toml-decoder"], input=source, capture_output=True, text=True)
print(result.stderr, file=sys.stderr)
assert result.returncode != 0
assert result.stdout == ""
=== END AC inline-table-closed ===

=== AC arrays-inline-suite-slice ===
Intent: The toml-test array and inline-table slices execute and every case passes.
Suite: scoped
Requires: executable=sh; scope=test
Requires: executable=go; scope=test
Requires: executable=toml-test; scope=test

import subprocess
import sys

pattern = "valid/array/**|valid/inline-table/**|invalid/array/**|invalid/inline-table/**"
result = subprocess.run(
    ["sh", "sources/stage_test.sh", pattern],
    capture_output=True, text=True,
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC arrays-inline-suite-slice ===

## User Acceptance

- None.

## Guardrails

- An inline table written with a trailing comma is rejected.
- Inline tables may not span lines unless a value inside them does.
