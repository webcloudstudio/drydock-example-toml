package toml

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDecodeBasicDocument(t *testing.T) {
	got, err := Decode("answer = 42\nready = true\nlabel = \"ok\"\n")
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	want := Document{
		"answer": Value{Type: "integer", Value: "42"},
		"ready":  Value{Type: "bool", Value: "true"},
		"label":  Value{Type: "string", Value: "ok"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Decode = %#v, want %#v", got, want)
	}
}

func TestDecodeEmptyDocumentIsEmptyObject(t *testing.T) {
	got, err := Decode("# comment\n\n")
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	if len(got) != 0 {
		t.Fatalf("Decode returned %d keys, want empty document", len(got))
	}
}

func TestDecodeMalformedDocumentReturnsError(t *testing.T) {
	if _, err := Decode("answer = [\n"); err == nil {
		t.Fatal("Decode returned success for malformed document")
	}
}

func TestValueUsesTaggedJSONSchema(t *testing.T) {
	encoded, err := json.Marshal(Value{Type: "integer", Value: "42"})
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	if string(encoded) != `{"type":"integer","value":"42"}` {
		t.Fatalf("Marshal = %s, want tagged value JSON", encoded)
	}
}

func TestDecodeStringFormsAndEscapes(t *testing.T) {
	got, err := Decode("basic = \"a\\n\\u03bb\"\nliteral = 'a\\nb'\nmultibasic = \"\"\"\nhello \\\n+world\"\"\"\nmultiliteral = '''\nraw\\ntext'''\n")
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	want := map[string]any{
		"basic": Value{"string", "a\nλ"}, "literal": Value{"string", `a\nb`},
		"multibasic": Value{"string", "hello +world"}, "multiliteral": Value{"string", "raw\\ntext"},
	}
	if !reflect.DeepEqual(got, Document(want)) {
		t.Fatalf("Decode = %#v, want %#v", got, want)
	}
}

func TestDecodeScalarLexingAndNormalization(t *testing.T) {
	got, err := Decode("decimal = +1_000\nhex = 0x2a\nfloat = 1.5e+2\nzero = -0\n")
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	for key, want := range map[string]Value{
		"decimal": {"integer", "1000"}, "hex": {"integer", "42"}, "float": {"float", "1.5e+2"}, "zero": {"integer", "0"},
	} {
		if got[key] != want {
			t.Errorf("%s = %#v, want %#v", key, got[key], want)
		}
	}
}

func TestDecodeRejectsInvalidEncodingAndBoundaries(t *testing.T) {
	for name, input := range map[string]string{"control": "a = \"x\u0001\"", "unicode-space": "\u3000a = 1", "bare-cr": "# x\r"} {
		t.Run(name, func(t *testing.T) {
			if _, err := Decode(input); err == nil {
				t.Fatal("Decode accepted invalid input")
			}
		})
	}
}
