# Evidence: Block 5 · Service (block-5)

- block type: block
- date: 2026-08-16
- resulting state: closed/verified
- story points (combined assembled cost): 7497
- execution id: 20260816.124257.241Z-fa66cd05

## Stories built
- Handle TOML comments, whitespace, line endings, UTF-8, and control-character rules. (parser-lexical-boundaries) [story]

## Missing context
- context: toml-v1.0.0.md

## Stacked context
- compass: COMPASS.md (SP 2942)
- implements: FEATURE-Lexical-Boundaries.md (SP 566)
- context: stage_test.sh (SP 161)
- context: ARCHITECTURE_compact.md (SP 123)
- stack: go.md (SP 3646)

## Build directory changes
- internal/toml/toml.go
- internal/toml/toml_test.go
- toml-decoder

## Pre-build acceptance observation
- RED: lexical-comments-valid-suite (FEATURE-Lexical-Boundaries.md)
  intent: The authoritative conformance suite passes every valid comment and lexical-boundary case.
  return code: 1
  stdout:
    FAIL valid/comment/everywhere
         line 7: invalid table
     
         Exit 1
    
         input sent to parser-cmd (PID 1709125):
            1 │[0m # Top comment.
            2 │[0m   # Top comment.
            3 │[0m # Top comment.
            4 │[0m
            5 │[0m # [no-extraneous-groups-please]
            6 │[0m
            7 │[0m [group] # Comment
            8 │[0m answer = 42 # Comment
            9 │[0m # no-extraneous-keys-please = 999
           10 │[0m # Inbetween comment.
           11 │[0m more = [ # Comment
           12 │[0m   # What about multiple # comments?
           13 │[0m   # Can you handle it?
           14 │[0m   #
           15 │[0m           # Evil.
           16 │[0m # Evil.
           17 │[0m   42, 42, # Comments within arrays are fun.
           18 │[0m   # What about multiple # comments?
           19 │[0m   # Can you handle it?
           20 │[0m   #
           21 │[0m           # Evil.
           22 │[0m # Evil.
           23 │[0m # ] Did I fool you?
           24 │[0m ] # Hopefully not.
           25 │[0m
           26 │[0m # Make sure the space between the datetime and "#" isn't lexed.
           27 │[0m dt = 1979-05-27T07:32:12-07:00  # c
           28 │[0m d = 1979-05-27 # Comment
           29 │[0m
           30 │[0m [[aot]] # Comment
           31 │[0m k = 98 # Comment
           32 │[0m [[aot]]# Comment
           33 │[0m k = 99# Comment
    
         output from parser-cmd (PID 1709125) (stderr):
            1 │[0m line 7: invalid table
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/comment/tricky
         line 1: invalid table
     
         Exit 1
    
         input sent to parser-cmd (PID 1709123):
            1 │[0m [section]#attached comment
            2 │[0m #[notsection]
            3 │[0m one = "11"#cmt
            4 │[0m two = "22#"
            5 │[0m three = '#'
            6 │[0m
            7 │[0m four = """# no comment
            8 │[0m # nor this
            9 │[0m #also not comment"""#is_comment
           10 │[0m
           11 │[0m five = 5.5#66
           12 │[0m six = 6#7
           13 │[0m 8 = "eight"
           14 │[0m #nine = 99
           15 │[0m ten = 10e2#1
           16 │[0m eleven = 1.11e1#23
           17 │[0m
           18 │[0m ["hash#tag"]
           19 │[0m "#!" = "hash bang"
           20 │[0m arr3 = [ "#", '#', """###""" ]
           21 │[0m arr4 = [ 1,# 9, 9,
           22 │[0m 2#,9
           23 │[0m ,#9
           24 │[0m 3#]
           25 │[0m ,4]
           26 │[0m arr5 = [[[[#["#"],
           27 │[0m ["#"]]]]#]
           28 │[0m ]
           29 │[0m tbl1 = { "#" = '}#'}#}}
    
         output from parser-cmd (PID 1709123) (stderr):
            1 │[0m line 1: invalid table
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 877
      valid tests:   5 passed,  2 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['sh', 'sources/stage_test.sh', 'valid/comment/**'], returncode=1, stdout='FAIL valid/comment/everywhere\n     line 7: invalid table\n \n     Exit 1\n\n     input sent to parser-cmd (PID 1709125):\n        1 │\x1b[0m # Top comment.\n        2 │\x1b[0m   # Top comment.\n        3 │\x1b[0m # Top comment.\n        4 │\x1b[0m\n        5 │\x1b[0m # [no-extraneous-groups-please]\n … (+2948 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "lexical-comments-valid-suite.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- GREEN (prepassed): lexical-boundaries-invalid-suite (FEATURE-Lexical-Boundaries.md)
  intent: The authoritative conformance suite rejects every invalid control-character and encoding case.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 884
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- RED: lexical-boundaries-suites (FEATURE-Lexical-Boundaries.md)
  intent: The valid comment and invalid lexical-boundary conformance slices complete successfully.
  return code: 1
  stdout:
    FAIL valid/comment/everywhere
         line 7: invalid table
     
         Exit 1
    
         input sent to parser-cmd (PID 1709478):
            1 │[0m # Top comment.
            2 │[0m   # Top comment.
            3 │[0m # Top comment.
            4 │[0m
            5 │[0m # [no-extraneous-groups-please]
            6 │[0m
            7 │[0m [group] # Comment
            8 │[0m answer = 42 # Comment
            9 │[0m # no-extraneous-keys-please = 999
           10 │[0m # Inbetween comment.
           11 │[0m more = [ # Comment
           12 │[0m   # What about multiple # comments?
           13 │[0m   # Can you handle it?
           14 │[0m   #
           15 │[0m           # Evil.
           16 │[0m # Evil.
           17 │[0m   42, 42, # Comments within arrays are fun.
           18 │[0m   # What about multiple # comments?
           19 │[0m   # Can you handle it?
           20 │[0m   #
           21 │[0m           # Evil.
           22 │[0m # Evil.
           23 │[0m # ] Did I fool you?
           24 │[0m ] # Hopefully not.
           25 │[0m
           26 │[0m # Make sure the space between the datetime and "#" isn't lexed.
           27 │[0m dt = 1979-05-27T07:32:12-07:00  # c
           28 │[0m d = 1979-05-27 # Comment
           29 │[0m
           30 │[0m [[aot]] # Comment
           31 │[0m k = 98 # Comment
           32 │[0m [[aot]]# Comment
           33 │[0m k = 99# Comment
    
         output from parser-cmd (PID 1709478) (stderr):
            1 │[0m line 7: invalid table
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/comment/tricky
         line 1: invalid table
     
         Exit 1
    
         input sent to parser-cmd (PID 1709470):
            1 │[0m [section]#attached comment
            2 │[0m #[notsection]
            3 │[0m one = "11"#cmt
            4 │[0m two = "22#"
            5 │[0m three = '#'
            6 │[0m
            7 │[0m four = """# no comment
            8 │[0m # nor this
            9 │[0m #also not comment"""#is_comment
           10 │[0m
           11 │[0m five = 5.5#66
           12 │[0m six = 6#7
           13 │[0m 8 = "eight"
           14 │[0m #nine = 99
           15 │[0m ten = 10e2#1
           16 │[0m eleven = 1.11e1#23
           17 │[0m
           18 │[0m ["hash#tag"]
           19 │[0m "#!" = "hash bang"
           20 │[0m arr3 = [ "#", '#', """###""" ]
           21 │[0m arr4 = [ 1,# 9, 9,
           22 │[0m 2#,9
           23 │[0m ,#9
           24 │[0m 3#]
           25 │[0m ,4]
           26 │[0m arr5 = [[[[#["#"],
           27 │[0m ["#"]]]]#]
           28 │[0m ]
           29 │[0m tbl1 = { "#" = '}#'}#}}
    
         output from parser-cmd (PID 1709470) (stderr):
            1 │[0m line 1: invalid table
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 877
      valid tests:   5 passed,  2 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['sh', 'sources/stage_test.sh', 'valid/comment/**'], returncode=1, stdout='FAIL valid/comment/everywhere\n     line 7: invalid table\n \n     Exit 1\n\n     input sent to parser-cmd (PID 1709478):\n        1 │\x1b[0m # Top comment.\n        2 │\x1b[0m   # Top comment.\n        3 │\x1b[0m # Top comment.\n        4 │\x1b[0m\n        5 │\x1b[0m # [no-extraneous-groups-please]\n … (+2948 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "lexical-boundaries-suites.py", line 15, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- PASS: lexical-comments-valid-suite (FEATURE-Lexical-Boundaries.md)
  intent: The authoritative conformance suite passes every valid comment and lexical-boundary case.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 877
      valid tests:   7 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: lexical-boundaries-invalid-suite (FEATURE-Lexical-Boundaries.md)
  intent: The authoritative conformance suite rejects every invalid control-character and encoding case.
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
- PASS: lexical-boundaries-suites (FEATURE-Lexical-Boundaries.md)
  intent: The valid comment and invalid lexical-boundary conformance slices complete successfully.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 877
      valid tests:   7 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
    
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
Implemented comment stripping, UTF-8/control validation, CRLF handling, lexical whitespace checks, and comment-aware arrays/inline tables. Focused suites pass; Go tests and vet pass.

BLOCKERS:
- Full suite still has unrelated table/inline-table failures: 198 valid passed, 7 failed; 442 invalid passed, 32 failed.
