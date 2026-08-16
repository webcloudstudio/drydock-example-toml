# Evidence: Block 3 · Service (block-3)

- block type: block
- date: 2026-08-15
- resulting state: closed/verified
- story points (combined assembled cost): 8724
- execution id: 20260816.010051.090Z-0d65277f

## Stories built
- Parse all TOML basic, literal, and multiline string forms. (parser-strings) [story]

## Acceptance tooling authorization
- FEATURE-Parser-Strings.md#parser-strings-valid: executable=sh; scope=test; authorization=existing Target environment
- FEATURE-Parser-Strings.md#parser-strings-valid: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Parser-Strings.md#parser-strings-invalid: executable=sh; scope=test; authorization=existing Target environment
- FEATURE-Parser-Strings.md#parser-strings-invalid: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Parser-Strings.md#parser-strings-contract: executable=go; scope=test; authorization=existing Target environment

## Reusable compacts
- stage_test_compact.md

## Stacked context
- compass: COMPASS.md (SP 2899)
- implements: FEATURE-Parser-Strings.md (SP 633)
- context: stage_test.sh (SP 161)
- context: ARCHITECTURE_compact.md (SP 128)
- stack: go.md (SP 3646)
- stack: common_compact.md (SP 1179)

## Build directory changes
- cmd/toml-decoder/main.go
- internal/toml/toml.go
- internal/toml/toml_test.go
- toml-decoder

