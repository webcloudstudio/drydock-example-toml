package toml

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

// scalarValueDoc is the tagged model a document `n = <value>` must decode to.
func scalarValueDoc(kind, value string) map[string]any {
	return map[string]any{"n": map[string]any{"type": kind, "value": value}}
}

func TestDecodeBooleans(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "true", input: "true", want: "true"},
		{name: "false", input: "false", want: "false"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := decodeString("n = " + tt.input + "\n")
			if err != nil {
				t.Fatalf("Decode(%q) error = %v, want nil", tt.input, err)
			}
			want := scalarValueDoc("bool", tt.want)
			if !reflect.DeepEqual(doc, want) {
				t.Errorf("Decode(%q) = %v, want %v", tt.input, doc, want)
			}
		})
	}
}

func TestDecodeIntegerForms(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "zero", input: "0", want: "0"},
		{name: "positive", input: "42", want: "42"},
		{name: "explicit plus", input: "+99", want: "99"},
		{name: "negative", input: "-17", want: "-17"},
		{name: "underscore groups", input: "1_000_000", want: "1000000"},
		{name: "signed underscores", input: "-1_2", want: "-12"},
		{name: "hex", input: "0xDEADBEEF", want: "3735928559"},
		{name: "hex lowercase with underscore", input: "0xdead_beef", want: "3735928559"},
		{name: "octal", input: "0o755", want: "493"},
		{name: "octal with underscore", input: "0o7_7", want: "63"},
		{name: "binary", input: "0b11010110", want: "214"},
		{name: "binary with underscore", input: "0b1101_0110", want: "214"},
		{name: "largest signed 64-bit", input: "9223372036854775807", want: "9223372036854775807"},
		{name: "smallest signed 64-bit", input: "-9223372036854775808", want: "-9223372036854775808"},
		{name: "largest hex signed 64-bit", input: "0x7FFFFFFFFFFFFFFF", want: "9223372036854775807"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := decodeString("n = " + tt.input + "\n")
			if err != nil {
				t.Fatalf("Decode(%q) error = %v, want nil", tt.input, err)
			}
			want := scalarValueDoc("integer", tt.want)
			if !reflect.DeepEqual(doc, want) {
				t.Errorf("Decode(%q) = %v, want %v", tt.input, doc, want)
			}
		})
	}
}

func TestDecodeFloatForms(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "simple", input: "3.14", want: "3.14"},
		{name: "zero", input: "0.0", want: "0"},
		{name: "negative zero", input: "-0.0", want: "-0"},
		{name: "explicit plus", input: "+1.5", want: "1.5"},
		{name: "negative", input: "-0.5", want: "-0.5"},
		{name: "underscore", input: "6_626.0", want: "6626"},
		{name: "exponent", input: "1e6", want: "1e+06"},
		{name: "signed exponent", input: "6.626e-34", want: "6.626e-34"},
		{name: "capital exponent with plus", input: "5E+22", want: "5e+22"},
		{name: "zero prefixed exponent", input: "1e05", want: "100000"},
		{name: "zero integer part with exponent", input: "0e0", want: "0"},
		{name: "fraction with exponent", input: "-2.5e3", want: "-2500"},
		{name: "smallest denormal", input: "5e-324", want: "5e-324"},
		{name: "largest finite", input: "1.7976931348623157e308", want: "1.7976931348623157e+308"},
		{name: "positive infinity", input: "inf", want: "inf"},
		{name: "explicit plus infinity", input: "+inf", want: "inf"},
		{name: "negative infinity", input: "-inf", want: "-inf"},
		{name: "not a number", input: "nan", want: "nan"},
		{name: "signed not a number", input: "+nan", want: "nan"},
		{name: "negative not a number", input: "-nan", want: "nan"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := decodeString("n = " + tt.input + "\n")
			if err != nil {
				t.Fatalf("Decode(%q) error = %v, want nil", tt.input, err)
			}
			want := scalarValueDoc("float", tt.want)
			if !reflect.DeepEqual(doc, want) {
				t.Errorf("Decode(%q) = %v, want %v", tt.input, doc, want)
			}
		})
	}
}

