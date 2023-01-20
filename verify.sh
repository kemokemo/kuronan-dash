#!/bin/sh

cspell lint '**/*.go' -c ./.cspell/cspell.cjs --gitignore --show-suggestions --no-progress
if [ $? -ne 0 ]; then
    echo "## failed to run 'cspell'. ##"
    exit $?
fi

go vet -v ./...
if [ $? -ne 0 ]; then
    echo "## failed to run 'go vet'. ##"
    exit $?
fi

golangci-lint run
if [ $? -ne 0 ]; then
    echo "## failed to run 'golangci-lint'. ##"
    exit $?
fi

go test -v -covermode=count ./...
if [ $? -ne 0 ]; then
    echo "## failed to run 'go test'. ##"
    exit $?
fi

go build
if [ $? -ne 0 ]; then
    echo "## failed to run 'go build'. ##"
    exit $?
fi
