# FEATURE: Conformance

| Field       | Value |
|-------------|-------|
| Version     | 20261021 V1 |
| Description | Verify the completed decoder against the complete supplied TOML conformance suite. |
| Depends On  | FEATURE-Parser-Strings.md, FEATURE-Parser-Lexical-Scalars.md, FEATURE-Parser-Lexical-Boundaries.md, FEATURE-Parser-Keys.md, FEATURE-Parser-Tables.md, FEATURE-Parser-Arrays-Inline.md, FEATURE-Parser-Datetimes.md |
| Provides    | complete TOML conformance verification |
| Consumes    | completed toml-decoder |

## Purpose

Run the pinned upstream `toml-test` v2.2.0 suite without filtering or reinterpretation and use its exit status as the completed decoder verdict.

## Verification Contract

- `sources/full_test.sh` builds the decoder and runs the complete suite.
- The harness assets remain unchanged.
- `TOML_TEST_VERSION` identifies the expected suite version.
- A zero exit status is required; any non-zero status fails conformance.
- The suite is the sole release-level correctness measurement.

## Programmatic Acceptance

Requires: executable=go; scope=test
Requires: executable=sh; scope=test
Requires: executable=toml-test; scope=test

=== AC complete-conformance ===
Intent: Run the complete pinned TOML v1.0.0 conformance suite after all parser capabilities are implemented.
Suite: full

import os
import subprocess

result = subprocess.run(
    ["sh", "sources/full_test.sh"],
    capture_output=True,
    text=True,
    env={**os.environ, "TOML_TEST_VERSION": "v2.2.0"},
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC complete-conformance ===

## User Acceptance

- None.

## Guardrails

- Do not filter, skip, reinterpret, or replace any supplied conformance case.
- Do not modify the supplied harness files.
- Do not treat printed summary text as the release verdict; use the harness exit status.
