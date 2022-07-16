run-usersrv:
	go run src/user-service/cmd/main/main.go
run-productsrv:
	go run src/product-service/cmd/main/main.go
run-auctionsrv:
	go run src/auction-service/cmd/main/main.go

gen-protoc-admin:
	protoc --go_out=. --go-grpc_out=. .\pkg\proto\admin.proto
gen-protoc-product:
	protoc --go_out=. --go-grpc_out=. .\pkg\proto\product.proto
gen-protoc-auction:
	protoc --go_out=. --go-grpc_out=. .\pkg\proto\auction.proto