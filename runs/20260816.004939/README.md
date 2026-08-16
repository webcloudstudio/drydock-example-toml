# Toml: PASSED

5 of 6 receipt claims proven. Open `index.html` for the linked proof kit; verify it with `sha256sum -c SHA256SUMS`.

## Receipt

| Claim | Verdict | Recorded outcome | Proof |
|---|---|---|---|
| Lifecycle completed | PASS | 18 lifecycle commands executed; the run ended at 18-score-release. | [result.json](result.json) |
| External conformance suite passed | PASS | sh sources/full_test.sh exited 0. | [evidence/commands/15-test.stdout.log](evidence/commands/15-test.stdout.log) |
| Target completion check passed | PASS | drydock status toml --check exited 0. | [evidence/commands/11-initial-complete.stdout.log](evidence/commands/11-initial-complete.stdout.log) |
| Acceptance score passed | FAIL | drydock score acceptance exited 1. | [evidence/commands/16-score-acceptance.stdout.log](evidence/commands/16-score-acceptance.stdout.log) |
| Release score passed | PASS | drydock score release exited 0. | [evidence/commands/18-score-release.stdout.log](evidence/commands/18-score-release.stdout.log) |
| Integrity verification passed | PASS | 344 files digested; verify with sha256sum -c SHA256SUMS. | [SHA256SUMS](SHA256SUMS) |

## Run facts

- Drydock: `0.1.8` (commit `3f2d0d7fa06faa2e89efcfd9139ce10daa74b1c4`)
- Provider and model: `codex` / `gpt-5.6-luna`
- Platform: `Linux-6.6.87.2-microsoft-standard-WSL2-x86_64-with-glibc2.35` on Python `3.12.13`
- Target: `toml`
- Run: `20260816.004939`
- Ran: 2026-08-15 23:26:00 EDT — 2026-08-15 23:27:29 EDT
- Elapsed: 89.5s
- LLM calls: 27
- Tokens: cached 10,483,712; uncached 820,937; output 110,979
- LLM elapsed: 2709.5s
- Build passes: 1; repairs: 6 attempts allowed
- Conformance: passed
- Verdict: expected PASSED, observed PASSED
- Advisory scores: acceptance=exit 1, build-report=exit 0, release=exit 0

## RUN SUMMARY

- Input specification: [`sources/toml-v1.0.0.md`](sources/toml-v1.0.0.md)
- Delivered Code: [`build/toml/`](build/toml)
- Test Results: [`evidence/commands/15-test.stdout.log`](evidence/commands/15-test.stdout.log)

## RUN NOTES:

- One run is evidence of one run. It is not a benchmark.
- It is not a security certification of the delivered code.

## Commands

