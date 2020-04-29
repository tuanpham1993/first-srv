package main

import (
	"context"
	"fmt"

	proto "github.com/tuanpham1993/first-srv/proto"

	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/micro/go-micro/v2"
)

type Say struct{}

func (w *Say) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (w *Say) All(ctx context.Context, req *empty.Empty, rsp *proto.Response) error {
	fmt.Println("All from srv")
	rsp.Msg = "Hello everyone"
	return nil
}

func main() {
	fmt.Println("foo")
	service := micro.NewService(
		micro.Name("go.micro.srv.wc"),
	)

	service.Init()
	// proto.Re
	proto.RegisterSayHandler(service.Server(), new(Say))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("bar")
}
