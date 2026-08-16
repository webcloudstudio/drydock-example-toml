package toml

import (
	"reflect"
	"testing"
)

func TestDecodeEmptyDocument(t *testing.T) {
	got, err := Decode("")
	if err != nil {
		t.Fatalf("Decode(\"\") returned error: %v", err)
	}
	if want := map[string]any{}; !reflect.DeepEqual(got, want) {
		t.Fatalf("Decode(\"\") = %#v, want %#v", got, want)
	}
}

func TestDecodeStringsAndDottedKeys(t *testing.T) {
	got, err := Decode("basic = \"line\\n\"\nliteral = 'a # b'\nmessage = \"\"\"\\\n  hello\\\n  world\"\"\"\nroot.\"with.dot\" = true\n")
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	if got["basic"].(scalar).Value != "line\n" || got["literal"].(scalar).Value != "a # b" || got["message"].(scalar).Value != "helloworld" {
		t.Fatalf("unexpected string values: %#v", got)
	}
	root := got["root"].(map[string]any)
	if root["with.dot"].(scalar).Value != "true" {
		t.Fatalf("unexpected quoted dotted key: %#v", got)
	}
}

func TestDecodeRejectsKeyRedefinitionAndInvalidStrings(t *testing.T) {
	for _, input := range []string{"a = 1\na = 2\n", "a = 1\na.b = 2\n", "a = \"\\q\"\n", "a = '''unterminated\n"} {
		if _, err := Decode(input); err == nil {
			t.Fatalf("Decode(%q) accepted invalid TOML", input)
		}
	}
}

func TestDecodeScalarValuesAndRejections(t *testing.T) {
	input := "truth=true\nhex=0xdead_beef\noctal=0o755\nbinary=0b1010_0101\nf=3.1415\ne=-6.02e+23\nz=-0.0\ni=+inf\nn=nan\n"
	got, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	want := map[string]scalar{
		"truth": {Type: "bool", Value: "true"}, "hex": {Type: "integer", Value: "3735928559"},
		"octal": {Type: "integer", Value: "493"}, "binary": {Type: "integer", Value: "165"},
		"f": {Type: "float", Value: "3.1415"}, "e": {Type: "float", Value: "-6.02e+23"},
		"z": {Type: "float", Value: "-0.0"}, "i": {Type: "float", Value: "+inf"}, "n": {Type: "float", Value: "nan"},
	}
	for key, expected := range want {
		if actual, ok := got[key].(scalar); !ok || actual != expected {
			t.Errorf("%s = %#v, want %#v", key, got[key], expected)
		}
	}
	invalid := []string{
		"a = TRUE\n", "a = 01\n", "a = 1__0\n", "a = 1_\n", "a = 0x_1\n",
		"a = 0x10000000000000000\n", "a = 9223372036854775808\n", "a = 1.\n", "a = .1\n",
		"a = 1e\n", "a = 1e+\n", "a = 1e2.0\n", "a = +NaN\n",
		"a = +0xff\n", "a = -0o7\n",
	}
	for _, text := range invalid {
		if _, err := Decode(text); err == nil {
			t.Errorf("Decode(%q) accepted invalid scalar", text)
		}
	}
}

func TestDecodeLexicalBoundaries(t *testing.T) {
	input := "# comment with\t a hash # still comment\r\n[group] # attached\r\nvalue = \"hash # in string\" # trailing\r\n"
	got, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode returned error for valid lexical boundaries: %v", err)
	}
	group, ok := got["group"].(map[string]any)
	if !ok {
		t.Fatalf("group = %#v, want table", got["group"])
	}
	if actual := group["value"].(scalar).Value; actual != "hash # in string" {
		t.Fatalf("value = %q, want %q", actual, "hash # in string")
	}
}

func TestDecodeRejectsInvalidUTF8AndControls(t *testing.T) {
	for _, input := range []string{
		string([]byte{'a', '=', ' ', 0xff}),
		"a = 1 # bad\x01 comment\n",
		"a = 1\x00\n",
		"a = \"bad\x7f\"\n",
	} {
		if _, err := Decode(input); err == nil {
			t.Errorf("Decode(%q) accepted invalid lexical input", input)
		}
	}
}

func TestDecodeTablesAndDatetimes(t *testing.T) {
	input := "[[albums]]\nname = \"Born to Run\"\n[[albums.songs]]\nname = \"Jungleland\"\n[[albums]]\nname = \"Born in the USA\"\n[[albums.songs]]\nname = \"Glory Days\"\nwhen = 1987-07-05t17:45:00.123456789999z\nlocal = 1987-07-05 17:45:00.6\n"
	got, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	albums := got["albums"].([]any)
	if len(albums) != 2 || len(albums[0].(map[string]any)["songs"].([]any)) != 1 {
		t.Fatalf("albums = %#v, want two ordered albums with one first song", albums)
	}
	song := albums[1].(map[string]any)["songs"].([]any)[0].(map[string]any)
	if got := song["when"].(scalar); got != (scalar{Type: "datetime", Value: "1987-07-05T17:45:00.123456789Z"}) {
		t.Fatalf("when = %#v, want truncated offset datetime", got)
	}
	if got := song["local"].(scalar); got != (scalar{Type: "datetime-local", Value: "1987-07-05T17:45:00.6"}) {
		t.Fatalf("local = %#v, want local datetime", got)
	}
}

func TestDecodeArraysAllowNestingCommentsAndTrailingCommas(t *testing.T) {
	input := "values = [\n  1, # first\n  {name = \"one\", tags = [\"a\", \"b\"]},\n  [true, false,],\n]\n"
	got, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	values, ok := got["values"].([]any)
	if !ok || len(values) != 3 {
		t.Fatalf("values = %#v, want three heterogeneous elements", got["values"])
	}
	if actual := values[0].(scalar); actual != (scalar{Type: "integer", Value: "1"}) {
		t.Fatalf("values[0] = %#v, want integer 1", actual)
	}
	item := values[1].(map[string]any)
	if actual := item["name"].(scalar); actual.Value != "one" {
		t.Fatalf("inline item name = %#v, want one", actual)
	}
	tags := item["tags"].([]any)
	if len(tags) != 2 || tags[1].(scalar).Value != "b" {
		t.Fatalf("inline item tags = %#v, want [a b]", tags)
	}
	flags := values[2].([]any)
	if len(flags) != 2 || flags[0].(scalar).Value != "true" || flags[1].(scalar).Value != "false" {
		t.Fatalf("nested flags = %#v, want [true false]", flags)
	}
}

func TestDecodeInlineTablesSupportDottedKeysAndRemainClosed(t *testing.T) {
	valid, err := Decode("point = { position.x = 1, position.y = 2 }\n")
	if err != nil {
		t.Fatalf("Decode returned error for dotted inline table: %v", err)
	}
	point := valid["point"].(map[string]any)["position"].(map[string]any)
	if point["x"].(scalar).Value != "1" || point["y"].(scalar).Value != "2" {
		t.Fatalf("point.position = %#v, want x and y", point)
	}

	for _, input := range []string{
		"point = { position = { x = 1 } }\n[point.position]\ny = 2\n",
		"point = { a = 1, a.b = 2 }\n",
		"point = { a = 1, }\n",
		"point = {\n a = 1\n}\n",
	} {
		if _, err := Decode(input); err == nil {
			t.Errorf("Decode(%q) accepted invalid inline-table structure", input)
		}
	}
}
