# FEATURE: Decoder Invalid Input

| Field       | Value |
|-------------|-------|
| Version     | 20260814 V1 |
| Description | Defines decoder behavior for malformed TOML input. |
| Depends On  | FEATURE-Decoder-Contract.md |
| Provides    | invalid-input diagnostic boundary |
| Consumes    | cmd/toml-decoder, parser error boundary |

## Workflow

When `internal/toml` rejects the input, `cmd/toml-decoder` writes a diagnostic to standard error and exits non-zero. It does not emit a successful JSON document on standard output. The diagnostic wording is not part of the contract.

## Failure Contract

Malformed syntax, invalid UTF-8, prohibited control characters, duplicate definitions, structural conflicts, and unsupported values all use the same process-level failure boundary: non-zero exit status, diagnostic on stderr, and no successful result on stdout.

## Guardrails

- Invalid TOML is rejected rather than accepted permissively.
- Failure diagnostics are written to stderr.
- The command does not claim success for invalid input.
- Diagnostic wording is not used as an acceptance oracle.

## Programmatic Acceptance

=== AC invalid-exit-status ===
Intent: Malformed TOML exits non-zero.
import subprocess

source = 'broken = [\n'
result = subprocess.run(
    ["./toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
assert result.returncode != 0
=== END AC invalid-exit-status ===

=== AC invalid-diagnostic-stream ===
Intent: Malformed TOML produces a non-empty diagnostic on stderr.
import subprocess

source = 'broken = [\n'
result = subprocess.run(
    ["./toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
assert result.returncode != 0
assert result.stderr != ""
=== END AC invalid-diagnostic-stream ===

=== AC invalid-no-success-result ===
Intent: Malformed TOML does not produce a successful JSON result on stdout.
import subprocess

source = 'broken = [\n'
result = subprocess.run(
    ["./toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
assert result.returncode != 0
assert result.stdout == ""
=== END AC invalid-no-success-result ===

## User Acceptance

- None.

## Guardrails

- None.
