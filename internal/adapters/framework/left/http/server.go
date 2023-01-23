package http

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Adapter struct{}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (a Adapter) Run() {

	http.Handle("/metrics", promhttp.Handler())

	log.Println("serving http on port 9080")
	if err := http.ListenAndServe(":9080", nil); err != nil {
		log.Fatalf("unable to serve http at 9080: %v", err)
	}
}
