package main

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
)

func TestRunValidDocumentEmitsTaggedJSONOnly(t *testing.T) {
	var output, diagnostics bytes.Buffer

	err := run(strings.NewReader("answer = 42\n"), &output, &diagnostics)

	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	if got, want := output.String(), "{\"answer\":{\"type\":\"integer\",\"value\":\"42\"}}\n"; got != want {
		t.Fatalf("output = %q, want %q", got, want)
	}
	if diagnostics.Len() != 0 {
		t.Fatalf("diagnostics = %q, want empty", diagnostics.String())
	}
}

func TestRunMalformedDocumentReportsErrorWithoutJSON(t *testing.T) {
	var output, diagnostics bytes.Buffer

	err := run(strings.NewReader("answer = [\n"), &output, &diagnostics)

	if err == nil {
		t.Fatal("run returned success for malformed TOML")
	}
	if output.Len() != 0 {
		t.Fatalf("output = %q, want empty", output.String())
	}
	if diagnostics.Len() == 0 {
		t.Fatal("diagnostics is empty, want an error diagnostic")
	}
}

func TestRunInputReadFailureReportsErrorWithoutJSON(t *testing.T) {
	var output, diagnostics bytes.Buffer
	readErr := errors.New("input unavailable")

	err := run(failingReader{err: readErr}, &output, &diagnostics)

	if !errors.Is(err, readErr) {
		t.Fatalf("run error = %v, want wrapped input error", err)
	}
	if output.Len() != 0 {
		t.Fatalf("output = %q, want empty", output.String())
	}
	if diagnostics.Len() == 0 {
		t.Fatal("diagnostics is empty, want an error diagnostic")
	}
}

type failingReader struct {
	err error
}

func (r failingReader) Read([]byte) (int, error) {
	return 0, r.err
}

var _ io.Reader = failingReader{}
