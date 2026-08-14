# Evidence: Block 1 · Foundational (block-1)

- block type: block
- date: 2026-08-14
- resulting state: closed/verified
- story points (combined assembled cost): 9050
- execution id: 20260814.050854.031Z-9dfacef7

## Stories built
- Define the Go parser modules, command boundary, tagged JSON contract, and standard-library constraints. (architecture) [story]

## Reusable compacts
- ARCHITECTURE_compact.md

## Stacked context
- compass: COMPASS.md (SP 2629)
- implements: ARCHITECTURE.md (SP 905)
- stack: go.md (SP 3646)
- stack: common.md (SP 1807)

## Build directory changes
- cmd/toml-decoder/main.go
- go.mod
- internal/toml/decode.go
- internal/toml/value.go
- internal/toml/value_test.go

## Pre-build acceptance observation
- RED: architecture-module (ARCHITECTURE.md)
  intent: The completed build contains the declared Go module and command boundary.
  return code: 1
  stderr:
    Traceback (most recent call last):
      File "architecture-module.py", line 6, in <module>
        assert module.is_file()
               ^^^^^^^^^^^^^^^^
    AssertionError
- RED: architecture-dependencies (ARCHITECTURE.md)
  intent: The Go module has no third-party dependency declarations.
  return code: 1
  stderr:
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
- RED: architecture-command-build (ARCHITECTURE.md)
  intent: The declared decoder command builds successfully.
  return code: 1
  stdout:
    go: cannot find main module, but found .git/config in /mnt/c/Users/barlo/projects/drydock
    	to create a module there, run:
    	cd ../../../../../.. && go mod init
  stderr:
    Traceback (most recent call last):
      File "architecture-command-build.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- PASS: architecture-module (ARCHITECTURE.md)
  intent: The completed build contains the declared Go module and command boundary.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
- PASS: architecture-dependencies (ARCHITECTURE.md)
  intent: The Go module has no third-party dependency declarations.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
- PASS: architecture-command-build (ARCHITECTURE.md)
  intent: The declared decoder command builds successfully.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0

## Build summary
<reusable-compact filename="ARCHITECTURE.md">
Go module `github.com/owner/toml-decoder` uses standard library only. `cmd/toml-decoder` reads TOML from stdin, emits tagged JSON to stdout, reports invalid input on stderr, and exits non-zero. `internal/toml` owns parsing and explicit value kinds while preserving integer lexemes.
</reusable-compact>

RESULT: SUCCESS

FILES CHANGED:
- go.mod
- cmd/toml-decoder/main.go
- internal/toml/value.go
- internal/toml/decode.go
- internal/toml/value_test.go

SUMMARY:
Implemented the dependency-free Go module, parser boundary, decoder command, tagged JSON model, and focused tests. Go tests, vet, formatting, build, and all architecture acceptance checks pass.

BLOCKERS:
- None
