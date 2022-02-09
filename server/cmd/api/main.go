package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"env"`
	Version     string `json:"version"`
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "port to listen on")
	flag.StringVar((*string)(&cfg.env), "env", "development", "environment")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgresql://postgres:XV9xY5hsHiYdR6QuGp4J@containers-us-west-9.railway.app:7241/railway", "database connection string")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	logger.Printf("Starting server on port %d", cfg.port)

	err := srv.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}

}
