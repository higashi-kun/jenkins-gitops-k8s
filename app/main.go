package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// HealthResponse は /healthz エンドポイントのレスポンス形式。
type HealthResponse struct {
	Status string `json:"status"`
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(HealthResponse{Status: "ok"}); err != nil {
		log.Printf("failed to encode healthz response: %v", err)
	}
}

func main() {
	http.HandleFunc("/healthz", healthzHandler)
	log.Println("server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
