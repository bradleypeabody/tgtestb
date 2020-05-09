#!/bin/bash

echo "Compiling wasm with tinygo..."
tinygo build -o main.wasm -target=wasm main_wasm.go

echo "Copying wasm_exec.js into place..."
cp `tinygo env TINYGOROOT`/targets/wasm_exec.js .

echo "Running devserver..."
go run devserver.go
