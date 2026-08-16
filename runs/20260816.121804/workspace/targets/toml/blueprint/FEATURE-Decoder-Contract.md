# FEATURE: Decoder Contract

| Field       | Value |
|-------------|-------|
| Version     | 20260816 V1 |
| Description | Defines the stdin, tagged-JSON, diagnostic-stream, and exit-status contract for the TOML decoder. |
| Depends On  | ARCHITECTURE.md |
| Provides    | stdin TOML decoder, tagged JSON stdout, stderr diagnostics |
| Consumes    | internal/toml |

## Workflow

The executable reads the complete TOML document from standard input and delegates parsing to `internal/toml`. On success it writes exactly one tagged JSON document to standard output and exits zero. On invalid TOML it writes a diagnostic to standard error, writes no successful result to standard output, and exits non-zero. The executable accepts no required arguments and has no runtime configuration.

## Output Contract

- Tables are JSON objects.
- Arrays are JSON arrays.
- Every scalar is encoded as `{"type": "<TOML_TYPE>", "value": "<TOML_VALUE>"}`.
- Scalar `value` fields are always JSON strings.
- Integer values must remain lossless across the signed 64-bit range.
- Offset and local date/time values use the required TOML type tags and formatting.

## Programmatic Acceptance

=== AC decoder-valid-contract ===
Intent: A valid TOML document produces the required tagged JSON on stdout with no stderr diagnostic.

import json
import subprocess

source = "answer = 42\n"
expected = {"answer": {"type": "integer", "value": "42"}}
result = subprocess.run(
    ["./toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
assert result.stderr == ""
assert json.loads(result.stdout) == expected
=== END AC decoder-valid-contract ===

=== AC decoder-invalid-contract ===
Intent: Invalid TOML produces a non-zero exit, no successful stdout result, and a stderr diagnostic.

import subprocess

source = "answer = [\n"
result = subprocess.run(
    ["./toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode != 0
assert result.stdout == ""
assert result.stderr != ""
=== END AC decoder-invalid-contract ===

=== AC decoder-stage-contract ===
Intent: The supplied decoder contract gate passes in full.

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_contract.sh", "decoder"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC decoder-stage-contract ===

## User Acceptance

- None.

## Guardrails

- Successful JSON is written only to stdout.
- Diagnostics are written only to stderr.
- Invalid input must never exit successfully.
- The executable must not read configuration files or produce runtime side effects.
