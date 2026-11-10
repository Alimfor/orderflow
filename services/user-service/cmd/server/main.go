package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alimfor/orderflow/services/user-service/internal/server"
)

func main() {
	router := server.NewRouter()

	srv := &http.Server{
		Addr:              ":8081",
		Handler:           router,
		ReadTimeout:       5 * time.Second,
		IdleTimeout:       120 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)

	defer stop()

	serverErrors := make(chan error, 1)

	go func() {
		log.Println("Server started on :8081")
		serverErrors <- srv.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		log.Println("Shutdown signal received")
	case err := <-serverErrors:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Server failed: %v", err)
		}
		return
	}

	stop()

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	defer cancel()

	shutdownErr := srv.Shutdown(shutdownCtx)

	if shutdownErr != nil {
		log.Printf("Server shutdown failed: %v", shutdownErr)

		if closeErr := srv.Close(); closeErr != nil {
			log.Printf("Force close failed: %v", closeErr)
		}

		return
	}

	serverErr := <-serverErrors

	if serverErr != nil &&
		!errors.Is(serverErr, http.ErrServerClosed) {
		log.Printf("HTTP server error: %v", serverErr)
		return
	}

	if shutdownErr != nil {
		log.Println("Server stopped forcefully")
		return
	}

	log.Println("Server stopped successfully")
}
