#!/usr/bin/env bash
CURDIR=$(cd `dirname $0`; pwd)
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

gofmt -w src

go install xMgr

export GOPATH="$OLDGOPATH"

echo 'finished'