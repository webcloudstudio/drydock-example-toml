# Scorecard: toml

## Specification

> **NOT YET RUN**
> drydock score spec toml

## Acceptance

> **PASSED**
> drydock score ac toml · 2026-10-14T13:38:50.502+00:00

PASS 12   FAIL 0   UNVERIFIED 0   PREPASSED 23   (35 AC total)

## Build

> **RECORDED**
> drydock score build toml

Blocks: 10   Repaired: 0   Failed: 0
Tokens: 7,879,917 (cached 7,643,170, output 236,427)   Cache hit rate: 100.0%

## Release

> **RELEASED**
> drydock score release toml · 2026-10-14T13:39:06.567+00:00

- Verdict: PASSED

toml: PASSED — 1 of 1

  st-001  MET       governed_gate full: outcome PASS, return_code 0, timed_out false, build_identity ab50ead6181bb8ac; stdout: valid tests 205 passed, 0 failed; invalid tests 474 passed, 0 failed; encoder tests not run (no encoder command given); stderr: toml-test v2.2.0 under go1.24.6 linux/amd64
  reported: Build directory has uncommitted changes

## Project acceptance

| ID | Type | Criterion | Verdict | Observed |
|---|---|---|---|---|
| st-001 | technical | The completed parser shall make sh sources/full_test.sh exit zero; that script's exit status is the sole acceptance verdict. | MET | governed_gate full: outcome PASS, return_code 0, timed_out false, build_identity ab50ead6181bb8ac; stdout: valid tests 205 passed, 0 failed; invalid tests 474 passed, 0 failed; encoder tests not run (no encoder command given); stderr: toml-test v2.2.0 under go1.24.6 linux/amd64 |

**Advisory warnings:**
- Build directory has uncommitted changes

**Ranked improvements:**
1. The build tree has uncommitted changes (code.dirty is true), so the recorded identity 4c804881d05911b5264a944ef5c279462f85c438 does not fully describe the graded tree. Commit or stash before grading so the verdict is tied to an exact revision.
2. The encoder suite did not run because no encoder command was supplied. If encoding is in scope for release, add an encoder command to the gate so that part is also covered.

**Material decisions:**

- ACCEPTANCE.json not among imported sources — 
- Scoped suite pattern — 
