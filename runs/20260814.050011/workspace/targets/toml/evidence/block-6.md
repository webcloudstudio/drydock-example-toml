# Evidence: Block 6 · Service (block-6)

- block type: block
- date: 2026-08-14
- resulting state: closed/verified
- story points (combined assembled cost): 8955
- execution id: 20260814.052805.919Z-3c07ffab

## Stories built
- Parse TOML integers, floating-point values, special floats, and booleans. (lexical-numbers) [story]
- Parse offset datetime, local datetime, local date, and local time values. (lexical-datetime) [story]

## Missing context
- context: toml-v1.0.0.md

## Stacked context
- compass: COMPASS.md (SP 2629)
- implements: FEATURE-Lexical-Numbers.md (SP 612)
- context: stage_test.sh (SP 161)
- context: ARCHITECTURE_compact.md (SP 108)
- stack: go.md (SP 3646)
- stack: common_compact.md (SP 1179)
- implements: FEATURE-Lexical-Datetime.md (SP 491)

## Build directory changes
- internal/toml/decode.go
- internal/toml/value_test.go
- toml-decoder

## Pre-build acceptance observation
- RED: lexical-numbers-integer (FEATURE-Lexical-Numbers.md)
  intent: The implementation passes the authoritative TOML integer conformance slice.
  return code: 1
  stdout:
    FAIL valid/integer/long
         Values for key "int64-max-neg" don't match:
           Expected:     -9223372036854775808
           Your encoder: -9223372036854775807
    
         input sent to parser-cmd (PID 772794):
            1 │[0m # int64 "should" be supported, but is not mandatory. It's fine to skip this
            2 │[0m # test.
            3 │[0m int64-max     = 9223372036854775807
            4 │[0m int64-max-neg = -9223372036854775808
    
         output from parser-cmd (PID 772794) (stdout):
            1 │[0m {
            2 │[0m   "int64-max":     {"type": "integer", "value": "9223372036854775807"},
            3 │[0m   "int64-max-neg": {"type": "integer", "value": "-9223372036854775807"}
            4 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "int64-max":     {"type": "integer", "value": "9223372036854775807"},
            3 │[0m     "int64-max-neg": {"type": "integer", "value": "-9223372036854775808"}
            4 │[0m }
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 878
      valid tests:   5 passed,  1 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
    
    Traceback (most recent call last):
      File "lexical-numbers-integer.py", line 12, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: lexical-numbers-float (FEATURE-Lexical-Numbers.md)
  intent: The implementation passes the authoritative TOML float conformance slice.
  return code: 1
  stdout:
    FAIL valid/float/inf-and-nan
         Values for key "infinity_neg" don't match:
           Expected:     -Inf
           Your encoder: +Inf
    
         input sent to parser-cmd (PID 773011):
            1 │[0m # We don't encode +nan and -nan back with the signs; many languages don't
            2 │[0m # support a sign on NaN (it doesn't really make much sense).
            3 │[0m nan = nan
            4 │[0m nan_neg = -nan
            5 │[0m nan_plus = +nan
            6 │[0m infinity = inf
            7 │[0m infinity_neg = -inf
            8 │[0m infinity_plus = +inf
    
         output from parser-cmd (PID 773011) (stdout):
            1 │[0m {
            2 │[0m   "infinity":      {"type": "float", "value": "inf"},
            3 │[0m   "infinity_neg":  {"type": "float", "value": "inf"},
            4 │[0m   "infinity_plus": {"type": "float", "value": "inf"},
            5 │[0m   "nan":           {"type": "float", "value": "nan"},
            6 │[0m   "nan_neg":       {"type": "float", "value": "nan"},
            7 │[0m   "nan_plus":      {"type": "float", "value": "nan"}
            8 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "infinity":      {"type": "float", "value": "inf"},
            3 │[0m     "infinity_neg":  {"type": "float", "value": "-inf"},
            4 │[0m     "infinity_plus": {"type": "float", "value": "inf"},
            5 │[0m     "nan":           {"type": "float", "value": "nan"},
            6 │[0m     "nan_neg":       {"type": "float", "value": "nan"},
            7 │[0m     "nan_plus":      {"type": "float", "value": "nan"}
            8 │[0m }
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 877
      valid tests:   6 passed,  1 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
    
    Traceback (most recent call last):
      File "lexical-numbers-float.py", line 13, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- GREEN (prepassed): lexical-numbers-bool (FEATURE-Lexical-Numbers.md)
  intent: The implementation passes the authoritative TOML boolean conformance slice.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 883
      valid tests:   1 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- RED: lexical-datetime-valid (FEATURE-Lexical-Datetime.md)
  intent: The implementation passes the authoritative valid TOML datetime conformance slice.
  return code: 1
  stdout:
    FAIL valid/datetime/datetime
         line 4: invalid value "1987-07-05t17:45:00z"
     
         Exit 1
    
         input sent to parser-cmd (PID 773414):
            1 │[0m space = 1987-07-05 17:45:00Z
            2 │[0m
            3 │[0m # ABNF is case-insensitive, both "Z" and "z" must be supported.
            4 │[0m lower = 1987-07-05t17:45:00z
    
         output from parser-cmd (PID 773414) (stderr):
            1 │[0m line 4: invalid value "1987-07-05t17:45:00z"
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 875
      valid tests:   8 passed,  1 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
    
    Traceback (most recent call last):
      File "lexical-datetime-valid.py", line 13, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: lexical-datetime-invalid (FEATURE-Lexical-Datetime.md)
  intent: The implementation rejects the authoritative invalid TOML datetime conformance slice.
  return code: 1
  stdout:
    FAIL invalid/datetime/offset-overflow-minute
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 773712):
            1 │[0m d = 1985-06-18 17:04:07+12:60
    
         output from parser-cmd (PID 773712) (stdout):
            1 │[0m {
            2 │[0m   "d": {"type": "datetime", "value": "1985-06-18 17:04:07+12:60"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/datetime/time-no-leads
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 773721):
            1 │[0m # Leading 0 is always required.
            2 │[0m d = 2023-10-01T1:32:00Z
    
         output from parser-cmd (PID 773721) (stdout):
            1 │[0m {
            2 │[0m   "d": {"type": "datetime", "value": "2023-10-01T1:32:00Z"}
            3 │[0m }
    
         want:
           Exit code 1
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 849
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  33 passed,  2 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
    
    Traceback (most recent call last):
      File "lexical-datetime-invalid.py", line 13, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- PASS: lexical-numbers-integer (FEATURE-Lexical-Numbers.md)
  intent: The implementation passes the authoritative TOML integer conformance slice.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 878
      valid tests:   6 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: lexical-numbers-float (FEATURE-Lexical-Numbers.md)
  intent: The implementation passes the authoritative TOML float conformance slice.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 877
      valid tests:   7 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: lexical-numbers-bool (FEATURE-Lexical-Numbers.md)
  intent: The implementation passes the authoritative TOML boolean conformance slice.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 883
      valid tests:   1 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: lexical-datetime-valid (FEATURE-Lexical-Datetime.md)
  intent: The implementation passes the authoritative valid TOML datetime conformance slice.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 875
      valid tests:   9 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: lexical-datetime-invalid (FEATURE-Lexical-Datetime.md)
  intent: The implementation rejects the authoritative invalid TOML datetime conformance slice.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 849
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  35 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true

## Build summary
RESULT: SUCCESS

FILES CHANGED:
- internal/toml/decode.go
- internal/toml/value_test.go

SUMMARY:
Implemented strict TOML numeric, boolean, and datetime lexing with lossless int64 handling, special floats, separators, fractional truncation, and validation. All declared acceptance slices, tests, and `go vet` pass.

BLOCKERS:
- Full suite still has failures in out-of-scope key, table, string, and inline-table behavior.
