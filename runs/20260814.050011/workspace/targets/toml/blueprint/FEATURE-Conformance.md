# FEATURE: Complete TOML Conformance

| Field       | Value |
|-------------|-------|
| Version     | 20260814 V1 |
| Description | Verifies the completed decoder against the pinned TOML 1.0.0 conformance suite. |
| Depends On  | FEATURE-Decoder-Invalid-Input.md, FEATURE-Tagged-JSON.md, FEATURE-Lexical-Strings.md, FEATURE-Lexical-Numbers.md, FEATURE-Lexical-Datetime.md, FEATURE-Lexical-Whitespace-Comments.md, FEATURE-Key-Forms.md, FEATURE-Key-Semantics.md, FEATURE-Arrays.md, FEATURE-Inline-Tables.md, FEATURE-Inline-Table-Closure.md, FEATURE-Standard-Tables.md, FEATURE-Table-Redefinition.md, FEATURE-Arrays-of-Tables.md, FEATURE-Arrays-of-Tables-Ordering.md |
| Provides    | complete TOML 1.0.0 conformance verdict |
| Consumes    | cmd/toml-decoder, TOML parser |

## Purpose

The completed Go decoder passes every valid and invalid case in the installed `toml-test` v2.2.0 TOML 1.0 suite.

## Verification

The supplied `sources/full_test.sh` builds `cmd/toml-decoder`, invokes the pinned harness without filters, and uses the harness exit status as the acceptance verdict. The supplied scoring assets remain unchanged.

## Programmatic Acceptance

Requires: executable=go; scope=test
Requires: executable=sh; scope=test

=== AC complete-conformance ===
Intent: The completed decoder passes the complete pinned TOML 1.0.0 conformance suite.
Suite: full

import os
import subprocess
import sys

result = subprocess.run(
    ["sh", "sources/full_test.sh"],
    capture_output=True,
    text=True,
    env={**os.environ, "TOML_TEST_VERSION": "v2.2.0"},
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC complete-conformance ===

## User Acceptance

- None.

## Guardrails

- The full suite runs unfiltered.
- The harness exit status is the sole acceptance verdict.
- The supplied scoring scripts are not modified.
