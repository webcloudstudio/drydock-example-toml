# Evidence: Block 5 · Service (block-5)

- block type: block
- date: 2026-08-14
- resulting state: closed/verified
- story points (combined assembled cost): 8629
- execution id: 20260814.052011.978Z-ac5a3505

## Stories built
- Parse TOML basic, multiline basic, literal, and multiline literal strings. (lexical-strings) [story]

## Reusable compacts
- stage_test_compact.md

## Missing context
- context: toml-v1.0.0.md

## Stacked context
- compass: COMPASS.md (SP 2629)
- implements: FEATURE-Lexical-Strings.md (SP 846)
- context: stage_test.sh (SP 161)
- context: ARCHITECTURE_compact.md (SP 108)
- stack: go.md (SP 3646)
- stack: common_compact.md (SP 1179)

## Build directory changes
- internal/toml/decode.go
- internal/toml/value_test.go
- toml-decoder

## Pre-build acceptance observation
- RED: strings-conformance (FEATURE-Lexical-Strings.md)
  intent: The implementation passes the authoritative TOML 1.0.0 string conformance slice.
  return code: 1
  stdout:
    FAIL valid/string/ends-in-whitespace-escape
         Values for key "beee" don't match:
           Expected:     heeee
         geeee
           Your encoder: heeee
         geeee\
    
         input sent to parser-cmd (PID 758702):
            1 │[0m beee = """
            2 │[0m heeee
            3 │[0m geeee\
            4 │[0m
            5 │[0m
            6 │[0m       """
    
         output from parser-cmd (PID 758702) (stdout):
            1 │[0m {
            2 │[0m   "beee": {"type": "string", "value": "heeee\ngeeee\\  \n\n\n      "}
            3 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "beee": {"type": "string", "value": "heeee\ngeeee"}
            3 │[0m }
    
    FAIL valid/string/escape-tricky
         Values for key "lit_multiline_not_unicode" don't match:
           Expected:     \u007f
           Your encoder:
         \u007f
    
         input sent to parser-cmd (PID 758612):
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
    
         output from parser-cmd (PID 758612) (stdout):
            1 │[0m {
            2 │[0m   "end_esc":                   {"type": "string", "value": "String does not end here\" but ends here\\"},
            3 │[0m   "lit_end_esc":               {"type": "string", "value": "String ends here\\"},
            4 │[0m   "lit_multiline_end":         {"type": "string", "value": "There is no escape\\"},
            5 │[0m   "lit_multiline_not_unicode": {"type": "string", "value": "\n\\u007f"},
            6 │[0m   "multiline_end_esc":         {"type": "string", "value": "When will it end? \\\"\"\"...\"\"\\\" should be here\\\""},
            7 │[0m   "multiline_not_unicode":     {"type": "string", "value": "\\\\u0041"},
            8 │[0m   "multiline_unicode":         {"type": "string", "value": "\\u00a0"}
            9 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "end_esc":                   {"type": "string", "value": "String does not end here\" but ends here\\"},
            3 │[0m     "lit_end_esc":               {"type": "string", "value": "String ends here\\"},
            4 │[0m     "lit_multiline_end":         {"type": "string", "value": "There is no escape\\"},
            5 │[0m     "lit_multiline_not_unicode": {"type": "string", "value": "\\u007f"},
            6 │[0m     "multiline_end_esc":         {"type": "string", "value": "When will it end? \"\"\"...\"\"\" should be here\""},
            7 │[0m     "multiline_not_unicode":     {"type": "string", "value": "\\u0041"},
            8 │[0m     "multiline_unicode":         {"type": "string", "value": " "}
            9 │[0m }
    
    FAIL valid/string/multibyte-escape
         Values for key "ml-basic-1" don't match:
           Expected:     ɑ € 𐫱 ɑ€𐫱
           Your encoder: \u0251 \u20ac \U00010AF1 \u0251\u20ac\U00010AF1
    
         input sent to parser-cmd (PID 758592):
            1 │[0m # Test each multibyte length: 2, 3, and 4 bytes:
            2 │[0m # ɑ € 𐫱
            3 │[0m
            4 │[0m basic-1    = "\u0251 \u20ac \U00010AF1 \u0251\u20ac\U00010AF1"
            5 │[0m ml-basic-1 = """\u0251 \u20ac \U00010AF1 \u0251\u20ac\U00010AF1"""
            6 │[0m
            7 │[0m # Again, but only using \U
            8 │[0m basic-2    = "\U00000251 \U000020ac \U00010AF1 \U00000251\U000020ac\U00010AF1"
            9 │[0m ml-basic-2 = """\U00000251 \U000020ac \U00010AF1 \U00000251\U000020ac\U00010AF1"""
    
         output from parser-cmd (PID 758592) (stdout):
            1 │[0m {
            2 │[0m   "basic-1":    {"type": "string", "value": "ɑ € 𐫱 ɑ€𐫱"},
            3 │[0m   "basic-2":    {"type": "string", "value": "ɑ € 𐫱 ɑ€𐫱"},
            4 │[0m   "ml-basic-1": {"type": "string", "value": "\\u0251 \\u20ac \\U00010AF1 \\u0251\\u20ac\\U00010AF1"},
            5 │[0m   "ml-basic-2": {
            6 │[0m     "type":  "string",
            7 │[0m     "value": "\\U00000251 \\U000020ac \\U00010AF1 \\U00000251\\U000020ac\\U00010AF1"
            8 │[0m   }
            9 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "basic-1":    {"type": "string", "value": "ɑ € 𐫱 ɑ€𐫱"},
            3 │[0m     "ml-basic-1": {"type": "string", "value": "ɑ € 𐫱 ɑ€𐫱"},
            4 │[0m     "basic-2":    {"type": "string", "value": "ɑ € 𐫱 ɑ€𐫱"},
            5 │[0m     "ml-basic-2": {"type": "string", "value": "ɑ € 𐫱 ɑ€𐫱"}
            6 │[0m }
    
    FAIL valid/string/multiline
         Values for key "equivalent_three" don't match:
           Expected:     The quick brown fox jumps over the lazy dog.
           Your encoder:        The quick brown        fox jumps over        the lazy dog.
    
         input sent to parser-cmd (PID 758569):
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
    
         output from parser-cmd (PID 758569) (stdout):
            1 │[0m {
            2 │[0m   "equivalent_one": {"type": "string", "value": "The quick brown fox jumps over the lazy dog."},
            3 │[0m   "equivalent_two": {"type": "string", "value": "The quick brown \n\n  fox jumps over     the lazy dog."},
            4 │[0m   "escape-bs-1":    {"type": "string", "value": "a \\b"},
            5 │[0m   "escape-bs-2":    {"type": "string", "value": "a \\\\b"},
            6 │[0m   "escape-bs-3":    {"type": "string", "value": "a \\\\\\  b"},
            7 │[0m   "keep-ws-before": {"type": "string", "value": "a   \t   b"},
            8 │[0m   "no-space":       {"type": "string", "value": "a    b"},
            9 │[0m   "equivalent_three": {
           10 │[0m     "type":  "string",
           11 │[0m     "value": "       The quick brown        fox jumps over        the lazy dog.       "
           12 │[0m   },
           13 │[0m   "whitespace-after-bs": {
           14 │[0m     "type":  "string",
           15 │[0m     "value": "       The quick brown        fox jumps over \\   \n       the lazy dog.\\\t\n       "
           16 │[0m   }
           17 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "equivalent_one":      {"type": "string", "value": "The quick brown fox jumps over the lazy dog."},
            3 │[0m     "equivalent_three":    {"type": "string", "value": "The quick brown fox jumps over the lazy dog."},
            4 │[0m     "equivalent_two":      {"type": "string", "value": "The quick brown fox jumps over the lazy dog."},
            5 │[0m     "escape-bs-1":         {"type": "string", "value": "a \\\nb"},
            6 │[0m     "escape-bs-2":         {"type": "string", "value": "a \\b"},
            7 │[0m     "escape-bs-3":         {"type": "string", "value": "a \\\\\n  b"},
            8 │[0m     "keep-ws-before":      {"type": "string", "value": "a   \tb"},
            9 │[0m     "no-space":            {"type": "string", "value": "ab"},
           10 │[0m     "whitespace-after-bs": {"type": "string", "value": "The quick brown fox jumps over the lazy dog."}
           11 │[0m }
    
    FAIL valid/string/multiline-empty
         Values for key "empty-3" don't match:
           Expected:
           Your encoder:
    
         input sent to parser-cmd (PID 758588):
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
    
         output from parser-cmd (PID 758588) (stdout):
            1 │[0m {
            2 │[0m   "empty-1": {"type": "string", "value": ""},
            3 │[0m   "empty-2": {"type": "string", "value": ""},
            4 │[0m   "empty-3": {"type": "string", "value": "    "},
            5 │[0m   "empty-4": {"type": "string", "value": "      \\  \n   "}
            6 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "empty-1": {"type": "string", "value": ""},
            3 │[0m     "empty-2": {"type": "string", "value": ""},
            4 │[0m     "empty-3": {"type": "string", "value": ""},
            5 │[0m     "empty-4": {"type": "string", "value": ""}
            6 │[0m }
    
    FAIL valid/string/multiline-escaped-crlf
         Values for key "0" don't match:
           Expected:
           Your encoder: \
    
         input sent to parser-cmd (PID 758587):
            1 │[0m # The following line should be an unescaped backslash followed by a Windows
            2 │[0m # newline sequence ("\r\n")
            3 │[0m 0="""\
            4 │[0m """
    
         output from parser-cmd (PID 758587) (stdout):
            1 │[0m {
            2 │[0m   "0": {"type": "string", "value": "\\\r\n"}
            3 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "0": {"type": "string", "value": ""}
            3 │[0m }
    
    FAIL valid/string/multiline-quotes
         unterminated value
     
         Exit 1
    
         input sent to parser-cmd (PID 758609):
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
    
         output from parser-cmd (PID 758609) (stderr):
            1 │[0m unterminated value
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/nl
         Values for key "nl_end" don't match:
           Expected:     value
     
           Your encoder: value\n
    
         input sent to parser-cmd (PID 758590):
            1 │[0m nl_mid = "val\nue"
            2 │[0m nl_end = """value\n"""
            3 │[0m
            4 │[0m lit_nl_end = '''value\n'''
            5 │[0m lit_nl_mid = 'val\nue'
            6 │[0m lit_nl_uni = 'val\ue'
    
         output from parser-cmd (PID 758590) (stdout):
            1 │[0m {
            2 │[0m   "lit_nl_end": {"type": "string", "value": "value\\n"},
            3 │[0m   "lit_nl_mid": {"type": "string", "value": "val\\nue"},
            4 │[0m   "lit_nl_uni": {"type": "string", "value": "val\\ue"},
            5 │[0m   "nl_end":     {"type": "string", "value": "value\\n"},
            6 │[0m   "nl_mid":     {"type": "string", "value": "val\nue"}
            7 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "lit_nl_end": {"type": "string", "value": "value\\n"},
            3 │[0m     "lit_nl_mid": {"type": "string", "value": "val\\nue"},
            4 │[0m     "lit_nl_uni": {"type": "string", "value": "val\\ue"},
            5 │[0m     "nl_end":     {"type": "string", "value": "value\n"},
            6 │[0m     "nl_mid":     {"type": "string", "value": "val\nue"}
            7 │[0m }
    
    FAIL valid/string/raw-multiline
         unterminated value
     
         Exit 1
    
         input sent to parser-cmd (PID 758608):
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
    
         output from parser-cmd (PID 758608) (stderr):
            1 │[0m unterminated value
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/start-mb
         Values for key "s6" don't match:
           Expected:     §
           Your encoder:
         §
    
         input sent to parser-cmd (PID 758610):
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
    
         output from parser-cmd (PID 758610) (stdout):
            1 │[0m {
            2 │[0m   "s1": {"type": "string", "value": "§"},
            3 │[0m   "s2": {"type": "string", "value": "§"},
            4 │[0m   "s3": {"type": "string", "value": "§"},
            5 │[0m   "s4": {"type": "string", "value": "§"},
            6 │[0m   "s5": {"type": "string", "value": "§"},
            7 │[0m   "s6": {"type": "string", "value": "\n§"},
            8 │[0m   "s7": {"type": "string", "value": "§"}
            9 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "s1": {"type": "string", "value": "§"},
            3 │[0m     "s2": {"type": "string", "value": "§"},
            4 │[0m     "s3": {"type": "string", "value": "§"},
            5 │[0m     "s4": {"type": "string", "value": "§"},
            6 │[0m     "s5": {"type": "string", "value": "§"},
            7 │[0m     "s6": {"type": "string", "value": "§"},
            8 │[0m     "s7": {"type": "string", "value": "§"}
            9 │[0m }
    
    FAIL valid/string/unicode-escape
         Values for key "ml-a" don't match:
           Expected:     a
           Your encoder: \u0061
    
         input sent to parser-cmd (PID 758586):
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
    
         output from parser-cmd (PID 758586) (stdout):
            1 │[0m {
            2 │[0m   "a":          {"type": "string", "value": "a"},
            3 │[0m   "b":          {"type": "string", "value": "b"},
            4 │[0m   "c":          {"type": "string", "value": "c"},
            5 │[0m   "delta-1":    {"type": "string", "value": "δ"},
            6 │[0m   "delta-2":    {"type": "string", "value": "δ"},
            7 │[0m   "ml-a":       {"type": "string", "value": "\\u0061"},
            8 │[0m   "ml-b":       {"type": "string", "value": "\\u0062"},
            9 │[0m   "ml-c":       {"type": "string", "value": "\\U00000063"},
           10 │[0m   "ml-delta-1": {"type": "string", "value": "\\u03B4"},
           11 │[0m   "ml-delta-2": {"type": "string", "value": "\\U000003B4"},
           12 │[0m   "ml-null-1":  {"type": "string", "value": "\\u0000"},
           13 │[0m   "ml-null-2":  {"type": "string", "value": "\\U00000000"},
           14 │[0m   "null-1":     {"type": "string", "value": "\u0000"},
           15 │[0m   "null-2":     {"type": "string", "value": "\u0000"}
           16 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "a":          {"type": "string", "value": "a"},
            3 │[0m     "b":          {"type": "string", "value": "b"},
            4 │[0m     "c":          {"type": "string", "value": "c"},
            5 │[0m     "delta-1":    {"type": "string", "value": "δ"},
            6 │[0m     "delta-2":    {"type": "string", "value": "δ"},
            7 │[0m     "ml-a":       {"type": "string", "value": "a"},
            8 │[0m     "ml-b":       {"type": "string", "value": "b"},
            9 │[0m     "ml-c":       {"type": "string", "value": "c"},
           10 │[0m     "ml-delta-1": {"type": "string", "value": "δ"},
           11 │[0m     "ml-delta-2": {"type": "string", "value": "δ"},
           12 │[0m     "ml-null-1":  {"type": "string", "value": "\u0000"},
           13 │[0m     "ml-null-2":  {"type": "string", "value": "\u0000"},
           14 │[0m     "null-1":     {"type": "string", "value": "\u0000"},
           15 │[0m     "null-2":     {"type": "string", "value": "\u0000"}
           16 │[0m }
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 861
      valid tests:  12 passed, 11 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
  stderr:
    Traceback (most recent call last):
      File "strings-conformance.py", line 12, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: strings-invalid-conformance (FEATURE-Lexical-Strings.md)
  intent: The implementation rejects the authoritative invalid string cases.
  return code: 1
  stdout:
    FAIL invalid/string/bad-escape-01
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 758967):
            1 │[0m invalid-escape = "This string has a bad \a escape character."
    
         output from parser-cmd (PID 758967) (stdout):
            1 │[0m {
            2 │[0m   "invalid-escape": {"type": "string", "value": "This string has a bad \u0007 escape character."}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/string/bad-uni-esc-ml-01
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 758956):
            1 │[0m bad-uni-esc-ml-01 = """val\ue"""
    
         output from parser-cmd (PID 758956) (stdout):
            1 │[0m {
            2 │[0m   "bad-uni-esc-ml-01": {"type": "string", "value": "val\\ue"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/string/bad-uni-esc-ml-02
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 758924):
            1 │[0m bad-uni-esc-ml-02 = """val\Ux"""
    
         output from parser-cmd (PID 758924) (stdout):
            1 │[0m {
            2 │[0m   "bad-uni-esc-ml-02": {"type": "string", "value": "val\\Ux"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/string/bad-uni-esc-ml-03
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 759105):
            1 │[0m bad-uni-esc-ml-03 = """val\U0000000"""
    
         output from parser-cmd (PID 759105) (stdout):
            1 │[0m {
            2 │[0m   "bad-uni-esc-ml-03": {"type": "string", "value": "val\\U0000000"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/string/bad-uni-esc-ml-04
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 758892):
            1 │[0m bad-uni-esc-ml-04 = """val\U0000"""
    
         output from parser-cmd (PID 758892) (stdout):
            1 │[0m {
            2 │[0m   "bad-uni-esc-ml-04": {"type": "string", "value": "val\\U0000"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/string/bad-uni-esc-ml-05
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 758958):
            1 │[0m bad-uni-esc-ml-05 = """val\Ugggggggg"""
    
         output from parser-cmd (PID 758958) (stdout):
            1 │[0m {
            2 │[0m   "bad-uni-esc-ml-05": {"type": "string", "value": "val\\Ugggggggg"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/string/bad-uni-esc-ml-06
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 758957):
            1 │[0m bad-uni-esc-ml-06 = """This string contains a non scalar unicode codepoint \uD801"""
    
         output from parser-cmd (PID 758957) (stdout):
            1 │[0m {
            2 │[0m   "bad-uni-esc-ml-06": {"type": "string", "value": "This string contains a non scalar unicode codepoint \\uD801"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/string/bad-uni-esc-ml-07
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 758893):
            1 │[0m bad-uni-esc-ml-07 = """\uabag"""
    
         output from parser-cmd (PID 758893) (stdout):
            1 │[0m {
            2 │[0m   "bad-uni-esc-ml-07": {"type": "string", "value": "\\uabag"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/string/basic-byte-escapes
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 759046):
            1 │[0m answer = "\x33"
    
         output from parser-cmd (PID 759046) (stdout):
            1 │[0m {
            2 │[0m   "answer": {"type": "string", "value": "3"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/string/basic-multiline-out-of-range-unicode-escape-01
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 758962):
            1 │[0m a = """\UFFFFFFFF"""
    
         output from parser-cmd (PID 758962) (stdout):
            1 │[0m {
            2 │[0m   "a": {"type": "string", "value": "\\UFFFFFFFF"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/string/basic-multiline-out-of-range-unicode-escape-02
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 758959):
            1 │[0m a = """\U00D80000"""
    
         output from parser-cmd (PID 758959) (stdout):
            1 │[0m {
            2 │[0m   "a": {"type": "string", "value": "\\U00D80000"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/string/basic-multiline-unknown-escape
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 758969):
            1 │[0m a = """\@"""
    
         output from parser-cmd (PID 758969) (stdout):
            1 │[0m {
            2 │[0m   "a": {"type": "string", "value": "\\@"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/string/multiline-bad-escape-01
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 759257):
            1 │[0m k = """t\a"""
    
         output from parser-cmd (PID 759257) (stdout):
            1 │[0m {
            2 │[0m   "k": {"type": "string", "value": "t\\a"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/string/multiline-bad-escape-02
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 759090):
            1 │[0m # \<Space> is not a valid escape.
            2 │[0m k = """t\ t"""
    
         output from parser-cmd (PID 759090) (stdout):
            1 │[0m {
            2 │[0m   "k": {"type": "string", "value": "t\\ t"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/string/multiline-bad-escape-03
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 759153):
            1 │[0m # \<Space> is not a valid escape.
            2 │[0m k = """t\ """
    
         output from parser-cmd (PID 759153) (stdout):
            1 │[0m {
            2 │[0m   "k": {"type": "string", "value": "t\\ "}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/string/multiline-escape-space-01
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 759106):
            1 │[0m a = """
            2 │[0m   foo \ \n
            3 │[0m   bar"""
    
         output from parser-cmd (PID 759106) (stdout):
            1 │[0m {
            2 │[0m   "a": {"type": "string", "value": "  foo \\ \\n\n  bar"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/string/multiline-escape-space-02
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 759107):
            1 │[0m bee = """
            2 │[0m hee \
            3 │[0m
            4 │[0m gee \   """
    
         output from parser-cmd (PID 759107) (stdout):
            1 │[0m {
            2 │[0m   "bee": {"type": "string", "value": "hee \ngee \\   "}
            3 │[0m }
    
         want:
           Exit code 1
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 807
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  60 passed, 17 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
  stderr:
    Traceback (most recent call last):
      File "strings-invalid-conformance.py", line 12, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- GREEN (prepassed): strings-process (FEATURE-Lexical-Strings.md)
  intent: A basic string is accepted through the decoder process boundary.
  return code: 0
- GREEN (prepassed): strings-literal (FEATURE-Lexical-Strings.md)
  intent: A literal string preserves its backslashes as supplied input.
  return code: 0

## Post-build programmatic acceptance
- PASS: strings-conformance (FEATURE-Lexical-Strings.md)
  intent: The implementation passes the authoritative TOML 1.0.0 string conformance slice.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 861
      valid tests:  23 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:   0 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: strings-invalid-conformance (FEATURE-Lexical-Strings.md)
  intent: The implementation rejects the authoritative invalid string cases.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/build/toml/toml-decoder] [no encoder]
    skipped tests: 807
      valid tests:   0 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests:  77 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: strings-process (FEATURE-Lexical-Strings.md)
  intent: A basic string is accepted through the decoder process boundary.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
- PASS: strings-literal (FEATURE-Lexical-Strings.md)
  intent: A literal string preserves its backslashes as supplied input.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0

## Build summary
<reusable-compact filename="stage_test.sh">
Runs `go build -o toml-decoder ./cmd/toml-decoder`, then executes `sources/run_conformance.sh -run "$1"` against TOML_TEST_VERSION v2.2.0. Requires exactly one test-pattern argument and uses the command exit status as the verdict.
</reusable-compact>

RESULT: SUCCESS

FILES CHANGED:
- internal/toml/decode.go
- internal/toml/value_test.go

SUMMARY:
Implemented TOML basic, multiline basic, literal, and multiline literal strings. Valid slice: 23/23; invalid slice: 77/77. Go tests, race tests, vet, build, and formatting pass.

BLOCKERS:
- None
