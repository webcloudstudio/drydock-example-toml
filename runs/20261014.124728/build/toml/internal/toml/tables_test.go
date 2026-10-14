package toml

import (
	"errors"
	"reflect"
	"testing"
)

func TestDecodeBuildsStandardTables(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[string]any
	}{
		{
			name:  "header holds its keys",
			input: "[owner]\nname = \"Regina\"\n",
			want: map[string]any{
				"owner": map[string]any{"name": stringLeaf("Regina")},
			},
		},
		{
			name:  "empty header is an empty table",
			input: "[t]\n",
			want:  map[string]any{"t": map[string]any{}},
		},
		{
			name:  "dotted header creates implicit super-tables",
			input: "[a.b.c]\nz = 1\n",
			want: map[string]any{
				"a": map[string]any{
					"b": map[string]any{"c": map[string]any{"z": intLeaf("1")}},
				},
			},
		},
		{
			name:  "root keys stay at the root beside a table",
			input: "top = \"t\"\n[t]\nb = \"v\"\n",
			want: map[string]any{
				"top": stringLeaf("t"),
				"t":   map[string]any{"b": stringLeaf("v")},
			},
		},
		{
			name:  "spaces and quoted parts in a header name",
			input: "[ a . \"b c\" ]\nk = \"v\"\n",
			want: map[string]any{
				"a": map[string]any{"b c": map[string]any{"k": stringLeaf("v")}},
			},
		},
		{
			name:  "empty quoted name is a valid key",
			input: "[\"\"]\nk = \"v\"\n",
			want:  map[string]any{"": map[string]any{"k": stringLeaf("v")}},
		},
		{
			name:  "sibling headers share their parent",
			input: "[a.b]\nx = \"1\"\n[a.c]\ny = \"2\"\n",
			want: map[string]any{
				"a": map[string]any{
					"b": map[string]any{"x": stringLeaf("1")},
					"c": map[string]any{"y": stringLeaf("2")},
				},
			},
		},
		{
			name:  "implicit super-table may be defined later",
			input: "[x.y.z.w]\nk = 1\n[x]\nv = 2\n",
			want: map[string]any{
				"x": map[string]any{
					"v": intLeaf("2"),
					"y": map[string]any{"z": map[string]any{"w": map[string]any{"k": intLeaf("1")}}},
				},
			},
		},
		{
			name:  "header may add a sub-table under a dotted-key table",
			input: "[fruit]\napple.color = \"red\"\n[fruit.apple.texture]\nsmooth = true\n",
			want: map[string]any{
				"fruit": map[string]any{
					"apple": map[string]any{
						"color":   stringLeaf("red"),
						"texture": map[string]any{"smooth": map[string]any{"type": "bool", "value": "true"}},
					},
				},
			},
		},
		{
			name:  "comments and CRLF line endings between headers",
			input: "# first\r\n[a]\r\nx = \"1\" # note\r\n\r\n[b]\r\n",
			want: map[string]any{
				"a": map[string]any{"x": stringLeaf("1")},
				"b": map[string]any{},
			},
		},
		{
			name:  "final header without a newline",
			input: "[a]\nx = \"1\"",
			want:  map[string]any{"a": map[string]any{"x": stringLeaf("1")}},
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

func TestDecodeBuildsArraysOfTables(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[string]any
	}{
		{
			name:  "one element",
			input: "[[p]]\nsku = 1\n",
			want: map[string]any{
				"p": []any{map[string]any{"sku": intLeaf("1")}},
			},
		},
		{
			name:  "two headers make two elements in order",
			input: "[[products]]\nsku = 738594937\n\n[[products]]\nsku = 284758393\n",
			want: map[string]any{
				"products": []any{
					map[string]any{"sku": intLeaf("738594937")},
					map[string]any{"sku": intLeaf("284758393")},
				},
			},
		},
		{
			name:  "empty elements are still elements",
			input: "[[a]]\n[[a]]\n",
			want:  map[string]any{"a": []any{map[string]any{}, map[string]any{}}},
		},
		{
			name:  "sub-table attaches to the most recent element",
			input: "[[fruit]]\nname = \"apple\"\n[fruit.physical]\ncolor = \"red\"\n[[fruit]]\nname = \"banana\"\n[fruit.physical]\ncolor = \"yellow\"\n",
			want: map[string]any{
				"fruit": []any{
					map[string]any{
						"name":     stringLeaf("apple"),
						"physical": map[string]any{"color": stringLeaf("red")},
					},
					map[string]any{
						"name":     stringLeaf("banana"),
						"physical": map[string]any{"color": stringLeaf("yellow")},
					},
				},
			},
		},
		{
			name:  "implicit super-table of an array element",
			input: "[[albums.songs]]\nname = \"Glory Days\"\n",
			want: map[string]any{
				"albums": map[string]any{
					"songs": []any{map[string]any{"name": stringLeaf("Glory Days")}},
				},
			},
		},
		{
			name:  "array elements may nest arrays of their own",
			input: "[[a]]\n[[a.b]]\nx = \"1\"\n[[a.b]]\nx = \"2\"\n[[a]]\n",
			want: map[string]any{
				"a": []any{
					map[string]any{"b": []any{
						map[string]any{"x": stringLeaf("1")},
						map[string]any{"x": stringLeaf("2")},
					}},
					map[string]any{},
				},
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

func TestDecodeRejectsTableRedefinition(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "same header twice", input: "[fruit]\napple = 1\n\n[fruit]\norange = 2\n"},
		{name: "explicit table after its sub-table", input: "[a.b]\n[a]\n[a]\n"},
		{name: "header after its own dotted-key definition", input: "[t1]\nt2.t3.v = 0\n[t1.t2]\n"},
		{name: "header below a dotted-key table it defined", input: "[t1]\nt2.t3.v = 0\n[t1.t2.t3]\n"},
		{name: "header after a dotted-key table at the root", input: "a.b = 1\n[a]\n"},
		{name: "dotted key extends a header table", input: "[a.b.c]\nz = 9\n[a]\nb.c.t = \"x\"\n"},
		{name: "dotted key extends an implicit table", input: "[x.y.z]\n[x]\ny.w = 1\n"},
		{name: "dotted key extends a table under a header", input: "[a.b.c.d]\nz = 9\n[a]\nb.c.d.k.t = \"x\"\n"},
		{name: "same header twice under an array element", input: "[[a]]\n[a.b]\n[a.b]\n"},
		{name: "header redefines a key with a value", input: "[a]\nb = 1\n[a.b]\nc = 2\n"},
		{name: "header walks through a value", input: "a = \"x\"\n[a.b]\n"},
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

func TestDecodeRejectsArrayOfTablesConflicts(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "array header after a standard table", input: "[a]\n[[a]]\n"},
		{name: "standard table header after an array", input: "[[a]]\n[a]\n"},
		{name: "array header after an implicit super-table", input: "[[a.b]]\n[[a]]\n"},
		{name: "array header over a value", input: "a = 1\n[[a]]\n"},
		{name: "dotted key extends an array of tables", input: "[[a.b]]\n[a]\nb.y = 2\n"},
		{name: "array header walks through a value", input: "a = \"x\"\n[[a.b]]\n"},
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

func TestDecodeRejectsMalformedHeaders(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "empty header", input: "[]\n"},
		{name: "empty array header", input: "[[]]\n"},
		{name: "missing close bracket", input: "[a\n"},
		{name: "missing close bracket after dotted name", input: "[a.b\n"},
		{name: "missing second close bracket", input: "[[a]\n"},
		{name: "space between close brackets", input: "[[a] ]\n"},
		{name: "space between open brackets", input: "[ [a]]\n"},
		{name: "text after close bracket", input: "[a]b]\n"},
		{name: "text after header", input: "[error] this shouldn't be here\n"},
		{name: "trailing dot", input: "[a.]\n"},
		{name: "leading dot", input: "[.a]\n"},
		{name: "space inside a bare name", input: "[invalid key]\n"},
		{name: "pound sign in a bare name", input: "[key#group]\n"},
		{name: "header spans a line break", input: "[tbl\n]\n"},
		{name: "invalid bare character", input: "[!]\nk = 1\n"},
		{name: "header at end of input", input: "[a"},
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

// intLeaf is the tagged model for an integer value, shared by the table tests.
func intLeaf(v string) map[string]any {
	return map[string]any{"type": "integer", "value": v}
}
