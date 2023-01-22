package scylla

import (
	"log"
	"os"
	"time"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/teleporttacos/proto/pb"
)

type Adapter struct {
	session *gocql.Session
}

// NewAdapter creates a new connection to the database
func NewAdapter() (*Adapter, error) {

	// TODO: pull in nodes from env
	cluster := gocql.NewCluster("scylla-node1", "scylla-node2", "scylla-node3")
	cluster.Timeout = 5 * time.Second
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

func (da Adapter) GetMenuItem() (pb.MenuItem, error) {

	var menuItem pb.MenuItem

	// seed the tacos table
	err := da.session.Query("SELECT * FROM tacos.menu;").Scan(&menuItem.Id, &menuItem.Description, &menuItem.Name, &menuItem.Price)
	if err != nil {
		log.Printf("unable get menu item from the database: %v", err)
		return menuItem, err
	}

	log.Printf("found menuItem: %v", menuItem.Name)

	return menuItem, nil
}

// SeedDatabase migrates and seeds the database with some demo data
func (da Adapter) SeedDatabase() error {

	log.Println("migrating and seeding the db. please wait...")

	// drop keyspace if exists to start this demo from scratch
	if err := da.session.Query("DROP KEYSPACE IF EXISTS tacos;").Exec(); err != nil {
		log.Fatalf("unable to drop the keyspace: %v", err)
	}

	// migrate the database schema

	// create tacos keyspace
	if err := da.session.Query("CREATE KEYSPACE IF NOT EXISTS tacos WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 3};").Exec(); err != nil {
		log.Fatalf("unable to create keyspace: %v", err)
	}

	// create tables for the menu and orders
	if err := da.session.Query("CREATE TABLE IF NOT EXISTS tacos.menu (id UUID PRIMARY KEY, name text, description text, price float)").Exec(); err != nil {
		log.Fatalf("unable to create tacos.menu table: %v", err)
	}

	if err := da.session.Query("CREATE TABLE IF NOT EXISTS tacos.orders (id UUID PRIMARY KEY, count int, created_at timestamp, menu_item text, price float, teleport_alt float, teleport_lat float, teleport_long float, updated_at timestamp)").Exec(); err != nil {
		log.Fatalf("unable to create tacos.orders table: %v", err)
	}

	menuId := uuid.New()

	menuItem := pb.MenuItem{
		Name:        "The Taco",
		Description: "Classic Taco",
		Price:       5.00,
	}

	// seed the tacos table
	err := da.session.Query("INSERT INTO tacos.menu(id, name, description, price) VALUES (?,?,?,?);", menuId.String(), menuItem.Name, menuItem.Description, menuItem.Price).Exec()
	if err != nil {
		log.Fatalf("unable to seed tacos.menu table: %v", err)
	}

	log.Println("finished migrating and seeding the db.")

	return nil
}

// // GetMenu returns a list of menu items from the DB
// func (da Adapter) GetMenu(pb.EmptyRequest) (pb.MenuItem, error) {

// 	var menuitem = pb.MenuItem{}

// 	da.session.Query()

// 	return menuitem, nil
// }
