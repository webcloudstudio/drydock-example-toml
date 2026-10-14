# Evidence: Block 10 · Feature (block-10)

- block type: block
- date: 2026-10-14
- resulting state: closed/verified
- story points (combined assembled cost): 9119
- execution id: 20261014.133552.021Z-fb0d01d8

## Stories built
- The complete unfiltered toml-test suite after all implementation stories. (conformance) [story]

## Acceptance tooling authorization
- FEATURE-Conformance.md#conformance-full-suite: executable=sh; scope=test; authorization=existing Target environment
- FEATURE-Conformance.md#conformance-full-suite: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Conformance.md#conformance-full-suite: executable=toml-test; scope=test; authorization=existing Target environment
- FEATURE-Conformance.md#conformance-no-third-party: executable=go; scope=test; authorization=existing Target environment

## Stacked context
- compass: COMPASS.md (SP 945)
- implements: FEATURE-Conformance.md (SP 739)
- context: ARCHITECTURE_compact.md (SP 244)
- stack: go.md (SP 5918)
- stack: common_compact.md (SP 1199)

## Build directory changes
- toml-decoder

## Pre-build acceptance observation
- GREEN (prepassed): conformance-full-suite (FEATURE-Conformance.md)
  intent: The complete unfiltered conformance suite exits zero.
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20261014.124728/build/toml/toml-decoder] [no encoder]
      valid tests: 205 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests: 474 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- GREEN (prepassed): conformance-no-third-party (FEATURE-Conformance.md)
  intent: go.mod declares no required module, so no third-party TOML module is present.
  return code: 0

## Post-build programmatic acceptance
- PASS: conformance-full-suite (FEATURE-Conformance.md)
  intent: The complete unfiltered conformance suite exits zero.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    toml-test v2.2.0 [/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20261014.124728/build/toml/toml-decoder] [no encoder]
      valid tests: 205 passed,  0 failed
    encoder tests: no encoder command given
    invalid tests: 474 passed,  0 failed
  stderr:
    harness: /home/barlo/.local/bin/toml-test toml-test v2.2.0; go1.24.6 linux/amd64; race=false; cgo=true
- PASS: conformance-no-third-party (FEATURE-Conformance.md)
  intent: go.mod declares no required module, so no third-party TOML module is present.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0

## Build summary
I'll start by inspecting the build directory and the staged conformance scripts.The scoring entry point is `sh sources/full_test.sh`. Before running it, I'll check the harness version and the current decoder state.The decoder is partial and the existing tests pass. Now I'll run the full suite, the declared verdict, in the background and capture the output to a file outside the build directory.The background run finished. I'll read its output.The full suite exits 0: 205/205 valid and 474/474 invalid pass. I'll confirm the second declared assertion and formatting, then check which files changed.The full conformance suite passes and the step needed no source changes. The verdict is `sh sources/full_test.sh` exiting 0.

- **Full suite:** toml-test v2.2.0, `-toml 1.0`, against the built decoder. Valid tests: 205 passed, 0 failed. Invalid tests: 474 passed, 0 failed. Exit status 0.
- **`conformance-no-third-party`:** the declared Python check passes, since `go.mod` contains no `require`.
- **Existing project tests:** `go test ./...` passed earlier in this session (`cmd/toml-decoder` and `internal/toml`).
- **Formatting:** `gofmt -l .` prints nothing.

I didn't write new tests. The two declared assertions are already covered, and the staged suite is the coverage for this story. The `toml-decoder` binary is rebuilt by `full_test.sh` on each run. It was already modified in the working tree when I started, so I haven't touched it beyond that rebuild. I didn't stage or commit anything.

RESULT: SUCCESS

FILES CHANGED:
- toml-decoder (binary rebuilt by `sources/full_test.sh`; no source changes)

SUMMARY:
Ran `sh sources/full_test.sh` unfiltered from the application root. It exited 0 with 205/205 valid and 474/474 invalid toml-test cases passing under toml-test v2.2.0. The `conformance-no-third-party` assertion passes (go.mod has no `require`). `gofmt -l .` is clean and `go test ./...` was green. Source files were not modified, and no new tests were written because the staged suite covers this story.

BLOCKERS:
- None.
