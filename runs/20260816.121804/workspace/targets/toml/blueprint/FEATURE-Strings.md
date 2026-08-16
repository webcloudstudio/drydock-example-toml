# FEATURE: TOML Strings

| Field       | Value |
|-------------|-------|
| Version     | 20260816 V1 |
| Description | Defines TOML 1.0.0 basic, literal, multiline basic, and multiline literal string parsing. |
| Depends On  | ARCHITECTURE.md |
| Provides    | TOML string parsing |
| Consumes    | internal/toml |

## Scope

Support basic and literal strings, multiline delimiters, required escapes, Unicode scalar escapes, newline trimming, line-ending continuation, embedded quotes, preserved whitespace, and prohibited-control validation according to TOML 1.0.0.

## Programmatic Acceptance

=== AC strings-valid-suite ===
Intent: The authoritative conformance suite passes every valid string case owned by this feature.

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/string/**"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC strings-valid-suite ===

=== AC strings-invalid-suite ===
Intent: The authoritative conformance suite rejects every invalid string case owned by this feature.

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "invalid/string/**"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC strings-invalid-suite ===

=== AC strings-bounded-suites ===
Intent: Both string conformance invocations complete successfully as separate executable checks.

import subprocess

patterns = ["valid/string/**", "invalid/string/**"]
for pattern in patterns:
    result = subprocess.run(
        ["sh", "sources/stage_test.sh", pattern],
        capture_output=True,
        text=True,
    )
    print(result.stdout)
    print(result.stderr)
    assert result.returncode == 0
=== END AC strings-bounded-suites ===

## User Acceptance

- None.

## Guardrails

- Basic-string escapes must be decoded only when TOML 1.0.0 permits them.
- Literal strings must preserve backslashes and reject unescaped delimiters.
- Multiline strings must enforce delimiter, trimming, continuation, and control-character rules.
- Invalid UTF-8 and prohibited controls must not be accepted as string content.
