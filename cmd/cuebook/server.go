package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"example.com/cuebook-rehearsal-lab/internal/app"
	"example.com/cuebook-rehearsal-lab/internal/format"
	"example.com/cuebook-rehearsal-lab/internal/model"
)

const defaultHTTPPort = "8080"

type healthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

type composeRequest struct {
	Show     string `json:"show"`
	Director string `json:"director"`
}

type reviewRequest struct {
	Show       string `json:"show"`
	Department string `json:"department"`
	Author     string `json:"author"`
}

type completeRequest struct {
	Show     string `json:"show"`
	Director string `json:"director"`
}

func serve(service *app.Service, args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("serve does not accept arguments")
	}

	port := strings.TrimSpace(os.Getenv("PORT"))
	if port == "" {
		port = defaultHTTPPort
	}
	return http.ListenAndServe(":"+port, newHTTPHandler(service))
}

func newHTTPHandler(service *app.Service) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", health)
	mux.HandleFunc("/runs", composeHandler(service))
	mux.HandleFunc("/reviews", reviewHandler(service))
	mux.HandleFunc("/complete", completeHandler(service))
	mux.HandleFunc("/", health)
	return mux
}

func health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w, http.MethodGet)
		return
	}
	writeJSON(w, http.StatusOK, healthResponse{
		Status:  "healthy",
		Service: "cuebook-rehearsal-lab",
	})
}

func composeHandler(service *app.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			methodNotAllowed(w, http.MethodPost)
			return
		}
		var request composeRequest
		if err := decodeJSON(r, &request); err != nil {
			badRequest(w, err)
			return
		}
		if strings.TrimSpace(request.Director) == "" {
			request.Director = "Mira"
		}
		response, err := service.Assemble(request.Show, request.Director)
		if err != nil {
			badRequest(w, err)
			return
		}
		payload, err := format.AssembleJSON(response)
		writeFormattedJSON(w, http.StatusCreated, payload, err)
	}
}

func reviewHandler(service *app.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			methodNotAllowed(w, http.MethodPost)
			return
		}
		var request reviewRequest
		if err := decodeJSON(r, &request); err != nil {
			badRequest(w, err)
			return
		}
		if strings.TrimSpace(request.Author) == "" {
			request.Author = "department-lead"
		}
		response, err := service.Review(request.Show, model.Department(strings.TrimSpace(request.Department)), request.Author)
		if err != nil {
			badRequest(w, err)
			return
		}
		payload, err := format.ReviewJSON(response)
		writeFormattedJSON(w, http.StatusOK, payload, err)
	}
}

func completeHandler(service *app.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			methodNotAllowed(w, http.MethodPost)
			return
		}
		var request completeRequest
		if err := decodeJSON(r, &request); err != nil {
			badRequest(w, err)
			return
		}
		if strings.TrimSpace(request.Director) == "" {
			request.Director = "Mira"
		}
		result, err := service.Complete(request.Show, request.Director)
		if err != nil {
			badRequest(w, err)
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{
			"message": "cue book published: " + app.ScenarioLabel(result),
		})
	}
}

func decodeJSON(r *http.Request, target any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("invalid JSON request: %w", err)
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		if err == nil {
			return fmt.Errorf("request body must contain one JSON object")
		}
		return fmt.Errorf("invalid JSON request: %w", err)
	}
	return nil
}

func methodNotAllowed(w http.ResponseWriter, allowed string) {
	w.Header().Set("Allow", allowed)
	writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
}

func badRequest(w http.ResponseWriter, err error) {
	writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
}

func writeFormattedJSON(w http.ResponseWriter, status int, payload []byte, err error) {
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to encode response"})
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_, _ = w.Write(append(payload, '\n'))
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
