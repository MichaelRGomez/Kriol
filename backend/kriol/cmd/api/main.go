//Filename: kriol/backend/kriol/cmd/api/main.go

package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// version number
const version = "1.0.0"

// configuration settings
type config struct {
	port int
	env  string   //development, staging, production
	db   struct { //database field
		dsn string
	}
}

// dependency injection
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	//reading the flags
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development | staging | production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("APPLETREE_DB_DSN"), "PostgreSQL DSN")
	flag.Parse()

	//creating logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//create the connection pool
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	//instance of app struct
	app := &application{
		config: cfg,
		logger: logger,
	}

	//HTTP server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	//starting our server
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

// OpenDB() function returns a *sql.DB connection pool
func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	//Create a context with a 5-second timeout deadline
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
