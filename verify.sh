#!/bin/sh

go vet -v ./...
if [ $? -ne 0 ]; then
    echo "failed to run 'go vet'."
    exit $?
fi

golangci-lint run
if [ $? -ne 0 ]; then
    echo "failed to run 'golangci-lint'."
    exit $?
fi

go test -v ./...
if [ $? -ne 0 ]; then
    echo "failed to run 'go test'."
    exit $?
fi

go build
if [ $? -ne 0 ]; then
    echo "failed to run 'go build'."
    exit $?
fi
