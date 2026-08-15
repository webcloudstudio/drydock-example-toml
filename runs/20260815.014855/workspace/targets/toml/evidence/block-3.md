# Evidence: Block 3 · Service (block-3)

- block type: block
- date: 2026-08-14
- resulting state: closed/verified
- story points (combined assembled cost): 9513
- execution id: 20260815.015856.448Z-283ea0ec

## Stories built
- Parse all TOML string forms. (parser-strings) [story]
- Parse booleans, integers, and floating-point values. (parser-lexical-scalars) [story]
- Enforce comments, control characters, UTF-8, and lexical boundaries. (parser-lexical-boundaries) [story]

## Acceptance tooling authorization
- FEATURE-Parser-Strings.md#parser-strings-conformance: executable=sh; scope=test; authorization=existing Target environment
- FEATURE-Parser-Strings.md#parser-strings-conformance: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Parser-Lexical-Scalars.md#parser-lexical-scalars-conformance: executable=sh; scope=test; authorization=existing Target environment
- FEATURE-Parser-Lexical-Scalars.md#parser-lexical-scalars-conformance: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Parser-Lexical-Boundaries.md#parser-lexical-boundaries-conformance: executable=sh; scope=test; authorization=existing Target environment
- FEATURE-Parser-Lexical-Boundaries.md#parser-lexical-boundaries-conformance: executable=go; scope=test; authorization=existing Target environment

## Reusable compacts
- stage_test_compact.md

## Missing context
- context: toml-v1.0.0.md

## Stacked context
- compass: COMPASS.md (SP 2932)
- implements: FEATURE-Parser-Strings.md (SP 430)
- context: stage_test.sh (SP 161)
- context: ARCHITECTURE_compact.md (SP 122)
- stack: go.md (SP 3646)
- stack: common_compact.md (SP 1179)
- implements: FEATURE-Parser-Lexical-Scalars.md (SP 455)
- implements: FEATURE-Parser-Lexical-Boundaries.md (SP 431)

## Build directory changes
- internal/toml/toml.go
- internal/toml/toml_test.go
- toml-decoder

## Pre-build acceptance observation
- GREEN (prepassed): parser-strings-conformance (FEATURE-Parser-Strings.md)
  intent: The authoritative TOML conformance suite passes every valid and invalid string case owned by this story.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260815.014855/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- GREEN (prepassed): parser-lexical-scalars-conformance (FEATURE-Parser-Lexical-Scalars.md)
  intent: The authoritative TOML conformance suite passes every valid and invalid boolean, integer, and float case owned by this story.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260815.014855/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- GREEN (prepassed): parser-lexical-boundaries-conformance (FEATURE-Parser-Lexical-Boundaries.md)
  intent: The authoritative TOML conformance suite passes every valid comment and invalid control-character and encoding case owned by this story.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260815.014855/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true

## Post-build programmatic acceptance
- PASS: parser-strings-conformance (FEATURE-Parser-Strings.md)
  intent: The authoritative TOML conformance suite passes every valid and invalid string case owned by this story.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260815.014855/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: parser-lexical-scalars-conformance (FEATURE-Parser-Lexical-Scalars.md)
  intent: The authoritative TOML conformance suite passes every valid and invalid boolean, integer, and float case owned by this story.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260815.014855/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: parser-lexical-boundaries-conformance (FEATURE-Parser-Lexical-Boundaries.md)
  intent: The authoritative TOML conformance suite passes every valid comment and invalid control-character and encoding case owned by this story.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260815.014855/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true

## Agent self-report (advisory)
The build agent declared a failure. This is advisory only; the programmatic acceptance above is the authority for this block's outcome.
- summary: Full conformance is incomplete due to deferred arrays, tables, and datetime support.
- detail:
    `sh sources/full_test.sh` reported 76 valid passed and 129 failed; failures are primarily later-block features such as arrays and tables. Implement those subsequent parser blocks before rerunning the full suite.

## Build summary
<reusable-compact filename="stage_test.sh">
Runs one governed conformance slice: `sh sources/stage_test.sh '<toml-test -run pattern>'`. Requires exactly one pattern argument, builds `./cmd/toml-decoder`, sets `TOML_TEST_VERSION` (default `v2.2.0`), and uses the decoder’s exit status as the verdict.
</reusable-compact>

RESULT: FAILED

FILES CHANGED:
- internal/toml/toml.go
- internal/toml/toml_test.go

SUMMARY:
Implemented TOML strings, scalar lexing, UTF-8/control validation, comments, and lexical boundaries. Go tests and vet pass. String, scalar, control, and encoding cases pass when filtered with `*`. Full suite remains incomplete because arrays, tables, and datetimes are later build blocks.

BLOCKERS:
- Valid comment cases containing tables/arrays still fail, as those parser features are not yet implemented.

FAILURE_SUMMARY: Full conformance is incomplete due to deferred arrays, tables, and datetime support.
FAILURE_DETAIL: `sh sources/full_test.sh` reported 76 valid passed and 129 failed; failures are primarily later-block features such as arrays and tables. Implement those subsequent parser blocks before rerunning the full suite.
