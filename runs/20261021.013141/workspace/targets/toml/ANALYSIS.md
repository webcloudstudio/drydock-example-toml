# Blueprint Analysis: Toml

## Commander Expectations

- assert the product parses TOML v1.0.0 input and passes every supplied valid and invalid conformance case.

## Crew

| Crew | Charge |
|---|---|
| Commander | Defines intent and decides what done means. |
| Team Lead | Confirms epic completeness and stakeholder expectations. |
| Planning Crew | Authors atomic specifications and the ordered Manifest. |
| Shipyard Crew | Builds the tickets without synchronous Commander access. |

## Story List

### Feature: Architecture and Decoder Contract

| ID | Story | High-level AC |
|---|---|---|
| ARCHITECTURE-001 | Establish the standard-library-only Go module and decoder command | `go build ./cmd/toml-decoder` succeeds without third-party dependencies. |
| DECODER-CONTRACT-001 | Implement stdin, tagged-JSON, diagnostics, and exit-status behavior | Valid TOML produces tagged JSON on stdout with no stderr output; invalid TOML produces diagnostics on stderr and exits non-zero. |

### Feature: Lexical Parsing

| ID | Story | High-level AC |
|---|---|---|
| PARSER-STRINGS-001 | Parse basic, literal, and multiline TOML strings | Valid TOML string cases decode correctly, including escapes, multiline trimming, and literal content. |
| PARSER-LEXICAL-SCALARS-001 | Parse booleans, integers, and floats | Valid scalar cases decode with the required TOML type and lossless string value representation. |
| PARSER-LEXICAL-BOUNDARIES-001 | Enforce comments, control characters, and UTF-8 boundaries | Comments and permitted whitespace are accepted; forbidden control characters and invalid encoding are rejected. |

### Feature: Document Structure

| ID | Story | High-level AC |
|---|---|---|
| PARSER-KEYS-001 | Parse key forms and key-definition semantics | Bare, quoted, dotted, and numeric-looking keys produce the required object structure; duplicate and invalid definitions are rejected. |
| PARSER-TABLES-001 | Parse tables and arrays of tables | Standard tables, nested tables, arrays of tables, ordering, and redefinition rules produce the required structure or reject invalid input. |
| PARSER-ARRAYS-INLINE-001 | Parse arrays, inline tables, and inline-table closure | Nested and mixed arrays, trailing array commas, inline tables, and closure rules conform to TOML 1.0. |
| PARSER-DATETIMES-001 | Parse offset and local date/time forms | Offset datetime, local datetime, local date, and local time values receive the required tagged types and normalized values. |

### Feature: Conformance Verification

| ID | Story | High-level AC |
|---|---|---|
| CONFORMANCE-001 | Verify the completed decoder against the complete supplied suite | The unfiltered supplied conformance command exits zero with every supplied valid and invalid case passing. |

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
| ARCHITECTURE-001 | plan | `sh sources/stage_contract.sh architecture` | Build Instructions: Architecture and module layout |
| DECODER-CONTRACT-001 | plan | `sh sources/stage_contract.sh decoder` | Build Instructions: Interface Contract |
| PARSER-STRINGS-001 | plan | `sh sources/stage_test.sh valid/string/*` | TOML v1.0.0 String section |
| PARSER-LEXICAL-SCALARS-001 | plan | `sh sources/stage_test.sh valid/bool/*,valid/integer/*,valid/float/*,invalid/bool/*,invalid/integer/*,invalid/float/*` | TOML v1.0.0 Boolean, Integer, and Float sections |
| PARSER-LEXICAL-BOUNDARIES-001 | plan | `sh sources/stage_test.sh valid/comment/*,valid/encoding/*,invalid/comment/*,invalid/control/*,invalid/encoding/*` | TOML v1.0.0 Spec and Comment sections |
| PARSER-KEYS-001 | plan | `sh sources/stage_test.sh valid/key/*,invalid/key/*` | TOML v1.0.0 Keys section |
| PARSER-TABLES-001 | plan | `sh sources/stage_test.sh valid/table/*,invalid/table/*,valid/spec-1.0.0/*,invalid/spec-1.0.0/*` | TOML v1.0.0 Table and Array of Tables sections |
| PARSER-ARRAYS-INLINE-001 | plan | `sh sources/stage_test.sh valid/array/*,valid/inline-table/*,invalid/array/*,invalid/inline-table/*` | TOML v1.0.0 Array and Inline Table sections |
| PARSER-DATETIMES-001 | plan | `sh sources/stage_test.sh valid/datetime/*,invalid/datetime/*,invalid/local-date/*,invalid/local-datetime/*,invalid/local-time/*` | TOML v1.0.0 date and time sections |
| CONFORMANCE-001 | plan | `sh sources/full_test.sh` | Build Instructions: complete unfiltered conformance verification |

## Surfaced Acceptance Criteria

The analyze step has surfaced these acceptance criteria for `drydock plan` to fold into the relevant story's typed specification.

