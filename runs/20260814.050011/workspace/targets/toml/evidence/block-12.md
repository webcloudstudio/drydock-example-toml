# Evidence: Block 12 · Service (block-12)

- block type: block
- date: 2026-08-14
- resulting state: closed/verified
- story points (combined assembled cost): 8313
- execution id: 20260814.055701.009Z-c7e9a43b

## Stories built
- Enforce array-of-table ordering and conflict rules. (tables-arrays-ordering) [story]

## Missing context
- context: toml-v1.0.0.md

## Stacked context
- compass: COMPASS.md (SP 2629)
- implements: FEATURE-Arrays-of-Tables-Ordering.md (SP 541)
- context: stage_test.sh (SP 161)
- context: ARCHITECTURE_compact.md (SP 108)
- stack: go.md (SP 3646)
- stack: common_compact.md (SP 1179)

## Build directory changes
- internal/toml/decode.go
- internal/toml/value_test.go
- toml-decoder

## Pre-build acceptance observation
- GREEN (prepassed): arrays-of-tables-ordering-conformance (FEATURE-Arrays-of-Tables-Ordering.md)
  intent: The implementation rejects the authoritative invalid table-ordering and conflict cases.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 822
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  62 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- UNVERIFIED: arrays-of-tables-ordering-exit (FEATURE-Arrays-of-Tables-Ordering.md)
  intent: An invalid array-of-table ordering document exits unsuccessfully.
  return code: 1
  error: malformed check: the assertion itself raised NameError (name 'subprocess' is not defined. Did you forget to import 'subprocess'?) in its own frame, before reaching the code under test. No implementation can satisfy it. Each check runs as its own script in its own process, so a name bound by another check is not in scope. Repair the assertion in the Blueprint specification.
  stderr:
    Traceback (most recent call last):
      File "arrays-of-tables-ordering-exit.py", line 2, in <module>
        result = subprocess.run(
                 ^^^^^^^^^^
    NameError: name 'subprocess' is not defined. Did you forget to import 'subprocess'?

## Post-build programmatic acceptance
- PASS: arrays-of-tables-ordering-conformance (FEATURE-Arrays-of-Tables-Ordering.md)
  intent: The implementation rejects the authoritative invalid table-ordering and conflict cases.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 822
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  62 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- UNVERIFIED: arrays-of-tables-ordering-exit (FEATURE-Arrays-of-Tables-Ordering.md)
  intent: An invalid array-of-table ordering document exits unsuccessfully.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 1
  error: malformed check: the assertion itself raised NameError (name 'subprocess' is not defined. Did you forget to import 'subprocess'?) in its own frame, before reaching the code under test. No implementation can satisfy it. Each check runs as its own script in its own process, so a name bound by another check is not in scope. Repair the assertion in the Blueprint specification.
  stderr:
    Traceback (most recent call last):
      File "arrays-of-tables-ordering-exit.py", line 2, in <module>
        result = subprocess.run(
                 ^^^^^^^^^^
    NameError: name 'subprocess' is not defined. Did you forget to import 'subprocess'?

## Build summary
RESULT: SUCCESS

FILES CHANGED:
- internal/toml/decode.go
- internal/toml/value_test.go

SUMMARY:
Added static-array ancestor validation for nested array-of-table headers and regression tests for ordering/conflict rules. Scoped suite: 62/62 invalid cases passed. Go tests, race tests, vet, and build passed.

BLOCKERS:
- Full suite retains unrelated pre-existing failures in multiline strings and numeric validation.
