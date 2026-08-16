# FEATURE: TOML Datetimes

| Field       | Value |
|-------------|-------|
| Version     | 20260816 V1 |
| Description | Parse and format TOML offset and local date and time values. |
| Depends On  | ARCHITECTURE.md |
| Provides    | TOML datetime, local datetime, local date, and local time parsing |
| Consumes    | internal/toml |

## Supported Values

Parse offset datetimes, local datetimes, local dates, and local times according to TOML 1.0.0 and RFC 3339-compatible forms. Accept `T` or permitted space separation for offset and local datetimes.

Emit the required tagged types: `datetime`, `datetime-local`, `date-local`, and `time-local`. Offset datetimes use RFC 3339 formatting; local values omit unavailable offset or date components.

## Precision and Validation

Require valid calendar and clock fields. Preserve supported fractional precision and truncate excess fractional seconds rather than rounding. Reject malformed separators, invalid dates or times, invalid offsets, and unsupported local forms.

## Programmatic Acceptance

=== AC datetimes-valid ===
Intent: The decoder passes all authoritative valid TOML datetime cases owned by this feature.
Suite: scoped
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/datetime/**"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC datetimes-valid ===

=== AC datetimes-invalid ===
Intent: The decoder rejects all authoritative invalid TOML datetime, local-date, local-datetime, and local-time cases owned by this feature.
Suite: scoped
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh",
     "invalid/datetime/**|invalid/local-date/**|invalid/local-datetime/**|invalid/local-time/**"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC datetimes-invalid ===

## User Acceptance

- None.

## Guardrails

- Emit datetime values with the required TOML type tags.
- Truncate excess fractional precision; never round it.
- Do not attach an offset to local datetime, date, or time values.
