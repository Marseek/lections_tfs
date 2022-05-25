package main

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://user:passwd@localhost:5442/fintech?sslmode=disable")
	if err != nil {
		log.Fatalf("can't connect to db: %s", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("can't ping db: %s", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("can't connect to db: %s", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://6_migrations/db/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("can't creaate migration object: %s", err)
	}
	err = m.Up()
	if err != nil {
		log.Fatalf("migrations didin't pass: %s", err)
	}
	log.Printf("%s", "all migrations pass")
}
