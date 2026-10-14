package main

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// buildDecoderBinary compiles the command into a fresh temporary directory so
// no test leaves a binary behind in the module.
func buildDecoderBinary(t *testing.T) string {
	t.Helper()
	bin := filepath.Join(t.TempDir(), "toml-decoder")
	build := exec.Command("go", "build", "-o", bin, "./cmd/toml-decoder")
	build.Dir = moduleRoot
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("go build failed: %v\n%s", err, out)
	}
	return bin
}

// runDecoderBinary feeds stdin to bin and returns its stdout and exit error.
func runDecoderBinary(t *testing.T, bin, stdin string) ([]byte, error) {
	t.Helper()
	cmd := exec.Command(bin)
	cmd.Stdin = strings.NewReader(stdin)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	return stdout.Bytes(), err
}

func TestBinaryDecodesValueBeforeTrailingComment(t *testing.T) {
	bin := buildDecoderBinary(t)
	stdout, err := runDecoderBinary(t, bin, "n = 1 # note\n")
	if err != nil {
		t.Fatalf("binary exit = %v, want 0", err)
	}
	var decoded map[string]map[string]string
	if err := json.Unmarshal(stdout, &decoded); err != nil {
		t.Fatalf("stdout %q is not JSON: %v", stdout, err)
	}
	want := map[string]map[string]string{"n": {"type": "integer", "value": "1"}}
	if !reflect.DeepEqual(decoded, want) {
		t.Errorf("stdout decodes to %v, want %v", decoded, want)
	}
}

func TestBinaryRejectsControlCharacterInComment(t *testing.T) {
	bin := buildDecoderBinary(t)
	stdout, err := runDecoderBinary(t, bin, "# \x01\nn = 1\n")
	assertExitNonZero(t, err)
	if len(stdout) != 0 {
		t.Errorf("stdout = %q, want empty", stdout)
	}
}

func TestBinaryRejectsInvalidUTF8(t *testing.T) {
	bin := buildDecoderBinary(t)
	stdout, err := runDecoderBinary(t, bin, "\xff")
	assertExitNonZero(t, err)
	if len(stdout) != 0 {
		t.Errorf("stdout = %q, want empty", stdout)
	}
}

func assertExitNonZero(t *testing.T, err error) {
	t.Helper()
	exitErr, ok := err.(*exec.ExitError)
	if !ok || exitErr.ExitCode() == 0 {
		t.Fatalf("binary exit = %v, want non-zero exit", err)
	}
}
