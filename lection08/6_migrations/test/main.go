package main

import (
	"log"
	"time"

	"github.com/gocql/gocql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/cassandra"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cluster := gocql.NewCluster("localhost")
	keyspace := "estore"
	cluster.Keyspace = ""
	cluster.Timeout = 60 * time.Second

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal("while creating session: ", err)
	}
	log.Println("session established")

	q := `CREATE KEYSPACE IF NOT EXISTS ` + keyspace + ` WITH` +
		`  REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': '1'}` +
		`  AND DURABLE_WRITES = TRUE`
	//q = `DROP KEYSPACE ` + keyspace
	err = session.Query(q).Exec()
	if err != nil {
		log.Println("while creating QUERY: ", err)
	}
	session.Close()

	cluster = gocql.NewCluster("localhost")
	cluster.Keyspace = keyspace
	cluster.Timeout = 60 * time.Second
	session, err = cluster.CreateSession()
	defer session.Close()
	if err != nil {
		log.Fatal("while creating session: ", err)
	}

	log.Println("keyspase created")

	// Create driver instance from db.
	// Check each driver if it supports the WithInstance function.
	// `import "github.com/golang-migrate/migrate/v4/database/postgres"`
	driver, err := cassandra.WithInstance(session, &cassandra.Config{KeyspaceName: "estore"})
	if err != nil {
		log.Fatal(err)
	}

	// Read migrations from /home/mattes/migrations and connect to a local postgres database.
	m, err := migrate.NewWithDatabaseInstance("file://6_migrations/db/migrations", "cassandra", driver)
	if err != nil {
		log.Fatal("Read migrations from", err)
	}
	log.Println("before UP\n")
	// Migrate all the way up ...
	if err := m.Up(); err != nil {
		log.Fatal("while applying migrations: ", err)
	}
	log.Println("migrations done\n")
}
func main2() {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "estore"

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal("while creating session: ", err)
	}
	log.Println("session established")
	// Create driver instance from db.
	// Check each driver if it supports the WithInstance function.
	// `import "github.com/golang-migrate/migrate/v4/database/postgres"`
	driver, err := cassandra.WithInstance(session, &cassandra.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Read migrations from /home/mattes/migrations and connect to a local postgres database.
	m, err := migrate.NewWithDatabaseInstance("file://6_migrations/db/migrations", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	// Migrate all the way up ...
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
