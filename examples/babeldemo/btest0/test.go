// Basic test of math service with no HTTP
package main

import (
	"fmt"
	"github.com/babelrpc/lib-go/examples/babeldemo/impl"
	"github.com/babelrpc/lib-go/examples/babeldemo/math"
)

func main() {
	// Create the first service
	floatSvc := new(math.FloatService)
	floatSvc.SvcObj = new(impl.FloatServiceImpl)

	// Create request
	req := new(math.FloatServiceAddRequest)
	req.Init()
	A := 3.4
	B := 6.2
	req.A = &A
	req.B = &B

	// Create response
	rsp := new(math.FloatServiceAddResponse)
	rsp.Init()

	// Add
	err := floatSvc.Add(req, rsp)
	if err != nil {
		fmt.Printf("Oops, error: %s\n", err)
	} else {
		fmt.Printf("Value = %f\n", *rsp.Value)
	}
}
