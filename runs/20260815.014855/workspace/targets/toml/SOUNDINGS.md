# Soundings

Per-assertion acceptance board, one row per Blueprint Programmatic Acceptance check.
`drydock plan` projects every assertion as `— UNVERIFIED`; `drydock score ac` sets the
Status column from deterministic proof results. A rerun of `drydock plan` resets Status to
`— UNVERIFIED`; rescore to refresh.

| Status | Blueprint | AC Id | Text | Evidence | Verified At |
|---|---|---|---|---|---|
| ✓ PASS | ARCHITECTURE.md | architecture-module | The project has a standard-library-only Go module and the fixed decoder build entry point. |  | 2026-08-15T02:18:54+00:00 |
| ✓ PASS | ARCHITECTURE.md | architecture-parser-boundary | The decoder command and internal parser boundary compile together. |  | 2026-08-15T02:18:54+00:00 |
| ~ PREPASSED | FEATURE-Decoder-Contract.md | decoder-valid-contract | A valid TOML document produces tagged JSON on stdout without stderr diagnostics. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-15T02:18:54+00:00 |
| ~ PREPASSED | FEATURE-Decoder-Contract.md | decoder-invalid-contract | Malformed TOML produces a non-zero exit, stderr diagnostics, and no stdout document. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-15T02:18:54+00:00 |
| ~ PREPASSED | FEATURE-Decoder-Contract.md | decoder-stage-contract | The supplied decoder contract gate passes for valid and malformed stdin documents. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-15T02:18:54+00:00 |
| ~ PREPASSED | FEATURE-Parser-Strings.md | parser-strings-conformance | The authoritative TOML conformance suite passes every valid and invalid string case owned by this story. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-15T02:18:54+00:00 |
| ~ PREPASSED | FEATURE-Parser-Lexical-Scalars.md | parser-lexical-scalars-conformance | The authoritative TOML conformance suite passes every valid and invalid boolean, integer, and float case owned by this story. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-15T02:18:54+00:00 |
| ~ PREPASSED | FEATURE-Parser-Lexical-Boundaries.md | parser-lexical-boundaries-conformance | The authoritative TOML conformance suite passes every valid comment and invalid control-character and encoding case owned by this story. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-15T02:18:54+00:00 |
| ~ PREPASSED | FEATURE-Parser-Keys.md | parser-keys-conformance | The implementation passes the complete authoritative key conformance slice. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-15T02:18:54+00:00 |
| ~ PREPASSED | FEATURE-Parser-Tables.md | parser-tables-conformance | The implementation passes the complete authoritative table conformance slice. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-15T02:18:54+00:00 |
| ~ PREPASSED | FEATURE-Parser-Datetimes.md | parser-datetimes-conformance | The implementation passes the complete authoritative date/time conformance slice. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-15T02:18:54+00:00 |
| ~ PREPASSED | FEATURE-Parser-Arrays-Inline.md | parser-arrays-inline-conformance | The implementation passes the complete authoritative array and inline-table conformance slice. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-15T02:18:54+00:00 |
| ~ PREPASSED | FEATURE-Conformance.md | complete-conformance | The completed implementation passes the complete authoritative TOML v1.0.0 conformance suite. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-15T02:18:54+00:00 |
