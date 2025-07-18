package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hashicorp/consul/api"
	fieldmask_utils "github.com/mennanov/fieldmask-utils"
	pb "github.com/testProject/pb"
	_ "github.com/testProject/pb/author"
	bookpb "github.com/testProject/pb/book"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
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

// consul 定义一个consul结构体，其内部有一个`*api.Client`字段。
type consul struct {
	client *api.Client
}

// NewConsul 连接至consul服务返回一个consul对象
func NewConsul(addr string) (*consul, error) {
	cfg := api.DefaultConfig()
	cfg.Address = addr
	c, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &consul{c}, nil
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

	// 注册健康检查服务
	healthcheck := health.NewServer()
	healthpb.RegisterHealthServer(s, healthcheck)

	conImpl, err := NewConsul("127.0.0.1:8500")
	if err != nil {
		log.Fatalln(err)
	}
	localIP, err := GetOutboundIP()
	if err != nil {
		log.Fatalln(err)
	}
	conImpl.RegisterService("BookAndHello", localIP.String(), 8091)

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

// GetOutboundIP 获取本机的出口IP
func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}

// RegisterService 将gRPC服务注册到consul
func (c *consul) RegisterService(serviceName string, ip string, port int) error {
	localIP, err := GetOutboundIP()
	if err != nil {
		log.Fatalln(err)
	}
	check := &api.AgentServiceCheck{
		GRPC:     fmt.Sprintf("%s:%d", localIP.String(), port), // 这里一定是外部可以访问的地址
		Timeout:  "10s",                                        // 超时时间
		Interval: "10s",                                        // 运行检查的频率
		// 指定时间后自动注销不健康的服务节点
		// 最小超时时间为1分钟，收获不健康服务的进程每30秒运行一次，因此触发注销的时间可能略长于配置的超时时间。
		DeregisterCriticalServiceAfter: "1m",
	}
	srv := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", serviceName, ip, port), // 服务唯一ID
		Name:    serviceName,                                    // 服务名称
		Tags:    []string{"Meng", "hello", "books"},             // 为服务打标签
		Address: ip,
		Port:    port,
		Check:   check,
	}
	return c.client.Agent().ServiceRegister(srv)
}
