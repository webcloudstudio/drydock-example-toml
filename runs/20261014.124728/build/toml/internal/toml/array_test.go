package toml

import (
	"errors"
	"reflect"
	"testing"
)

// arrayLeaf is the JSON model for an array value, shared by the array and inline tests.
// An array is a plain list; only its scalar elements carry the tagged form.
func arrayLeaf(elems ...any) []any {
	if elems == nil {
		return []any{}
	}
	return elems
}

func TestDecodeBuildsArrays(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[string]any
	}{
		{
			name:  "empty array",
			input: "a = []\n",
			want:  map[string]any{"a": arrayLeaf()},
		},
		{
			name:  "one element",
			input: "a = [1]\n",
			want:  map[string]any{"a": arrayLeaf(intLeaf("1"))},
		},
		{
			name:  "two integers",
			input: "xs = [1, 2]\n",
			want:  map[string]any{"xs": arrayLeaf(intLeaf("1"), intLeaf("2"))},
		},
		{
			name:  "spaces and tabs around elements",
			input: "a = [ 1 ,\t2 ]\n",
			want:  map[string]any{"a": arrayLeaf(intLeaf("1"), intLeaf("2"))},
		},
		{
			name:  "mixed types in one array",
			input: "a = [1, \"x\", true, 1.5]\n",
			want: map[string]any{"a": arrayLeaf(
				intLeaf("1"),
				stringLeaf("x"),
				map[string]any{"type": "bool", "value": "true"},
				map[string]any{"type": "float", "value": "1.5"},
			)},
		},
		{
			name:  "trailing comma",
			input: "a = [1, 2,]\n",
			want:  map[string]any{"a": arrayLeaf(intLeaf("1"), intLeaf("2"))},
		},
		{
			name:  "nested arrays",
			input: "a = [[1], [2, 3], []]\n",
			want: map[string]any{"a": arrayLeaf(
				arrayLeaf(intLeaf("1")),
				arrayLeaf(intLeaf("2"), intLeaf("3")),
				arrayLeaf(),
			)},
		},
		{
			name:  "comma and bracket inside a string element",
			input: "a = [\"x,]\", 'y]']\n",
			want:  map[string]any{"a": arrayLeaf(stringLeaf("x,]"), stringLeaf("y]"))},
		},
		{
			name:  "multi-line array with comments and trailing comma",
			input: "a = [\n  # first\n  1, # one\n  2,\n  # last\n]\n",
			want:  map[string]any{"a": arrayLeaf(intLeaf("1"), intLeaf("2"))},
		},
		{
			name:  "multi-line array with CRLF line breaks",
			input: "a = [\r\n  1,\r\n  2\r\n]\r\n",
			want:  map[string]any{"a": arrayLeaf(intLeaf("1"), intLeaf("2"))},
		},
		{
			name:  "array element is an inline table",
			input: "a = [{ x = 1 }, { x = 2 }]\n",
			want: map[string]any{"a": arrayLeaf(
				map[string]any{"x": intLeaf("1")},
				map[string]any{"x": intLeaf("2")},
			)},
		},
		{
			name:  "array in a dotted key",
			input: "a.b = [1]\n",
			want:  map[string]any{"a": map[string]any{"b": arrayLeaf(intLeaf("1"))}},
		},
		{
			name:  "array after a header",
			input: "[t]\na = [1]\n",
			want:  map[string]any{"t": map[string]any{"a": arrayLeaf(intLeaf("1"))}},
		},
		{
			name:  "array of tables element holds its own array",
			input: "[[p]]\na = [1]\n[[p]]\na = [2]\n",
			want: map[string]any{"p": []any{
				map[string]any{"a": arrayLeaf(intLeaf("1"))},
				map[string]any{"a": arrayLeaf(intLeaf("2"))},
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

func TestDecodeRejectsMalformedArrays(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "unterminated", input: "a = [1\n"},
		{name: "only opening bracket", input: "a = [\n"},
		{name: "missing comma", input: "a = [1 2]\n"},
		{name: "double comma", input: "a = [1,,2]\n"},
		{name: "leading comma", input: "a = [,1]\n"},
		{name: "lone comma", input: "a = [,]\n"},
		{name: "extra closing bracket", input: "a = [1]]\n"},
		{name: "comment runs past the closing bracket", input: "a = [1 # ]\n"},
		{name: "control character in array comment", input: "a = [# x\x01\n1]\n"},
		{name: "bare carriage return in array", input: "a = [1\r2]\n"},
		{name: "missing value after comma", input: "a = [1, = 2]\n"},
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
