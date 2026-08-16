# Evidence: Block 5 · Service (block-5)

- block type: block
- date: 2026-08-15
- resulting state: closed/verified
- story points (combined assembled cost): 8685
- execution id: 20260816.011402.538Z-7d77748f

## Stories built
- Enforce TOML comments, control-character, and UTF-8 boundaries. (parser-lexical-boundaries) [story]

## Acceptance tooling authorization
- FEATURE-Parser-Lexical-Boundaries.md#parser-boundaries-valid: executable=sh; scope=test; authorization=existing Target environment
- FEATURE-Parser-Lexical-Boundaries.md#parser-boundaries-valid: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Parser-Lexical-Boundaries.md#parser-boundaries-invalid-control: executable=sh; scope=test; authorization=existing Target environment
- FEATURE-Parser-Lexical-Boundaries.md#parser-boundaries-invalid-control: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Parser-Lexical-Boundaries.md#parser-boundaries-invalid-encoding: executable=sh; scope=test; authorization=existing Target environment
- FEATURE-Parser-Lexical-Boundaries.md#parser-boundaries-invalid-encoding: executable=go; scope=test; authorization=existing Target environment

## Stacked context
- compass: COMPASS.md (SP 2899)
- implements: FEATURE-Parser-Lexical-Boundaries.md (SP 615)
- context: stage_test.sh (SP 161)
- context: ARCHITECTURE_compact.md (SP 128)
- stack: go.md (SP 3646)
- stack: common_compact.md (SP 1179)

## Build directory changes
- internal/toml/toml.go
- internal/toml/toml_test.go
- toml-decoder

