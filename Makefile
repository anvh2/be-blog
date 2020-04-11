BIN = blog

clean:
	rm -rf $(BIN)

build: 
	GOOS=linux GOARCH=amd64 go build -o $(BIN)

genpb:
	protoc -I/usr/local/include -Iidl \
		-I$$GOPATH/src \
		-I$$GOPATH/src/github.com/gogo/protobuf/protobuf \
		-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
		--gogo_out=plugins=grpc:grpc-gen/blog \
		--swagger_out=logtostderr=true:grpc-gen/blog \
		--grpc-gateway_out=logtostderr=true:grpc-gen/blog \
		idl/blog.proto

	protoc -I/usr/local/include -Iidl \
		-I$$GOPATH/src \
		-I$$GOPATH/src/github.com/gogo/protobuf/protobuf \
		-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
		--gogo_out=plugins=grpc:grpc-gen/user \
		--swagger_out=logtostderr=true:grpc-gen/user \
		--grpc-gateway_out=logtostderr=true:grpc-gen/user \
		idl/user.proto
runLocal: 
	go run main.go blogs --config config.dev.toml

deploy: clean build