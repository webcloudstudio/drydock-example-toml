package toml

import (
	"errors"
	"reflect"
	"testing"
)

// stringValueDoc is the tagged model a document `s = <value>` must decode to.
func stringValueDoc(value string) map[string]any {
	return map[string]any{"s": map[string]any{"type": "string", "value": value}}
}

func TestDecodeStringFormsYieldExpectedValue(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Basic strings
		{name: "basic simple", input: `"text"`, want: "text"},
		{name: "basic empty", input: `""`, want: ""},
		{name: "basic hash is content", input: `"a # b"`, want: "a # b"},
		{name: "basic tab is allowed", input: "\"a\tb\"", want: "a\tb"},
		{name: "basic non-ascii", input: `"é€😀"`, want: "é€😀"},
		{name: "basic escape b", input: `"\b"`, want: "\b"},
		{name: "basic escape t", input: `"\t"`, want: "\t"},
		{name: "basic escape n", input: `"\n"`, want: "\n"},
		{name: "basic escape f", input: `"\f"`, want: "\f"},
		{name: "basic escape r", input: `"\r"`, want: "\r"},
		{name: "basic escape quote", input: `"\""`, want: `"`},
		{name: "basic escape backslash", input: `"\\"`, want: `\`},
		{name: "basic escaped backslash then quote", input: `"\\\""`, want: `\"`},
		{name: "basic unicode escape 4", input: `"\u00E9"`, want: "é"},
		{name: "basic unicode escape 8", input: `"\U0001F600"`, want: "😀"},
		{name: "basic nul escape", input: `"\u0000"`, want: "\x00"},
		{name: "basic last scalar", input: `"\U0010FFFF"`, want: "\U0010FFFF"},
		{name: "basic escape is not a unicode escape", input: `"\\u0041"`, want: `\u0041`},

		// Multi-line basic strings
		{name: "multiline trims opening newline", input: "\"\"\"\nabc\"\"\"", want: "abc"},
		{name: "multiline trims opening CRLF", input: "\"\"\"\r\nabc\"\"\"", want: "abc"},
		{name: "multiline keeps interior newline", input: "\"\"\"a\nb\"\"\"", want: "a\nb"},
		{name: "multiline empty", input: `""""""`, want: ""},
		{name: "multiline empty after trimmed newline", input: "\"\"\"\n\"\"\"", want: ""},
		{name: "multiline quotes inside", input: `"""a "quoted" b"""`, want: `a "quoted" b`},
		{name: "multiline one extra quote before close", input: `"""a""""`, want: `a"`},
		{name: "multiline two extra quotes before close", input: `"""a"""""`, want: `a""`},
		{name: "multiline escaped quote", input: `"""\""""`, want: `"`},
		{name: "multiline unicode escape", input: `"""\u0041"""`, want: "A"},
		{name: "multiline line-ending backslash joins lines", input: "\"\"\"one \\\n\n  two\"\"\"", want: "one two"},
		{name: "multiline line-ending backslash with trailing spaces", input: "\"\"\"a\\  \n b\"\"\"", want: "ab"},

		// Literal strings
		{name: "literal simple", input: `'C:\path'`, want: `C:\path`},
		{name: "literal empty", input: `''`, want: ""},
		{name: "literal escapes are not processed", input: `'\b\t\n\\\u0041'`, want: `\b\t\n\\\u0041`},
		{name: "literal double quote is content", input: `'a "b"'`, want: `a "b"`},

		// Multi-line literal strings
		{name: "multiline literal trims opening newline", input: "'''\nabc'''", want: "abc"},
		{name: "multiline literal keeps newline and backslash", input: "'''a\nb\\n'''", want: "a\nb\\n"},
		{name: "multiline literal single quote", input: `''''one quote''''`, want: `'one quote'`},
		{name: "multiline literal two extra quotes", input: `'''a'''''`, want: `a''`},
		{name: "multiline literal spaced quotes", input: `''' 'one quote' '''`, want: ` 'one quote' `},
		{name: "multiline literal empty", input: `''''''`, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := decodeString("s = " + tt.input + "\n")
			if err != nil {
				t.Fatalf("Decode(%q) error = %v, want nil", tt.input, err)
			}
			want := stringValueDoc(tt.want)
			if !reflect.DeepEqual(doc, want) {
				t.Errorf("Decode(%q) = %v, want %v", tt.input, doc, want)
			}
		})
	}
}

func TestDecodeRejectsInvalidStringForms(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		// Basic strings
		{name: "basic unknown escape q", input: `"\q"`},
		{name: "basic hex escape", input: `"\x41"`},
		{name: "basic slash escape", input: `"\/"`},
		{name: "basic escape with space", input: `"\ "`},
		{name: "basic unterminated", input: `"abc`},
		{name: "basic raw newline", input: "\"a\nb\""},
		{name: "basic raw carriage return", input: "\"a\rb\""},
		{name: "basic raw DEL", input: "\"a\x7fb\""},
		{name: "basic raw control", input: "\"a\x01b\""},
		{name: "basic short unicode escape", input: `"\u12"`},
		{name: "basic non-hex unicode escape", input: `"\u12G4"`},
		{name: "basic surrogate unicode escape", input: `"\uD800"`},
		{name: "basic beyond unicode range", input: `"\U00110000"`},
		{name: "basic 8-digit overflow", input: `"\UFFFFFFFF"`},

		// Multi-line basic strings
		{name: "multiline unknown escape", input: `"""\q"""`},
		{name: "multiline escape then space", input: `"""\ x"""`},
		{name: "multiline unterminated", input: `"""abc`},
		{name: "multiline too many closing quotes", input: `"""a""""""`},
		{name: "multiline bare carriage return", input: "\"\"\"a\rb\"\"\""},
		{name: "multiline surrogate escape", input: `"""\uDFFF"""`},

		// Literal strings
		{name: "literal unterminated", input: `'abc`},
		{name: "literal raw newline", input: "'a\nb'"},
		{name: "literal raw control", input: "'a\x01b'"},

		// Multi-line literal strings
		{name: "multiline literal unterminated", input: `'''abc`},
		{name: "multiline literal too many closing quotes", input: `'''a''''''`},
		{name: "multiline literal raw control", input: "'''a\x01'''"},

		// Other value types
		{name: "array missing comma", input: `["x" "y"]`},
		{name: "array leading comma", input: `[,"x"]`},
		{name: "inline table trailing comma", input: `{a = "x",}`},
		{name: "inline table missing value", input: `{a = }`},
		{name: "bare word", input: `text`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := decodeString("s = " + tt.input + "\n")
			if !errors.Is(err, ErrSyntax) {
				t.Fatalf("Decode(%q) error = %v, want %v", tt.input, err, ErrSyntax)
			}
			if doc != nil {
				t.Errorf("Decode(%q) = %v, want nil document", tt.input, doc)
			}
		})
	}
}
