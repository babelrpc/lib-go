package main

import (
	"fmt"
	"github.com/babelrpc/lib-go/examples/babelOverJsonRpc/gen"
	"net"
	"net/rpc/jsonrpc"
	"time"
)

type BabelCaller interface {
	Call(serviceMethod string, args interface{}, reply interface{}) error
	Close() error
	//Go(serviceMethod string, args interface{}, reply interface{}, done chan *rpc.Call) *rpc.Call
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8222")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	u := &gen.User{}
	u.Init()
	u.Name = new(string)
	*u.Name = "Michael Lore"
	u.Age = new(int32)
	*u.Age = 47
	u.EmailAddress = new(string)
	*u.EmailAddress = "michael.lore@gmail.com"

	req := &gen.UserServiceAddUserRequest{User: u}
	req.Init()
	reply := &gen.UserServiceAddUserResponse{}
	reply.Init()

	c := jsonrpc.NewClient(conn)

	cr := &gen.UserServiceClearRequest{}
	cr.Init()
	br := &gen.UserServiceClearResponse{}
	br.Init()

	err = c.Call("UserService.Clear", cr, br)
	if err != nil {
		fmt.Printf("Clear Error %s\n", err)
	}

	for i := 0; i < 10; i++ {
		t0 := time.Now()
		err = c.Call("UserService.AddUser", req, reply)
		t1 := time.Now()
		fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
		if err != nil {
			fmt.Printf("ERROR %s\n", err)
		} else {
			fmt.Printf("Added user %d\n", *reply.Value)
		}
	}

	ar := &gen.UserServiceGetUserRequest{}
	ar.Id = new(int32)
	*ar.Id = 1
	ars := &gen.UserServiceGetUserResponse{}

	err = c.Call("UserService.GetUser", ar, ars)
	if err != nil {
		fmt.Printf("ERROR %s\n", err)
	} else {
		fmt.Printf("Got user %#v\n", ars.Value)
	}

}
