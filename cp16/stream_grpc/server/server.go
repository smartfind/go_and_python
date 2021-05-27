package main

import (
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"

	"awesomeProject_1/cp16/stream_grpc/proto"
)

const PORT = ":50052"

type server struct {
}

func (s *server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		_ = res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	return nil
}

func (s *server) PutStream(clientstr proto.Greeter_PutStreamServer) error {
	return nil
}

func (s *server) AllStream(allstr proto.Greeter_AllStreamServer) error {
	return nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
