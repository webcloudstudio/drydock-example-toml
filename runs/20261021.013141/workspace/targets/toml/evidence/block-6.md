# Evidence: Block 6 · Service (block-6)

- block type: block
- date: 2026-10-20
- resulting state: closed/verified
- story points (combined assembled cost): 5305
- execution id: 20261021.014805.430Z-3ffad54a

## Stories built
- Verify the completed decoder against the complete supplied TOML suite. (conformance) [story]

## Stacked context
- stack: go_compact.md (SP 1173)
- compass: COMPASS.md (SP 521)
- implements: FEATURE-Conformance.md (SP 469)
- context: full_test.sh (SP 113)
- context: run_conformance.sh (SP 677)
- context: setup_harness.sh (SP 323)
- context: ARCHITECTURE_compact.md (SP 136)
- stack: common.md (SP 1812)

## Build directory changes
- internal/toml/toml.go
- internal/toml/toml_test.go
- toml-decoder

## Pre-build acceptance observation
- RED: complete-conformance (FEATURE-Conformance.md)
  intent: Run the complete pinned TOML v1.0.0 conformance suite after all parser capabilities are implemented.
  return code: 1
  stdout:
    FAIL valid/integer/float64-max
         Values for key "max_int" don't match:
           Expected:     9007199254740991
           Your encoder: 9_007_199_254_740_991
    
         input sent to parser-cmd (PID 126809):
            1 │[0m # Maximum and minimum safe float64 natural numbers. Mainly here for
            2 │[0m # -int-as-float.
            3 │[0m max_int =  9_007_199_254_740_991
            4 │[0m min_int = -9_007_199_254_740_991
    
         output from parser-cmd (PID 126809) (stdout):
            1 │[0m {
            2 │[0m   "max_int": {"type": "integer", "value": "9_007_199_254_740_991"},
            3 │[0m   "min_int": {"type": "integer", "value": "-9_007_199_254_740_991"}
            4 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "max_int": {"type": "integer", "value": "9007199254740991"},
            3 │[0m     "min_int": {"type": "integer", "value": "-9007199254740991"}
            4 │[0m }
    
    FAIL valid/integer/integer
         Values for key "posanswer" don't match:
           Expected:     42
           Your encoder: +42
    
         input sent to parser-cmd (PID 126714):
            1 │[0m answer = 42
            2 │[0m posanswer = +42
            3 │[0m neganswer = -42
            4 │[0m zero = 0
    
         output from parser-cmd (PID 126714) (stdout):
            1 │[0m {
            2 │[0m   "answer":    {"type": "integer", "value": "42"},
            3 │[0m   "neganswer": {"type": "integer", "value": "-42"},
            4 │[0m   "posanswer": {"type": "integer", "value": "+42"},
            5 │[0m   "zero":      {"type": "integer", "value": "0"}
            6 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "answer":    {"type": "integer", "value": "42"},
            3 │[0m     "neganswer": {"type": "integer", "value": "-42"},
            4 │[0m     "posanswer": {"type": "integer", "value": "42"},
            5 │[0m     "zero":      {"type": "integer", "value": "0"}
            6 │[0m }
    
    FAIL valid/integer/literals
         Values for key "bin1" don't match:
           Expected:     214
           Your encoder: 0b11010110
    
         input sent to parser-cmd (PID 126801):
            1 │[0m bin1 = 0b11010110
            2 │[0m bin2 = 0b1_0_1
            3 │[0m
            4 │[0m oct1 = 0o01234567
            5 │[0m oct2 = 0o755
            6 │[0m oct3 = 0o7_6_5
            7 │[0m
            8 │[0m hex1 = 0xDEADBEEF
            9 │[0m hex2 = 0xdeadbeef
           10 │[0m hex3 = 0xdead_beef
           11 │[0m hex4 = 0x00987
    
         output from parser-cmd (PID 126801) (stdout):
            1 │[0m {
            2 │[0m   "bin1": {"type": "integer", "value": "0b11010110"},
            3 │[0m   "bin2": {"type": "integer", "value": "0b1_0_1"},
            4 │[0m   "hex1": {"type": "integer", "value": "0xDEADBEEF"},
            5 │[0m   "hex2": {"type": "integer", "value": "0xdeadbeef"},
            6 │[0m   "hex3": {"type": "integer", "value": "0xdead_beef"},
            7 │[0m   "hex4": {"type": "integer", "value": "0x00987"},
            8 │[0m   "oct1": {"type": "integer", "value": "0o01234567"},
            9 │[0m   "oct2": {"type": "integer", "value": "0o755"},
           10 │[0m   "oct3": {"type": "integer", "value": "0o7_6_5"}
           11 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "bin1": {"type": "integer", "value": "214"},
            3 │[0m     "bin2": {"type": "integer", "value": "5"},
            4 │[0m     "hex1": {"type": "integer", "value": "3735928559"},
            5 │[0m     "hex2": {"type": "integer", "value": "3735928559"},
            6 │[0m     "hex3": {"type": "integer", "value": "3735928559"},
            7 │[0m     "hex4": {"type": "integer", "value": "2439"},
            8 │[0m     "oct1": {"type": "integer", "value": "342391"},
            9 │[0m     "oct2": {"type": "integer", "value": "493"},
           10 │[0m     "oct3": {"type": "integer", "value": "501"}
           11 │[0m }
    
    FAIL valid/integer/underscore
         Values for key "kilo" don't match:
           Expected:     1000
           Your encoder: 1_000
    
         input sent to parser-cmd (PID 126740):
            1 │[0m kilo = 1_000
            2 │[0m x = 1_1_1_1
    
         output from parser-cmd (PID 126740) (stdout):
            1 │[0m {
            2 │[0m   "kilo": {"type": "integer", "value": "1_000"},
            3 │[0m   "x":    {"type": "integer", "value": "1_1_1_1"}
            4 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "kilo": {"type": "integer", "value": "1000"},
            3 │[0m     "x":    {"type": "integer", "value": "1111"}
            4 │[0m }
    
    FAIL valid/integer/zero
         Values for key "a2" don't match:
           Expected:     0
           Your encoder: 0o00
    
         input sent to parser-cmd (PID 126741):
            1 │[0m d1 = 0
            2 │[0m d2 = +0
            3 │[0m d3 = -0
            4 │[0m
            5 │[0m h1 = 0x0
            6 │[0m h2 = 0x00
            7 │[0m h3 = 0x00000
            8 │[0m
            9 │[0m o1 = 0o0
           10 │[0m a2 = 0o00
           11 │[0m a3 = 0o00000
           12 │[0m
           13 │[0m b1 = 0b0
           14 │[0m b2 = 0b00
           15 │[0m b3 = 0b00000
    
         output from parser-cmd (PID 126741) (stdout):
            1 │[0m {
            2 │[0m   "a2": {"type": "integer", "value": "0o00"},
            3 │[0m   "a3": {"type": "integer", "value": "0o00000"},
            4 │[0m   "b1": {"type": "integer", "value": "0b0"},
            5 │[0m   "b2": {"type": "integer", "value": "0b00"},
            6 │[0m   "b3": {"type": "integer", "value": "0b00000"},
            7 │[0m   "d1": {"type": "integer", "value": "0"},
            8 │[0m   "d2": {"type": "integer", "value": "+0"},
            9 │[0m   "d3": {"type": "integer", "value": "-0"},
           10 │[0m   "h1": {"type": "integer", "value": "0x0"},
           11 │[0m   "h2": {"type": "integer", "value": "0x00"},
           12 │[0m   "h3": {"type": "integer", "value": "0x00000"},
           13 │[0m   "o1": {"type": "integer", "value": "0o0"}
           14 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "a2": {"type": "integer", "value": "0"},
            3 │[0m     "a3": {"type": "integer", "value": "0"},
            4 │[0m     "b1": {"type": "integer", "value": "0"},
            5 │[0m     "b2": {"type": "integer", "value": "0"},
            6 │[0m     "b3": {"type": "integer", "value": "0"},
            7 │[0m     "d1": {"type": "integer", "value": "0"},
            8 │[0m     "d2": {"type": "integer", "value": "0"},
            9 │[0m     "d3": {"type": "integer", "value": "0"},
           10 │[0m     "h1": {"type": "integer", "value": "0"},
           11 │[0m     "h2": {"type": "integer", "value": "0"},
           12 │[0m     "h3": {"type": "integer", "value": "0"},
           13 │[0m     "o1": {"type": "integer", "value": "0"}
           14 │[0m }
    
    FAIL valid/spec-1.0.0/integer-0
         Values for key "int1" don't match:
           Expected:     99
           Your encoder: +99
    
         input sent to parser-cmd (PID 126988):
            1 │[0m int1 = +99
            2 │[0m int2 = 42
            3 │[0m int3 = 0
            4 │[0m int4 = -17
    
         output from parser-cmd (PID 126988) (stdout):
            1 │[0m {
            2 │[0m   "int1": {"type": "integer", "value": "+99"},
            3 │[0m   "int2": {"type": "integer", "value": "42"},
            4 │[0m   "int3": {"type": "integer", "value": "0"},
            5 │[0m   "int4": {"type": "integer", "value": "-17"}
            6 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "int1": {"type": "integer", "value": "99"},
            3 │[0m     "int2": {"type": "integer", "value": "42"},
            4 │[0m     "int3": {"type": "integer", "value": "0"},
            5 │[0m     "int4": {"type": "integer", "value": "-17"}
            6 │[0m }
    
    FAIL valid/spec-1.0.0/integer-1
         Values for key "int5" don't match:
           Expected:     1000
           Your encoder: 1_000
    
         input sent to parser-cmd (PID 126989):
            1 │[0m int5 = 1_000
            2 │[0m int6 = 5_349_221
            3 │[0m int7 = 53_49_221  # Indian number system grouping
            4 │[0m int8 = 1_2_3_4_5  # VALID but discouraged
    
         output from parser-cmd (PID 126989) (stdout):
            1 │[0m {
            2 │[0m   "int5": {"type": "integer", "value": "1_000"},
            3 │[0m   "int6": {"type": "integer", "value": "5_349_221"},
            4 │[0m   "int7": {"type": "integer", "value": "53_49_221"},
            5 │[0m   "int8": {"type": "integer", "value": "1_2_3_4_5"}
            6 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "int5": {"type": "integer", "value": "1000"},
            3 │[0m     "int6": {"type": "integer", "value": "5349221"},
            4 │[0m     "int7": {"type": "integer", "value": "5349221"},
            5 │[0m     "int8": {"type": "integer", "value": "12345"}
            6 │[0m }
    
    FAIL valid/spec-1.0.0/integer-2
         Values for key "bin1" don't match:
           Expected:     214
           Your encoder: 0b11010110
    
         input sent to parser-cmd (PID 126990):
            1 │[0m # hexadecimal with prefix `0x`
            2 │[0m hex1 = 0xDEADBEEF
            3 │[0m hex2 = 0xdeadbeef
            4 │[0m hex3 = 0xdead_beef
            5 │[0m
            6 │[0m # octal with prefix `0o`
            7 │[0m oct1 = 0o01234567
            8 │[0m oct2 = 0o755 # useful for Unix file permissions
            9 │[0m
           10 │[0m # binary with prefix `0b`
           11 │[0m bin1 = 0b11010110
    
         output from parser-cmd (PID 126990) (stdout):
            1 │[0m {
            2 │[0m   "bin1": {"type": "integer", "value": "0b11010110"},
            3 │[0m   "hex1": {"type": "integer", "value": "0xDEADBEEF"},
            4 │[0m   "hex2": {"type": "integer", "value": "0xdeadbeef"},
            5 │[0m   "hex3": {"type": "integer", "value": "0xdead_beef"},
            6 │[0m   "oct1": {"type": "integer", "value": "0o01234567"},
            7 │[0m   "oct2": {"type": "integer", "value": "0o755"}
            8 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "bin1": {"type": "integer", "value": "214"},
            3 │[0m     "hex1": {"type": "integer", "value": "3735928559"},
            4 │[0m     "hex2": {"type": "integer", "value": "3735928559"},
            5 │[0m     "hex3": {"type": "integer", "value": "3735928559"},
            6 │[0m     "oct1": {"type": "integer", "value": "342391"},
            7 │[0m     "oct2": {"type": "integer", "value": "493"}
            8 │[0m }
    
    FAIL valid/spec-1.0.0/string-3
         Values for key "str2" don't match:
           Expected:     The quick brown fox jumps over the lazy dog.
           Your encoder: The quick brown
     
           fox jumps over the lazy dog.
    
         input sent to parser-cmd (PID 127051):
            1 │[0m # The following strings are byte-for-byte equivalent:
            2 │[0m str1 = "The quick brown fox jumps over the lazy dog."
            3 │[0m
            4 │[0m str2 = """
            5 │[0m The quick brown \
            6 │[0m
            7 │[0m
            8 │[0m   fox jumps over \
            9 │[0m     the lazy dog."""
           10 │[0m
           11 │[0m str3 = """\
           12 │[0m        The quick brown \
           13 │[0m        fox jumps over \
           14 │[0m        the lazy dog.\
           15 │[0m        """
    
         output from parser-cmd (PID 127051) (stdout):
            1 │[0m {
            2 │[0m   "str1": {"type": "string", "value": "The quick brown fox jumps over the lazy dog."},
            3 │[0m   "str2": {"type": "string", "value": "The quick brown \n\n  fox jumps over the lazy dog."},
            4 │[0m   "str3": {"type": "string", "value": "The quick brown fox jumps over the lazy dog."}
            5 │[0m }
    
         want:
            1 │[0m {
            2 │[0m     "str1": {"type": "string", "value": "The quick brown fox jumps over the lazy dog."},
            3 │[0m     "str2": {"type": "string", "value": "The quick brown fox jumps over the lazy dog."},
            4 │[0m     "str3": {"type": "string", "value": "The quick brown fox jumps over the lazy dog."}
            5 │[0m }
    
    FAIL valid/spec-1.0.0/string-4
         decode TOML: line 7: trailing characters
     
         Exit 1
    
         input sent to parser-cmd (PID 127052):
            1 │[0m str4 = """Here are two quotation marks: "". Simple enough."""
            2 │[0m # str5 = """Here are three quotation marks: """."""  # INVALID
            3 │[0m str5 = """Here are three quotation marks: ""\"."""
            4 │[0m str6 = """Here are fifteen quotation marks: ""\"""\"""\"""\"""\"."""
            5 │[0m
            6 │[0m # "This," she said, "is just a pointless statement."
            7 │[0m str7 = """"This," she said, "is just a pointless statement.""""
    
         output from parser-cmd (PID 127052) (stderr):
            1 │[0m decode TOML: line 7: trailing characters
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/spec-1.0.0/string-6
         decode TOML: line 1: trailing characters
     
         Exit 1
    
         input sent to parser-cmd (PID 127138):
            1 │[0m regex2 = '''I [dw]on't need \d{2} apples'''
            2 │[0m lines  = '''
            3 │[0m The first newline is
            4 │[0m trimmed in raw strings.
            5 │[0m    All other whitespace
            6 │[0m    is preserved.
            7 │[0m '''
    
         output from parser-cmd (PID 127138) (stderr):
            1 │[0m decode TOML: line 1: trailing characters
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/spec-1.0.0/string-7
         decode TOML: line 7: trailing characters
     
         Exit 1
    
         input sent to parser-cmd (PID 127086):
            1 │[0m quot15 = '''Here are fifteen quotation marks: """""""""""""""'''
            2 │[0m
            3 │[0m # apos15 = '''Here are fifteen apostrophes: ''''''''''''''''''  # INVALID
            4 │[0m apos15 = "Here are fifteen apostrophes: '''''''''''''''"
            5 │[0m
            6 │[0m # 'That,' she said, 'is still pointless.'
            7 │[0m str = ''''That,' she said, 'is still pointless.''''
    
         output from parser-cmd (PID 127086) (stderr):
            1 │[0m decode TOML: line 7: trailing characters
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/ends-in-whitespace-escape
         decode TOML: line 1: invalid escape
     
         Exit 1
    
         input sent to parser-cmd (PID 127286):
            1 │[0m beee = """
            2 │[0m heeee
            3 │[0m geeee\
            4 │[0m
            5 │[0m
            6 │[0m       """
    
         output from parser-cmd (PID 127286) (stderr):
            1 │[0m decode TOML: line 1: invalid escape
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/escape-tricky
         decode TOML: line 10: unterminated value
     
         Exit 1
    
         input sent to parser-cmd (PID 127241):
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
    
         output from parser-cmd (PID 127241) (stderr):
            1 │[0m decode TOML: line 10: unterminated value
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/multiline
         decode TOML: line 17: invalid escape
     
         Exit 1
    
         input sent to parser-cmd (PID 127184):
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
    
         output from parser-cmd (PID 127184) (stderr):
            1 │[0m decode TOML: line 17: invalid escape
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/multiline-empty
         decode TOML: line 11: invalid escape
     
         Exit 1
    
         input sent to parser-cmd (PID 127163):
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
    
         output from parser-cmd (PID 127163) (stderr):
            1 │[0m decode TOML: line 11: invalid escape
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/multiline-quotes
         decode TOML: line 4: trailing characters
     
         Exit 1
    
         input sent to parser-cmd (PID 127183):
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
    
         output from parser-cmd (PID 127183) (stderr):
            1 │[0m decode TOML: line 4: trailing characters
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL valid/string/raw-multiline
         decode TOML: line 2: trailing characters
     
         Exit 1
    
         input sent to parser-cmd (PID 127231):
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
    
         output from parser-cmd (PID 127231) (stderr):
            1 │[0m decode TOML: line 2: trailing characters
            2 │[0m
            3 │[0m Exit 1
    
         want:
              <empty>
    
    FAIL invalid/control/bare-cr
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 127598):
            1 │[0m # The following line contains a single carriage return control character
            2 │[0m 
    
         output from parser-cmd (PID 127598) (stdout):
            1 │[0m {}
    
         want:
           Exit code 1
    
    FAIL invalid/control/comment-cr
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 127596):
            1 │[0m comment-cr   = "Carriage return in comment" # 
    a=1
    
         output from parser-cmd (PID 127596) (stdout):
            1 │[0m {
            2 │[0m   "comment-cr": {"type": "string", "value": "Carriage return in comment"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/control/rawstring-cr
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 127746):
            1 │[0m rawstring-cr   = 'null
    '
    
         output from parser-cmd (PID 127746) (stdout):
            1 │[0m {
            2 │[0m   "rawstring-cr": {"type": "string", "value": "null\r"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/control/string-cr
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 127760):
            1 │[0m string-cr   = "null
    "
    
         output from parser-cmd (PID 127760) (stdout):
            1 │[0m {
            2 │[0m   "string-cr": {"type": "string", "value": "null\r"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/encoding/ideographic-space
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 127986):
            1 │[0m # First on next line is U+3000 IDEOGRAPHIC SPACE
            2 │[0m 　foo = "bar"
    
         output from parser-cmd (PID 127986) (stdout):
            1 │[0m {
            2 │[0m   "foo": {"type": "string", "value": "bar"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/exp-dot-02
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 128203):
            1 │[0m exp-dot-02 = 1.e2
    
         output from parser-cmd (PID 128203) (stdout):
            1 │[0m {
            2 │[0m   "exp-dot-02": {"type": "float", "value": "1.e2"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/exp-dot-03
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 128055):
            1 │[0m exp-dot-03 = 3.e+20
    
         output from parser-cmd (PID 128055) (stdout):
            1 │[0m {
            2 │[0m   "exp-dot-03": {"type": "float", "value": "3.e+20"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/leading-dot-neg
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 128106):
            1 │[0m leading-dot-neg = -.12345
    
         output from parser-cmd (PID 128106) (stdout):
            1 │[0m {
            2 │[0m   "leading-dot-neg": {"type": "float", "value": "-.12345"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/leading-dot-plus
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 128107):
            1 │[0m leading-dot-plus = +.12345
    
         output from parser-cmd (PID 128107) (stdout):
            1 │[0m {
            2 │[0m   "leading-dot-plus": {"type": "float", "value": "+.12345"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/leading-zero
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 128144):
            1 │[0m leading-zero = 03.14
    
         output from parser-cmd (PID 128144) (stdout):
            1 │[0m {
            2 │[0m   "leading-zero": {"type": "float", "value": "03.14"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/leading-zero-neg
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 128197):
            1 │[0m leading-zero-neg = -03.14
    
         output from parser-cmd (PID 128197) (stdout):
            1 │[0m {
            2 │[0m   "leading-zero-neg": {"type": "float", "value": "-03.14"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/leading-zero-plus
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 128143):
            1 │[0m leading-zero-plus = +03.14
    
         output from parser-cmd (PID 128143) (stdout):
            1 │[0m {
            2 │[0m   "leading-zero-plus": {"type": "float", "value": "+03.14"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/integer/leading-zero-01
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 128471):
            1 │[0m leading-zero-01 = 01
    
         output from parser-cmd (PID 128471) (stdout):
            1 │[0m {
            2 │[0m   "leading-zero-01": {"type": "integer", "value": "01"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/integer/leading-zero-02
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 128470):
            1 │[0m leading-zero-02 = 00
    
         output from parser-cmd (PID 128470) (stdout):
            1 │[0m {
            2 │[0m   "leading-zero-02": {"type": "integer", "value": "00"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/integer/leading-zero-03
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 128472):
            1 │[0m leading-zero-03 = 0_0
    
         output from parser-cmd (PID 128472) (stdout):
            1 │[0m {
            2 │[0m   "leading-zero-03": {"type": "integer", "value": "0_0"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/integer/leading-zero-sign-01
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 128489):
            1 │[0m leading-zero-sign-01 = -01
    
         output from parser-cmd (PID 128489) (stdout):
            1 │[0m {
            2 │[0m   "leading-zero-sign-01": {"type": "integer", "value": "-01"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/integer/leading-zero-sign-02
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 128551):
            1 │[0m leading-zero-sign-02 = +01
    
         output from parser-cmd (PID 128551) (stdout):
            1 │[0m {
            2 │[0m   "leading-zero-sign-02": {"type": "integer", "value": "+01"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/integer/leading-zero-sign-03
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 128493):
            1 │[0m leading-zero-sign-03 = +0_1
    
         output from parser-cmd (PID 128493) (stdout):
            1 │[0m {
            2 │[0m   "leading-zero-sign-03": {"type": "integer", "value": "+0_1"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/integer/negative-bin
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 128599):
            1 │[0m negative-bin = -0b11010110
    
         output from parser-cmd (PID 128599) (stdout):
            1 │[0m {
            2 │[0m   "negative-bin": {"type": "integer", "value": "-0b11010110"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/integer/negative-hex
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 128501):
            1 │[0m negative-hex = -0xff
    
         output from parser-cmd (PID 128501) (stdout):
            1 │[0m {
            2 │[0m   "negative-hex": {"type": "integer", "value": "-0xff"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/integer/negative-oct
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 128504):
            1 │[0m negative-oct = -0o755
    
         output from parser-cmd (PID 128504) (stdout):
            1 │[0m {
            2 │[0m   "negative-oct": {"type": "integer", "value": "-0o755"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/integer/positive-bin
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 128505):
            1 │[0m positive-bin = +0b11010110
    
         output from parser-cmd (PID 128505) (stdout):
            1 │[0m {
            2 │[0m   "positive-bin": {"type": "integer", "value": "+0b11010110"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/integer/positive-hex
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 128509):
            1 │[0m positive-hex = +0xff
    
         output from parser-cmd (PID 128509) (stdout):
            1 │[0m {
            2 │[0m   "positive-hex": {"type": "integer", "value": "+0xff"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/integer/positive-oct
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 128559):
            1 │[0m positive-oct = +0o755
    
         output from parser-cmd (PID 128559) (stdout):
            1 │[0m {
            2 │[0m   "positive-oct": {"type": "integer", "value": "+0o755"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/table/append-with-dotted-keys-01
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 129525):
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
    
         output from parser-cmd (PID 129525) (stdout):
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
    
         input sent to parser-cmd (PID 129526):
            1 │[0m # This is the same issue as in injection-1.toml, except that nests one level
            2 │[0m # deeper. See that file for a more complete description.
            3 │[0m
            4 │[0m [a.b.c.d]
            5 │[0m   z = 9
            6 │[0m
            7 │[0m [a]
            8 │[0m   b.c.d.k.t = "Using dotted keys to add to [a.b.c.d] after explicitly defining it above is not allowed"
    
         output from parser-cmd (PID 129526) (stdout):
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
    
         input sent to parser-cmd (PID 129548):
            1 │[0m [a.b.c]
            2 │[0m z = 9
            3 │[0m
            4 │[0m [[totally_unrelated]]
            5 │[0m x = 123
            6 │[0m
            7 │[0m [a]
            8 │[0m b.c.t = "this should NOT be allowed"
    
         output from parser-cmd (PID 129548) (stdout):
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
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20261021.013141/build/toml/toml-decoder] [no encoder]
      valid tests: 187 passed, 18 failed
    encoder tests: no encoder command given
    invalid tests: 447 passed, 27 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['sh', 'sources/full_test.sh'], returncode=1, stdout='FAIL valid/integer/float64-max\n     Values for key "max_int" don\'t match:\n       Expected:     9007199254740991\n       Your encoder: 9_007_199_254_740_991\n\n     input sent to parser-cmd (PID 126809):\n        1 │\x1b[0m # Maximum and minimum safe float64 natural numbers. Mainly here for\n        2 │\x1b[0m # -int-as-… (+35072 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "complete-conformance.py", line 12, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- PASS: complete-conformance (FEATURE-Conformance.md)
  intent: Run the complete pinned TOML v1.0.0 conformance suite after all parser capabilities are implemented.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20261021.013141/build/toml/toml-decoder] [no encoder]
      valid tests: 205 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests: 474 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true

## Build summary
RESULT: SUCCESS

FILES CHANGED:
- internal/toml/toml.go
- internal/toml/toml_test.go
- toml-decoder

SUMMARY:
Implemented conformance fixes. Go tests pass, and toml-test v2.2.0 passes all 205 valid and 474 invalid cases.

BLOCKERS:
- None
