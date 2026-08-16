# FEATURE: Lexical Boundaries

| Field       | Value |
|-------------|-------|
| Version     | 20260816 V1 |
| Description | Defines TOML comment, whitespace, line-ending, encoding, and control-character handling. |
| Depends On  | ARCHITECTURE.md |
| Provides    | TOML lexical boundary validation |
| Consumes    | internal/toml |

## Scope

Handle spaces and tabs, LF and CRLF line endings, comments outside strings, multiline lexical boundaries, valid UTF-8 documents, and rejection of prohibited control characters in comments and applicable value contexts.

## Programmatic Acceptance

=== AC lexical-comments-valid-suite ===
Intent: The authoritative conformance suite passes every valid comment and lexical-boundary case.

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/comment/**"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC lexical-comments-valid-suite ===

=== AC lexical-boundaries-invalid-suite ===
Intent: The authoritative conformance suite rejects every invalid control-character and encoding case.

import subprocess

result = subprocess.run(
    [
        "sh",
        "sources/stage_test.sh",
        "invalid/control/**|invalid/encoding/**",
    ],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC lexical-boundaries-invalid-suite ===

=== AC lexical-boundaries-suites ===
Intent: The valid comment and invalid lexical-boundary conformance slices complete successfully.

import subprocess

patterns = [
    "valid/comment/**",
    "invalid/control/**|invalid/encoding/**",
]
for pattern in patterns:
    result = subprocess.run(
        ["sh", "sources/stage_test.sh", pattern],
        capture_output=True,
        text=True,
    )
    print(result.stdout)
    print(result.stderr)
    assert result.returncode == 0
=== END AC lexical-boundaries-suites ===

## User Acceptance

- None.

## Guardrails

- Comments must not consume content inside strings.
- Only TOML-permitted whitespace and line endings may be accepted.
- Invalid UTF-8 must be rejected.
- Prohibited control characters must be rejected in every applicable lexical context.
