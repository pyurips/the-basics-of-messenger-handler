package utilities

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"the_basics_of_messenger_handler/entities"
)

func JSONRequisitionParser(sender any, c *gin.Context) (error, error) {
	shouldBindJson := c.ShouldBindJSON(sender)
	if shouldBindJson != nil {
		return shouldBindJson, nil
	}

	_, err := json.Marshal(sender)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func DotEnvHandler() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file not found")
	}

	accessToken := os.Getenv("ACCESS_TOKEN")
	emulator := os.Getenv("EMULATOR")

	if accessToken == "" {
		log.Fatal("ACCESS_TOKEN is not set in .env file")
	}

	if emulator == "" {
		log.Fatal("EMULATOR is not set in .env file")
	}
}

func MessageTypeCheck(sender *entities.Sender, c *gin.Context) error {
	if sender.MessageType != "text" && sender.MessageType != "button" {
		return errors.New("invalid message type")
	}
	return nil
}
