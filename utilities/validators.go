package utilities

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
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
