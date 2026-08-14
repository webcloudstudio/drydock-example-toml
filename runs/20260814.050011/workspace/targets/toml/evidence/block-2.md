# Evidence: Block 2 · Service (block-2)

- block type: block
- date: 2026-08-14
- resulting state: closed/verified
- story points (combined assembled cost): 8233
- execution id: 20260814.051109.985Z-4688b3f9

## Stories built
- Build the stdin-to-tagged-JSON decoder command. (decoder-contract) [story]

## Stacked context
- compass: COMPASS.md (SP 2629)
- implements: FEATURE-Decoder-Contract.md (SP 582)
- context: ARCHITECTURE_compact.md (SP 108)
- stack: go.md (SP 3646)
- stack: common_compact.md (SP 1179)

## Build directory changes
- cmd/toml-decoder/main.go
- cmd/toml-decoder/main_test.go
- internal/toml/decode.go
- toml-decoder

## Pre-build acceptance observation
- GREEN (prepassed): decoder-valid-stdin (FEATURE-Decoder-Contract.md)
  intent: A valid TOML document is accepted through stdin and produces a JSON object.
  return code: 0
- RED: decoder-argument-free (FEATURE-Decoder-Contract.md)
  intent: Supplying an argument is rejected by the argument-free command contract.
  return code: 1
  stderr:
    Traceback (most recent call last):
      File "decoder-argument-free.py", line 9, in <module>
        assert result.returncode != 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- GREEN (prepassed): decoder-stream-separation (FEATURE-Decoder-Contract.md)
  intent: Successful decoding does not write diagnostics to stderr.
  return code: 0

## Post-build programmatic acceptance
- PASS: decoder-valid-stdin (FEATURE-Decoder-Contract.md)
  intent: A valid TOML document is accepted through stdin and produces a JSON object.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
- PASS: decoder-argument-free (FEATURE-Decoder-Contract.md)
  intent: Supplying an argument is rejected by the argument-free command contract.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
- PASS: decoder-stream-separation (FEATURE-Decoder-Contract.md)
  intent: Successful decoding does not write diagnostics to stderr.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0

## Agent self-report (advisory)
The build agent declared a failure. This is advisory only; the programmatic acceptance above is the authority for this block's outcome.
- summary: Authoritative conformance suite failed.
- detail:
    `sh sources/full_test.sh` completed with exit status 1 due to remaining parser grammar and table-semantics failures. Further TOML 1.0 parser implementation is required.

## Build summary
RESULT: FAILED

FILES CHANGED:
- cmd/toml-decoder/main.go
- cmd/toml-decoder/main_test.go
- internal/toml/decode.go

SUMMARY:
Implemented stdin decoding, tagged JSON output, argument rejection, diagnostics, structured values, tables, arrays, and process-level tests. Unit tests and `go vet` pass.

Full conformance: 158 valid passed / 47 failed; 345 invalid passed / 129 failed.

BLOCKERS:
- Full TOML 1.0 conformance is not yet passing.

FAILURE_SUMMARY: Authoritative conformance suite failed.
FAILURE_DETAIL: `sh sources/full_test.sh` completed with exit status 1 due to remaining parser grammar and table-semantics failures. Further TOML 1.0 parser implementation is required.
