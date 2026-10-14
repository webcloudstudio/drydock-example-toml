# FEATURE: Conformance

| Field       | Value |
|-------------|-------|
| Version     | 20261014 V1 |
| Description | The terminal assembly story: the complete unfiltered toml-test suite run through the supplied scoring entry point. |
| Depends On  | ARCHITECTURE.md, FEATURE-Decoder-Contract.md, FEATURE-Parser-Strings.md, FEATURE-Parser-Lexical-Scalars.md, FEATURE-Parser-Lexical-Boundaries.md, FEATURE-Parser-Keys.md, FEATURE-Parser-Tables.md, FEATURE-Parser-Arrays-Inline.md, FEATURE-Parser-Datetimes.md |
| Provides    | sh sources/full_test.sh verdict |
| Consumes    | toml-decoder stdin/stdout contract |

## Purpose

This is the only story that runs the complete suite. It runs `sh sources/full_test.sh` from the application root with no filter and takes its exit status as the verdict. The exit status is the whole verdict: the summary text and case counts are not asserted.

## Programmatic Acceptance

=== AC conformance-full-suite ===
Intent: The complete unfiltered conformance suite exits zero.
Suite: full
Requires: executable=sh; scope=test
Requires: executable=go; scope=test
Requires: executable=toml-test; scope=test

import subprocess
import sys

result = subprocess.run(
    ["sh", "sources/full_test.sh"],
    capture_output=True, text=True,
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC conformance-full-suite ===

=== AC conformance-no-third-party ===
Intent: go.mod declares no required module, so no third-party TOML module is present.
Requires: executable=go; scope=test

from pathlib import Path

text = Path("go.mod").read_text(encoding="utf-8")
assert "require" not in text
=== END AC conformance-no-third-party ===

## User Acceptance

- None.

## Guardrails

- The scoring assets under `sources/` are read-only and are never modified or filtered.
- The conformance verdict is the exit status of `sh sources/full_test.sh` and nothing else.
