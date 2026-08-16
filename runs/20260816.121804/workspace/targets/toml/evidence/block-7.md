# Evidence: Block 7 · Service (block-7)

- block type: block
- date: 2026-08-16
- resulting state: closed/verified
- story points (combined assembled cost): 7658
- execution id: 20260816.125657.658Z-e4cf954c

## Stories built
- Parse arrays and inline tables with closure and conflict semantics. (parser-arrays-inline) [story]

## Acceptance tooling authorization
- FEATURE-Arrays-Inline.md#arrays-valid: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Arrays-Inline.md#inline-valid: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Arrays-Inline.md#arrays-invalid: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Arrays-Inline.md#inline-invalid: executable=go; scope=test; authorization=existing Target environment

## Missing context
- context: toml-v1.0.0.md

## Stacked context
- compass: COMPASS.md (SP 2942)
- implements: FEATURE-Arrays-Inline.md (SP 718)
- context: stage_test.sh (SP 161)
- context: ARCHITECTURE_compact.md (SP 123)
- stack: go.md (SP 3646)

## Build directory changes
- internal/toml/toml_test.go

## Pre-build acceptance observation
- GREEN (prepassed): arrays-valid (FEATURE-Arrays-Inline.md)
  intent: The decoder passes all authoritative valid TOML array cases owned by this feature.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 863
      valid tests:  21 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- GREEN (prepassed): inline-valid (FEATURE-Arrays-Inline.md)
  intent: The decoder passes all authoritative valid TOML inline-table cases owned by this feature.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 867
      valid tests:  17 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- GREEN (prepassed): arrays-invalid (FEATURE-Arrays-Inline.md)
  intent: The decoder rejects all authoritative invalid TOML array cases owned by this feature.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 858
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  26 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- GREEN (prepassed): inline-invalid (FEATURE-Arrays-Inline.md)
  intent: The decoder rejects all authoritative invalid TOML inline-table cases owned by this feature.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 856
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  28 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true

## Post-build programmatic acceptance
- PASS: arrays-valid (FEATURE-Arrays-Inline.md)
  intent: The decoder passes all authoritative valid TOML array cases owned by this feature.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 863
      valid tests:  21 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: inline-valid (FEATURE-Arrays-Inline.md)
  intent: The decoder passes all authoritative valid TOML inline-table cases owned by this feature.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 867
      valid tests:  17 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: arrays-invalid (FEATURE-Arrays-Inline.md)
  intent: The decoder rejects all authoritative invalid TOML array cases owned by this feature.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 858
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  26 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: inline-invalid (FEATURE-Arrays-Inline.md)
  intent: The decoder rejects all authoritative invalid TOML inline-table cases owned by this feature.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 856
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  28 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true

## Build summary
RESULT: SUCCESS

FILES CHANGED:
- internal/toml/toml_test.go

SUMMARY:
Added regression tests for nested/commented arrays, trailing commas, dotted inline tables, closure conflicts, and invalid inline-table syntax. All gates pass:
- Full suite: 205 valid, 474 invalid
- Scoped array/inline suites: all pass
- `go test`, `gofmt`, and `go vet`: pass

BLOCKERS:
- None
