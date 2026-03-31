package main

import (
	"log"
	"net/http"

	"github.com/DenariusData/API-5SEM-BACKEND/internal/config"
	"github.com/DenariusData/API-5SEM-BACKEND/internal/handler"
)

func main() {
	cfg := config.Load()

	mux := http.NewServeMux()
	handler.Register(mux)

	addr := ":" + cfg.Port
	log.Printf("Server starting on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
