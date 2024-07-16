package handlers

import (
	"net/http"
	"strconv"

	"the_basics_of_messenger_handler/entities"
	"the_basics_of_messenger_handler/utilities"

	"github.com/gin-gonic/gin"
)

func SendMessage(c *gin.Context) {
	sender := entities.Sender{}
	bindError, marshalError := utilities.JSONRequisitionParser(&sender, c)
	if bindError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error when binding JSON"})
		return
	}
	if marshalError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when marshalling JSON"})
		utilities.CreateLogContent(strconv.Itoa(http.StatusInternalServerError), sender.UserId, marshalError.Error())
		return
	}

	messageTypeError := utilities.MessageTypeCheck(&sender, c)
	if messageTypeError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message type"})
		utilities.CreateLogContent(strconv.Itoa(http.StatusBadRequest), sender.UserId, messageTypeError.Error())
		return
	}

	if sender.MessageType == "text" {
		response, err := sender.SendText()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when making a request to the external API"})
			utilities.CreateLogContent(strconv.Itoa(http.StatusInternalServerError), sender.UserId, err.Error())
			return
		}
		c.JSON(response.StatusCode, nil)
	}

	if sender.MessageType == "button" {
		response, err := sender.SendButton()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when making a request to the external API"})
			utilities.CreateLogContent(strconv.Itoa(http.StatusInternalServerError), sender.UserId, err.Error())
			return
		}
		c.JSON(response.StatusCode, nil)
	}
}
