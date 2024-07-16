package emulator

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeEmulator() {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		accessToken := c.Query("access_token")
		if accessToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "access token is required"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	r.Run(":8081")
}

func initializeWebhook() {}
