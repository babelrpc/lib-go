// Added monitoring page
package main

import (
	"github.com/ancientlore/kubismus"
	"github.com/babelrpc/lib-go/babel"
	"github.com/babelrpc/lib-go/examples/babeldemo/harness"
	"github.com/babelrpc/lib-go/examples/babeldemo/impl"
	"github.com/babelrpc/lib-go/examples/babeldemo/math"
	"log"
	"net/http"
)

func main() {

	// Register the first service
	floatSvc := new(math.FloatService)
	floatSvc.SvcObj = new(impl.FloatServiceImpl)
	babel.Register(floatSvc)

	// Register the second service
	fractSvc := new(math.FractionService)
	fractSvc.SvcObj = new(impl.FractionServiceImpl)
	babel.Register(fractSvc)

	// Set some info for the monitoring page
	kubismus.Setup("Math Service", "")
	kubismus.Note("Babel base path", "http://localhost:9999"+babel.DefaultHttpPath)
	kubismus.Note("Monitoring path", "http://localhost:9999"+kubismus.DefaultPath)
	kubismus.Note("Test harness path", "http://localhost:9999/test/")
	kubismus.Define("Requests", kubismus.COUNT, "Requests/sec")

	// Set up Babel HTTP handlers and serve HTTP
	http.Handle(babel.DefaultHttpPath, kubismus.HttpRequestMetric("Requests", babel.DefaultServer))
	kubismus.HandleHTTP()
	http.HandleFunc("/test/", harness.ServeHTTP)
	http.HandleFunc("/", frame)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func frame(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<!DOCTYPE html>
<html>
<body>
<iframe src="/kubismus/" width="1024px" height="160px">
<p>Your browser does not support iframes.</p>
</iframe>
<br/>
<iframe src="/test/" width="1024px" height="440px">
<p>Your browser does not support iframes.</p>
</iframe>
</body>
</html>`))
}
