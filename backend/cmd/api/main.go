package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/dmslmvsk/daily-tracker/backend/internal/api"
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



    log.Println("Database connection, Ping and Migrations — Success!")

    store := repository.New(db)
    router := api.NewRouter(store)
    log.Println("Server is starting on port 8080...")

    err = http.ListenAndServe(":8080",router)
    if err != nil {
        log.Fatal("Server failed to start: ",err)
    }
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