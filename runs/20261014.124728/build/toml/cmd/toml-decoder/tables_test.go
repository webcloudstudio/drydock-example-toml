package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBinaryDecodesStandardTableAsNestedObject(t *testing.T) {
	bin := buildDecoderBinary(t)
	stdout, err := runDecoderBinary(t, bin, "[owner]\nname = \"Regina\"\n")
	if err != nil {
		t.Fatalf("binary exit = %v, want 0", err)
	}
	var decoded map[string]map[string]map[string]string
	if err := json.Unmarshal(stdout, &decoded); err != nil {
		t.Fatalf("stdout %q is not JSON: %v", stdout, err)
	}
	want := map[string]map[string]map[string]string{
		"owner": {"name": {"type": "string", "value": "Regina"}},
	}
	if !reflect.DeepEqual(decoded, want) {
		t.Errorf("decoded = %v, want %v", decoded, want)
	}
}

func TestBinaryDecodesArrayOfTablesAsJSONArray(t *testing.T) {
	bin := buildDecoderBinary(t)
	stdout, err := runDecoderBinary(t, bin, "[[products]]\nsku = 738594937\n\n[[products]]\nsku = 284758393\n")
	if err != nil {
		t.Fatalf("binary exit = %v, want 0", err)
	}
	var decoded map[string][]map[string]map[string]string
	if err := json.Unmarshal(stdout, &decoded); err != nil {
		t.Fatalf("stdout %q is not JSON: %v", stdout, err)
	}
	want := map[string][]map[string]map[string]string{
		"products": {
			{"sku": {"type": "integer", "value": "738594937"}},
			{"sku": {"type": "integer", "value": "284758393"}},
		},
	}
	if !reflect.DeepEqual(decoded, want) {
		t.Errorf("decoded = %v, want %v", decoded, want)
	}
}

func TestBinaryRejectsRedefinedTableWithEmptyStdout(t *testing.T) {
	bin := buildDecoderBinary(t)
	stdout, err := runDecoderBinary(t, bin, "[fruit]\napple = 1\n\n[fruit]\norange = 2\n")
	if err == nil {
		t.Fatalf("binary exit = 0, want non-zero for redefined table")
	}
	if len(stdout) != 0 {
		t.Errorf("stdout = %q, want empty", stdout)
	}
}
