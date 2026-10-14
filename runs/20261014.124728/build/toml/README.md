# toml

A TOML v1.0.0 decoder that reads TOML on stdin and writes tagged JSON to stdout, written in Go with the standard library only and verified against the toml-test conformance suite.

## Intent

A standard-library Go decoder for TOML v1.0.0. It reads a TOML document on stdin and writes the tagged-JSON description of it to stdout, so that conformance can be judged by the upstream toml-test suite. It serves anyone needing a faithful, dependency-free TOML reader, and it is judged by the exit status of the supplied scoring script. Rejecting invalid input counts as much as accepting valid input.

## What It Does

This is the only story that runs the complete suite. It runs `sh sources/full_test.sh` from the application root with no filter and takes its exit status as the verdict. The exit status is the whole verdict: the summary text and case counts are not asserted.

The decoder reads TOML from stdin and writes tagged JSON to stdout with exit status 0. On invalid TOML it writes a diagnostic to stderr, writes nothing to stdout, and exits non-zero. It takes no arguments and reads no configuration.

Parse arrays of mixed value types, with nesting, optional trailing comma, and multi-line layout with comments. Parse inline tables on a single line with no trailing comma. An inline table is closed: keys cannot be added to it through dotted keys or table headers.

Decode RFC 3339 offset date-times, accepting either `T` or a space as the delimiter. Decode local date-times, local dates, and local times without an offset. Fractional seconds beyond the supported precision are truncated, not rounded.

Parse bare keys (`A-Za-z0-9_-`), quoted keys (basic or literal strings), and dotted keys that create intermediate tables. Reject defining the same key twice. A bare key and a quoted key with the same name are equivalent. A key that holds an integer cannot later be used as a table.

A hash begins a comment that runs to end of line, except inside a string. Control characters other than tab are rejected in comments and literal strings. Input that is not valid UTF-8 is rejected.

Decode booleans, decimal and prefixed integers (`0x`, `0o`, `0b`) with underscores, and floats with fractional and exponent parts, plus the special floats `inf`, `nan`, and their signed forms. Integers must fit signed 64-bit and are rejected otherwise.

Decode basic strings with the listed escapes only, multi-line basic strings with newline trimming and the line-ending backslash rule, literal strings with no escaping, and multi-line literal strings. Reject any other backslash escape and any control character not permitted by the specification.

Parse `[table]` headers and `[[array]]` headers. Create implicit super-tables for `[a.b.c]`. Each `[[name]]` header appends a new table element to the array named `name`, and later sub-tables attach to the most recent element. Reject a standard table defined twice, and reject a `[[name]]` header that conflicts with a static table or static array.

## Next Steps

- Complete the first-run checks above.
