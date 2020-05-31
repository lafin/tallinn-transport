package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/lafin/tallinn-transport/provider"
)

func main() {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		tallinnTransport, err := provider.GetTallinnTransport()
		if err != nil {
			log.Printf("[ERROR] get tallinn transport, %s", err)
		}
		elronTransport, err := provider.GetElronTransport()
		if err != nil {
			log.Printf("[ERROR] get elron transport, %s", err)
		}
		if err != nil {
			log.Printf("[ERROR] get transport, %s", err)
		}

		body, err := json.Marshal(append(tallinnTransport, elronTransport...))
		if err != nil {
			log.Printf("[ERROR] marshal json, %s", err)
		}
		_, err = w.Write(body)
		if err != nil {
			log.Printf("[ERROR] write body, %s", err)
		}
	})
	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", 3000),
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       30 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("[ERROR] start http server, %s", err)
	}
}
