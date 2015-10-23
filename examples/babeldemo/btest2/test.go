// Added embedded test harness
package main

import (
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

	// Set up Babel HTTP handlers and serve HTTP
	babel.HandleHTTP()
	http.HandleFunc("/test/", harness.ServeHTTP)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
