# Blueprint Analysis: TOML Parser

## Commander Expectations

- assert the parser passes every supplied valid and invalid TOML conformance case.
- assert the decoder reads TOML from stdin and emits tagged JSON on stdout.
- assert invalid TOML produces a diagnostic on stderr and a non-zero exit.

## Crew

| Crew | Charge |
|---|---|
| Commander | Defines intent and decides what done means. |
| Team Lead | Confirms epic completeness and stakeholder expectations. |
| Planning Crew | Authors atomic specifications and the ordered Manifest. |
| Shipyard Crew | Builds the tickets without synchronous Commander access. |

## Story List

### Feature: Foundation and Decoder Contract

| ID | Story | High-level AC |
|---|---|---|
| FOUNDATION-001 | Establish the standard-library-only Go module and decoder command | The module builds successfully and declares no third-party dependencies. |
| FOUNDATION-002 | Implement the stdin-to-tagged-JSON decoder contract | Valid input produces tagged JSON on stdout; invalid input produces stderr diagnostics and exits non-zero. |

### Feature: Lexical Parsing

| ID | Story | High-level AC |
|---|---|---|
| LEXICAL-001 | Parse basic, literal, and multiline TOML strings | All supplied valid and invalid string cases are handled according to TOML 1.0.0. |
| LEXICAL-002 | Parse booleans, integers, and floating-point values | Supported scalar forms, bases, underscores, special floats, and range rules conform to TOML 1.0.0. |
| LEXICAL-003 | Enforce comments, control-character, UTF-8, and lexical boundaries | Invalid control characters, malformed encoding, and malformed lexical forms are rejected. |

### Feature: Document Structure and Values

| ID | Story | High-level AC |
|---|---|---|
| STRUCTURE-001 | Parse bare, quoted, and dotted keys | Key forms map to the correct nested objects and duplicate or conflicting definitions are rejected. |
| STRUCTURE-002 | Parse standard tables and arrays of tables | Table creation, nesting, ordering, redefinition, and array-element targeting follow TOML 1.0.0. |
| STRUCTURE-003 | Parse arrays and inline tables | Nested, mixed, multiline, trailing-comma arrays and closed inline tables are handled correctly. |
| STRUCTURE-004 | Parse offset and local date/time values | Date/time types are recognized and serialized using the required tagged representations. |

### Feature: Conformance Verification

| ID | Story | High-level AC |
|---|---|---|
| CONFORMANCE-001 | Verify the completed parser against the complete supplied conformance suite | The unfiltered supplied suite exits zero with every valid and invalid case passing. |

## Surfaced Acceptance Criteria

None.

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
| `sources/INSTRUCTIONS.md` | instruction-to-test | `sources/full_test.sh` | Defines the stdin/stdout contract and identifies the full test command as the sole verdict. | Foundation and conformance stories must implement and invoke the supplied contract. |
| `sources/full_test.sh` | dependency | `sources/run_conformance.sh` | Builds `./cmd/toml-decoder`, sets `DECODER`, and delegates to the runner. | The command path and build layout are fixed before conformance verification. |
| `sources/run_conformance.sh` | test-kit-to-implementation | `cmd/toml-decoder` | Pipes the TOML corpus into the decoder and checks valid/invalid behavior. | The decoder must remain a filter with no argument-dependent interface. |
| `sources/toml-v1.0.0.md` | reference-to-replacement | `internal/toml` | Normative syntax and semantic rules define parser behavior. | All parser stories implement the specification directly. |
| `sources/stage_contract.sh` | acceptance contract | `cmd/toml-decoder` | Provides architecture and decoder smoke gates, including malformed-input rejection. | Foundation stories must satisfy these governed gates before feature work is accepted. |
| `sources/stage_test.sh` | test-kit-to-implementation | parser feature stories | Runs bounded `toml-test` slices for individual feature groups. | Feature stories use scoped suite runs; only CONFORMANCE-001 runs the full suite. |

## Source Roles

