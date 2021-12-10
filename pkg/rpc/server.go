package rpc

import (
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
)

type Server struct {
	addr      string
	functions map[string]reflect.Value
}

func NewServer(addr string) *Server {
	return &Server{addr: addr, functions: make(map[string]reflect.Value)}
}

func (s *Server) Register(fnName string, fFunc interface{}) {
	if _, ok := s.functions[fnName]; ok {
		return
	}

	s.functions[fnName] = reflect.ValueOf(fFunc)
}

func (s *Server) Execute(req RPCdata) RPCdata {
	f, ok := s.functions[req.Name]
	if !ok {
		e := fmt.Sprintf("func %s not Registered", req.Name)
		log.Println(e)
		return RPCdata{Name: req.Name, Args: nil, Err: e}
	}

	log.Printf("func %s is called\n", req.Name)
	inArgs := make([]reflect.Value, len(req.Args))
	for i := range req.Args {
		inArgs[i] = reflect.ValueOf(req.Args[i])
	}

	out := f.Call(inArgs)
	resArgs := make([]interface{}, len(out)-1)
	for i := 0; i < len(out)-1; i++ {
		resArgs[i] = out[i].Interface()
	}

	var er string
	if _, ok := out[len(out)-1].Interface().(error); ok {
		er = out[len(out)-1].Interface().(error).Error()
	}
	return RPCdata{Name: req.Name, Args: resArgs, Err: er}
}

func (s *Server) Run() {
	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Printf("listen on %s err: %v\n", s.addr, err)
		return
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("accept err: %v\n", err)
			continue
		}
		go func() {
			newSession := NewSession(conn)
			for {
				req, err := newSession.Read()
				if err != nil {
					if err != io.EOF {
						log.Printf("read err: %v\n", err)
						return
					}
				}

				decReq, err := Decode(req)
				if err != nil {
					log.Printf("Error Decoding the Payload err: %v\n", err)
					return
				}
				resP := s.Execute(decReq)
				b, err := Encode(resP)
				if err != nil {
					log.Printf("Error Encoding the Payload for response err: %v\n", err)
					return
				}
				err = newSession.Send(b)
				if err != nil {
					log.Printf("transport write err: %v\n", err)
				}
			}
		}()
	}
}
