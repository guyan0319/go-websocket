package grpc

import (
	"github.com/smallnest/rpcx/server"
	"go-websocket/conf"
)

type Server struct{}

func (S *Server)SendMessages()  {

}

func (S *Server)GetUserList()  {

}

func Start() {
	s := server.NewServer()
	s.Register( new(Server), "")
	err := s.Serve("tcp", GetRpcAddress())
	if err != nil {
		panic(err)
	}
}
func GetRpcAddress()  string {
	return  conf.Cfg.ServerIp+":"+conf.Cfg.RpcPort
}
