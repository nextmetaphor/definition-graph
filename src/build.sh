#!/bin/sh

export CGO_ENABLED=1

#apk add build-base
#
if [ "$#" -gt 0 ]; then
    export GOOS=$1
fi
if [ "$#" -gt 1 ]; then
    export GOARCH=$2
fi

printf "### go fmt ###\n"
go fmt ./...

printf "\n### go vet ###\n"
#go vet ./...

printf "\n### go build ###\n"
go build -o grfn