# Evidence: Block 11 · Service (block-11)

- block type: block
- date: 2026-08-14
- resulting state: closed/verified
- story points (combined assembled cost): 8785
- execution id: 20260814.055428.590Z-e38d6c51

## Stories built
- Enforce standard-table redefinition and structure-conflict rules. (tables-redefinition) [story]
- Parse arrays of TOML tables and nested table elements. (tables-arrays) [story]

## Missing context
- context: toml-v1.0.0.md

## Stacked context
- compass: COMPASS.md (SP 2629)
- implements: FEATURE-Table-Redefinition.md (SP 425)
- context: stage_test.sh (SP 161)
- context: ARCHITECTURE_compact.md (SP 108)
- stack: go.md (SP 3646)
- stack: common_compact.md (SP 1179)
- implements: FEATURE-Arrays-of-Tables.md (SP 546)

## Build directory changes
- internal/toml/value_test.go

## Pre-build acceptance observation
- GREEN (prepassed): table-redefinition-invalid (FEATURE-Table-Redefinition.md)
  intent: The implementation rejects every authoritative invalid table-redefinition case.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 822
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  62 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- GREEN (prepassed): table-redefinition-valid (FEATURE-Table-Redefinition.md)
  intent: The implementation preserves all valid standard-table definitions while enforcing redefinition rules.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 860
      valid tests:  24 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- GREEN (prepassed): arrays-of-tables-conformance (FEATURE-Arrays-of-Tables.md)
  intent: The implementation passes the authoritative TOML table conformance slice owned by this story.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 860
      valid tests:  24 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- GREEN (prepassed): arrays-of-tables-build (FEATURE-Arrays-of-Tables.md)
  intent: The decoder builds successfully after array-of-table support is implemented.
  return code: 0

## Post-build programmatic acceptance
- PASS: table-redefinition-invalid (FEATURE-Table-Redefinition.md)
  intent: The implementation rejects every authoritative invalid table-redefinition case.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 822
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  62 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: table-redefinition-valid (FEATURE-Table-Redefinition.md)
  intent: The implementation preserves all valid standard-table definitions while enforcing redefinition rules.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 860
      valid tests:  24 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: arrays-of-tables-conformance (FEATURE-Arrays-of-Tables.md)
  intent: The implementation passes the authoritative TOML table conformance slice owned by this story.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 860
      valid tests:  24 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: arrays-of-tables-build (FEATURE-Arrays-of-Tables.md)
  intent: The decoder builds successfully after array-of-table support is implemented.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0

## Build summary
RESULT: SUCCESS

FILES CHANGED:
- internal/toml/value_test.go

SUMMARY:
Added regression tests for ordered/nested arrays of tables, empty elements, and table structure conflicts. Go tests, vet, formatting, and both declared table conformance gates pass:
- Valid: 24 passed, 0 failed
- Invalid: 62 passed, 0 failed

BLOCKERS:
- Full suite retains unrelated pre-existing failures in multiline quotes, float lexical validation, and signed non-decimal integers.
