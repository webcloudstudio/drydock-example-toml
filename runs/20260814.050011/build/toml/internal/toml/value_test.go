package toml

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDecodeEmptyDocumentIsEmptyObject(t *testing.T) {
	got, err := Decode("\n# comment\n")
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	if !reflect.DeepEqual(got, Object{}) {
		t.Fatalf("Decode returned %#v, want empty object", got)
	}
}

func TestTaggedPreservesScalarLexemes(t *testing.T) {
	got, err := Tagged(Object{"n": Value{Kind: Integer, Lexeme: "9223372036854775807"}})
	if err != nil {
		t.Fatalf("Tagged returned error: %v", err)
	}
	want := map[string]any{"n": map[string]any{"type": "integer", "value": "9223372036854775807"}}
	data, err := json.Marshal(got)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if !reflect.DeepEqual(decoded, want) {
		t.Fatalf("JSON = %#v, want %#v", decoded, want)
	}
}

func TestTaggedEncodesEveryScalarKindWithStringPayload(t *testing.T) {
	input := Object{
		"string":         Value{Kind: String, Lexeme: "hello"},
		"integer":        Value{Kind: Integer, Lexeme: "42"},
		"float":          Value{Kind: Float, Lexeme: "3.5"},
		"bool":           Value{Kind: Bool, Lexeme: "true"},
		"datetime":       Value{Kind: Datetime, Lexeme: "1979-05-27T07:32:00Z"},
		"datetime-local": Value{Kind: DatetimeLocal, Lexeme: "1979-05-27T07:32:00"},
		"date-local":     Value{Kind: DateLocal, Lexeme: "1979-05-27"},
		"time-local":     Value{Kind: TimeLocal, Lexeme: "07:32:00"},
	}

	got, err := Tagged(input)
	if err != nil {
		t.Fatalf("Tagged returned error: %v", err)
	}
	data, err := json.Marshal(got)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}

	var decoded map[string]map[string]any
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	for name, want := range input {
		scalar, ok := want.(Value)
		if !ok {
			t.Fatalf("input %s is %T, want Value", name, want)
		}
		encoded := decoded[name]
		if encoded["type"] != string(scalar.Kind) {
			t.Errorf("%s type = %#v, want %q", name, encoded["type"], scalar.Kind)
		}
		if encoded["value"] != scalar.Lexeme {
			t.Errorf("%s value = %#v, want %q", name, encoded["value"], scalar.Lexeme)
		}
	}
}

func TestTaggedPreservesNestedAndEmptyContainers(t *testing.T) {
	input := Object{
		"empty-table": Object{},
		"empty-array": []any{},
		"nested": Object{
			"values": []any{
				Value{Kind: Integer, Lexeme: "1"},
				Object{"name": Value{Kind: String, Lexeme: "x"}},
			},
		},
	}

	got, err := Tagged(input)
	if err != nil {
		t.Fatalf("Tagged returned error: %v", err)
	}
	data, err := json.Marshal(got)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}

	var decoded map[string]any
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if len(decoded["empty-table"].(map[string]any)) != 0 {
		t.Errorf("empty table = %#v, want empty object", decoded["empty-table"])
	}
	if len(decoded["empty-array"].([]any)) != 0 {
		t.Errorf("empty array = %#v, want empty array", decoded["empty-array"])
	}
	nested := decoded["nested"].(map[string]any)
	values := nested["values"].([]any)
	if values[0].(map[string]any)["value"] != "1" {
		t.Errorf("nested integer = %#v, want tagged value 1", values[0])
	}
	if values[1].(map[string]any)["name"].(map[string]any)["value"] != "x" {
		t.Errorf("nested string = %#v, want tagged value x", values[1])
	}
}

func TestDecodeRejectsDuplicateKeys(t *testing.T) {
	if _, err := Decode("a = 1\na = 2"); err == nil {
		t.Fatal("Decode accepted duplicate key")
	}
}

