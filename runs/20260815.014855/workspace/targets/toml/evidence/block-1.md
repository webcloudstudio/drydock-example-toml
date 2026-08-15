# Evidence: Block 1 · Foundational (block-1)

- block type: block
- date: 2026-08-14
- resulting state: closed/verified
- story points (combined assembled cost): 9535
- execution id: 20260815.015512.078Z-a4d1b111

## Stories built
- Establish the standard-library-only Go module and decoder command. (architecture) [story]

## Reusable compacts
- ARCHITECTURE_compact.md
- stage_contract_compact.md

## Stacked context
- compass: COMPASS.md (SP 2932)
- implements: ARCHITECTURE.md (SP 778)
- context: stage_contract.sh (SP 317)
- stack: go.md (SP 3646)
- stack: common.md (SP 1807)

## Build directory changes
- cmd/toml-decoder/main.go
- go.mod
- internal/toml/toml.go
- internal/toml/toml_test.go
- toml-decoder

## Pre-build acceptance observation
- RED: architecture-module (ARCHITECTURE.md)
  intent: The project has a standard-library-only Go module and the fixed decoder build entry point.
  return code: 1
  stderr:
    Traceback (most recent call last):
      File "architecture-module.py", line 7, in <module>
        assert go_mod.is_file()
               ^^^^^^^^^^^^^^^^
    AssertionError
- RED: architecture-parser-boundary (ARCHITECTURE.md)
  intent: The decoder command and internal parser boundary compile together.
  return code: 1
  stderr:
    Traceback (most recent call last):
      File "architecture-parser-boundary.py", line 6, in <module>
        assert parser.is_dir()
               ^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- PASS: architecture-module (ARCHITECTURE.md)
  intent: The project has a standard-library-only Go module and the fixed decoder build entry point.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
- PASS: architecture-parser-boundary (ARCHITECTURE.md)
  intent: The decoder command and internal parser boundary compile together.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0

## Build summary
<reusable-compact filename="ARCHITECTURE.md">
Go module `github.com/owner/toml-decoder`, Go 1.22+, standard library only. `cmd/toml-decoder` reads stdin, calls `internal/toml.Decode`, emits tagged JSON to stdout, and reports parse errors on stderr with non-zero exit. Tables/arrays map to JSON containers; scalar values use `{type, value}` string fields. No third-party dependencies.
</reusable-compact>

<reusable-compact filename="stage_contract.sh">
`stage_contract.sh architecture` requires dependency-free `go.mod` and successful decoder compilation. `stage_contract.sh decoder` verifies `answer = 42` produces tagged JSON, malformed TOML exits non-zero, stdout stays empty on failure, and stderr is non-empty.
</reusable-compact>

RESULT: SUCCESS

FILES CHANGED:
- go.mod
- cmd/toml-decoder/main.go
- internal/toml/toml.go
- internal/toml/toml_test.go
- toml-decoder

SUMMARY:
Established the standard-library-only Go module, decoder entry point, internal parser boundary, tagged value model, initial parsing behavior, and deterministic tests. Architecture acceptance, decoder smoke contract, formatting, vet, and race tests pass.

BLOCKERS:
- None
