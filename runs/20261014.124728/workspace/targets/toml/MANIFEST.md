# MANIFEST: toml
updated:     2026-10-14T12:51:29+00:00
state:       approved
blocks:      10

applied_specs: |
  ARCHITECTURE.md sha256=97687a3f941d71ba070274e7fe1318c1951bc871777f1e85bac25312dc5ed387 commit=601a291786af88ca57995fb5dbeb34d82a8bfafc applied_by=block-1 applied_at=2026-10-14T12:56:33+00:00 build_sha256=4d5c8ae4286085f9b9194856c7dee3162ed524065ed834c68507e46cfa96fc99
  ARCHITECTURE_compact.md sha256=5b8855b225fea341d444ecc7ca611183da3a15c92e858898c5a8e682c9817ebf commit=356aede28e1d9a7cb3204cc10413fb6f2754aad7 applied_by=block-10 applied_at=2026-10-14T13:37:57+00:00 build_sha256=5b8855b225fea341d444ecc7ca611183da3a15c92e858898c5a8e682c9817ebf
  FEATURE-Conformance.md sha256=aced4eb3e95a72718bb6e4d469a0f8cab2a6152b747991a5b3ead9a819dd49b9 commit=601a291786af88ca57995fb5dbeb34d82a8bfafc applied_by=block-10 applied_at=2026-10-14T13:37:57+00:00 build_sha256=87ce8050491e659a90c23ce7aed7e43afd8e9fcb65c4ff9521e7937f6340be85
  FEATURE-Decoder-Contract.md sha256=0b4eafd6199276643f842fc039427edeb25a21d31ed2e8ef41c9b6d51315179c commit=601a291786af88ca57995fb5dbeb34d82a8bfafc applied_by=block-8 applied_at=2026-10-14T13:32:12+00:00 build_sha256=0d368ff4a99c2894255e1f02ec7eddb3f78cfac472697854c462e4680e46cd3f
  FEATURE-Parser-Arrays-Inline.md sha256=6fc45d37577cb6bc8eb6f723ed5b9369b11f7be1375967213724562fa8e92bb4 commit=601a291786af88ca57995fb5dbeb34d82a8bfafc applied_by=block-7 applied_at=2026-10-14T13:29:28+00:00 build_sha256=a971858e3af3629dbbeb74e75a261ebd774fedfad08d90e3f5afe5fd2280bd65
  FEATURE-Parser-Datetimes.md sha256=92ea7b03b39157bed042f0d9f6ec27dda75998118e6bf5652acc945cba5296eb commit=601a291786af88ca57995fb5dbeb34d82a8bfafc applied_by=block-9 applied_at=2026-10-14T13:35:42+00:00 build_sha256=d7c351eef8817f8c719b4a84fa6f2c940851c1b2a16b700d00d2f01598815725
  FEATURE-Parser-Keys.md sha256=aaf9bf6d4ecbc8be7b3ed0905d4c44680b775ef27271286386e99463e6019366 commit=601a291786af88ca57995fb5dbeb34d82a8bfafc applied_by=block-5 applied_at=2026-10-14T13:18:17+00:00 build_sha256=2a84b8e97338404bca7ae5d43655e993bc7c8960dc0742dbc1514c1695a5497e
  FEATURE-Parser-Lexical-Boundaries.md sha256=590cf8c0d4077cd3001bd05c90d8392d6c12fcebfc4e79aa9fb1c6f839b316a9 commit=601a291786af88ca57995fb5dbeb34d82a8bfafc applied_by=block-4 applied_at=2026-10-14T13:05:59+00:00 build_sha256=e9b7e756221979c6b20f049a5c0556cd5e5b98f92fee84e4d968c0a4b9e669ce
  FEATURE-Parser-Lexical-Scalars.md sha256=b587b41c7cadd9f1ac78d1153716f40df7eb8f00263dd820de19a85ef4ad4cde commit=601a291786af88ca57995fb5dbeb34d82a8bfafc applied_by=block-3 applied_at=2026-10-14T13:03:12+00:00 build_sha256=d15eb618156682a7472c473a413abfca0c0b39f946534b26b864e0a02e962c4c
  FEATURE-Parser-Strings.md sha256=165b41774d545509acdce9b04e126c2d5dd8d1205f815d6ed88fde4fdfd814ea commit=601a291786af88ca57995fb5dbeb34d82a8bfafc applied_by=block-2 applied_at=2026-10-14T13:00:27+00:00 build_sha256=8e925d834a07291185fa009a130760a069506aed09ad5af68ebbaca376e8155c
  FEATURE-Parser-Tables.md sha256=fdd2f88fe46b1481b02d45777e39a01fd1b54fe0b754328c896cc8f204dfa535 commit=601a291786af88ca57995fb5dbeb34d82a8bfafc applied_by=block-6 applied_at=2026-10-14T13:23:46+00:00 build_sha256=5eff78968f9cb581c15b648451ff5413e68ebde1396dd33380914db42b3dbbf4
