# Blueprint Analysis: TOML 1.0.0 Parser

## Commander Expectations

- assert the parser accepts every supplied valid TOML 1.0.0 case and rejects every supplied invalid case.
- assert the decoder implements the stdin-to-tagged-JSON filter contract.
- assert the implementation uses only the Go standard library.

## Crew

| Crew | Charge |
|---|---|
| Commander | Defines intent and decides what done means. |
| Team Lead | Confirms epic completeness and stakeholder expectations. |
| Planning Crew | Authors atomic specifications and the ordered Manifest. |
| Shipyard Crew | Builds the tickets without synchronous Commander access. |

## Story List

### Feature: Decoder Contract

| ID | Story | High-level AC |
|---|---|---|
| DECODER-001 | Build the TOML decoder command | `cmd/toml-decoder` reads TOML from stdin, emits tagged JSON on stdout, and exits zero for valid input. |
| DECODER-002 | Report invalid input | Invalid TOML produces a diagnostic on stderr and a non-zero exit without emitting a successful result. |
| DECODER-003 | Encode parsed values as tagged JSON | Tables, arrays, and all supported TOML values use the required JSON structure and string-valued payloads. |

### Feature: TOML Lexical Values

| ID | Story | High-level AC |
|---|---|---|
| LEXICAL-001 | Parse TOML strings | Basic, multiline basic, literal, and multiline literal strings follow TOML 1.0.0 escaping, trimming, UTF-8, and control-character rules. |
| LEXICAL-002 | Parse numbers and booleans | Decimal, hexadecimal, octal, binary, floating-point, special float, and boolean values follow TOML 1.0.0 syntax and lossless integer requirements. |
| LEXICAL-003 | Parse date and time values | Offset datetime, local datetime, local date, and local time values are recognized and encoded using the required representations. |
| LEXICAL-004 | Process whitespace and comments | Whitespace, line endings, comments, and continuation rules are handled without altering valid values or accepting prohibited control characters. |

### Feature: Keys and Key-Value Semantics

| ID | Story | High-level AC |
|---|---|---|
| KEYS-001 | Parse bare, quoted, and dotted keys | Key forms and whitespace around dotted components produce the required nested structure. |
| KEYS-002 | Enforce key definition rules | Duplicate keys, invalid keys, and attempts to turn scalar or closed structures into tables are rejected. |

### Feature: Arrays and Inline Tables

| ID | Story | High-level AC |
|---|---|---|
| STRUCTURES-001 | Parse arrays | Nested, mixed-type, multiline, commented, empty, and trailing-comma arrays are parsed correctly. |
| STRUCTURES-002 | Parse inline tables | Nested inline tables and dotted keys are parsed, while multiline and trailing-comma violations are rejected. |
| STRUCTURES-003 | Enforce inline-table closure | Keys and subtables cannot be added outside an already-defined inline table. |

### Feature: Tables

| ID | Story | High-level AC |
|---|---|---|
| TABLES-001 | Parse standard tables | Explicit and implicitly created tables, including empty and nested tables, are represented correctly. |
| TABLES-002 | Enforce table redefinition rules | Duplicate tables and conflicts between tables, values, and arrays are rejected. |
| TABLES-003 | Parse arrays of tables | Array-of-table headers create ordered table elements and support nested tables and nested arrays of tables. |
| TABLES-004 | Enforce array-of-table ordering rules | Invalid parent ordering and conflicts with statically defined arrays or normal tables are rejected. |

### Feature: Conformance Verification

| ID | Story | High-level AC |
|---|---|---|
| CONFORMANCE-001 | Run the complete TOML conformance suite | **Suite: full** — the supplied `sources/full_test.sh` exits zero with every supplied valid and invalid case passing. |

## Surfaced Acceptance Criteria

| ID | Story ID | Criterion |
|---|---|---|
| AC-001 | DECODER-001 | The program accepts no arguments, has no configuration or side effects, and operates as a stdin/stdout filter. |
| AC-002 | DECODER-003 | Offset datetimes use RFC 3339 encoding; local datetimes omit the offset; local dates and times use their date and time representations. |
| AC-003 | LEXICAL-001 | The parser accepts only valid UTF-8 documents and rejects prohibited control characters and invalid string escapes. |
| AC-004 | LEXICAL-002 | Integers are handled losslessly across the supported signed 64-bit range. |
| AC-005 | CONFORMANCE-001 | The complete suite is invoked unfiltered, and its exit status is the acceptance verdict. |

## Source Inventory

| Path | Content kind | Disposition | Reason |
|---|---|---|---|
| `sources/INSTRUCTIONS.md` | markdown | analyzed | readable UTF-8 |
| `sources/full_test.sh` | code | analyzed | readable UTF-8 |
| `sources/run_conformance.sh` | code | analyzed | readable UTF-8 |
| `sources/setup_harness.sh` | code | analyzed | readable UTF-8 |
| `sources/stage_test.sh` | code | analyzed | readable UTF-8 |
| `sources/toml-v1.0.0.md` | markdown | analyzed | readable UTF-8 |

