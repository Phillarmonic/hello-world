package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

// EnvVar represents a single environment variable
type EnvVar struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	http.HandleFunc("/container-name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, hostname)
	})

	http.HandleFunc("/server-ip", func(w http.ResponseWriter, r *http.Request) {
		serverIP := os.Getenv("SERVER_IP")
		if serverIP == "" {
			serverIP = "unknown"
		}
		fmt.Fprint(w, serverIP)
	})

	// New endpoint for environment variables
	http.HandleFunc("/env-vars", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Get all environment variables
		envVars := []EnvVar{}
		for _, env := range os.Environ() {
			parts := strings.SplitN(env, "=", 2)
			key := parts[0]
			value := parts[1]

			// Mask sensitive values
			if strings.Contains(strings.ToLower(key), "key") ||
				strings.Contains(strings.ToLower(key), "token") ||
				strings.Contains(strings.ToLower(key), "secret") ||
				strings.Contains(strings.ToLower(key), "password") {
				value = "********"
			}

			envVars = append(envVars, EnvVar{Key: key, Value: value})
		}

		// Sort by key
		sort.Slice(envVars, func(i, j int) bool {
			return envVars[i].Key < envVars[j].Key
		})

		json.NewEncoder(w).Encode(envVars)
	})

	// Serve the static HTML file
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	log.Printf("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
