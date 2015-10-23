/*
	Package babel provides access to the exported methods of an object across an
	http connection.  A server registers an object, making it visible
	as a service with the name of the type of the object.  After registration, exported
	methods of the object will be accessible remotely.  A server may register multiple
	objects (services) of different types but it is an error to register multiple
	objects of the same type.

	Only methods that satisfy these criteria will be made available for remote access;
	other methods will be ignored:

		- the method is exported.
		- the method has two arguments, both exported (or builtin) types.
		- the method's second argument is a pointer.
		- the method has return type error.

	In effect, the method must look schematically like

		func (t *T) MethodName(argType T1, replyType *T2) error

	where T, T1 and T2 can be marshaled by encoding/json.

	NOTE: Half of this code is adapted from Go's rpc/jsonrpc package. The goal was to be
	able to use Babel as either a normal Babel HTTP server or a Go jsonrpc server.
*/
package babel

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"time"
	"unicode"
	"unicode/utf8"
)

// This generates the error model
//go:generate babel -lang go -model -output $GOPATH/src $GOPATH/etc/babeltemplates/error.babel

const (
	// Defaults used by HandleHTTP
	DefaultHttpPath = "/_babel_/"
)

// Precompute the reflect type for error.  Can't use error directly
// because Typeof takes an empty interface value.  This is annoying.
var typeOfError = reflect.TypeOf((*error)(nil)).Elem()

type methodType struct {
	method    reflect.Method
	ArgType   reflect.Type
	ReplyType reflect.Type
}

type service struct {
	name   string                 // name of service
	rcvr   reflect.Value          // receiver of methods for the service
	typ    reflect.Type           // type of the receiver
	method map[string]*methodType // registered methods
}

// Server represents a Babel Server.
type Server struct {
	mu         sync.RWMutex // protects the serviceMap
	serviceMap map[string]*service
}

// NewServer returns a new Server.
func NewServer() *Server {
	return &Server{serviceMap: make(map[string]*service)}
}

// DefaultServer is the default instance of *Server.
var DefaultServer = NewServer()

// Is this an exported - upper case - name?
func isExported(name string) bool {
	rune, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(rune)
}

// Is this type exported or a builtin?
func isExportedOrBuiltinType(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// PkgPath will be non-empty even for an exported type,
	// so we need to check the type name as well.
	return isExported(t.Name()) || t.PkgPath() == ""
}

// Register publishes in the server the set of methods of the
// receiver value that satisfy the following conditions:
//	- exported method
//	- two arguments, both of exported type
//	- the second argument is a pointer
//	- one return value, of type error
// It returns an error if the receiver is not an exported type or has
// no suitable methods. It also logs the error using package log.
// The client accesses each method using a string of the form "Type.Method",
// where Type is the receiver's concrete type.
func (server *Server) Register(rcvr interface{}) error {
	return server.register(rcvr, "", false)
}

// RegisterName is like Register but uses the provided name for the type
// instead of the receiver's concrete type.
func (server *Server) RegisterName(name string, rcvr interface{}) error {
	return server.register(rcvr, name, true)
}

func (server *Server) register(rcvr interface{}, name string, useName bool) error {
	server.mu.Lock()
	defer server.mu.Unlock()
	if server.serviceMap == nil {
		server.serviceMap = make(map[string]*service)
	}
	s := new(service)
	s.typ = reflect.TypeOf(rcvr)
	s.rcvr = reflect.ValueOf(rcvr)
	sname := reflect.Indirect(s.rcvr).Type().Name()
	if useName {
		sname = name
	}
	if sname == "" {
		s := "babel.Register: no service name for type " + s.typ.String()
		log.Print(s)
		return errors.New(s)
	}
	if !isExported(sname) && !useName {
		s := "babel.Register: type " + sname + " is not exported"
		log.Print(s)
		return errors.New(s)
	}
	if _, present := server.serviceMap[sname]; present {
		return errors.New("babel: service already defined: " + sname)
	}
	s.name = sname

	// Install the methods
	s.method = suitableMethods(s.typ, true)

	if len(s.method) == 0 {
		str := ""

		// To help the user, see if a pointer receiver would work.
		method := suitableMethods(reflect.PtrTo(s.typ), false)
		if len(method) != 0 {
			str = "babel.Register: type " + sname + " has no exported methods of suitable type (hint: pass a pointer to value of that type)"
		} else {
			str = "babel.Register: type " + sname + " has no exported methods of suitable type"
		}
		log.Print(str)
		return errors.New(str)
	}
	server.serviceMap[s.name] = s
	log.Print("babel: registered service " + s.name)
	return nil
}

