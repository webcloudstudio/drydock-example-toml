# ARCHITECTURE: TOML Decoder

| Field       | Value |
|-------------|-------|
| Version     | 20261014 V1 |
| Description | Module layout, package boundaries, and technology stack for a standard-library TOML v1.0.0 decoder. |
| Depends On  | None |
| Provides    | cmd/toml-decoder, internal/toml |
| Consumes    | None |

## Purpose

A command-line filter. It reads a TOML document on stdin, decodes it with a hand-written parser in `internal/toml`, and writes tagged JSON to stdout. Invalid input produces a diagnostic on stderr and a non-zero exit. The program takes no arguments, reads no configuration, and has no side effects.

## Modules

| Path | Responsibility |
|---|---|
| `go.mod` | Module `github.com/owner/toml-decoder`. Declares no `require` directives. Go 1.22 or newer. |
| `cmd/toml-decoder` | `package main`. Reads stdin, calls `internal/toml`, encodes the result to stdout, and writes errors to stderr. |
| `internal/toml` | Hand-written TOML v1.0.0 lexer, parser, and tagged-value model. Not importable outside the module. |

## Technology Stack

- Go, standard library only. Imports limited to `encoding/json`, `strconv`, `strings`, `unicode`, `unicode/utf8`, `time`, `math`, `fmt`, `io`, `os`, `errors`.
- No third-party TOML module, including `BurntSushi/toml`, `pelletier/go-toml`, and `naoina/toml`.
- The conformance harness is shell and Go tooling staged under `sources/`. It is scoring instrumentation, not part of the deliverable.

## Module Ownership

| Boundary | Owner | May access |
|---|---|---|
| stdin, stdout, stderr, exit status | `cmd/toml-decoder` | `os`, `io`, `encoding/json` |
| TOML syntax and semantics | `internal/toml` | `strings`, `strconv`, `unicode`, `unicode/utf8`, `time`, `math` |
| Persistence, configuration, external services | None | Not applicable |

## Programmatic Acceptance

=== AC architecture-go-mod ===
Intent: go.mod declares no require directive and names a module.
Requires: executable=go; scope=test

from pathlib import Path

text = Path("go.mod").read_text(encoding="utf-8")
assert "require" not in text
assert "module" in text
=== END AC architecture-go-mod ===

=== AC architecture-build ===
Intent: The decoder command builds using only the standard library.
Requires: executable=go; scope=test

import subprocess
import sys

build = subprocess.run(
    ["go", "build", "-o", "toml-decoder", "./cmd/toml-decoder"],
    capture_output=True, text=True,
)
print(build.stdout)
print(build.stderr, file=sys.stderr)
assert build.returncode == 0
=== END AC architecture-build ===

=== AC architecture-entry-point ===
Intent: The internal parser package is present and the entry point is a main package.
Requires: executable=go; scope=test

from pathlib import Path

entry = Path("cmd/toml-decoder/main.go").read_text(encoding="utf-8")
assert "package main" in entry
assert Path("internal/toml").is_dir()
=== END AC architecture-entry-point ===

## User Acceptance

- None.

## Guardrails

- Never import a third-party TOML module.
- `go.mod` declares no `require` directive.
- The scoring assets under `sources/` are read-only and are never modified.
