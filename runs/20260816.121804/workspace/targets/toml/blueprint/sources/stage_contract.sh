#!/bin/sh
# stage_contract.sh — governed gates for boundaries not represented by one toml-test slice.
set -eu

if [ $# -ne 1 ]; then
  echo "usage: stage_contract.sh <architecture|decoder>" >&2
  exit 2
fi

case "$1" in
  architecture)
    test -f go.mod
    if grep -Eq '^[[:space:]]*require[[:space:](]' go.mod; then
      echo "go.mod declares a forbidden third-party dependency" >&2
      exit 1
    fi
    go build ./cmd/toml-decoder
    ;;
  decoder)
    go build -o toml-decoder ./cmd/toml-decoder
    stage_dir="$(mktemp -d)"
    trap 'rm -rf "$stage_dir"' EXIT HUP INT TERM

    printf 'answer = 42\n' | ./toml-decoder >"$stage_dir/valid.json" 2>"$stage_dir/valid.err"
    test ! -s "$stage_dir/valid.err"
    python - "$stage_dir/valid.json" <<'PY'
import json
import sys

with open(sys.argv[1], encoding="utf-8") as stream:
    decoded = json.load(stream)
assert decoded == {"answer": {"type": "integer", "value": "42"}}
PY

    if printf 'answer = [\n' | ./toml-decoder >"$stage_dir/invalid.out" 2>"$stage_dir/invalid.err"; then
      echo "malformed TOML exited successfully" >&2
      exit 1
    fi
    test ! -s "$stage_dir/invalid.out"
    test -s "$stage_dir/invalid.err"
    ;;
  *)
    echo "unknown stage contract: $1" >&2
    exit 2
    ;;
esac
