run-usersrv:
	go run src/user-service/cmd/main/main.go

run-productsrv:
	go run src/product-service/cmd/main/main.go

gen-protoc-admin:
	protoc --go_out=. --go-grpc_out=. .\pkg\proto\admin.proto
