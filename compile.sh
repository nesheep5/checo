#!/bin/bash                                                                                                                                                                                                                                                                                                         
gox \
    -os="darwin linux windows" \
    -arch="386 amd64" \
    -output "pkg/{{.OS}}_{{.Arch}}/{{.Dir}}" \
    ./cmd/checo