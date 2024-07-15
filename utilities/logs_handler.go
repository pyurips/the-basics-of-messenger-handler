package utilities

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CreateLogContent(logType string, userId string, pageId string, content string) {
	fileName := ""
	if pageId != "" {
		fileName = fmt.Sprintf("logs/%s:%s.log", userId, pageId)
	} else {
		fileName = fmt.Sprintf("logs/%s:nil.log", userId)
	}

	os.MkdirAll(filepath.Dir(fileName), 0755)
	file, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	logger := log.New(file, fmt.Sprintf("[%s]", logType), log.LUTC)
	logMessage := fmt.Sprintf("content: %s", content)
	logger.Println(logMessage)
}
