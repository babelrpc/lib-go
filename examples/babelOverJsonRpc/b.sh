#!/bin/bash

babel -output $GOPATH/src -lang go -inc "*.babel"
babel -output $GOPATH/src/github.com/babelrpc/lib-go/examples/babelOverJsonRpc/test/services -lang test -inc "*.babel"
