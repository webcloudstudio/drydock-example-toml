# Evidence: Block 9 · Service (block-9)

- block type: block
- date: 2026-08-15
- resulting state: closed/verified
- story points (combined assembled cost): 8881
- execution id: 20260816.013008.609Z-2aa9dd3b

## Stories built
- Parse offset and local TOML date/time values with required encoding. (parser-datetimes) [story]

## Stacked context
- compass: COMPASS.md (SP 2899)
- implements: FEATURE-Parser-Datetimes.md (SP 791)
- context: stage_test.sh (SP 161)
- context: ARCHITECTURE_compact.md (SP 128)
- stack: go.md (SP 3646)
- stack: common_compact.md (SP 1179)

## Build directory changes
- internal/toml/toml.go
- internal/toml/toml_test.go
- toml-decoder

## Pre-build acceptance observation
- GREEN (prepassed): parser-datetimes-conformance (FEATURE-Parser-Datetimes.md)
  intent: The implementation passes the authoritative TOML datetime cases owned by this story.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.004939/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- GREEN (prepassed): parser-datetimes-types (FEATURE-Parser-Datetimes.md)
  intent: Offset and local date/time inputs receive the declared tagged types.
  return code: 0
  stdout:
    {"local_date":{"type":"date-local","value":"1979-05-27"},"local_datetime":{"type":"datetime-local","value":"1979-05-27T07:32:00"},"local_time":{"type":"time-local","value":"07:32:00"},"offset":{"type":"datetime","value":"1979-05-27T07:32:00Z"}}
- RED: parser-datetimes-invalid (FEATURE-Parser-Datetimes.md)
  intent: Invalid local date/time syntax is rejected with a nonzero exit status.
  return code: 1
  stdout:
    {"value":{"type":"date-local","value":"1979-02-30"}}
  stderr:
    Traceback (most recent call last):
      File "parser-datetimes-invalid.py", line 12, in <module>
        assert result.returncode != 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- PASS: parser-datetimes-conformance (FEATURE-Parser-Datetimes.md)
  intent: The implementation passes the authoritative TOML datetime cases owned by this story.
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
- PASS: parser-datetimes-types (FEATURE-Parser-Datetimes.md)
  intent: Offset and local date/time inputs receive the declared tagged types.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    {"local_date":{"type":"date-local","value":"1979-05-27"},"local_datetime":{"type":"datetime-local","value":"1979-05-27T07:32:00"},"local_time":{"type":"time-local","value":"07:32:00"},"offset":{"type":"datetime","value":"1979-05-27T07:32:00Z"}}
- PASS: parser-datetimes-invalid (FEATURE-Parser-Datetimes.md)
  intent: Invalid local date/time syntax is rejected with a nonzero exit status.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    line 1: invalid date

## Agent self-report (advisory)
The build agent declared a failure. This is advisory only; the programmatic acceptance above is the authority for this block's outcome.
- summary: Full conformance suite fails outside the datetime story.
- detail:
    Datetime-scoped verification passes, but sources/full_test.sh exits nonzero due unrelated integer normalization, float validation, and table semantics failures. Those areas must be fixed in their respective build blocks.

## Build summary
RESULT: FAILED

FILES CHANGED:
- internal/toml/toml.go
- internal/toml/toml_test.go

SUMMARY:
Implemented strict TOML datetime parsing, validation, RFC 3339 encoding, and fractional truncation. Datetime tests passed: 9 valid, 35 invalid datetime, 12 invalid dates, 15 invalid local datetimes, and 8 invalid local times. Unit tests and vet pass.

BLOCKERS:
- Full suite retains 8 valid and 22 invalid failures in pre-existing integer, float, and table behavior.

FAILURE_SUMMARY: Full conformance suite fails outside the datetime story.
FAILURE_DETAIL: Datetime-scoped verification passes, but sources/full_test.sh exits nonzero due unrelated integer normalization, float validation, and table semantics failures. Those areas must be fixed in their respective build blocks.