## story 1: Standard-library Go module and decoder command boundary.
id:           architecture
evidence: evidence/block-1.md
summary:      Standard-library Go module and decoder command boundary.
type:         foundational
kind:         capability
phase:        1
block:        1
implements:   ARCHITECTURE.md
covers:       architecture
accepts:      st-001
instructions: |
  Create go.mod with the module path github.com/owner/toml-decoder and no require directives.
  Create cmd/toml-decoder/main.go as a package main that reads stdin, delegates to internal/toml, and encodes to stdout.
  Create internal/toml as a package with a stub Decode function so the module builds with the standard library only.
stack:        go.md, common.md
stack_mode:   builder
size:         10961
provides:     go.mod, cmd/toml-decoder, internal/toml
acceptance:   yes
state: closed/verified

## story 2: TOML string forms: basic, multi-line basic, literal, multi-line literal.
id: parser-strings
evidence: evidence/block-2.md
summary: TOML string forms: basic, multi-line basic, literal, multi-line literal.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-Parser-Strings.md
covers: parser-strings
accepts: st-001
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 9255
provides: string value decoding in internal/toml
instructions: |
  Implement the four string forms in internal/toml per sources/toml-v1.0.0.md "String".
  Basic strings accept only the listed escapes; any other backslash escape is rejected.
  Multi-line basic strings trim a newline after the opening delimiter and apply the line-ending backslash rule.
acceptance: yes
depends: architecture
state: closed/verified

## story 3: Booleans, integers, and floats.
id: parser-lexical-scalars
evidence: evidence/block-3.md
summary: Booleans, integers, and floats.
type: service
kind: capability
phase: 1
block: 3
implements: FEATURE-Parser-Lexical-Scalars.md
covers: parser-lexical-scalars
accepts: st-001
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 9239
provides: boolean, integer, and float decoding in internal/toml
instructions: |
  Implement booleans, integers (signed 64-bit, lossless, with 0x, 0o, 0b prefixes and underscores), and floats (IEEE 754 binary64, inf, nan, exponents).
  Reject out-of-range integers and malformed float forms.
acceptance: yes
depends: architecture
state: closed/verified

## story 4: Comments, control characters, and UTF-8 encoding.
id: parser-lexical-boundaries
evidence: evidence/block-4.md
summary: Comments, control characters, and UTF-8 encoding.
type: service
kind: capability
phase: 1
block: 4
implements: FEATURE-Parser-Lexical-Boundaries.md
covers: parser-lexical-boundaries
accepts: st-001
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 9445
provides: comment handling and control-character and UTF-8 validation
instructions: |
  Treat # as a comment start outside strings to end of line.
  Reject control characters other than tab in comments and literal strings.
  Reject input that is not valid UTF-8.
acceptance: yes
depends: parser-strings, parser-lexical-scalars
state: closed/verified

## story 5: Key forms and key-definition semantics.
id: parser-keys
evidence: evidence/block-5.md
summary: Key forms and key-definition semantics.
type: service
kind: capability
phase: 1
block: 5
implements: FEATURE-Parser-Keys.md
covers: parser-keys
accepts: st-001
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 9528
provides: bare, quoted, and dotted key parsing with redefinition checks
instructions: |
  Parse bare keys (A-Za-z0-9_-), quoted keys (basic or literal strings), and dotted keys.
  Reject defining the same key twice; bare and quoted keys with the same name are equivalent.
