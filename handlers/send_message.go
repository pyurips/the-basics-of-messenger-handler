package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"the_basics_of_messenger_handler/entities"

	"github.com/gin-gonic/gin"
)

func SendMessage(c *gin.Context) {
	var sender entities.Sender
	if err := c.ShouldBindJSON(&sender); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error when binding JSON"})
		return
	}

	jsonData, err := json.Marshal(sender)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when marshalling JSON"})
		return
	}

	response, err := http.Post("externalAPIURL", "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when making a request to the external API"})
		return
	}
	defer response.Body.Close()
	c.JSON(response.StatusCode, nil)
}
