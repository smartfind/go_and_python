package main

import (
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
	_ = rpc.RegisterName("HelloService", &HelloService{})
	//3、启动服务
	conn, _ := listener.Accept() //当一个新的连接进来的时候
	rpc.ServeConn(conn)

	//一连串代码大部分都是net的包，似乎和rpc没有关系
	//直接rpc包不行，实现rpc调用需要解决几个问题：1、call id；2、序列化和反序列化
	//可以跨语言调用：1、go语言的的rpc的序列化协议是什么（GOB）；2、能否替换成常见的序列化
}
