# Evidence: Block 2 · Service (block-2)

- block type: block
- date: 2026-08-15
- resulting state: closed/verified
- story points (combined assembled cost): 9552
- execution id: 20260816.005855.689Z-33607cc9

## Stories built
- Implement the stdin, tagged-JSON, diagnostic, and exit-status contract. (decoder-contract) [story]

## Acceptance tooling authorization
- FEATURE-Decoder-Contract.md#decoder-valid-contract: executable=sh; scope=test; authorization=existing Target environment
- FEATURE-Decoder-Contract.md#decoder-valid-contract: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-Decoder-Contract.md#decoder-invalid-contract: executable=sh; scope=test; authorization=existing Target environment
- FEATURE-Decoder-Contract.md#decoder-no-arguments: executable=go; scope=test; authorization=existing Target environment

## Stacked context
- compass: COMPASS.md (SP 2899)
- implements: FEATURE-Decoder-Contract.md (SP 665)
- context: stage_contract.sh (SP 317)
- context: ARCHITECTURE_compact.md (SP 128)
- stack: go.md (SP 3646)
- stack: common.md (SP 1807)

## Build directory changes
- cmd/toml-decoder/main.go
- internal/toml/toml.go
- internal/toml/toml_test.go

## Pre-build acceptance observation
- RED: decoder-valid-contract (FEATURE-Decoder-Contract.md)
  intent: Valid stdin produces the required tagged JSON state and no diagnostic.
  return code: 1
  stderr:
    Traceback (most recent call last):
      File "decoder-valid-contract.py", line 11, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: decoder-invalid-contract (FEATURE-Decoder-Contract.md)
  intent: Malformed stdin is rejected without a successful stdout payload.
  return code: 1
  stderr:
    Traceback (most recent call last):
      File "decoder-invalid-contract.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- GREEN (prepassed): decoder-no-arguments (FEATURE-Decoder-Contract.md)
  intent: The command boundary rejects unsupported command-line arguments.
  return code: 0

## Post-build programmatic acceptance
- PASS: decoder-valid-contract (FEATURE-Decoder-Contract.md)
  intent: Valid stdin produces the required tagged JSON state and no diagnostic.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
- PASS: decoder-invalid-contract (FEATURE-Decoder-Contract.md)
  intent: Malformed stdin is rejected without a successful stdout payload.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
- PASS: decoder-no-arguments (FEATURE-Decoder-Contract.md)
  intent: The command boundary rejects unsupported command-line arguments.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0

## Build summary
RESULT: SUCCESS

FILES CHANGED:
- cmd/toml-decoder/main.go
- internal/toml/toml.go
- internal/toml/toml_test.go

SUMMARY:
Implemented stdin/stdout/stderr handling, tagged JSON output, diagnostics, nonzero exits, argument rejection, and minimal integer decoding. All declared contract checks, `go test -race`, `go vet`, and formatting checks pass.

BLOCKERS:
- None
