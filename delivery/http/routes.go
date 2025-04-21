package http

import (
	"github.com/gin-gonic/gin"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Summary Home page
// @Tags common
// @Success 200 {object} map[string]string
// @Router / [get]
func HomeHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello friend."})
}

// @Summary Server availability check
// @Tags common
// @Success 200 {object} map[string]string
// @Router /ping [get]
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func SetupRoutes(r *gin.Engine) {
	r.GET("/", HomeHandler)
	r.GET("/ping", PingHandler)

	// Swagger
	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
