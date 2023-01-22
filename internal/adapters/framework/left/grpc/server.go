package rpc

import (
	"log"
	"net"

	"github.com/teleporttacos/internal/ports"
	"github.com/teleporttacos/proto/pb"
	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

// Run registers the grpc server and bind the port
func (a Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("failed to listen on port 9999", err)
	}

	tacoServiceServer := a
	grpcServer := grpc.NewServer()
	pb.RegisterTacoServiceServer(grpcServer, tacoServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc server on port 9999: %v", err)
	}
}