| ID | Story ID | Criterion |
|---|---|---|
| AC-001 | DECODER-CONTRACT-001 | The decoder takes no arguments, reads TOML only from stdin, writes no decoded output for invalid input, and produces no side effects. |
| AC-002 | PARSER-LEXICAL-BOUNDARIES-001 | The parser accepts only valid UTF-8 documents and rejects prohibited control characters in comments and strings. |
| AC-003 | PARSER-KEYS-001 | Key definitions are unique, and quoted and bare keys with the same spelling cannot be defined twice. |
| AC-004 | PARSER-TABLES-001 | A table or array-of-tables definition cannot conflict with an existing value, table, static array, or incompatible table form. |
| AC-005 | PARSER-ARRAYS-INLINE-001 | Inline tables cannot be extended after their closing delimiter, and malformed multiline inline-table forms are rejected. |

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
| `sources/INSTRUCTIONS.md` | instruction-to-test | `sources/full_test.sh` | Defines the implementation contract and names the supplied scoring entry point. | Build the decoder at `cmd/toml-decoder`; preserve the supplied runner. |
| `sources/full_test.sh` | dependency | `sources/run_conformance.sh` | Builds the executable and invokes the conformance runner. | Full verification follows successful compilation. |
| `sources/run_conformance.sh` | test-kit-to-implementation | `cmd/toml-decoder` | Passes stdin cases to the decoder and judges exit status and JSON. | Decoder output and exit behavior must remain machine-readable and deterministic. |
| `sources/setup_harness.sh` | test-kit-to-implementation | `sources/run_conformance.sh` | Installs the pinned `toml-test` harness used by the runner. | Harness installation is an external precondition for execution, not product code. |
| `sources/stage_test.sh` | test-kit-to-implementation | parser stories | Provides scoped `-run` gates for story boundaries. | Implementation stories use bounded conformance slices. |
| `sources/stage_contract.sh` | test-kit-to-implementation | architecture and decoder stories | Verifies compilation and valid/invalid stream behavior. | Early stories use the governed contract checks. |
| `sources/toml-v1.0.0.md` | reference-to-replacement | `internal/toml` | Normatively defines TOML syntax, semantics, and invalid cases. | Parser behavior is authored from the specification without third-party TOML modules. |

## Source Roles

| Path | Role | Plan disposition | Build disposition |
|---|---|---|---|
| `sources/INSTRUCTIONS.md` | author intent | context | prompt-only |
| `sources/full_test.sh` | conformance harness | context | stage |
| `sources/run_conformance.sh` | conformance harness | context | stage |
| `sources/setup_harness.sh` | test helper | context | stage |
| `sources/stage_contract.sh` | test helper | context | stage |
| `sources/stage_test.sh` | conformance harness | context | stage |
| `sources/toml-v1.0.0.md` | normative specification and conformance test suite | context | prompt-only |

## Planning Instructions

### Delivery Shape

Deliver a standard-library-only Go command-line filter. It accepts TOML bytes on stdin, parses them into nested tables, arrays, and tagged scalar values, and emits JSON on stdout. Invalid documents produce stderr diagnostics and a non-zero exit. The supplied shell harness and pinned `toml-test` suite provide scoped and terminal verification.

### Story Realization Map

| Story ID | Blueprint scope | Evidence | Related files | Capability |
|---|---|---|---|---|
| ARCHITECTURE-001 | Go module and command boundary | `sources/INSTRUCTIONS.md` | `go.mod`, `cmd/toml-decoder` | capability |
| DECODER-CONTRACT-001 | Filter I/O and tagged encoding | `sources/INSTRUCTIONS.md`, `sources/stage_contract.sh` | `cmd/toml-decoder`, `internal/toml` | acceptance contract |
| PARSER-STRINGS-001 | TOML string grammar | `sources/toml-v1.0.0.md` String section | `internal/toml` | capability |
| PARSER-LEXICAL-SCALARS-001 | Boolean, integer, float grammar | `sources/toml-v1.0.0.md` scalar sections | `internal/toml` | capability |
| PARSER-LEXICAL-BOUNDARIES-001 | Comments, controls, UTF-8 | `sources/toml-v1.0.0.md` Spec and Comment sections | `internal/toml` | capability |
| PARSER-KEYS-001 | Key grammar and uniqueness | `sources/toml-v1.0.0.md` Keys section | `internal/toml` | capability |
| PARSER-TABLES-001 | Tables and arrays of tables | `sources/toml-v1.0.0.md` Table and Array of Tables sections | `internal/toml` | capability |
| PARSER-ARRAYS-INLINE-001 | Arrays and inline tables | `sources/toml-v1.0.0.md` Array and Inline Table sections | `internal/toml` | capability |
| PARSER-DATETIMES-001 | Date/time forms | `sources/toml-v1.0.0.md` datetime sections | `internal/toml` | capability |
| CONFORMANCE-001 | Complete suite verification | `sources/full_test.sh`, `sources/run_conformance.sh` | `sources/*`, built decoder | acceptance contract |

### Test and Acceptance Strategy

Stories use focused parser tests and the corresponding `stage_test.sh` feature slice where available. Architecture and decoder behavior use `stage_contract.sh`. `CONFORMANCE-001` is the sole terminal verification story and runs the unfiltered suite with `Suite: full`; all other conformance checks remain scoped.

### Sequencing and Dependencies

Create the Go module and command first, then implement the decoder contract, lexical parsing, keys, tables, arrays and inline tables, and datetimes. Build before every governed test. Install the pinned harness before conformance execution. The complete suite runs only after every implementation story.

### Source Conflicts and Gaps

No conflicting product definitions or blockers were found. Project identity is proposed from the supplied source and remains subject to Commander confirmation. No authentication, persistence, external service, or deployment platform is required by the sources.

## Analysis Notes
generated: 2026-10-20
blueprint: /mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20261021.013141/workspace/targets/toml/blueprint

Quality: Ready
  blockers: 0
  questions: 0
  features: 4
  stories: 10
  stack: Go
  display_name: TOML 1.0.0 Parser
  short_description: A standard-library-only Go parser that decodes TOML 1.0.0 from stdin into tagged JSON.

None.