## Pre-build acceptance observation
- RED: parser-boundaries-valid (FEATURE-Parser-Lexical-Boundaries.md)
  intent: The authoritative TOML suite passes every valid comment and boundary case.
  return code: 1
  stdout:
    FAIL valid/comment/everywhere
         line 7: expected key-value pair
     
         Exit 1
    
         input sent to parser-cmd (PID 1467266):
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
    
         output from parser-cmd (PID 1467266) (stderr):
            1 │[0m line 7: expected key-value pair
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/comment/tricky
         line 1: expected key-value pair
     
         Exit 1
    
         input sent to parser-cmd (PID 1467259):
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
    
         output from parser-cmd (PID 1467259) (stderr):
            1 │[0m line 1: expected key-value pair
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.004939/build/toml/toml-decoder] [no encoder]
    skipped tests: 877
      valid tests:   5 passed,  2 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
  stderr:
    Traceback (most recent call last):
      File "parser-boundaries-valid.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: parser-boundaries-invalid-control (FEATURE-Parser-Lexical-Boundaries.md)
  intent: The authoritative TOML suite rejects every invalid control-character case.
  return code: 1
  stdout:
    FAIL invalid/control/bare-cr
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1467542):
            1 │[0m # The following line contains a single carriage return control character
            2 │[0m 
    
         output from parser-cmd (PID 1467542) (stdout):
            1 │[0m {}
    
         want:
           Exit code 1
    
    FAIL invalid/control/comment-cr
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1467545):
            1 │[0m comment-cr   = "Carriage return in comment" # 
    a=1
    
         output from parser-cmd (PID 1467545) (stdout):
            1 │[0m {
            2 │[0m   "a":          {"type": "integer", "value": "1"},
            3 │[0m   "comment-cr": {"type": "string", "value": "Carriage return in comment"}
            4 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/control/comment-del
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1467539):
            1 │[0m comment-del  = "0x7f"   # 
    
         output from parser-cmd (PID 1467539) (stdout):
            1 │[0m {
            2 │[0m   "comment-del": {"type": "string", "value": "0x7f"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/control/comment-ff
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1467615):
            1 │[0m comment-ff   = "0x7f"   # 
    
    
         output from parser-cmd (PID 1467615) (stdout):
            1 │[0m {
            2 │[0m   "comment-ff": {"type": "string", "value": "0x7f"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/control/comment-lf
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1467525):
            1 │[0m comment-lf   = "ctrl-P" # 
    
         output from parser-cmd (PID 1467525) (stdout):
            1 │[0m {
            2 │[0m   "comment-lf": {"type": "string", "value": "ctrl-P"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/control/comment-null
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1467534):
            1 │[0m comment-null = "null"   #  
    
         output from parser-cmd (PID 1467534) (stdout):
            1 │[0m {
            2 │[0m   "comment-null": {"type": "string", "value": "null"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/control/comment-us
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1467491):
            1 │[0m comment-us   = "ctrl-_" # 
    
         output from parser-cmd (PID 1467491) (stdout):
            1 │[0m {
            2 │[0m   "comment-us": {"type": "string", "value": "ctrl-_"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/control/multi-del
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1467493):
            1 │[0m multi-del  = """null"""
    
         output from parser-cmd (PID 1467493) (stdout):
            1 │[0m {
            2 │[0m   "multi-del": {"type": "string", "value": "null"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/control/only-ff
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1467490):
            1 │[0m 
    
    
         output from parser-cmd (PID 1467490) (stdout):
            1 │[0m {}
    
         want:
           Exit code 1
    
    FAIL invalid/control/only-vt
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1467518):
            1 │[0m 
    
    
         output from parser-cmd (PID 1467518) (stdout):
            1 │[0m {}
    
         want:
           Exit code 1
    
    FAIL invalid/control/rawmulti-del
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1467515):
            1 │[0m rawmulti-del  = '''null'''
    
         output from parser-cmd (PID 1467515) (stdout):
            1 │[0m {
            2 │[0m   "rawmulti-del": {"type": "string", "value": "null"}
            3 │[0m }
    
         want:
           Exit code 1
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.004939/build/toml/toml-decoder] [no encoder]
    skipped tests: 852
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  21 passed, 11 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
  stderr:
    Traceback (most recent call last):
      File "parser-boundaries-invalid-control.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: parser-boundaries-invalid-encoding (FEATURE-Parser-Lexical-Boundaries.md)
  intent: The authoritative TOML suite rejects every invalid UTF-8 encoding case.
  return code: 1
  stdout:
    FAIL invalid/encoding/ideographic-space
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1467834):
            1 │[0m # First on next line is U+3000 IDEOGRAPHIC SPACE
            2 │[0m 　foo = "bar"
    
         output from parser-cmd (PID 1467834) (stdout):
            1 │[0m {
            2 │[0m   "foo": {"type": "string", "value": "bar"}
            3 │[0m }
    
         want:
           Exit code 1
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.004939/build/toml/toml-decoder] [no encoder]
    skipped tests: 870
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  13 passed,  1 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
  stderr:
    Traceback (most recent call last):
      File "parser-boundaries-invalid-encoding.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- PASS: parser-boundaries-valid (FEATURE-Parser-Lexical-Boundaries.md)
  intent: The authoritative TOML suite passes every valid comment and boundary case.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.004939/build/toml/toml-decoder] [no encoder]
    skipped tests: 877
      valid tests:   7 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: parser-boundaries-invalid-control (FEATURE-Parser-Lexical-Boundaries.md)
  intent: The authoritative TOML suite rejects every invalid control-character case.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.004939/build/toml/toml-decoder] [no encoder]
    skipped tests: 852
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  32 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: parser-boundaries-invalid-encoding (FEATURE-Parser-Lexical-Boundaries.md)
  intent: The authoritative TOML suite rejects every invalid UTF-8 encoding case.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.004939/build/toml/toml-decoder] [no encoder]
    skipped tests: 870
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  14 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true

## Repair attempts
- attempt 0 (initial build): failed; 2/3 checks; 46/46 cases model=gpt-5.6-luna; execution 20260816.010953.453Z-b14dc4e6
- attempt 1 (repair 1): built; 3/3 checks; 46/46 cases model=gpt-5.6-luna; execution 20260816.011402.538Z-7d77748f

## Build summary
RESULT: SUCCESS

FILES CHANGED:
- internal/toml/toml.go
- toml-decoder

SUMMARY:
Implemented table headers, arrays, inline tables, datetime recognition, comment-aware multiline parsing, and strict boundary handling.

Passed:
- Valid comments: 7/7
- Invalid controls: 32/32
- Invalid UTF-8: 14/14
- Go tests, race tests, vet, and build

BLOCKERS:
- None
