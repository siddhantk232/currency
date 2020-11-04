.PHONY: protos

protos: 
	protoc -I protos/ \
	--go-grpc_opt=paths=source_relative --go-grpc_out=protos/currency \
	--go_opt=paths=source_relative --go_out=protos/currency protos/currency.proto

