package main

import (
	"log"

	rpc "github.com/teleporttacos/internal/adapters/framework/left/grpc"
	"github.com/teleporttacos/internal/adapters/framework/left/http"
	"github.com/teleporttacos/internal/adapters/framework/right/scylla"
	"github.com/teleporttacos/internal/application/core/geo"
	"github.com/teleporttacos/internal/application/core/geo/api"
)

func main() {

	// wire up db
	dbAdapter, err := scylla.NewAdapter()
	if err != nil {
		log.Fatalf("failed to connect to DB - %v", err)
	}

	defer dbAdapter.CloseDBConnection()

	dbAdapter.SeedDatabase()

	// wire up the core layer
	core := geo.New()

	// wire up the api layer
	appAPI := api.NewApplication(dbAdapter, core)

	// wire up and run http server
	httpAdapter := http.NewAdapter()

	// run an http sever in another thread to
	// access the Prometheus exporter metrics
	go httpAdapter.Run()

	// wire up and run grpc server
	gRPCAdapter := rpc.NewAdapter(appAPI)
	gRPCAdapter.Run()
}
