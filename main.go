package main

import (
	"context"
	proto "github.com/tuanpham1993/firstsrv/proto"
	"fmt"

	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/micro/go-micro/v2"
)

type Say struct{}

func (w *Say) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (w *Say) All(ctx context.Context, req *empty.Empty, rsp *proto.Response) error {
	rsp.Msg = "Hello everyone"
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.welcome"),
	)

	// proto.Re
	proto.RegisterSayHandler(service.Server(), new(Say))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
