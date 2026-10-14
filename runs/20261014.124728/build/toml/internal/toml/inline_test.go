package toml

import (
	"errors"
	"reflect"
	"testing"
)

func TestDecodeBuildsInlineTables(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[string]any
	}{
		{
			name:  "empty inline table",
			input: "t = {}\n",
			want:  map[string]any{"t": map[string]any{}},
		},
		{
			name:  "one pair",
			input: "point = { x = 1 }\n",
			want:  map[string]any{"point": map[string]any{"x": intLeaf("1")}},
		},
		{
			name:  "two pairs of different types",
			input: "t = { a = 1, b = \"v\" }\n",
			want: map[string]any{"t": map[string]any{
				"a": intLeaf("1"),
				"b": stringLeaf("v"),
			}},
		},
		{
			name:  "nested inline table",
			input: "t = { a = { b = 1 } }\n",
			want:  map[string]any{"t": map[string]any{"a": map[string]any{"b": intLeaf("1")}}},
		},
		{
			name:  "dotted keys inside the braces share one table",
			input: "t = { a.b = 1, a.c = 2 }\n",
			want: map[string]any{"t": map[string]any{
				"a": map[string]any{"b": intLeaf("1"), "c": intLeaf("2")},
			}},
		},
		{
			name:  "array value inside an inline table",
			input: "t = { a = [1, 2] }\n",
			want:  map[string]any{"t": map[string]any{"a": arrayLeaf(intLeaf("1"), intLeaf("2"))}},
		},
		{
			name:  "multi-line array inside an inline table",
			input: "t = { a = [\n  1,\n  2,\n] }\n",
			want:  map[string]any{"t": map[string]any{"a": arrayLeaf(intLeaf("1"), intLeaf("2"))}},
		},
		{
			name:  "inline table beside a header that does not extend it",
			input: "[a]\nb = { c = 1 }\n[a.d]\ne = 2\n",
			want: map[string]any{"a": map[string]any{
				"b": map[string]any{"c": intLeaf("1")},
				"d": map[string]any{"e": intLeaf("2")},
			}},
		},
		{
			name:  "same inline table name in each array of tables element",
			input: "[[arr]]\nt = { a = 1 }\n[[arr]]\nt = { b = 2 }\n",
			want: map[string]any{"arr": []any{
				map[string]any{"t": map[string]any{"a": intLeaf("1")}},
				map[string]any{"t": map[string]any{"b": intLeaf("2")}},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := decodeString(tt.input)
			if err != nil {
				t.Fatalf("Decode(%q) error = %v", tt.input, err)
			}
			if !reflect.DeepEqual(doc, tt.want) {
				t.Errorf("Decode(%q) = %v, want %v", tt.input, doc, tt.want)
			}
		})
	}
}

func TestDecodeRejectsMalformedInlineTables(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "unterminated", input: "t = { a = 1\n"},
		{name: "trailing comma", input: "t = { a = 1, }\n"},
		{name: "lone comma", input: "t = { , }\n"},
		{name: "missing comma", input: "t = { a = 1 b = 2 }\n"},
		{name: "missing value", input: "t = { a = }\n"},
		{name: "missing equals", input: "t = { a 1 }\n"},
		{name: "duplicate key", input: "t = { a = 1, a = 2 }\n"},
		{name: "line break before first key", input: "t = {\na = 1 }\n"},
		{name: "line break before closing brace", input: "t = { a = 1\n}\n"},
		{name: "line break after comma", input: "t = { a = 1,\nb = 2 }\n"},
		{name: "text after closing brace", input: "t = { a = 1 }}\n"},
		{name: "comment inside braces", input: "t = { a = 1 # c\n}\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := decodeString(tt.input)
			if !errors.Is(err, ErrSyntax) {
				t.Fatalf("Decode(%q) error = %v, want %v", tt.input, err, ErrSyntax)
			}
			if doc != nil {
				t.Errorf("Decode(%q) = %v, want nil document", tt.input, doc)
			}
		})
	}
}

func TestDecodeRejectsExtendingClosedInlineTable(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "header defines the table", input: "t = { a = 1 }\n[t]\nb = 2\n"},
		{name: "header extends the table", input: "t = { a = 1 }\n[t.sub]\n"},
		{name: "array header on the table", input: "t = { a = 1 }\n[[t]]\n"},
		{name: "dotted key adds a key", input: "t = { a = 1 }\nt.b = 2\n"},
		{name: "dotted key extends a nested dotted table", input: "t = { a.b = 1 }\nt.a.c = 2\n"},
		{name: "header walks through a nested inline table", input: "t = { a = { b = 1 } }\n[t.a.c]\n"},
		{name: "dotted key extends an inline table value", input: "a.b = { c = 1 }\na.b.d = 2\n"},
		{name: "header under a parent that holds the inline table", input: "[x]\nt = { a = 1 }\n[x.t.sub]\n"},
		{name: "inline table redefines a key", input: "t = { a = 1 }\nt = { b = 2 }\n"},
		{name: "header defines a key inside an array element table", input: "[[arr]]\nt = { a = 1 }\n[arr.t]\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := decodeString(tt.input)
			if !errors.Is(err, ErrSyntax) {
				t.Fatalf("Decode(%q) error = %v, want %v", tt.input, err, ErrSyntax)
			}
			if doc != nil {
				t.Errorf("Decode(%q) = %v, want nil document", tt.input, doc)
			}
		})
	}
}
