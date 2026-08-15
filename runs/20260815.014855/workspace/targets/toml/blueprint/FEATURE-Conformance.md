# FEATURE: Complete Conformance Verification

| Field       | Value |
|-------------|-------|
| Version     | 20260815 V1 |
| Description | Verifies the completed decoder against the complete supplied TOML v1.0.0 conformance suite. |
| Depends On  | ARCHITECTURE.md, FEATURE-Decoder-Contract.md, FEATURE-Parser-Strings.md, FEATURE-Parser-Lexical-Scalars.md, FEATURE-Parser-Lexical-Boundaries.md, FEATURE-Parser-Parser-Keys.md, FEATURE-Parser-Tables.md, FEATURE-Parser-Arrays-Inline.md, FEATURE-Parser-Datetimes.md |
| Provides    | complete TOML v1.0.0 conformance verdict |
| Consumes    | completed stdin TOML decoder |

## Purpose

Provide the terminal, unfiltered verification of the complete TOML parser implementation.

## Verification

- Build the fixed `cmd/toml-decoder` entry point.
- Run the supplied `sources/full_test.sh` command without filtering or reinterpretation.
- Require the authoritative `toml-test` v2.2.0 suite to exit successfully.
- Treat the suite exit status as the complete release verification result.

## Programmatic Acceptance

=== AC complete-conformance ===
Intent: The completed implementation passes the complete authoritative TOML v1.0.0 conformance suite.
Suite: full
Requires: executable=go; scope=test
Requires: executable=toml-test; scope=test

import subprocess

result = subprocess.run(
    ["sh", "sources/full_test.sh"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC complete-conformance ===

## User Acceptance

- None.

## Guardrails

- Do not filter, skip, reinterpret, or replace the complete conformance verdict.
- Do not assert or depend on a fixed case count.
