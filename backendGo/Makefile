all:
	protoc \
		--go_out=protos/ --go-grpc_out=protos/ \
		protos/*.proto
	protoc \
		--grpc-gateway_out=protos/ \
		protos/membership.proto

clean:
	rm *.pb.go *.pb.gw.go *.swagger.json