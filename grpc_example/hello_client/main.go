package main

import (
	"context"
	"flag"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	pb "github.com/testProject/pb"
	bookbp "github.com/testProject/pb/book"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"
)

// hello_client

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "127.0.0.1:8091", "the address to connect to")
	name = flag.String("name", "ExName", "Name to greet")
)

func getHotBooksName(c bookbp.BookServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 客户端流式RPC
	stream, err := c.GetHotBooks(ctx)
	if err != nil {
		log.Fatalf("c.LotsOfGreetings failed, err: %v", err)
	}
	names := []string{"七米", "q1mi", "沙河娜扎"}
	for _, name := range names {
		// 发送流式数据
		err := stream.Send(&bookbp.HotBooksRequest{Request: name})
		if err != nil {
			log.Fatalf("c.LotsOfGreetings stream.Send(%v) failed, err: %v", name, err)
		}
	}
	stream.CloseSend()
	//in, err := stream.Recv()
	//stream, err = c.GetHotBooks(ctx, &bookbp.HotBooksRequest{Request: "Meng"})
	//if err != nil {
	//	log.Fatalf("c.LotsOfReplies failed, err: %v", err)
	//}
	for {
		// 接收服务端返回的流式数据，当收到io.EOF或错误时退出
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("c.LotsOfReplies failed, err: %v", err)
		}
		log.Printf("got reply: %q\n", res.GetBookName())
	}
}

func main() {
	//flag.Parse()
	//// 连接到server端，此处禁用安全传输
	//conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//defer conn.Close()
	//c := bookbp.NewBookServiceClient(conn)
	//
	//// 执行RPC调用并打印收到的响应数据
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//// 构造请求参数（BookMessage）
	//req := &bookbp.Book{
	//	Title:      "BookExTitle",
	//	Price:      &bookbp.Price{MarketPrice: 49, SalePrice: 29},
	//	AuthorInfo: &author.Info{Name: "Mr. Zhang"},
	//	//Idx:        &wrapperspb.StringValue{Value: "s123abc"},
	//}
	////r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	//r, err := c.Create(ctx, req)
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Greeting: %s", r.Result)
	//
	//updateField := []string{"title"}
	//reqUpdate := &bookbp.BookUpdateMsg{
	//	UpdateMask: &fieldmaskpb.FieldMask{Paths: updateField},
	//	Book: &bookbp.Book{
	//		Title:      "BookExTitle_Update",
	//		Price:      &bookbp.Price{MarketPrice: 49, SalePrice: 29},
	//		AuthorInfo: &author.Info{Name: "Mr. Zhang"},
	//	},
	//}
	//rU, err := c.Update(ctx, reqUpdate)
	//
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Greeting: %s", rU.Result)
	//creds, err := credentials.NewClientTLSFromFile("../cert/server.crt", "")
	//if err != nil {
	//	log.Fatalf("Failed to create TLS credentials %v", err)
	//}
	//conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(
		// consul服务
		"consul://localhost:8500/BookAndHello?healthy=true",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	// 执行RPC调用并打印收到的响应数据
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Meng"})
	if err != nil {
		log.Fatalf("SayHello failed, err: %v", err)
	} else {
		fmt.Println(r)
	}
}
