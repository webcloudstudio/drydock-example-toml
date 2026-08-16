# FEATURE: Conformance

| Field       | Value |
|-------------|-------|
| Version     | 20260816 V1 |
| Description | Verifies complete TOML 1.0.0 conformance against the installed authoritative suite. |
| Depends On  | ARCHITECTURE.md, FEATURE-Decoder-Contract.md, FEATURE-Parser-Strings.md, FEATURE-Parser-Lexical-Scalars.md, FEATURE-Parser-Lexical-Boundaries.md, FEATURE-Parser-Keys.md, FEATURE-Parser-Tables.md, FEATURE-Parser-Arrays-Inline.md, FEATURE-Parser-Datetimes.md |
| Provides    | complete TOML 1.0.0 conformance verification |
| Consumes    | cmd/toml-decoder |

## Purpose

The terminal verification runs the supplied unfiltered `sources/full_test.sh` entry point against the installed `toml-test` v2.2.0 suite. The harness exit status is the complete release verdict.

## Verification

The supplied scripts remain unchanged. `sources/full_test.sh` builds the decoder and invokes `sources/run_conformance.sh`, which selects the required harness version and executes the complete TOML 1.0 suite.

## Programmatic Acceptance

Requires: executable=go; scope=test
Requires: executable=sh; scope=test
Requires: executable=toml-test; scope=test

=== AC conformance-full ===
Intent: The completed implementation passes the complete authoritative TOML 1.0.0 suite.
Suite: full

import subprocess

result = subprocess.run(
    ["sh", "sources/full_test.sh"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC conformance-full ===

=== AC conformance-harness-integrity ===
Intent: The supplied full-test entry point remains present and executable as a shell script.
import os
from pathlib import Path

harness = Path("sources/full_test.sh")
assert harness.is_file()
assert os.access(harness, os.R_OK)
=== END AC conformance-harness-integrity ===

=== AC conformance-decoder-build ===
Intent: The decoder builds successfully through the same Go command used by the scoring harness.
import subprocess

result = subprocess.run(
    ["go", "build", "-o", "toml-decoder", "./cmd/toml-decoder"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC conformance-decoder-build ===

## User Acceptance

- None.

## Guardrails

- The complete unfiltered suite is the sole release verdict.
- Supplied scoring and harness scripts must not be modified.
- The installed harness must be TOML Test v2.2.0 when its version is available.
- No valid or invalid case may be skipped, reinterpreted, or filtered during final verification.
