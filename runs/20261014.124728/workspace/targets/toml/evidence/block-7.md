# Evidence: Block 7 · Service (block-7)

- block type: block
- date: 2026-10-14
- resulting state: closed/verified
- story points (combined assembled cost): 8625
- execution id: 20261014.132405.837Z-8f9e01f1

## Stories built
- Arrays and inline tables, including inline-table closure. (parser-arrays-inline) [story]

## Acceptance tooling authorization
- FEATURE-Parser-Arrays-Inline.md#arrays-round-trip: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Parser-Arrays-Inline.md#inline-table-round-trip: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Parser-Arrays-Inline.md#inline-table-closed: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Parser-Arrays-Inline.md#arrays-inline-suite-slice: executable=sh; scope=test; authorization=existing Target environment
- FEATURE-Parser-Arrays-Inline.md#arrays-inline-suite-slice: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Parser-Arrays-Inline.md#arrays-inline-suite-slice: executable=toml-test; scope=test; authorization=existing Target environment

## Stacked context
- compass: COMPASS.md (SP 945)
- implements: FEATURE-Parser-Arrays-Inline.md (SP 1428)
- context: ARCHITECTURE_compact.md (SP 244)
- stack: go.md (SP 5918)

## Build directory changes
- cmd/toml-decoder/arrays_test.go
- internal/toml/array.go
- internal/toml/array_test.go
- internal/toml/inline.go
- internal/toml/inline_test.go
- internal/toml/keys.go
- internal/toml/string_test.go
- internal/toml/tables.go
- internal/toml/toml.go
- internal/toml/toml_test.go
- toml-decoder

## Pre-build acceptance observation
- RED: arrays-round-trip (FEATURE-Parser-Arrays-Inline.md)
  intent: An array of two integers encodes as a JSON array of tagged values.
  return code: 1
  stderr:
    toml-decoder: decode: invalid TOML: line 1: expected a value
    
    --- drydock: values at failure ---
      result = CompletedProcess(args=['./toml-decoder'], returncode=1, stdout='', stderr='toml-decoder: decode: invalid TOML: line 1: expected a value\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "arrays-round-trip.py", line 18, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: inline-table-round-trip (FEATURE-Parser-Arrays-Inline.md)
  intent: An inline table on one line encodes as a nested JSON object.
  return code: 1
  stderr:
    toml-decoder: decode: invalid TOML: line 1: expected a value
    
    --- drydock: values at failure ---
      result = CompletedProcess(args=['./toml-decoder'], returncode=1, stdout='', stderr='toml-decoder: decode: invalid TOML: line 1: expected a value\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "inline-table-round-trip.py", line 18, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- GREEN (prepassed): inline-table-closed (FEATURE-Parser-Arrays-Inline.md)
  intent: Adding a key to an inline table through a later table header is rejected.
  return code: 0
  stderr:
    toml-decoder: decode: invalid TOML: line 1: expected a value
- GREEN (prepassed): arrays-inline-suite-slice (FEATURE-Parser-Arrays-Inline.md)
  intent: The toml-test array and inline-table slices execute and every case passes.
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
- PASS: arrays-round-trip (FEATURE-Parser-Arrays-Inline.md)
  intent: An array of two integers encodes as a JSON array of tagged values.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: inline-table-round-trip (FEATURE-Parser-Arrays-Inline.md)
  intent: An inline table on one line encodes as a nested JSON object.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: inline-table-closed (FEATURE-Parser-Arrays-Inline.md)
  intent: Adding a key to an inline table through a later table header is rejected.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    toml-decoder: decode: invalid TOML: line 2: table "t" is already defined
- UNVERIFIED: arrays-inline-suite-slice (FEATURE-Parser-Arrays-Inline.md)
  intent: The toml-test array and inline-table slices execute and every case passes.
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

## Agent self-report (advisory)
The build agent declared a failure. This is advisory only; the programmatic acceptance above is the authority for this block's outcome.
- summary: Scoped conformance slice fails on valid/array/array, which contains datetime values this block does not implement.
- detail:
    The declared arrays-inline-suite-slice pattern uses `|`, which toml-test does not accept, so the declared invocation runs zero cases and passes vacuously. Re-run with comma-separated patterns: 37 of 38 valid array and inline-table cases pass, and all 54 invalid cases pass. The one failure needs datetime scalars (1987-07-05T17:45:00Z, 1979-05-27T07:32:00, 2006-06-01, 11:00:00), which are outside this block's stories. To rerun green, either implement datetime scalars in a feature that precedes this one, or have the owner scope valid/array/array out of this gate and fix the pattern separator in the criterion.

## Build summary
I'll start by inspecting the build directory and the staged sources.The scoped suite is `stage_test.sh`, and `toml-test` is installed. Next I'll read the existing code, since this block extends the parser built in earlier blocks.The parser has no array or inline-table support yet: `value()` falls through to `scalar()`, which rejects `[` and `{`. Before designing the change, I'll check the existing tests for anything that depends on that rejection or on the `map[string]any` value type, and look at the test style.Two earlier-block tests assert that arrays and inline tables are rejected (`toml_test.go:97`, `string_test.go:122-123`). Those encode the old "not yet supported" boundary, so they'll need to move to valid cases. I'll check their context and the test helpers.I have what I need. My plan:

- **Value model:** `value()` returns `any` so an array can hold a mix of scalars, nested arrays, and inline tables during parsing. Arrays export as `{"type":"array","value":[...]}`, which is the format COMPASS requires ("every TOML value is encoded as `{"type","value"}`") and toml-test checks for.
- **Closure:** add an `originInline` origin. Dotted keys already refuse non-dotted tables, and `enterHeader` already refuses anything not implicit. I need to add the check that `headerTable` refuses to walk through an inline table.

I'll start with the origin and the header check in `tables.go`.Next, `toml.go` and `keys.go`: `value()` and `keyValue()` now return `any`, and `insert` accepts `any`.Now the two new files: `array.go` for arrays and the shared gap-skipping, and `inline.go` for inline tables.Now I'll update the two earlier-block rejection lists. Valid arrays and inline tables move out of the rejection lists, and I'll add malformed forms of each in their place.Next, the tests. I'll write table-driven tests in the same style as the existing ones: one behavior per case, expected values written out, and the decoded tree compared with `reflect.DeepEqual`.I'll add one binary-level test for the stdout contract, matching the existing `tables_test.go` in `cmd`, and add a comment on the empty `case ']'` in `array()` so the close-bracket path reads clearly.Now the Go gates: `gofmt`, `go vet`, build, and the full test suite with the race detector.One failure, and it's a typo in my test input: `'y]''` leaves a stray quote. The decoder rejected it correctly. I'll fix the input to `'y]'`.Gates are green. Next is the declared scoped conformance slice, run through the invocation the acceptance criterion names (`stage_test.sh` with the array and inline-table pattern).The exit status is 0, but the run reports `0 passed, 0 failed` with 884 skipped, so the pattern matched nothing. A zero-case pass isn't evidence, so I need to see how `toml-test -run` parses patterns before I trust this.The help text settles the format question. It says "TOML arrays correspond to JSON arrays", and only scalars use the `{"type","value"}` wrapper. So my `{"type":"array"}` wrapper is wrong for this harness, and the declared acceptance criterion (a plain JSON array of tagged values) was right. I'll check the `-run` syntax too.Two findings change the plan:

