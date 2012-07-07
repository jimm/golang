#!/bin/bash

cd $(dirname $0)
export GOPATH=$(pwd)
go install $(basename $GOPATH)
bin/$(basename $GOPATH) $*
