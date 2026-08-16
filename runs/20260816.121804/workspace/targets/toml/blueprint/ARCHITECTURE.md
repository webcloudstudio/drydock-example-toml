# ARCHITECTURE: TOML 1.0.0 Parser

| Field       | Value |
|-------------|-------|
| Version     | 20260816 V1 |
| Description | Standard-library-only Go architecture for a TOML 1.0.0 command-line decoder. |
| Depends On  | — |
| Provides    | cmd/toml-decoder, internal/toml, Go module |
| Consumes    | — |

## Intent

The project is a Go command-line filter that reads a TOML 1.0.0 document from standard input and emits deterministic tagged JSON. Parsing is implemented in `internal/toml`, while `cmd/toml-decoder` owns process I/O, JSON encoding, diagnostics, and exit status.

## Modules and Boundaries

| Module | Responsibility |
|---|---|
| `cmd/toml-decoder` | Read stdin, invoke the parser, encode successful results, and report failures. |
| `internal/toml` | Parse TOML 1.0.0 syntax and semantics into tagged values and tables. |
| `go.mod` | Declare the Go module and toolchain floor without third-party dependencies. |

The command must not read configuration files, access the network, persist data, or modify supplied scoring assets. The parser must use only the Go standard library.

## Tagged Value Model

Tables map to JSON objects and arrays map to JSON arrays. Every scalar value is represented as an object containing string-valued `type` and `value` fields. Supported type tags are `string`, `integer`, `float`, `bool`, `datetime`, `datetime-local`, `date-local`, and `time-local`.

## Technology Stack

- Go 1.22 or newer, using `go.md`.
- POSIX shell for supplied acceptance scripts, using `common.md`.

## Programmatic Acceptance

=== AC architecture-build ===
Intent: The declared Go module and decoder command compile without third-party module requirements.

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_contract.sh", "architecture"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC architecture-build ===

=== AC architecture-module ===
Intent: The command package is buildable through the required repository path.

import subprocess

result = subprocess.run(
    ["go", "build", "./cmd/toml-decoder"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC architecture-module ===

=== AC architecture-dependencies ===
Intent: The module declares no forbidden third-party dependency.

from pathlib import Path

module_text = Path("go.mod").read_text(encoding="utf-8")
require_lines = [
    line for line in module_text.splitlines()
    if line.lstrip().startswith("require ")
]
assert not require_lines
=== END AC architecture-dependencies ===

## User Acceptance

- None.

## Guardrails

- The parser must not import a third-party TOML library.
- The `cmd/toml-decoder` build path must remain unchanged.
- Supplied scoring scripts and harness assets must not be modified.
