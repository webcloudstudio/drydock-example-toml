<!-- Compacted from stage_contract.sh sha256=c95f3a314d9d8099bf27a64404a7eba213710c2948d41c1ce6c44f87f225227b on 2026-08-14 by drydock build agent -->

`stage_contract.sh architecture` requires dependency-free `go.mod` and successful decoder compilation. `stage_contract.sh decoder` verifies `answer = 42` produces tagged JSON, malformed TOML exits non-zero, stdout stays empty on failure, and stderr is non-empty.
