# Scorecard: toml

## Specification

> **NOT YET RUN**
> drydock score spec toml

## Acceptance

> **PASSED**
> drydock score ac toml · 2026-10-21T01:50:56.666+00:00

PASS 5   FAIL 0   UNVERIFIED 0   PREPASSED 13   (18 AC total)

## Build

> **RECORDED**
> drydock score build toml

Blocks: 6   Repaired: 0   Failed: 0
Tokens: 5,635,616 (cached 5,251,840, output 66,150)   Cache hit rate: 94.3%

## Release

> **RELEASED**
> drydock score release toml · 2026-10-21T01:51:08.252+00:00

- Verdict: PASSED

toml: PASSED — 1 of 1

  st-001  MET       sh sources/full_test.sh: return code 0; governed gate stdout: 205 valid tests passed and 474 invalid tests passed
  reported: Build directory has uncommitted changes

## Project acceptance

| ID | Type | Criterion | Verdict | Observed |
|---|---|---|---|---|
| st-001 | technical | The completed parser shall make sh sources/full_test.sh exit zero; that script's exit status is the sole acceptance verdict. | MET | sh sources/full_test.sh: return code 0; governed gate stdout: 205 valid tests passed and 474 invalid tests passed |

**Advisory warnings:**
- Build directory has uncommitted changes

**Ranked improvements:**
1. Review and commit the reported uncommitted changes before release.
