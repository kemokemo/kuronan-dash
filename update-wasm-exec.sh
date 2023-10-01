#!/bin/sh

export goVersion=1.21.1

wget https://github.com/golang/go/archive/refs/tags/go$goVersion.zip
unzip go$goVersion.zip

cp ./go-go$goVersion/misc/wasm/wasm_exec.js ./public/

rm go$goVersion.zip
rm -rf go-go$goVersion/
