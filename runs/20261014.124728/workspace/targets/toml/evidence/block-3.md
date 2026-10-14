# Evidence: Block 3 · Service (block-3)

- block type: block
- date: 2026-10-14
- resulting state: closed/verified
- story points (combined assembled cost): 8271
- execution id: 20261014.130041.922Z-b7ed98fe

## Stories built
- Booleans, integers, and floats. (parser-lexical-scalars) [story]

## Acceptance tooling authorization
- FEATURE-Parser-Lexical-Scalars.md#scalars-integer-round-trip: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Parser-Lexical-Scalars.md#scalars-integer-overflow: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Parser-Lexical-Scalars.md#scalars-suite-slice: executable=sh; scope=test; authorization=existing Target environment
- FEATURE-Parser-Lexical-Scalars.md#scalars-suite-slice: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Parser-Lexical-Scalars.md#scalars-suite-slice: executable=toml-test; scope=test; authorization=existing Target environment

## Stacked context
- compass: COMPASS.md (SP 945)
- implements: FEATURE-Parser-Lexical-Scalars.md (SP 1085)
- context: ARCHITECTURE_compact.md (SP 244)
- stack: go.md (SP 5918)

## Build directory changes
- internal/toml/scalar.go
- internal/toml/scalar_test.go
- internal/toml/string_test.go
- internal/toml/toml.go
- internal/toml/toml_test.go
- toml-decoder