## Pre-build acceptance observation
- RED: parser-strings-valid (FEATURE-Parser-Strings.md)
  intent: The authoritative TOML suite passes every valid string case.
  return code: 1
  stdout:
    FAIL valid/string/basic-escape-01
         line 2: unsupported value
     
         Exit 1
    
         input sent to parser-cmd (PID 1452558):
            1 │[0m # Escape "
            2 │[0m test = "\"one\""
    
         output from parser-cmd (PID 1452558) (stderr):
            1 │[0m line 2: unsupported value
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/basic-escape-02
         line 2: unsupported value
     
         Exit 1
    
         input sent to parser-cmd (PID 1452559):
            1 │[0m # Escape \ and then "
            2 │[0m test = "\\\"one"
    
         output from parser-cmd (PID 1452559) (stderr):
            1 │[0m line 2: unsupported value
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/basic-escape-03
         line 2: unsupported value
     
         Exit 1
    
         input sent to parser-cmd (PID 1452551):
            1 │[0m # Escape \ four times and then "
            2 │[0m test = "\\\\\\\\\"one"
    
         output from parser-cmd (PID 1452551) (stderr):
            1 │[0m line 2: unsupported value
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/empty
         line 1: invalid integer
     
         Exit 1
    
         input sent to parser-cmd (PID 1452549):
            1 │[0m answer = ""
    
         output from parser-cmd (PID 1452549) (stderr):
            1 │[0m line 1: invalid integer
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/ends-in-whitespace-escape
         line 1: invalid integer
     
         Exit 1
    
         input sent to parser-cmd (PID 1452582):
            1 │[0m beee = """
            2 │[0m heeee
            3 │[0m geeee\
            4 │[0m
            5 │[0m
            6 │[0m       """
    
         output from parser-cmd (PID 1452582) (stderr):
            1 │[0m line 1: invalid integer
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/escape-tricky
         line 1: unsupported value
     
         Exit 1
    
         input sent to parser-cmd (PID 1452560):
            1 │[0m end_esc = "String does not end here\" but ends here\\"
            2 │[0m lit_end_esc = 'String ends here\'
            3 │[0m
            4 │[0m multiline_unicode = """
            5 │[0m \u00a0"""
            6 │[0m
            7 │[0m multiline_not_unicode = """
            8 │[0m \\u0041"""
            9 │[0m
           10 │[0m multiline_end_esc = """When will it end? \"""...""\" should be here\""""
           11 │[0m
           12 │[0m lit_multiline_not_unicode = '''
           13 │[0m \u007f'''
           14 │[0m
           15 │[0m lit_multiline_end = '''There is no escape\'''
    
         output from parser-cmd (PID 1452560) (stderr):
            1 │[0m line 1: unsupported value
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/escaped-escape
         line 1: invalid integer
     
         Exit 1
    
         input sent to parser-cmd (PID 1452624):
            1 │[0m answer = "\\x64"
    
         output from parser-cmd (PID 1452624) (stderr):
            1 │[0m line 1: invalid integer
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/escapes
         line 1: unsupported value
     
         Exit 1
    
         input sent to parser-cmd (PID 1452620):
            1 │[0m backspace     = "|\b."
            2 │[0m tab           = "|\t."
            3 │[0m newline       = "|\n."
            4 │[0m formfeed      = "|\f."
            5 │[0m carriage      = "|\r."
            6 │[0m quote         = "|\"."
            7 │[0m backslash     = "|\\."
            8 │[0m delete        = "|\u007F."
            9 │[0m unitseparator = "|\u001F."
           10 │[0m
           11 │[0m # \u is escaped, so should NOT be interperted as a \u escape.
           12 │[0m notunicode1   = "|\\u."
           13 │[0m notunicode2   = "|\u005Cu."
           14 │[0m notunicode3   = "|\\u0075."
           15 │[0m notunicode4   = "|\\\u0075."
    
         output from parser-cmd (PID 1452620) (stderr):
            1 │[0m line 1: unsupported value
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/multibyte
         line 4: invalid integer
     
         Exit 1
    
         input sent to parser-cmd (PID 1452563):
            1 │[0m # Test each multibyte length: 2, 3, and 4 bytes:
            2 │[0m # ɑ € 𐫱
            3 │[0m
            4 │[0m basic    = "ɑ € 𐫱 ɑ€𐫱"
            5 │[0m raw      = 'ɑ € 𐫱 ɑ€𐫱'
            6 │[0m ml-basic = """ɑ € 𐫱 ɑ€𐫱"""
            7 │[0m ml-raw   = '''ɑ € 𐫱 ɑ€𐫱'''
    
         output from parser-cmd (PID 1452563) (stderr):
            1 │[0m line 4: invalid integer
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/multibyte-escape
         line 4: invalid integer
     
         Exit 1
    
         input sent to parser-cmd (PID 1452626):
            1 │[0m # Test each multibyte length: 2, 3, and 4 bytes:
            2 │[0m # ɑ € 𐫱
            3 │[0m
            4 │[0m basic-1    = "\u0251 \u20ac \U00010AF1 \u0251\u20ac\U00010AF1"
            5 │[0m ml-basic-1 = """\u0251 \u20ac \U00010AF1 \u0251\u20ac\U00010AF1"""
            6 │[0m
            7 │[0m # Again, but only using \U
            8 │[0m basic-2    = "\U00000251 \U000020ac \U00010AF1 \U00000251\U000020ac\U00010AF1"
            9 │[0m ml-basic-2 = """\U00000251 \U000020ac \U00010AF1 \U00000251\U000020ac\U00010AF1"""
    
         output from parser-cmd (PID 1452626) (stderr):
            1 │[0m line 4: invalid integer
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/multiline
         line 3: unsupported value
     
         Exit 1
    
         input sent to parser-cmd (PID 1452550):
            1 │[0m # NOTE: this file includes some literal tab characters.
            2 │[0m
            3 │[0m equivalent_one = "The quick brown fox jumps over the lazy dog."
            4 │[0m equivalent_two = """
            5 │[0m The quick brown \
            6 │[0m
            7 │[0m
            8 │[0m   fox jumps over \
            9 │[0m     the lazy dog."""
           10 │[0m
           11 │[0m equivalent_three = """\
           12 │[0m        The quick brown \
           13 │[0m        fox jumps over \
           14 │[0m        the lazy dog.\
           15 │[0m        """
           16 │[0m
           17 │[0m whitespace-after-bs = """\
           18 │[0m        The quick brown \
           19 │[0m        fox jumps over \
           20 │[0m        the lazy dog.\
           21 │[0m        """
           22 │[0m
           23 │[0m no-space = """a\
           24 │[0m     b"""
           25 │[0m
           26 │[0m # Has tab character.
           27 │[0m keep-ws-before = """a   	\
           28 │[0m    b"""
           29 │[0m
           30 │[0m escape-bs-1 = """a \\
           31 │[0m b"""
           32 │[0m
           33 │[0m escape-bs-2 = """a \\\
           34 │[0m b"""
           35 │[0m
           36 │[0m escape-bs-3 = """a \\\\
           37 │[0m   b"""
    
         output from parser-cmd (PID 1452550) (stderr):
            1 │[0m line 3: unsupported value
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/multiline-empty
         line 1: invalid integer
     
         Exit 1
    
         input sent to parser-cmd (PID 1452564):
            1 │[0m empty-1 = """"""
            2 │[0m
            3 │[0m # A newline immediately following the opening delimiter will be trimmed.
            4 │[0m empty-2 = """
            5 │[0m """
            6 │[0m
            7 │[0m # \ at the end of line trims newlines as well; note that last \ is followed by
            8 │[0m # two spaces, which are ignored.
            9 │[0m empty-3 = """\
           10 │[0m     """
           11 │[0m empty-4 = """\
           12 │[0m    \
           13 │[0m    \
           14 │[0m    """
    
         output from parser-cmd (PID 1452564) (stderr):
            1 │[0m line 1: invalid integer
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/multiline-escaped-crlf
         line 3: invalid integer
     
         Exit 1
    
         input sent to parser-cmd (PID 1452617):
            1 │[0m # The following line should be an unescaped backslash followed by a Windows
            2 │[0m # newline sequence ("\r\n")
            3 │[0m 0="""\
            4 │[0m """
    
         output from parser-cmd (PID 1452617) (stderr):
            1 │[0m line 3: invalid integer
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/multiline-quotes
         line 4: unsupported value
     
         Exit 1
    
         input sent to parser-cmd (PID 1452562):
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
    
         output from parser-cmd (PID 1452562) (stderr):
            1 │[0m line 4: unsupported value
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/nl
         line 1: unsupported value
     
         Exit 1
    
         input sent to parser-cmd (PID 1452585):
            1 │[0m nl_mid = "val\nue"
            2 │[0m nl_end = """value\n"""
            3 │[0m
            4 │[0m lit_nl_end = '''value\n'''
            5 │[0m lit_nl_mid = 'val\nue'
            6 │[0m lit_nl_uni = 'val\ue'
    
         output from parser-cmd (PID 1452585) (stderr):
            1 │[0m line 1: unsupported value
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/quoted-unicode
         line 2: unsupported value
     
         Exit 1
    
         input sent to parser-cmd (PID 1452547):
            1 │[0m
            2 │[0m escaped_string = "\u0000 \u0008 \u000c \U00000041 \u007f \u0080 \u00ff \ud7ff \ue000 \uffff \U00010000 \U0010ffff"
            3 │[0m not_escaped_string = '\u0000 \u0008 \u000c \U00000041 \u007f \u0080 \u00ff \ud7ff \ue000 \uffff \U00010000 \U0010ffff'
            4 │[0m
            5 │[0m basic_string = "~  ÿ ퟿  ￿ 𐀀 􏿿"
            6 │[0m literal_string = '~  ÿ ퟿  ￿ 𐀀 􏿿'
    
         output from parser-cmd (PID 1452547) (stderr):
            1 │[0m line 2: unsupported value
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/raw
         line 1: unsupported value
     
         Exit 1
    
         input sent to parser-cmd (PID 1452546):
            1 │[0m backspace = 'This string has a \b backspace character.'
            2 │[0m tab = 'This string has a \t tab character.'
            3 │[0m unescaped_tab = 'This string has an 	 unescaped tab character.'
            4 │[0m newline = 'This string has a \n new line character.'
            5 │[0m formfeed = 'This string has a \f form feed character.'
            6 │[0m carriage = 'This string has a \r carriage return character.'
            7 │[0m slash = 'This string has a \/ slash character.'
            8 │[0m backslash = 'This string has a \\ backslash character.'
    
         output from parser-cmd (PID 1452546) (stderr):
            1 │[0m line 1: unsupported value
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/raw-empty
         line 1: invalid integer
     
         Exit 1
    
         input sent to parser-cmd (PID 1452621):
            1 │[0m empty = ''
    
         output from parser-cmd (PID 1452621) (stderr):
            1 │[0m line 1: invalid integer
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/raw-multiline
         line 2: unsupported value
     
         Exit 1
    
         input sent to parser-cmd (PID 1452553):
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
    
         output from parser-cmd (PID 1452553) (stderr):
            1 │[0m line 2: unsupported value
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/simple
         line 1: unsupported value
     
         Exit 1
    
         input sent to parser-cmd (PID 1452613):
            1 │[0m answer = "You are not drinking enough whisky."
    
         output from parser-cmd (PID 1452613) (stderr):
            1 │[0m line 1: unsupported value
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/start-mb
         line 4: invalid integer
     
         Exit 1
    
         input sent to parser-cmd (PID 1452561):
            1 │[0m # Start first line with a multibyte character.
            2 │[0m #
            3 │[0m # https://github.com/marzer/tomlplusplus/issues/190
            4 │[0m s1 = "§"
            5 │[0m s2 = '§'
            6 │[0m s3 = """\
            7 │[0m §"""
            8 │[0m s4 = """
            9 │[0m §"""
           10 │[0m s5 = """§"""
           11 │[0m s6 = '''
           12 │[0m §'''
           13 │[0m s7 = '''§'''
    
         output from parser-cmd (PID 1452561) (stderr):
            1 │[0m line 4: invalid integer
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/unicode-escape
         line 1: invalid integer
     
         Exit 1
    
         input sent to parser-cmd (PID 1452565):
            1 │[0m delta-1 = "\u03B4"
            2 │[0m delta-2 = "\U000003B4"
            3 │[0m a       = "\u0061"
            4 │[0m b       = "\u0062"
            5 │[0m c       = "\U00000063"
            6 │[0m null-1  = "\u0000"
            7 │[0m null-2  = "\U00000000"
            8 │[0m
            9 │[0m ml-delta-1 = """\u03B4"""
           10 │[0m ml-delta-2 = """\U000003B4"""
           11 │[0m ml-a       = """\u0061"""
           12 │[0m ml-b       = """\u0062"""
           13 │[0m ml-c       = """\U00000063"""
           14 │[0m ml-null-1  = """\u0000"""
           15 │[0m ml-null-2  = """\U00000000"""
    
         output from parser-cmd (PID 1452565) (stderr):
            1 │[0m line 1: invalid integer
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/with-pound
         line 1: unsupported value
     
         Exit 1
    
         input sent to parser-cmd (PID 1452548):
            1 │[0m pound = "We see no # comments here."
            2 │[0m poundcomment = "But there are # some comments here." # Did I # mess you up?
    
         output from parser-cmd (PID 1452548) (stderr):
            1 │[0m line 1: unsupported value
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.004939/build/toml/toml-decoder] [no encoder]
    skipped tests: 861
      valid tests:   0 passed, 23 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
  stderr:
    Traceback (most recent call last):
      File "parser-strings-valid.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- GREEN (prepassed): parser-strings-invalid (FEATURE-Parser-Strings.md)
  intent: The authoritative TOML suite rejects every invalid string case.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.004939/build/toml/toml-decoder] [no encoder]
    skipped tests: 807
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  77 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- RED: parser-strings-contract (FEATURE-Parser-Strings.md)
  intent: String parsing remains reachable through the stdin decoder boundary.
  return code: 1
  stderr:
    Traceback (most recent call last):
      File "parser-strings-contract.py", line 15, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- PASS: parser-strings-valid (FEATURE-Parser-Strings.md)
  intent: The authoritative TOML suite passes every valid string case.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.004939/build/toml/toml-decoder] [no encoder]
    skipped tests: 861
      valid tests:  23 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: parser-strings-invalid (FEATURE-Parser-Strings.md)
  intent: The authoritative TOML suite rejects every invalid string case.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.004939/build/toml/toml-decoder] [no encoder]
    skipped tests: 807
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  77 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: parser-strings-contract (FEATURE-Parser-Strings.md)
  intent: String parsing remains reachable through the stdin decoder boundary.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0

## Build summary
<reusable-compact filename="stage_test.sh">
- Runs one authoritative TOML conformance scope.
- Requires exactly one `-run` pattern argument.
- Builds `./cmd/toml-decoder`, then invokes `sources/run_conformance.sh -run "$1"`.
- Exit status is the complete verdict; no output parsing.
</reusable-compact>

RESULT: SUCCESS

FILES CHANGED:
- cmd/toml-decoder/main.go
- internal/toml/toml.go
- internal/toml/toml_test.go

SUMMARY:
Implemented TOML basic, literal, multiline, escape, Unicode, continuation, and malformed-string handling. Valid string tests: 23 passed. Invalid string tests: 77 passed. Contract, race tests, vet, and formatting checks pass.

BLOCKERS:
- None
