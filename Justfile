
example ARG:
    go run {{ARG}}

run:
    go run main.go

all-examples:
    #!/usr/bin/env bash -euxo pipefail
    for file in $(ls examples); do
        go run "examples/$file"
    done