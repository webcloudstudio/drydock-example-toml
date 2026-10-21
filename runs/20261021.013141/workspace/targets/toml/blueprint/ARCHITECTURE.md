# ARCHITECTURE: TOML Decoder Architecture

| Field       | Value |
|-------------|-------|
| Version     | 20261021 V1 |
| Description | Standard-library-only Go architecture for the TOML 1.0.0 decoder. |
| Depends On  | — |
| Provides    | Go module, cmd/toml-decoder, internal/toml |
| Consumes    | — |

## Purpose

The project is a Go command-line filter that parses TOML 1.0.0 from standard input and emits the required tagged JSON representation to standard output.

## Modules and Boundaries

| Module | Responsibility |
|---|---|
| `cmd/toml-decoder` | Process startup, stdin reading, parser delegation, JSON encoding, diagnostics, and exit status. |
| `internal/toml` | TOML lexical analysis, parsing, document structure, semantic validation, and tagged values. |

The command package must not contain TOML parsing logic. The parser package must use only the Go standard library and must not read files, environment configuration, or external services.

## Interface

The executable accepts no required arguments. It reads the complete TOML document from stdin.

Valid input produces one JSON document on stdout and exits zero. Invalid input produces a diagnostic on stderr, no decoded JSON on stdout, and exits non-zero.

## Technology Stack

- Go 1.22 or newer, governed by `go.md`.
- POSIX shell support for the supplied harness, governed by `common.md`.
- Go module dependencies are restricted to the standard library; `go.mod` declares no `require` dependencies.

## Module Ownership

| Boundary | Owner | Allowed access |
|---|---|---|
| Process startup and streams | `cmd/toml-decoder` | `os`, `io`, `encoding/json`, and `internal/toml` |
| TOML parsing | `internal/toml` | Go standard-library parsing and validation packages |
| Persistence | None | The decoder has no persistent state. |
| Configuration | None | The decoder reads no configuration files or environment settings. |
| File store | None | The decoder reads only stdin. |
| External services | None | The decoder has no external service calls. |

## Programmatic Acceptance

=== AC architecture-gate ===
Intent: The Go module and decoder command satisfy the supplied architecture contract.

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_contract.sh", "architecture"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC architecture-gate ===

=== AC architecture-artifacts ===
Intent: The foundation exposes the declared Go module and command package.

from pathlib import Path

assert Path("go.mod").is_file()
assert Path("cmd/toml-decoder").is_dir()
assert Path("internal/toml").is_dir()
=== END AC architecture-artifacts ===

## User Acceptance

- None.

## Guardrails

- Do not add a third-party TOML parsing dependency.
- Do not move parser responsibilities into the command package.
- Do not modify supplied harness files.
