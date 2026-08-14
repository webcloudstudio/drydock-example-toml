# Evidence: Block 4 · Service (block-4)

- block type: block
- date: 2026-08-14
- resulting state: closed/verified
- story points (combined assembled cost): 8516
- execution id: 20260814.051752.394Z-723f1677

## Stories built
- Encode parsed TOML values using the required tagged JSON representation. (decoder-tagged-json) [story]

## Stacked context
- compass: COMPASS.md (SP 2629)
- implements: FEATURE-Tagged-JSON.md (SP 886)
- context: ARCHITECTURE_compact.md (SP 108)
- stack: go.md (SP 3646)
- stack: common_compact.md (SP 1179)

## Build directory changes
- internal/toml/value_test.go

## Pre-build acceptance observation
- GREEN (prepassed): tagged-scalars (FEATURE-Tagged-JSON.md)
  intent: Scalar TOML values are emitted with type names and string payloads.
  return code: 0
- GREEN (prepassed): tagged-structures (FEATURE-Tagged-JSON.md)
  intent: Tables and arrays retain structural JSON forms while their scalar members remain tagged.
  return code: 0
- GREEN (prepassed): tagged-empty-containers (FEATURE-Tagged-JSON.md)
  intent: Empty arrays and tables encode as empty JSON containers.
  return code: 0
- GREEN (prepassed): tagged-datetime-types (FEATURE-Tagged-JSON.md)
  intent: Date and time values use the distinct required tagged type names.
  return code: 0

## Post-build programmatic acceptance
- PASS: tagged-scalars (FEATURE-Tagged-JSON.md)
  intent: Scalar TOML values are emitted with type names and string payloads.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
- PASS: tagged-structures (FEATURE-Tagged-JSON.md)
  intent: Tables and arrays retain structural JSON forms while their scalar members remain tagged.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
- PASS: tagged-empty-containers (FEATURE-Tagged-JSON.md)
  intent: Empty arrays and tables encode as empty JSON containers.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
- PASS: tagged-datetime-types (FEATURE-Tagged-JSON.md)
  intent: Date and time values use the distinct required tagged type names.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0

## Agent self-report (advisory)
The build agent declared a failure. This is advisory only; the programmatic acceptance above is the authority for this block's outcome.
- summary: Full conformance suite remains failing outside tagged-JSON coverage.
- detail:
    The tagged-JSON acceptance checks pass, but the required full suite reports parser failures in strings, keys, datetimes, floats, tables, and invalid-input handling. Those behaviors require work in another build block before rerunning.

## Build summary
RESULT: FAILED

FILES CHANGED:
- internal/toml/value_test.go

SUMMARY:
Added coverage for all tagged scalar types, string payloads, nested structures, and empty containers. Focused tests, acceptance checks, `go test`, `go vet`, formatting, and build pass.

BLOCKERS:
- Full `toml-test v2.2.0` conformance fails due to pre-existing parser behavior: 47 valid and 129 invalid cases failed.

FAILURE_SUMMARY: Full conformance suite remains failing outside tagged-JSON coverage.
FAILURE_DETAIL: The tagged-JSON acceptance checks pass, but the required full suite reports parser failures in strings, keys, datetimes, floats, tables, and invalid-input handling. Those behaviors require work in another build block before rerunning.
