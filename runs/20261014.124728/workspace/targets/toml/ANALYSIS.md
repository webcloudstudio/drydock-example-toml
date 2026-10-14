# Blueprint Analysis: toml

## Commander Expectations

- assert the decoder reads TOML from stdin and writes tagged JSON to stdout
- assert the decoder exits non-zero and writes a diagnostic to stderr on invalid TOML
- assert the full unfiltered conformance suite, run through `sh sources/full_test.sh`, is the acceptance verdict
- assert the parser is hand-written and uses no third-party TOML module

## Crew

| Crew | Charge |
|---|---|
| Commander | Defines intent and decides what done means. |
| Team Lead | Confirms epic completeness and stakeholder expectations. |
| Planning Crew | Authors atomic specifications and the ordered Manifest. |
| Shipyard Crew | Builds the tickets without synchronous Commander access. |

## Story List

The story IDs below are the Commander-owned governed boundary IDs from `INSTRUCTIONS.md`, used unchanged. Each boundary is one story and is not split.

### Feature: Decoder Contract

| ID | Story | High-level AC |
|---|---|---|
| architecture | Standard-library Go module and decoder command boundary | `go.mod` has no `require` directive; `./cmd/toml-decoder` builds |
| decoder-contract | Stdin, tagged JSON, diagnostic-stream, and exit-status contract | Stdin in, tagged JSON on stdout, diagnostic on stderr with non-zero exit on invalid input |

### Feature: Lexical Grammar

| ID | Story | High-level AC |
|---|---|---|
| parser-strings | TOML string forms | Basic, multi-line basic, literal, and multi-line literal strings decode per spec |
| parser-lexical-scalars | Booleans, integers, and floats | Scalar forms, underscores, prefixes, special floats, and range rules decode per spec |
| parser-lexical-boundaries | Comments, control characters, and encoding | Comments and strings reject forbidden control characters; input must be valid UTF-8 |

### Feature: Document Model

| ID | Story | High-level AC |
|---|---|---|
| parser-keys | Key forms and key-definition semantics | Bare, quoted, and dotted keys decode; redefinition is rejected |
| parser-tables | Standard tables, redefinition, arrays of tables, and ordering | Table headers, implicit super-tables, and array-of-tables rules hold |
| parser-arrays-inline | Arrays, inline tables, and inline-table closure | Arrays and inline tables decode; inline tables cannot be extended |
| parser-datetimes | Offset and local date/time forms | Offset and local date, time, and datetime forms decode with the required precision rules |

### Feature: Conformance

| ID | Story | High-level AC |
|---|---|---|
| conformance | The complete unfiltered suite after every implementation stage | `sh sources/full_test.sh` is run unfiltered; its exit status is the verdict |

## Story Guidance

Named stories planning must produce. Authoritative record: `STORY_GUIDANCE.json`.

| Story ID | Provenance | Gate | Note |
|---|---|---|---|
| architecture | commander | `sh sources/stage_contract.sh architecture` | — |
| decoder-contract | commander | `sh sources/stage_contract.sh decoder` | — |
| parser-lexical-scalars | commander | `sh sources/stage_test.sh valid/bool/**|valid/integer/**|valid/float/**|invalid/bool/**|invalid/integer/**|invalid/float/**` | — |
| parser-strings | commander | `sh sources/stage_test.sh valid/string/**|invalid/string/**` | — |
| parser-lexical-boundaries | commander | `sh sources/stage_test.sh valid/comment/**|invalid/control/**|invalid/encoding/**` | — |
| parser-keys | commander | `sh sources/stage_test.sh valid/key/**|invalid/key/**` | — |
| parser-tables | commander | `sh sources/stage_test.sh valid/table/**|invalid/table/**` | — |
| parser-arrays-inline | commander | `sh sources/stage_test.sh valid/array/**|valid/inline-table/**|invalid/array/**|invalid/inline-table/**` | — |
| parser-datetimes | commander | `sh sources/stage_test.sh valid/datetime/**|invalid/datetime/**|invalid/local-date/**|invalid/local-datetime/**|invalid/local-time/**` | — |
| conformance | commander | `sh sources/full_test.sh` | — |

## Surfaced Acceptance Criteria

