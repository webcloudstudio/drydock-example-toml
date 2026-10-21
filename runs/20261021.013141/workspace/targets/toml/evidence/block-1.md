# Evidence: Block 1 · Foundational (block-1)

- block type: block
- date: 2026-10-20
- resulting state: closed/verified
- story points (combined assembled cost): 7880
- execution id: 20261021.013514.533Z-e70ca1b7

## Stories built
- Establish the standard-library-only Go module and decoder command boundary. (architecture) [story]

## Reusable compacts
- ARCHITECTURE_compact.md

## Missing context
- context: INSTRUCTIONS.md

## Stacked context
- implements: go.md (SP 4743)
- compass: COMPASS.md (SP 521)
- implements: ARCHITECTURE.md (SP 728)
- stack: common.md (SP 1812)

## Build directory changes
- cmd/toml-decoder/main.go
- go.mod
- internal/toml/toml.go
- internal/toml/toml_test.go
- toml-decoder

## Pre-build acceptance observation
- RED: architecture-gate (ARCHITECTURE.md)
  intent: The Go module and decoder command satisfy the supplied architecture contract.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['sh', 'sources/stage_contract.sh', 'architecture'], returncode=1, stdout='', stderr='')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "architecture-gate.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: architecture-artifacts (ARCHITECTURE.md)
  intent: The foundation exposes the declared Go module and command package.
  return code: 1
  stderr:
    Traceback (most recent call last):
      File "architecture-artifacts.py", line 3, in <module>
        assert Path("go.mod").is_file()
               ^^^^^^^^^^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- PASS: architecture-gate (ARCHITECTURE.md)
  intent: The Go module and decoder command satisfy the supplied architecture contract.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: architecture-artifacts (ARCHITECTURE.md)
  intent: The foundation exposes the declared Go module and command package.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0

## Build summary
<reusable-compact filename="ARCHITECTURE.md">
- Go 1.22+ standard-library-only module.
- `cmd/toml-decoder`: stdin/stdout process boundary, diagnostics on stderr, non-zero on invalid TOML.
- `internal/toml`: TOML lexing, parsing, validation, document structure, and tagged values; no file, environment, or service access.
- No third-party dependencies or persistence.
- Required artifacts: `go.mod`, `cmd/toml-decoder/`, `internal/toml/`.
</reusable-compact>

RESULT: SUCCESS

FILES CHANGED:
- go.mod
- cmd/toml-decoder/main.go
- internal/toml/toml.go
- internal/toml/toml_test.go
- toml-decoder

SUMMARY:
Created the standard-library-only Go module and decoder/parser boundaries. `gofmt`, race tests, vet, builds, architecture contract, and decoder smoke tests pass.

BLOCKERS:
- None
