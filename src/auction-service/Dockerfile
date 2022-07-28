FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.* /
RUN go mod download

COPY pkg/ /pkg/
COPY src/auction-service/ /src/auction-service/

RUN CGO_ENABLED=0 GOOS=linux go build -o /main /src/auction-service/cmd/main/main.go

FROM scratch

WORKDIR /

COPY --from=build /pkg /pkg/
COPY --from=build /main /main

EXPOSE 3031
EXPOSE 50053

ENTRYPOINT ["/main"]