package scylla

import (
	"log"
	"os"

	"github.com/gocql/gocql"
)

type Adapter struct {
	session *gocql.Session
}

// NewAdapter creates a new connection to the database
func NewAdapter() (*Adapter, error) {

	// TODO: pull in nodes from env
	cluster := gocql.NewCluster("scylla-node1.scylla-summit-2023_demo", "scylla-node2.scylla-summit-2023_demo", "scylla-node3.scylla-summit-2023_demo")
	cluster.Keyspace = "tacos"
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("unable to establish DB connection: %v", err)
		return nil, err
	}

	return &Adapter{session: session}, nil
}

// CloseDBConnection closes the DB connection
func (da Adapter) CloseDBConnection() {

	log.Println("closing the db")
	da.session.Close()

	// shut down the service since it can't do anything without a DB
	os.Exit(0)
}

// GetMenu returns a list of menu items from the DB
func (da Adapter) GetMenu()
