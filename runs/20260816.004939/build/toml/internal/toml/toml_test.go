package toml

import "testing"

func TestDecodeEmptyDocument(t *testing.T) {
	got, err := Decode("")
	if err != nil {
		t.Fatalf("Decode(empty) returned error: %v", err)
	}
	if got == nil {
		t.Fatal("Decode(empty) returned nil document")
	}
	if len(got) != 0 {
		t.Fatalf("Decode(empty) returned %v, want empty document", got)
	}
}

func TestDecodeInteger(t *testing.T) {
	got, err := Decode("answer = 42\n")
	if err != nil {
		t.Fatalf("Decode(integer) returned error: %v", err)
	}
	want := Value{Type: "integer", Value: "42"}
	if got["answer"] != want {
		t.Fatalf("Decode(integer) = %v, want %v", got["answer"], want)
	}
}

func TestDecodeMalformedArrayReturnsError(t *testing.T) {
	if _, err := Decode("answer = [\n"); err == nil {
		t.Fatal("Decode(malformed array) returned nil error")
	}
}

func TestDecodeNormalizesIntegerValue(t *testing.T) {
	got, err := Decode("answer = 1_000\n")
	if err != nil {
		t.Fatalf("Decode(integer spelling) returned error: %v", err)
	}
	if got["answer"] != (Value{Type: "integer", Value: "1000"}) {
		t.Fatalf("Decode(integer spelling) = %v", got["answer"])
	}
}

func TestDecodeScalars(t *testing.T) {
	tests := []struct {
		name, input string
		want        Value
	}{
		{"true", "value = true", Value{Type: "bool", Value: "true"}},
		{"false", "value = false", Value{Type: "bool", Value: "false"}},
		{"hex", "value = 0xDEAD_BEEF", Value{Type: "integer", Value: "3735928559"}},
		{"octal", "value = 0o755", Value{Type: "integer", Value: "493"}},
		{"binary", "value = 0b1010", Value{Type: "integer", Value: "10"}},
		{"float", "value = 3.1415", Value{Type: "float", Value: "3.1415"}},
		{"exponent", "value = -2.5e+3", Value{Type: "float", Value: "-2.5e+3"}},
		{"special", "value = -inf", Value{Type: "float", Value: "-inf"}},
		{"negative zero", "value = -0.0", Value{Type: "float", Value: "-0.0"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.input)
			if err != nil {
				t.Fatalf("Decode(%q) returned error: %v", tt.input, err)
			}
			if got["value"] != tt.want {
				t.Fatalf("Decode(%q) = %v, want %v", tt.input, got["value"], tt.want)
			}
		})
	}
}

func TestDecodeDateTimes(t *testing.T) {
	tests := []struct {
		name, input string
		want        Value
	}{
		{"offset UTC", "value = 1979-05-27T07:32:00Z", Value{"datetime", "1979-05-27T07:32:00Z"}},
		{"offset numeric", "value = 1979-05-27 07:32:00-07:00", Value{"datetime", "1979-05-27T07:32:00-07:00"}},
		{"local datetime", "value = 1979-05-27t07:32:00.123456789987", Value{"datetime-local", "1979-05-27T07:32:00.123456789"}},
		{"local date", "value = 1979-05-27", Value{"date-local", "1979-05-27"}},
		{"local time", "value = 07:32:00.123456789987", Value{"time-local", "07:32:00.123456789"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.input)
			if err != nil {
				t.Fatalf("Decode(%q) returned error: %v", tt.input, err)
			}
			if got["value"] != tt.want {
				t.Fatalf("Decode(%q) = %v, want %v", tt.input, got["value"], tt.want)
			}
		})
	}
}

func TestDecodeRejectsInvalidDateTimes(t *testing.T) {
	inputs := []string{
		"value = 1979-02-30",
		"value = 1979-05-27T24:00:00Z",
		"value = 1979-05-27T07:32:60Z",
		"value = 1979-05-27T07:32:00+24:00",
		"value = 07:32:00.abc",
	}
	for _, input := range inputs {
		if _, err := Decode(input); err == nil {
			t.Errorf("Decode(%q) accepted invalid datetime", input)
		}
	}
}

