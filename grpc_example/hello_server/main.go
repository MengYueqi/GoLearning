package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	fieldmask_utils "github.com/mennanov/fieldmask-utils"
	pb "github.com/testProject/pb"
	_ "github.com/testProject/pb/author"
	bookpb "github.com/testProject/pb/book"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
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
	if bookMsg.GetIdx() == nil {
		return &bookpb.BookCreateResponse{Result: "The Book Name is: " + bookMsg.Title +
			"\nThe Book Author is: " + bookMsg.AuthorInfo.GetName() +
			"\nThe sale Price is: " + strconv.Itoa(int(bookMsg.Price.GetSalePrice())) + "\nThe index is: UNKNOWN!",
		}, nil
	} else {
		return &bookpb.BookCreateResponse{Result: "The Book Name is: " + bookMsg.Title +
			"\nThe Book Author is: " + bookMsg.AuthorInfo.GetName() +
			"\nThe sale Price is: " + strconv.Itoa(int(bookMsg.Price.GetSalePrice())) + "\nThe index is: " + bookMsg.Idx.GetValue(),
		}, nil
	}
}

func (s *BookServiceImpl) Update(ctx context.Context, bookUpdateMsg *bookpb.BookUpdateMsg) (*bookpb.BookUpdateResponse, error) {
	mask, _ := fieldmask_utils.MaskFromProtoFieldMask(bookUpdateMsg.UpdateMask, generator.CamelCase)
	var bookDst = make(map[string]interface{})
	// 将数据读取到map[string]interface{}
	// fieldmask-utils支持读取到结构体等，更多用法可查看文档。
	err := fieldmask_utils.StructToMap(mask, bookUpdateMsg.Book, bookDst)
	if err != nil {
		return &bookpb.BookUpdateResponse{Result: "Update Book Error: " + err.Error()}, nil
	}
	// do update with bookDst
	resultMsg := fmt.Sprintf("bookDst:%#v\n", bookDst)
	fmt.Println(resultMsg)
	return &bookpb.BookUpdateResponse{Result: resultMsg}, nil
}

func (s *BookServiceImpl) GetHotBooks(stream bookpb.BookService_GetHotBooksServer) error {
	hotBooksName := []string{"ABCEnglish", "AAAChinese", "CCCMath"}
	reply := ""
	for {
		// 接收客户端发来的流式数据
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		reply += res.GetRequest()
	}
	for _, word := range hotBooksName {
		// 使用Send方法返回多个数据
		if err := stream.Send(&bookpb.HotBooksResponse{BookName: reply + word}); err != nil {
			return err
		}
		//time.Sleep(1 * time.Second)
	}
	return nil
}

func main() {
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", "localhost:8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	// 加载证书
	//creds, err := credentials.NewServerTLSFromFile("../cert/server.crt", "../cert/server.key")
	//if err != nil {
	//	log.Fatalf("Failed to generate credentials %v", err)
	//}
	////s := grpc.NewServer(grpc.Creds(creds))                  // 创建gRPC服务器
	// 创建一个gRPC server对象
	s := grpc.NewServer()
	bookpb.RegisterBookServiceServer(s, &BookServiceImpl{}) // 在gRPC服务端注册服务
	pb.RegisterGreeterServer(s, &server{})
	// 启动服务
	go func() {
		err = s.Serve(lis)
		if err != nil {
			fmt.Printf("failed to serve: %v", err)
			return
		}
	}()
	// 创建一个连接到我们刚刚启动的 gRPC 服务器的客户端连接
	// gRPC-Gateway 就是通过它来代理请求（将HTTP请求转为RPC请求）
	conn, err := grpc.NewClient(
		"localhost:8972",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// 注册Greeter
	err = pb.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	err = bookpb.RegisterBookServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register BookService gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}
	// 8090端口提供gRPC-Gateway服务
	log.Println("Serving gRPC-Gateway on http://192.168.0.1:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
