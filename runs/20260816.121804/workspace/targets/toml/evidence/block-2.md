# Evidence: Block 2 · Service (block-2)

- block type: block
- date: 2026-08-16
- resulting state: closed/verified
- story points (combined assembled cost): 8951
- execution id: 20260816.122612.488Z-8023b632

## Stories built
- Implement the TOML decoder stdin, tagged-JSON, diagnostic, and exit-status contract. (decoder-contract) [story]

## Stacked context
- compass: COMPASS.md (SP 2942)
- implements: FEATURE-Decoder-Contract.md (SP 694)
- context: stage_contract.sh (SP 317)
- context: ARCHITECTURE_compact.md (SP 123)
- stack: go.md (SP 3646)
- stack: common_compact.md (SP 1179)

## Build directory changes
- cmd/toml-decoder/main.go
- cmd/toml-decoder/main_test.go
- internal/toml/toml.go
- toml-decoder

## Pre-build acceptance observation
- RED: decoder-valid-contract (FEATURE-Decoder-Contract.md)
  intent: A valid TOML document produces the required tagged JSON on stdout with no stderr diagnostic.
  return code: 1
  stdout:
    {}
  stderr:
    --- drydock: values at failure ---
      expected = {'answer': {'type': 'integer', 'value': '42'}}
      result = CompletedProcess(args=['./toml-decoder'], returncode=0, stdout='{}\n', stderr='')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "decoder-valid-contract.py", line 16, in <module>
        assert json.loads(result.stdout) == expected
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: decoder-invalid-contract (FEATURE-Decoder-Contract.md)
  intent: Invalid TOML produces a non-zero exit, no successful stdout result, and a stderr diagnostic.
  return code: 1
  stdout:
    {}
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['./toml-decoder'], returncode=0, stdout='{}\n', stderr='')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "decoder-invalid-contract.py", line 12, in <module>
        assert result.returncode != 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: decoder-stage-contract (FEATURE-Decoder-Contract.md)
  intent: The supplied decoder contract gate passes in full.
  return code: 1
  stdout:
    Traceback (most recent call last):
      File "<stdin>", line 6, in <module>
    AssertionError
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['sh', 'sources/stage_contract.sh', 'decoder'], returncode=1, stdout='', stderr='Traceback (most recent call last):\n  File "<stdin>", line 6, in <module>\nAssertionError\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "decoder-stage-contract.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- PASS: decoder-valid-contract (FEATURE-Decoder-Contract.md)
  intent: A valid TOML document produces the required tagged JSON on stdout with no stderr diagnostic.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    {"answer":{"type":"integer","value":"42"}}
- PASS: decoder-invalid-contract (FEATURE-Decoder-Contract.md)
  intent: Invalid TOML produces a non-zero exit, no successful stdout result, and a stderr diagnostic.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0
  stdout:
    line 1: invalid value "["
- PASS: decoder-stage-contract (FEATURE-Decoder-Contract.md)
  intent: The supplied decoder contract gate passes in full.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python3
  provisioning: not required
  return code: 0

## Agent self-report (advisory)
The build agent declared a failure. This is advisory only; the programmatic acceptance above is the authority for this block's outcome.
- summary: Full TOML conformance is incomplete
- detail:
    The current parser only supports basic scalar key/value pairs; later parser stages must implement arrays, tables, floats, datetimes, and strict validation before rerunning the full suite.

## Build summary
RESULT: FAILED

FILES CHANGED:
- cmd/toml-decoder/main.go
- cmd/toml-decoder/main_test.go
- internal/toml/toml.go

SUMMARY:
Implemented and verified the decoder stdin, tagged-JSON, stderr diagnostic, and exit-status contract. Decoder and architecture gates pass. Full conformance fails on parser features assigned to later stages.

BLOCKERS:
- Full suite remains failing: 167 valid and 59 invalid cases failed.

FAILURE_SUMMARY: Full TOML conformance is incomplete
FAILURE_DETAIL: The current parser only supports basic scalar key/value pairs; later parser stages must implement arrays, tables, floats, datetimes, and strict validation before rerunning the full suite.
