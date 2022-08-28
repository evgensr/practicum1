package genproto

// //go:generate go install -v google.golang.org/protobuf/cmd/protoc-gen-go
// //go:generate go install -v google.golang.org/grpc/cmd/protoc-gen-go-grpc
// //go:generate go install -v github.com/gogo/protobuf/protoc-gen-gofast

// //go:generate protoc  --go_out=../../   --go-grpc_out=../../  -I ../../api   ../../api/proto/short.proto
// //go:generate protoc  --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import  ../../api/proto/short.proto
// //go:generate  protoc  --go_out=../../pkg/api --go_opt=paths=source_relative --go-grpc_out=../../pkg/api --go-grpc_opt=paths=source_relative -I../../api/proto/short.proto

// //go:generate  protoc   --go_out=../../internal/pb --go_opt=paths=source_relative --go-grpc_out=../../internal/pb --go-grpc_opt=paths=source_relative --proto_path=internal/proto/short.proto