| ID | Story ID | Criterion |
|---|---|---|
| AC-001 | architecture | `go.mod` declares no `require` directive, and `go build ./cmd/toml-decoder` succeeds using only standard-library imports. |
| AC-002 | architecture | The decoder entry point is the `cmd/toml-decoder` main package, and the parser lives under `internal/toml`. |
| AC-003 | decoder-contract | The decoder takes no arguments, reads no configuration file, and writes only to stdout and stderr. |
| AC-004 | decoder-contract | Tables encode as JSON objects, arrays as JSON arrays, and each value as `{"type": ..., "value": ...}` with `value` as a JSON string. |
| AC-005 | decoder-contract | Offset datetimes encode as RFC 3339, local datetimes encode without an offset, local dates as the date part, and local times as the time part. |
| AC-006 | parser-strings | Basic strings decode `\b \t \n \f \r \" \\ \uXXXX \UXXXXXXXX`, and any other backslash escape is rejected. |
| AC-007 | parser-strings | A multi-line basic string trims a newline immediately after its opening delimiter and applies the line-ending backslash rule. |
| AC-008 | parser-lexical-scalars | Integers in the signed 64-bit range decode losslessly, and values outside that range are rejected. |
| AC-009 | parser-lexical-scalars | The special floats `inf`, `+inf`, `-inf`, `nan`, `+nan`, and `-nan` decode as IEEE 754 values. |
| AC-010 | parser-lexical-boundaries | Control characters other than tab are rejected in comments and literal strings. |
| AC-011 | parser-keys | Defining the same key twice is rejected, and bare and quoted keys with the same name are equivalent. |
| AC-012 | parser-tables | Defining a standard table twice is rejected, and a `[[name]]` header that conflicts with a static table or array is rejected. |
| AC-013 | parser-arrays-inline | Inline tables reject later extension through dotted keys or table headers. |
| AC-014 | parser-datetimes | Fractional seconds beyond the supported precision are truncated, not rounded. |
| AC-015 | conformance | The verdict is the exit status of `sh sources/full_test.sh` run from the application root. |

## Source Inventory

| Path | Content kind | Disposition | Reason |
|---|---|---|---|
| `sources/INSTRUCTIONS.md` | markdown | analyzed | readable UTF-8 |
| `sources/full_test.sh` | code | analyzed | readable UTF-8 |
| `sources/run_conformance.sh` | code | analyzed | readable UTF-8 |
| `sources/setup_harness.sh` | code | analyzed | readable UTF-8 |
| `sources/stage_contract.sh` | code | analyzed | readable UTF-8 |
| `sources/stage_test.sh` | code | analyzed | readable UTF-8 |
| `sources/toml-v1.0.0.md` | markdown | analyzed | readable UTF-8 |

## Relationship Model

| Source or group | Relationship type | Related source or group | Evidence | Delivery implication |
|---|---|---|---|---|
| `sources/INSTRUCTIONS.md` | instruction-to-test | `sources/full_test.sh` | Run Harness section names `full_test.sh` as the scoring entry point | Conformance story gates on `full_test.sh` unchanged. |
| `sources/full_test.sh` | dependency | `sources/run_conformance.sh` | `exec sh sources/run_conformance.sh` | Harness must be staged with the entry point. |
| `sources/run_conformance.sh` | dependency | `sources/setup_harness.sh` | Exit 2 when `toml-test` is absent; version check against `TOML_TEST_VERSION` | Harness install is a precondition, not a build charge. |
| `sources/stage_test.sh` | dependency | `sources/run_conformance.sh` | `exec sh sources/run_conformance.sh -run "$1"` | Parser stories are bounded by pattern through this runner. |
| `sources/stage_contract.sh` | instruction-to-test | `cmd/toml-decoder` | `architecture` and `decoder` stage gates | Governs the architecture and decoder-contract stories. |
| `sources/toml-v1.0.0.md` | reference-to-replacement | `internal/toml` | Spec is normative; INSTRUCTIONS forbids third-party TOML modules | Parser is written from the spec, not from a library. |

## Source Roles

| Path | Role | Plan disposition | Build disposition |
|---|---|---|---|
| sources/INSTRUCTIONS.md | author intent | compass | prompt-only |
| sources/toml-v1.0.0.md | normative specification | context | prompt-only |
| sources/full_test.sh | conformance harness | context | stage |
| sources/run_conformance.sh | conformance harness | context | stage |
| sources/setup_harness.sh | test helper | context | stage |
| sources/stage_contract.sh | test helper | context | stage |
| sources/stage_test.sh | test helper | context | stage |

## Planning Instructions

### Delivery Shape

A command-line filter. It reads a TOML document from stdin, decodes it with a hand-written parser in `internal/toml`, and writes tagged JSON to stdout. Invalid input produces a stderr diagnostic and a non-zero exit. The program has no arguments, configuration, or persistence.

### Story Realization Map

