package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const version = "1.0.0" // keep track version of this app

// AppStatus to send back JSON formatted data about the current condition of the server
type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type config struct {
	port int
	env  string
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production)")
	flag.Parse()

	fmt.Println("app is currently running")

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprint(w, "hollaaa")
		currentStatus := AppStatus{
			Status:      "Available",
			Environment: cfg.env,
			Version:     version,
		}

		js, err := json.MarshalIndent(currentStatus, "", "\t")
		if err != nil {
			log.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)

	})

	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", *&cfg.port), nil)
	if err != nil {
		fmt.Println(err)
	}
}