func TestDecodeRejectsInvalidScalars(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		// Integers
		{name: "integer one past signed 64-bit max", input: "9223372036854775808"},
		{name: "integer one past signed 64-bit min", input: "-9223372036854775809"},
		{name: "integer far beyond range", input: "99999999999999999999999"},
		{name: "hex beyond signed 64-bit", input: "0x8000000000000000"},
		{name: "hex overflow to 16 digits", input: "0xFFFFFFFFFFFFFFFF"},
		{name: "octal beyond signed 64-bit", input: "0o1000000000000000000000"},
		{name: "binary beyond signed 64-bit", input: "0b" + strings.Repeat("1", 64)},
		{name: "leading zero", input: "01"},
		{name: "leading zero with sign", input: "-00"},
		{name: "trailing underscore", input: "1_"},
		{name: "leading underscore", input: "_1"},
		{name: "double underscore", input: "1__0"},
		{name: "underscore after prefix", input: "0x_1"},
		{name: "underscore after sign", input: "+_1"},
		{name: "zero with underscore", input: "0_0"},
		{name: "plus on hex", input: "+0x1"},
		{name: "minus on hex", input: "-0x1"},
		{name: "plus on octal", input: "+0o7"},
		{name: "plus on binary", input: "+0b1"},
		{name: "uppercase hex prefix", input: "0XFF"},
		{name: "digit out of octal range", input: "0o8"},
		{name: "digit out of binary range", input: "0b2"},
		{name: "empty hex", input: "0x"},
		{name: "empty octal", input: "0o"},
		{name: "empty binary", input: "0b"},
		{name: "hex invalid digit", input: "0xG"},

		// Floats
		{name: "float without digits after point", input: "1."},
		{name: "float without digits before point", input: ".5"},
		{name: "float with signed point only", input: "-."},
		{name: "float exponent without digits", input: "1e"},
		{name: "float exponent sign without digits", input: "1e+"},
		{name: "float point before exponent without fraction", input: "1.e5"},
		{name: "float leading zero before fraction", input: "00.5"},
		{name: "float leading zero before exponent-free integer", input: "01.5"},
		{name: "float underscore after point", input: "1._5"},
		{name: "float trailing underscore in fraction", input: "1.5_"},
		{name: "float underscore before exponent", input: "1_e5"},
		{name: "float underscore after exponent sign", input: "1e+_5"},
		{name: "float exponent only", input: "e5"},
		{name: "float overflow", input: "1e309"},
		{name: "float negative overflow", input: "-1e309"},
		{name: "float with hex-style exponent", input: "1p5"},
		{name: "special float mixed case", input: "Inf"},
		{name: "special float nan mixed case", input: "NaN"},
		{name: "special float with trailing text", input: "infinity"},
		{name: "special float with underscore", input: "in_f"},

		// Other tokens
		{name: "bool mixed case", input: "True"},
		{name: "bool with trailing text", input: "trueish"},
		{name: "bare sign", input: "+"},
		{name: "bare minus", input: "-"},
		{name: "value with trailing junk", input: "1abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := decodeString("n = " + tt.input + "\n")
			if !errors.Is(err, ErrSyntax) {
				t.Fatalf("Decode(%q) error = %v, want %v", tt.input, err, ErrSyntax)
			}
			if doc != nil {
				t.Errorf("Decode(%q) = %v, want nil document", tt.input, doc)
			}
		})
	}
}

func TestDecodeScalarsInOneDocument(t *testing.T) {
	doc, err := decodeString("b = false\ni = -3\nf = 2.5\ns = \"x\"\n")
	if err != nil {
		t.Fatalf("Decode error = %v, want nil", err)
	}
	want := map[string]any{
		"b": map[string]any{"type": "bool", "value": "false"},
		"i": map[string]any{"type": "integer", "value": "-3"},
		"f": map[string]any{"type": "float", "value": "2.5"},
		"s": map[string]any{"type": "string", "value": "x"},
	}
	if !reflect.DeepEqual(doc, want) {
		t.Errorf("Decode = %v, want %v", doc, want)
	}
}

func TestDecodeScalarFollowedByCommentAndCRLF(t *testing.T) {
	doc, err := decodeString("i = 7 # seven\r\nf = 1.0\r\n")
	if err != nil {
		t.Fatalf("Decode error = %v, want nil", err)
	}
	want := map[string]any{
		"i": map[string]any{"type": "integer", "value": "7"},
		"f": map[string]any{"type": "float", "value": "1"},
	}
	if !reflect.DeepEqual(doc, want) {
		t.Errorf("Decode = %v, want %v", doc, want)
	}
}
