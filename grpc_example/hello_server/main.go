package main

import (
	"context"
	"fmt"
	pb "github.com/testProject/pb"
	_ "github.com/testProject/pb/author"
	bookpb "github.com/testProject/pb/book"
	"google.golang.org/grpc"
	"net"
)

// hello server

type server struct {
	pb.UnimplementedGreeterServer
}

// BookServiceImpl 是你的服务实现
type BookServiceImpl struct {
	bookpb.UnimplementedBookServiceServer // 嵌入 UnimplementedBookServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Reply: "Hello " + in.Name}, nil
}

func (s *BookServiceImpl) Create(ctx context.Context, bookMsg *bookpb.Book) (*bookpb.BookCreateResponse, error) {
	return &bookpb.BookCreateResponse{Result: "The Book Name is: " + bookMsg.Title}, nil
}

func main() {
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()                                   // 创建gRPC服务器
	bookpb.RegisterBookServiceServer(s, &BookServiceImpl{}) // 在gRPC服务端注册服务
	// 启动服务
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