- **architecture**: scope is `go.mod` and the `cmd/toml-decoder` boundary. Evidence: INSTRUCTIONS "Interface Contract" minimal shape and "Definition of Done". Related: `stage_contract.sh architecture`. Requires an acceptance contract and a stdlib-only module.
- **decoder-contract**: scope is stdin, tagged JSON, stderr, and exit status. Evidence: INSTRUCTIONS "Tagged JSON encoding" table. Related: `stage_contract.sh decoder`. Requires a capability.
- **parser-strings**: scope is TOML string forms. Evidence: spec "String". Related: `valid/string`, `invalid/string` groups via `stage_test.sh`. Requires a capability and a test slice.
- **parser-lexical-scalars**: scope is booleans, integers, and floats. Evidence: spec "Integer", "Float", "Boolean". Related: `valid/bool`, `valid/integer`, `valid/float` and their invalid groups. Requires a capability and a test slice.
- **parser-lexical-boundaries**: scope is comments, control characters, and UTF-8. Evidence: spec "Spec" and "Comment". Related: `valid/comment`, `invalid/comment`, `invalid/control`, `invalid/encoding`. Requires a capability and a test slice.
- **parser-keys**: scope is bare, quoted, and dotted keys. Evidence: spec "Keys". Related: `valid/key`, `invalid/key`. Requires a capability and a test slice.
- **parser-tables**: scope is standard tables, array-of-tables, and ordering. Evidence: spec "Table", "Array of Tables". Related: `valid/table`, `invalid/table`. Requires a capability and a test slice.
- **parser-arrays-inline**: scope is arrays and inline tables. Evidence: spec "Array", "Inline Table". Related: `valid/array`, `valid/inline-table`, and their invalid groups. Requires a capability and a test slice.
- **parser-datetimes**: scope is the four datetime forms. Evidence: spec datetime sections. Related: `valid/datetime`, `invalid/datetime`, `invalid/local-date`, `invalid/local-datetime`, `invalid/local-time`. Requires a capability and a test slice.
- **conformance**: scope is the terminal unfiltered run. Evidence: `full_test.sh`. Requires a proof gate, not a capability.

### Test and Acceptance Strategy

- The terminal story `conformance` is the only story that runs the complete suite. It depends on all nine implementation stories and declares `Suite: full`. Its gate is `sh sources/full_test.sh`, with no selector.
- Every other story is bounded. The `architecture` and `decoder-contract` stories run `sh sources/stage_contract.sh <architecture|decoder>`. The parser stories run `sh sources/stage_test.sh <group patterns>` scoped to the groups they own, and declare `Suite: scoped`.
- Verdicts are exit statuses. No story asserts the text of the summary line or a case count.
- A stdlib-only, no-third-party guardrail is verified by `stage_contract.sh architecture` and by SEA st-002.
- No source supplies a numeric release threshold. The universal "all cases pass" requirement is kept in its universal form and proven by the full-suite run.

### Sequencing and Dependencies

- `architecture` → `decoder-contract` → `parser-strings` → `parser-lexical-scalars` → `parser-lexical-boundaries` → `parser-keys` → `parser-tables` → `parser-arrays-inline` → `parser-datetimes` → `conformance`.
- `decoder-contract` must follow `parser-keys` and `parser-lexical-scalars` in the Manifest. Its governed gate `stage_contract.sh decoder` decodes `answer = 42` and rejects an unterminated array, so it cannot pass before key and integer parsing exist. Its position in the list is governed; the Manifest order carries this constraint.
- `parser-tables` depends on `parser-keys`. `parser-arrays-inline` depends on `parser-strings`, `parser-lexical-scalars`, and `parser-keys`.
- `conformance` depends on all implementation stories.

### Source Conflicts and Gaps

- No contradictions among the imported sources. The toml-test version `v2.2.0` is consistent across `full_test.sh`, `run_conformance.sh`, `stage_test.sh`, and `setup_harness.sh`.
- The story boundaries are taken from the list in `INSTRUCTIONS.md`. The `ACCEPTANCE.json` file it cites is not among the imported sources, so Plan should preserve those boundary IDs as written.
- The sample module path `github.com/owner/toml-decoder` is a placeholder. The builder uses it as the module path, because the sources name no other.
- `stage_contract.sh` uses `python` to check JSON. That is harness tooling only and does not make Python part of the deliverable.
- The `setup_harness.sh` install requires network access. That is a one-time precondition, not a product dependency.

## Analysis Notes
generated: 2026-10-14
blueprint: /mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20261014.124728/workspace/targets/toml/blueprint

Quality: Ready
  blockers: 0
  questions: 0
  features: 4
  stories: 10
  stack: Go, standard library only
  display_name: not proposed
  short_description: A TOML v1.0.0 decoder that reads TOML on stdin and writes tagged JSON to stdout, written in Go with the standard library only and verified against the toml-test conformance suite.

Project type: `cli`. The program is a stdin-to-stdout filter with no subcommands and no arguments.
Core baseline (cli): the program has no commands or sub-verbs and takes no arguments, so the sources define no help surface. This is satisfied, not a gap.
Gap checklist: no persistence, auth, or external service calls are involved in the product. The deployment target is derived as a local `go build -o toml-decoder ./cmd/toml-decoder` executable.
SEA_TRIALS st-001 is preserved from the prior project contract with the same ID and criterion.
