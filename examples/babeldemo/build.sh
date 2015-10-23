#!/bin/bash

echo "*** Compiling Babel files ***"
babel -output $GOPATH/src -lang go "*.babel"
babel -output $GOPATH/src/github.com/babelrpc/lib-go/examples/babeldemo/test/services -lang test "*.babel"

echo "*** Packaging test harness files ***"
binder -package harness test/* test/*/* > harness/testharness.go

echo "*** Building generated code ***"

pushd math >/dev/null
go install
popd >/dev/null

pushd harness >/dev/null
go install
popd >/dev/null

pushd impl >/dev/null
go install
popd >/dev/null
