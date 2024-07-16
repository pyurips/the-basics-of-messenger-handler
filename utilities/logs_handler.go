package utilities

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CreateLogContent(logType string, userId string, content string) {
	fileName := fmt.Sprintf("logs/%s.log", userId)

	os.MkdirAll(filepath.Dir(fileName), 0755)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer file.Close()

	logger := log.New(file, fmt.Sprintf("[%s] ", logType), log.Ldate|log.Ltime|log.LUTC)
	logMessage := fmt.Sprintf("content: %s", content)
	logger.Println(logMessage)
	loggerSeparator := log.New(file, "------------------------------------------------------", 0)
	loggerSeparator.Println()
}
