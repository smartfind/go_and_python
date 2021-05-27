package main

import (
	"awesomeProject_1/cp15/new_helloeworld/client_proxy"
	"fmt"
)

func main() {
	//1、建立连接
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")
	//只想写业务逻辑，不想关注每个函数名称
	//客户端部分
	var reply *string = new(string)
	err := client.Hello("xiaoming", reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(*reply)
}
