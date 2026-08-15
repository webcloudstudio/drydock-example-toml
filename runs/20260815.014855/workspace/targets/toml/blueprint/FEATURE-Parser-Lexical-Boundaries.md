# FEATURE: Parser Lexical Boundaries

| Field       | Value |
|-------------|-------|
| Version     | 20260815 V1 |
| Description | Defines TOML comment, encoding, control-character, whitespace, and token-boundary validation. |
| Depends On  | FEATURE-Parser-Lexical-Scalars.md |
| Provides    | TOML comment handling, UTF-8 validation, control-character rejection, lexical termination |
| Consumes    | stdin TOML decoder, TOML string lexer, TOML scalar lexer |

## Behavior

The parser treats hashes as comments only outside strings, accepts TOML whitespace and LF or CRLF newlines, validates the complete document as UTF-8, rejects forbidden control characters in comments and values, and requires valid lexical termination between key/value pairs, headers, comments, and end of document.

## Programmatic Acceptance

=== AC parser-lexical-boundaries-conformance ===
Intent: The authoritative TOML conformance suite passes every valid comment and invalid control-character and encoding case owned by this story.
Suite: scoped
Requires: executable=sh; scope=test
Requires: executable=go; scope=test

import os
import subprocess

result = subprocess.run(
    [
        "sh",
        "sources/stage_test.sh",
        "valid/comment/**|invalid/control/**|invalid/encoding/**",
    ],
    capture_output=True,
    text=True,
    env={**os.environ},
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC parser-lexical-boundaries-conformance ===

## User Acceptance

- None.

## Guardrails

- Never treat a hash inside a string as a comment marker.
- Reject invalid UTF-8 and forbidden control characters.
- Reject adjacent or unterminated tokens rather than silently accepting trailing content.
