package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"env"`
	Version     string `json:"version"`
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "port to listen on")
	flag.StringVar((*string)(&cfg.env), "env", "development", "environment")
	flag.Parse()
	fmt.Println("Running version:", version)

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		currentStatus := AppStatus{
			Status:      "OK",
			Environment: cfg.env,
			Version:     version,
		}
		js, err := json.MarshalIndent(currentStatus, "", "\t")
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