func TestDecodeStandardTablesAndImplicitParents(t *testing.T) {
	got, err := Decode("[a.b.c]\nvalue = 1\n[a]\nroot = 2\n")
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	a := got["a"].(Object)
	if a["root"] != (Value{Integer, "2"}) {
		t.Fatalf("root table value = %#v", a["root"])
	}
	if a["b"].(Object)["c"].(Object)["value"] != (Value{Integer, "1"}) {
		t.Fatalf("nested table = %#v", a)
	}
}

func TestDecodeRejectsExtensionsOfInlineTables(t *testing.T) {
	for _, input := range []string{
		"type = {name = 'Nail'}\ntype.edible = false\n",
		"product = {name = 'Nail'}\n[product.details]\n",
	} {
		if _, err := Decode(input); err == nil {
			t.Errorf("Decode accepted inline-table extension %q", input)
		}
	}
}

func TestDecodeAllowsIndependentInlineTablesInTableArrayElements(t *testing.T) {
	_, err := Decode("[[items]]\nvalue = {a.b = 1}\n[[items]]\nvalue = {a.b = 2}\n")
	if err != nil {
		t.Fatalf("Decode rejected independent array-table elements: %v", err)
	}
}

func TestDecodeArraysOfTablesPreserveElementsAndNestedTargets(t *testing.T) {
	input := "[[fruits]]\nname = \"apple\"\n[fruits.physical]\ncolor = \"red\"\n[[fruits.varieties]]\nname = \"red delicious\"\n[[fruits]]\nname = \"banana\"\n[fruits.physical]\ncolor = \"yellow\"\n[[fruits.varieties]]\nname = \"plantain\"\n"
	got, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	fruits, ok := got["fruits"].([]any)
	if !ok || len(fruits) != 2 {
		t.Fatalf("fruits = %#v, want two ordered elements", got["fruits"])
	}
	first := fruits[0].(Object)
	second := fruits[1].(Object)
	if first["name"] != (Value{String, "apple"}) || second["name"] != (Value{String, "banana"}) {
		t.Fatalf("fruit order = %#v, want apple then banana", fruits)
	}
	if first["physical"].(Object)["color"] != (Value{String, "red"}) || second["physical"].(Object)["color"] != (Value{String, "yellow"}) {
		t.Fatalf("physical tables attached to wrong elements: %#v", fruits)
	}
	firstVarieties := first["varieties"].([]any)
	secondVarieties := second["varieties"].([]any)
	if firstVarieties[0].(Object)["name"] != (Value{String, "red delicious"}) || secondVarieties[0].(Object)["name"] != (Value{String, "plantain"}) {
		t.Fatalf("nested array tables attached to wrong elements: %#v", fruits)
	}
}

func TestDecodeArraysOfTablesKeepEmptyElements(t *testing.T) {
	got, err := Decode("[[items]]\n[[items]]\nvalue = 2\n")
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	items := got["items"].([]any)
	if len(items) != 2 || len(items[0].(Object)) != 0 {
		t.Fatalf("items = %#v, want an empty first element and a second element", items)
	}
	if items[1].(Object)["value"] != (Value{Integer, "2"}) {
		t.Fatalf("second item = %#v, want value 2", items[1])
	}
}

func TestDecodeRejectsTableStructureConflicts(t *testing.T) {
	inputs := []string{
		"value = 1\n[value]\n",
		"[value]\n[value]\n",
		"value = {nested = 1}\n[value.more]\n",
		"value = []\n[[value]]\n",
		"[[value]]\n[value]\n",
		"a.b = 1\n[a]\n",
	}
	for _, input := range inputs {
		if _, err := Decode(input); err == nil {
			t.Errorf("Decode accepted conflicting table structure %q", input)
		}
	}
}

