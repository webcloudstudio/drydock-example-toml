# Evidence: Block 13 · Service (block-13)

- block type: block
- date: 2026-08-14
- resulting state: closed/verified
- story points (combined assembled cost): 9099
- execution id: 20260814.055941.201Z-4b157393

## Stories built
- Run the complete pinned TOML 1.0.0 conformance suite. (conformance) [story]

## Stacked context
- compass: COMPASS.md (SP 2629)
- implements: FEATURE-Conformance.md (SP 468)
- context: full_test.sh (SP 113)
- context: run_conformance.sh (SP 575)
- context: setup_harness.sh (SP 323)
- context: ARCHITECTURE_compact.md (SP 108)
- stack: go.md (SP 3646)
- stack: common_compact.md (SP 1179)

## Build directory changes
- internal/toml/decode.go
- toml-decoder

## Pre-build acceptance observation
- RED: complete-conformance (FEATURE-Conformance.md)
  intent: The completed decoder passes the complete pinned TOML 1.0.0 conformance suite.
  return code: 1
  stdout:
    FAIL valid/spec-1.0.0/string-4
         line 7: unterminated multiline string
     
         Exit 1
    
         input sent to parser-cmd (PID 888242):
            1 │[0m str4 = """Here are two quotation marks: "". Simple enough."""
            2 │[0m # str5 = """Here are three quotation marks: """."""  # INVALID
            3 │[0m str5 = """Here are three quotation marks: ""\"."""
            4 │[0m str6 = """Here are fifteen quotation marks: ""\"""\"""\"""\"""\"."""
            5 │[0m
            6 │[0m # "This," she said, "is just a pointless statement."
            7 │[0m str7 = """"This," she said, "is just a pointless statement.""""
    
         output from parser-cmd (PID 888242) (stderr):
            1 │[0m line 7: unterminated multiline string
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/spec-1.0.0/string-7
         line 7: unterminated multiline string
     
         Exit 1
    
         input sent to parser-cmd (PID 888269):
            1 │[0m quot15 = '''Here are fifteen quotation marks: """""""""""""""'''
            2 │[0m
            3 │[0m # apos15 = '''Here are fifteen apostrophes: ''''''''''''''''''  # INVALID
            4 │[0m apos15 = "Here are fifteen apostrophes: '''''''''''''''"
            5 │[0m
            6 │[0m # 'That,' she said, 'is still pointless.'
            7 │[0m str = ''''That,' she said, 'is still pointless.''''
    
         output from parser-cmd (PID 888269) (stderr):
            1 │[0m line 7: unterminated multiline string
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/multiline-quotes
         line 4: unterminated multiline string
     
         Exit 1
    
         input sent to parser-cmd (PID 888431):
            1 │[0m # Make sure that quotes inside multiline strings are allowed, including right
            2 │[0m # after the opening '''/""" and before the closing '''/"""
            3 │[0m
            4 │[0m lit_one = ''''one quote''''
            5 │[0m lit_two = '''''two quotes'''''
            6 │[0m lit_one_space = ''' 'one quote' '''
            7 │[0m lit_two_space = ''' ''two quotes'' '''
            8 │[0m
            9 │[0m one = """"one quote""""
           10 │[0m two = """""two quotes"""""
           11 │[0m one_space = """ "one quote" """
           12 │[0m two_space = """ ""two quotes"" """
           13 │[0m
           14 │[0m mismatch1 = """aaa'''bbb"""
           15 │[0m mismatch2 = '''aaa"""bbb'''
           16 │[0m
           17 │[0m # Three opening """, then one escaped ", then two "" (allowed), and then three
           18 │[0m # closing """
           19 │[0m escaped = """lol\""""""
           20 │[0m
           21 │[0m five-quotes = """
           22 │[0m Closing with five quotes
           23 │[0m """""
           24 │[0m four-quotes = """
           25 │[0m Closing with four quotes
           26 │[0m """"
    
         output from parser-cmd (PID 888431) (stderr):
            1 │[0m line 4: unterminated multiline string
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/raw-multiline
         line 13: unterminated multiline string
     
         Exit 1
    
         input sent to parser-cmd (PID 888436):
            1 │[0m # Single ' should be allowed.
            2 │[0m oneline = '''This string has a ' quote character.'''
            3 │[0m
            4 │[0m # A newline immediately following the opening delimiter will be trimmed.
            5 │[0m firstnl = '''
            6 │[0m This string has a ' quote character.'''
            7 │[0m
            8 │[0m # All other whitespace and newline characters remain intact.
            9 │[0m multiline = '''
           10 │[0m This string
           11 │[0m has ' a quote character
           12 │[0m and more than
           13 │[0m one newline
           14 │[0m in it.'''
           15 │[0m
           16 │[0m # Tab character in literal string does not need to be escaped
           17 │[0m multiline_with_tab = '''First line
           18 │[0m 	 Followed by a tab'''
           19 │[0m
           20 │[0m this-str-has-apostrophes='''' there's one already
           21 │[0m '' two more
           22 │[0m '''''
    
         output from parser-cmd (PID 888436) (stderr):
            1 │[0m line 13: unterminated multiline string
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL invalid/float/exp-dot-02
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 889197):
            1 │[0m exp-dot-02 = 1.e2
    
         output from parser-cmd (PID 889197) (stdout):
            1 │[0m {
            2 │[0m   "exp-dot-02": {"type": "float", "value": "1.e2"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/exp-dot-03
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 889199):
            1 │[0m exp-dot-03 = 3.e+20
    
         output from parser-cmd (PID 889199) (stdout):
            1 │[0m {
            2 │[0m   "exp-dot-03": {"type": "float", "value": "3.e+20"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/leading-dot
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 889294):
            1 │[0m leading-dot = .12345
    
         output from parser-cmd (PID 889294) (stdout):
            1 │[0m {
            2 │[0m   "leading-dot": {"type": "float", "value": ".12345"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/leading-dot-neg
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 889283):
            1 │[0m leading-dot-neg = -.12345
    
         output from parser-cmd (PID 889283) (stdout):
            1 │[0m {
            2 │[0m   "leading-dot-neg": {"type": "float", "value": "-.12345"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/leading-dot-plus
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 889284):
            1 │[0m leading-dot-plus = +.12345
    
         output from parser-cmd (PID 889284) (stdout):
            1 │[0m {
            2 │[0m   "leading-dot-plus": {"type": "float", "value": "+.12345"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/leading-zero
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 889321):
            1 │[0m leading-zero = 03.14
    
         output from parser-cmd (PID 889321) (stdout):
            1 │[0m {
            2 │[0m   "leading-zero": {"type": "float", "value": "03.14"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/leading-zero-neg
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 889317):
            1 │[0m leading-zero-neg = -03.14
    
         output from parser-cmd (PID 889317) (stdout):
            1 │[0m {
            2 │[0m   "leading-zero-neg": {"type": "float", "value": "-03.14"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/leading-zero-plus
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 889318):
            1 │[0m leading-zero-plus = +03.14
    
         output from parser-cmd (PID 889318) (stdout):
            1 │[0m {
            2 │[0m   "leading-zero-plus": {"type": "float", "value": "+03.14"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/trailing-dot
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 889384):
            1 │[0m trailing-dot = 1.
    
         output from parser-cmd (PID 889384) (stdout):
            1 │[0m {
            2 │[0m   "trailing-dot": {"type": "float", "value": "1."}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/trailing-dot-01
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 889356):
            1 │[0m trailing-point = 1.
    
         output from parser-cmd (PID 889356) (stdout):
            1 │[0m {
            2 │[0m   "trailing-point": {"type": "float", "value": "1."}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/trailing-dot-02
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 889357):
            1 │[0m a = 1.
            2 │[0m b = 2
    
         output from parser-cmd (PID 889357) (stdout):
            1 │[0m {
            2 │[0m   "a": {"type": "float", "value": "1."},
            3 │[0m   "b": {"type": "integer", "value": "2"}
            4 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/trailing-dot-min
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 889359):
            1 │[0m trailing-dot-min = -1.
    
         output from parser-cmd (PID 889359) (stdout):
            1 │[0m {
            2 │[0m   "trailing-dot-min": {"type": "float", "value": "-1."}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/trailing-dot-plus
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 889435):
            1 │[0m trailing-dot-plus = +1.
    
         output from parser-cmd (PID 889435) (stdout):
            1 │[0m {
            2 │[0m   "trailing-dot-plus": {"type": "float", "value": "+1."}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/integer/negative-bin
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 889663):
            1 │[0m negative-bin = -0b11010110
    
         output from parser-cmd (PID 889663) (stdout):
            1 │[0m {
            2 │[0m   "negative-bin": {"type": "integer", "value": "-214"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/integer/negative-hex
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 889665):
            1 │[0m negative-hex = -0xff
    
         output from parser-cmd (PID 889665) (stdout):
            1 │[0m {
            2 │[0m   "negative-hex": {"type": "integer", "value": "-255"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/integer/negative-oct
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 889667):
            1 │[0m negative-oct = -0o755
    
         output from parser-cmd (PID 889667) (stdout):
            1 │[0m {
            2 │[0m   "negative-oct": {"type": "integer", "value": "-493"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/integer/positive-bin
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 889690):
            1 │[0m positive-bin = +0b11010110
    
         output from parser-cmd (PID 889690) (stdout):
            1 │[0m {
            2 │[0m   "positive-bin": {"type": "integer", "value": "214"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/integer/positive-hex
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 889695):
            1 │[0m positive-hex = +0xff
    
         output from parser-cmd (PID 889695) (stdout):
            1 │[0m {
            2 │[0m   "positive-hex": {"type": "integer", "value": "255"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/integer/positive-oct
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 889713):
            1 │[0m positive-oct = +0o755
    
         output from parser-cmd (PID 889713) (stdout):
            1 │[0m {
            2 │[0m   "positive-oct": {"type": "integer", "value": "493"}
            3 │[0m }
    
         want:
           Exit code 1
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
      valid tests: 201 passed,  4 failed
    encoder tests: no encoder command given
    invalid tests: 455 passed, 19 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
    
    Traceback (most recent call last):
      File "complete-conformance.py", line 13, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- PASS: complete-conformance (FEATURE-Conformance.md)
  intent: The completed decoder passes the complete pinned TOML 1.0.0 conformance suite.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
      valid tests: 205 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests: 474 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true

## Build summary
RESULT: SUCCESS

FILES CHANGED:
- internal/toml/decode.go

SUMMARY:
Full pinned TOML 1.0.0 conformance passed:
- Valid: 205/205
- Invalid: 474/474

Also passed `go test -race ./...`, `go vet ./...`, and `go build ./...`.

BLOCKERS:
- None
