#!/bin/bash

# Author      : Balaji Pothula <balan.pothula@gmail.com>,
# Date        : Thursday, 28 August 2025,
# Description : go file compile commands.

# get index page
#
# GOOS=linux    : build for Linux
# GOARCH=amd64  : build for x86_64 (change if targeting ARM, etc.)
# CGO_ENABLED=0 : ensures no libc/glibc dependencies
# -buildvcs     : Whether to stamp binaries with version control information.
# -ldflags      : arguments to pass on each go tool link invocation.
## -s : Omit symbol table and debug information
## -w : Omit DWARF symbol table
# -tags         : a comma-separated list of additional build tags to consider satisfied during the build.
## netgo    : forces Go’s pure-Go DNS resolver
## osusergo : avoids cgo lookups for user/group DB.
# -trimpath     : remove all file system paths from the resulting executable.
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs=false -ldflags="-s -w" -tags netgo,osusergo -trimpath -o main main.go

# Ultimate Packer for eXecutables
#
# --best           : compress best (can be slow for big files)
# -f               : force compression of suspicious files
# --lzma           : enabling LZMA (Lempel–Ziv–MArkov chain) algorithm
# --preserve-build : copy .gnu.note.build-id to compressed output
# -q               : be quiet
# --no-backup      : no backup files
# -t               : test compressed file
upx --best -f --lzma --preserve-build -q --no-backup -t main
