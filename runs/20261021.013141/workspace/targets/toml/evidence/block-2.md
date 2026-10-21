# Evidence: Block 2 · Service (block-2)

- block type: block
- date: 2026-10-20
- resulting state: closed/verified
- story points (combined assembled cost): 4584
- execution id: 20261021.013610.964Z-acd37d69

## Stories built
- Implement the stdin, tagged-JSON, diagnostics, and exit-status contract. (decoder-contract) [story]

## Missing context
- context: INSTRUCTIONS.md

## Stacked context
- stack: go_compact.md (SP 1173)
- compass: COMPASS.md (SP 521)
- implements: FEATURE-Decoder-Contract.md (SP 510)
- context: stage_contract.sh (SP 317)
- context: ARCHITECTURE_compact.md (SP 136)
- stack: common.md (SP 1812)

## Build directory changes
- cmd/toml-decoder/main.go
- cmd/toml-decoder/main_test.go
- internal/toml/toml.go
- internal/toml/toml_test.go
- toml-decoder

## Pre-build acceptance observation
- GREEN (prepassed): decoder-contract-gate (FEATURE-Decoder-Contract.md)
  intent: The executable satisfies the supplied valid-input and invalid-input stream contract.
  return code: 0
- GREEN (prepassed): decoder-executable (FEATURE-Decoder-Contract.md)
  intent: The decoder executable is produced at the declared command boundary.
  return code: 0

## Post-build programmatic acceptance
- PASS: decoder-contract-gate (FEATURE-Decoder-Contract.md)
  intent: The executable satisfies the supplied valid-input and invalid-input stream contract.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: decoder-executable (FEATURE-Decoder-Contract.md)
  intent: The decoder executable is produced at the declared command boundary.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0

## Build summary
RESULT: SUCCESS

FILES CHANGED:
- cmd/toml-decoder/main.go
- cmd/toml-decoder/main_test.go
- internal/toml/toml.go
- internal/toml/toml_test.go
- toml-decoder

SUMMARY:
Implemented tagged TOML decoding, stdin/stdout contract, diagnostics, exit statuses, argument rejection, tables, arrays, and scalar types. Verified with Go tests, vet, build, and decoder contract gate.

BLOCKERS:
- None
