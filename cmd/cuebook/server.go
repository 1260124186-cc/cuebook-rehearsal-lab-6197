package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const defaultHTTPPort = "8080"

type healthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

func serve(args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("serve does not accept arguments")
	}

	port := strings.TrimSpace(os.Getenv("PORT"))
	if port == "" {
		port = defaultHTTPPort
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", health)
	mux.HandleFunc("/", health)

	return http.ListenAndServe(":"+port, mux)
}

func health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(healthResponse{
		Status:  "ok",
		Service: "cuebook-rehearsal-lab",
	})
}
