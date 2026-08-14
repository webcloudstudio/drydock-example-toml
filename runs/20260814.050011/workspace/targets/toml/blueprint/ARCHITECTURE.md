# ARCHITECTURE: TOML 1.0.0 Parser

| Field       | Value |
|-------------|-------|
| Version     | 20260814 V1 |
| Description | Defines the standard-library Go architecture for the TOML parser and decoder command. |
| Depends On  | — |
| Provides    | parser module boundary, decoder command boundary, tagged JSON contract |
| Consumes    | — |

## Intent

The application is a deterministic command-line filter that reads TOML 1.0.0 from standard input and emits tagged JSON on standard output. Invalid TOML produces a diagnostic on standard error and a non-zero exit status.

## Module Boundaries

| Module | Responsibility |
|---|---|
| `cmd/toml-decoder` | Reads stdin, invokes the parser, encodes successful results, and maps failures to process behavior. |
| `internal/toml` | Parses TOML 1.0.0, validates lexical and structural rules, and exposes the internal value model. |
| JSON encoder | Converts internal values to the required tagged JSON representation. |

The parser separates lexical scanning, structural parsing, semantic validation, value representation, and JSON encoding responsibilities. The internal value model uses explicit value kinds and preserves integer lexemes losslessly.

## Technology Stack

| Technology | Use |
|---|---|
| Go | Parser, command, tests, and standard-library integration. |
| Shell | Supplied conformance harness entry points. |

The Go module declares no external dependencies. Runtime behavior uses only the Go standard library.

## Process Contract

The decoder accepts no arguments or configuration. It reads the complete TOML document from stdin. Valid input produces one JSON document on stdout and exit status `0`. Invalid input produces no successful JSON result, writes diagnostics to stderr, and exits non-zero. The command has no filesystem or network side effects.

## Tagged JSON Contract

Tables are JSON objects and arrays are JSON arrays. Each scalar TOML value is represented as an object with `type` and `value` fields. The `value` field is always a JSON string. Supported scalar types are `string`, `integer`, `float`, `bool`, `datetime`, `datetime-local`, `date-local`, and `time-local`.

## Guardrails

- The implementation uses Go with no `go.mod` dependencies.
- The parser does not import a third-party TOML parser.
- The parser preserves signed 64-bit integer values without floating-point conversion.
- The decoder remains an argument-free stdin/stdout filter.
- Supplied scoring scripts remain unchanged.

## Programmatic Acceptance

=== AC architecture-module ===
Intent: The completed build contains the declared Go module and command boundary.
from pathlib import Path

module = Path("go.mod")
command = Path("cmd/toml-decoder")
parser = Path("internal/toml")
assert module.is_file()
assert command.is_dir()
assert parser.is_dir()
=== END AC architecture-module ===

=== AC architecture-dependencies ===
Intent: The Go module has no third-party dependency declarations.
from pathlib import Path

module_text = Path("go.mod").read_text(encoding="utf-8")
assert "require " not in module_text
assert "github.com/BurntSushi/toml" not in module_text
assert "github.com/pelletier/go-toml" not in module_text
=== END AC architecture-dependencies ===

=== AC architecture-command-build ===
Intent: The declared decoder command builds successfully.
import subprocess

result = subprocess.run(
    ["go", "build", "./cmd/toml-decoder"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC architecture-command-build ===

## User Acceptance

- None.

## Guardrails

- None.
