package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/owner/toml-decoder/internal/toml"
)

const moduleRoot = "../.."

func TestGoModDeclaresModuleWithoutRequire(t *testing.T) {
	data, err := os.ReadFile(filepath.Join(moduleRoot, "go.mod"))
	if err != nil {
		t.Fatalf("read go.mod: %v", err)
	}
	text := string(data)
	if strings.Contains(text, "require") {
		t.Errorf("go.mod declares a require directive:\n%s", text)
	}
	if !strings.Contains(text, "module github.com/owner/toml-decoder") {
		t.Errorf("go.mod does not declare the module path:\n%s", text)
	}
}

func TestEntryPointIsPackageMain(t *testing.T) {
	data, err := os.ReadFile("main.go")
	if err != nil {
		t.Fatalf("read main.go: %v", err)
	}
	if !strings.Contains(string(data), "package main") {
		t.Errorf("main.go is not a package main")
	}
	info, err := os.Stat(filepath.Join(moduleRoot, "internal", "toml"))
	if err != nil || !info.IsDir() {
		t.Errorf("internal/toml is not a directory: %v", err)
	}
}

func TestDecoderCommandBuildsWithStandardLibrary(t *testing.T) {
	bin := filepath.Join(t.TempDir(), "toml-decoder")
	cmd := exec.Command("go", "build", "-o", bin, "./cmd/toml-decoder")
	cmd.Dir = moduleRoot
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("go build failed: %v\n%s", err, out)
	}
	if _, err := os.Stat(bin); err != nil {
		t.Fatalf("built binary missing: %v", err)
	}
}

func TestRunWritesTaggedJSONForString(t *testing.T) {
	var out bytes.Buffer
	if err := run(strings.NewReader("s = \"text\"\n"), &out); err != nil {
		t.Fatalf("run error = %v, want nil", err)
	}
	var decoded map[string]map[string]string
	if err := json.Unmarshal(out.Bytes(), &decoded); err != nil {
		t.Fatalf("run output %q is not JSON: %v", out.String(), err)
	}
	want := map[string]map[string]string{"s": {"type": "string", "value": "text"}}
	if !reflect.DeepEqual(decoded, want) {
		t.Errorf("run output = %v, want %v", decoded, want)
	}
}

func TestRunReportsSyntaxErrorWithEmptyStdout(t *testing.T) {
	var out bytes.Buffer
	err := run(strings.NewReader("bad = \"\\q\"\n"), &out)
	if !errors.Is(err, toml.ErrSyntax) {
		t.Fatalf("run error = %v, want wrapping %v", err, toml.ErrSyntax)
	}
	if out.Len() != 0 {
		t.Errorf("run wrote %q to stdout on failure, want nothing", out.String())
	}
}

func TestBinaryExitsNonZeroWithEmptyStdout(t *testing.T) {
	bin := filepath.Join(t.TempDir(), "toml-decoder")
	build := exec.Command("go", "build", "-o", bin, "./cmd/toml-decoder")
	build.Dir = moduleRoot
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("go build failed: %v\n%s", err, out)
	}

	cmd := exec.Command(bin)
	cmd.Stdin = strings.NewReader("bad = \"\\q\"\n")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	var exitErr *exec.ExitError
	if !errors.As(err, &exitErr) || exitErr.ExitCode() == 0 {
		t.Fatalf("binary exit = %v, want non-zero exit", err)
	}
	if stdout.Len() != 0 {
		t.Errorf("stdout = %q, want empty", stdout.String())
	}
	if stderr.Len() == 0 {
		t.Errorf("stderr is empty, want a diagnostic")
	}
}
