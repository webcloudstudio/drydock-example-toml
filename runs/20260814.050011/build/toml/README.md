# toml

A Go command-line parser that converts TOML 1.0.0 from stdin into tagged JSON and rejects invalid documents.

## Intent

Build a Go command-line TOML 1.0.0 parser for users and harnesses that need a deterministic stdin-to-tagged-JSON decoder. The parser must implement the supplied specification and distinguish valid documents from invalid documents through its output and exit behavior.

## What It Does

The parser enforces TOML 1.0.0 ordering and conflict rules for arrays of tables.

The parser recognizes `[[table]]` headers and creates ordered table elements. Each subsequent key/value pair applies to the newest element. Nested standard tables and nested arrays of tables target the current enclosing element.

The completed Go decoder passes every valid and invalid case in the installed `toml-test` v2.2.0 TOML 1.0 suite.

## Stack

Go

## Next Steps

- Complete the first-run checks above.
