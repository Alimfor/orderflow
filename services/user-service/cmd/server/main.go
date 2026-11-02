package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/alimfor/orderflow/services/user-service/internal/server"
)

func main() {
	router := server.NewRouter()

	srv := &http.Server{
		Addr:        ":8081",
		Handler:     router,
		ReadTimeout: 5 * time.Second,
		IdleTimeout: 120 * time.Second,
	}

	log.Println("User service started on :8081")

	if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
