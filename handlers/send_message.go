package handlers

import (
	"bytes"
	"net/http"

	"the_basics_of_messenger_handler/entities"
	"the_basics_of_messenger_handler/utilities"

	"github.com/gin-gonic/gin"
)

func SendMessage(c *gin.Context) {
	sender := new(entities.Sender)
	jsonData, bindError, marshalError := utilities.JSONRequisitionParser(sender, c)
	if bindError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error when binding JSON"})
		return
	}
	if marshalError != nil {
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
