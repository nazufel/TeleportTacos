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

	log.Printf("successfully connected to scylla nodes: %v", cluster.Hosts)

	return &Adapter{session: session}, nil
}

// CloseDBConnection closes the DB connection
func (da Adapter) CloseDBConnection() {

	log.Println("closing the db connection")
	da.session.Close()

	// shut down the service since it can't do anything without a DB
	os.Exit(0)
}

// GetMenuItem retreives a single item from the database
// in this case, the only item.
func (da Adapter) GetMenuItem(m *pb.MenuItemRequest) (*pb.MenuItemResponse, error) {

	var returnItem pb.MenuItemResponse

	// run query
	err := da.session.Query("SELECT * FROM tacos.menu;").Scan(&returnItem.Id, &returnItem.Description, &returnItem.Name, &returnItem.Price)

	// if query failed, wait and restart the loop
	if err != nil {
		log.Printf("failed to query database: %v", err)
	}

	log.Printf("retrieved menu item name: %v", returnItem.Name)
	return &returnItem, nil
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

	// setting a tombstone on this table becuse the demo will be writing a
	// lot of data really fast and I don't want to be filling up the contianer's volume
	if err := da.session.Query("CREATE TABLE IF NOT EXISTS tacos.orders (id UUID PRIMARY KEY, count int, created_at timestamp, menu_item text, payment float, teleport_alt float, teleport_lat float, teleport_long float, updated_at timestamp) WITH default_time_to_live = 600").Exec(); err != nil {
		log.Fatalf("unable to create tacos.orders table: %v", err)
	}

	menuId := uuid.New()

	menuItem := pb.MenuItemRequest{
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

// PlaceOrder func places an order to teleport tacos
func (da Adapter) PlaceOrder(o *pb.OrderRequest) (*pb.OrderResponse, error) {

	orderId := uuid.New()

	log.Printf("submitting order: %v to the database", orderId)

	var res pb.OrderResponse
	// TODO: get timestamps working, skipping for now
	err := da.session.Query("INSERT INTO tacos.orders(id, count, menu_item, payment, teleport_alt, teleport_lat, teleport_long) VALUES (?,?,?,?,?,?,?);", orderId.String(), o.Count, o.MenuItem, o.Payment, o.TeleportAlt, o.TeleportLat, o.TeleportLong).Exec()
	if err != nil {
		// log.Printf("unable to place order %v: %v", orderId.String(), err)
		return &res, err
	}

	log.Printf("successfully placed order: %v", orderId)

	return &res, nil
}
