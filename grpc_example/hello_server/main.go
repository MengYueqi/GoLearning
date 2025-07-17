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
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
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
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8091")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// 创建一个gRPC server对象
	s := grpc.NewServer()
	// 注册Greeter service到server
	bookpb.RegisterBookServiceServer(s, &BookServiceImpl{})
	pb.RegisterGreeterServer(s, &server{})

	// gRPC-Gateway mux
	gwmux := runtime.NewServeMux()
	dops := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = bookpb.RegisterBookServiceHandlerFromEndpoint(context.Background(), gwmux, "127.0.0.1:8091", dops)
	if err != nil {
		log.Fatalln("Failed to register gwmux:", err)
	}
	err = pb.RegisterGreeterHandlerFromEndpoint(context.Background(), gwmux, "127.0.0.1:8091", dops)
	if err != nil {
		log.Fatalln("Failed to register gwmux:", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	// 定义HTTP server配置
	gwServer := &http.Server{
		Addr:    "127.0.0.1:8091",
		Handler: grpcHandlerFunc(s, mux), // 请求的统一入口
	}
	log.Println("Serving on http://127.0.0.1:8091")
	log.Fatalln(gwServer.Serve(lis)) // 启动HTTP服务
}

// grpcHandlerFunc 将gRPC请求和HTTP请求分别调用不同的handler处理
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}
