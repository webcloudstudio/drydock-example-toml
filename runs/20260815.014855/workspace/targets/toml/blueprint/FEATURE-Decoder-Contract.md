# FEATURE: Decoder Contract

| Field       | Value |
|-------------|-------|
| Version     | 20260815 V1 |
| Description | Defines the stdin, tagged-JSON, diagnostic, and exit-status contract for the decoder. |
| Depends On  | ARCHITECTURE.md |
| Provides    | stdin TOML decoder, tagged JSON output, invalid-input diagnostics |
| Consumes    | Go module, cmd/toml-decoder build entry point, internal/toml parser boundary |

## Workflow

The executable reads exactly one TOML document from standard input. For valid input it parses through `internal/toml` and emits one JSON document on stdout with no diagnostic output. For invalid input it emits a diagnostic on stderr, emits no JSON document on stdout, and exits non-zero. It accepts no arguments and performs no side effects.

## Output Contract

- Tables become JSON objects.
- Arrays become JSON arrays.
- Scalar values become `{"type": "<TOML_TYPE>", "value": "<TOML_VALUE>"}`.
- Scalar `value` fields are JSON strings, including integers, floats, and booleans.
- Offset datetimes use RFC 3339 representation; local date/time values use their corresponding local representations.

## Programmatic Acceptance

=== AC decoder-valid-contract ===
Intent: A valid TOML document produces tagged JSON on stdout without stderr diagnostics.

import json
import subprocess

source = "answer = 42\n"
result = subprocess.run(
    ["./toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
decoded = json.loads(result.stdout)
expected_key = "answer"
expected_type = "integer"
expected_value = "42"
assert decoded[expected_key]["type"] == expected_type
assert decoded[expected_key]["value"] == expected_value
assert result.stderr == ""
=== END AC decoder-valid-contract ===

=== AC decoder-invalid-contract ===
Intent: Malformed TOML produces a non-zero exit, stderr diagnostics, and no stdout document.

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
Intent: The supplied decoder contract gate passes for valid and malformed stdin documents.

import os
import subprocess

result = subprocess.run(
    ["sh", "sources/stage_contract.sh", "decoder"],
    capture_output=True,
    text=True,
    env={**os.environ},
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC decoder-stage-contract ===

## User Acceptance

- None.

## Guardrails

- Emit no diagnostics on stdout for valid input.
- Reject malformed TOML rather than producing partial JSON.
- Do not inspect command-line arguments or require configuration.
