# Release Scorecard: toml

- Verdict: PASSED
- Code identity: 5151bbfe82258ab8a72545405acbfc5decc38099

## Project acceptance

| ID | Type | Criterion | Verdict | Observed |
|---|---|---|---|---|
| st-001 | technical | The completed parser shall make sh sources/full_test.sh exit zero; that script's exit status is the sole acceptance verdict. | MET | sh sources/full_test.sh: return code 0; toml-test: 205 valid tests and 474 invalid tests passed |

## Failures

- None.

## Manual verification required

- None.

## Could not judge

- None.

## Advisory warnings

- Build directory has uncommitted changes

## Ranked improvements

1. Review the reported uncommitted changes before release.
