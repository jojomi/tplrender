#!/bin/sh
set -ex

COVERFILE=/tmp/go-test-coverage

go test ./... -covermode=count "-coverprofile=${COVERFILE}" -v
go tool cover "-func=${COVERFILE}"
go tool cover "-html=${COVERFILE}"
