#!/bin/bash

OS="darwin linux"
ARCH="amd64"

for GOOS in $OS; do
    for GOARCH in $ARCH; do
        OS="$(tr '[:lower:]' '[:upper:]' <<< ${GOOS:0:1})${GOOS:1}"
        architecture="${OS}-${GOARCH}"
        echo "Building ${architecture}"
        export GOOS=$GOOS
        export GOARCH=$GOARCH
        go get
        go build -o bin/gone-${architecture}
    done
done
