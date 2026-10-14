# FEATURE: Parser Datetimes

| Field       | Value |
|-------------|-------|
| Version     | 20261014 V1 |
| Description | Offset date-time, local date-time, local date, and local time forms, with the fractional-second truncation rule. |
| Depends On  | ARCHITECTURE.md, FEATURE-Parser-Lexical-Scalars.md |
| Provides    | offset datetime, local datetime, local date, and local time decoding |
| Consumes    | None |

## Purpose

Decode RFC 3339 offset date-times, accepting either `T` or a space as the delimiter. Decode local date-times, local dates, and local times without an offset. Fractional seconds beyond the supported precision are truncated, not rounded.

## Programmatic Acceptance

=== AC datetimes-local-date-round-trip ===
Intent: A local date encodes as a date-local tagged value holding the date part.
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

key = "d"
value = "1979-05-27"
source = f"{key} = {value}\n"
result = subprocess.run(["./toml-decoder"], input=source, capture_output=True, text=True)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
decoded = json.loads(result.stdout)
assert decoded == {key: {"type": "date-local", "value": value}}
=== END AC datetimes-local-date-round-trip ===

=== AC datetimes-local-time-round-trip ===
Intent: A local time encodes as a time-local tagged value holding the time part.
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

key = "t"
value = "07:32:00"
source = f"{key} = {value}\n"
result = subprocess.run(["./toml-decoder"], input=source, capture_output=True, text=True)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
decoded = json.loads(result.stdout)
assert decoded == {key: {"type": "time-local", "value": value}}
=== END AC datetimes-local-time-round-trip ===

=== AC datetimes-invalid-date-rejected ===
Intent: A date with an out-of-range month is rejected with a non-zero exit.
Requires: executable=go; scope=test

import subprocess
import sys

build = subprocess.run(
    ["go", "build", "-o", "toml-decoder", "./cmd/toml-decoder"],
    capture_output=True, text=True,
)
print(build.stderr, file=sys.stderr)
assert build.returncode == 0

key = "d"
source = f"{key} = 1979-13-01\n"
result = subprocess.run(["./toml-decoder"], input=source, capture_output=True, text=True)
print(result.stderr, file=sys.stderr)
assert result.returncode != 0
assert result.stdout == ""
=== END AC datetimes-invalid-date-rejected ===

=== AC datetimes-suite-slice ===
Intent: The toml-test datetime slices execute and every datetime case passes.
Suite: scoped
Requires: executable=sh; scope=test
Requires: executable=go; scope=test
Requires: executable=toml-test; scope=test

import subprocess
import sys

pattern = "valid/datetime/**|invalid/datetime/**|invalid/local-date/**|invalid/local-datetime/**|invalid/local-time/**"
result = subprocess.run(
    ["sh", "sources/stage_test.sh", pattern],
    capture_output=True, text=True,
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC datetimes-suite-slice ===

## User Acceptance

- None.

## Guardrails

- Fractional seconds are truncated, never rounded.
- Precision is not invented: a value with fewer digits is not padded.
