# syntax=docker/dockerfile:1

# build stage
FROM golang:1.19-alpine as builder
WORKDIR /go/src

# copy program files
COPY ./cmd/client/main.go ./cmd/client/
COPY proto/*.proto ./proto/
COPY go.mod go.sum ./

# install dependencies and build the app
RUN apk add protobuf
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go mod download

# build protobuf files
RUN mkdir ./proto/pb
# build protobuf
RUN protoc -I=./proto --go_out=./proto/ --go-grpc_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:./proto/pb ./proto/tacos.proto

# build the program
ENV CGO_ENABLED=0 
ENV GOOS=linux
RUN go build -a -installsuffix cgo -o client ./cmd/client/main.go

# final phase: copy the built binary into a scratch image to run
FROM scratch

## set the working directory
WORKDIR /go/bin/sever

## copy the combiled binary to the smaller image
COPY --from=builder /go/src/client .

## start the binary when the container starts
ENTRYPOINT ["./client"]