# FEATURE: Parser Tables

| Field       | Value |
|-------------|-------|
| Version     | 20261014 V1 |
| Description | Standard table headers, implicit super-tables, arrays of tables, and their redefinition and ordering rules. |
| Depends On  | ARCHITECTURE.md, FEATURE-Parser-Keys.md |
| Provides    | table headers and arrays of tables in internal/toml |
| Consumes    | None |

## Purpose

Parse `[table]` headers and `[[array]]` headers. Create implicit super-tables for `[a.b.c]`. Each `[[name]]` header appends a new table element to the array named `name`, and later sub-tables attach to the most recent element. Reject a standard table defined twice, and reject a `[[name]]` header that conflicts with a static table or static array.

## Programmatic Acceptance

=== AC tables-header-round-trip ===
Intent: A standard table header creates a nested table holding its keys.
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

table = "owner"
key = "name"
payload = "Regina"
source = f'[{table}]\n{key} = "{payload}"\n'
result = subprocess.run(["./toml-decoder"], input=source, capture_output=True, text=True)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
decoded = json.loads(result.stdout)
assert decoded == {table: {key: {"type": "string", "value": payload}}}
=== END AC tables-header-round-trip ===

=== AC tables-array-of-tables ===
Intent: Two array-of-tables headers produce a JSON array of two objects.
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

name = "products"
key = "sku"
first = "738594937"
second = "284758393"
source = f"[[{name}]]\n{key} = {first}\n\n[[{name}]]\n{key} = {second}\n"
result = subprocess.run(["./toml-decoder"], input=source, capture_output=True, text=True)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
decoded = json.loads(result.stdout)
assert decoded == {name: [
    {key: {"type": "integer", "value": first}},
    {key: {"type": "integer", "value": second}},
]}
=== END AC tables-array-of-tables ===

=== AC tables-duplicate-header-rejected ===
Intent: Defining the same standard table twice is rejected with a non-zero exit.
Requires: executable=go; scope=test

import subprocess
import sys

build = subprocess.run(
    ["go", "build", "-o", "toml-decoder", "./cmd/toml-decoder"],
    capture_output=True, text=True,
)
print(build.stderr, file=sys.stderr)
assert build.returncode == 0

table = "fruit"
source = f"[{table}]\napple = 1\n\n[{table}]\norange = 2\n"
result = subprocess.run(["./toml-decoder"], input=source, capture_output=True, text=True)
print(result.stderr, file=sys.stderr)
assert result.returncode != 0
assert result.stdout == ""
=== END AC tables-duplicate-header-rejected ===

=== AC tables-suite-slice ===
Intent: The toml-test table slices execute and every table case passes.
Suite: scoped
Requires: executable=sh; scope=test
Requires: executable=go; scope=test
Requires: executable=toml-test; scope=test

import subprocess
import sys

pattern = "valid/table/**|invalid/table/**"
result = subprocess.run(
    ["sh", "sources/stage_test.sh", pattern],
    capture_output=True, text=True,
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC tables-suite-slice ===

## User Acceptance

- None.

## Guardrails

- A static table or static array may not be redefined by a `[[name]]` header.
- Implicit super-tables created by a header may later be defined explicitly.
