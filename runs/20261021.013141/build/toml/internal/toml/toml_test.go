package toml

import (
	"reflect"
	"testing"
)

func TestDecodeIntegerAssignment(t *testing.T) {
	got, err := Decode([]byte("answer = 42\n"))
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	want := Value{Type: "integer", Value: "42"}
	if got["answer"] != want {
		t.Fatalf("Decode returned %#v, want %#v", got, want)
	}
}

func TestDecodeRejectsMalformedAssignment(t *testing.T) {
	if _, err := Decode([]byte("answer = [\n")); err == nil {
		t.Fatal("Decode accepted malformed TOML")
	}
}

func TestDecodeStringFormsAndComments(t *testing.T) {
	doc := "basic = \"line\\n# content\"\n" + "literal = 'C:\\\\tmp#file' # trailing comment\n" + "multi = \"\"\"\nfirst\nsecond\"\"\"\n" + "literal_multi = '''\nraw\\n# still content\n'''\n"
	got, err := Decode([]byte(doc))
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]any{"basic": Value{"string", "line\n# content"}, "literal": Value{"string", `C:\\tmp#file`}, "multi": Value{"string", "first\nsecond"}, "literal_multi": Value{"string", "raw\\n# still content\n"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v, want %#v", got, want)
	}
}

func TestDecodeScalarsAndRejectsMalformedNumbers(t *testing.T) {
	got, err := Decode([]byte("truth = false\nhex = 0xDEAD_BEEF\nfloat = -1_000.25e+2\n"))
	if err != nil {
		t.Fatal(err)
	}
	if got["truth"] != (Value{"bool", "false"}) || got["hex"] != (Value{"integer", "3735928559"}) || got["float"] != (Value{"float", "-1_000.25e+2"}) {
		t.Fatalf("unexpected scalar values: %#v", got)
	}
	for _, input := range []string{"x = 1__0\n", "x = 1_e2\n", "x = 0x_FF\n", "x = 9223372036854775808\n", "x = truth\n"} {
		if _, err := Decode([]byte(input)); err == nil {
			t.Errorf("accepted malformed scalar %q", input)
		}
	}
}

func TestDecodeRejectsInvalidUTF8AndControls(t *testing.T) {
	if _, err := Decode([]byte{'x', ' ', '=', ' ', '"', 0xff, '"'}); err == nil {
		t.Fatal("accepted invalid UTF-8")
	}
	if _, err := Decode([]byte("x = \"bad\x01\"\n")); err == nil {
		t.Fatal("accepted prohibited control character")
	}
}

func TestDecodeKeyFormsAndDottedWhitespace(t *testing.T) {
	doc := `"" = "empty"
"a b" = "quoted"
123 = "numeric-looking"
outer . "inner-key" = "nested"
`
	got, err := Decode([]byte(doc))
	if err != nil {
		t.Fatal(err)
	}
	if got[""] != (Value{"string", "empty"}) || got["a b"] != (Value{"string", "quoted"}) || got["123"] != (Value{"string", "numeric-looking"}) {
		t.Fatalf("unexpected keys: %#v", got)
	}
	outer := got["outer"].(map[string]any)
	if outer["inner-key"] != (Value{"string", "nested"}) {
		t.Fatalf("unexpected dotted key: %#v", outer)
	}
}

func TestDecodeRejectsDottedTableConflicts(t *testing.T) {
	for _, doc := range []string{
		"name = \"Tom\"\nname = \"Pradyun\"\n",
		"a.b = 1\n[a]\nc = 2\n",
		"[fruit]\napple.color = \"red\"\n[fruit.apple]\n",
	} {
		if _, err := Decode([]byte(doc)); err == nil {
			t.Fatalf("accepted conflicting TOML: %q", doc)
		}
	}
}

func TestDecodeArraysOfTablesResolveMostRecentElement(t *testing.T) {
	doc := `[[fruits]]
name = "apple"
[[fruits.varieties]]
name = "red delicious"
[[fruits]]
name = "banana"
[[fruits.varieties]]
name = "cavendish"
`
	got, err := Decode([]byte(doc))
	if err != nil {
		t.Fatal(err)
	}
	fruits := got["fruits"].([]any)
	if len(fruits) != 2 {
		t.Fatalf("got %d fruits", len(fruits))
	}
	first := fruits[0].(map[string]any)["varieties"].([]any)
	second := fruits[1].(map[string]any)["varieties"].([]any)
	if len(first) != 1 || len(second) != 1 {
		t.Fatalf("wrong nested array lengths: %#v", got)
	}
}

func TestDecodeArraysAllowNestingCommentsAndTrailingCommas(t *testing.T) {
	doc := "values = [\n  [],\n  [1, 2,], # a permitted array trailing comma\n  [{ name = \"Nail\", meta = { edible = false } }],\n]\n"
	got, err := Decode([]byte(doc))
	if err != nil {
		t.Fatal(err)
	}
	values := got["values"].([]any)
	if len(values) != 3 || len(values[0].([]any)) != 0 || len(values[1].([]any)) != 2 {
		t.Fatalf("unexpected arrays: %#v", got)
	}
	item := values[2].([]any)[0].(map[string]any)
	meta := item["meta"].(map[string]any)
	if meta["edible"] != (Value{Type: "bool", Value: "false"}) {
		t.Fatalf("unexpected nested inline table: %#v", item)
	}
}

func TestDecodeRejectsInlineTableTrailingCommaMultilineAndExtension(t *testing.T) {
	for _, doc := range []string{
		"x = { a = 1, }\n",
		"x = {\n a = 1\n}\n",
		"[product]\ntype = { name = \"Nail\" }\ntype.edible = false\n",
	} {
		if _, err := Decode([]byte(doc)); err == nil {
			t.Errorf("accepted invalid inline table document: %q", doc)
		}
	}
}

func TestDecodeTemporalTypesAndFractionTruncation(t *testing.T) {
	doc := "odt = 1979-05-27 07:32:00.12345678999Z\n" +
		"ldt = 1979-05-27T07:32:00.12345678999\n" +
		"date = 1979-05-27\ntime = 07:32:00.5\n"
	got, err := Decode([]byte(doc))
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]Value{
		"odt":  {"datetime", "1979-05-27T07:32:00.123456789Z"},
		"ldt":  {"datetime-local", "1979-05-27T07:32:00.123456789"},
		"date": {"date-local", "1979-05-27"},
		"time": {"time-local", "07:32:00.5"},
	}
	for key, expected := range want {
		if got[key] != expected {
			t.Errorf("%s = %#v, want %#v", key, got[key], expected)
		}
	}
}

func TestDecodeRejectsInvalidTemporalValues(t *testing.T) {
	for _, doc := range []string{
		"date = 1979-02-30\n",
		"time = 24:00:00\n",
		"datetime = 1979-05-27T07:32:00+24:00\n",
		"datetime = 1979-05-27T07:32:00.\n",
	} {
		if _, err := Decode([]byte(doc)); err == nil {
			t.Errorf("accepted invalid temporal document: %q", doc)
		}
	}
}
