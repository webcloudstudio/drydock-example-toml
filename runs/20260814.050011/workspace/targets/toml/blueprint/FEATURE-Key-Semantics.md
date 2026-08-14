# FEATURE: Key Semantics

| Field       | Value |
|-------------|-------|
| Version     | 20260814 V1 |
| Description | Enforces TOML key uniqueness and scalar-to-table definition rules. |
| Depends On  | FEATURE-Key-Forms.md, FEATURE-Lexical-Numbers.md, FEATURE-Lexical-Datetime.md |
| Provides    | key-definition validation |
| Consumes    | TOML key paths, parser value model |

## Behavior

The parser records direct key definitions and implicitly created tables. It rejects duplicate keys, attempts to redefine equivalent bare and quoted keys, invalid key/value boundaries, and attempts to extend scalar values or otherwise closed structures as tables. Valid out-of-order dotted-key definitions remain accepted where TOML 1.0.0 permits them.

## Programmatic Acceptance

Requires: executable=go; scope=test  
Requires: executable=sh; scope=test

=== AC key-semantics-valid ===
Intent: The implementation passes the authoritative valid TOML key and specification conformance slices.
Suite: scoped

import os
import subprocess
import sys

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "valid/key/*"],
    capture_output=True,
    text=True,
    env={**os.environ, "TOML_TEST_VERSION": "v2.2.0"},
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC key-semantics-valid ===

=== AC key-semantics-invalid ===
Intent: The implementation rejects duplicate and structurally invalid key definitions in the authoritative suite.
Suite: scoped

import os
import subprocess
import sys

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "invalid/key/*"],
    capture_output=True,
    text=True,
    env={**os.environ, "TOML_TEST_VERSION": "v2.2.0"},
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC key-semantics-invalid ===

=== AC key-semantics-boundary ===
Intent: The implementation rejects invalid key/value boundaries in the authoritative specification suite.
Suite: scoped

import os
import subprocess
import sys

result = subprocess.run(
    ["sh", "sources/stage_test.sh", "invalid/spec-1.0.0/*"],
    capture_output=True,
    text=True,
    env={**os.environ, "TOML_TEST_VERSION": "v2.2.0"},
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC key-semantics-boundary ===

## User Acceptance

- None.

## Guardrails

- Equivalent key spellings cannot define the same key more than once.
- A scalar value cannot later become a table or parent of a table.
