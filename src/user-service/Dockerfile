FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.* /
RUN go mod download

COPY pkg/ /pkg/
COPY src/user-service/ /src/user-service/

RUN CGO_ENABLED=0 GOOS=linux go build -o /main /src/user-service/cmd/main/main.go

FROM scratch

WORKDIR /

COPY --from=build /main /main

EXPOSE 3000
EXPOSE 50051

ENTRYPOINT ["/main"]