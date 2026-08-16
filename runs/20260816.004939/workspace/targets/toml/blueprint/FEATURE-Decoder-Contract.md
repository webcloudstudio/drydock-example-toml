# FEATURE: Decoder Contract

| Field       | Value |
|-------------|-------|
| Version     | 20260816 V1 |
| Description | Defines the stdin, tagged-JSON, diagnostic, and exit-status contract of the decoder command. |
| Depends On  | ARCHITECTURE.md |
| Provides    | stdin TOML filter, tagged JSON output, invalid-input exit contract |
| Consumes    | internal/toml.Decode |

## Workflow

The command reads the complete TOML document from stdin. On valid input it writes exactly one tagged JSON document to stdout and exits zero. On invalid input it writes a diagnostic to stderr, writes no successful JSON payload to stdout, and exits nonzero.

The command accepts no positional arguments or options and has no configuration or side effects.

## Tagged JSON Contract

Every scalar TOML value is encoded as:

```json
{"type":"<TOML_TYPE>","value":"<TOML_VALUE>"}
```

Integer, float, and boolean values retain their TOML spelling in the string-valued `value` field. Tables and arrays retain their object and array structure.

## Programmatic Acceptance

=== AC decoder-valid-contract ===
Intent: Valid stdin produces the required tagged JSON state and no diagnostic.
Requires: executable=sh; scope=test
Requires: executable=python3; scope=test

import json
import subprocess

result = subprocess.run(
    ["sh", "sources/stage_contract.sh", "decoder"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC decoder-valid-contract ===

=== AC decoder-invalid-contract ===
Intent: Malformed stdin is rejected without a successful stdout payload.
Requires: executable=sh; scope=test

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_contract.sh", "decoder"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC decoder-invalid-contract ===

=== AC decoder-no-arguments ===
Intent: The command boundary rejects unsupported command-line arguments.
Requires: executable=go; scope=test

import subprocess

subprocess.run(
    ["go", "build", "-o", "toml-decoder", "./cmd/toml-decoder"],
    check=True,
)
result = subprocess.run(
    ["./toml-decoder", "--unsupported"],
    input="answer = 42\n",
    capture_output=True,
    text=True,
)
assert result.returncode != 0
=== END AC decoder-no-arguments ===

## User Acceptance

- None.

## Guardrails

- Successful output is written only to stdout.
- Diagnostics are written only to stderr.
- Invalid TOML must never produce a successful payload or exit zero.
- The command must have no configuration, persistence, network, or filesystem side effects.
