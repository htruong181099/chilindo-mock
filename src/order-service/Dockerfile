FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.* /
RUN go mod download

COPY pkg/ /pkg/
COPY src/order-service/ /src/order-service/

RUN CGO_ENABLED=0 GOOS=linux go build -o /main /src/order-service/cmd/main/main.go

FROM scratch

WORKDIR /

COPY --from=build /pkg /pkg/
COPY --from=build /main /main

EXPOSE 3032

ENTRYPOINT ["/main"]