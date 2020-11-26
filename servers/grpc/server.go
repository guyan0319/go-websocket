package grpc

import (
	"flag"
	"github.com/smallnest/rpcx/server"
)
var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

type Arith struct{}
func Start()  {


	s := server.NewServer()
	//s.Register(new(Arith), "")
	s.RegisterName("Arith", new(Arith), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}


}