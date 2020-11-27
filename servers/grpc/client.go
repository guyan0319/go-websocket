package grpc

import (
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

func SendMessages()  {





}

func RpcConnect() client.XClient {
	d := client.NewPeer2PeerDiscovery("tcp@"+GetRpcAddress(), "")
	opt := client.DefaultOption
	opt.SerializeType = protocol.JSON
	return  client.NewXClient("Server", client.Failtry, client.RandomSelect, d, opt)
}