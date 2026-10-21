# FEATURE: Decoder Contract

| Field       | Value |
|-------------|-------|
| Version     | 20261021 V1 |
| Description | Defines the stdin, tagged-JSON, diagnostic, and exit-status contract for the TOML decoder. |
| Depends On  | ARCHITECTURE.md |
| Provides    | stdin TOML decoder, tagged JSON stdout, stderr diagnostics |
| Consumes    | Go module, cmd/toml-decoder |

## Purpose

Provide a deterministic command-line filter for TOML input. The command delegates parsing to `internal/toml` and serializes the resulting tagged representation as JSON.

## Contract

- Tables encode as JSON objects.
- Arrays encode as JSON arrays.
- Every TOML scalar value encodes as an object with `type` and string-valued `value` fields.
- Supported type tags are `string`, `integer`, `float`, `bool`, `datetime`, `datetime-local`, `date-local`, and `time-local`.
- Valid input emits JSON only on stdout and exits zero.
- Invalid input emits a diagnostic on stderr, emits no decoded output, and exits non-zero.
- The command takes no arguments, reads no configuration files, and produces no side effects.

## Programmatic Acceptance

Requires: executable=sh; scope=test
Requires: executable=go; scope=test

=== AC decoder-contract-gate ===
Intent: The executable satisfies the supplied valid-input and invalid-input stream contract.

import subprocess

result = subprocess.run(
    ["sh", "sources/stage_contract.sh", "decoder"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC decoder-contract-gate ===

=== AC decoder-executable ===
Intent: The decoder executable is produced at the declared command boundary.

from pathlib import Path

assert Path("toml-decoder").is_file()
=== END AC decoder-executable ===

## User Acceptance

- None.

## Guardrails

- Never emit decoded output for invalid TOML.
- Never write diagnostics to stdout.
- Never weaken invalid-input rejection to improve acceptance.
- Read TOML exclusively from stdin and emit JSON exclusively to stdout.