func TestDecodeRejectsArrayOfTablesOrderingAndKindConflicts(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "child table before parent array element",
			input: "[fruit.physical]\ncolor = \"red\"\n[[fruit]]\nname = \"apple\"\n",
		},
		{
			name:  "child array before parent array element",
			input: "[[fruit.varieties]]\nname = \"red delicious\"\n[[fruit]]\nname = \"apple\"\n",
		},
		{
			name:  "static array cannot become array of tables",
			input: "fruit = []\n[[fruit]]\nname = \"apple\"\n",
		},
		{
			name:  "static array cannot contain array of tables",
			input: "fruit = [{name = \"apple\"}]\n[[fruit.varieties]]\nname = \"red delicious\"\n",
		},
		{
			name:  "array of tables cannot become normal table",
			input: "[[fruit]]\nname = \"apple\"\n[fruit]\ncolor = \"red\"\n",
		},
		{
			name:  "normal table cannot become array of tables",
			input: "[fruit]\nname = \"apple\"\n[[fruit]]\ncolor = \"red\"\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if _, err := Decode(test.input); err == nil {
				t.Fatalf("Decode accepted invalid table ordering: %q", test.input)
			}
		})
	}
}

func TestDecodeAllowsArrayOfTablesUnderNormalTable(t *testing.T) {
	input := "[fruit]\nname = \"apple\"\n[[fruit.varieties]]\nname = \"red delicious\"\n"
	got, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode rejected array of tables under normal table: %v", err)
	}
	fruit := got["fruit"].(Object)
	varieties := fruit["varieties"].([]any)
	if len(varieties) != 1 || varieties[0].(Object)["name"] != (Value{String, "red delicious"}) {
		t.Fatalf("varieties = %#v, want one red delicious element", varieties)
	}
}

func TestDecodeLexicalNumbers(t *testing.T) {
	input := "decimal = -9223372036854775808\nhex = 0xDEAD_BEEF\nfloat = 1_000.5e+2\ninf = -inf\nnan = +nan\nenabled = true\n"
	got, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	want := map[string]Value{
		"decimal": {Integer, "-9223372036854775808"},
		"hex":     {Integer, "3735928559"},
		"float":   {Float, "1_000.5e+2"},
		"inf":     {Float, "-inf"},
		"nan":     {Float, "nan"},
		"enabled": {Bool, "true"},
	}
	for key, expected := range want {
		if got[key] != expected {
			t.Errorf("%s = %#v, want %#v", key, got[key], expected)
		}
	}
}

func TestDecodeLexicalDatetimes(t *testing.T) {
	input := "offset = 1979-05-27t07:32:00.123456789123z\nlocal = 1979-05-27 07:32:00\ndate = 1979-05-27\ntime = 07:32:00.123456789123\n"
	got, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	checks := map[string]Value{
		"offset": {Datetime, "1979-05-27t07:32:00.123456789123z"},
		"local":  {DatetimeLocal, "1979-05-27 07:32:00"},
		"date":   {DateLocal, "1979-05-27"},
		"time":   {TimeLocal, "07:32:00.123456789123"},
	}
	for key, expected := range checks {
		if got[key] != expected {
			t.Errorf("%s = %#v, want %#v", key, got[key], expected)
		}
	}
}

func TestDecodeRejectsMalformedLexicalNumbersAndDatetimes(t *testing.T) {
	for _, input := range []string{
		"value = 01\n", "value = 1_.0\n", "value = 0x_FF\n",
		"value = 1985-06-18 17:04:07+12:60\n", "value = 2020-02-30\n",
	} {
		if _, err := Decode(input); err == nil {
			t.Errorf("Decode accepted malformed value %q", input)
		}
	}
}

func TestDecodeStringForms(t *testing.T) {
	input := "basic = \"line\\nvalue\"\n" +
		"literal = 'C:\\\\Users\\\\nodejs'\n" +
		"multiline = \"\"\"\nfirst \\\n+" + "second\"\"\"\n" +
		"multiline-literal = '''\nraw\\nvalue'''\n"
	input = "basic = \"line\\nvalue\"\n" + "literal = 'C:\\\\Users\\\\nodejs'\n" + "multiline = \"\"\"\nfirst\nsecond\"\"\"\n" + "multiline-literal = '''\nraw\\nvalue'''\n"
	got, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	want := map[string]string{
		"basic":             "line\nvalue",
		"literal":           `C:\\Users\\nodejs`,
		"multiline":         "first\nsecond",
		"multiline-literal": "raw\\nvalue",
	}
	for key, expected := range want {
		value, ok := got[key].(Value)
		if !ok {
			t.Fatalf("%s has type %T, want Value", key, got[key])
		}
		if value.Kind != String || value.Lexeme != expected {
			t.Errorf("%s = %#v, want string %q", key, value, expected)
		}
	}
}

