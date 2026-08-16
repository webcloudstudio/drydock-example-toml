package main

import (
	"encoding/json"
	"os/exec"
	"reflect"
	"testing"
)

func TestDecoderContract(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantStatus int
		wantJSON   map[string]any
	}{
		{name: "valid", input: "answer = 42\n", wantJSON: map[string]any{"answer": map[string]any{"type": "integer", "value": "42"}}},
		{name: "empty", input: "", wantJSON: map[string]any{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", "run", ".")
			cmd.Stdin = stringsReader(tt.input)
			stdout, stderr, err := runCommand(cmd)
			if tt.wantStatus == 0 && err != nil {
				t.Fatalf("decoder failed: %v; stderr=%q", err, stderr)
			}
			var got map[string]any
			if err := json.Unmarshal(stdout, &got); err != nil {
				t.Fatalf("stdout is not JSON: %v", err)
			}
			if !reflect.DeepEqual(got, tt.wantJSON) {
				t.Fatalf("decoded JSON = %#v, want %#v", got, tt.wantJSON)
			}
			if len(stderr) != 0 {
				t.Fatalf("unexpected stderr: %q", stderr)
			}
		})
	}
}

func TestDecoderInvalidInput(t *testing.T) {
	cmd := exec.Command("go", "run", ".")
	cmd.Stdin = stringsReader("answer = [\n")
	stdout, stderr, err := runCommand(cmd)
	if err == nil {
		t.Fatal("invalid TOML exited successfully")
	}
	if len(stdout) != 0 || len(stderr) == 0 {
		t.Fatalf("invalid TOML streams: stdout=%q stderr=%q", stdout, stderr)
	}
}
