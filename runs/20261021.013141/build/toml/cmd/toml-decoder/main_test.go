package main

import (
	"bytes"
	"testing"
)

func TestRunRejectsArguments(t *testing.T) {
	var out, errOut bytes.Buffer
	if got := run([]string{"--help"}, bytes.NewReader(nil), &out, &errOut); got == 0 {
		t.Fatal("run accepted an argument")
	}
	if out.Len() != 0 {
		t.Fatalf("argument error wrote stdout: %q", out.String())
	}
}

func TestRunInvalidInputHasNoDecodedOutput(t *testing.T) {
	var out, errOut bytes.Buffer
	if got := run(nil, bytes.NewBufferString("answer = [\n"), &out, &errOut); got == 0 {
		t.Fatal("run accepted malformed TOML")
	}
	if out.Len() != 0 {
		t.Fatalf("invalid input wrote stdout: %q", out.String())
	}
	if errOut.Len() == 0 {
		t.Fatal("invalid input produced no diagnostic")
	}
}

func TestRunEmitsTaggedJSONForValidInput(t *testing.T) {
	var out, errOut bytes.Buffer
	if got := run(nil, bytes.NewBufferString("answer = 42\n"), &out, &errOut); got != 0 {
		t.Fatalf("run returned %d: %s", got, errOut.String())
	}
	if out.String() != "{\"answer\":{\"type\":\"integer\",\"value\":\"42\"}}\n" {
		t.Fatalf("unexpected JSON output: %q", out.String())
	}
	if errOut.Len() != 0 {
		t.Fatalf("valid input wrote stderr: %q", errOut.String())
	}
}
