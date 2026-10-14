# Evidence: Block 1 · Foundational (block-1)

- block type: block
- date: 2026-10-14
- resulting state: closed/verified
- story points (combined assembled cost): 10953
- execution id: 20261014.125538.882Z-7dab416f

## Stories built
- Standard-library Go module and decoder command boundary. (architecture) [story]

## Acceptance tooling authorization
- ARCHITECTURE.md#architecture-go-mod: executable=go; scope=test; authorization=existing Target environment
- ARCHITECTURE.md#architecture-build: executable=go; scope=test; authorization=existing Target environment
- ARCHITECTURE.md#architecture-entry-point: executable=go; scope=test; authorization=existing Target environment

## Reusable compacts
- ARCHITECTURE_compact.md

## Stacked context
- compass: COMPASS.md (SP 945)
- implements: ARCHITECTURE.md (SP 1196)
- stack: go.md (SP 5918)
- stack: common.md (SP 2769)

## Build directory changes
- cmd/toml-decoder/main.go
- cmd/toml-decoder/main_test.go
- go.mod
- internal/toml/toml.go
- internal/toml/toml_test.go

## Pre-build acceptance observation
- RED: architecture-go-mod (ARCHITECTURE.md)
  intent: go.mod declares no require directive and names a module.
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
      File "architecture-go-mod.py", line 3, in <module>
        text = Path("go.mod").read_text(encoding="utf-8")
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1027, in read_text
        with self.open(mode='r', encoding=encoding, errors=errors) as f:
             ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1013, in open
        return io.open(self, mode, buffering, encoding, errors, newline)
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    FileNotFoundError: [Errno 2] No such file or directory: 'go.mod'
- RED: architecture-build (ARCHITECTURE.md)
  intent: The decoder command builds using only the standard library.
  return code: 1
  stderr:
    go: cannot find main module, but found .git/config in /mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20261014.124728/build/toml
    	to create a module there, run:
    	go mod init
    
    --- drydock: values at failure ---
      build = CompletedProcess(args=['go', 'build', '-o', 'toml-decoder', './cmd/toml-decoder'], returncode=1, stdout='', stderr='go: cannot find main module, but found .git/config in /mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20261014.124728/build/toml\n\tto create a module there, run:\n\tgo mod init\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "architecture-build.py", line 10, in <module>
        assert build.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: architecture-entry-point (ARCHITECTURE.md)
  intent: The internal parser package is present and the entry point is a main package.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      self = PosixPath('cmd/toml-decoder/main.go')
      mode = 'r'
      buffering = -1
      encoding = 'utf-8'
      errors = None
      newline = None
    --- drydock: end values ---
    Traceback (most recent call last):
      File "architecture-entry-point.py", line 3, in <module>
        entry = Path("cmd/toml-decoder/main.go").read_text(encoding="utf-8")
                ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1027, in read_text
        with self.open(mode='r', encoding=encoding, errors=errors) as f:
             ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1013, in open
        return io.open(self, mode, buffering, encoding, errors, newline)
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    FileNotFoundError: [Errno 2] No such file or directory: 'cmd/toml-decoder/main.go'

## Post-build programmatic acceptance
- PASS: architecture-go-mod (ARCHITECTURE.md)
  intent: go.mod declares no require directive and names a module.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: architecture-build (ARCHITECTURE.md)
  intent: The decoder command builds using only the standard library.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: architecture-entry-point (ARCHITECTURE.md)
  intent: The internal parser package is present and the entry point is a main package.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0

## Build summary
The build directory holds only the staged `sources/` and a git repo, with Go 1.24.6 available. I'll create the module, the command, and the stub package, then write tests that build and run the binary and read back its output.Next, the tests. Each one acts on the system and reads the state back. Binary builds go to `t.TempDir()` so nothing is left in the build directory.Everything passes. Source tree is untouched (`sources/` does not appear in status).

<reusable-compact filename="ARCHITECTURE.md">
Module: github.com/owner/toml-decoder (go 1.22, no require directives, stdlib only).
Packages:
- cmd/toml-decoder (package main): reads stdin, calls toml.Decode, JSON-encodes result to stdout; errors to stderr, exit 1. No args, no config, no side effects.
- internal/toml: Decode(r io.Reader) (map[string]any, error); ErrNotImplemented sentinel (stub in this block).
Constraints: no third-party TOML modules; malformed input exits non-zero; scoring assets under sources/ are read-only.
</reusable-compact>

<blueprint-decisions>
[{"spec":"ARCHITECTURE.md","severity":"Low","subject":"Stub Decode signature","decision":"Decode returns (map[string]any, error) and always returns ErrNotImplemented, so the command's JSON encoding path is wired and testable before the parser exists. The tagged-value model will replace map[string]any in a later block. Is that acceptable, or should the stub use a different shape?"}]
</blueprint-decisions>

RESULT: SUCCESS

FILES CHANGED:
- go.mod
- cmd/toml-decoder/main.go
- cmd/toml-decoder/main_test.go
- internal/toml/toml.go
- internal/toml/toml_test.go

SUMMARY:
- `go.mod` declares `github.com/owner/toml-decoder` with `go 1.22` and no `require` directives.
- `cmd/toml-decoder/main.go` is a `package main` that reads stdin, calls `toml.Decode`, JSON-encodes the result to stdout, and writes errors to stderr with exit 1.
- `internal/toml/toml.go` is a stub `Decode` returning `ErrNotImplemented`, so the module builds on the standard library alone.
- Tests cover the three declared acceptance assertions: go.mod has no `require` and names the module, `main.go` is `package main` and `internal/toml` exists, and `go build ./cmd/toml-decoder` succeeds. Extra tests check the stub's error path, `run`'s stdout on failure, and that the built binary exits non-zero with empty stdout and a stderr diagnostic. Binaries are built into `t.TempDir()`, so nothing is left in the build directory.
- Verification: `gofmt -l .` prints nothing, `go vet` and `go build ./...` pass, and `go test -race -cover ./...` passes (all 5 command tests and 1 stub test green; coverage 42.9% for the command, 100% for `internal/toml`).
- I did not run the staged `sources/` conformance scripts; this block's criteria do not invoke them, and `sources/` is unmodified.
- The Python acceptance snippets from ARCHITECTURE.md were not executed literally. Their equivalent checks are in the Go tests above.

BLOCKERS:
- None.