## Relationship Model

| Source or group | Relationship type | Related source or group | Evidence | Delivery implication |
|---|---|---|---|---|
| `sources/toml-v1.0.0.md` | instruction-to-test | `sources/run_conformance.sh` | The specification defines TOML 1.0.0 behavior; the harness runs `toml-test test -toml 1.0`. | Implementation stories must follow the specification and use scoped conformance slices. |
| `sources/full_test.sh` | test-kit-to-implementation | `cmd/toml-decoder` | The script builds `./cmd/toml-decoder` and supplies it through `DECODER`. | The decoder command must exist before terminal verification. |
| `sources/run_conformance.sh` | test-kit-to-implementation | `cmd/toml-decoder` | The harness pipes TOML to the configured decoder and checks valid and invalid outcomes. | Decoder process behavior is an integration contract. |
| `sources/stage_test.sh` | test-kit-to-implementation | parser features | The stage gate accepts a `-run` pattern for one owned suite slice. | Feature stories may run bounded conformance patterns. |
| `sources/setup_harness.sh` | dependency | `toml-test` | The setup script installs the pinned `v2.2.0` harness. | Harness installation is an external prerequisite, not parser functionality. |

## Source Roles

| Path | Role | Plan disposition | Build disposition |
|---|---|---|---|
| `sources/INSTRUCTIONS.md` | author intent | compass | prompt-only |
| `sources/full_test.sh` | conformance harness | context | stage |
| `sources/run_conformance.sh` | conformance harness | context | stage |
| `sources/setup_harness.sh` | test helper | context | stage |
| `sources/stage_test.sh` | conformance harness | context | stage |
| `sources/toml-v1.0.0.md` | normative specification and conformance test suite | context | prompt-only |

## Planning Instructions

### Delivery Shape

Deliver a Go command-line filter with an internal TOML parser. It consumes UTF-8 TOML on stdin and produces tagged JSON on stdout, or a diagnostic and non-zero exit for invalid input. The staged shell harness builds the command and executes the pinned upstream TOML 1.0.0 conformance suite.

### Story Realization Map

- `DECODER-001` — `cmd/toml-decoder`; evidence `sources/INSTRUCTIONS.md`; capability and CLI contract.
- `DECODER-002` — `cmd/toml-decoder` and parser error boundary; evidence `sources/INSTRUCTIONS.md`; capability and acceptance contract.
- `DECODER-003` — `internal/toml` plus JSON output; evidence `sources/INSTRUCTIONS.md`; capability.
- `LEXICAL-001` through `LEXICAL-004` — `internal/toml`; evidence `sources/toml-v1.0.0.md`; parser capabilities with scoped conformance tests.
- `KEYS-001` and `KEYS-002` — `internal/toml`; evidence `sources/toml-v1.0.0.md`; parser capability and validation.
- `STRUCTURES-001` through `STRUCTURES-003` — `internal/toml`; evidence `sources/toml-v1.0.0.md`; parser capability and validation.
- `TABLES-001` through `TABLES-004` — `internal/toml`; evidence `sources/toml-v1.0.0.md`; parser capability and validation.
- `CONFORMANCE-001` — `sources/full_test.sh`, `sources/run_conformance.sh`, and the completed decoder; terminal acceptance contract.

### Test and Acceptance Strategy

Parser stories use focused unit tests and bounded `toml-test` patterns through `sources/stage_test.sh`, declaring `Suite: scoped`. Decoder contract stories use process-level stdin, stdout, stderr, and exit-code tests. `CONFORMANCE-001` is the sole terminal story declaring `Suite: full`; it runs `sh sources/full_test.sh` unfiltered and gates on the harness exit status.

### Sequencing and Dependencies

Create the Go module and decoder contract first. Implement lexical values before composite structures, then keys and table semantics, followed by scoped conformance verification. Build the decoder before any harness execution. Install the pinned `toml-test` binary before acceptance. `CONFORMANCE-001` depends on every parser and decoder story.

### Source Conflicts and Gaps

No conflicting product definitions were found. Product identity is proposed for Commander confirmation because the metadata fields are blank. The source does not define a separate deployment target; the inferred target is local command-line execution.

## Analysis Notes
generated: 2026-08-14
blueprint: /mnt/c/Users/barlo/projects/drydock/uat/Toml/runs/20260814.050011/workspace/targets/toml/blueprint

Quality: Questions
  blockers: 0
  questions: 1
  features: 6
  stories: 18
  stack: Go
  display_name: TOML 1.0.0 Parser
  short_description: A Go command-line parser that converts TOML 1.0.0 from stdin into tagged JSON and rejects invalid documents.

None.
