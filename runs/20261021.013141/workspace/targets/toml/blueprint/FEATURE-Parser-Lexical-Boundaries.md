# FEATURE: Parser Lexical Boundaries

| Field       | Value |
|-------------|-------|
| Version     | 20261021 V1 |
| Description | Enforces TOML comment, whitespace, UTF-8, and control-character boundaries. |
| Depends On  | FEATURE-Decoder-Contract.md |
| Provides    | TOML comment and encoding validation |
| Consumes    | stdin TOML decoder |

## Scope

Implement LF and CRLF line handling, permitted tab and space whitespace, full-line and trailing comments, and comment markers that occur inside strings.

The parser must accept only valid UTF-8 documents and reject prohibited control characters in comments and strings while permitting the controls explicitly allowed by TOML 1.0.0.

## Programmatic Acceptance

Requires: executable=sh; scope=test
Requires: executable=go; scope=test
Requires: executable=toml-test; scope=test

=== AC parser-lexical-boundaries-conformance ===
Intent: The decoder passes the authoritative comment, control-character, and encoding cases selected by `valid/comment/**|invalid/control/**|invalid/encoding/**`; the selector is expected to match the boundary corpus and the scoped run must execute successfully.

import os
import subprocess

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/comment/**|invalid/control/**|invalid/encoding/**"],
    capture_output=True,
    text=True,
    env={**os.environ, "TOML_TEST_VERSION": "v2.2.0"},
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC parser-lexical-boundaries-conformance ===

=== AC parser-lexical-boundaries-executable ===
Intent: The decoder executable remains available at the command boundary after lexical boundary support is added.

from pathlib import Path

assert Path("toml-decoder").is_file()
=== END AC parser-lexical-boundaries-executable ===

## User Acceptance

- None.

## Guardrails

- Reject invalid UTF-8 input.
- Reject prohibited control characters in comments and strings.
- Treat hash characters inside strings as string content.
- Do not alter supplied conformance harness files.
