# Evidence: Block 1 · Foundational (block-1)

- block type: block
- date: 2026-08-15
- resulting state: closed/verified
- story points (combined assembled cost): 7225
- execution id: 20260816.005728.172Z-cb1d49ff

## Stories built
- Establish the standard-library-only Go module and decoder command. (architecture) [story]

## Acceptance tooling authorization
- ARCHITECTURE.md#architecture-module: executable=sh; scope=test; authorization=existing Target environment
- ARCHITECTURE.md#architecture-boundary: executable=go; scope=test; authorization=existing Target environment

## Reusable compacts
- ARCHITECTURE_compact.md

## Stacked context
- compass: COMPASS.md (SP 2899)
- implements: ARCHITECTURE.md (SP 623)
- stack: go.md (SP 3646)

## Build directory changes
- cmd/toml-decoder/main.go
- go.mod
- internal/toml/toml.go
- internal/toml/toml_test.go

## Pre-build acceptance observation
- RED: architecture-module (ARCHITECTURE.md)
  intent: The project has a compilable standard-library-only Go module and decoder command.
  return code: 1
  stderr:
    Traceback (most recent call last):
      File "architecture-module.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: architecture-boundary (ARCHITECTURE.md)
  intent: The decoder command is buildable at its declared path.
  return code: 1
  stdout:
    go: cannot find main module, but found .git/config in /mnt/c/Users/barlo/projects/drydock/uat/Toml
    	to create a module there, run:
    	cd ../../../.. && go mod init
  stderr:
    Traceback (most recent call last):
      File "architecture-boundary.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- PASS: architecture-module (ARCHITECTURE.md)
  intent: The project has a compilable standard-library-only Go module and decoder command.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
- PASS: architecture-boundary (ARCHITECTURE.md)
  intent: The decoder command is buildable at its declared path.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0

## Build summary
<reusable-compact filename="ARCHITECTURE.md">
Go module `github.com/owner/toml-decoder` using Go 1.22+ and no dependencies. `cmd/toml-decoder` reads stdin, calls `internal/toml.Decode`, emits tagged JSON to stdout, and reports errors to stderr with nonzero exit. `internal/toml` owns TOML parsing and tagged values: scalar values use string `type`/`value` fields; tables are objects and arrays are arrays.
</reusable-compact>

RESULT: SUCCESS

FILES CHANGED:
- go.mod
- cmd/toml-decoder/main.go
- internal/toml/toml.go
- internal/toml/toml_test.go

SUMMARY:
Created the standard-library-only Go module, decoder command boundary, tagged value model, and foundational tests. `go test`, `go vet`, `go build`, formatting checks, and the architecture acceptance contract pass.

BLOCKERS:
- None
