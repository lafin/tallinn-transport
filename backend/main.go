package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	cache "github.com/go-pkgz/expirable-cache"
	"github.com/lafin/tallinn-transport/provider"
)

func main() {
	c, _ := cache.NewCache(cache.TTL(time.Minute * 5))
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		var err error
		var tallinnTransport []provider.Transport
		tallinnTransport, err = provider.GetTallinnTransport()
		if err != nil {
			log.Printf("[ERROR] get tallinn transport, %s", err)
			cached, ok := c.Get("tallinnTransport")
			if ok {
				tallinnTransport = cached.([]provider.Transport)
			}
		} else {
			c.Set("tallinnTransport", tallinnTransport, time.Minute*5)
		}

		var elronTransport []provider.Transport
		elronTransport, err = provider.GetElronTransport()
		if err != nil {
			log.Printf("[ERROR] get elron transport, %s", err)
			cached, ok := c.Get("tallinnTransport")
			if ok {
				elronTransport = cached.([]provider.Transport)
			}
		} else {
			c.Set("elronTransport", elronTransport, time.Minute*5)
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
