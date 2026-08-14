# FEATURE: Tagged JSON

| Field       | Value |
|-------------|-------|
| Version     | 20260814 V1 |
| Description | Defines tagged JSON encoding for TOML tables, arrays, and scalar values. |
| Depends On  | FEATURE-Decoder-Contract.md |
| Provides    | tagged JSON encoding |
| Consumes    | parser value model, cmd/toml-decoder |

## Encoding Rules

Tables become JSON objects and arrays become JSON arrays. Every scalar value becomes an object containing a `type` field and a string-valued `value` field. Integer, float, and boolean payloads remain strings rather than native JSON numbers or booleans.

The encoder supports the scalar type names `string`, `integer`, `float`, `bool`, `datetime`, `datetime-local`, `date-local`, and `time-local`.

## Guardrails

- Empty tables encode as `{}`.
- Empty arrays encode as `[]`.
- Nested tables and arrays retain their structure.
- Every scalar payload is a JSON string.
- Offset datetimes use RFC 3339 representation.
- Local datetimes omit an offset.
- Local dates and local times use their respective representations.

## Programmatic Acceptance

=== AC tagged-scalars ===
Intent: Scalar TOML values are emitted with type names and string payloads.
import json
import subprocess

source = 'text = "hello"\ncount = 42\nactive = true\n'
result = subprocess.run(
    ["./toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
decoded = json.loads(result.stdout)
assert result.returncode == 0
assert decoded["text"]["type"] == "string"
assert decoded["text"]["value"] == "hello"
assert decoded["count"]["type"] == "integer"
assert decoded["count"]["value"] == "42"
assert decoded["active"]["type"] == "bool"
assert decoded["active"]["value"] == "true"
=== END AC tagged-scalars ===

=== AC tagged-structures ===
Intent: Tables and arrays retain structural JSON forms while their scalar members remain tagged.
import json
import subprocess

source = 'items = [1, 2]\n[meta]\nname = "x"\n'
result = subprocess.run(
    ["./toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
decoded = json.loads(result.stdout)
assert result.returncode == 0
assert isinstance(decoded["items"], list)
assert isinstance(decoded["meta"], dict)
assert decoded["items"][0]["type"] == "integer"
assert decoded["meta"]["name"]["type"] == "string"
=== END AC tagged-structures ===

=== AC tagged-empty-containers ===
Intent: Empty arrays and tables encode as empty JSON containers.
import json
import subprocess

source = 'empty_array = []\n[empty_table]\n'
result = subprocess.run(
    ["./toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
decoded = json.loads(result.stdout)
assert result.returncode == 0
assert decoded["empty_array"] == []
assert decoded["empty_table"] == {}
=== END AC tagged-empty-containers ===

=== AC tagged-datetime-types ===
Intent: Date and time values use the distinct required tagged type names.
import json
import subprocess

source = (
    'offset = 1979-05-27T07:32:00Z\n'
    'local = 1979-05-27T07:32:00\n'
    'date = 1979-05-27\n'
    'time = 07:32:00\n'
)
result = subprocess.run(
    ["./toml-decoder"],
    input=source,
    capture_output=True,
    text=True,
)
decoded = json.loads(result.stdout)
assert result.returncode == 0
assert decoded["offset"]["type"] == "datetime"
assert decoded["local"]["type"] == "datetime-local"
assert decoded["date"]["type"] == "date-local"
assert decoded["time"]["type"] == "time-local"
=== END AC tagged-datetime-types ===

## User Acceptance

- None.

## Guardrails

- None.
