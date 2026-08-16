# Evidence: Block 1 · Foundational (block-1)

- block type: block
- date: 2026-08-16
- resulting state: closed/verified
- story points (combined assembled cost): 9153
- execution id: 20260816.122441.637Z-c478c6cc

## Stories built
- Establish the standard-library-only Go module and decoder command. (architecture) [story]

## Reusable compacts
- ARCHITECTURE_compact.md

## Stacked context
- compass: COMPASS.md (SP 2942)
- implements: ARCHITECTURE.md (SP 721)
- stack: go.md (SP 3646)
- stack: common.md (SP 1807)

## Build directory changes
- cmd/toml-decoder/main.go
- go.mod
- internal/toml/toml.go
- internal/toml/toml_test.go
- toml-decoder

## Pre-build acceptance observation
- RED: architecture-build (ARCHITECTURE.md)
  intent: The declared Go module and decoder command compile without third-party module requirements.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['sh', 'sources/stage_contract.sh', 'architecture'], returncode=1, stdout='', stderr='')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "architecture-build.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: architecture-module (ARCHITECTURE.md)
  intent: The command package is buildable through the required repository path.
  return code: 1
  stdout:
    go: cannot find main module, but found .git/config in /mnt/c/Users/barlo/projects/drydock/uat/Toml
    	to create a module there, run:
    	cd ../../../.. && go mod init
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['go', 'build', './cmd/toml-decoder'], returncode=1, stdout='', stderr='go: cannot find main module, but found .git/config in /mnt/c/Users/barlo/projects/drydock/uat/Toml\n\tto create a module there, run:\n\tcd ../../../.. && go mod init\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "architecture-module.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: architecture-dependencies (ARCHITECTURE.md)
  intent: The module declares no forbidden third-party dependency.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      self = PosixPath('go.mod')
      mode = 'r'
      buffering = -1
      encoding = 'utf-8'
      errors = None
      newline = None
    --- drydock: end values ---
    Traceback (most recent call last):
      File "architecture-dependencies.py", line 3, in <module>
        module_text = Path("go.mod").read_text(encoding="utf-8")
                      ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1027, in read_text
        with self.open(mode='r', encoding=encoding, errors=errors) as f:
             ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1013, in open
        return io.open(self, mode, buffering, encoding, errors, newline)
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    FileNotFoundError: [Errno 2] No such file or directory: 'go.mod'

## Post-build programmatic acceptance
- PASS: architecture-build (ARCHITECTURE.md)
  intent: The declared Go module and decoder command compile without third-party module requirements.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
- PASS: architecture-module (ARCHITECTURE.md)
  intent: The command package is buildable through the required repository path.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
- PASS: architecture-dependencies (ARCHITECTURE.md)
  intent: The module declares no forbidden third-party dependency.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0

## Build summary
<reusable-compact filename="ARCHITECTURE.md">
Go module `github.com/owner/toml-decoder` using Go 1.22+ and no dependencies. `cmd/toml-decoder` reads stdin, invokes `internal/toml`, emits tagged JSON on stdout, and reports errors on stderr with non-zero exit. `internal/toml` owns TOML parsing. Scalars use `{type, value}` string fields; tables and arrays map to JSON objects and arrays.
</reusable-compact>

RESULT: SUCCESS

FILES CHANGED:
- go.mod
- cmd/toml-decoder/main.go
- internal/toml/toml.go
- internal/toml/toml_test.go

SUMMARY:
Created the standard-library-only Go module and decoder command scaffold. `go test`, `go vet`, `go build`, formatting checks, and the architecture contract all pass.

BLOCKERS:
- None
