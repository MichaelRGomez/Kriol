//Filename: kriol/backend/kriol/cmd/api/main.go

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// version number
const version = "1.0.0"

// configuration settings
type config struct {
	port int
	env  string //development, staging, production
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
	flag.Parse()

	//creating logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//instance of app struct
	app := &application{
		config: cfg,
		logger: logger,
	}

	//server mux
	mux := http.NewServeMux()
	mux.HandleFunc("v1/healthcheck", app.healthcheckHandler)

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
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
