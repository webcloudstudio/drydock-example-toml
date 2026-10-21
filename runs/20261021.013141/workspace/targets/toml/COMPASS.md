# COMPASS: Toml

## Compass

Build a Go command-line filter that parses TOML v1.0.0 and emits the required tagged JSON representation. It serves the supplied conformance harness and must correctly accept valid documents while rejecting invalid ones.

## Constraints

- Use Go and the standard library only.
- Provide the decoder at `cmd/toml-decoder`.
- Read TOML exclusively from stdin and emit JSON exclusively to stdout.
- Invalid TOML must produce stderr diagnostics and a non-zero exit.
- The supplied harness files are read-only and must remain unchanged.
- Conform to TOML v1.0.0 semantics and the pinned `toml-test` suite.

## Guardrails

- Never add a third-party TOML parsing dependency.
- Never weaken invalid-input rejection to improve permissive acceptance.
- Never emit decoded output for invalid input.
- Never change, filter, skip, or reinterpret the supplied scoring harness.
- Preserve lossless tagged representations required by the decoder contract.

<!-- drydock:build-write-guardrail:start -->
## Build Write Guardrail

- Authorized build directory: `/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20261021.013141/build/toml`
- Authorized Target directory: `/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20261021.013141/workspace/targets/toml`
- Build agents have permission to create, modify, and remove files required by the active build block inside these authorized directories.
- No path outside these authorized directories may be modified.
- Protected Drydock artifacts:
  - `/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20261021.013141/workspace/targets/toml/blueprint/`
  - `/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20261021.013141/workspace/targets/toml/MANIFEST.md`
  - `/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20261021.013141/workspace/targets/toml/COMPASS.md`
  - `/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20261021.013141/workspace/targets/toml/QuarterDeck/`
  - `/mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20261021.013141/workspace/targets/toml/evidence/`
<!-- drydock:build-write-guardrail:end -->
