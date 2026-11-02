package main

import (
	"log"

	"github.com/alimfor/orderflow/services/user-service/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/health", handler.Health)

	log.Println("User Service started on port 8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatalln(err)
	}
}
