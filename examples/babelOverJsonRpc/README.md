Babel for Go Example
====================

Babel for Go supports Go's builtin [jsonrpc package](http://golang.org/pkg/net/rpc/jsonrpc/) as well as Babel's HTTP format.

- Use `b.sh` to run the Babel commands (you may need to fix [test/services/UserService.json](test/services/UserService.json) afterward due to a bug with test harness generation - there is a trailing comma).
- Go to the [gen](gen) folder and run `go install`
- Go to the [server](server) folder and run `go run babjs.go` to start the server.
- Go to the [client](client) folder and run `go run babjc.go` to exercise a json-rpc client.
- Point your browser at [http://localhost:8333/test/](http://localhost:8333/test/) to try Babel's test harness.
