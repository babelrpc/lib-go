Babel and Go
============

An initial version of Babel for Go supports server-side development. Some caveats:

- `decimal` and `int64` do not serialize correctly yet. They are supposed to have quotes. The Go serializer doesn't do this correctly for pointer types.
- `datetime` serializes with nanoseconds (too many digits). For example, "2006-01-02T15:04:05.999999999Z07:00". Perhaps we should support this because it's super accurate.
- `enum` is supported but shows up in the data structure as a string. There are handy functions generated to make between the string values and constants representing the enum values.
- `decimal` is not yet supported. I thought it would map nicely to `big.Rat`, but that doesn't seem to the same type.
- No client is generated yet.
- There isn't library infrastructure to support anything other than JSON. This is partly due to being able to support Go's jsonrpc package.
- The code doesn't yet figure out which imports are needed. Go doesn't allow extra imports, so when using multiple Babel files or includes you are likely to run into issues.
- Code is plaed into folders based on the namespace. It's recommended to pass `-output $GOPATH/src` on the command line. Thus, a namespace of `github.com/me/Foo` would end up with files in `$GOPATH/src/github.com/me/foo`.
- Biggest caveat: not much testing has been done yet.

Babel for Go generates a class that can be used with a Babel server or with Go's built-in [jsonrpc package](http://golang.org/pkg/net/rpc/jsonrpc/).

Sample Application
------------------

See the [demo application](examples/babeldemo) or [JSON-RPC sample application](examples/babelOverJsonRpc) for more details. The sample includes a jsonrpc client (see `babjc.go`) and a server that supports both Babel HTTP and jsonrpc (see `babjs.go`). It also serves the test harness.

Steps to get started
--------------------

- You'll need to rebuild Babel using [build.sh](build.sh)
- Until a time issue is fixed, you'll need to edit [babel/errorModel.go](babel/errorModel.go) and uncomment the `"time"` import.
- Go to [babel](babel) and run `go install`.

Interesting Notes
-----------------

- I will probably code it up to call `go fmt` at some point.
- I was surprised how several types just serialized right out of the box. Byte arrays (`[]byte`) drop right into the base64 format we use. Times are obvious (although have too much precision per our spec).
- In Go it's a pain to have everything generated as a pointer so that it can be nullable. It certainly makes it consistent with Babel's philosophy, but the code is yuckier.

Have fun.

