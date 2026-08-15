# Evidence: Block 2 · Service (block-2)

- block type: block
- date: 2026-08-14
- resulting state: closed/verified
- story points (combined assembled cost): 8978
- execution id: 20260815.015711.540Z-414b87b5

## Stories built
- Implement the stdin-to-tagged-JSON decoder contract. (decoder-contract) [story]

## Stacked context
- compass: COMPASS.md (SP 2932)
- implements: FEATURE-Decoder-Contract.md (SP 722)
- context: stage_contract.sh (SP 317)
- context: ARCHITECTURE_compact.md (SP 122)
- stack: go.md (SP 3646)
- stack: common_compact.md (SP 1179)

## Build directory changes
- cmd/toml-decoder/main.go
- cmd/toml-decoder/main_test.go
- toml-decoder

## Pre-build acceptance observation
- GREEN (prepassed): decoder-valid-contract (FEATURE-Decoder-Contract.md)
  intent: A valid TOML document produces tagged JSON on stdout without stderr diagnostics.
  return code: 0
  stdout:
    {"answer":{"type":"integer","value":"42"}}
- GREEN (prepassed): decoder-invalid-contract (FEATURE-Decoder-Contract.md)
  intent: Malformed TOML produces a non-zero exit, stderr diagnostics, and no stdout document.
  return code: 0
  stdout:
    line 1: unsupported value
- GREEN (prepassed): decoder-stage-contract (FEATURE-Decoder-Contract.md)
  intent: The supplied decoder contract gate passes for valid and malformed stdin documents.
  return code: 0

## Post-build programmatic acceptance
- PASS: decoder-valid-contract (FEATURE-Decoder-Contract.md)
  intent: A valid TOML document produces tagged JSON on stdout without stderr diagnostics.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    {"answer":{"type":"integer","value":"42"}}
- PASS: decoder-invalid-contract (FEATURE-Decoder-Contract.md)
  intent: Malformed TOML produces a non-zero exit, stderr diagnostics, and no stdout document.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    line 1: unsupported value
- PASS: decoder-stage-contract (FEATURE-Decoder-Contract.md)
  intent: The supplied decoder contract gate passes for valid and malformed stdin documents.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0

## Build summary
RESULT: SUCCESS

FILES CHANGED:
- cmd/toml-decoder/main.go
- cmd/toml-decoder/main_test.go

SUMMARY:
Implemented and tested stdin decoding, tagged JSON output, diagnostics, and non-zero error behavior. `go test -race`, `go vet`, formatting, and `stage_contract.sh decoder` all pass.

BLOCKERS:
- None
