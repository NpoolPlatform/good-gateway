#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

go get -u github.com/ugorji/go/codec@latest
go mod tidy -compat=1.17
git diff --exit-code go.*
