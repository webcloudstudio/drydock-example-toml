package toml

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestDecodeCommentRunsToEndOfLine(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "trailing comment with quotes and equals", input: "s = \"v\" # a = \"b\"\n", want: "v"},
		{name: "trailing comment with non-ASCII", input: "s = \"v\" # café 😀 中文\n", want: "v"},
		{name: "trailing comment with tab", input: "s = \"v\" #\tx\ty\n", want: "v"},
		{name: "trailing comment with CRLF", input: "s = \"v\" # c\r\n", want: "v"},
		{name: "trailing comment at end of input", input: "s = \"v\" # end", want: "v"},
		{name: "empty trailing comment", input: "s = \"v\" #\n", want: "v"},
		{name: "comment line before assignment", input: "# head\ns = \"v\"\n", want: "v"},
		{name: "comment line after assignment", input: "s = \"v\"\n# tail", want: "v"},
		{name: "hash inside basic string is content", input: "s = \"a#b\" # c\n", want: "a#b"},
		{name: "hash inside literal string is content", input: "s = 'a#b' # c\n", want: "a#b"},
		{name: "hash inside multi-line string is content", input: "s = \"\"\"a # b\"\"\" # c\n", want: "a # b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := decodeString(tt.input)
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

func TestDecodeCommentOnlyDocumentHasNoKeys(t *testing.T) {
	for _, input := range []string{"# only\n", "# only", "#\n#\n", "\t# tabbed\r\n"} {
		t.Run(fmt.Sprintf("%q", input), func(t *testing.T) {
			doc, err := decodeString(input)
			if err != nil {
				t.Fatalf("Decode(%q) error = %v, want nil", input, err)
			}
			if len(doc) != 0 {
				t.Errorf("Decode(%q) = %v, want no keys", input, doc)
			}
		})
	}
}

func TestDecodeAllowsTabInCommentAndLiteralStrings(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "tab in literal string", input: "s = 'a\tb'\n", want: "a\tb"},
		{name: "tab in multi-line literal string", input: "s = '''a\tb'''\n", want: "a\tb"},
		{name: "tab in multi-line basic string", input: "s = \"\"\"a\tb\"\"\"\n", want: "a\tb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := decodeString(tt.input)
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

// disallowedControls lists every control byte TOML forbids inside comments and strings.
// Tab (U+0009) is allowed and line feed (U+000A) ends a line, so neither appears here.
var disallowedControls = func() []byte {
	var cs []byte
	for c := byte(0x00); c < 0x20; c++ {
		if c != '\t' && c != '\n' {
			cs = append(cs, c)
		}
	}
	return append(cs, 0x7F)
}()

func TestDecodeRejectsControlCharactersInCommentsAndLiterals(t *testing.T) {
	contexts := []struct {
		name   string
		format string
	}{
		{name: "comment line", format: "# a%cb\ns = \"v\"\n"},
		{name: "trailing comment", format: "s = \"v\" # a%cb\n"},
		{name: "literal string", format: "s = 'a%cb'\n"},
		{name: "multi-line literal string", format: "s = '''a%cb'''\n"},
	}
	for _, ctx := range contexts {
		for _, c := range disallowedControls {
			t.Run(fmt.Sprintf("%s U+%04X", ctx.name, c), func(t *testing.T) {
				input := fmt.Sprintf(ctx.format, c)
				doc, err := decodeString(input)
				if !errors.Is(err, ErrSyntax) {
					t.Fatalf("Decode(%q) error = %v, want %v", input, err, ErrSyntax)
				}
				if doc != nil {
					t.Errorf("Decode(%q) = %v, want nil document", input, doc)
				}
			})
		}
	}
}

func TestDecodeRejectsBareCarriageReturnInComment(t *testing.T) {
	doc, err := decodeString("s = \"v\" # a\rb\n")
	if !errors.Is(err, ErrSyntax) {
		t.Fatalf("Decode error = %v, want %v", err, ErrSyntax)
	}
	if doc != nil {
		t.Errorf("Decode = %v, want nil document", doc)
	}
}

func TestDecodeRejectsInvalidUTF8(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "lone continuation byte", input: "s = \"v\"\n\x80"},
		{name: "byte 0xFF", input: "\xff = \"v\"\n"},
		{name: "truncated two-byte sequence", input: "s = \"\xc3\"\n"},
		{name: "truncated three-byte sequence", input: "s = \"\xe2\x82\"\n"},
		{name: "overlong slash", input: "s = \"\xc0\xaf\"\n"},
		{name: "encoded surrogate", input: "s = \"\xed\xa0\x80\"\n"},
		{name: "beyond U+10FFFF", input: "s = \"\xf4\x90\x80\x80\"\n"},
		{name: "invalid byte in comment", input: "# \xff\ns = \"v\"\n"},
		{name: "invalid byte in trailing comment", input: "s = \"v\" # \xfe\n"},
		{name: "invalid byte in literal string", input: "s = 'a\xffb'\n"},
		{name: "invalid byte in multi-line string", input: "s = \"\"\"a\xffb\"\"\"\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := decodeString(tt.input)
			if !errors.Is(err, ErrSyntax) {
				t.Fatalf("Decode(%q) error = %v, want %v", tt.input, err, ErrSyntax)
			}
			if doc != nil {
				t.Errorf("Decode(%q) = %v, want nil document; invalid UTF-8 must never decode partially", tt.input, doc)
			}
		})
	}
}

func TestDecodeAcceptsValidUTF8InCommentsAndStrings(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "multi-byte comment", input: "# ünïcödé 😀 中文 ✓\ns = \"ok\"\n", want: "ok"},
		{name: "last scalar value in comment", input: "# \U0010FFFF\ns = \"ok\"\n", want: "ok"},
		{name: "multi-byte literal string", input: "s = 'é€😀'\n", want: "é€😀"},
		{name: "multi-byte trailing comment", input: "s = \"ok\" # ✓\n", want: "ok"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := decodeString(tt.input)
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
