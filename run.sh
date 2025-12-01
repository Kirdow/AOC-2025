#!/bin/bash

set -e

go mod tidy
go build -o ./aoc cmd/cli/main.go
./aoc
