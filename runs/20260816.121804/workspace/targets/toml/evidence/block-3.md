# Evidence: Block 3 · Service (block-3)

- block type: block
- date: 2026-08-16
- resulting state: closed/verified
- story points (combined assembled cost): 8020
- execution id: 20260816.123439.459Z-c6e09f01

## Stories built
- Parse TOML basic, literal, and multiline string values. (parser-strings) [story]
- Parse TOML key forms and enforce key-definition semantics. (parser-keys) [story]

## Acceptance tooling authorization
- FEATURE-Keys.md#keys-valid: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Keys.md#keys-invalid: executable=go; scope=test; authorization=existing Target environment

## Reusable compacts
- stage_test_compact.md

## Missing context
- context: toml-v1.0.0.md

## Stacked context
- compass: COMPASS.md (SP 2942)
- implements: FEATURE-Strings.md (SP 556)
- context: stage_test.sh (SP 161)
- context: ARCHITECTURE_compact.md (SP 123)
- stack: go.md (SP 3646)
- implements: FEATURE-Keys.md (SP 477)

## Build directory changes
- internal/toml/toml.go
- internal/toml/toml_test.go
- toml-decoder

## Pre-build acceptance observation
- RED: strings-valid-suite (FEATURE-Strings.md)
  intent: The authoritative conformance suite passes every valid string case owned by this feature.
  return code: 1
  stdout:
    FAIL valid/string/ends-in-whitespace-escape
         line 1: invalid string: invalid syntax
     
         Exit 1
    
         input sent to parser-cmd (PID 1666750):
            1 │[0m beee = """
            2 │[0m heeee
            3 │[0m geeee\
            4 │[0m
            5 │[0m
            6 │[0m       """
    
         output from parser-cmd (PID 1666750) (stderr):
            1 │[0m line 1: invalid string: invalid syntax
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/escape-tricky
         line 4: invalid string: invalid syntax
     
         Exit 1
    
         input sent to parser-cmd (PID 1666787):
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
    
         output from parser-cmd (PID 1666787) (stderr):
            1 │[0m line 4: invalid string: invalid syntax
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/multibyte
         line 6: invalid string: invalid syntax
     
         Exit 1
    
         input sent to parser-cmd (PID 1666813):
            1 │[0m # Test each multibyte length: 2, 3, and 4 bytes:
            2 │[0m # ɑ € 𐫱
            3 │[0m
            4 │[0m basic    = "ɑ € 𐫱 ɑ€𐫱"
            5 │[0m raw      = 'ɑ € 𐫱 ɑ€𐫱'
            6 │[0m ml-basic = """ɑ € 𐫱 ɑ€𐫱"""
            7 │[0m ml-raw   = '''ɑ € 𐫱 ɑ€𐫱'''
    
         output from parser-cmd (PID 1666813) (stderr):
            1 │[0m line 6: invalid string: invalid syntax
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/multibyte-escape
         line 5: invalid string: invalid syntax
     
         Exit 1
    
         input sent to parser-cmd (PID 1666727):
            1 │[0m # Test each multibyte length: 2, 3, and 4 bytes:
            2 │[0m # ɑ € 𐫱
            3 │[0m
            4 │[0m basic-1    = "\u0251 \u20ac \U00010AF1 \u0251\u20ac\U00010AF1"
            5 │[0m ml-basic-1 = """\u0251 \u20ac \U00010AF1 \u0251\u20ac\U00010AF1"""
            6 │[0m
            7 │[0m # Again, but only using \U
            8 │[0m basic-2    = "\U00000251 \U000020ac \U00010AF1 \U00000251\U000020ac\U00010AF1"
            9 │[0m ml-basic-2 = """\U00000251 \U000020ac \U00010AF1 \U00000251\U000020ac\U00010AF1"""
    
         output from parser-cmd (PID 1666727) (stderr):
            1 │[0m line 5: invalid string: invalid syntax
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/multiline
         line 4: invalid string: invalid syntax
     
         Exit 1
    
         input sent to parser-cmd (PID 1666783):
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
    
         output from parser-cmd (PID 1666783) (stderr):
            1 │[0m line 4: invalid string: invalid syntax
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/multiline-empty
         line 1: invalid string: invalid syntax
     
         Exit 1
    
         input sent to parser-cmd (PID 1666817):
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
    
         output from parser-cmd (PID 1666817) (stderr):
            1 │[0m line 1: invalid string: invalid syntax
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/multiline-escaped-crlf
         line 3: invalid value "\"\"\"\\"
     
         Exit 1
    
         input sent to parser-cmd (PID 1666746):
            1 │[0m # The following line should be an unescaped backslash followed by a Windows
            2 │[0m # newline sequence ("\r\n")
            3 │[0m 0="""\
            4 │[0m """
    
         output from parser-cmd (PID 1666746) (stderr):
            1 │[0m line 3: invalid value "\"\"\"\\"
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/multiline-quotes
         line 9: invalid string: invalid syntax
     
         Exit 1
    
         input sent to parser-cmd (PID 1666816):
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
    
         output from parser-cmd (PID 1666816) (stderr):
            1 │[0m line 9: invalid string: invalid syntax
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/nl
         line 2: invalid string: invalid syntax
     
         Exit 1
    
         input sent to parser-cmd (PID 1666749):
            1 │[0m nl_mid = "val\nue"
            2 │[0m nl_end = """value\n"""
            3 │[0m
            4 │[0m lit_nl_end = '''value\n'''
            5 │[0m lit_nl_mid = 'val\nue'
            6 │[0m lit_nl_uni = 'val\ue'
    
         output from parser-cmd (PID 1666749) (stderr):
            1 │[0m line 2: invalid string: invalid syntax
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/raw-multiline
         line 6: expected key-value pair
     
         Exit 1
    
         input sent to parser-cmd (PID 1666786):
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
    
         output from parser-cmd (PID 1666786) (stderr):
            1 │[0m line 6: expected key-value pair
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/start-mb
         line 6: invalid value "\"\"\"\\"
     
         Exit 1
    
         input sent to parser-cmd (PID 1666782):
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
    
         output from parser-cmd (PID 1666782) (stderr):
            1 │[0m line 6: invalid value "\"\"\"\\"
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/unicode-escape
         line 9: invalid string: invalid syntax
     
         Exit 1
    
         input sent to parser-cmd (PID 1666812):
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
    
         output from parser-cmd (PID 1666812) (stderr):
            1 │[0m line 9: invalid string: invalid syntax
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/with-pound
         line 1: invalid value "\"We see no"
     
         Exit 1
    
         input sent to parser-cmd (PID 1666726):
            1 │[0m pound = "We see no # comments here."
            2 │[0m poundcomment = "But there are # some comments here." # Did I # mess you up?
    
         output from parser-cmd (PID 1666726) (stderr):
            1 │[0m line 1: invalid value "\"We see no"
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 861
      valid tests:  10 passed, 13 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['sh', 'sources/stage_test.sh', 'valid/string/**'], returncode=1, stdout='FAIL valid/string/ends-in-whitespace-escape\n     line 1: invalid string: invalid syntax\n \n     Exit 1\n\n     input sent to parser-cmd (PID 1666750):\n        1 │\x1b[0m beee = """\n        2 │\x1b[0m heeee\n        3 │\x1b[0m geeee\\\n        4 │\x1b[0m\n        5 │\x1b[0m\n        6 │\x1b[0m       … (+12018 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "strings-valid-suite.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: strings-invalid-suite (FEATURE-Strings.md)
  intent: The authoritative conformance suite rejects every invalid string case owned by this feature.
  return code: 1
  stdout:
    FAIL invalid/string/bad-escape-01
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1667198):
            1 │[0m invalid-escape = "This string has a bad \a escape character."
    
         output from parser-cmd (PID 1667198) (stdout):
            1 │[0m {
            2 │[0m   "invalid-escape": {"type": "string", "value": "This string has a bad \u0007 escape character."}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/string/basic-byte-escapes
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1667131):
            1 │[0m answer = "\x33"
    
         output from parser-cmd (PID 1667131) (stdout):
            1 │[0m {
            2 │[0m   "answer": {"type": "string", "value": "3"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/string/literal-multiline-quotes-01
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1667199):
            1 │[0m a = '''6 apostrophes: ''''''
    
         output from parser-cmd (PID 1667199) (stdout):
            1 │[0m {
            2 │[0m   "a": {"type": "string", "value": "''6 apostrophes: '''''"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/string/literal-multiline-quotes-02
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1667361):
            1 │[0m a = '''15 apostrophes: ''''''''''''''''''
    
         output from parser-cmd (PID 1667361) (stdout):
            1 │[0m {
            2 │[0m   "a": {"type": "string", "value": "''15 apostrophes: '''''''''''''''''"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/string/multiline-lit-no-close-02
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1667258):
            1 │[0m x='''
    
         output from parser-cmd (PID 1667258) (stdout):
            1 │[0m {
            2 │[0m   "x": {"type": "string", "value": "'"}
            3 │[0m }
    
         want:
           Exit code 1
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 807
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  72 passed,  5 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['sh', 'sources/stage_test.sh', 'invalid/string/**'], returncode=1, stdout='FAIL invalid/string/bad-escape-01\n     Expected an error, but no error was reported.\n\n     input sent to parser-cmd (PID 1667198):\n        1 │\x1b[0m invalid-escape = "This string has a bad \\a escape character."\n\n     output from parser-cmd (PID 1667198) (stdout):\n        1 │\x1b[0m {\n       … (+2206 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "strings-invalid-suite.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: strings-bounded-suites (FEATURE-Strings.md)
  intent: Both string conformance invocations complete successfully as separate executable checks.
  return code: 1
  stdout:
    FAIL valid/string/ends-in-whitespace-escape
         line 1: invalid string: invalid syntax
     
         Exit 1
    
         input sent to parser-cmd (PID 1667702):
            1 │[0m beee = """
            2 │[0m heeee
            3 │[0m geeee\
            4 │[0m
            5 │[0m
            6 │[0m       """
    
         output from parser-cmd (PID 1667702) (stderr):
            1 │[0m line 1: invalid string: invalid syntax
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/escape-tricky
         line 4: invalid string: invalid syntax
     
         Exit 1
    
         input sent to parser-cmd (PID 1667631):
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
    
         output from parser-cmd (PID 1667631) (stderr):
            1 │[0m line 4: invalid string: invalid syntax
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/multibyte
         line 6: invalid string: invalid syntax
     
         Exit 1
    
         input sent to parser-cmd (PID 1667698):
            1 │[0m # Test each multibyte length: 2, 3, and 4 bytes:
            2 │[0m # ɑ € 𐫱
            3 │[0m
            4 │[0m basic    = "ɑ € 𐫱 ɑ€𐫱"
            5 │[0m raw      = 'ɑ € 𐫱 ɑ€𐫱'
            6 │[0m ml-basic = """ɑ € 𐫱 ɑ€𐫱"""
            7 │[0m ml-raw   = '''ɑ € 𐫱 ɑ€𐫱'''
    
         output from parser-cmd (PID 1667698) (stderr):
            1 │[0m line 6: invalid string: invalid syntax
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/multibyte-escape
         line 5: invalid string: invalid syntax
     
         Exit 1
    
         input sent to parser-cmd (PID 1667660):
            1 │[0m # Test each multibyte length: 2, 3, and 4 bytes:
            2 │[0m # ɑ € 𐫱
            3 │[0m
            4 │[0m basic-1    = "\u0251 \u20ac \U00010AF1 \u0251\u20ac\U00010AF1"
            5 │[0m ml-basic-1 = """\u0251 \u20ac \U00010AF1 \u0251\u20ac\U00010AF1"""
            6 │[0m
            7 │[0m # Again, but only using \U
            8 │[0m basic-2    = "\U00000251 \U000020ac \U00010AF1 \U00000251\U000020ac\U00010AF1"
            9 │[0m ml-basic-2 = """\U00000251 \U000020ac \U00010AF1 \U00000251\U000020ac\U00010AF1"""
    
         output from parser-cmd (PID 1667660) (stderr):
            1 │[0m line 5: invalid string: invalid syntax
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/multiline
         line 4: invalid string: invalid syntax
     
         Exit 1
    
         input sent to parser-cmd (PID 1667695):
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
    
         output from parser-cmd (PID 1667695) (stderr):
            1 │[0m line 4: invalid string: invalid syntax
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/multiline-empty
         line 1: invalid string: invalid syntax
     
         Exit 1
    
         input sent to parser-cmd (PID 1667697):
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
    
         output from parser-cmd (PID 1667697) (stderr):
            1 │[0m line 1: invalid string: invalid syntax
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/multiline-escaped-crlf
         line 3: invalid value "\"\"\"\\"
     
         Exit 1
    
         input sent to parser-cmd (PID 1667696):
            1 │[0m # The following line should be an unescaped backslash followed by a Windows
            2 │[0m # newline sequence ("\r\n")
            3 │[0m 0="""\
            4 │[0m """
    
         output from parser-cmd (PID 1667696) (stderr):
            1 │[0m line 3: invalid value "\"\"\"\\"
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/multiline-quotes
         line 9: invalid string: invalid syntax
     
         Exit 1
    
         input sent to parser-cmd (PID 1667659):
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
    
         output from parser-cmd (PID 1667659) (stderr):
            1 │[0m line 9: invalid string: invalid syntax
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/nl
         line 2: invalid string: invalid syntax
     
         Exit 1
    
         input sent to parser-cmd (PID 1667704):
            1 │[0m nl_mid = "val\nue"
            2 │[0m nl_end = """value\n"""
            3 │[0m
            4 │[0m lit_nl_end = '''value\n'''
            5 │[0m lit_nl_mid = 'val\nue'
            6 │[0m lit_nl_uni = 'val\ue'
    
         output from parser-cmd (PID 1667704) (stderr):
            1 │[0m line 2: invalid string: invalid syntax
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/raw-multiline
         line 6: expected key-value pair
     
         Exit 1
    
         input sent to parser-cmd (PID 1667699):
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
    
         output from parser-cmd (PID 1667699) (stderr):
            1 │[0m line 6: expected key-value pair
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/start-mb
         line 6: invalid value "\"\"\"\\"
     
         Exit 1
    
         input sent to parser-cmd (PID 1667662):
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
    
         output from parser-cmd (PID 1667662) (stderr):
            1 │[0m line 6: invalid value "\"\"\"\\"
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/unicode-escape
         line 9: invalid string: invalid syntax
     
         Exit 1
    
         input sent to parser-cmd (PID 1667663):
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
    
         output from parser-cmd (PID 1667663) (stderr):
            1 │[0m line 9: invalid string: invalid syntax
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/with-pound
         line 1: invalid value "\"We see no"
     
         Exit 1
    
         input sent to parser-cmd (PID 1667630):
            1 │[0m pound = "We see no # comments here."
            2 │[0m poundcomment = "But there are # some comments here." # Did I # mess you up?
    
         output from parser-cmd (PID 1667630) (stderr):
            1 │[0m line 1: invalid value "\"We see no"
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 861
      valid tests:  10 passed, 13 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['sh', 'sources/stage_test.sh', 'valid/string/**'], returncode=1, stdout='FAIL valid/string/ends-in-whitespace-escape\n     line 1: invalid string: invalid syntax\n \n     Exit 1\n\n     input sent to parser-cmd (PID 1667702):\n        1 │\x1b[0m beee = """\n        2 │\x1b[0m heeee\n        3 │\x1b[0m geeee\\\n        4 │\x1b[0m\n        5 │\x1b[0m\n        6 │\x1b[0m       … (+12018 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "strings-bounded-suites.py", line 12, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: keys-valid (FEATURE-Keys.md)
  intent: The decoder passes all authoritative valid TOML key cases owned by this feature.
  return code: 1
  stdout:
    FAIL valid/key/alphanum
         line 10: tables are not supported yet
     
         Exit 1
    
         input sent to parser-cmd (PID 1667968):
            1 │[0m alpha = "a"
            2 │[0m 123 = "num"
            3 │[0m 000111 = "leading"
            4 │[0m 10e3 = "false float"
            5 │[0m one1two2 = "mixed"
            6 │[0m with-dash = "dashed"
            7 │[0m under_score = "___"
            8 │[0m 34-11 = 23
            9 │[0m
           10 │[0m [2018_10]
           11 │[0m 001 = 1
           12 │[0m
           13 │[0m [a-a-a]
           14 │[0m _ = false
    
         output from parser-cmd (PID 1667968) (stderr):
            1 │[0m line 10: tables are not supported yet
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/key/case-sensitive
         line 3: tables are not supported yet
     
         Exit 1
    
         input sent to parser-cmd (PID 1667974):
            1 │[0m sectioN = "NN"
            2 │[0m
            3 │[0m [section]
            4 │[0m name = "lower"
            5 │[0m NAME = "upper"
            6 │[0m Name = "capitalized"
            7 │[0m
            8 │[0m [Section]
            9 │[0m name = "different section!!"
           10 │[0m "μ" = "greek small letter mu"
           11 │[0m "Μ" = "greek capital letter MU"
           12 │[0m M = "latin letter M"
    
         output from parser-cmd (PID 1667974) (stderr):
            1 │[0m line 3: tables are not supported yet
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/key/dotted-01
         Could not find key "many" in parser output.
    
         input sent to parser-cmd (PID 1667973):
            1 │[0m name.first = "Arthur"
            2 │[0m "name".'last' = "Dent"
            3 │[0m
            4 │[0m many.dots.dot.dot.dot = 42
    
         output from parser-cmd (PID 1667973) (stdout):
            1 │[0m {
            2 │[0m   "\"name\".'last'":       {"type": "string", "value": "Dent"},
            3 │[0m   "many.dots.dot.dot.dot": {"type": "integer", "value": "42"},
            4 │[0m   "name.first":            {"type": "string", "value": "Arthur"}
            5 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "many": {
            3 │[0m         "dots": {
            4 │[0m             "dot": {
            5 │[0m                 "dot": {
            6 │[0m                     "dot": {"type": "integer", "value": "42"}
            7 │[0m                 }
            8 │[0m             }
            9 │[0m         }
           10 │[0m     },
           11 │[0m     "name": {
           12 │[0m         "first": {"type": "string", "value": "Arthur"},
           13 │[0m         "last":  {"type": "string", "value": "Dent"}
           14 │[0m     }
           15 │[0m }
    
    FAIL valid/key/dotted-02
         Could not find key "count" in parser output.
    
         input sent to parser-cmd (PID 1667889):
            1 │[0m # Note: this file contains literal tab characters.
            2 │[0m
            3 │[0m # Space are ignored, and key parts can be quoted.
            4 │[0m count.a       = 1
            5 │[0m count . b     = 2
            6 │[0m "count"."c"   = 3
            7 │[0m "count" . "d" = 4
            8 │[0m 'count'.'e'   = 5
            9 │[0m 'count' . 'f' = 6
           10 │[0m "count".'g'   = 7
           11 │[0m "count" . 'h' = 8
           12 │[0m count.'i'     = 9
           13 │[0m count 	.	 'j'	   = 10
           14 │[0m "count".k     = 11
           15 │[0m "count" . l   = 12
    
         output from parser-cmd (PID 1667889) (stdout):
            1 │[0m {
            2 │[0m   "\"count\" . \"d\"": {"type": "integer", "value": "4"},
            3 │[0m   "\"count\" . 'h'": {"type": "integer", "value": "8"},
            4 │[0m   "\"count\" . l":   {"type": "integer", "value": "12"},
            5 │[0m   "\"count\".\"c\"": {"type": "integer", "value": "3"},
            6 │[0m   "\"count\".'g'":   {"type": "integer", "value": "7"},
            7 │[0m   "\"count\".k":     {"type": "integer", "value": "11"},
            8 │[0m   "'count' . 'f'":   {"type": "integer", "value": "6"},
            9 │[0m   "'count'.'e'":     {"type": "integer", "value": "5"},
           10 │[0m   "count \t.\t 'j'": {"type": "integer", "value": "10"},
           11 │[0m   "count . b":       {"type": "integer", "value": "2"},
           12 │[0m   "count.'i'":       {"type": "integer", "value": "9"},
           13 │[0m   "count.a":         {"type": "integer", "value": "1"}
           14 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "count": {
            3 │[0m         "a": {"type": "integer", "value": "1"},
            4 │[0m         "b": {"type": "integer", "value": "2"},
            5 │[0m         "c": {"type": "integer", "value": "3"},
            6 │[0m         "d": {"type": "integer", "value": "4"},
            7 │[0m         "e": {"type": "integer", "value": "5"},
            8 │[0m         "f": {"type": "integer", "value": "6"},
            9 │[0m         "g": {"type": "integer", "value": "7"},
           10 │[0m         "h": {"type": "integer", "value": "8"},
           11 │[0m         "i": {"type": "integer", "value": "9"},
           12 │[0m         "j": {"type": "integer", "value": "10"},
           13 │[0m         "k": {"type": "integer", "value": "11"},
           14 │[0m         "l": {"type": "integer", "value": "12"}
           15 │[0m     }
           16 │[0m }
    
    FAIL valid/key/dotted-03
         line 3: tables are not supported yet
     
         Exit 1
    
         input sent to parser-cmd (PID 1667935):
            1 │[0m top.key = 1
            2 │[0m
            3 │[0m [tbl]
            4 │[0m a.b.c = 42.666
            5 │[0m
            6 │[0m [a.few.dots]
            7 │[0m polka.dot = "again?"
            8 │[0m polka.dance-with = "Dot"
    
         output from parser-cmd (PID 1667935) (stderr):
            1 │[0m line 3: tables are not supported yet
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/key/dotted-04
         line 3: tables are not supported yet
     
         Exit 1
    
         input sent to parser-cmd (PID 1667933):
            1 │[0m top.key = 1
            2 │[0m
            3 │[0m [[arr]]
            4 │[0m a.b.c=1
            5 │[0m a.b.d=2
            6 │[0m
            7 │[0m [[arr]]
            8 │[0m a.b.c=3
            9 │[0m a.b.d=4
    
         output from parser-cmd (PID 1667933) (stderr):
            1 │[0m line 3: tables are not supported yet
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/key/dotted-empty
         line 3: tables are not supported yet
     
         Exit 1
    
         input sent to parser-cmd (PID 1667972):
            1 │[0m ''.x = "empty.x"
            2 │[0m x."" = "x.empty"
            3 │[0m [a]
            4 │[0m "".'' = "empty.empty"
    
         output from parser-cmd (PID 1667972) (stderr):
            1 │[0m line 3: tables are not supported yet
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/key/empty-01
         Could not find key "" in parser output.
    
         input sent to parser-cmd (PID 1667939):
            1 │[0m "" = "blank"
    
         output from parser-cmd (PID 1667939) (stdout):
            1 │[0m {
            2 │[0m   "\"\"": {"type": "string", "value": "blank"}
            3 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "": {"type": "string", "value": "blank"}
            3 │[0m }
    
    FAIL valid/key/empty-02
         Could not find key "" in parser output.
    
         input sent to parser-cmd (PID 1667975):
            1 │[0m '' = "blank"
    
         output from parser-cmd (PID 1667975) (stdout):
            1 │[0m {
            2 │[0m   "''": {"type": "string", "value": "blank"}
            3 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "": {"type": "string", "value": "blank"}
            3 │[0m }
    
    FAIL valid/key/empty-03
         Could not find key "" in parser output.
    
         input sent to parser-cmd (PID 1667986):
            1 │[0m ''=0
    
         output from parser-cmd (PID 1667986) (stdout):
            1 │[0m {
            2 │[0m   "''": {"type": "integer", "value": "0"}
            3 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "": {"type": "integer", "value": "0"}
            3 │[0m }
    
    FAIL valid/key/escapes
         line 6: tables are not supported yet
     
         Exit 1
    
         input sent to parser-cmd (PID 1667934):
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
    
         output from parser-cmd (PID 1667934) (stderr):
            1 │[0m line 6: tables are not supported yet
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/key/like-date
         line 18: tables are not supported yet
     
         Exit 1
    
         input sent to parser-cmd (PID 1667982):
            1 │[0m # '-' is a valid character in keys: make a key that looks like a date.
            2 │[0m 2001-02-03   = 1
            3 │[0m "2001-02-04" = 2
            4 │[0m '2001-02-05' = 3
            5 │[0m
            6 │[0m # Also include datetime and time for good measure; these need to be quoted as
            7 │[0m # ':' isn't a valid bare key.
            8 │[0m "2001-02-06T15:16:17+01:00" = 4
            9 │[0m "2001-02-07T15:16:17"       = 5
           10 │[0m "15:16:17"                  = 6
           11 │[0m
           12 │[0m # Dotted keys
           13 │[0m a.2001-02-08 = 7
           14 │[0m a.2001-02-09.2001-02-10 = 8
           15 │[0m 2001-02-11.a.2001-02-12 = 9
           16 │[0m
           17 │[0m # Table names
           18 │[0m [2002-01-02]
           19 │[0m k = 10
           20 │[0m
           21 │[0m [2002-01-02.2024-01-03]
           22 │[0m k = 11
           23 │[0m
           24 │[0m [[2002-01-04]]
           25 │[0m k = 12
    
         output from parser-cmd (PID 1667982) (stderr):
            1 │[0m line 18: tables are not supported yet
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/key/numeric-02
         Could not find key "1" in parser output.
    
         input sent to parser-cmd (PID 1667919):
            1 │[0m 1.2   = true
    
         output from parser-cmd (PID 1667919) (stdout):
            1 │[0m {
            2 │[0m   "1.2": {"type": "bool", "value": "true"}
            3 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "1": {
            3 │[0m         "2": {"type": "bool", "value": "true"}
            4 │[0m     }
            5 │[0m }
    
    FAIL valid/key/numeric-04
         Could not find key "01" in parser output.
    
         input sent to parser-cmd (PID 1667937):
            1 │[0m 01.23 = true
    
         output from parser-cmd (PID 1667937) (stdout):
            1 │[0m {
            2 │[0m   "01.23": {"type": "bool", "value": "true"}
            3 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "01": {
            3 │[0m         "23": {"type": "bool", "value": "true"}
            4 │[0m     }
            5 │[0m }
    
    FAIL valid/key/numeric-05
         Could not find key "23" in parser output.
    
         input sent to parser-cmd (PID 1667936):
            1 │[0m 23.01 = true
    
         output from parser-cmd (PID 1667936) (stdout):
            1 │[0m {
            2 │[0m   "23.01": {"type": "bool", "value": "true"}
            3 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "23": {
            3 │[0m         "01": {"type": "bool", "value": "true"}
            4 │[0m     }
            5 │[0m }
    
    FAIL valid/key/quoted-dots
         line 4: tables are not supported yet
     
         Exit 1
    
         input sent to parser-cmd (PID 1667918):
            1 │[0m plain = 1
            2 │[0m "with.dot" = 2
            3 │[0m
            4 │[0m [plain_table]
            5 │[0m plain = 3
            6 │[0m "with.dot" = 4
            7 │[0m
            8 │[0m [table.withdot]
            9 │[0m plain = 5
           10 │[0m "key.with.dots" = 6
           11 │[0m "escaped\u002edot" = 7
    
         output from parser-cmd (PID 1667918) (stderr):
            1 │[0m line 4: tables are not supported yet
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/key/quoted-unicode
         Could not find key "\x00" in parser output.
    
         input sent to parser-cmd (PID 1667938):
            1 │[0m
            2 │[0m "\u0000" = "null"
            3 │[0m '\u0000' = "different key"
            4 │[0m "\u0008 \u000c \U00000041 \u007f \u0080 \u00ff \ud7ff \ue000 \uffff \U00010000 \U0010ffff" = "escaped key"
            5 │[0m
            6 │[0m "~  ÿ ퟿  ￿ 𐀀 􏿿" = "basic key"
            7 │[0m 'l ~  ÿ ퟿  ￿ 𐀀 􏿿' = "literal key"
    
         output from parser-cmd (PID 1667938) (stdout):
            1 │[0m {
            2 │[0m   "\"\\u0000\"":                                                                                {"type": "string", "value": "null"},
            3 │[0m   "\"\\u0008 \\u000c \\U00000041 \\u007f \\u0080 \\u00ff \\ud7ff \\ue000 \\uffff \\U00010000 \\U0010ffff\"": {"type": "string", "value": "escaped key"},
            4 │[0m   "\"~  ÿ ퟿  ￿ 𐀀 􏿿\"":                                                                          {"type": "string", "value": "basic key"},
            5 │[0m   "'\\u0000'":                                                                                  {"type": "string", "value": "different key"},
            6 │[0m   "'l ~  ÿ ퟿  ￿ 𐀀 􏿿'":                                                                          {"type": "string", "value": "literal key"}
            7 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "\u0000":               {"type": "string", "value": "null"},
            3 │[0m     "\b \f A   ÿ ퟿  ￿ 𐀀 􏿿": {"type": "string", "value": "escaped key"},
            4 │[0m     "\\u0000":              {"type": "string", "value": "different key"},
            5 │[0m     "l ~  ÿ ퟿  ￿ 𐀀 􏿿":      {"type": "string", "value": "literal key"},
            6 │[0m     "~  ÿ ퟿  ￿ 𐀀 􏿿":        {"type": "string", "value": "basic key"}
            7 │[0m }
    
    FAIL valid/key/space
         line 6: tables are not supported yet
     
         Exit 1
    
         input sent to parser-cmd (PID 1667971):
            1 │[0m # Keep whitespace inside quotes keys at all positions.
            2 │[0m "a b"   = 1
            3 │[0m " c d " = 2
            4 │[0m "  much 		  whitespace  	\n  \r\n  " = 3
            5 │[0m
            6 │[0m [ " tbl " ]
            7 │[0m "\ttab\ttab\t" = "tab"
    
         output from parser-cmd (PID 1667971) (stderr):
            1 │[0m line 6: tables are not supported yet
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/key/special-chars
         line 1: invalid value "~!@$^&*()_+-`1234567890[]|/?><.,;:'=\" = 1"
     
         Exit 1
    
         input sent to parser-cmd (PID 1667920):
            1 │[0m "=~!@$^&*()_+-`1234567890[]|/?><.,;:'=" = 1
    
         output from parser-cmd (PID 1667920) (stderr):
            1 │[0m line 1: invalid value "~!@$^&*()_+-`1234567890[]|/?><.,;:'=\" = 1"
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/key/start
         line 4: tables are not supported yet
     
         Exit 1
    
         input sent to parser-cmd (PID 1667970):
            1 │[0m # Table and keys can start with any character; there is no requirement for it to
            2 │[0m # start with a letter.
            3 │[0m
            4 │[0m [-key]
            5 │[0m -key = 1
            6 │[0m
            7 │[0m [_key]
            8 │[0m _key = 2
            9 │[0m
           10 │[0m [1key]
           11 │[0m 1key = 3
           12 │[0m
           13 │[0m [-]
           14 │[0m - = 4
           15 │[0m
           16 │[0m [_]
           17 │[0m _ = 5
           18 │[0m
           19 │[0m [1]
           20 │[0m 1 = 6
           21 │[0m
           22 │[0m [---]
           23 │[0m --- = 7
           24 │[0m
           25 │[0m [___]
           26 │[0m ___ = 8
           27 │[0m
           28 │[0m [111]
           29 │[0m 111 = 9
           30 │[0m
           31 │[0m [inline]
           32 │[0m --- = {--- = 10, ___ = 11, 111 = 12}
    
         output from parser-cmd (PID 1667970) (stderr):
            1 │[0m line 4: tables are not supported yet
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 856
      valid tests:   8 passed, 20 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['sh', 'sources/stage_test.sh', 'valid/key/**'], returncode=1, stdout='FAIL valid/key/alphanum\n     line 10: tables are not supported yet\n \n     Exit 1\n\n     input sent to parser-cmd (PID 1667968):\n        1 │\x1b[0m alpha = "a"\n        2 │\x1b[0m 123 = "num"\n        3 │\x1b[0m 000111 = "leading"\n        4 │\x1b[0m 10e3 = "false float"\n        5 │\x1b[0m one1two2 = … (+17319 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "keys-valid.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: keys-invalid (FEATURE-Keys.md)
  intent: The decoder rejects all authoritative invalid TOML key cases owned by this feature.
  return code: 1
  stdout:
    FAIL invalid/key/bare-invalid-character-01
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1668300):
            1 │[0m ! = 123
    
         output from parser-cmd (PID 1668300) (stdout):
            1 │[0m {
            2 │[0m   "!": {"type": "integer", "value": "123"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/key/bare-invalid-character-02
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1668241):
            1 │[0m bare!key = 123
    
         output from parser-cmd (PID 1668241) (stdout):
            1 │[0m {
            2 │[0m   "bare!key": {"type": "integer", "value": "123"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/key/dot
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1668296):
            1 │[0m . = 1
    
         output from parser-cmd (PID 1668296) (stdout):
            1 │[0m {
            2 │[0m   ".": {"type": "integer", "value": "1"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/key/dotdot
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1668238):
            1 │[0m .. = 1
    
         output from parser-cmd (PID 1668238) (stdout):
            1 │[0m {
            2 │[0m   "..": {"type": "integer", "value": "1"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/key/dotted-redefine-table-01
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1668215):
            1 │[0m a = false
            2 │[0m a.b = true
    
         output from parser-cmd (PID 1668215) (stdout):
            1 │[0m {
            2 │[0m   "a":   {"type": "bool", "value": "false"},
            3 │[0m   "a.b": {"type": "bool", "value": "true"}
            4 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/key/dotted-redefine-table-02
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1668246):
            1 │[0m # Defined a.b as int
            2 │[0m a.b = 1
            3 │[0m # Tries to access it as table: error
            4 │[0m a.b.c = 2
    
         output from parser-cmd (PID 1668246) (stdout):
            1 │[0m {
            2 │[0m   "a.b":   {"type": "integer", "value": "1"},
            3 │[0m   "a.b.c": {"type": "integer", "value": "2"}
            4 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/key/duplicate-keys-03
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1668240):
            1 │[0m spelling   = "favorite"
            2 │[0m "spelling" = "favourite"
    
         output from parser-cmd (PID 1668240) (stdout):
            1 │[0m {
            2 │[0m   "\"spelling\"": {"type": "string", "value": "favourite"},
            3 │[0m   "spelling":   {"type": "string", "value": "favorite"}
            4 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/key/duplicate-keys-04
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1668295):
            1 │[0m spelling   = "favorite"
            2 │[0m 'spelling' = "favourite"
    
         output from parser-cmd (PID 1668295) (stdout):
            1 │[0m {
            2 │[0m   "'spelling'": {"type": "string", "value": "favourite"},
            3 │[0m   "spelling":   {"type": "string", "value": "favorite"}
            4 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/key/duplicate-keys-05
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1668197):
            1 │[0m a        = 1
            2 │[0m "\u0061" = 1
    
         output from parser-cmd (PID 1668197) (stdout):
            1 │[0m {
            2 │[0m   "\"\\u0061\"": {"type": "integer", "value": "1"},
            3 │[0m   "a":        {"type": "integer", "value": "1"}
            4 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/key/duplicate-keys-06
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1668363):
            1 │[0m "a'b"      = 1
            2 │[0m "a\u0027b" = 2
    
         output from parser-cmd (PID 1668363) (stdout):
            1 │[0m {
            2 │[0m   "\"a'b\"":    {"type": "integer", "value": "1"},
            3 │[0m   "\"a\\u0027b\"": {"type": "integer", "value": "2"}
            4 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/key/escape
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1668239):
            1 │[0m \u00c0 = "latin capital letter A with grave"
    
         output from parser-cmd (PID 1668239) (stdout):
            1 │[0m {
            2 │[0m   "\\u00c0": {"type": "string", "value": "latin capital letter A with grave"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/key/multiline-key-01
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1668247):
            1 │[0m """key""" = 1
    
         output from parser-cmd (PID 1668247) (stdout):
            1 │[0m {
            2 │[0m   "\"\"\"key\"\"\"": {"type": "integer", "value": "1"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/key/multiline-key-02
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1668220):
            1 │[0m '''key''' = 1
    
         output from parser-cmd (PID 1668220) (stdout):
            1 │[0m {
            2 │[0m   "'''key'''": {"type": "integer", "value": "1"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/key/multiline-key-04
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1668244):
            1 │[0m '''key''' = '''v'''
    
         output from parser-cmd (PID 1668244) (stdout):
            1 │[0m {
            2 │[0m   "'''key'''": {"type": "string", "value": "''v''"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/key/partial-quoted
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1668490):
            1 │[0m partial"quoted" = 5
    
         output from parser-cmd (PID 1668490) (stdout):
            1 │[0m {
            2 │[0m   "partial\"quoted\"": {"type": "integer", "value": "5"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/key/space
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1668420):
            1 │[0m a b = 1
    
         output from parser-cmd (PID 1668420) (stdout):
            1 │[0m {
            2 │[0m   "a b": {"type": "integer", "value": "1"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/key/special-character
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1668427):
            1 │[0m μ = "greek small letter mu"
    
         output from parser-cmd (PID 1668427) (stdout):
            1 │[0m {
            2 │[0m   "μ": {"type": "string", "value": "greek small letter mu"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/key/start-dot
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1668474):
            1 │[0m .key = 1
    
         output from parser-cmd (PID 1668474) (stdout):
            1 │[0m {
            2 │[0m   ".key": {"type": "integer", "value": "1"}
            3 │[0m }
    
         want:
           Exit code 1
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 823
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  43 passed, 18 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['sh', 'sources/stage_test.sh', 'invalid/key/**'], returncode=1, stdout='FAIL invalid/key/bare-invalid-character-01\n     Expected an error, but no error was reported.\n\n     input sent to parser-cmd (PID 1668300):\n        1 │\x1b[0m ! = 123\n\n     output from parser-cmd (PID 1668300) (stdout):\n        1 │\x1b[0m {\n        2 │\x1b[0m   "!": {"type": "integer", "value": "… (+7655 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "keys-invalid.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- PASS: strings-valid-suite (FEATURE-Strings.md)
  intent: The authoritative conformance suite passes every valid string case owned by this feature.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 861
      valid tests:  23 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: strings-invalid-suite (FEATURE-Strings.md)
  intent: The authoritative conformance suite rejects every invalid string case owned by this feature.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 807
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  77 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: strings-bounded-suites (FEATURE-Strings.md)
  intent: Both string conformance invocations complete successfully as separate executable checks.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 861
      valid tests:  23 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 807
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  77 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: keys-valid (FEATURE-Keys.md)
  intent: The decoder passes all authoritative valid TOML key cases owned by this feature.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 856
      valid tests:  28 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: keys-invalid (FEATURE-Keys.md)
  intent: The decoder rejects all authoritative invalid TOML key cases owned by this feature.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.121804/build/toml/toml-decoder] [no encoder]
    skipped tests: 823
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  61 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true

## Repair attempts
- attempt 0 (initial build): failed; 2/5 checks; 138/138 cases model=gpt-5.6-luna; execution 20260816.122829.516Z-b05144b1
- attempt 1 (repair 1): built; 5/5 checks; 215/215 cases model=gpt-5.6-luna; execution 20260816.123439.459Z-c6e09f01

## Build summary
<reusable-compact filename="stage_test.sh">
Runs one scoped TOML conformance slice. Requires exactly one `-run` pattern argument, builds `./cmd/toml-decoder`, invokes `sources/run_conformance.sh`, and uses the command exit status as the verdict. Defaults `TOML_TEST_VERSION` to `v2.2.0`.
</reusable-compact>

RESULT: SUCCESS

FILES CHANGED:
- internal/toml/toml.go
- toml-decoder

SUMMARY:
Fixed multiline quote-run parsing and line-comment boundary handling. Scoped acceptance passes: strings valid 23/23, strings invalid 77/77, keys valid 28/28, keys invalid 61/61. Go tests and vet pass.

BLOCKERS:
- Full conformance remains outside this block’s string/key scope and still has unrelated array, scalar, datetime, and table failures.
