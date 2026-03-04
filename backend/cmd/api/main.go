package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/dmslmvsk/daily-tracker/backend/internal/repository"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
    dbURL := os.Getenv("DB_SOURCE")
    if dbURL == "" {
        log.Fatal("DB_SOURCE is not set")
    }

    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        log.Fatal("cannot connect to db:", err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal("cannot ping db:", err)
    }

    runDBMigration("file://migrations", dbURL)


    store := repository.New(db)
    _ = store

    log.Println("Database connection, Ping and Migrations — Success!")

    select {}
}

func runDBMigration(migrationURL string, dbURL string) {
	m, err := migrate.New(migrationURL, dbURL)
	if err != nil {
		log.Fatal("cannot create migration instance:", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up:", err)
	}

	log.Println("db migrated successfully")
}