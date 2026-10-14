# Evidence: Block 8 · Service (block-8)

- block type: block
- date: 2026-10-14
- resulting state: closed/verified
- story points (combined assembled cost): 9945
- execution id: 20261014.132940.698Z-0a155b5f

## Stories built
- Stdin, tagged JSON, diagnostic-stream, and exit-status contract. (decoder-contract) [story]

## Acceptance tooling authorization
- FEATURE-Decoder-Contract.md#decoder-stage-gate: executable=sh; scope=test; authorization=existing Target environment
- FEATURE-Decoder-Contract.md#decoder-stage-gate: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Decoder-Contract.md#decoder-stage-gate: executable=python; scope=test; authorization=existing Target environment
- FEATURE-Decoder-Contract.md#decoder-scalar-round-trip: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Decoder-Contract.md#decoder-table-round-trip: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Decoder-Contract.md#decoder-invalid-exit: executable=go; scope=test; authorization=existing Target environment

## Stacked context
- compass: COMPASS.md (SP 945)
- implements: FEATURE-Decoder-Contract.md (SP 1521)
- context: ARCHITECTURE_compact.md (SP 244)
- stack: go.md (SP 5918)
- stack: common_compact.md (SP 1199)

## Build directory changes
- cmd/toml-decoder/contract_test.go
- cmd/toml-decoder/main.go
- toml-decoder

## Pre-build acceptance observation
- GREEN (prepassed): decoder-stage-gate (FEATURE-Decoder-Contract.md)
  intent: The governed decoder gate passes: the decoder decodes a scalar and rejects an unterminated array.
  return code: 0
- GREEN (prepassed): decoder-scalar-round-trip (FEATURE-Decoder-Contract.md)
  intent: A scalar integer is written to stdout as a tagged value and exit status is zero.
  return code: 0
- GREEN (prepassed): decoder-table-round-trip (FEATURE-Decoder-Contract.md)
  intent: A standard table encodes as a nested JSON object with tagged values.
  return code: 0
- GREEN (prepassed): decoder-invalid-exit (FEATURE-Decoder-Contract.md)
  intent: Malformed TOML exits non-zero, writes nothing to stdout, and writes a diagnostic to stderr.
  return code: 0
  stderr:
    toml-decoder: decode: invalid TOML: line 2: unterminated array

## Post-build programmatic acceptance
- PASS: decoder-stage-gate (FEATURE-Decoder-Contract.md)
  intent: The governed decoder gate passes: the decoder decodes a scalar and rejects an unterminated array.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: decoder-scalar-round-trip (FEATURE-Decoder-Contract.md)
  intent: A scalar integer is written to stdout as a tagged value and exit status is zero.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: decoder-table-round-trip (FEATURE-Decoder-Contract.md)
  intent: A standard table encodes as a nested JSON object with tagged values.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: decoder-invalid-exit (FEATURE-Decoder-Contract.md)
  intent: Malformed TOML exits non-zero, writes nothing to stdout, and writes a diagnostic to stderr.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    toml-decoder: decode: invalid TOML: line 2: unterminated array

## Build summary
I'll start by inspecting the build directory and the staged sources.The directory already has a built parser and a command package. Next I'll read the current entry point and the `toml` package API.The stdin-to-JSON path already exists in `main.go`. Before deciding what's missing, I'll run the existing test suite and the declared decoder gate to see where things stand.Tests, vet, gofmt, and the stage gate are all clean. The spec's encoding rules go further than the existing tests check, so I'll probe the built binary directly against them.The contract mostly holds already: tables, arrays, empty `{}`/`[]`, and the invalid-input exit path all behave as specified. Two things stand out. The entry point silently ignores command-line arguments, but the spec says it accepts none. Datetime values are rejected, which belongs to a scalar feature outside this block. I'll check the existing command tests for style and whether any cover arguments.Nothing in the command checks arguments, and no test covers that contract. I'll make the one code change this block needs: reject any argument with a usage diagnostic and exit 2, while keeping `run` unchanged. Then I'll add tests for the declared stdin, stdout, stderr, and exit behavior.Next, the tests. They act on the built binary and read back stdout as parsed JSON, not as scraped text.Now I'll run gofmt, vet, and the full test suite with the race detector, then run each declared acceptance assertion as written.Everything passes. I'll confirm the new tests actually ran and that only the intended files changed, with `sources/` untouched.The decoder contract is in place and every declared acceptance assertion passes. The one code change adds the "accepts no arguments" rule, which the entry point did not enforce.

<blueprint-decisions>
[{"spec":"FEATURE-Decoder-Contract.md","severity":"Low","subject":"Exit status for command-line arguments","decision":"The spec requires that the decoder accept no arguments but does not name an exit code for rejecting them. I exit 2, the conventional usage-error code, with a diagnostic on stderr and nothing on stdout. Decode failures keep exit 1. Is 2 acceptable, or should arguments use 1 like other failures?"}]
</blueprint-decisions>

RESULT: SUCCESS

FILES CHANGED:
- cmd/toml-decoder/main.go
- cmd/toml-decoder/contract_test.go
- toml-decoder (rebuilt binary; the acceptance scripts rebuild it, and it was already modified in the working tree before this step)

SUMMARY:
- **Argument rejection (new):** `main` now calls `checkNoArgs(os.Args[1:])` before decoding. Any argument prints `toml-decoder: takes no arguments, got [...]` to stderr and exits 2. Before this change, arguments were silently ignored.
- **Contract tests (new):** `contract_test.go` drives the built binary and reads stdout back as parsed JSON. It covers: an empty document gives `{}`; a table gives a nested tagged object; an array gives an array of tagged values; an empty array gives `[]`; a successful run writes nothing to stderr; and arguments give a non-zero exit, empty stdout, and a stderr diagnostic. It also has unit tests for `checkNoArgs`.
- **Acceptance assertions, all passing:**
  - `decoder-stage-gate` (`sh sources/stage_contract.sh decoder`): exit 0.
  - `decoder-scalar-round-trip`: pass.
  - `decoder-table-round-trip`: pass.
  - `decoder-invalid-exit`: pass.
- **Project checks:** `gofmt -l .` is empty, `go vet ./...` is clean, and `go test -race -cover ./...` passes in both packages (cmd 53.8%, internal/toml 97.0%). Coverage in `cmd` dropped from 57.1% because the new `main` lines are uncovered; `main` itself is not exercised by tests.
- **Suite runs:** I did not run the staged `toml-test` suite. This block's declared criteria don't invoke it, so I left it alone.
- **Not in this block:** datetime values (`1979-05-27T07:32:00Z`) are rejected with a syntax error. The spec lists datetime as a valid tagged type, but datetime scalar parsing belongs to a later feature, so I did not implement it here.
- **Git:** nothing committed or staged. `sources/` is unmodified.

BLOCKERS:
- None.
