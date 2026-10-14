# COMPASS: toml

## Compass
A standard-library Go decoder for TOML v1.0.0. It reads a TOML document on stdin and writes the tagged-JSON description of it to stdout, so that conformance can be judged by the upstream toml-test suite. It serves anyone needing a faithful, dependency-free TOML reader, and it is judged by the exit status of the supplied scoring script. Rejecting invalid input counts as much as accepting valid input.

## Constraints
- Filter contract only: stdin in, stdout out, no arguments, no configuration files, no side effects.
- Standard library only; `go.mod` declares no required modules.
- No network use at test time once the harness is installed.
- Scoring assets under `sources/` are read-only and hash-verified.
- Every TOML value is encoded as `{"type", "value"}` with `value` as a JSON string.
- Diagnostics go to stderr; malformed input exits non-zero.

## Guardrails
- Never import a third-party TOML module, including `BurntSushi/toml`, `pelletier/go-toml`, `naoina/toml`, or any other.
- Never modify, skip, filter, or reinterpret the scoring scripts or harness inputs.
- Never special-case fixture input or hard-code expected outputs.
- Never assert on the text of the conformance summary; the exit status is the verdict.
- Never assert a suite case count.
- Reject invalid TOML as strictly as valid TOML is accepted.

<!-- drydock:build-write-guardrail:start -->
## Build Write Guardrail

- Authorized build directory: `/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20261014.124728/build/toml`
- Authorized Target directory: `/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20261014.124728/workspace/targets/toml`
- Build agents have permission to create, modify, and remove files required by the active build block inside these authorized directories.
- No path outside these authorized directories may be modified.
- Protected Drydock artifacts:
  - `/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20261014.124728/workspace/targets/toml/blueprint/`
  - `/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20261014.124728/workspace/targets/toml/MANIFEST.md`
  - `/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20261014.124728/workspace/targets/toml/COMPASS.md`
  - `/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20261014.124728/workspace/targets/toml/QuarterDeck/`
  - `/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20261014.124728/workspace/targets/toml/evidence/`
<!-- drydock:build-write-guardrail:end -->
