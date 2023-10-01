#!/bin/sh

pushd ..

cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./public
GOOS=js GOARCH=wasm go build -ldflags="-X 'main.Revision=$(git rev-parse --short HEAD)'" -o ./public/kuronan-dash.wasm

popd
