package grpc

import (
	"context"
	"flag"
	"fmt"
	"github.com/smallnest/rpcx/server"
	"go-websocket/conf"
)

type RpcServer struct{}

func (R *RpcServer)SendMessages(ctx context.Context, args Args, reply *Reply) error {
	//ws.SendMessages(args.AppId, args.UserId, args.Data)
	return nil
}

//func (R *RpcServer)GetUserList()  {
//
//}

func Start() {
	flag.Parse()
	s := server.NewServer()
	s.Register( new(RpcServer), "")
	err := s.Serve("tcp",GetRpcAddress())
	if err != nil {
		panic(err)
	}
	fmt.Println("rpc服务启动")
}
func GetRpcAddress()  string {
	return  conf.Cfg.RpcIp+":"+conf.Cfg.RpcPort
}
