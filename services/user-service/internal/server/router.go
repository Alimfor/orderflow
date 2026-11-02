package server

import (
	"net/http"

	"github.com/alimfor/orderflow/services/user-service/internal/health"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("GET /health", health.Handler)
	return router
}