| # | Command | Exit | Elapsed | Output |
|---|---|---|---|---|
| 01-init | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python3 -m drydock init toml` | 0 | 1.7s | [stdout](evidence/commands/01-init.stdout.log) · [stderr](evidence/commands/01-init.stderr.log) |
| 02-import-sources | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python3 -m drydock import toml sources --format markdown` | 0 | 2.3s | [stdout](evidence/commands/02-import-sources.stdout.log) · [stderr](evidence/commands/02-import-sources.stderr.log) |
| 03-analyze | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python3 -m drydock analyze toml` | 0 | 64.8s | [stdout](evidence/commands/03-analyze.stdout.log) · [stderr](evidence/commands/03-analyze.stderr.log) · [llm](evidence/llm_logs/20260816.004945.256Z_toml_analyze_codex.llm.log) |
| 04-plan | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python3 -m drydock plan toml --override` | 0 | 395.3s | [stdout](evidence/commands/04-plan.stdout.log) · [stderr](evidence/commands/04-plan.stderr.log) · [llm](evidence/llm_logs/20260816.005050.540Z_toml_plan_codex.llm.log) · [llm](evidence/llm_logs/20260816.005148.127Z_toml_plan_codex.llm.log) · [llm](evidence/llm_logs/20260816.005236.283Z_toml_plan_codex.llm.log) · [llm](evidence/llm_logs/20260816.005339.472Z_toml_plan_codex.llm.log) · [llm](evidence/llm_logs/20260816.005452.169Z_toml_lineage_attribute_codex.llm.log) · [llm](evidence/llm_logs/20260816.005515.010Z_toml_lineage_attribute_codex.llm.log) · [llm](evidence/llm_logs/20260816.005530.089Z_toml_lineage_attribute_codex.llm.log) · [llm](evidence/llm_logs/20260816.005557.468Z_toml_lineage_attribute_codex.llm.log) · [llm](evidence/llm_logs/20260816.005611.866Z_toml_lineage_attribute_codex.llm.log) · [llm](evidence/llm_logs/20260816.005623.968Z_toml_lineage_attribute_codex.llm.log) · [llm](evidence/llm_logs/20260816.005639.611Z_toml_lineage_attribute_codex.llm.log) |
| 05-after-plan-build-status | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python3 -m drydock build status toml` | 0 | 0.5s | [stdout](evidence/commands/05-after-plan-build-status.stdout.log) · [stderr](evidence/commands/05-after-plan-build-status.stderr.log) |
| 06-after-plan-target-status | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python3 -m drydock status toml` | 0 | 0.6s | [stdout](evidence/commands/06-after-plan-target-status.stdout.log) · [stderr](evidence/commands/06-after-plan-target-status.stderr.log) |
| 07-after-plan-workspace-status | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python3 -m drydock status` | 0 | 0.4s | [stdout](evidence/commands/07-after-plan-workspace-status.stdout.log) · [stderr](evidence/commands/07-after-plan-workspace-status.stderr.log) |
| 08-initial-ready | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python3 -m drydock status toml --ready` | 0 | 0.5s | [stdout](evidence/commands/08-initial-ready.stdout.log) · [stderr](evidence/commands/08-initial-ready.stderr.log) |
| 09-initial-build-1 | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python3 -m drydock build toml --override --repair-attempts 6` | 0 | 2410.7s | [stdout](evidence/commands/09-initial-build-1.stdout.log) · [stderr](evidence/commands/09-initial-build-1.stderr.log) · [llm](evidence/llm_logs/20260816.005728.172Z_toml_build_codex.llm.log) · [llm](evidence/llm_logs/20260816.005855.689Z_toml_build_codex.llm.log) · [llm](evidence/llm_logs/20260816.010051.090Z_toml_build_codex.llm.log) · [llm](evidence/llm_logs/20260816.010655.417Z_toml_build_codex.llm.log) · [llm](evidence/llm_logs/20260816.010953.453Z_toml_build_codex.llm.log) · [llm](evidence/llm_logs/20260816.011402.538Z_toml_build_codex.llm.log) · [llm](evidence/llm_logs/20260816.011726.759Z_toml_build_codex.llm.log) · [llm](evidence/llm_logs/20260816.012022.052Z_toml_build_codex.llm.log) · [llm](evidence/llm_logs/20260816.012414.115Z_toml_build_codex.llm.log) · [llm](evidence/llm_logs/20260816.013008.609Z_toml_build_codex.llm.log) · [llm](evidence/llm_logs/20260816.013330.042Z_toml_build_codex.llm.log) |
| 10-initial-ready | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python3 -m drydock status toml --ready` | 1 | 0.3s | [stdout](evidence/commands/10-initial-ready.stdout.log) · [stderr](evidence/commands/10-initial-ready.stderr.log) |
| 11-initial-complete | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python3 -m drydock status toml --check` | 0 | 0.3s | [stdout](evidence/commands/11-initial-complete.stdout.log) · [stderr](evidence/commands/11-initial-complete.stderr.log) |
| 12-after-initial-build-build-status | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python3 -m drydock build status toml` | 0 | 0.5s | [stdout](evidence/commands/12-after-initial-build-build-status.stdout.log) · [stderr](evidence/commands/12-after-initial-build-build-status.stderr.log) |
| 13-after-initial-build-target-status | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python3 -m drydock status toml` | 0 | 0.6s | [stdout](evidence/commands/13-after-initial-build-target-status.stdout.log) · [stderr](evidence/commands/13-after-initial-build-target-status.stderr.log) |
| 14-after-initial-build-workspace-status | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python3 -m drydock status` | 0 | 0.3s | [stdout](evidence/commands/14-after-initial-build-workspace-status.stdout.log) · [stderr](evidence/commands/14-after-initial-build-workspace-status.stderr.log) |
| 15-test | `sh sources/full_test.sh` | 0 | 2.9s | [stdout](evidence/commands/15-test.stdout.log) · [stderr](evidence/commands/15-test.stderr.log) |
| 16-score-acceptance | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python -m drydock score ac toml` | 1 | 62.4s | [stdout](evidence/commands/16-score-acceptance.stdout.log) · [stderr](evidence/commands/16-score-acceptance.stderr.log) |
| 17-score-build-report | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python -m drydock score build toml` | 0 | 0.5s | [stdout](evidence/commands/17-score-build-report.stdout.log) · [stderr](evidence/commands/17-score-build-report.stderr.log) |
| 18-score-release | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python -m drydock score release toml` | 0 | 15.0s | [stdout](evidence/commands/18-score-release.stdout.log) · [stderr](evidence/commands/18-score-release.stderr.log) · [llm](evidence/llm_logs/20260816.032708.567Z_toml_score-release_codex.llm.log) |

## Evidence

- [`evidence/commands/`](evidence/commands) — stdout and stderr of every command
- [`evidence/prompts/`](evidence/prompts) — the assembled prompt for every LLM call
- [`evidence/provider_raw/`](evidence/provider_raw) — unmodified provider transcripts
- [`evidence/llm_logs/`](evidence/llm_logs) — call banners and token accounting
- [`result.json`](result.json) — the machine-readable record of this run
