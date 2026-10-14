package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBinaryDecodesBareKey(t *testing.T) {
	bin := buildDecoderBinary(t)
	stdout, err := runDecoderBinary(t, bin, "bare-key_1 = 5\n")
	if err != nil {
		t.Fatalf("binary exit = %v, want 0", err)
	}
	var decoded map[string]map[string]string
	if err := json.Unmarshal(stdout, &decoded); err != nil {
		t.Fatalf("stdout %q is not JSON: %v", stdout, err)
	}
	want := map[string]map[string]string{"bare-key_1": {"type": "integer", "value": "5"}}
	if !reflect.DeepEqual(decoded, want) {
		t.Errorf("decoded = %v, want %v", decoded, want)
	}
}

func TestBinaryDecodesDottedKeyAsNestedTable(t *testing.T) {
	bin := buildDecoderBinary(t)
	stdout, err := runDecoderBinary(t, bin, "physical.color = \"orange\"\n")
	if err != nil {
		t.Fatalf("binary exit = %v, want 0", err)
	}
	var decoded map[string]map[string]map[string]string
	if err := json.Unmarshal(stdout, &decoded); err != nil {
		t.Fatalf("stdout %q is not JSON: %v", stdout, err)
	}
	want := map[string]map[string]map[string]string{
		"physical": {"color": {"type": "string", "value": "orange"}},
	}
	if !reflect.DeepEqual(decoded, want) {
		t.Errorf("decoded = %v, want %v", decoded, want)
	}
}

func TestBinaryRejectsRedefinedKeyWithEmptyStdout(t *testing.T) {
	bin := buildDecoderBinary(t)
	stdout, err := runDecoderBinary(t, bin, "name = \"Tom\"\nname = \"Pradyun\"\n")
	if err == nil {
		t.Fatalf("binary exit = 0, want non-zero for redefined key")
	}
	if len(stdout) != 0 {
		t.Errorf("stdout = %q, want empty", stdout)
	}
}
