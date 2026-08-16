# Soundings

Per-assertion acceptance board, one row per Blueprint Programmatic Acceptance check.
`drydock plan` projects every assertion as `— UNVERIFIED`; `drydock score ac` sets the
Status column from deterministic proof results. A rerun of `drydock plan` resets Status to
`— UNVERIFIED`; rescore to refresh.

| Status | Blueprint | AC Id | Text | Evidence | Verified At |
|---|---|---|---|---|---|
| ✓ PASS | ARCHITECTURE.md | architecture-module | The project has a compilable standard-library-only Go module and decoder command. |  | 2026-08-16T03:26:02+00:00 |
| ✓ PASS | ARCHITECTURE.md | architecture-boundary | The decoder command is buildable at its declared path. |  | 2026-08-16T03:26:02+00:00 |
| ✓ PASS | FEATURE-Decoder-Contract.md | decoder-valid-contract | Valid stdin produces the required tagged JSON state and no diagnostic. |  | 2026-08-16T03:26:02+00:00 |
| ✓ PASS | FEATURE-Decoder-Contract.md | decoder-invalid-contract | Malformed stdin is rejected without a successful stdout payload. |  | 2026-08-16T03:26:02+00:00 |
| ✗ FAIL | FEATURE-Decoder-Contract.md | decoder-no-arguments | The command boundary rejects unsupported command-line arguments. | assert result.returncode != 0 → AssertionError | 2026-08-16T03:26:02+00:00 |
| ✓ PASS | FEATURE-Parser-Strings.md | parser-strings-valid | The authoritative TOML suite passes every valid string case. |  | 2026-08-16T03:26:02+00:00 |
| ~ PREPASSED | FEATURE-Parser-Strings.md | parser-strings-invalid | The authoritative TOML suite rejects every invalid string case. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-16T03:26:02+00:00 |
| ✓ PASS | FEATURE-Parser-Strings.md | parser-strings-contract | String parsing remains reachable through the stdin decoder boundary. |  | 2026-08-16T03:26:02+00:00 |
| ~ PREPASSED | FEATURE-Parser-Lexical-Scalars.md | parser-scalars-valid | The authoritative TOML suite passes every valid boolean, integer, and float case. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-16T03:26:02+00:00 |
| ~ PREPASSED | FEATURE-Parser-Lexical-Scalars.md | parser-scalars-invalid | The authoritative TOML suite rejects every invalid boolean, integer, and float case. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-16T03:26:02+00:00 |
| ~ PREPASSED | FEATURE-Parser-Lexical-Scalars.md | parser-scalars-integer-contract | A signed 64-bit integer is emitted without numeric precision loss. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-16T03:26:02+00:00 |
| ✓ PASS | FEATURE-Parser-Lexical-Boundaries.md | parser-boundaries-valid | The authoritative TOML suite passes every valid comment and boundary case. |  | 2026-08-16T03:26:02+00:00 |
| ✓ PASS | FEATURE-Parser-Lexical-Boundaries.md | parser-boundaries-invalid-control | The authoritative TOML suite rejects every invalid control-character case. |  | 2026-08-16T03:26:02+00:00 |
| ✓ PASS | FEATURE-Parser-Lexical-Boundaries.md | parser-boundaries-invalid-encoding | The authoritative TOML suite rejects every invalid UTF-8 encoding case. |  | 2026-08-16T03:26:02+00:00 |
| ~ PREPASSED | FEATURE-Parser-Keys.md | parser-keys-conformance | The implementation passes the authoritative TOML key cases owned by this story. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-16T03:26:02+00:00 |
| ✓ PASS | FEATURE-Parser-Keys.md | parser-keys-invalid-extension | A scalar key cannot later be extended as a dotted table path. |  | 2026-08-16T03:26:02+00:00 |
| ~ PREPASSED | FEATURE-Parser-Keys.md | parser-keys-duplicate | Repeating an already-defined key is rejected with a nonzero exit status. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-16T03:26:02+00:00 |
| ~ PREPASSED | FEATURE-Parser-Tables.md | parser-tables-conformance | The implementation passes the authoritative TOML table and array-of-table cases owned by this story. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-16T03:26:02+00:00 |
| ✓ PASS | FEATURE-Parser-Tables.md | parser-tables-redefinition | Redefining a standard table is rejected. |  | 2026-08-16T03:26:02+00:00 |
| ~ PREPASSED | FEATURE-Parser-Tables.md | parser-tables-array-conflict | Defining an array of tables where a static table already exists is rejected. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-16T03:26:02+00:00 |
| ~ PREPASSED | FEATURE-Parser-Arrays-Inline.md | parser-arrays-inline-conformance | The implementation passes the authoritative array and inline-table cases owned by this story. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-16T03:26:02+00:00 |
| ~ PREPASSED | FEATURE-Parser-Arrays-Inline.md | parser-arrays-trailing-comma | Arrays permit a trailing comma and preserve their parsed elements. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-16T03:26:02+00:00 |
| ✓ PASS | FEATURE-Parser-Arrays-Inline.md | parser-inline-closure | An inline table cannot be extended by a later assignment. |  | 2026-08-16T03:26:02+00:00 |
| ~ PREPASSED | FEATURE-Parser-Datetimes.md | parser-datetimes-conformance | The implementation passes the authoritative TOML datetime cases owned by this story. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-16T03:26:02+00:00 |
| ~ PREPASSED | FEATURE-Parser-Datetimes.md | parser-datetimes-types | Offset and local date/time inputs receive the declared tagged types. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-16T03:26:02+00:00 |
| ✓ PASS | FEATURE-Parser-Datetimes.md | parser-datetimes-invalid | Invalid local date/time syntax is rejected with a nonzero exit status. |  | 2026-08-16T03:26:02+00:00 |
| ✓ PASS | FEATURE-Conformance.md | conformance-full | The completed implementation passes the complete authoritative TOML 1.0.0 suite. |  | 2026-08-16T03:26:02+00:00 |
| ~ PREPASSED | FEATURE-Conformance.md | conformance-harness-integrity | The supplied full-test entry point remains present and executable as a shell script. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-16T03:26:02+00:00 |
| ~ PREPASSED | FEATURE-Conformance.md | conformance-decoder-build | The decoder builds successfully through the same Go command used by the scoring harness. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-16T03:26:02+00:00 |
