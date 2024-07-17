package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"/v1/message": "Use this endpoint to send a new message.",
		"/v1/receive": "Use this endpoint to receive messages via webhook.",
		"obs":         "Please refer to the README.md for more information.",
	})
}
