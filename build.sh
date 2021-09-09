#!/bin/bash
SRC_DIR=$(pwd)
export GOPATH=$SRC_DIR
export GOROOT=/usr/lib/go

go build -o go_server main