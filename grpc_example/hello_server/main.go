package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	fieldmask_utils "github.com/mennanov/fieldmask-utils"
	pb "github.com/testProject/pb"
	_ "github.com/testProject/pb/author"
	bookpb "github.com/testProject/pb/book"
	"google.golang.org/grpc"
	"net"
	"strconv"
	"time"
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

func (s *BookServiceImpl) GetHotBooks(booksReq *bookpb.HotBooksRequest, stream bookpb.BookService_GetHotBooksServer) error {
	hotBooksName := []string{"ABCEnglish", "AAAChinese", "CCCMath"}
	for _, word := range hotBooksName {
		// 使用Send方法返回多个数据
		if err := stream.Send(&bookpb.HotBooksResponse{BookName: booksReq.Request + word}); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
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