func TestDecodeRejectsMalformedScalars(t *testing.T) {
	inputs := []string{
		"value = True", "value = TRUE", "value = 01", "value = 1__0",
		"value = 0x", "value = 0x_1", "value = 0b102", "value = 0o8",
		"value = -0x1", "value = +0b1", "value = +0o7",
		"value = 9223372036854775808", "value = 1.0.0", "value = 1e",
		"value = .5", "value = 01.0", "value = inf.0", "value = nanfoo",
	}
	for _, input := range inputs {
		if _, err := Decode(input); err == nil {
			t.Errorf("Decode(%q) accepted malformed scalar", input)
		}
	}
}

func TestDecodeStrings(t *testing.T) {
	tests := []struct {
		name, input, want string
	}{
		{"basic", `message = "hello"`, "hello"},
		{"escapes", `message = "\b\t\n\f\r\"\\"`, "\b\t\n\f\r\"\\"},
		{"unicode", `message = "\u03B1\U0001F600"`, "α😀"},
		{"literal", `message = 'a\nb#c'`, `a\nb#c`},
		{"multiline basic", "message = \"\"\"hello\nworld\"\"\"", "hello\nworld"},
		{"multiline literal", "message = '''hello\nworld'''", "hello\nworld"},
		{"continuation", "message = \"first\\\n  second\"", "firstsecond"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.input)
			if err != nil {
				t.Fatalf("Decode(%q) returned error: %v", tt.input, err)
			}
			if got["message"] != (Value{Type: "string", Value: tt.want}) {
				t.Fatalf("Decode(%q) = %v, want %q", tt.input, got["message"], tt.want)
			}
		})
	}
}

func TestDecodeRejectsMalformedStrings(t *testing.T) {
	inputs := []string{
		`message = "bad\q"`,
		`message = "bad\uD800"`,
		`message = "bad`,
		`message = "bad` + "\nnext = 1",
		`message = "bad` + string([]byte{0x01}) + `"`,
		`message = '''bad`,
	}
	for _, input := range inputs {
		if _, err := Decode(input); err == nil {
			t.Errorf("Decode(%q) accepted malformed string", input)
		}
	}
}

func TestDecodeCommentsAndLineBoundaries(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Value
	}{
		{name: "attached comment", input: "value=42# comment", want: Value{Type: "integer", Value: "42"}},
		{name: "hash in basic string", input: `value = "a#b" # comment`, want: Value{Type: "string", Value: "a#b"}},
		{name: "hash in literal string", input: `value = '#b' # comment`, want: Value{Type: "string", Value: "#b"}},
		{name: "tabs are whitespace", input: "\tvalue\t=\ttrue\t# comment", want: Value{Type: "bool", Value: "true"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.input)
			if err != nil {
				t.Fatalf("Decode(%q) returned error: %v", tt.input, err)
			}
			if got["value"] != tt.want {
				t.Fatalf("Decode(%q) = %v, want %v", tt.input, got["value"], tt.want)
			}
		})
	}
}

func TestDecodeRejectsLexicalBoundaryViolations(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "adjacent statements", input: "a = 1 b = 2"},
		{name: "lone carriage return", input: "a = 1\rb = 2"},
		{name: "null", input: "# bad" + string([]byte{0})},
		{name: "vertical tab", input: "a = 1" + string([]byte{0x0b})},
		{name: "form feed", input: "# bad" + string([]byte{0x0c})},
		{name: "delete", input: "# bad" + string([]byte{0x7f})},
		{name: "invalid UTF-8", input: string([]byte{'a', '=', ' ', 0xff})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := Decode(tt.input); err == nil {
				t.Fatalf("Decode(%q) accepted invalid lexical boundary", tt.input)
			}
		})
	}
}

func TestDecodeKeys(t *testing.T) {
	input := `bare = 1
"quoted key" = 2
outer . "inner.key" = 3
"equals=key"=4
`
	got, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode(keys) returned error: %v", err)
	}
	if got["bare"] != (Value{Type: "integer", Value: "1"}) {
		t.Fatalf("bare key = %v", got["bare"])
	}
	if got["quoted key"] != (Value{Type: "integer", Value: "2"}) {
		t.Fatalf("quoted key = %v", got["quoted key"])
	}
	outer, ok := got["outer"].(map[string]any)
	if !ok {
		t.Fatalf("dotted key parent = %T, want map[string]any", got["outer"])
	}
	if outer["inner.key"] != (Value{Type: "integer", Value: "3"}) {
		t.Fatalf("dotted quoted key = %v", outer["inner.key"])
	}
	if got["equals=key"] != (Value{Type: "integer", Value: "4"}) {
		t.Fatalf("quoted key containing equals = %v", got["equals=key"])
	}
}

