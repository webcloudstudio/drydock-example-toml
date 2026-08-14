package main

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"strings"
	"testing"
)

func TestDecoderValidStdinProducesJSON(t *testing.T) {
	cmd := exec.Command("go", "run", ".")
	cmd.Stdin = strings.NewReader("answer = 42\n")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("decoder failed: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(out, &decoded); err != nil {
		t.Fatalf("stdout is not JSON: %v", err)
	}
	if decoded["answer"].(map[string]any)["value"] != "42" {
		t.Fatalf("decoded JSON = %#v", decoded)
	}
}

func TestDecoderRejectsArguments(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "--help")
	if err := cmd.Run(); err == nil {
		t.Fatal("decoder accepted an argument")
	}
}

func TestDecoderSuccessfulRunLeavesStderrEmpty(t *testing.T) {
	cmd := exec.Command("go", "run", ".")
	cmd.Stdin = strings.NewReader("enabled = true\n")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("decoder failed: %v", err)
	}
	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q", stderr.String())
	}
}

func TestDecoderInvalidInputUsesFailureBoundary(t *testing.T) {
	cmd := exec.Command("go", "run", ".")
	cmd.Stdin = strings.NewReader("broken = [\n")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	stdout, err := cmd.Output()
	if err == nil {
		t.Fatal("decoder accepted malformed TOML")
	}
	if stderr.Len() == 0 {
		t.Fatal("decoder produced no diagnostic on stderr")
	}
	if len(stdout) != 0 {
		t.Fatalf("decoder produced successful stdout: %q", stdout)
	}
}
