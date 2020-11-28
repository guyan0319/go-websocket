package grpc

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"go-websocket/conf"
	"go-websocket/servers/msgs"
	"log"
)
type Args struct {
	AppId    uint32 `json:"appId,omitempty"`
	UserId   string `json:"userId,omitempty"`
	Data string `json:"data,omitempty"`
}
type Reply struct {

}
func SendMessages(server *msgs.Server,appId uint32,userId,data string,)  {
	xclient:=RpcConnect("RpcServer",server.Ip)
	defer xclient.Close()
	args := Args{
		AppId:appId,
		UserId:userId,
		Data:data,
	}
	reply := &Reply{}
	err := xclient.Call(context.Background(), "SendMessages", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
}

func RpcConnect(servicePath ,ip string) client.XClient {
	d := client.NewPeer2PeerDiscovery("tcp@"+ip+":"+conf.Cfg.RpcPort, "")
	opt := client.DefaultOption
	opt.SerializeType = protocol.JSON
	return  client.NewXClient(servicePath, client.Failtry, client.RandomSelect, d, opt)
}