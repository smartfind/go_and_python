package main

import (
	"awesomeProject_1/cp15/new_helloeworld/handler"
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	//返回值是通过修改reply的值
	*reply = "hello, " + request
	return nil
}

func main() {
	//1、实例化一个server
	listener, _ := net.Listen("tcp", ":1234")
	//2、注册逻辑，handler
	_ = rpc.RegisterName(handler.HelloServiceName, &HelloService{})
	//3、启动服务
	for {
		conn, _ := listener.Accept() //当一个新的连接进来的时候
		go rpc.ServeConn(conn)
	}

}
