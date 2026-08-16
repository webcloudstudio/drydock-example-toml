# FEATURE: Complete TOML Conformance

| Field       | Value |
|-------------|-------|
| Version     | 20260816 V1 |
| Description | Assemble the completed decoder and verify the complete TOML 1.0.0 conformance suite. |
| Depends On  | ARCHITECTURE.md, FEATURE-Decoder-Contract.md, FEATURE-Strings.md, FEATURE-Lexical-Scalars.md, FEATURE-Lexical-Boundaries.md, FEATURE-Keys.md, FEATURE-Tables.md, FEATURE-Arrays-Inline.md, FEATURE-Datetimes.md |
| Provides    | complete TOML 1.0.0 conformance verdict |
| Consumes    | cmd/toml-decoder, sources/full_test.sh |

## Assembly

Use the completed `cmd/toml-decoder` and parser implementation with the supplied scoring assets unchanged. The final verification must execute the unfiltered `sources/full_test.sh` command from the application root.

## Release Verification

The complete authoritative TOML 1.0.0 suite is the sole project acceptance gate. Its exit status must be zero; do not substitute filtered runs or assert on summary text.

## Programmatic Acceptance

=== AC conformance-full ===
Intent: The completed decoder passes the supplied unfiltered TOML 1.0.0 conformance suite.
Suite: full
Requires: executable=go; scope=test

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

## User Acceptance

- None.

## Guardrails

- Run `sh sources/full_test.sh` without filtering, skipping, or reinterpretation.
- Preserve all supplied scoring and harness assets.
- Do not introduce a third-party TOML parser.
