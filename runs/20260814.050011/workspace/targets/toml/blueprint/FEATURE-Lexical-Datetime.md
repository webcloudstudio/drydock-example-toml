# FEATURE: Lexical Datetime

| Field       | Value |
|-------------|-------|
| Version     | 20260814 V1 |
| Description | Parses TOML offset datetime, local datetime, local date, and local time values. |
| Depends On  | ARCHITECTURE.md |
| Provides    | TOML datetime values |
| Consumes    | parser module boundary |

## Behavior

The parser recognizes offset datetimes, local datetimes, local dates, and local times according to TOML 1.0.0. It accepts permitted `T`, `t`, and space separators, preserves local values without an offset, emits the required tagged representations, and truncates excess fractional precision rather than rounding it. Invalid date and time forms are rejected.

## Programmatic Acceptance

Requires: executable=go; scope=test  
Requires: executable=sh; scope=test

=== AC lexical-datetime-valid ===
Intent: The implementation passes the authoritative valid TOML datetime conformance slice.
Suite: scoped

import os
import subprocess
import sys

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/datetime/*"],
    capture_output=True,
    text=True,
    env={**os.environ, "TOML_TEST_VERSION": "v2.2.0"},
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC lexical-datetime-valid ===

=== AC lexical-datetime-invalid ===
Intent: The implementation rejects the authoritative invalid TOML datetime conformance slice.
Suite: scoped

import os
import subprocess
import sys

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "invalid/datetime/*"],
    capture_output=True,
    text=True,
    env={**os.environ, "TOML_TEST_VERSION": "v2.2.0"},
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC lexical-datetime-invalid ===

## User Acceptance

- None.

## Guardrails

- Local date and time values do not acquire an implicit timezone.
- Fractional seconds beyond supported precision are truncated, not rounded.
