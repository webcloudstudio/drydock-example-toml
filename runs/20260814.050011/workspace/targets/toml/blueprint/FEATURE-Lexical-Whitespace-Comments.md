# FEATURE: Lexical Whitespace and Comments

| Field       | Value |
|-------------|-------|
| Version     | 20260814 V1 |
| Description | Processes TOML whitespace, comments, line endings, continuations, and encoding rules. |
| Depends On  | ARCHITECTURE.md |
| Provides    | TOML document lexical boundaries |
| Consumes    | parser module boundary |

## Behavior

The parser treats spaces and tabs as whitespace, supports LF and CRLF line endings, removes comments only outside strings, and handles multiline continuation rules. It accepts valid UTF-8 documents and rejects invalid UTF-8, prohibited control characters, malformed line boundaries, and invalid lexical continuation forms.

## Programmatic Acceptance

Requires: executable=go; scope=test  
Requires: executable=sh; scope=test

=== AC lexical-whitespace-comments-valid ===
Intent: The implementation passes the authoritative valid TOML comment conformance slice.
Suite: scoped

import os
import subprocess
import sys

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/comment/*"],
    capture_output=True,
    text=True,
    env={**os.environ, "TOML_TEST_VERSION": "v2.2.0"},
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC lexical-whitespace-comments-valid ===

=== AC lexical-whitespace-comments-control ===
Intent: The implementation rejects prohibited control characters in the authoritative suite.
Suite: scoped

import os
import subprocess
import sys

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "invalid/control/*"],
    capture_output=True,
    text=True,
    env={**os.environ, "TOML_TEST_VERSION": "v2.2.0"},
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC lexical-whitespace-comments-control ===

=== AC lexical-whitespace-comments-encoding ===
Intent: The implementation rejects invalid encodings in the authoritative suite.
Suite: scoped

import os
import subprocess
import sys

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "invalid/encoding/*"],
    capture_output=True,
    text=True,
    env={**os.environ, "TOML_TEST_VERSION": "v2.2.0"},
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC lexical-whitespace-comments-encoding ===

## User Acceptance

- None.

## Guardrails

- Comments never alter string contents.
- Invalid UTF-8 and prohibited control characters are rejected.
