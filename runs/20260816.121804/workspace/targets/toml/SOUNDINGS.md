# Soundings

Per-assertion acceptance board, one row per Blueprint Programmatic Acceptance check.
`drydock plan` projects every assertion as `— UNVERIFIED`; `drydock score ac` sets the
Status column from deterministic proof results. A rerun of `drydock plan` resets Status to
`— UNVERIFIED`; rescore to refresh.

| Status | Blueprint | AC Id | Text | Evidence | Verified At |
|---|---|---|---|---|---|
| ✓ PASS | ARCHITECTURE.md | architecture-build | The declared Go module and decoder command compile without third-party module requirements. |  | 2026-08-16T13:04:11+00:00 |
| ✓ PASS | ARCHITECTURE.md | architecture-module | The command package is buildable through the required repository path. |  | 2026-08-16T13:04:11+00:00 |
| ✓ PASS | ARCHITECTURE.md | architecture-dependencies | The module declares no forbidden third-party dependency. |  | 2026-08-16T13:04:11+00:00 |
| ✓ PASS | FEATURE-Decoder-Contract.md | decoder-valid-contract | A valid TOML document produces the required tagged JSON on stdout with no stderr diagnostic. |  | 2026-08-16T13:04:11+00:00 |
| ✓ PASS | FEATURE-Decoder-Contract.md | decoder-invalid-contract | Invalid TOML produces a non-zero exit, no successful stdout result, and a stderr diagnostic. |  | 2026-08-16T13:04:11+00:00 |
| ✓ PASS | FEATURE-Decoder-Contract.md | decoder-stage-contract | The supplied decoder contract gate passes in full. |  | 2026-08-16T13:04:11+00:00 |
| ✓ PASS | FEATURE-Strings.md | strings-valid-suite | The authoritative conformance suite passes every valid string case owned by this feature. |  | 2026-08-16T13:04:11+00:00 |
| ✓ PASS | FEATURE-Strings.md | strings-invalid-suite | The authoritative conformance suite rejects every invalid string case owned by this feature. |  | 2026-08-16T13:04:11+00:00 |
| ✓ PASS | FEATURE-Strings.md | strings-bounded-suites | Both string conformance invocations complete successfully as separate executable checks. |  | 2026-08-16T13:04:11+00:00 |
| ✓ PASS | FEATURE-Keys.md | keys-valid | The decoder passes all authoritative valid TOML key cases owned by this feature. |  | 2026-08-16T13:04:11+00:00 |
| ✓ PASS | FEATURE-Keys.md | keys-invalid | The decoder rejects all authoritative invalid TOML key cases owned by this feature. |  | 2026-08-16T13:04:11+00:00 |
| ~ PREPASSED | FEATURE-Lexical-Scalars.md | scalars-valid-suite | The authoritative conformance suite passes all valid boolean, integer, and float cases. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-16T13:04:11+00:00 |
| ~ PREPASSED | FEATURE-Lexical-Scalars.md | scalars-invalid-suite | The authoritative conformance suite rejects all invalid boolean, integer, and float cases. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-16T13:04:11+00:00 |
| ~ PREPASSED | FEATURE-Lexical-Scalars.md | scalars-bounded-suites | The valid and invalid scalar conformance slices each complete successfully. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-16T13:04:11+00:00 |
| ✓ PASS | FEATURE-Lexical-Boundaries.md | lexical-comments-valid-suite | The authoritative conformance suite passes every valid comment and lexical-boundary case. |  | 2026-08-16T13:04:11+00:00 |
| ~ PREPASSED | FEATURE-Lexical-Boundaries.md | lexical-boundaries-invalid-suite | The authoritative conformance suite rejects every invalid control-character and encoding case. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-16T13:04:11+00:00 |
| ✓ PASS | FEATURE-Lexical-Boundaries.md | lexical-boundaries-suites | The valid comment and invalid lexical-boundary conformance slices complete successfully. |  | 2026-08-16T13:04:11+00:00 |
| ✓ PASS | FEATURE-Tables.md | tables-valid | The decoder passes all authoritative valid TOML table and array-of-table cases owned by this feature. |  | 2026-08-16T13:04:11+00:00 |
| ✓ PASS | FEATURE-Tables.md | tables-invalid | The decoder rejects all authoritative invalid TOML table and array-of-table cases owned by this feature. |  | 2026-08-16T13:04:11+00:00 |
| ✓ PASS | FEATURE-Datetimes.md | datetimes-valid | The decoder passes all authoritative valid TOML datetime cases owned by this feature. |  | 2026-08-16T13:04:11+00:00 |
| ~ PREPASSED | FEATURE-Datetimes.md | datetimes-invalid | The decoder rejects all authoritative invalid TOML datetime, local-date, local-datetime, and local-time cases owned by this feature. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-16T13:04:11+00:00 |
| ~ PREPASSED | FEATURE-Arrays-Inline.md | arrays-valid | The decoder passes all authoritative valid TOML array cases owned by this feature. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-16T13:04:11+00:00 |
| ~ PREPASSED | FEATURE-Arrays-Inline.md | inline-valid | The decoder passes all authoritative valid TOML inline-table cases owned by this feature. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-16T13:04:11+00:00 |
| ~ PREPASSED | FEATURE-Arrays-Inline.md | arrays-invalid | The decoder rejects all authoritative invalid TOML array cases owned by this feature. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-16T13:04:11+00:00 |
| ~ PREPASSED | FEATURE-Arrays-Inline.md | inline-invalid | The decoder rejects all authoritative invalid TOML inline-table cases owned by this feature. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-16T13:04:11+00:00 |
| ~ PREPASSED | FEATURE-Conformance.md | conformance-full | The completed decoder passes the supplied unfiltered TOML 1.0.0 conformance suite. | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-08-16T13:04:11+00:00 |
