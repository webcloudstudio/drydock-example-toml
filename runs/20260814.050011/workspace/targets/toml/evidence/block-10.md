# Evidence: Block 10 · Service (block-10)

- block type: block
- date: 2026-08-14
- resulting state: closed/verified
- story points (combined assembled cost): 8690
- execution id: 20260814.054821.850Z-b97cfb83

## Stories built
- Enforce closure of TOML inline tables. (structures-inline-table-closure) [story]
- Parse standard TOML table headers and implicit tables. (tables-standard) [story]

## Missing context
- context: toml-v1.0.0.md

## Stacked context
- compass: COMPASS.md (SP 2629)
- implements: FEATURE-Inline-Table-Closure.md (SP 432)
- context: stage_test.sh (SP 161)
- context: ARCHITECTURE_compact.md (SP 108)
- stack: go.md (SP 3646)
- stack: common_compact.md (SP 1179)
- implements: FEATURE-Standard-Tables.md (SP 439)

## Build directory changes
- internal/toml/decode.go
- internal/toml/value_test.go
- toml-decoder

## Pre-build acceptance observation
- GREEN (prepassed): inline-closure-invalid (FEATURE-Inline-Table-Closure.md)
  intent: The implementation rejects authoritative invalid cases that extend or redefine inline tables.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 856
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  28 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- GREEN (prepassed): inline-closure-boundary (FEATURE-Inline-Table-Closure.md)
  intent: The implementation passes the authoritative valid inline-table cases while preserving inline-table closure semantics.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 867
      valid tests:  17 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- RED: standard-tables-valid (FEATURE-Standard-Tables.md)
  intent: The implementation passes the authoritative valid TOML standard-table conformance slice.
  return code: 1
  stdout:
    FAIL valid/table/array-nest
         line 4: table conflict
     
         Exit 1
    
         input sent to parser-cmd (PID 838453):
            1 │[0m [[albums]]
            2 │[0m name = "Born to Run"
            3 │[0m
            4 │[0m   [[albums.songs]]
            5 │[0m   name = "Jungleland"
            6 │[0m
            7 │[0m   [[albums.songs]]
            8 │[0m   name = "Meeting Across the River"
            9 │[0m
           10 │[0m [[albums]]
           11 │[0m name = "Born in the USA"
           12 │[0m
           13 │[0m   [[albums.songs]]
           14 │[0m   name = "Glory Days"
           15 │[0m
           16 │[0m   [[albums.songs]]
           17 │[0m   name = "Dancing in the Dark"
    
         output from parser-cmd (PID 838453) (stderr):
            1 │[0m line 4: table conflict
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/table/array-table-array
         line 2: table conflict
     
         Exit 1
    
         input sent to parser-cmd (PID 838488):
            1 │[0m [[a]]
            2 │[0m     [[a.b]]
            3 │[0m         [a.b.c]
            4 │[0m             d = "val0"
            5 │[0m     [[a.b]]
            6 │[0m         [a.b.c]
            7 │[0m             d = "val1"
    
         output from parser-cmd (PID 838488) (stderr):
            1 │[0m line 2: table conflict
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/table/names
         line 2: table redefinition
     
         Exit 1
    
         input sent to parser-cmd (PID 838454):
            1 │[0m [a.b.c]
            2 │[0m [a."b.c"]
            3 │[0m [a.'d.e']
            4 │[0m [a.' x ']
            5 │[0m [ d.e.f ]
            6 │[0m [ g . h . i ]
            7 │[0m [ j . "ʞ" . 'l' ]
            8 │[0m
            9 │[0m [x.1.2]
    
         output from parser-cmd (PID 838454) (stderr):
            1 │[0m line 2: table redefinition
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/table/names-with-values
         line 4: table redefinition
     
         Exit 1
    
         input sent to parser-cmd (PID 838484):
            1 │[0m [a.b.c]
            2 │[0m key = 1
            3 │[0m
            4 │[0m [a."b.c"]
            5 │[0m key = 2
            6 │[0m
            7 │[0m [a.'d.e']
            8 │[0m key = 3
            9 │[0m
           10 │[0m [a.' x ']
           11 │[0m key = 4
           12 │[0m
           13 │[0m [ d.e.f ]
           14 │[0m key = 5
           15 │[0m
           16 │[0m [ g . h . i ]
           17 │[0m key = 6
           18 │[0m
           19 │[0m [ j . "ʞ" . 'l' ]
           20 │[0m key = 7
           21 │[0m
           22 │[0m [x.1.2]
           23 │[0m key = 8
    
         output from parser-cmd (PID 838484) (stderr):
            1 │[0m line 4: table redefinition
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 860
      valid tests:  20 passed,  4 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
  stderr:
    Traceback (most recent call last):
      File "standard-tables-valid.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: standard-tables-invalid (FEATURE-Standard-Tables.md)
  intent: The implementation rejects invalid standard-table header and hierarchy cases in the authoritative suite.
  return code: 1
  stdout:
    FAIL invalid/table/append-with-dotted-keys-01
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 838715):
            1 │[0m # First a.b.c defines a table: a.b.c = {z=9}
            2 │[0m #
            3 │[0m # Then we define a.b.c.t = "str" to add a str to the above table, making it:
            4 │[0m #
            5 │[0m #   a.b.c = {z=9, t="..."}
            6 │[0m #
            7 │[0m # While this makes sense, logically, it was decided this is not valid TOML as
            8 │[0m # it's too confusing/convoluted.
            9 │[0m #
           10 │[0m # See: https://github.com/toml-lang/toml/issues/846
           11 │[0m #      https://github.com/toml-lang/toml/pull/859
           12 │[0m
           13 │[0m [a.b.c]
           14 │[0m   z = 9
           15 │[0m
           16 │[0m [a]
           17 │[0m   b.c.t = "Using dotted keys to add to [a.b.c] after explicitly defining it above is not allowed"
    
         output from parser-cmd (PID 838715) (stdout):
            1 │[0m {
            2 │[0m   "a": {
            3 │[0m     "b": {
            4 │[0m       "c": {
            5 │[0m         "z": {"type": "integer", "value": "9"},
            6 │[0m         "t": {
            7 │[0m           "type":  "string",
            8 │[0m           "value": "Using dotted keys to add to [a.b.c] after explicitly defining it above is not allowed"
            9 │[0m         }
           10 │[0m       }
           11 │[0m     }
           12 │[0m   }
           13 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/table/append-with-dotted-keys-02
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 838750):
            1 │[0m # This is the same issue as in injection-1.toml, except that nests one level
            2 │[0m # deeper. See that file for a more complete description.
            3 │[0m
            4 │[0m [a.b.c.d]
            5 │[0m   z = 9
            6 │[0m
            7 │[0m [a]
            8 │[0m   b.c.d.k.t = "Using dotted keys to add to [a.b.c.d] after explicitly defining it above is not allowed"
    
         output from parser-cmd (PID 838750) (stdout):
            1 │[0m {
            2 │[0m   "a": {
            3 │[0m     "b": {
            4 │[0m       "c": {
            5 │[0m         "d": {
            6 │[0m           "z": {"type": "integer", "value": "9"},
            7 │[0m           "k": {
            8 │[0m             "t": {
            9 │[0m               "type":  "string",
           10 │[0m               "value": "Using dotted keys to add to [a.b.c.d] after explicitly defining it above is not allowed"
           11 │[0m             }
           12 │[0m           }
           13 │[0m         }
           14 │[0m       }
           15 │[0m     }
           16 │[0m   }
           17 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/table/append-with-dotted-keys-08
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 838849):
            1 │[0m [a.b.c]
            2 │[0m z = 9
            3 │[0m
            4 │[0m [[totally_unrelated]]
            5 │[0m x = 123
            6 │[0m
            7 │[0m [a]
            8 │[0m b.c.t = "this should NOT be allowed"
    
         output from parser-cmd (PID 838849) (stdout):
            1 │[0m {
            2 │[0m   "totally_unrelated": [{
            3 │[0m     "x": {"type": "integer", "value": "123"}
            4 │[0m   }],
            5 │[0m   "a": {
            6 │[0m     "b": {
            7 │[0m       "c": {
            8 │[0m         "t": {"type": "string", "value": "this should NOT be allowed"},
            9 │[0m         "z": {"type": "integer", "value": "9"}
           10 │[0m       }
           11 │[0m     }
           12 │[0m   }
           13 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/table/array-implicit
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 838789):
            1 │[0m # This test is a bit tricky. It should fail because the first use of
            2 │[0m # `[[albums.songs]]` without first declaring `albums` implies that `albums`
            3 │[0m # must be a table. The alternative would be quite weird. Namely, it wouldn't
            4 │[0m # comply with the TOML spec: "Each double-bracketed sub-table will belong to
            5 │[0m # the most *recently* defined table element *above* it."
            6 │[0m #
            7 │[0m # This is in contrast to the *valid* test, table-array-implicit where
            8 │[0m # `[[albums.songs]]` works by itself, so long as `[[albums]]` isn't declared
            9 │[0m # later. (Although, `[albums]` could be.)
           10 │[0m [[albums.songs]]
           11 │[0m name = "Glory Days"
           12 │[0m
           13 │[0m [[albums]]
           14 │[0m name = "Born in the USA"
    
         output from parser-cmd (PID 838789) (stdout):
            1 │[0m {
            2 │[0m   "albums": [{
            3 │[0m     "name": {"type": "string", "value": "Born in the USA"}
            4 │[0m   }]
            5 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/table/duplicate-key-06
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 838719):
            1 │[0m [tbl]
            2 │[0m [[tbl]]
    
         output from parser-cmd (PID 838719) (stdout):
            1 │[0m {
            2 │[0m   "tbl": [{}]
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/table/newline-01
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 838922):
            1 │[0m [tbl
            2 │[0m ]
            3 │[0m k = 1
    
         output from parser-cmd (PID 838922) (stdout):
            1 │[0m {
            2 │[0m   "tbl": {
            3 │[0m     "k": {"type": "integer", "value": "1"}
            4 │[0m   }
            5 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/table/newline-03
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 838918):
            1 │[0m ["tbl"
            2 │[0m ]
            3 │[0m k = 1
    
         output from parser-cmd (PID 838918) (stdout):
            1 │[0m {
            2 │[0m   "tbl": {
            3 │[0m     "k": {"type": "integer", "value": "1"}
            4 │[0m   }
            5 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/table/newline-05
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 838919):
            1 │[0m [tbl
            2 │[0m .sub]
            3 │[0m k = 1
    
         output from parser-cmd (PID 838919) (stdout):
            1 │[0m {
            2 │[0m   "tbl": {
            3 │[0m     "sub": {
            4 │[0m       "k": {"type": "integer", "value": "1"}
            5 │[0m     }
            6 │[0m   }
            7 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/table/overwrite-bool-with-array
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 838988):
            1 │[0m a=true
            2 │[0m [[a]]
    
         output from parser-cmd (PID 838988) (stdout):
            1 │[0m {
            2 │[0m   "a": [{}]
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/table/redefine-02
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 838943):
            1 │[0m [t1]
            2 │[0m t2.t3.v = 0
            3 │[0m [t1.t2]
    
         output from parser-cmd (PID 838943) (stdout):
            1 │[0m {
            2 │[0m   "t1": {
            3 │[0m     "t2": {
            4 │[0m       "t3": {
            5 │[0m         "v": {"type": "integer", "value": "0"}
            6 │[0m       }
            7 │[0m     }
            8 │[0m   }
            9 │[0m }
    
         want:
           Exit code 1
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 822
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  52 passed, 10 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
  stderr:
    Traceback (most recent call last):
      File "standard-tables-invalid.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- PASS: inline-closure-invalid (FEATURE-Inline-Table-Closure.md)
  intent: The implementation rejects authoritative invalid cases that extend or redefine inline tables.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 856
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  28 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: inline-closure-boundary (FEATURE-Inline-Table-Closure.md)
  intent: The implementation passes the authoritative valid inline-table cases while preserving inline-table closure semantics.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 867
      valid tests:  17 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: standard-tables-valid (FEATURE-Standard-Tables.md)
  intent: The implementation passes the authoritative valid TOML standard-table conformance slice.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 860
      valid tests:  24 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: standard-tables-invalid (FEATURE-Standard-Tables.md)
  intent: The implementation rejects invalid standard-table header and hierarchy cases in the authoritative suite.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 822
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  62 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true

## Build summary
RESULT: SUCCESS

FILES CHANGED:
- internal/toml/decode.go
- internal/toml/value_test.go

SUMMARY:
Implemented standard tables, implicit/nested table resolution, array-of-table contexts, dotted-key provenance, and inline-table closure enforcement. Added focused tests.

Scoped suites passed:
- Tables: 24 valid, 62 invalid
- Inline tables: 17 valid, 28 invalid
- Go tests passed

BLOCKERS:
- Full suite retains unrelated pre-existing failures in multiline strings, numeric validation, and integer-sign handling.
