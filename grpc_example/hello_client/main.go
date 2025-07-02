package main

import (
	"context"
	"flag"
	"log"
	"time"

	author "github.com/testProject/pb/author"
	bookbp "github.com/testProject/pb/book"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// hello_client

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "127.0.0.1:8972", "the address to connect to")
	name = flag.String("name", "ExName", "Name to greet")
)

func main() {
	flag.Parse()
	// 连接到server端，此处禁用安全传输
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := bookbp.NewBookServiceClient(conn)

	// 执行RPC调用并打印收到的响应数据
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 构造请求参数（BookMessage）
	req := &bookbp.Book{Title: "BookExTitle",
		Price:      &bookbp.Price{MarketPrice: 49, SalePrice: 29},
		AuthorInfo: &author.Info{Name: "Mr. Zhang"}}
	//r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	r, err := c.Create(ctx, req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Result)
}
