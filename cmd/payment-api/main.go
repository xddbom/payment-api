package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	delivery "payment-api/delivery/http"
	_ "payment-api/docs"
)

// @title Payment API
// @version 1.0

// @host localhost:8080
// @BasePath /

func main() {
	r := gin.Default()

	delivery.SetupRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Println("Server failed.")
	}
	fmt.Println("Server is successfully started!")
}
