# Blueprint Analysis: toml

## Commander Expectations

- assert the TOML 1.0.0 decoder accepts every supplied valid case and rejects every supplied invalid case.
- assert the decoder follows the stdin-to-tagged-JSON contract.
- assert the implementation uses only the Go standard library.

## Crew

| Crew | Charge |
|---|---|
| Commander | Defines intent and decides what done means. |
| Team Lead | Confirms epic completeness and stakeholder expectations. |
| Planning Crew | Authors atomic specifications and the ordered Manifest. |
| Shipyard Crew | Builds the tickets without synchronous Commander access. |

## Story List

### Feature: Decoder Foundation

| ID | Story | High-level AC |
|---|---|---|
| DECODER-001 | Establish the standard-library Go module and decoder command | The module builds with no third-party dependencies and exposes `cmd/toml-decoder`. |
| DECODER-002 | Implement the stdin, tagged-JSON, diagnostic, and exit-status contract | Valid input produces tagged JSON on stdout with exit zero; invalid input produces stderr diagnostics, no stdout payload, and nonzero exit. |

### Feature: Lexical TOML Values

| ID | Story | High-level AC |
|---|---|---|
| LEXICAL-001 | Parse TOML strings | Basic, literal, multiline basic, and multiline literal strings follow TOML 1.0.0 syntax and escape rules. |
| LEXICAL-002 | Parse booleans, integers, and floats | Supported scalar forms, bases, underscores, special floats, and lossless signed 64-bit integers are handled correctly. |
| LEXICAL-003 | Enforce comments, control characters, and UTF-8 boundaries | Comments and whitespace are handled correctly and forbidden control characters or invalid UTF-8 are rejected. |

### Feature: Keys and Tables

| ID | Story | High-level AC |
|---|---|---|
| STRUCTURE-001 | Parse bare, quoted, and dotted keys | Key forms and dotted-key table creation follow TOML 1.0.0 semantics, including duplicate-key rejection. |
| STRUCTURE-002 | Parse standard tables and table ordering | Tables, implicit parents, redefinitions, and ordering rules are enforced. |
| STRUCTURE-003 | Parse arrays of tables | Array-of-table creation, nesting, most-recent-element targeting, and conflicts are enforced. |

### Feature: Composite Values and Datetimes

| ID | Story | High-level AC |
|---|---|---|
| COMPOSITE-001 | Parse arrays and inline tables | Nested and mixed arrays, comments, multiline arrays, trailing array commas, inline tables, and inline-table closure rules are supported. |
| COMPOSITE-002 | Parse offset and local date/time values | Offset datetimes, local datetimes, local dates, and local times are recognized and encoded using the required tagged representations. |

### Feature: Conformance Delivery

| ID | Story | High-level AC |
|---|---|---|
| CONFORMANCE-001 | Verify complete TOML 1.0.0 conformance | `sources/full_test.sh` completes successfully against the unfiltered installed suite; assertion heading declares `Suite: full`. |

## Surfaced Acceptance Criteria

| ID | Story ID | Criterion |
|---|---|---|
| AC-001 | DECODER-002 | The decoder emits every TOML value as an object containing the required `type` and string-valued `value`, while tables and arrays retain object and array structure. |
| AC-002 | DECODER-002 | The decoder accepts no command-line arguments and has no configuration or side effects. |
| AC-003 | COMPOSITE-002 | Offset datetime values are encoded as RFC 3339; local datetime, date, and time values omit unrelated offset or date components as specified. |

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
| sources/INSTRUCTIONS.md | instruction-to-test | sources/full_test.sh | The instructions define `full_test.sh` as the sole scoring entry point. | Preserve the supplied entry point and make the decoder buildable at its expected path. |
| sources/full_test.sh | dependency | sources/run_conformance.sh | The full script builds `cmd/toml-decoder` and invokes the conformance runner. | Build and contract stories must precede terminal verification. |
| sources/run_conformance.sh | test-kit-to-implementation | cmd/toml-decoder | The harness supplies TOML on stdin and compares decoder output or exit status. | The decoder must remain a filter with the exact process contract. |
| sources/setup_harness.sh | dependency | sources/run_conformance.sh | The runner requires the installed `toml-test` binary and version. | Stage the setup and runner assets; verification depends on harness availability. |
| sources/toml-v1.0.0.md | reference-to-replacement | internal/toml | The specification defines lexical, structural, and encoding behavior. | Parser stories implement the normative specification without third-party modules. |
| sources/stage_contract.sh | test-kit-to-implementation | cmd/toml-decoder | The stage contract checks valid tagged output and invalid-input failure behavior. | Decoder contract must satisfy the governed smoke gate. |
| sources/stage_test.sh | test-kit-to-implementation | parser stories | The stage runner scopes conformance tests to a supplied pattern. | Feature stories use scoped suite slices; only terminal conformance runs the full suite. |

