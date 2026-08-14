# Evidence: Block 8 · Service (block-8)

- block type: block
- date: 2026-08-14
- resulting state: closed/verified
- story points (combined assembled cost): 8697
- execution id: 20260814.053827.709Z-e3024daf

## Stories built
- Parse bare, quoted, and dotted TOML keys. (keys-forms) [story]
- Parse TOML arrays and nested array values. (structures-arrays) [story]

## Missing context
- context: toml-v1.0.0.md

## Stacked context
- compass: COMPASS.md (SP 2629)
- implements: FEATURE-Key-Forms.md (SP 462)
- context: stage_test.sh (SP 161)
- context: ARCHITECTURE_compact.md (SP 108)
- stack: go.md (SP 3646)
- stack: common_compact.md (SP 1179)
- implements: FEATURE-Arrays.md (SP 417)

## Build directory changes
- internal/toml/decode.go
- internal/toml/value_test.go
- toml-decoder

## Pre-build acceptance observation
- RED: key-forms-valid (FEATURE-Key-Forms.md)
  intent: The implementation passes the authoritative valid TOML key conformance slice.
  return code: 1
  stdout:
    FAIL valid/key/escapes
         line 1: newline in key
     
         Exit 1
    
         input sent to parser-cmd (PID 800855):
            1 │[0m "\n" = "newline"
            2 │[0m "\b" = "bell"
            3 │[0m "\u00c0" = "latin capital letter A with grave"
            4 │[0m "\"" = "just a quote"
            5 │[0m
            6 │[0m ["backsp\b\b"]
            7 │[0m
            8 │[0m ["\"quoted\""]
            9 │[0m quote = true
           10 │[0m
           11 │[0m ["a.b"."\u00c0"]
    
         output from parser-cmd (PID 800855) (stderr):
            1 │[0m line 1: newline in key
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/key/space
         line 4: newline in key
     
         Exit 1
    
         input sent to parser-cmd (PID 800795):
            1 │[0m # Keep whitespace inside quotes keys at all positions.
            2 │[0m "a b"   = 1
            3 │[0m " c d " = 2
            4 │[0m "  much 		  whitespace  	\n  \r\n  " = 3
            5 │[0m
            6 │[0m [ " tbl " ]
            7 │[0m "\ttab\ttab\t" = "tab"
    
         output from parser-cmd (PID 800795) (stderr):
            1 │[0m line 4: newline in key
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 856
      valid tests:  26 passed,  2 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
    
    Traceback (most recent call last):
      File "key-forms-valid.py", line 13, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- GREEN (prepassed): key-forms-invalid (FEATURE-Key-Forms.md)
  intent: The implementation rejects malformed key forms in the authoritative suite.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 823
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  61 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- GREEN (prepassed): arrays-valid (FEATURE-Arrays.md)
  intent: The implementation passes the authoritative valid TOML array conformance slice.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 863
      valid tests:  21 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- RED: arrays-invalid (FEATURE-Arrays.md)
  intent: The implementation rejects every invalid TOML array case in the authoritative suite.
  return code: 1
  stdout:
    FAIL invalid/array/extending-table
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 801825):
            1 │[0m a = [{ b = 1 }]
            2 │[0m
            3 │[0m # Cannot extend tables within static arrays
            4 │[0m # https://github.com/toml-lang/toml/issues/908
            5 │[0m [a.c]
            6 │[0m foo = 1
    
         output from parser-cmd (PID 801825) (stdout):
            1 │[0m {
            2 │[0m   "a": [{
            3 │[0m     "b": {"type": "integer", "value": "1"},
            4 │[0m     "c": {
            5 │[0m       "foo": {"type": "integer", "value": "1"}
            6 │[0m     }
            7 │[0m   }]
            8 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/array/tables-01
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 801832):
            1 │[0m # INVALID TOML DOC
            2 │[0m fruit = []
            3 │[0m
            4 │[0m [[fruit]] # Not allowed
    
         output from parser-cmd (PID 801832) (stdout):
            1 │[0m {
            2 │[0m   "fruit": [{}]
            3 │[0m }
    
         want:
           Exit code 1
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 858
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  24 passed,  2 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
  stderr:
    Traceback (most recent call last):
      File "arrays-invalid.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- PASS: key-forms-valid (FEATURE-Key-Forms.md)
  intent: The implementation passes the authoritative valid TOML key conformance slice.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 856
      valid tests:  28 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: key-forms-invalid (FEATURE-Key-Forms.md)
  intent: The implementation rejects malformed key forms in the authoritative suite.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 823
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  61 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: arrays-valid (FEATURE-Arrays.md)
  intent: The implementation passes the authoritative valid TOML array conformance slice.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 863
      valid tests:  21 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: arrays-invalid (FEATURE-Arrays.md)
  intent: The implementation rejects every invalid TOML array case in the authoritative suite.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 858
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  26 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true

## Build summary
RESULT: SUCCESS

FILES CHANGED:
- internal/toml/decode.go
- internal/toml/value_test.go

SUMMARY:
Implemented robust bare, quoted, and dotted key parsing; nested arrays; comments; mixed values; trailing commas; and malformed-input rejection. Added focused Go tests.

Validation passed:

- Key valid: 28/28
- Key invalid: 61/61
- Array valid: 21/21
- Array invalid: 26/26
- `go test -race ./...`
- `go vet ./...`

BLOCKERS:
- Full conformance still has unrelated failures in table lifecycle, floats, inline tables, integers, and strings outside this build block.
