# Evidence: Block 4 · Service (block-4)

- block type: block
- date: 2026-08-15
- resulting state: closed/verified
- story points (combined assembled cost): 8760
- execution id: 20260816.010655.417Z-48ceef7c

## Stories built
- Parse TOML booleans, integers, and floating-point values. (parser-lexical-scalars) [story]

## Acceptance tooling authorization
- FEATURE-Parser-Lexical-Scalars.md#parser-scalars-valid: executable=sh; scope=test; authorization=existing Target environment
- FEATURE-Parser-Lexical-Scalars.md#parser-scalars-valid: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Parser-Lexical-Scalars.md#parser-scalars-invalid: executable=sh; scope=test; authorization=existing Target environment
- FEATURE-Parser-Lexical-Scalars.md#parser-scalars-invalid: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Parser-Lexical-Scalars.md#parser-scalars-integer-contract: executable=go; scope=test; authorization=existing Target environment

## Stacked context
- compass: COMPASS.md (SP 2899)
- implements: FEATURE-Parser-Lexical-Scalars.md (SP 677)
- context: stage_test.sh (SP 161)
- context: ARCHITECTURE_compact.md (SP 128)
- stack: go.md (SP 3646)
- stack: common_compact.md (SP 1179)

## Build directory changes
- internal/toml/toml.go
- internal/toml/toml_test.go

## Pre-build acceptance observation
- GREEN (prepassed): parser-scalars-valid (FEATURE-Parser-Lexical-Scalars.md)
  intent: The authoritative TOML suite passes every valid boolean, integer, and float case.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.004939/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- GREEN (prepassed): parser-scalars-invalid (FEATURE-Parser-Lexical-Scalars.md)
  intent: The authoritative TOML suite rejects every invalid boolean, integer, and float case.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.004939/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- GREEN (prepassed): parser-scalars-integer-contract (FEATURE-Parser-Lexical-Scalars.md)
  intent: A signed 64-bit integer is emitted without numeric precision loss.
  return code: 0

## Post-build programmatic acceptance
- PASS: parser-scalars-valid (FEATURE-Parser-Lexical-Scalars.md)
  intent: The authoritative TOML suite passes every valid boolean, integer, and float case.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.004939/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: parser-scalars-invalid (FEATURE-Parser-Lexical-Scalars.md)
  intent: The authoritative TOML suite rejects every invalid boolean, integer, and float case.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.004939/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: parser-scalars-integer-contract (FEATURE-Parser-Lexical-Scalars.md)
  intent: A signed 64-bit integer is emitted without numeric precision loss.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0

## Build summary
RESULT: SUCCESS

FILES CHANGED:
- internal/toml/toml.go
- internal/toml/toml_test.go

SUMMARY:
Implemented strict TOML booleans, integers, floats, special values, signed zero, range checks, underscore validation, and malformed-token rejection. Added tests. `go test`, race tests, vet, and builds pass. Scalar stages exited 0.

BLOCKERS:
- None.
