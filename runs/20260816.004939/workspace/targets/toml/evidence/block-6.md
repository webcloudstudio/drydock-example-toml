# Evidence: Block 6 · Service (block-6)

- block type: block
- date: 2026-08-15
- resulting state: closed/verified
- story points (combined assembled cost): 8714
- execution id: 20260816.011726.759Z-f3150fe3

## Stories built
- Parse bare, quoted, and dotted keys with duplicate protection. (parser-keys) [story]

## Stacked context
- compass: COMPASS.md (SP 2899)
- implements: FEATURE-Parser-Keys.md (SP 637)
- context: stage_test.sh (SP 161)
- context: ARCHITECTURE_compact.md (SP 128)
- stack: go.md (SP 3646)
- stack: common_compact.md (SP 1179)

## Build directory changes
- internal/toml/toml.go
- internal/toml/toml_test.go
- toml-decoder

## Pre-build acceptance observation
- GREEN (prepassed): parser-keys-conformance (FEATURE-Parser-Keys.md)
  intent: The implementation passes the authoritative TOML key cases owned by this story.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.004939/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- RED: parser-keys-invalid-extension (FEATURE-Parser-Keys.md)
  intent: A scalar key cannot later be extended as a dotted table path.
  return code: 1
  stdout:
    {"apple":{"type":"integer","value":"1"},"fruit":{"smooth":{"type":"bool","value":"true"}}}
  stderr:
    Traceback (most recent call last):
      File "parser-keys-invalid-extension.py", line 12, in <module>
        assert result.returncode != 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- GREEN (prepassed): parser-keys-duplicate (FEATURE-Parser-Keys.md)
  intent: Repeating an already-defined key is rejected with a nonzero exit status.
  return code: 0
  stdout:
    line 2: duplicate key "name"

## Post-build programmatic acceptance
- PASS: parser-keys-conformance (FEATURE-Parser-Keys.md)
  intent: The implementation passes the authoritative TOML key cases owned by this story.
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
- PASS: parser-keys-invalid-extension (FEATURE-Parser-Keys.md)
  intent: A scalar key cannot later be extended as a dotted table path.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    line 2: table redefinition
- PASS: parser-keys-duplicate (FEATURE-Parser-Keys.md)
  intent: Repeating an already-defined key is rejected with a nonzero exit status.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    line 2: duplicate key "name"

## Build summary
RESULT: SUCCESS

FILES CHANGED:
- internal/toml/toml.go
- internal/toml/toml_test.go

SUMMARY:
Implemented bare, quoted, dotted, and whitespace-separated keys with duplicate and scalar-extension protection. Added deterministic key tests. `go test -race ./...`, vet, build, and declared key stage passed.

BLOCKERS:
- None
