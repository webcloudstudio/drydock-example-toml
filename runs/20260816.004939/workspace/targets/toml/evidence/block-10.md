# Evidence: Block 10 · Service (block-10)

- block type: block
- date: 2026-08-15
- resulting state: closed/verified
- story points (combined assembled cost): 9683
- execution id: 20260816.013330.042Z-abd5c8bf

## Stories built
- Verify complete TOML 1.0.0 conformance against the installed suite. (conformance) [story]

## Stacked context
- compass: COMPASS.md (SP 2899)
- implements: FEATURE-Conformance.md (SP 643)
- context: full_test.sh (SP 113)
- context: run_conformance.sh (SP 677)
- context: setup_harness.sh (SP 323)
- context: ARCHITECTURE_compact.md (SP 128)
- stack: go.md (SP 3646)
- stack: common_compact.md (SP 1179)

## Build directory changes
- internal/toml/toml.go
- internal/toml/toml_test.go
- toml-decoder

## Pre-build acceptance observation
- RED: conformance-full (FEATURE-Conformance.md)
  intent: The completed implementation passes the complete authoritative TOML 1.0.0 suite.
  return code: 1
  stdout:
    FAIL valid/integer/float64-max
         Values for key "max_int" don't match:
           Expected:     9007199254740991
           Your encoder: 9_007_199_254_740_991
    
         input sent to parser-cmd (PID 1519906):
            1 │[0m # Maximum and minimum safe float64 natural numbers. Mainly here for
            2 │[0m # -int-as-float.
            3 │[0m max_int =  9_007_199_254_740_991
            4 │[0m min_int = -9_007_199_254_740_991
    
         output from parser-cmd (PID 1519906) (stdout):
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
    
         input sent to parser-cmd (PID 1519961):
            1 │[0m answer = 42
            2 │[0m posanswer = +42
            3 │[0m neganswer = -42
            4 │[0m zero = 0
    
         output from parser-cmd (PID 1519961) (stdout):
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
    
         input sent to parser-cmd (PID 1519964):
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
    
         output from parser-cmd (PID 1519964) (stdout):
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
    
         input sent to parser-cmd (PID 1520002):
            1 │[0m kilo = 1_000
            2 │[0m x = 1_1_1_1
    
         output from parser-cmd (PID 1520002) (stdout):
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
    
         input sent to parser-cmd (PID 1519908):
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
    
         output from parser-cmd (PID 1519908) (stdout):
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
    
         input sent to parser-cmd (PID 1520173):
            1 │[0m int1 = +99
            2 │[0m int2 = 42
            3 │[0m int3 = 0
            4 │[0m int4 = -17
    
         output from parser-cmd (PID 1520173) (stdout):
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
    
         input sent to parser-cmd (PID 1520200):
            1 │[0m int5 = 1_000
            2 │[0m int6 = 5_349_221
            3 │[0m int7 = 53_49_221  # Indian number system grouping
            4 │[0m int8 = 1_2_3_4_5  # VALID but discouraged
    
         output from parser-cmd (PID 1520200) (stdout):
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
    
         input sent to parser-cmd (PID 1520208):
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
    
         output from parser-cmd (PID 1520208) (stdout):
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
    
    FAIL invalid/float/exp-dot-02
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1521186):
            1 │[0m exp-dot-02 = 1.e2
    
         output from parser-cmd (PID 1521186) (stdout):
            1 │[0m {
            2 │[0m   "exp-dot-02": {"type": "float", "value": "1.e2"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/exp-dot-03
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1521362):
            1 │[0m exp-dot-03 = 3.e+20
    
         output from parser-cmd (PID 1521362) (stdout):
            1 │[0m {
            2 │[0m   "exp-dot-03": {"type": "float", "value": "3.e+20"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/trailing-dot
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1521394):
            1 │[0m trailing-dot = 1.
    
         output from parser-cmd (PID 1521394) (stdout):
            1 │[0m {
            2 │[0m   "trailing-dot": {"type": "float", "value": "1."}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/trailing-dot-01
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1521395):
            1 │[0m trailing-point = 1.
    
         output from parser-cmd (PID 1521395) (stdout):
            1 │[0m {
            2 │[0m   "trailing-point": {"type": "float", "value": "1."}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/trailing-dot-02
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1521400):
            1 │[0m a = 1.
            2 │[0m b = 2
    
         output from parser-cmd (PID 1521400) (stdout):
            1 │[0m {
            2 │[0m   "a": {"type": "float", "value": "1."},
            3 │[0m   "b": {"type": "integer", "value": "2"}
            4 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/trailing-dot-min
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1521399):
            1 │[0m trailing-dot-min = -1.
    
         output from parser-cmd (PID 1521399) (stdout):
            1 │[0m {
            2 │[0m   "trailing-dot-min": {"type": "float", "value": "-1."}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/float/trailing-dot-plus
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1521403):
            1 │[0m trailing-dot-plus = +1.
    
         output from parser-cmd (PID 1521403) (stdout):
            1 │[0m {
            2 │[0m   "trailing-dot-plus": {"type": "float", "value": "+1."}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/integer/negative-bin
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1521735):
            1 │[0m negative-bin = -0b11010110
    
         output from parser-cmd (PID 1521735) (stdout):
            1 │[0m {
            2 │[0m   "negative-bin": {"type": "integer", "value": "-0b11010110"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/integer/negative-hex
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1522000):
            1 │[0m negative-hex = -0xff
    
         output from parser-cmd (PID 1522000) (stdout):
            1 │[0m {
            2 │[0m   "negative-hex": {"type": "integer", "value": "-0xff"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/integer/negative-oct
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1521883):
            1 │[0m negative-oct = -0o755
    
         output from parser-cmd (PID 1521883) (stdout):
            1 │[0m {
            2 │[0m   "negative-oct": {"type": "integer", "value": "-0o755"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/integer/positive-bin
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1521734):
            1 │[0m positive-bin = +0b11010110
    
         output from parser-cmd (PID 1521734) (stdout):
            1 │[0m {
            2 │[0m   "positive-bin": {"type": "integer", "value": "+0b11010110"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/integer/positive-hex
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1522077):
            1 │[0m positive-hex = +0xff
    
         output from parser-cmd (PID 1522077) (stdout):
            1 │[0m {
            2 │[0m   "positive-hex": {"type": "integer", "value": "+0xff"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/integer/positive-oct
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1521777):
            1 │[0m positive-oct = +0o755
    
         output from parser-cmd (PID 1521777) (stdout):
            1 │[0m {
            2 │[0m   "positive-oct": {"type": "integer", "value": "+0o755"}
            3 │[0m }
    
         want:
           Exit code 1
    
    FAIL invalid/spec-1.0.0/table-9-0
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1522307):
            1 │[0m [fruit]
            2 │[0m apple.color = "red"
            3 │[0m apple.taste.sweet = true
            4 │[0m
            5 │[0m [fruit.apple]  # INVALID
            6 │[0m # [fruit.apple.taste]  # INVALID
            7 │[0m
            8 │[0m [fruit.apple.texture]  # you can add sub-tables
            9 │[0m smooth = true
    
         output from parser-cmd (PID 1522307) (stdout):
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
    
         input sent to parser-cmd (PID 1522268):
            1 │[0m [fruit]
            2 │[0m apple.color = "red"
            3 │[0m apple.taste.sweet = true
            4 │[0m
            5 │[0m # [fruit.apple]  # INVALID
            6 │[0m [fruit.apple.taste]  # INVALID
            7 │[0m
            8 │[0m [fruit.apple.texture]  # you can add sub-tables
            9 │[0m smooth = true
    
         output from parser-cmd (PID 1522268) (stdout):
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
    
    FAIL invalid/table/append-with-dotted-keys-01
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1522840):
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
    
         output from parser-cmd (PID 1522840) (stdout):
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
    
         input sent to parser-cmd (PID 1522713):
            1 │[0m # This is the same issue as in injection-1.toml, except that nests one level
            2 │[0m # deeper. See that file for a more complete description.
            3 │[0m
            4 │[0m [a.b.c.d]
            5 │[0m   z = 9
            6 │[0m
            7 │[0m [a]
            8 │[0m   b.c.d.k.t = "Using dotted keys to add to [a.b.c.d] after explicitly defining it above is not allowed"
    
         output from parser-cmd (PID 1522713) (stdout):
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
    
         input sent to parser-cmd (PID 1522669):
            1 │[0m [a.b.c]
            2 │[0m z = 9
            3 │[0m
            4 │[0m [[totally_unrelated]]
            5 │[0m x = 123
            6 │[0m
            7 │[0m [a]
            8 │[0m b.c.t = "this should NOT be allowed"
    
         output from parser-cmd (PID 1522669) (stdout):
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
    
    FAIL invalid/table/duplicate-key-04
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1522794):
            1 │[0m [fruit]
            2 │[0m apple.color = "red"
            3 │[0m
            4 │[0m [fruit.apple] # INVALID
    
         output from parser-cmd (PID 1522794) (stdout):
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
    
         input sent to parser-cmd (PID 1522795):
            1 │[0m [fruit]
            2 │[0m apple.taste.sweet = true
            3 │[0m
            4 │[0m [fruit.apple.taste] # INVALID
    
         output from parser-cmd (PID 1522795) (stdout):
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
    
    FAIL invalid/table/redefine-02
         Expected an error, but no error was reported.
    
         input sent to parser-cmd (PID 1522948):
            1 │[0m [t1]
            2 │[0m t2.t3.v = 0
            3 │[0m [t1.t2]
    
         output from parser-cmd (PID 1522948) (stdout):
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
    
         input sent to parser-cmd (PID 1522981):
            1 │[0m [t1]
            2 │[0m t2.t3.v = 0
            3 │[0m [t1.t2.t3]
    
         output from parser-cmd (PID 1522981) (stdout):
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
    
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.004939/build/toml/toml-decoder] [no encoder]
      valid tests: 197 passed,  8 failed
    encoder tests: no encoder command given
    invalid tests: 452 passed, 22 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
  stderr:
    Traceback (most recent call last):
      File "conformance-full.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- GREEN (prepassed): conformance-harness-integrity (FEATURE-Conformance.md)
  intent: The supplied full-test entry point remains present and executable as a shell script.
  return code: 0
- GREEN (prepassed): conformance-decoder-build (FEATURE-Conformance.md)
  intent: The decoder builds successfully through the same Go command used by the scoring harness.
  return code: 0

## Post-build programmatic acceptance
- PASS: conformance-full (FEATURE-Conformance.md)
  intent: The completed implementation passes the complete authoritative TOML 1.0.0 suite.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.004939/build/toml/toml-decoder] [no encoder]
      valid tests: 205 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests: 474 passed,  0 failed
    
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: conformance-harness-integrity (FEATURE-Conformance.md)
  intent: The supplied full-test entry point remains present and executable as a shell script.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
- PASS: conformance-decoder-build (FEATURE-Conformance.md)
  intent: The decoder builds successfully through the same Go command used by the scoring harness.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0

## Build summary
RESULT: SUCCESS

FILES CHANGED:
- internal/toml/toml.go
- internal/toml/toml_test.go
- toml-decoder

SUMMARY:
- Full TOML Test v2.2.0 suite passed: 205 valid, 474 invalid.
- `go test -race`, `go vet`, and build passed.
- Harness scripts remained unchanged.

BLOCKERS:
- None