| Path | Role | Plan disposition | Build disposition |
|---|---|---|---|
| `sources/INSTRUCTIONS.md` | author intent | compass | prompt-only |
| `sources/toml-v1.0.0.md` | normative specification | context | prompt-only |
| `sources/full_test.sh` | conformance harness | context | stage |
| `sources/run_conformance.sh` | conformance harness | context | stage |
| `sources/setup_harness.sh` | test helper | context | stage |
| `sources/stage_contract.sh` | acceptance contract | context | stage |
| `sources/stage_test.sh` | acceptance contract | context | stage |

## Planning Instructions

### Delivery Shape

Deliver a standard-library-only Go command-line filter. It accepts a UTF-8 TOML document on stdin, parses it into an internal document model, emits tagged JSON on stdout, and reports invalid input on stderr with a non-zero exit. The supplied shell scripts and `toml-test` corpus provide staged acceptance execution.

### Story Realization Map

| Story ID | Blueprint scope | Evidence | Related files | Delivery kind |
|---|---|---|---|---|
| FOUNDATION-001 | Go module and decoder command | `sources/INSTRUCTIONS.md`, `sources/stage_contract.sh` | `go.mod`, `cmd/toml-decoder` | capability, acceptance contract |
| FOUNDATION-002 | stdin, JSON, diagnostics, exit status | `sources/INSTRUCTIONS.md`, `sources/stage_contract.sh` | `cmd/toml-decoder`, `internal/toml` | capability, acceptance contract |
| LEXICAL-001 | all TOML string forms | `sources/toml-v1.0.0.md` string section | `internal/toml` | capability, scoped harness |
| LEXICAL-002 | booleans, integers, floats | `sources/toml-v1.0.0.md` scalar sections | `internal/toml` | capability, scoped harness |
| LEXICAL-003 | comments, controls, encoding, lexical rejection | `sources/toml-v1.0.0.md`, `sources/stage_contract.sh` | `internal/toml` | capability, scoped harness |
| STRUCTURE-001 | key forms and definitions | `sources/toml-v1.0.0.md` keys section | `internal/toml` | capability, scoped harness |
| STRUCTURE-002 | tables and arrays of tables | `sources/toml-v1.0.0.md` table sections | `internal/toml` | capability, scoped harness |
| STRUCTURE-003 | arrays and inline tables | `sources/toml-v1.0.0.md` array and inline-table sections | `internal/toml` | capability, scoped harness |
| STRUCTURE-004 | all date/time forms | `sources/toml-v1.0.0.md` date/time sections | `internal/toml` | capability, scoped harness |
| CONFORMANCE-001 | complete unfiltered verification | `sources/full_test.sh`, `sources/run_conformance.sh` | staged `sources/*`, `cmd/toml-decoder` | terminal acceptance contract |

### Test and Acceptance Strategy

Foundation stories use `sources/stage_contract.sh`. Lexical and structure stories run only their corresponding bounded `toml-test` slices through `sources/stage_test.sh` and declare `Suite: scoped`. CONFORMANCE-001 is the sole terminal verification story, depends on every implementation story, declares `Suite: full`, and runs `sh sources/full_test.sh`.

### Sequencing and Dependencies

Build FOUNDATION-001 before all tests. Implement FOUNDATION-002 before parser slices. Parser lexical behavior precedes structural semantics; all parser stories precede CONFORMANCE-001. The staged harness scripts must remain unchanged. `toml-test` v2.2.0 must be installed before conformance execution.

### Source Conflicts and Gaps

No conflicting product definitions were found. The product purpose, language, interface, specification, and acceptance command are defined. Project identity fields are blank and require Commander confirmation.

## Analysis Notes
generated: 2026-08-14
blueprint: /mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260815.014855/workspace/targets/toml/blueprint

Quality: Questions
  blockers: 0
  questions: 1
  features: 4
  stories: 10
  stack: Go, standard library, POSIX shell, toml-test v2.2.0
  display_name: TOML Parser
  short_description: A Go command-line TOML v1.0.0 parser that converts stdin into tagged JSON and rejects invalid documents.

No blockers. Identity confirmation remains open.
