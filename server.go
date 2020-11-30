package main

import (
	"context"
	"flag"
	"fmt"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "192.168.9.93:8972", "server address")
)

type Arith struct{}

// the second parameter is not a pointer
func (t *Arith) Mul(ctx context.Context, args example.Args, reply *example.Reply) error {
	reply.C = args.A * args.B
	fmt.Println("C=", reply.C)
	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer()
	s.Register(new(Arith), "")
	//s.RegisterName("Arith", new(Arith), "")
	fmt.Println(*addr)

	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}