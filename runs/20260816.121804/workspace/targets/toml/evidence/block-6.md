# Evidence: Block 6 · Service (block-6)

- block type: block
- date: 2026-08-16
- resulting state: closed/verified
- story points (combined assembled cost): 8046
- execution id: 20260816.124810.803Z-1cf982a0

## Stories built
- Parse standard tables and arrays of tables with ordering and conflict rules. (parser-tables) [story]
- Parse and format TOML offset and local date/time values. (parser-datetimes) [story]

## Acceptance tooling authorization
- FEATURE-Tables.md#tables-valid: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Tables.md#tables-invalid: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Datetimes.md#datetimes-valid: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Datetimes.md#datetimes-invalid: executable=go; scope=test; authorization=existing Target environment

## Missing context
- context: toml-v1.0.0.md

## Stacked context
- compass: COMPASS.md (SP 2942)
- implements: FEATURE-Tables.md (SP 503)
- context: stage_test.sh (SP 161)
- context: ARCHITECTURE_compact.md (SP 123)
- stack: go.md (SP 3646)
- implements: FEATURE-Datetimes.md (SP 560)

## Build directory changes
- internal/toml/toml.go
- internal/toml/toml_test.go
- toml-decoder

## Pre-build acceptance observation
- RED: tables-valid (FEATURE-Tables.md)
  intent: The decoder passes all authoritative valid TOML table and array-of-table cases owned by this feature.
  return code: 1
  stdout:
    FAIL valid/table/array-nest
         line 4: invalid table path
     
         Exit 1
    
         input sent to parser-cmd (PID 1725827):
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
    
         output from parser-cmd (PID 1725827) (stderr):
            1 │[0m line 4: invalid table path
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/table/array-table-array
         line 2: invalid table path
     
         Exit 1
    
         input sent to parser-cmd (PID 1725821):
            1 │[0m [[a]]
            2 │[0m     [[a.b]]
            3 │[0m         [a.b.c]
            4 │[0m             d = "val0"
            5 │[0m     [[a.b]]
            6 │[0m         [a.b.c]
            7 │[0m             d = "val1"
    
         output from parser-cmd (PID 1725821) (stderr):
            1 │[0m line 2: invalid table path
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 860
      valid tests:  22 passed,  2 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['sh', 'sources/stage_test.sh', 'valid/table/**'], returncode=1, stdout='FAIL valid/table/array-nest\n     line 4: invalid table path\n \n     Exit 1\n\n     input sent to parser-cmd (PID 1725827):\n        1 │\x1b[0m [[albums]]\n        2 │\x1b[0m name = "Born to Run"\n        3 │\x1b[0m\n        4 │\x1b[0m   [[albums.songs]]\n        5 │\x1b[0m   name = "Jungleland"\n      … (+1542 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "tables-valid.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: tables-invalid (FEATURE-Tables.md)
  intent: The decoder rejects all authoritative invalid TOML table and array-of-table cases owned by this feature.
  return code: 1
  stdout:
    FAIL invalid/table/append-with-dotted-keys-01
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1726104):
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
    
         output from parser-cmd (PID 1726104) (stdout):
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
    
         input sent to parser-cmd (PID 1726233):
            1 │[0m # This is the same issue as in injection-1.toml, except that nests one level
            2 │[0m # deeper. See that file for a more complete description.
            3 │[0m
            4 │[0m [a.b.c.d]
            5 │[0m   z = 9
            6 │[0m
            7 │[0m [a]
            8 │[0m   b.c.d.k.t = "Using dotted keys to add to [a.b.c.d] after explicitly defining it above is not allowed"
    
         output from parser-cmd (PID 1726233) (stdout):
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
    
    FAIL invalid/table/append-with-dotted-keys-04
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1726077):
            1 │[0m [dependencies.foo]
            2 │[0m version = "0.16"
            3 │[0m
            4 │[0m [dependencies]
            5 │[0m libc = "0.2"
            6 │[0m
            7 │[0m [dependencies]
            8 │[0m rand = "0.3.14"
    
         output from parser-cmd (PID 1726077) (stdout):
            1 │[0m {
            2 │[0m   "dependencies": {
            3 │[0m     "libc": {"type": "string", "value": "0.2"},
            4 │[0m     "rand": {"type": "string", "value": "0.3.14"},
            5 │[0m     "foo": {
            6 │[0m       "version": {"type": "string", "value": "0.16"}
            7 │[0m     }
            8 │[0m   }
            9 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/table/append-with-dotted-keys-08
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1726226):
            1 │[0m [a.b.c]
            2 │[0m z = 9
            3 │[0m
            4 │[0m [[totally_unrelated]]
            5 │[0m x = 123
            6 │[0m
            7 │[0m [a]
            8 │[0m b.c.t = "this should NOT be allowed"
    
         output from parser-cmd (PID 1726226) (stdout):
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
    
         input sent to parser-cmd (PID 1726162):
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
    
         output from parser-cmd (PID 1726162) (stdout):
            1 │[0m {
            2 │[0m   "albums": [{
            3 │[0m     "name": {"type": "string", "value": "Born in the USA"}
            4 │[0m   }]
            5 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/table/duplicate-key-01
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1726107):
            1 │[0m [a]
            2 │[0m b = 1
            3 │[0m
            4 │[0m [a]
            5 │[0m c = 2
    
         output from parser-cmd (PID 1726107) (stdout):
            1 │[0m {
            2 │[0m   "a": {
            3 │[0m     "b": {"type": "integer", "value": "1"},
            4 │[0m     "c": {"type": "integer", "value": "2"}
            5 │[0m   }
            6 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/table/duplicate-key-03
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1726166):
            1 │[0m [fruit]
            2 │[0m apple.color = "red"
            3 │[0m
            4 │[0m [[fruit.apple]]
    
         output from parser-cmd (PID 1726166) (stdout):
            1 │[0m {
            2 │[0m   "fruit": {"apple": [{}]}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/table/duplicate-key-04
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1726105):
            1 │[0m [fruit]
            2 │[0m apple.color = "red"
            3 │[0m
            4 │[0m [fruit.apple] # INVALID
    
         output from parser-cmd (PID 1726105) (stdout):
            1 │[0m {
            2 │[0m   "fruit": {
            3 │[0m     "apple": {
            4 │[0m       "color": {"type": "string", "value": "red"}
            5 │[0m     }
            6 │[0m   }
            7 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/table/duplicate-key-05
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1726112):
            1 │[0m [fruit]
            2 │[0m apple.taste.sweet = true
            3 │[0m
            4 │[0m [fruit.apple.taste] # INVALID
    
         output from parser-cmd (PID 1726112) (stdout):
            1 │[0m {
            2 │[0m   "fruit": {
            3 │[0m     "apple": {
            4 │[0m       "taste": {
            5 │[0m         "sweet": {"type": "bool", "value": "true"}
            6 │[0m       }
            7 │[0m     }
            8 │[0m   }
            9 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/table/duplicate-key-06
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1726101):
            1 │[0m [tbl]
            2 │[0m [[tbl]]
    
         output from parser-cmd (PID 1726101) (stdout):
            1 │[0m {
            2 │[0m   "tbl": [{}]
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/table/duplicate-key-09
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1726164):
            1 │[0m [a]
            2 │[0m foo="bar"
            3 │[0m [a.b]
            4 │[0m foo="bar"
            5 │[0m [a]
    
         output from parser-cmd (PID 1726164) (stdout):
            1 │[0m {
            2 │[0m   "a": {
            3 │[0m     "foo": {"type": "string", "value": "bar"},
            4 │[0m     "b": {
            5 │[0m       "foo": {"type": "string", "value": "bar"}
            6 │[0m     }
            7 │[0m   }
            8 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/table/overwrite-bool-with-array
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1726317):
            1 │[0m a=true
            2 │[0m [[a]]
    
         output from parser-cmd (PID 1726317) (stdout):
            1 │[0m {
            2 │[0m   "a": [{}]
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/table/redefine-02
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1726353):
            1 │[0m [t1]
            2 │[0m t2.t3.v = 0
            3 │[0m [t1.t2]
    
         output from parser-cmd (PID 1726353) (stdout):
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
    
    FAIL invalid/table/redefine-03
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1726355):
            1 │[0m [t1]
            2 │[0m t2.t3.v = 0
            3 │[0m [t1.t2.t3]
    
         output from parser-cmd (PID 1726355) (stdout):
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
    
    FAIL invalid/table/super-twice
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1726361):
            1 │[0m [a.b]
            2 │[0m [a]
            3 │[0m [a]
    
         output from parser-cmd (PID 1726361) (stdout):
            1 │[0m {
            2 │[0m   "a": {
            3 │[0m     "b": {}
            4 │[0m   }
            5 │[0m }
    
         want:
           Exit code 1
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 822
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  47 passed, 15 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['sh', 'sources/stage_test.sh', 'invalid/table/**'], returncode=1, stdout='FAIL invalid/table/append-with-dotted-keys-01\n     Expected an error, but no error was reported.\n\n     input sent to parser-cmd (PID 1726104):\n        1 │\x1b[0m # First a.b.c defines a table: a.b.c = {z=9}\n        2 │\x1b[0m #\n        3 │\x1b[0m # Then we define a.b.c.t = "str" to add a str to t… (+11446 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "tables-invalid.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: datetimes-valid (FEATURE-Datetimes.md)
  intent: The decoder passes all authoritative valid TOML datetime cases owned by this feature.
  return code: 1
  stdout:
    FAIL valid/datetime/datetime
         line 4: invalid value "1987-07-05t17:45:00z"
     
         Exit 1
    
         input sent to parser-cmd (PID 1726620):
            1 │[0m space = 1987-07-05 17:45:00Z
            2 │[0m
            3 │[0m # ABNF is case-insensitive, both "Z" and "z" must be supported.
            4 │[0m lower = 1987-07-05t17:45:00z
    
         output from parser-cmd (PID 1726620) (stderr):
            1 │[0m line 4: invalid value "1987-07-05t17:45:00z"
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 875
      valid tests:   8 passed,  1 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['sh', 'sources/stage_test.sh', 'valid/datetime/**'], returncode=1, stdout='FAIL valid/datetime/datetime\n     line 4: invalid value "1987-07-05t17:45:00z"\n \n     Exit 1\n\n     input sent to parser-cmd (PID 1726620):\n        1 │\x1b[0m space = 1987-07-05 17:45:00Z\n        2 │\x1b[0m\n        3 │\x1b[0m # ABNF is case-insensitive, both "Z" and "z" must be supported.\n    … (+624 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "datetimes-valid.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- GREEN (prepassed): datetimes-invalid (FEATURE-Datetimes.md)
  intent: The decoder rejects all authoritative invalid TOML datetime, local-date, local-datetime, and local-time cases owned by this feature.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true

## Post-build programmatic acceptance
- PASS: tables-valid (FEATURE-Tables.md)
  intent: The decoder passes all authoritative valid TOML table and array-of-table cases owned by this feature.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 860
      valid tests:  24 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: tables-invalid (FEATURE-Tables.md)
  intent: The decoder rejects all authoritative invalid TOML table and array-of-table cases owned by this feature.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 822
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  62 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: datetimes-valid (FEATURE-Datetimes.md)
  intent: The decoder passes all authoritative valid TOML datetime cases owned by this feature.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 875
      valid tests:   9 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: datetimes-invalid (FEATURE-Datetimes.md)
  intent: The decoder rejects all authoritative invalid TOML datetime, local-date, local-datetime, and local-time cases owned by this feature.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true

## Build summary
RESULT: SUCCESS

FILES CHANGED:
- internal/toml/toml.go
- internal/toml/toml_test.go

SUMMARY:
Implemented table lifecycle/array-of-table resolution and TOML datetime parsing, normalization, and precision truncation. Full suite passed: 205 valid, 474 invalid. Race tests and vet passed.

BLOCKERS:
- None