## Pre-build acceptance observation
- RED: scalars-integer-round-trip (FEATURE-Parser-Lexical-Scalars.md)
  intent: The largest signed 64-bit integer decodes losslessly as a tagged integer.
  return code: 1
  stderr:
    toml-decoder: decode: invalid TOML: line 1: expected a string value
    
    --- drydock: values at failure ---
      result = CompletedProcess(args=['./toml-decoder'], returncode=1, stdout='', stderr='toml-decoder: decode: invalid TOML: line 1: expected a string value\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "scalars-integer-round-trip.py", line 17, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- GREEN (prepassed): scalars-integer-overflow (FEATURE-Parser-Lexical-Scalars.md)
  intent: An integer one past the signed 64-bit maximum is rejected.
  return code: 0
  stderr:
    toml-decoder: decode: invalid TOML: line 1: expected a string value
- GREEN (prepassed): scalars-suite-slice (FEATURE-Parser-Lexical-Scalars.md)
  intent: The toml-test slices for booleans, integers, and floats execute and every case passes.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20261014.124728/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true

## Post-build programmatic acceptance
- PASS: scalars-integer-round-trip (FEATURE-Parser-Lexical-Scalars.md)
  intent: The largest signed 64-bit integer decodes losslessly as a tagged integer.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: scalars-integer-overflow (FEATURE-Parser-Lexical-Scalars.md)
  intent: An integer one past the signed 64-bit maximum is rejected.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    toml-decoder: decode: invalid TOML: line 1: integer 9223372036854775808 is outside the signed 64-bit range
- UNVERIFIED: scalars-suite-slice (FEATURE-Parser-Lexical-Scalars.md)
  intent: The toml-test slices for booleans, integers, and floats execute and every case passes.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  error: unverified acceptance: selector matched zero cases
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20261014.124728/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true

## Build summary
The existing tests reject `n = 1` and `b = true` as "not decoded yet," which this step changes, so I'll update those specific cases. Let me check the harness version and the test-only baseline first.Baseline is green and the harness is v2.2.0. Now I'll add the scalar decoder.Now I'll route non-string values through the scalar path and update the stale doc comments.Next, I'll remove the now-wrong rejection cases from the existing tests and write the scalar tests.Two of those tests are weak: the NaN test checks nothing real, and the overflow test asserts on error text. I'll remove both.That failure is a mistake in my test, not the code. `0x8000000000000000` is 2^63, outside signed 64-bit, and it's already in the reject list. I'll drop the bogus accept case.Unit tests, vet, and gofmt are clean, with coverage of 95.7% for the decoder package. Next I'll run the two integer acceptance checks as written in the spec.Both integer checks pass. Now the scoped conformance slice, using the exact invocation the spec declares.The harness exited 0, but it ran zero cases: "0 passed, 0 failed" with 884 skipped. That is a vacuous pass, and I won't report it as success. I'll check how toml-test interprets `-run` patterns.The declared pattern is the problem. `-run` takes comma-separated globs, and `|` is not a separator, so the spec's pattern matches nothing, and the harness reports success on zero cases. I'll rerun the same staged script with the equivalent comma-separated pattern to get a real verdict.RESULT: SUCCESS

Booleans, integers, and floats now decode in `internal/toml`. Both integer acceptance checks pass. The suite slice passes, but the pattern written in the spec selects nothing, so I ran the same script with a corrected pattern.

**Acceptance results**
- `scalars-integer-round-trip`: passes. `9223372036854775807` decodes to `{"type":"integer","value":"9223372036854775807"}`.
- `scalars-integer-overflow`: passes. `9223372036854775808` exits 1 with empty stdout.
- `scalars-suite-slice`: the criterion's exit-0 check passes, but the spec's pattern is wrong. `-run` takes comma-separated globs, and `|` is not a separator. With the spec's pattern the harness ran 0 cases and reported "0 passed, 0 failed" with 884 skipped, which is a vacuous green. With the same script and `valid/bool/*,valid/integer/*,valid/float/*,invalid/bool/*,invalid/integer/*,invalid/float/*`, the real verdict is 14 valid passed, 96 invalid passed, 0 failed, exit 0. The spec pattern needs fixing before this gate means anything.

**Other checks**
- `gofmt -l` is clean, `go vet` is clean, and `go test -race -cover ./...` passes. Coverage of `internal/toml` is 95.7%.
- I did not run `full_test.sh`, since it is not declared for this step.

**Choices I made**
- Scalars are read as a bare token and must match one full TOML grammar, so partial matches are rejected rather than truncated. Prefixed integers take no sign, and underscores appear only between digits.
- Integers are emitted as decimal strings after an int64 range check. Hex, octal, and binary inputs are converted to decimal.
- Floats are emitted with `strconv.FormatFloat(f, 'g', -1, 64)`, and special values as `nan`, `inf`, and `-inf`. Out-of-range floats are rejected. The suite passed with this format, so the format is accepted by the harness.
- I removed the existing "integer" and "boolean" rejection cases from `toml_test.go` and `string_test.go`, because this step makes those valid.

**Files changed**
- `internal/toml/scalar.go` (new)
- `internal/toml/scalar_test.go` (new)
- `internal/toml/toml.go` (routes non-string values to the scalar path; doc comments updated)
- `internal/toml/toml_test.go` (removed two stale rejection cases)
- `internal/toml/string_test.go` (removed two stale rejection cases)
- `toml-decoder` (rebuilt binary; it was already modified in the working tree)

<blueprint-decisions>
[{"spec":"FEATURE-Parser-Lexical-Scalars.md","severity":"Low","subject":"Float output format","decision":"Finite floats are emitted with strconv 'g' shortest form (for example 1e6 becomes 1e+06). The spec does not fix a float string format, and the staged suite accepted this format. Is that acceptable, or should floats use a fixed notation?"},{"spec":"FEATURE-Parser-Lexical-Scalars.md","severity":"Material","subject":"Scoped suite pattern","decision":"The declared pattern uses '|' separators, which toml-test does not support. It selects 0 cases and exits 0. I verified the slice with the comma-separated equivalent (110 cases, all passing). The criterion's pattern should be corrected on replan."}]
</blueprint-decisions>

FILES CHANGED:
- internal/toml/scalar.go
- internal/toml/scalar_test.go
- internal/toml/toml.go
- internal/toml/toml_test.go
- internal/toml/string_test.go
- toml-decoder

SUMMARY:
Implemented booleans, integers (signed 64-bit, lossless, with 0x/0o/0b prefixes and underscores), and floats (binary64, inf, nan, exponents) in the `internal/toml` decoder. Out-of-range integers and malformed float forms are rejected with ErrSyntax. Acceptance checks: integer round-trip passes, integer overflow passes, and the scoped suite slice passes 110/110 with the corrected pattern (14 valid, 96 invalid; the declared pattern itself selects 0 cases). gofmt, go vet, and go test -race pass.

BLOCKERS:
- The `scalars-suite-slice` pattern in the spec uses `|`, which toml-test does not parse, so the declared invocation runs zero cases and reports a false green. It needs correcting on replan.
