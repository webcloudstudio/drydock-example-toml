# FEATURE: Parser Datetimes

| Field       | Value |
|-------------|-------|
| Version     | 20260815 V1 |
| Description | Parses TOML offset and local date/time values into tagged representations. |
| Depends On  | FEATURE-Parser-Lexical-Boundaries.md |
| Provides    | offset datetime, local datetime, local date, and local time values |
| Consumes    | stdin TOML decoder, tagged JSON output |

## Purpose

Recognize TOML v1.0.0 date and time forms and serialize them using the decoder's required tagged JSON representation.

## Behavior

- Parse RFC 3339 offset datetimes with `Z` and numeric offsets.
- Parse local datetimes without an offset.
- Parse local dates and local times.
- Accept permitted fractional seconds and truncate unsupported precision rather than round.
- Reject malformed dates, times, offsets, and invalid local forms.
- Emit the types `datetime`, `datetime-local`, `date-local`, and `time-local`.
- Serialize values as the required RFC 3339 or date/time strings.

## Programmatic Acceptance

=== AC parser-datetimes-conformance ===
Intent: The implementation passes the complete authoritative date/time conformance slice.
Suite: scoped
Requires: executable=go; scope=test
Requires: executable=toml-test; scope=test

import subprocess

result = subprocess.run(
    [
        "sh",
        "sources/stage_test.sh",
        "valid/datetime/**|invalid/datetime/**|invalid/local-date/**|invalid/local-datetime/**|invalid/local-time/**",
    ],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC parser-datetimes-conformance ===

## User Acceptance

- None.

## Guardrails

- Do not attach local date/time values to an implicit timezone.
- Truncate excess fractional precision; never round it.
