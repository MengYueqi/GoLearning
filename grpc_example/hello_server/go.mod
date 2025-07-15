module hello_server

go 1.23.0

require (
	github.com/golang/protobuf v1.5.4
	github.com/mennanov/fieldmask-utils v1.1.2
	google.golang.org/grpc v1.73.0
)

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.1 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250603155806-513f23925822 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250603155806-513f23925822 // indirect
)

require (
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto v0.0.0-20250324211829-b45e905df463 // indirect
	//google.golang.org/genproto/googleapis/rpc v0.0.0-20250324211829-b45e905df463 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)

// 本地库
require github.com/testProject/pb v0.0.0

replace github.com/testProject/pb => ../pb
