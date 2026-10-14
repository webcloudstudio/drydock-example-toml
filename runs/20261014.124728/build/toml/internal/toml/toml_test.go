package toml

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func decodeString(input string) (map[string]any, error) {
	return Decode(strings.NewReader(input))
}

func TestDecodeEmptyDocumentHasNoKeys(t *testing.T) {
	doc, err := decodeString("")
	if err != nil {
		t.Fatalf("Decode(\"\") error = %v, want nil", err)
	}
	if len(doc) != 0 {
		t.Errorf("Decode(\"\") = %v, want no keys", doc)
	}
}

func TestDecodeReadsEachStringAssignment(t *testing.T) {
	doc, err := decodeString("a = \"one\"\nb = 'two'\n")
	if err != nil {
		t.Fatalf("Decode error = %v, want nil", err)
	}
	want := map[string]any{
		"a": map[string]any{"type": "string", "value": "one"},
		"b": map[string]any{"type": "string", "value": "two"},
	}
	if !reflect.DeepEqual(doc, want) {
		t.Errorf("Decode = %v, want %v", doc, want)
	}
}

func TestDecodeSkipsCommentsAndBlankLines(t *testing.T) {
	doc, err := decodeString("# leading comment\n\n   \n  s = \"v\" # trailing comment\n\n# end\n")
	if err != nil {
		t.Fatalf("Decode error = %v, want nil", err)
	}
	want := map[string]any{"s": map[string]any{"type": "string", "value": "v"}}
	if !reflect.DeepEqual(doc, want) {
		t.Errorf("Decode = %v, want %v", doc, want)
	}
}

func TestDecodeAcceptsFinalLineWithoutNewline(t *testing.T) {
	doc, err := decodeString("s = \"v\"")
	if err != nil {
		t.Fatalf("Decode error = %v, want nil", err)
	}
	want := map[string]any{"s": map[string]any{"type": "string", "value": "v"}}
	if !reflect.DeepEqual(doc, want) {
		t.Errorf("Decode = %v, want %v", doc, want)
	}
}

func TestDecodeAcceptsCRLFLineEndings(t *testing.T) {
	doc, err := decodeString("a = \"1\"\r\nb = \"2\"\r\n")
	if err != nil {
		t.Fatalf("Decode error = %v, want nil", err)
	}
	want := map[string]any{
		"a": map[string]any{"type": "string", "value": "1"},
		"b": map[string]any{"type": "string", "value": "2"},
	}
	if !reflect.DeepEqual(doc, want) {
		t.Errorf("Decode = %v, want %v", doc, want)
	}
}

func TestDecodeAcceptsQuotedKeys(t *testing.T) {
	doc, err := decodeString("\"quoted key\" = \"v\"\n'literal key' = \"w\"\n0 = \"x\"\n")
	if err != nil {
		t.Fatalf("Decode error = %v, want nil", err)
	}
	want := map[string]any{
		"quoted key":  map[string]any{"type": "string", "value": "v"},
		"literal key": map[string]any{"type": "string", "value": "w"},
		"0":           map[string]any{"type": "string", "value": "x"},
	}
	if !reflect.DeepEqual(doc, want) {
		t.Errorf("Decode = %v, want %v", doc, want)
	}
}

func TestDecodeRejectsMalformedDocuments(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "missing equals", input: "s \"v\"\n"},
		{name: "missing value", input: "s =\n"},
		{name: "key only", input: "=\"v\"\n"},
		{name: "duplicate key", input: "s = \"a\"\ns = \"b\"\n"},
		{name: "text after value", input: "s = \"a\" \"b\"\n"},
		{name: "bare carriage return", input: "s = \"a\"\rt = \"b\"\n"},
		{name: "invalid utf-8", input: "s = \"\xff\"\n"},
		{name: "control character in comment", input: "# a\x01\ns = \"v\"\n"},
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
