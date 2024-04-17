#!/bin/sh
# ===========================================================================
# File: build.sh
# Description: usage: ./build.sh [outdir]
# ===========================================================================

# exit when any command fails
set -e

. ./build_init.sh

OUTPUT_DIR=$(mkdir_output "$1")
OUTPUT_BINARY=$OUTPUT_DIR/bb

echo "Start building bb ${VERSION}..."

flags="-X 'github.com/bytebase/cli/cmd/version.version=${VERSION}'
-X 'github.com/bytebase/cli/cmd/version.goversion=$(go version)'
-X 'github.com/bytebase/cli/cmd/version.gitcommit=$(git rev-parse HEAD)'
-X 'github.com/bytebase/cli/cmd/version.buildtime=$(date -u +"%Y-%m-%dT%H:%M:%SZ")'
-X 'github.com/bytebase/cli/cmd/version.builduser=$(id -u -n)'"

# -ldflags="-w -s" means omit DWARF symbol table and the symbol table and debug information
go build --tags "release" -ldflags "-w -s $flags" -o ${OUTPUT_BINARY} ./main.go

echo "Completed building bb."

echo ""
echo "Printing version..."

${OUTPUT_BINARY} version