1. **Arrays are plain JSON arrays in this harness.** My wrapper is wrong. I'll drop it, which also matches the acceptance criterion `arrays-round-trip`.
2. **The declared pattern selects nothing.** `toml-test -run` separates patterns with commas, and `|` isn't a separator. `filepath.Match` treats `|` as a literal, so the declared invocation runs 0 cases and exits 0. That is a vacuous pass. I'll run the same scoped entry point with the pattern in the syntax the tool accepts, and report the declared string as a defect.

I'll fix the export first.Now the Go gates again, then the scoped suite with the comma-separated pattern, and the declared pattern alongside it so the vacuous result is on record.The real slice now runs: 37 of 38 valid array and inline-table cases pass, and all 54 invalid cases pass. The one failure is `valid/array/array`, which contains datetime values (`1987-07-05T17:45:00Z`, `1979-05-27T07:32:00`, `2006-06-01`, `11:00:00`). Datetimes aren't in this block's scope and the scalar parser doesn't support them yet. I'll confirm that below.

Next I'll run the three non-suite acceptance criteria exactly as written. I'll put the scripts in a scratch directory outside the build tree.All three criteria pass as written. The fourth criterion (`arrays-inline-suite-slice`) is the one I'm reporting as blocked. Before writing the summary I'll confirm the build tree holds only my intended changes and that no staged source was touched.The array and inline-table support is in and the Go gates are green, but the build step fails its scoped conformance criterion. The remaining failure is datetime values inside an array, which this block doesn't cover.

**What's blocked:** `valid/array/array` contains datetimes (`1987-07-05T17:45:00Z`, `1979-05-27T07:32:00`, `2006-06-01`, `11:00:00`), and the scalar parser only accepts strings, booleans, integers, and floats. Datetimes aren't among this block's stories. Passing this criterion needs a datetime feature in a later block, or a decision from you to scope that case out of this gate. I did not change the staged suite.

