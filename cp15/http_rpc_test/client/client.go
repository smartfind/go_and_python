package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	//1、建立连接
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		panic("连接失败")
	}
	var reply *string = new(string) //或者两边同时为string，string有默认值，指针默认值是nil
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call("HelloService.Hello", "xiaoming", reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(*reply)
}
