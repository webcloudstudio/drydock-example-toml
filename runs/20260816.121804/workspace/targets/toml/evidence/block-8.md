# Evidence: Block 8 · Feature (block-8)

- block type: block
- date: 2026-08-16
- resulting state: closed/verified
- story points (combined assembled cost): 9453
- execution id: 20260816.130124.690Z-22f45be7

## Stories built
- Assemble the completed parser and verify the complete TOML 1.0.0 suite. (conformance) [story]

## Acceptance tooling authorization
- FEATURE-Conformance.md#conformance-full: executable=go; scope=test; authorization=existing Target environment

## Stacked context
- compass: COMPASS.md (SP 2942)
- implements: FEATURE-Conformance.md (SP 409)
- context: full_test.sh (SP 113)
- context: run_conformance.sh (SP 677)
- context: setup_harness.sh (SP 323)
- context: ARCHITECTURE_compact.md (SP 123)
- stack: go.md (SP 3646)
- stack: common_compact.md (SP 1179)

## Build directory changes
- cmd/toml-decoder/main.go
- toml-decoder

## Pre-build acceptance observation
- GREEN (prepassed): conformance-full (FEATURE-Conformance.md)
  intent: The completed decoder passes the supplied unfiltered TOML 1.0.0 conformance suite.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
      valid tests: 205 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests: 474 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true

## Post-build programmatic acceptance
- PASS: conformance-full (FEATURE-Conformance.md)
  intent: The completed decoder passes the supplied unfiltered TOML 1.0.0 conformance suite.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
      valid tests: 205 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests: 474 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true

## Build summary
RESULT: SUCCESS

FILES CHANGED:
- cmd/toml-decoder/main.go

SUMMARY:
- Added rejection of unexpected command-line arguments.
- `go test`, `go test -race`, `go vet`, and build pass.
- Full TOML suite passes: 205 valid and 474 invalid, zero failures.

BLOCKERS:
- None
