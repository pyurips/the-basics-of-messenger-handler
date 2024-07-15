package handlers

import (
	"net/http"

	"the_basics_of_messenger_handler/entities"
	"the_basics_of_messenger_handler/utilities"

	"github.com/gin-gonic/gin"
)

func SendMessage(c *gin.Context) {
	sender := entities.Sender{}
	bindError, marshalError := utilities.JSONRequisitionParser(&sender, c)
	if bindError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error when binding JSON"})
		utilities.CreateLogContent(entities.ERROR, "sender.UserId", "PAGE_ID", bindError.Error())
		return
	}
	if marshalError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when marshalling JSON"})
		utilities.CreateLogContent(entities.ERROR, sender.UserId, "PAGE_ID", marshalError.Error())
		return
	}

	utilities.MessageTypeCheck(&sender, c)

	if sender.MessageType == "text" {
		response, err := sender.SendText()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when making a request to the external API"})
			utilities.CreateLogContent(entities.ERROR, sender.UserId, "PAGE_ID", err.Error())
			return
		}
		c.JSON(response.StatusCode, nil)
	}

	if sender.MessageType == "button" {
		response, err := sender.SendButton()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when making a request to the external API"})
			utilities.CreateLogContent(entities.ERROR, sender.UserId, "PAGE_ID", err.Error())
			return
		}
		c.JSON(response.StatusCode, nil)
	}
}
