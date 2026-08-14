# Evidence: Block 3 · Service (block-3)

- block type: block
- date: 2026-08-14
- resulting state: closed/verified
- story points (combined assembled cost): 8169
- execution id: 20260814.051622.557Z-0ad1458e

## Stories built
- Report invalid TOML through the decoder process boundary. (decoder-invalid-input) [story]

## Stacked context
- compass: COMPASS.md (SP 2629)
- implements: FEATURE-Decoder-Invalid-Input.md (SP 548)
- context: ARCHITECTURE_compact.md (SP 108)
- stack: go.md (SP 3646)
- stack: common_compact.md (SP 1179)

## Build directory changes
- cmd/toml-decoder/main_test.go

## Pre-build acceptance observation
- GREEN (prepassed): invalid-exit-status (FEATURE-Decoder-Invalid-Input.md)
  intent: Malformed TOML exits non-zero.
  return code: 0
- GREEN (prepassed): invalid-diagnostic-stream (FEATURE-Decoder-Invalid-Input.md)
  intent: Malformed TOML produces a non-empty diagnostic on stderr.
  return code: 0
- GREEN (prepassed): invalid-no-success-result (FEATURE-Decoder-Invalid-Input.md)
  intent: Malformed TOML does not produce a successful JSON result on stdout.
  return code: 0

## Post-build programmatic acceptance
- PASS: invalid-exit-status (FEATURE-Decoder-Invalid-Input.md)
  intent: Malformed TOML exits non-zero.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
- PASS: invalid-diagnostic-stream (FEATURE-Decoder-Invalid-Input.md)
  intent: Malformed TOML produces a non-empty diagnostic on stderr.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
- PASS: invalid-no-success-result (FEATURE-Decoder-Invalid-Input.md)
  intent: Malformed TOML does not produce a successful JSON result on stdout.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0

## Agent self-report (advisory)
The build agent declared a failure. This is advisory only; the programmatic acceptance above is the authority for this block's outcome.
- summary: Full conformance suite remains incomplete.
- detail:
    Existing parser behavior fails unrelated TOML syntax and semantic cases; broader parser implementation is required before the full-suite gate can pass.

## Build summary
RESULT: FAILED

FILES CHANGED:
- cmd/toml-decoder/main_test.go

SUMMARY:
Added process-level coverage proving malformed TOML exits non-zero, writes stderr diagnostics, and emits no stdout. Targeted tests, race tests, vet, formatting, and direct acceptance checks pass.

BLOCKERS:
- Full toml-test v2.2.0 suite fails: 47 valid and 129 invalid cases failed.

FAILURE_SUMMARY: Full conformance suite remains incomplete.
FAILURE_DETAIL: Existing parser behavior fails unrelated TOML syntax and semantic cases; broader parser implementation is required before the full-suite gate can pass.
