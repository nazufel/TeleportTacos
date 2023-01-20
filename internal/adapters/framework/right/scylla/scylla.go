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

	log.Println("migrating and seeding the db. please wait...")

	// drop keyspace if exists to start from scratch
	// if err := da.session.Query("DROP KEYSPACE IF EXISTS tacos;"); err != nil {
	// 	log.Fatalf("unable to drop the keyspace: %v", err)

	// }

	// migrate the database schema

	// create tacos keyspace
	if err := da.session.Query("CREATE KEYSPACE IF NOT EXISTS tacos WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 3};").Exec(); err != nil {
		log.Fatalf("unable to create keyspace: %v", err)
	}

	// create table for the menu
	if err := da.session.Query("CREATE TABLE IF NOT EXISTS tacos.menu (id UUID PRIMARY KEY, name text, description text, price float)").Exec(); err != nil {
		log.Fatalf("unable to create tacos.menu table: %v", err)
	}

	// seed the taco.menu table
	menuId := uuid.New()

	menuItem := pb.MenuItem{
		Name:        "The Taco",
		Description: "Classic taco",
		Price:       5.00,
	}

	err := da.session.Query("INSERT INTO tacos.menu(id, name, description, price) VALUES (?,?,?,?);", menuId, menuItem.Name, menuItem.Description, menuItem.Price).Exec()
	if err != nil {
		log.Fatalf("unable to seed tacos.menu table: %v", err)
	}

	// log.Printf("query: %v", query)
	// if err := da.session.Query("INSERT INTO tacos.menu (\"id\", \"name\", \"description\", \"price\") VALUES (21245886-5805-4212-b452-a0b8090acf34, 'taco', 'taco', 5.0);").Exec(); err != nil {
	// 	log.Fatalf("unable to seed tacos.menu table: %v", err)
	// }

	log.Println("finished migrating and seeding the db.")

	return nil
}

// // GetMenu returns a list of menu items from the DB
// func (da Adapter) GetMenu(pb.EmptyRequest) (pb.MenuItem, error) {

// 	var menuitem = pb.MenuItem{}

// 	da.session.Query()

// 	return menuitem, nil
// }
