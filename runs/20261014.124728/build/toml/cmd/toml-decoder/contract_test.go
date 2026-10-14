package main

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"reflect"
	"strings"
	"testing"
)

// runDecoderWithArgs runs bin with args and stdin, returning stdout, stderr, and the exit error.
func runDecoderWithArgs(t *testing.T, bin, stdin string, args ...string) ([]byte, []byte, error) {
	t.Helper()
	cmd := exec.Command(bin, args...)
	cmd.Stdin = strings.NewReader(stdin)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.Bytes(), stderr.Bytes(), err
}

func TestCheckNoArgsAcceptsEmptyArgs(t *testing.T) {
	if err := checkNoArgs(nil); err != nil {
		t.Errorf("checkNoArgs(nil) = %v, want nil", err)
	}
}

func TestCheckNoArgsRejectsAnyArgument(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{name: "one flag", args: []string{"-v"}},
		{name: "one path", args: []string{"config.toml"}},
		{name: "many", args: []string{"-x", "y", "z"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkNoArgs(tt.args); err == nil {
				t.Errorf("checkNoArgs(%q) = nil, want an error", tt.args)
			}
		})
	}
}

func TestBinaryRejectsArgumentsWithEmptyStdout(t *testing.T) {
	bin := buildDecoderBinary(t)
	stdout, stderr, err := runDecoderWithArgs(t, bin, "n = 1\n", "config.toml")
	assertExitNonZero(t, err)
	if len(stdout) != 0 {
		t.Errorf("stdout = %q, want empty", stdout)
	}
	if len(stderr) == 0 {
		t.Errorf("stderr is empty, want a diagnostic")
	}
}

func TestBinaryEmptyDocumentWritesEmptyObject(t *testing.T) {
	bin := buildDecoderBinary(t)
	stdout, stderr, err := runDecoderWithArgs(t, bin, "")
	if err != nil {
		t.Fatalf("binary exit = %v, want 0", err)
	}
	if len(stderr) != 0 {
		t.Errorf("stderr = %q, want empty on success", stderr)
	}
	var decoded map[string]any
	if err := json.Unmarshal(stdout, &decoded); err != nil {
		t.Fatalf("stdout %q is not JSON: %v", stdout, err)
	}
	if !reflect.DeepEqual(decoded, map[string]any{}) {
		t.Errorf("stdout decodes to %v, want empty object", decoded)
	}
}

func TestBinaryEncodesTableAsTaggedObject(t *testing.T) {
	bin := buildDecoderBinary(t)
	stdout, stderr, err := runDecoderWithArgs(t, bin, "[tbl]\na = 7\n")
	if err != nil {
		t.Fatalf("binary exit = %v, want 0; stderr %q", err, stderr)
	}
	var decoded map[string]any
	if err := json.Unmarshal(stdout, &decoded); err != nil {
		t.Fatalf("stdout %q is not JSON: %v", stdout, err)
	}
	want := map[string]any{
		"tbl": map[string]any{"a": map[string]any{"type": "integer", "value": "7"}},
	}
	if !reflect.DeepEqual(decoded, want) {
		t.Errorf("stdout decodes to %v, want %v", decoded, want)
	}
}

func TestBinaryEncodesArrayAsJSONArrayOfTaggedValues(t *testing.T) {
	bin := buildDecoderBinary(t)
	stdout, stderr, err := runDecoderWithArgs(t, bin, "a = [1, 2]\n")
	if err != nil {
		t.Fatalf("binary exit = %v, want 0; stderr %q", err, stderr)
	}
	var decoded map[string]any
	if err := json.Unmarshal(stdout, &decoded); err != nil {
		t.Fatalf("stdout %q is not JSON: %v", stdout, err)
	}
	want := map[string]any{
		"a": []any{
			map[string]any{"type": "integer", "value": "1"},
			map[string]any{"type": "integer", "value": "2"},
		},
	}
	if !reflect.DeepEqual(decoded, want) {
		t.Errorf("stdout decodes to %v, want %v", decoded, want)
	}
}

func TestBinaryEmptyArrayEncodesAsEmptyJSONArray(t *testing.T) {
	bin := buildDecoderBinary(t)
	stdout, stderr, err := runDecoderWithArgs(t, bin, "a = []\n")
	if err != nil {
		t.Fatalf("binary exit = %v, want 0; stderr %q", err, stderr)
	}
	var decoded map[string]any
	if err := json.Unmarshal(stdout, &decoded); err != nil {
		t.Fatalf("stdout %q is not JSON: %v", stdout, err)
	}
	want := map[string]any{"a": []any{}}
	if !reflect.DeepEqual(decoded, want) {
		t.Errorf("stdout decodes to %v, want %v", decoded, want)
	}
}
