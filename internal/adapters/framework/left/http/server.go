package http

import (
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Adapter struct{}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (a Adapter) Run() {

	// // Create non-global registry.
	// reg := prometheus.NewRegistry()

	// // Add go runtime metrics and process collectors.
	// reg.MustRegister(
	// 	collectors.NewGoCollector(),
	// 	collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	// )

	http.Handle("/metrics", promhttp.Handler())

	httpPort := ":" + os.Getenv("HTTP_SERVER_LISTEN_PORT")

	log.Printf("serving http on port %v", httpPort)
	if err := http.ListenAndServe(httpPort, nil); err != nil {
		log.Fatalf("unable to serve http at %v : %v", httpPort, err)
	}
}
