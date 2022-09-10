//filename: kriol/backend/kriol/cmd/api/main.go

package main

import (
	"flag"     //for logging
	"fmt"      //for printing to webpage
	"log"      //also for logging
	"net/http" //for webserver
	"os"       //also also for logging
	"time"     //timing the log reports
)

// The application version number
const version = "1.0.0"

// The configuration settings
type config struct {
	port int
	env  string //development, staging, production, etc
}

// Dependency Injection
type application struct {
	config config
	logger *log.Logger
}

// main
func main() {
	//configing config and environment status
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development | staging | production)")
	flag.Parse()

	//creating logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//application struct
	app := &application{
		config: cfg,
		logger: logger,
	}

	//creaing servr mux
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)
	mux.HandleFunc("/v1/createEntryHandler", app.createEntryHandler)
	mux.HandleFunc("/v1/showEntryHandler", app.showEntryHandler)

	//creating HTTP server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		//Handler: app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	//starting the server
	logger.Printf("sarting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}
