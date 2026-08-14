# Evidence: Block 9 · Service (block-9)

- block type: block
- date: 2026-08-14
- resulting state: closed/verified
- story points (combined assembled cost): 8928
- execution id: 20260814.054305.133Z-f31ab117

## Stories built
- Enforce TOML key uniqueness and scalar-to-table definition rules. (keys-semantics) [story]
- Parse TOML inline tables and nested inline-table keys. (structures-inline-tables) [story]

## Missing context
- context: toml-v1.0.0.md

## Stacked context
- compass: COMPASS.md (SP 2629)
- implements: FEATURE-Key-Semantics.md (SP 634)
- context: stage_test.sh (SP 161)
- context: ARCHITECTURE_compact.md (SP 108)
- stack: go.md (SP 3646)
- stack: common_compact.md (SP 1179)
- implements: FEATURE-Inline-Tables.md (SP 457)

## Build directory changes
- internal/toml/decode.go
- internal/toml/value_test.go
- toml-decoder

## Pre-build acceptance observation
- GREEN (prepassed): key-semantics-valid (FEATURE-Key-Semantics.md)
  intent: The implementation passes the authoritative valid TOML key and specification conformance slices.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 856
      valid tests:  28 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- GREEN (prepassed): key-semantics-invalid (FEATURE-Key-Semantics.md)
  intent: The implementation rejects duplicate and structurally invalid key definitions in the authoritative suite.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 823
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  61 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- RED: key-semantics-boundary (FEATURE-Key-Semantics.md)
  intent: The implementation rejects invalid key/value boundaries in the authoritative specification suite.
  return code: 1
  stdout:
    FAIL invalid/spec-1.0.0/inline-table-2-0
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 818632):
            1 │[0m [product]
            2 │[0m type = { name = "Nail" }
            3 │[0m type.edible = false  # INVALID
    
         output from parser-cmd (PID 818632) (stdout):
            1 │[0m {
            2 │[0m   "product": {
            3 │[0m     "type": {
            4 │[0m       "edible": {"type": "bool", "value": "false"},
            5 │[0m       "name":   {"type": "string", "value": "Nail"}
            6 │[0m     }
            7 │[0m   }
            8 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/spec-1.0.0/string-4-0
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 818607):
            1 │[0m str4 = """Here are two quotation marks: "". Simple enough."""
            2 │[0m str5 = """Here are three quotation marks: """."""  # INVALID
            3 │[0m str5 = """Here are three quotation marks: ""\"."""
            4 │[0m str6 = """Here are fifteen quotation marks: ""\"""\"""\"""\"""\"."""
            5 │[0m
            6 │[0m # "This," she said, "is just a pointless statement."
            7 │[0m str7 = """"This," she said, "is just a pointless statement.""""
    
         output from parser-cmd (PID 818607) (stdout):
            1 │[0m {
            2 │[0m   "str4": {"type": "string", "value": "Here are two quotation marks: \"\". Simple enough."},
            3 │[0m   "str6": {"type": "string", "value": "Here are fifteen quotation marks: \"\"\"\"\"\"\"\"\"\"\"\"\"\"\"."},
            4 │[0m   "str7": {"type": "string", "value": "\"This,\" she said, \"is just a pointless statement.\""},
            5 │[0m   "str5": {
            6 │[0m     "type":  "string",
            7 │[0m     "value": "Here are three quotation marks: \"\"\".\"\"\"  # INVALID\nstr5 = \"\"\"Here are three quotation marks: \"\"\"."
            8 │[0m   }
            9 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/spec-1.0.0/table-9-0
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 818631):
            1 │[0m [fruit]
            2 │[0m apple.color = "red"
            3 │[0m apple.taste.sweet = true
            4 │[0m
            5 │[0m [fruit.apple]  # INVALID
            6 │[0m # [fruit.apple.taste]  # INVALID
            7 │[0m
            8 │[0m [fruit.apple.texture]  # you can add sub-tables
            9 │[0m smooth = true
    
         output from parser-cmd (PID 818631) (stdout):
            1 │[0m {
            2 │[0m   "fruit": {
            3 │[0m     "apple": {
            4 │[0m       "color": {"type": "string", "value": "red"},
            5 │[0m       "taste": {
            6 │[0m         "sweet": {"type": "bool", "value": "true"}
            7 │[0m       },
            8 │[0m       "texture": {
            9 │[0m         "smooth": {"type": "bool", "value": "true"}
           10 │[0m       }
           11 │[0m     }
           12 │[0m   }
           13 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/spec-1.0.0/table-9-1
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 818608):
            1 │[0m [fruit]
            2 │[0m apple.color = "red"
            3 │[0m apple.taste.sweet = true
            4 │[0m
            5 │[0m # [fruit.apple]  # INVALID
            6 │[0m [fruit.apple.taste]  # INVALID
            7 │[0m
            8 │[0m [fruit.apple.texture]  # you can add sub-tables
            9 │[0m smooth = true
    
         output from parser-cmd (PID 818608) (stdout):
            1 │[0m {
            2 │[0m   "fruit": {
            3 │[0m     "apple": {
            4 │[0m       "color": {"type": "string", "value": "red"},
            5 │[0m       "taste": {
            6 │[0m         "sweet": {"type": "bool", "value": "true"}
            7 │[0m       },
            8 │[0m       "texture": {
            9 │[0m         "smooth": {"type": "bool", "value": "true"}
           10 │[0m       }
           11 │[0m     }
           12 │[0m   }
           13 │[0m }
    
         want:
           Exit code 1
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 876
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   4 passed,  4 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
    
    Traceback (most recent call last):
      File "key-semantics-boundary.py", line 13, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- GREEN (prepassed): inline-tables-valid (FEATURE-Inline-Tables.md)
  intent: The implementation passes the authoritative valid TOML inline-table conformance slice.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 867
      valid tests:  17 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- RED: inline-tables-invalid (FEATURE-Inline-Tables.md)
  intent: The implementation rejects every invalid TOML inline-table case in the authoritative suite.
  return code: 1
  stdout:
    FAIL invalid/inline-table/double-comma
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 819102):
            1 │[0m t = {x=3,,y=4}
    
         output from parser-cmd (PID 819102) (stdout):
            1 │[0m {
            2 │[0m   "t": {}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/inline-table/duplicate-key-03
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 819135):
            1 │[0m tbl = { fruit = { apple.color = "red" }, fruit.apple.texture = { smooth = true } }
    
         output from parser-cmd (PID 819135) (stdout):
            1 │[0m {
            2 │[0m   "tbl": {
            3 │[0m     "fruit": {
            4 │[0m       "apple": {
            5 │[0m         "color": {"type": "string", "value": "red"},
            6 │[0m         "texture": {
            7 │[0m           "smooth": {"type": "bool", "value": "true"}
            8 │[0m         }
            9 │[0m       }
           10 │[0m     }
           11 │[0m   }
           12 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/inline-table/empty-01
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 819248):
            1 │[0m t = {,}
    
         output from parser-cmd (PID 819248) (stdout):
            1 │[0m {
            2 │[0m   "t": {}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/inline-table/empty-02
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 819142):
            1 │[0m t = {,
            2 │[0m }
    
         output from parser-cmd (PID 819142) (stdout):
            1 │[0m {
            2 │[0m   "t": {}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/inline-table/empty-03
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 819141):
            1 │[0m t = {
            2 │[0m ,
            3 │[0m }
    
         output from parser-cmd (PID 819141) (stdout):
            1 │[0m {
            2 │[0m   "t": {}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/inline-table/linebreak-01
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 819190):
            1 │[0m # No newlines are allowed between the curly braces unless they are valid within
            2 │[0m # a value.
            3 │[0m simple = { a = 1
            4 │[0m }
    
         output from parser-cmd (PID 819190) (stdout):
            1 │[0m {
            2 │[0m   "simple": {
            3 │[0m     "a": {"type": "integer", "value": "1"}
            4 │[0m   }
            5 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/inline-table/linebreak-02
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 819104):
            1 │[0m t = {a=1,
            2 │[0m b=2}
    
         output from parser-cmd (PID 819104) (stdout):
            1 │[0m {
            2 │[0m   "t": {
            3 │[0m     "a": {"type": "integer", "value": "1"},
            4 │[0m     "b": {"type": "integer", "value": "2"}
            5 │[0m   }
            6 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/inline-table/linebreak-03
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 819101):
            1 │[0m t = {a=1
            2 │[0m ,b=2}
    
         output from parser-cmd (PID 819101) (stdout):
            1 │[0m {
            2 │[0m   "t": {
            3 │[0m     "a": {"type": "integer", "value": "1"},
            4 │[0m     "b": {"type": "integer", "value": "2"}
            5 │[0m   }
            6 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/inline-table/linebreak-04
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 819136):
            1 │[0m json_like = {
            2 │[0m           first = "Tom",
            3 │[0m           last = "Preston-Werner"
            4 │[0m }
    
         output from parser-cmd (PID 819136) (stdout):
            1 │[0m {
            2 │[0m   "json_like": {
            3 │[0m     "first": {"type": "string", "value": "Tom"},
            4 │[0m     "last":  {"type": "string", "value": "Preston-Werner"}
            5 │[0m   }
            6 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/inline-table/overwrite-02
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 819139):
            1 │[0m a={}
            2 │[0m # Inline tables are immutable and can't be extended
            3 │[0m [a.b]
    
         output from parser-cmd (PID 819139) (stdout):
            1 │[0m {
            2 │[0m   "a": {
            3 │[0m     "b": {}
            4 │[0m   }
            5 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/inline-table/overwrite-04
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 819191):
            1 │[0m inline-t = { nest = {} }
            2 │[0m
            3 │[0m [[inline-t.nest]]
    
         output from parser-cmd (PID 819191) (stdout):
            1 │[0m {
            2 │[0m   "inline-t": {"nest": [{}]}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/inline-table/overwrite-05
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 819185):
            1 │[0m inline-t = { nest = {} }
            2 │[0m
            3 │[0m [inline-t.nest]
    
         output from parser-cmd (PID 819185) (stdout):
            1 │[0m {
            2 │[0m   "inline-t": {
            3 │[0m     "nest": {}
            4 │[0m   }
            5 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/inline-table/overwrite-08
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 819183):
            1 │[0m tab = { inner = { dog = "best" }, inner.cat = "worst" }
    
         output from parser-cmd (PID 819183) (stdout):
            1 │[0m {
            2 │[0m   "tab": {
            3 │[0m     "inner": {
            4 │[0m       "cat": {"type": "string", "value": "worst"},
            5 │[0m       "dog": {"type": "string", "value": "best"}
            6 │[0m     }
            7 │[0m   }
            8 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/inline-table/trailing-comma
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 819100):
            1 │[0m # A terminating comma (also called trailing comma) is not permitted after the
            2 │[0m # last key/value pair in an inline table
            3 │[0m abc = { abc = 123, }
    
         output from parser-cmd (PID 819100) (stdout):
            1 │[0m {
            2 │[0m   "abc": {
            3 │[0m     "abc": {"type": "integer", "value": "123"}
            4 │[0m   }
            5 │[0m }
    
         want:
           Exit code 1
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 856
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  14 passed, 14 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
  stderr:
    Traceback (most recent call last):
      File "inline-tables-invalid.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- PASS: key-semantics-valid (FEATURE-Key-Semantics.md)
  intent: The implementation passes the authoritative valid TOML key and specification conformance slices.
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
- PASS: key-semantics-invalid (FEATURE-Key-Semantics.md)
  intent: The implementation rejects duplicate and structurally invalid key definitions in the authoritative suite.
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
- PASS: key-semantics-boundary (FEATURE-Key-Semantics.md)
  intent: The implementation rejects invalid key/value boundaries in the authoritative specification suite.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 876
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   8 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: inline-tables-valid (FEATURE-Inline-Tables.md)
  intent: The implementation passes the authoritative valid TOML inline-table conformance slice.
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
- PASS: inline-tables-invalid (FEATURE-Inline-Tables.md)
  intent: The implementation rejects every invalid TOML inline-table case in the authoritative suite.
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

## Build summary
RESULT: SUCCESS

FILES CHANGED:
- internal/toml/decode.go
- internal/toml/value_test.go

SUMMARY:
Implemented key uniqueness, scalar/table closure rules, nested inline tables, dotted inline keys, strict trailing-comma and newline validation. Scoped acceptance slices and Go tests pass.

BLOCKERS:
- Full conformance suite still reports unrelated pre-existing failures outside this build block.
