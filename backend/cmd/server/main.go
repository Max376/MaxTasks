package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	serviceName         = "maxtasks-api"
	databasePingTimeout = 2 * time.Second
)

// databasePinger keeps the readiness handler independent from a concrete SQL
// implementation, which makes its failure behavior straightforward to test.
type databasePinger interface {
	PingContext(context.Context) error
}

type serviceStatus struct {
	Service string           `json:"service"`
	Status  string           `json:"status"`
	Checks  *readinessChecks `json:"checks,omitempty"`
}

type readinessChecks struct {
	Database string `json:"database"`
}

func main() {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	db, err := openDatabase(os.Getenv("DATABASE_URL"))
	if err != nil {
		// Keep connection details, including credentials in a DSN, out of logs.
		log.Fatal("database configuration is invalid")
	}
	if db != nil {
		defer db.Close()
	}

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           newMux(db),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("MaxTasks API listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func openDatabase(databaseURL string) (*sql.DB, error) {
	if strings.TrimSpace(databaseURL) == "" {
		return nil, nil
	}

	return sql.Open("pgx", databaseURL)
}

func newMux(db databasePinger) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", healthHandler)
	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("GET /ready", readinessHandler(db))
	return mux
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, serviceStatus{
		Service: serviceName,
		Status:  "ok",
	})
}

func readinessHandler(db databasePinger) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		if db == nil {
			writeJSON(w, http.StatusServiceUnavailable, serviceStatus{
				Service: serviceName,
				Status:  "not_ready",
				Checks:  &readinessChecks{Database: "unavailable"},
			})
			return
		}

		requestContext, cancel := context.WithTimeout(request.Context(), databasePingTimeout)
		defer cancel()
		if err := db.PingContext(requestContext); err != nil {
			writeJSON(w, http.StatusServiceUnavailable, serviceStatus{
				Service: serviceName,
				Status:  "not_ready",
				Checks:  &readinessChecks{Database: "unavailable"},
			})
			return
		}

		writeJSON(w, http.StatusOK, serviceStatus{
			Service: serviceName,
			Status:  "ready",
			Checks:  &readinessChecks{Database: "ok"},
		})
	}
}

func writeJSON(w http.ResponseWriter, statusCode int, response serviceStatus) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(response)
}
