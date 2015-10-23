package main

import (
	"errors"
	"github.com/babelrpc/lib-go/babel"
	"github.com/babelrpc/lib-go/examples/babelOverJsonRpc/gen"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"sync"
)

type UserServiceImpl struct {
	Map     map[int32]*gen.User
	Counter int32
	m       sync.Mutex
	mm      sync.RWMutex
}

func NewUserServiceImpl() *UserServiceImpl {
	x := new(UserServiceImpl)
	x.Map = make(map[int32]*gen.User)
	return x
}

// Get a user given the ID
// ID:  User ID
func (svc *UserServiceImpl) GetUser(ID *int32) (*gen.User, error) {
	if ID == nil {
		return nil, errors.New("ID parameter cannot be null")
	}
	log.Printf("GetUser %d", *ID)
	svc.mm.RLock()
	defer svc.mm.RUnlock()
	return svc.Map[*ID], nil
}

// Add a user
// user:  Details of the user to add
func (svc *UserServiceImpl) AddUser(user *gen.User) (*int32, error) {
	if user == nil {
		return nil, errors.New("user parameter cannot be null")
	}
	svc.m.Lock()
	svc.Counter++
	x := svc.Counter
	svc.m.Unlock()
	user.Id = &x
	svc.mm.Lock()
	defer svc.mm.Unlock()
	svc.Map[x] = user
	log.Printf("AddUser %d", x)
	return &x, nil
}

// Clears the list
func (svc *UserServiceImpl) Clear() error {
	svc.m.Lock()
	svc.Counter = 0
	svc.m.Unlock()
	svc.mm.Lock()
	defer svc.mm.Unlock()
	svc.Map = make(map[int32]*gen.User)
	log.Printf("Clear")
	return nil
}

func main() {
	// Create Babel service objects
	svc := new(gen.UserService)
	svc.SvcObj = NewUserServiceImpl()

	// Register the service with RPC
	rpc.Register(svc)

	// Register the service with Babel
	babel.Register(svc)

	// Set up Babel HTTP handlers and serve HTTP
	babel.HandleHTTP()
	http.Handle("/test/", http.StripPrefix("/test/", http.FileServer(http.Dir("../test"))))
	go func() { log.Fatal(http.ListenAndServe(":8333", nil)) }()

	//rpc.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)

	// set up network listener for json rpc
	l, e := net.Listen("tcp", ":8222")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go jsonrpc.ServeConn(conn)
	}
}
