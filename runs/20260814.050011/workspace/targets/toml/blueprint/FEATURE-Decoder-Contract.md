# FEATURE: Decoder Contract

| Field       | Value |
|-------------|-------|
| Version     | 20260814 V1 |
| Description | Defines the argument-free stdin-to-tagged-JSON decoder command contract. |
| Depends On  | ARCHITECTURE.md |
| Provides    | cmd/toml-decoder |
| Consumes    | parser module boundary, tagged JSON contract |

## Workflow

The `cmd/toml-decoder` executable reads one TOML document from standard input and delegates parsing to `internal/toml`. For valid input it writes the encoded tagged JSON document to standard output and exits successfully. It accepts no command-line arguments, configuration, or alternate input source.

## Output Behavior

Tables and arrays retain their JSON structure. Scalar values use the tagged representation defined by `ARCHITECTURE.md`. The command does not write files, contact external services, or emit diagnostics for successful input.

## Guardrails

- The command accepts no arguments.
- TOML is read only from stdin.
- Successful output is written only to stdout.
- The command has no configuration or persistent side effects.
- Invalid-input handling is implemented by the decoder error boundary.

## Programmatic Acceptance

=== AC decoder-valid-stdin ===
Intent: A valid TOML document is accepted through stdin and produces a JSON object.
import json
import subprocess

source = 'answer = 42\n'
result = subprocess.run(
    ["./toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
decoded = json.loads(result.stdout)
assert result.returncode == 0
assert isinstance(decoded, dict)
=== END AC decoder-valid-stdin ===

=== AC decoder-argument-free ===
Intent: Supplying an argument is rejected by the argument-free command contract.
import subprocess

result = subprocess.run(
    ["./toml-decoder", "--help"],
    input="",
    capture_output=True,
    text=True,
)
assert result.returncode != 0
=== END AC decoder-argument-free ===

=== AC decoder-stream-separation ===
Intent: Successful decoding does not write diagnostics to stderr.
import subprocess

source = 'enabled = true\n'
result = subprocess.run(
    ["./toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
assert result.returncode == 0
assert result.stderr == ""
=== END AC decoder-stream-separation ===

## User Acceptance

- None.

## Guardrails

- None.
