package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"sync/atomic"
)

type Response struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func service() http.Handler {
	r := chi.NewRouter()

	// Ping
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		responseBody, e := json.Marshal(Response{
			StatusCode: 200,
			Message:    "Pong",
		})

		if e != nil {
			log.Fatalf("[ERROR] : marshal %s: ", e.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseBody)
	})

	// Headthz Check
	r.Get("/api/healthz", func(w http.ResponseWriter, r *http.Request) {
		if atomic.LoadInt32(&healthy) == 1 {
			responseBody, e := json.Marshal(Response{
				StatusCode: 200,
				Message:    "Server is Running",
			})

			if e != nil {
				log.Fatalf("[ERROR] : marshal %s: ", e.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(responseBody)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
	})
	return r
}
