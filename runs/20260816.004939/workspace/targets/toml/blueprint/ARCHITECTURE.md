# ARCHITECTURE: TOML 1.0.0 Parser

| Field       | Value |
|-------------|-------|
| Version     | 20260816 V1 |
| Description | Defines the standard-library-only Go modules and command boundary for the TOML decoder. |
| Depends On  | — |
| Provides    | Go module, cmd/toml-decoder, internal/toml.Decode |
| Consumes    | — |

## Intent

The project is a Go command-line filter that parses TOML 1.0.0 from standard input and emits the required tagged JSON representation. The implementation is self-contained and uses only the Go standard library.

## Modules and Boundaries

| Module | Responsibility |
|---|---|
| `cmd/toml-decoder` | Process stdin, invoke the parser, encode successful results as JSON, and report failures. |
| `internal/toml` | Parse TOML 1.0.0 and construct tagged values and table/array structures. |

The command accepts no arguments, configuration, network access, or persistent state. Invalid input must not produce a successful stdout payload.

## Tagged Value Model

Scalar values are represented as objects with string-valued `type` and `value` fields. Tables are JSON objects and arrays are JSON arrays. Supported scalar types are `string`, `integer`, `float`, `bool`, `datetime`, `datetime-local`, `date-local`, and `time-local`.

## Technology Stack

- Go 1.22 or newer, with `go.md` governing implementation.
- Shell harnesses, with `common.md` governing supplied verification scripts.
- No third-party Go modules.

## Programmatic Acceptance

=== AC architecture-module ===
Intent: The project has a compilable standard-library-only Go module and decoder command.
Requires: executable=sh; scope=test

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_contract.sh", "architecture"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC architecture-module ===

=== AC architecture-boundary ===
Intent: The decoder command is buildable at its declared path.
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(
    ["go", "build", "./cmd/toml-decoder"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC architecture-boundary ===

## User Acceptance

- None.

## Guardrails

- The Go module must declare no third-party dependencies.
- The parser must not import an external TOML implementation.
- The supplied conformance and scoring scripts remain unchanged.
