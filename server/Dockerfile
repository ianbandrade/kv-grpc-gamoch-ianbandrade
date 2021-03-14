FROM golang:1.16.2-alpine3.13 AS builder

WORKDIR /build
RUN go get google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

RUN apk add --no-cache -qt .build-deps protoc
COPY protofiles protofiles
RUN mkdir pb && \
    protoc --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    --proto_path=protofiles protofiles/*.proto && \
    rm -rf protofiles
RUN apk del -q .build-deps

COPY go.mod go.sum ./
RUN go mod vendor

ENV CGO_ENABLED 0
ENV GOOS linux

COPY main.go main.go
RUN go build -o grpc-server .


FROM alpine:3.13
COPY --from=builder /build/grpc-server /usr/local/bin

ENTRYPOINT ["grpc-server"]