// suitableMethods returns suitable Babel invoker methods of typ, it will report
// error using log if reportErr is true. The code generator produces methods
// that can be used with Babel or Go's builting jsonrpc package.
func suitableMethods(typ reflect.Type, reportErr bool) map[string]*methodType {
	methods := make(map[string]*methodType)
	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		mtype := method.Type
		mname := method.Name
		// Method must be exported.
		if method.PkgPath != "" {
			continue
		}
		// Method needs three ins: receiver, *args, *reply.
		if mtype.NumIn() != 3 {
			if reportErr {
				log.Println("method", mname, "has wrong number of ins:", mtype.NumIn())
			}
			continue
		}
		// First arg need not be a pointer.
		argType := mtype.In(1)
		if !isExportedOrBuiltinType(argType) {
			if reportErr {
				log.Println(mname, "argument type not exported:", argType)
			}
			continue
		}
		// Second arg must be a pointer.
		replyType := mtype.In(2)
		if replyType.Kind() != reflect.Ptr {
			if reportErr {
				log.Println("method", mname, "reply type not a pointer:", replyType)
			}
			continue
		}
		// Reply type must be exported.
		if !isExportedOrBuiltinType(replyType) {
			if reportErr {
				log.Println("method", mname, "reply type not exported:", replyType)
			}
			continue
		}
		// Method needs one out.
		if mtype.NumOut() != 1 {
			if reportErr {
				log.Println("method", mname, "has wrong number of outs:", mtype.NumOut())
			}
			continue
		}
		// The return type of the method must be error.
		if returnType := mtype.Out(0); returnType != typeOfError {
			if reportErr {
				log.Println("method", mname, "returns", returnType.String(), "not error")
			}
			continue
		}
		methods[mname] = &methodType{method: method, ArgType: argType, ReplyType: replyType}
	}
	return methods
}

// Register publishes the receiver's methods in the DefaultServer.
func Register(rcvr interface{}) error {
	return DefaultServer.Register(rcvr)
}

// RegisterName is like Register but uses the provided name for the type
// instead of the receiver's concrete type.
func RegisterName(name string, rcvr interface{}) error {
	return DefaultServer.RegisterName(name, rcvr)
}

func badRequest(w http.ResponseWriter, err error) {
	BabelError(w, "bad request: "+err.Error(), http.StatusBadRequest)
	//http.Error(w, "400 bad request", http.StatusBadRequest)
}

func BabelError(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	errModel := &ServiceError{}
	errModel.Init()
	when := time.Now()
	errModel.Time = &when
	errInfo := &Error{}
	errInfo.Init()
	errInfo.Message = &msg
	errModel.Errors = append(errModel.Errors, errInfo)
	buf, err := json.Marshal(errModel)
	if err != nil {
		fmt.Fprintf(w, "%d %s", code, msg)
	} else {
		w.Write(buf)
	}
}

// ServeHTTP implements an http.Handler that answers Babel requests.
func (server *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Print(req.URL.Path)
	if req.Method != "POST" {
		http.NotFound(w, req)
		return
	}
	arr := strings.Split(req.URL.Path, "/")
	if len(arr) < 2 {
		http.NotFound(w, req)
		return
	}
	arr = arr[len(arr)-2:]
	server.mu.RLock()
	service := server.serviceMap[arr[0]]
	server.mu.RUnlock()
	if service == nil {
		http.NotFound(w, req)
		return
	}
	mtype := service.method[arr[1]]
	if mtype == nil {
		http.NotFound(w, req)
		return
	}
	function := mtype.method.Func

	// Decode the argument value.
	var argv, replyv reflect.Value
	argIsValue := false // if true, need to indirect before calling.
	if mtype.ArgType.Kind() == reflect.Ptr {
		argv = reflect.New(mtype.ArgType.Elem())
	} else {
		argv = reflect.New(mtype.ArgType)
		argIsValue = true
	}
	// argv guaranteed to be a pointer now.

	buf, err := ioutil.ReadAll(req.Body)
	//log.Printf("%s REQUEST: %s", req.URL.Path, string(buf))
	if err != nil {
		badRequest(w, err)
		return
	}
	err = json.Unmarshal(buf, argv.Interface())
	if err != nil {
		badRequest(w, err)
		return
	}
	if argIsValue {
		argv = argv.Elem()
	}

	replyv = reflect.New(mtype.ReplyType.Elem())
	// need to call Init
	init := replyv.MethodByName("Init")
	init.Call([]reflect.Value{})

	// Invoke the method, providing a new value for the reply.
	returnValues := function.Call([]reflect.Value{service.rcvr, argv, replyv})

	// The return value for the method is an error.
	errInter := returnValues[0].Interface()
	errmsg := ""
	if errInter != nil {
		errmsg = errInter.(error).Error()
		BabelError(w, errmsg, http.StatusInternalServerError)
	} else {
		fld := reflect.Indirect(replyv).FieldByName("Value")
		if fld.IsValid() {
			buf, err = json.Marshal(fld.Interface())
			if err != nil {
				BabelError(w, err.Error(), http.StatusInternalServerError)
			} else {
				//log.Printf("%s RESPONSE: %s", req.URL.Path, string(buf))
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Write(buf)
			}
		}
	}
}

// HandleHTTP registers an HTTP handler for Babel messages on babelPath.
// It is still necessary to invoke http.Serve(), typically in a go statement.
func (server *Server) HandleHTTP(babelPath string) {
	http.Handle(babelPath, server)
	log.Print("babel: handling HTTP requests on " + babelPath)
}

// HandleHTTP registers an HTTP handler for RPC messages to DefaultServer
// on DefaultRPCPath and a debugging handler on DefaultDebugPath.
// It is still necessary to invoke http.Serve(), typically in a go statement.
func HandleHTTP() {
	DefaultServer.HandleHTTP(DefaultHttpPath)
}
