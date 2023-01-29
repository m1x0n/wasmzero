MAKEFLAGS += --silent

.DEFAULT_GOAL = help

NAME = wasmzero: WASM animation example with ebiten engine

.PHONY: help ## Shows this help
help:
	echo "${NAME}\n\n"
	echo "List of available make targets:\n"
	@printf "%-19s %s\n" "Target" "Description"
	@printf "%-19s %s\n" "------" "-----------"
	@grep '^.PHONY: .* ##' Makefile | sed 's/\.PHONY: \(.*\) ## \(.*\)/\1	\2/' | sort | expand -t20

.PHONY: build ## Builds binary
build:
	echo "Not implemented"

.PHONY: wasm ## Builds wasm
wasm:
	GOOS=js GOARCH=wasm go build -v -o wasmzero.wasm .

.PHONY: run ## Runs application without build
run:
	go run main.go