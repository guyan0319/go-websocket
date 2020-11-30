package main

import (
	"context"
	"log"

	"github.com/smallnest/rpcx/protocol"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/client"
)



func main() {

	d := client.NewPeer2PeerDiscovery("tcp@"+"192.168.9.93:8972", "")
	opt := client.DefaultOption
	opt.SerializeType = protocol.JSON

	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, opt)
	defer xclient.Close()

	args := example.Args{
		A: 10,
		B: 20,
	}

	reply := &example.Reply{}
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)

}