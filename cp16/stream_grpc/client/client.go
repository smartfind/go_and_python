package main

import "C"
import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"awesomeProject_1/cp16/stream_grpc/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	//服务端流模式
	c := proto.NewGreeterClient(conn)
	res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "慕课网"})
	for {
		a, err := res.Recv() //如果大家懂socket编程的话就明白 send recv
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a.Data)
	}

}
