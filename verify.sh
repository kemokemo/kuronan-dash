#!/bin/sh

verboseFlag=""
if [ "$1" = "-v" ]; then
    verboseFlag="-v"
else
    echo "If you need more detail information, please set -v flag"
fi

if [ ! -d "node_modules" ]; then
    npm i
fi
npx cspell lint '**/*.go' -c ./.cspell/cspell.cjs --gitignore --show-suggestions --no-progress
if [ $? -ne 0 ]; then
    echo "## failed to run 'cspell'. ##"
    exit $?
fi

go vet $verboseFlag ./...
if [ $? -ne 0 ]; then
    echo "## failed to run 'go vet'. ##"
    exit $?
fi

golangci-lint run
if [ $? -ne 0 ]; then
    echo "## failed to run 'golangci-lint'. ##"
    exit $?
fi

go test $verboseFlag -covermode=count ./...
if [ $? -ne 0 ]; then
    echo "## failed to run 'go test'. ##"
    exit $?
fi

go build
if [ $? -ne 0 ]; then
    echo "## failed to run 'go build'. ##"
    exit $?
fi
