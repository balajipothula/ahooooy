#!/bin/bash

# Author      : Balaji Pothula <balan.pothula@gmail.com>,
# Date        : Thursday, 28 August 2025,
# Description : go file compile commands.

# get index page
#
# GOOS=linux    : build for Linux
# GOARCH=amd64  : build for x86_64 (change if targeting ARM, etc.)
# CGO_ENABLED=0 : ensures no libc/glibc dependencies
# -s : Omit symbol table and debug information
# -w : Omit DWARF symbol table
# go tool link --help
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main main.go


# --best : Tells UPX to use the maximum compression ratio
# --lzma : Forces UPX to use the LZMA compression algorithm.
upx --best --lzma main
