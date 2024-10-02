#!/bin/bash

if command -v go &>/dev/null; then
    echo "Go is installed"
else
    echo "Go is not installed"
    echo "Please install Go and try again"
    exit 1
fi

go build -o build/whatwords whatwords.go
