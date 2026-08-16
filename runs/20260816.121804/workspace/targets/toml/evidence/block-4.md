# Evidence: Block 4 · Service (block-4)

- block type: block
- date: 2026-08-16
- resulting state: closed/verified
- story points (combined assembled cost): 7534
- execution id: 20260816.123731.533Z-77b80c17

## Stories built
- Parse TOML booleans, integers, and floating-point scalar values. (parser-lexical-scalars) [story]

## Missing context
- context: toml-v1.0.0.md

## Stacked context
- compass: COMPASS.md (SP 2942)
- implements: FEATURE-Lexical-Scalars.md (SP 602)
- context: stage_test.sh (SP 161)
- context: ARCHITECTURE_compact.md (SP 123)
- stack: go.md (SP 3646)

## Build directory changes
- internal/toml/toml.go
- internal/toml/toml_test.go
- toml-decoder

## Pre-build acceptance observation
- GREEN (prepassed): scalars-valid-suite (FEATURE-Lexical-Scalars.md)
  intent: The authoritative conformance suite passes all valid boolean, integer, and float cases.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- GREEN (prepassed): scalars-invalid-suite (FEATURE-Lexical-Scalars.md)
  intent: The authoritative conformance suite rejects all invalid boolean, integer, and float cases.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- GREEN (prepassed): scalars-bounded-suites (FEATURE-Lexical-Scalars.md)
  intent: The valid and invalid scalar conformance slices each complete successfully.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true

## Post-build programmatic acceptance
- PASS: scalars-valid-suite (FEATURE-Lexical-Scalars.md)
  intent: The authoritative conformance suite passes all valid boolean, integer, and float cases.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: scalars-invalid-suite (FEATURE-Lexical-Scalars.md)
  intent: The authoritative conformance suite rejects all invalid boolean, integer, and float cases.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: scalars-bounded-suites (FEATURE-Lexical-Scalars.md)
  intent: The valid and invalid scalar conformance slices each complete successfully.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true

## Agent self-report (advisory)
The build agent declared a failure. This is advisory only; the programmatic acceptance above is the authority for this block's outcome.
- summary: Full conformance suite has 62 valid and 33 invalid failures outside this scalar block.
- detail:
    Scalar-specific tests pass, but completing the full gate requires implementing arrays, tables, comments, strings, and datetime parsing in subsequent build blocks.

## Build summary
RESULT: FAILED

FILES CHANGED:
- internal/toml/toml.go
- internal/toml/toml_test.go

SUMMARY:
Implemented TOML booleans, canonical integers, base-prefixed integers, floats, exponents, special values, validation, and tests. Scalar suites pass; full suite remains failing in out-of-scope arrays, tables, strings, comments, and datetimes.

BLOCKERS:
- Full conformance suite is not yet green due to prior unimplemented parser features.

FAILURE_SUMMARY: Full conformance suite has 62 valid and 33 invalid failures outside this scalar block.
FAILURE_DETAIL: Scalar-specific tests pass, but completing the full gate requires implementing arrays, tables, comments, strings, and datetime parsing in subsequent build blocks.
