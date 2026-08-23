# AUTHORITATIVE PROJECT METADATA — FIELDS SHOULD BE CURRENT

name: Toml
display_name: drydock TOML Uat Case
short_description: Drydock uat - Toml proves that the Drydock framework, itself written in Python, drives a build in a language other than Python.  The kit consists of the TOML v1.0.0 specification as Markdown, the upstream toml-test conformance harness, and shell scripts that install it and score the decoder.  TECHNOLOGY_STACK.md specifies Go, and the build is governed by Rigging/stack/go.md; any other language works the same way once you supply its Rigging file (copy the python.md format).  It also proves staged scoring: architecture, decoder contract, and seven parser phases are each graded on their own slice of the suite before the full conformance run, and rejecting malformed input is graded as heavily as accepting valid input, so a merely permissive parser fails outright.  Run with drydock uat Toml.  Source Repo: github.com/toml-lang/toml and github.com/toml-lang/toml-test.  
stack: Go
version: 
build_state: 
build_sub_state: 
last_analyzed: 
last_planned: 
last_built: 
build_dir: 
brand: 
git_repo: 
release_tag: 
