# ARCHITECTURE: TOML Parser

| Field       | Value |
|-------------|-------|
| Version     | 20260815 V1 |
| Description | Standard-library-only Go architecture for a TOML v1.0.0 stdin decoder. |
| Depends On  | — |
| Provides    | Go module, cmd/toml-decoder build entry point, internal/toml parser boundary |
| Consumes    | — |

## Intent

The application is a command-line filter that reads one UTF-8 TOML document from standard input and emits a tagged JSON document. Parsing is owned by `internal/toml`; the fixed executable entry point is `cmd/toml-decoder`.

## Module Boundaries

| Module | Responsibility |
|---|---|
| `cmd/toml-decoder` | Read stdin, invoke the parser, encode tagged JSON, report errors, and set exit status. |
| `internal/toml` | Lex and parse TOML v1.0.0 into the tagged document model. |
| `sources/` | Supplied conformance and acceptance harness assets; read-only during implementation. |

The command must not contain TOML grammar logic. The parser must not depend on command-line arguments, configuration, network access, or third-party packages.

## Tagged Document Model

Tables are JSON objects and arrays are JSON arrays. Every scalar TOML value is represented as an object with string fields:

- `type`: one of `string`, `integer`, `float`, `bool`, `datetime`, `datetime-local`, `date-local`, or `time-local`.
- `value`: the TOML value represented as a JSON string.

## Technology Stack

- Go 1.22 or newer, using only the standard library.
- POSIX shell for supplied acceptance and conformance scripts.

## Build and Runtime Constraints

The module must compile with no `require` dependencies in `go.mod`. The executable reads only stdin, writes successful JSON only to stdout, writes diagnostics only to stderr, and exits non-zero for malformed TOML.

## Programmatic Acceptance

=== AC architecture-module ===
Intent: The project has a standard-library-only Go module and the fixed decoder build entry point.

import pathlib
import subprocess

root = pathlib.Path(".")
go_mod = root / "go.mod"
command = root / "cmd" / "toml-decoder"
assert go_mod.is_file()
assert command.is_dir()
result = subprocess.run(
    ["sh", "sources/stage_contract.sh", "architecture"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC architecture-module ===

=== AC architecture-parser-boundary ===
Intent: The decoder command and internal parser boundary compile together.

import pathlib
import subprocess

parser = pathlib.Path("internal") / "toml"
command = pathlib.Path("cmd") / "toml-decoder"
assert parser.is_dir()
assert command.is_dir()
result = subprocess.run(
    ["go", "build", "./cmd/toml-decoder"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC architecture-parser-boundary ===

## User Acceptance

- None.

## Guardrails

- Do not add a third-party TOML parser or any other third-party Go dependency.
- Preserve `cmd/toml-decoder` as the build entry point.
- Do not modify supplied scoring or conformance scripts.