func TestDecodeRejectsKeyRedefinitionAndScalarExtension(t *testing.T) {
	tests := []string{
		"name = \"Tom\"\nname = \"Pradyun\"\n",
		"fruit.apple = 1\nfruit.apple.smooth = true\n",
		"fruit = 1\nfruit.apple = true\n",
		"a.b = 1\na . b = 2\n",
		"\"a\"tail = 1\n",
	}
	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			if _, err := Decode(input); err == nil {
				t.Fatalf("Decode(%q) accepted invalid key definition", input)
			}
		})
	}
}

func TestDecodeTablesAndArraysOfTables(t *testing.T) {
	input := `[fruit]
name = "apple"
[fruit.physical]
color = "red"
[[fruit.variety]]
name = "red delicious"
[[fruit.variety]]
name = "granny smith"
`
	got, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode(tables) returned error: %v", err)
	}
	fruit := got["fruit"].(map[string]any)
	if fruit["name"] != (Value{Type: "string", Value: "apple"}) {
		t.Fatalf("fruit.name = %v", fruit["name"])
	}
	physical := fruit["physical"].(map[string]any)
	if physical["color"] != (Value{Type: "string", Value: "red"}) {
		t.Fatalf("fruit.physical.color = %v", physical["color"])
	}
	varieties := fruit["variety"].([]any)
	if len(varieties) != 2 {
		t.Fatalf("fruit.variety length = %d, want 2", len(varieties))
	}
	if varieties[1].(map[string]any)["name"] != (Value{Type: "string", Value: "granny smith"}) {
		t.Fatalf("second variety = %v", varieties[1])
	}
}

func TestDecodeNestedArraysOfTablesTargetMostRecentElement(t *testing.T) {
	input := `[[fruits]]
name = "apple"
[[fruits.varieties]]
name = "red"
[[fruits]]
name = "orange"
[[fruits.varieties]]
name = "navel"
`
	got, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode(nested arrays) returned error: %v", err)
	}
	fruits := got["fruits"].([]any)
	if len(fruits) != 2 {
		t.Fatalf("fruits length = %d, want 2", len(fruits))
	}
	first := fruits[0].(map[string]any)
	second := fruits[1].(map[string]any)
	if len(first["varieties"].([]any)) != 1 || len(second["varieties"].([]any)) != 1 {
		t.Fatalf("nested varieties targeted incorrectly: %v", fruits)
	}
}

func TestDecodeRejectsTableOrderingConflicts(t *testing.T) {
	tests := []string{
		"[fruit]\napple = 1\n[fruit]\norange = 2\n",
		"fruits = []\n[[fruits]]\nname = \"apple\"\n",
		"[fruit]\nname = \"apple\"\n[[fruit]]\nname = \"orange\"\n",
		"value = 1\n[value.child]\nname = \"bad\"\n",
	}
	for _, input := range tests {
		if _, err := Decode(input); err == nil {
			t.Errorf("Decode(%q) accepted table conflict", input)
		}
	}
}

func TestDecodeArraysAndInlineTables(t *testing.T) {
	input := `values = [
		1, # first
		[true, {name = "n"}],
	]
product = {name = "Nail", physical.color = "gray"}
`
	got, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode(arrays and inline tables) returned error: %v", err)
	}
	values := got["values"].([]any)
	if len(values) != 2 {
		t.Fatalf("values length = %d, want 2", len(values))
	}
	nested := values[1].([]any)
	if nested[0] != (Value{Type: "bool", Value: "true"}) {
		t.Fatalf("nested array value = %v", nested[0])
	}
	inline := nested[1].(map[string]any)
	if inline["name"] != (Value{Type: "string", Value: "n"}) {
		t.Fatalf("nested inline value = %v", inline)
	}
	product := got["product"].(map[string]any)
	physical := product["physical"].(map[string]any)
	if physical["color"] != (Value{Type: "string", Value: "gray"}) {
		t.Fatalf("dotted inline value = %v", physical["color"])
	}
}

func TestDecodeRejectsInlineTableExtensionsAndTrailingCommas(t *testing.T) {
	tests := []string{
		`product = {name = "Nail",}`,
		"[product]\ntype = { name = \"Nail\" }\ntype.edible = false\n",
	}
	for _, input := range tests {
		if _, err := Decode(input); err == nil {
			t.Errorf("Decode(%q) accepted invalid inline table", input)
		}
	}
}
