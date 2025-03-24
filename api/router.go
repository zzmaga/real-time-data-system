package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DataPayload struct {
	Source string  `json:"source"`
	Value  float64 `json:"value"`
}

var collectedData []DataPayload

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Приём данных
	router.POST("/collect", func(c *gin.Context) {
		var payload DataPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		collectedData = append(collectedData, payload)
		c.JSON(http.StatusOK, gin.H{"status": "data received"})
	})

	// Отображение всех собранных данных
	router.GET("/data", func(c *gin.Context) {
		c.JSON(http.StatusOK, collectedData)
	})

	return router
}
