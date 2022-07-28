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

gen-docker-user:
	docker build -t chilindo-usersrv -f .\src\user-service\Dockerfile .

run-docker-user:
	docker run chilindo-usersrv --env DB_CONNECTION_STRING="root:@Hoang_123456@tcp(localhost:3306)/chilindo?parseTime=true"

