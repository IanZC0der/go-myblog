PKG := "github.com/IanZC0der/go-myblog"

dep:
	@go mod tidy
run:
	@go run main.go
gen:
	@protoc -I=. --go_out=. --go_opt=module=${PKG} --go-grpc_out=. --go-grpc_opt=module=${PKG} apps/commentrpc/proto/comment.proto
	@protoc-go-inject-tag -input=apps/commentrpc/*.pb.go
help:
	@echo "Available commands:"
	@echo "  dep  - tidy up the dependencies using 'go mod tidy'"
	@echo "  run  - run the main application using 'go run main.go'"
	@echo "  gen  - generate gRPC files from proto definitions"