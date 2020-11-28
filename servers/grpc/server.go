package grpc

import (
	"github.com/smallnest/rpcx/server"
	"context"
	"go-websocket/conf"
)

type RpcServer struct{}


func (R *RpcServer)SendMessages(ctx context.Context, args Args, reply *Reply) error {

	



	return nil



}

func (R *RpcServer)GetUserList()  {

}

func Start() {
	s := server.NewServer()
	s.Register( new(RpcServer), "")
	err := s.Serve("tcp", GetRpcAddress())
	if err != nil {
		panic(err)
	}
}
func GetRpcAddress()  string {
	return  conf.Cfg.ServerIp+":"+conf.Cfg.RpcPort
}
