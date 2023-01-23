package rpc

import (
	"log"
	"net"
	"os"

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

	listenPort := ":" + os.Getenv("GRPC_SERVER_LISTEN_PORT")

	listen, err := net.Listen("tcp", listenPort)
	if err != nil {
		log.Fatalf("failed to listen on port %v - %v", listenPort, err)
	}

	tacoServiceServer := a
	grpcServer := grpc.NewServer()
	pb.RegisterTacoServiceServer(grpcServer, tacoServiceServer)

	log.Printf("serving grpc server on port: %v", listenPort)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc server on port %v - %v", listenPort, err)
	}
}