func TestDecodeRejectsInvalidStringEscapesAndControls(t *testing.T) {
	tests := []string{
		`value = "bad\q"`,
		`value = "bad\uD800"`,
		"value = \"bad\x00\"",
		"value = \"bad\nline\"",
		`value = "unterminated`,
	}
	for _, input := range tests {
		if _, err := Decode(input); err == nil {
			t.Errorf("Decode accepted invalid string %q", input)
		}
	}
}

func TestDecodeCommentsInsideMultilineContainers(t *testing.T) {
	input := "arr = [\n  1, # first\n  2,\n  \"#\", # hash is string content\n  \"\"\"###\"\"\"\n]\n"
	got, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	arr, ok := got["arr"].([]any)
	if !ok || len(arr) != 4 {
		t.Fatalf("arr = %#v, want four values", got["arr"])
	}
	if arr[2] != (Value{String, "#"}) || arr[3] != (Value{String, "###"}) {
		t.Fatalf("arr string values = %#v, want hash values", arr)
	}
}

func TestDecodeAcceptsCRLFAndMultilineContinuation(t *testing.T) {
	input := "message = \"\"\"first \\\r\n  second\"\"\"\r\nvalues = [1,\r\n  2]\r\n"
	got, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	if got["message"] != (Value{String, "first second"}) {
		t.Fatalf("message = %#v, want continued multiline string", got["message"])
	}
	if len(got["values"].([]any)) != 2 {
		t.Fatalf("values = %#v, want two values", got["values"])
	}
}

func TestDecodeRejectsInvalidUTF8AndDocumentControls(t *testing.T) {
	for name, input := range map[string]string{
		"invalid utf8": string([]byte{'a', '=', ' ', 0xff}),
		"nul":          "a = \"ok\"\x00",
		"form feed":    "a = 1\f",
		"delete":       "a = 1\x7f",
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := Decode(input); err == nil {
				t.Fatalf("Decode accepted %q", input)
			}
		})
	}
}

func TestDecodeKeyFormsCreateDecodedNestedPaths(t *testing.T) {
	got, err := Decode("bare-key = 1\n\"quoted.key\" = 2\n'raw.key' = 3\na . \"b.c\" . d = 4\n")
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	want := Object{
		"bare-key":   Value{Integer, "1"},
		"quoted.key": Value{Integer, "2"},
		"raw.key":    Value{Integer, "3"},
		"a":          Object{"b.c": Object{"d": Value{Integer, "4"}}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Decode = %#v, want %#v", got, want)
	}
}

func TestDecodeRejectsMalformedKeySeparators(t *testing.T) {
	for _, input := range []string{"a..b = 1", ".a = 1", "a. = 1", "a b = 1", "\"a\"b = 1", "\"unterminated = 1"} {
		if _, err := Decode(input); err == nil {
			t.Errorf("Decode accepted malformed key %q", input)
		}
	}
}

func TestDecodeArraysPreserveNestedOrderAndTrailingCommas(t *testing.T) {
	got, err := Decode("values = [1, [\"two\", 3],]\nempty = []\nmultiline = [\n  true, # comment\n  { nested = [4, 5,] },\n]\n")
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	values := got["values"].([]any)
	if len(values) != 2 || values[0] != (Value{Integer, "1"}) {
		t.Fatalf("values = %#v, want ordered nested array", values)
	}
	nested := values[1].([]any)
	if nested[0] != (Value{String, "two"}) || nested[1] != (Value{Integer, "3"}) {
		t.Fatalf("nested values = %#v, want string then integer", nested)
	}
	if len(got["empty"].([]any)) != 0 {
		t.Fatalf("empty array = %#v, want empty", got["empty"])
	}
}

func TestDecodeRejectsMalformedArrays(t *testing.T) {
	for _, input := range []string{"a = [, 1]", "a = [1,, 2]", "a = [1, 2", "a = [1] trailing"} {
		if _, err := Decode(input); err == nil {
			t.Errorf("Decode accepted malformed array %q", input)
		}
	}
}
