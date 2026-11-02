package health

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(Response{Status: "ok"})
	if err != nil {
		log.Fatal(err)
	}
}
