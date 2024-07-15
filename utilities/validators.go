package utilities

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func JSONRequisitionParser(sender any, c *gin.Context) ([]byte, error, error) {
	shouldBindJson := c.ShouldBindJSON(sender)
	if shouldBindJson != nil {
		return nil, shouldBindJson, nil
	}

	jsonData, err := json.Marshal(sender)
	if err != nil {
		return nil, nil, err
	}

	return jsonData, nil, nil
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
