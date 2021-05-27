package server_proxy

import (
	"awesomeProject_1/cp15/new_helloeworld/handler"
	"net/rpc"
)

func RegisterHelloService(srv handler.HelloService) error {
	return rpc.RegisterName(handler.HelloServiceName, srv)
}
