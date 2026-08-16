# FEATURE: Parser Datetimes

| Field       | Value |
|-------------|-------|
| Version     | 20260816 V1 |
| Description | Parses TOML offset and local date/time values and emits their required tagged representations. |
| Depends On  | ARCHITECTURE.md, FEATURE-Decoder-Contract.md, FEATURE-Parser-Lexical-Boundaries.md |
| Provides    | TOML offset datetime, local datetime, local date, and local time parsing |
| Consumes    | internal/toml.Decode |

## Purpose

The parser recognizes offset datetimes, local datetimes, local dates, and local times, including permitted separators, offsets, and fractional precision. It encodes values using the required tagged JSON representation.

## Workflow

1. Classify a scalar token as an offset datetime, local datetime, local date, or local time.
2. Validate calendar, clock, separator, offset, and fractional-second syntax.
3. Truncate unsupported fractional precision rather than rounding.
4. Emit offset datetimes as RFC 3339 values.
5. Emit local values without unrelated offset or date components.
6. Reject malformed and invalid local date/time forms.

## Programmatic Acceptance

Requires: executable=go; scope=test
Requires: executable=sh; scope=test
Requires: executable=toml-test; scope=test

=== AC parser-datetimes-conformance ===
Intent: The implementation passes the authoritative TOML datetime cases owned by this story.
Suite: scoped

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/datetime/**|invalid/datetime/**|invalid/local-date/**|invalid/local-datetime/**|invalid/local-time/**"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC parser-datetimes-conformance ===

=== AC parser-datetimes-types ===
Intent: Offset and local date/time inputs receive the declared tagged types.
import subprocess
import json

source = (
    "offset = 1979-05-27T07:32:00Z\n"
    "local_datetime = 1979-05-27T07:32:00\n"
    "local_date = 1979-05-27\n"
    "local_time = 07:32:00\n"
)
result = subprocess.run(
    ["./toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
decoded = json.loads(result.stdout)
assert decoded["offset"]["type"] == "datetime"
assert decoded["local_datetime"]["type"] == "datetime-local"
assert decoded["local_date"]["type"] == "date-local"
assert decoded["local_time"]["type"] == "time-local"
=== END AC parser-datetimes-types ===

=== AC parser-datetimes-invalid ===
Intent: Invalid local date/time syntax is rejected with a nonzero exit status.
import subprocess

source = "value = 1979-02-30\n"
result = subprocess.run(
    ["./toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode != 0
=== END AC parser-datetimes-invalid ===

## User Acceptance

- None.

## Guardrails

- Offset datetimes use RFC 3339 encoding.
- Local datetime, date, and time values omit unrelated components.
- Invalid calendar and clock values are rejected.
- Excess fractional precision is truncated, not rounded.
