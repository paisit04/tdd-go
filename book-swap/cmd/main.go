package main

import (
	"book-swap/db"
	"book-swap/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	log.Printf("working dir: %s", wd)

	migrationsPath := filepath.Join(wd, "/db/migrations")
	sourceURL := fmt.Sprintf("file://%s", migrationsPath)
	log.Printf("migrate: %s", sourceURL)

	port, ok := os.LookupEnv("BOOKSWAP_PORT")
	if !ok {
		log.Fatal("$BOOKSWAP_PORT not found")
	}

	postgresURL, ok := os.LookupEnv("BOOKSWAP_DB_URL")
	if !ok {
		log.Fatal("$BOOKSWAP_DB_URL not found")
	}

	m, err := migrate.New(sourceURL, postgresURL)
	if err != nil {
		log.Fatalf("migrate:%v", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration up:%v", err)
	}

	defer func() {
		m.Down()
	}()

	dbConn, err := gorm.Open(postgres.Open(postgresURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("db open:%v", err)
	}

	ps := db.NewPostingService()
	b := db.NewBookRepository(dbConn, ps)
	u := db.NewUserService(dbConn, b)
	h := handlers.NewHandler(b, u)

	router := handlers.ConfigureServer(h)
	log.Printf("Listening on :%s...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), router))
}
