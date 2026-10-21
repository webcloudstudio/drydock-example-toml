# FEATURE: Parser Datetimes

| Field       | Value |
|-------------|-------|
| Version     | 20261021 V1 |
| Description | Parse TOML offset and local date and time values. |
| Depends On  | FEATURE-Parser-Keys.md, FEATURE-Parser-Lexical-Scalars.md |
| Provides    | TOML offset datetime, local datetime, local date, and local time parsing |
| Consumes    | TOML key parsing, TOML scalar parsing |

## Purpose

Implement TOML v1.0.0 offset datetime, local datetime, local date, and local time values with tagged JSON representations and required precision handling.

## Behavior

- Offset datetimes accept RFC 3339 forms with `T` or space separators.
- Local datetimes omit the offset and retain local semantics.
- Local dates and local times receive their distinct tagged types.
- Fractional seconds beyond supported precision are truncated rather than rounded.
- Invalid calendar dates, clock values, offsets, and malformed fractional forms are rejected.
- Tagged values use `datetime`, `datetime-local`, `date-local`, and `time-local`.

## Programmatic Acceptance

Requires: executable=go; scope=test
Requires: executable=sh; scope=test
Requires: executable=toml-test; scope=test

=== AC parser-datetimes-conformance ===
Intent: Execute every valid datetime case and every invalid datetime, local-date, local-datetime, and local-time case in the authoritative TOML suite.
Suite: scoped
Probe: --list
Probe Count: lines

import os
import subprocess

selector = "valid/datetime/**|invalid/datetime/**|invalid/local-date/**|invalid/local-datetime/**|invalid/local-time/**"
result = subprocess.run(
    ["sh", "sources/stage_test.sh", selector],
    capture_output=True,
    text=True,
    env={**os.environ, "TOML_TEST_VERSION": "v2.2.0"},
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC parser-datetimes-conformance ===

=== AC parser-datetimes-input-boundaries ===
Intent: Verify that supplied date and time forms receive distinct tagged types and that an invalid calendar value is rejected.
import json
import subprocess

valid_source = (
    'odt = 1979-05-27T07:32:00Z\n'
    'ldt = 1979-05-27T07:32:00\n'
    'date = 1979-05-27\n'
    'time = 07:32:00\n'
)
invalid_source = 'date = 1979-02-30\n'

valid = subprocess.run(
    ["./toml-decoder"],
    input=valid_source,
    capture_output=True,
    text=True,
)
assert valid.returncode == 0
decoded = json.loads(valid.stdout)
assert decoded["odt"]["type"] == "datetime"
assert decoded["ldt"]["type"] == "datetime-local"
assert decoded["date"]["type"] == "date-local"
assert decoded["time"]["type"] == "time-local"

invalid = subprocess.run(
    ["./toml-decoder"],
    input=invalid_source,
    capture_output=True,
    text=True,
)
assert invalid.returncode != 0
assert invalid.stdout == ""
=== END AC parser-datetimes-input-boundaries ===

## User Acceptance

- None.

## Guardrails

- Preserve local versus offset datetime distinctions.
- Truncate unsupported fractional precision rather than rounding.
- Reject impossible dates and times.
