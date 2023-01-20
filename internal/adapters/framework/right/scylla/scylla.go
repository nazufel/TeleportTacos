package scylla

import (
	"log"
	"os"

	"github.com/TeleportTacos/proto/pb"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

type Adapter struct {
	session *gocql.Session
}

// NewAdapter creates a new connection to the database
func NewAdapter() (*Adapter, error) {

	// TODO: pull in nodes from env
	cluster := gocql.NewCluster("scylla-node1", "scylla-node2", "scylla-node3")
	cluster.Consistency = gocql.Any
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

// SeedDatabase migrates and seeds the database with some demo data
func (da Adapter) SeedDatabase() error {

	// migrate the database schema

	// create tacos keyspace
	da.session.Query("CREATE KEYSPACE tacos;").Exec()

	// create table for the menu
	da.session.Query("CREATE TABLE tacos.menu (id UUID PRIMARY KEY, name text, description text, price float)").Exec()

	// seed the taco.menu table
	menuId := uuid.New()

	menuitem := pb.MenuItem{
		Name:        "The Taco",
		Description: "Classic taco",
		Price:       5.00,
	}

	da.session.Query("INSERT INTO tacos.menu (id, name, description, price) VALUES (%v, %v, %v, %v)", menuId, menuitem.Name, menuitem.Description, menuitem.Price).Exec()

	return nil
}

// // GetMenu returns a list of menu items from the DB
// func (da Adapter) GetMenu(pb.EmptyRequest) (pb.MenuItem, error) {

// 	var menuitem = pb.MenuItem{}

// 	da.session.Query()

// 	return menuitem, nil
// }
