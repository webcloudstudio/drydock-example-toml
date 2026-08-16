# FEATURE: Lexical Scalars

| Field       | Value |
|-------------|-------|
| Version     | 20260816 V1 |
| Description | Defines TOML boolean, integer, and floating-point scalar parsing. |
| Depends On  | ARCHITECTURE.md |
| Provides    | TOML boolean, integer, and float parsing |
| Consumes    | internal/toml |

## Scope

Support lowercase booleans; signed decimal integers; hexadecimal, octal, and binary integers; underscores; signed 64-bit bounds; decimal fractions and exponents; signed zero; infinity; and NaN. Scalar values must be emitted with the corresponding TOML type tag and a JSON string representation.

## Programmatic Acceptance

=== AC scalars-valid-suite ===
Intent: The authoritative conformance suite passes all valid boolean, integer, and float cases.

import subprocess

result = subprocess.run(
    [
        "sh",
        "sources/stage_test.sh",
        "valid/bool/**|valid/integer/**|valid/float/**",
    ],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC scalars-valid-suite ===

=== AC scalars-invalid-suite ===
Intent: The authoritative conformance suite rejects all invalid boolean, integer, and float cases.

import subprocess

result = subprocess.run(
    [
        "sh",
        "sources/stage_test.sh",
        "invalid/bool/**|invalid/integer/**|invalid/float/**",
    ],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC scalars-invalid-suite ===

=== AC scalars-bounded-suites ===
Intent: The valid and invalid scalar conformance slices each complete successfully.

import subprocess

patterns = [
    "valid/bool/**|valid/integer/**|valid/float/**",
    "invalid/bool/**|invalid/integer/**|invalid/float/**",
]
for pattern in patterns:
    result = subprocess.run(
        ["sh", "sources/stage_test.sh", pattern],
        capture_output=True,
        text=True,
    )
    print(result.stdout)
    print(result.stderr)
    assert result.returncode == 0
=== END AC scalars-bounded-suites ===

## User Acceptance

- None.

## Guardrails

- Integer parsing must reject values outside the signed 64-bit range.
- Integer underscores must satisfy TOML placement rules.
- Floats must reject malformed decimal points and exponents.
- Special float tokens must remain lowercase and preserve the required type semantics.