**Criterion defect:** `arrays-inline-suite-slice` declares a `|`-separated pattern. `toml-test -run` separates patterns with commas, so the declared string matches 0 cases and exits 0. That is a vacuous pass, and I didn't count it. I ran the same `stage_test.sh` entry point with comma-separated patterns instead.

**Results**
- `gofmt`, `go vet`, and `go test -race -cover ./...` all pass (internal/toml 97.0%, cmd 57.1%).
- The three non-suite acceptance criteria pass as written: `arrays-round-trip`, `inline-table-round-trip`, and `inline-table-closed`.
- Scoped conformance, comma-separated: valid 37 passed / 1 failed (the datetime case above), invalid 54 passed / 0 failed.

**Choices I made**
- **Arrays encode as plain JSON arrays**, not `{"type":"array"}`. The toml-test v2.2.0 JSON description says "TOML arrays correspond to JSON arrays", and the `arrays-round-trip` criterion agrees. Scalars keep the tagged form.
- **Closure:** a new `originInline` origin. Dotted keys and headers can't extend an inline table, and a header can't walk through one.
- **Two earlier-block tests** asserted that arrays and inline tables are rejected (`toml_test.go`, `string_test.go`). Those are now valid cases, and I replaced them with malformed array and inline-table forms.

**Files changed**
- `internal/toml/tables.go`, `toml.go`, `keys.go` (modified)
- `internal/toml/array.go`, `inline.go` (new)
- `internal/toml/array_test.go`, `inline_test.go` (new)
- `internal/toml/toml_test.go`, `string_test.go` (modified)
- `cmd/toml-decoder/arrays_test.go` (new)
- `toml-decoder` (binary rebuilt by the stage script; it was already modified before I started)

`sources/` is untouched, and I did not commit.

<blueprint-decisions>
[{"spec":"FEATURE-Parser-Arrays-Inline.md","severity":"Low","subject":"Array encoding","decision":"Options were the tagged {\"type\":\"array\",\"value\":[...]} form and a plain JSON array. I implemented a plain JSON array, as toml-test v2.2.0's JSON description and the arrays-round-trip criterion specify. Tagged form is for scalars only."},{"spec":"FEATURE-Parser-Arrays-Inline.md","severity":"Low","subject":"Earlier rejection tests","decision":"Two earlier-block tests asserted that arrays and inline tables are rejected. I moved the valid forms to positive tests and replaced them with malformed-array and malformed-inline-table cases. Confirm this is acceptable."}]
</blueprint-decisions>

RESULT: FAILED

FILES CHANGED:
- internal/toml/tables.go
- internal/toml/toml.go
- internal/toml/keys.go
- internal/toml/array.go
- internal/toml/inline.go
- internal/toml/array_test.go
- internal/toml/inline_test.go
- internal/toml/toml_test.go
- internal/toml/string_test.go
- cmd/toml-decoder/arrays_test.go
- toml-decoder

SUMMARY:
Implemented arrays (mixed types, nesting, trailing comma, multi-line with comments and CRLF) and inline tables (single line, no trailing comma, dotted keys inside). Inline tables are closed via a new `originInline` origin: headers and dotted keys cannot extend them. Arrays export as plain JSON arrays. Go gates are green. Non-suite acceptance criteria `arrays-round-trip`, `inline-table-round-trip`, and `inline-table-closed` pass. Scoped conformance with comma-separated patterns: valid 37/38, invalid 54/54. The failing case, `valid/array/array`, needs datetime scalars.

BLOCKERS:
- `valid/array/array` requires datetime values, which the scalar parser does not support. Datetimes belong to a later feature.
- The declared `arrays-inline-suite-slice` pattern uses `|`, which `toml-test -run` does not support. It runs 0 cases and exits 0, so it passes vacuously.

FAILURE_SUMMARY: Scoped conformance slice fails on valid/array/array, which contains datetime values this block does not implement.
FAILURE_DETAIL: The declared arrays-inline-suite-slice pattern uses `|`, which toml-test does not accept, so the declared invocation runs zero cases and passes vacuously. Re-run with comma-separated patterns: 37 of 38 valid array and inline-table cases pass, and all 54 invalid cases pass. The one failure needs datetime scalars (1987-07-05T17:45:00Z, 1979-05-27T07:32:00, 2006-06-01, 11:00:00), which are outside this block's stories. To rerun green, either implement datetime scalars in a feature that precedes this one, or have the owner scope valid/array/array out of this gate and fix the pattern separator in the criterion.
