package handlers

import (
	"encoding/json"
	"the_basics_of_messenger_handler/entities"
	"the_basics_of_messenger_handler/utilities"

	"github.com/gin-gonic/gin"
)

func ReceiveMessage(c *gin.Context) {
	receiver := entities.Receiver{}
	utilities.JSONRequisitionParser(&receiver, c)
	receiverJSON, _ := json.Marshal(receiver)
	for _, messaging := range receiver.Messaging {
		utilities.CreateLogContent("WEBHOOK", messaging.Sender.ID, string(receiverJSON))
	}
}
