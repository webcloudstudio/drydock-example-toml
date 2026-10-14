package toml

import (
	"errors"
	"reflect"
	"testing"
)

func TestDecodeAcceptsBareKeys(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[string]any
	}{
		{
			name:  "letters digits underscore and dash",
			input: "bare-key_1 = \"v\"\n",
			want:  map[string]any{"bare-key_1": stringLeaf("v")},
		},
		{
			name:  "all digits",
			input: "1234 = \"v\"\n",
			want:  map[string]any{"1234": stringLeaf("v")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := decodeString(tt.input)
			if err != nil {
				t.Fatalf("Decode(%q) error = %v, want nil", tt.input, err)
			}
			if !reflect.DeepEqual(doc, tt.want) {
				t.Errorf("Decode(%q) = %v, want %v", tt.input, doc, tt.want)
			}
		})
	}
}

func TestDecodeDottedKeyCreatesIntermediateTables(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[string]any
	}{
		{
			name:  "two parts",
			input: "physical.color = \"orange\"\n",
			want: map[string]any{
				"physical": map[string]any{"color": stringLeaf("orange")},
			},
		},
		{
			name:  "three parts",
			input: "a.b.c = 1\n",
			want: map[string]any{
				"a": map[string]any{
					"b": map[string]any{"c": map[string]any{"type": "integer", "value": "1"}},
				},
			},
		},
		{
			name:  "siblings share an intermediate table",
			input: "a.b = \"x\"\na.c = \"y\"\n",
			want: map[string]any{
				"a": map[string]any{"b": stringLeaf("x"), "c": stringLeaf("y")},
			},
		},
		{
			name:  "spaces and tabs around parts are ignored",
			input: "a . b\t.\tc = \"v\"\n",
			want: map[string]any{
				"a": map[string]any{"b": map[string]any{"c": stringLeaf("v")}},
			},
		},
		{
			name:  "quoted parts may mix with bare parts",
			input: "\"a b\".'c'.d = \"v\"\n",
			want: map[string]any{
				"a b": map[string]any{"c": map[string]any{"d": stringLeaf("v")}},
			},
		},
		{
			name:  "dot inside a quoted part is not a separator",
			input: "\"a.b\" = \"v\"\n",
			want:  map[string]any{"a.b": stringLeaf("v")},
		},
		{
			name:  "dotted key after a sibling top-level key",
			input: "top = \"t\"\na.b = \"v\"\n",
			want: map[string]any{
				"top": stringLeaf("t"),
				"a":   map[string]any{"b": stringLeaf("v")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := decodeString(tt.input)
			if err != nil {
				t.Fatalf("Decode(%q) error = %v, want nil", tt.input, err)
			}
			if !reflect.DeepEqual(doc, tt.want) {
				t.Errorf("Decode(%q) = %v, want %v", tt.input, doc, tt.want)
			}
		})
	}
}

func TestDecodeTreatsBareAndQuotedKeysAsEquivalent(t *testing.T) {
	doc, err := decodeString("name = \"a\"\n\"name\" = \"b\"\n")
	if !errors.Is(err, ErrSyntax) {
		t.Fatalf("Decode error = %v, want %v", err, ErrSyntax)
	}
	if doc != nil {
		t.Errorf("Decode = %v, want nil document", doc)
	}

	doc, err = decodeString("\"name\".x = \"a\"\nname.y = \"b\"\n")
	if err != nil {
		t.Fatalf("Decode error = %v, want nil", err)
	}
	want := map[string]any{
		"name": map[string]any{"x": stringLeaf("a"), "y": stringLeaf("b")},
	}
	if !reflect.DeepEqual(doc, want) {
		t.Errorf("Decode = %v, want %v", doc, want)
	}
}

func TestDecodeRejectsKeyDefinitionErrors(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "same bare key twice", input: "name = \"Tom\"\nname = \"Pradyun\"\n"},
		{name: "bare then quoted with same name", input: "name = \"a\"\n\"name\" = \"b\"\n"},
		{name: "same dotted key twice", input: "a.b = 1\na.b = 2\n"},
		{name: "dotted key redefines a table", input: "a.b.c = 1\na.b = 2\n"},
		{name: "table then dotted key redefines leaf", input: "a.b = 1\na = 2\n"},
		{name: "integer used as a table", input: "a = 1\na.b = 2\n"},
		{name: "string used as a table", input: "s = \"x\"\ns.t = \"y\"\n"},
		{name: "integer used as a table two levels down", input: "a.b = 1\na.b.c = 2\n"},
		{name: "empty part before dot", input: ".b = \"v\"\n"},
		{name: "empty part after dot", input: "a. = \"v\"\n"},
		{name: "empty middle part", input: "a..b = \"v\"\n"},
		{name: "dot without key before equals", input: "a. \"v\"\n"},
		{name: "key part with invalid character", input: "a+b = \"v\"\n"},
		{name: "table header after dotted key", input: "a.b = \"v\"\n[a]\n"},
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

// stringLeaf is the tagged model for a string value, shared by the key tests.
func stringLeaf(v string) map[string]any {
	return map[string]any{"type": "string", "value": v}
}