## Source Roles

| Path | Role | Plan disposition | Build disposition |
|---|---|---|---|
| sources/INSTRUCTIONS.md | author intent | compass | prompt-only |
| sources/full_test.sh | conformance harness | context | stage |
| sources/run_conformance.sh | conformance harness | context | stage |
| sources/setup_harness.sh | test helper | context | stage |
| sources/stage_contract.sh | conformance harness | context | stage |
| sources/stage_test.sh | conformance harness | context | stage |
| sources/toml-v1.0.0.md | normative specification and conformance test suite | context | prompt-only |

## Planning Instructions

### Delivery Shape

Deliver a Go command-line filter that reads TOML from stdin, parses it with an internal standard-library-only parser, and writes tagged JSON to stdout. Invalid TOML produces a diagnostic on stderr and a nonzero exit. Supplied shell scripts and the installed `toml-test` binary provide scoped and terminal conformance verification.

### Story Realization Map

| Story ID | Blueprint scope | Evidence | Related files | Delivery kind |
|---|---|---|---|---|
| DECODER-001 | Go module and command scaffold | sources/INSTRUCTIONS.md | `go.mod`, `cmd/toml-decoder` | capability |
| DECODER-002 | Process and JSON boundary | sources/INSTRUCTIONS.md, sources/stage_contract.sh | `cmd/toml-decoder`, `internal/toml` | capability, acceptance contract |
| LEXICAL-001 | String lexer/parser | sources/toml-v1.0.0.md | `internal/toml` | capability |
| LEXICAL-002 | Scalar lexer/parser | sources/toml-v1.0.0.md | `internal/toml` | capability |
| LEXICAL-003 | Input validation and comment handling | sources/toml-v1.0.0.md | `internal/toml` | capability |
| STRUCTURE-001 | Key parsing and definition semantics | sources/toml-v1.0.0.md | `internal/toml` | capability |
| STRUCTURE-002 | Standard table state machine | sources/toml-v1.0.0.md | `internal/toml` | capability |
| STRUCTURE-003 | Array-of-tables state machine | sources/toml-v1.0.0.md | `internal/toml` | capability |
| COMPOSITE-001 | Array and inline-table parsing | sources/toml-v1.0.0.md | `internal/toml` | capability |
| COMPOSITE-002 | Date/time parsing and encoding | sources/toml-v1.0.0.md | `internal/toml` | capability |
| CONFORMANCE-001 | Complete unfiltered verification | sources/full_test.sh, sources/run_conformance.sh | `sources/*`, `toml-test` | acceptance contract |

### Test and Acceptance Strategy

Feature stories use focused parser tests and the corresponding scoped `stage_test.sh` pattern. The decoder contract uses `stage_contract.sh decoder`. The terminal conformance story alone runs `sources/full_test.sh` with `Suite: full`, and gates on the complete unfiltered suite and its exit status.

### Sequencing and Dependencies

Build the module and command first, then establish the process contract. Implement lexical values before structural tables and composite values; implement all parser capabilities before terminal conformance. Stage every shell harness required by verification. Harness installation is an external setup prerequisite and is not part of the deliverable.

### Source Conflicts and Gaps

No conflicting product definitions or implementation requirements were found. Project identity remains proposed for Commander confirmation. No persistence, authentication, UI, deployment target, or external service is described or required by the CLI filter.

## Analysis Notes
generated: 2026-08-15T00:00:00-04:00
blueprint: /mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260816.004939/workspace/targets/toml/blueprint

Quality: Questions
  blockers: 0
  questions: 1
  features: 4
  stories: 11
  stack: Go standard library
  display_name: TOML 1.0.0 Parser
  short_description: A standard-library-only Go command-line parser that converts TOML 1.0.0 input into tagged JSON.

None.
