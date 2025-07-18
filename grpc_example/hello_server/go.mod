module hello_server

go 1.23.8

toolchain go1.23.11

require (
	github.com/golang/protobuf v1.5.4
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.1
	github.com/hashicorp/consul/api v1.32.1
	github.com/mennanov/fieldmask-utils v1.1.2
	google.golang.org/grpc v1.73.0
)

require (
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	golang.org/x/exp v0.0.0-20250305212735-054e65f0b394 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250603155806-513f23925822 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250603155806-513f23925822 // indirect
)

require (
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.38.0
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto v0.0.0-20250324211829-b45e905df463 // indirect
	//google.golang.org/genproto/googleapis/rpc v0.0.0-20250324211829-b45e905df463 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)

// 本地库
require github.com/testProject/pb v0.0.0

replace github.com/testProject/pb => ../pb
