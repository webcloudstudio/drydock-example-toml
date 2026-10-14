package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBinaryDecodesLocalDateAsDateLocal(t *testing.T) {
	bin := buildDecoderBinary(t)
	stdout, err := runDecoderBinary(t, bin, "d = 1979-05-27\n")
	if err != nil {
		t.Fatalf("binary exit = %v, want 0", err)
	}
	var decoded map[string]map[string]string
	if err := json.Unmarshal(stdout, &decoded); err != nil {
		t.Fatalf("stdout %q is not JSON: %v", stdout, err)
	}
	want := map[string]map[string]string{"d": {"type": "date-local", "value": "1979-05-27"}}
	if !reflect.DeepEqual(decoded, want) {
		t.Errorf("decoded = %v, want %v", decoded, want)
	}
}

func TestBinaryDecodesLocalTimeAsTimeLocal(t *testing.T) {
	bin := buildDecoderBinary(t)
	stdout, err := runDecoderBinary(t, bin, "t = 07:32:00\n")
	if err != nil {
		t.Fatalf("binary exit = %v, want 0", err)
	}
	var decoded map[string]map[string]string
	if err := json.Unmarshal(stdout, &decoded); err != nil {
		t.Fatalf("stdout %q is not JSON: %v", stdout, err)
	}
	want := map[string]map[string]string{"t": {"type": "time-local", "value": "07:32:00"}}
	if !reflect.DeepEqual(decoded, want) {
		t.Errorf("decoded = %v, want %v", decoded, want)
	}
}

func TestBinaryRejectsInvalidDateWithEmptyStdout(t *testing.T) {
	bin := buildDecoderBinary(t)
	stdout, err := runDecoderBinary(t, bin, "d = 1979-13-01\n")
	assertExitNonZero(t, err)
	if len(stdout) != 0 {
		t.Errorf("stdout = %q, want empty", stdout)
	}
}