acceptance: yes
depends: parser-strings, parser-lexical-scalars
state: closed/verified

## story 6: Standard tables, redefinition, arrays of tables, and ordering.
id: parser-tables
evidence: evidence/block-6.md
summary: Standard tables, redefinition, arrays of tables, and ordering.
type: service
kind: capability
phase: 1
block: 6
implements: FEATURE-Parser-Tables.md
covers: parser-tables
accepts: st-001
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 9641
provides: table headers and arrays of tables in internal/toml
instructions: |
  Implement [table] and [[array of tables]] headers, implicit super-tables, and the redefinition rules.
  Reject defining a standard table twice and reject a [[name]] header that conflicts with a static table or array.
acceptance: yes
depends: parser-keys
state: closed/verified

## story 7: Arrays and inline tables, including inline-table closure.
id: parser-arrays-inline
evidence: evidence/block-7.md
summary: Arrays and inline tables, including inline-table closure.
type: service
kind: capability
phase: 1
block: 7
implements: FEATURE-Parser-Arrays-Inline.md
covers: parser-arrays-inline
accepts: st-001
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 9594
provides: array and inline-table parsing in internal/toml
instructions: |
  Parse arrays with optional trailing comma, nested arrays, and mixed types, including multi-line arrays with comments.
  Parse inline tables on one line with no trailing comma; reject later extension through dotted keys or table headers.
acceptance: yes
depends: parser-strings, parser-lexical-scalars, parser-keys
state: closed/verified

## story 8: Stdin, tagged JSON, diagnostic-stream, and exit-status contract.
id: decoder-contract
evidence: evidence/block-8.md
summary: Stdin, tagged JSON, diagnostic-stream, and exit-status contract.
type: service
kind: capability
phase: 1
block: 8
implements: FEATURE-Decoder-Contract.md
covers: decoder-contract
accepts: st-001
context: ARCHITECTURE_compact.md
stack: go.md, common.md
stack_mode: consumer
size: 10914
provides: toml-decoder stdin/stdout filter contract
consumes: toml.Decode
instructions: |
  Wire the stdin to tagged-JSON contract in cmd/toml-decoder.
  Encode tables as objects, arrays as arrays, and every value as {"type", "value"} with value as a JSON string.
  On invalid input write a diagnostic to stderr and exit non-zero; write nothing to stdout.
  Accept no arguments and read no configuration.
acceptance: yes
depends: architecture, parser-lexical-scalars, parser-keys, parser-arrays-inline
state: closed/verified

## story 9: Offset and local date, time, and datetime forms.
id: parser-datetimes
evidence: evidence/block-9.md
summary: Offset and local date, time, and datetime forms.
type: service
kind: capability
phase: 1
block: 9
implements: FEATURE-Parser-Datetimes.md
covers: parser-datetimes
accepts: st-001
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 9549
provides: offset datetime, local datetime, local date, and local time decoding
instructions: |
  Decode RFC 3339 offset datetimes (T or space delimiter), local datetimes, local dates, and local times.
  Truncate fractional seconds beyond supported precision; do not round.
acceptance: yes
depends: parser-lexical-scalars
state: closed/verified

## story 10: The complete unfiltered toml-test suite after all implementation stories.
id: conformance
evidence: evidence/block-10.md
summary: The complete unfiltered toml-test suite after all implementation stories.
type: feature
kind: test harness
phase: 1
block: 10
implements: FEATURE-Conformance.md
covers: conformance
accepts: st-001
context: ARCHITECTURE_compact.md
stack: go.md, common.md
stack_mode: consumer
size: 10088
instructions: |
  Assembly story. Run sh sources/full_test.sh unfiltered from the application root and assert it exits zero.
  This is the only story that runs the full suite; its acceptance declares Suite: full.
acceptance: yes
depends: architecture, decoder-contract, parser-strings, parser-lexical-scalars, parser-lexical-boundaries, parser-keys, parser-tables, parser-arrays-inline, parser-datetimes
state: closed/verified
