package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBinaryDecodesArrayAsJSONArray(t *testing.T) {
	bin := buildDecoderBinary(t)
	stdout, err := runDecoderBinary(t, bin, "xs = [1, 2]\n")
	if err != nil {
		t.Fatalf("binary exit = %v, want 0", err)
	}
	var decoded map[string][]map[string]string
	if err := json.Unmarshal(stdout, &decoded); err != nil {
		t.Fatalf("stdout %q is not JSON: %v", stdout, err)
	}
	want := map[string][]map[string]string{
		"xs": {
			{"type": "integer", "value": "1"},
			{"type": "integer", "value": "2"},
		},
	}
	if !reflect.DeepEqual(decoded, want) {
		t.Errorf("decoded = %v, want %v", decoded, want)
	}
}

func TestBinaryDecodesInlineTableAsNestedObject(t *testing.T) {
	bin := buildDecoderBinary(t)
	stdout, err := runDecoderBinary(t, bin, "point = { x = 1 }\n")
	if err != nil {
		t.Fatalf("binary exit = %v, want 0", err)
	}
	var decoded map[string]map[string]map[string]string
	if err := json.Unmarshal(stdout, &decoded); err != nil {
		t.Fatalf("stdout %q is not JSON: %v", stdout, err)
	}
	want := map[string]map[string]map[string]string{
		"point": {"x": {"type": "integer", "value": "1"}},
	}
	if !reflect.DeepEqual(decoded, want) {
		t.Errorf("decoded = %v, want %v", decoded, want)
	}
}

func TestBinaryRejectsExtendedInlineTableWithEmptyStdout(t *testing.T) {
	bin := buildDecoderBinary(t)
	stdout, err := runDecoderBinary(t, bin, "t = { a = 1 }\n[t]\nb = 2\n")
	if err == nil {
		t.Fatalf("binary exit = 0, want non-zero for extended inline table")
	}
	if len(stdout) != 0 {
		t.Errorf("stdout = %q, want empty", stdout)
	}
}
