package emulator

import (
	"bytes"
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var USERS [4]string = [4]string{"100", "101", "102", "103"}
var ACCESSTOKEN string = "1234567890"

func InitializeEmulator() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		accessToken := c.Query("access_token")
		recipient := Payload{}
		if err := c.ShouldBindJSON(&recipient); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
			return
		}

		if accessToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "access token is required"})
			return
		}

		if accessToken != ACCESSTOKEN {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid access token"})
			return
		}

		if err := usersCheck(recipient.Recipient.ID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user"})
			return
		}

		c.JSON(http.StatusOK, nil)
	})
	go initializeWebhook()
	r.Run(":8081")
}

func initializeWebhook() {
	for {
		http.Post("http://localhost:8080/v1/receive", "application/json", getWebhookData())
		randomSeconds := rand.Intn(11) + 10
		time.Sleep(time.Duration(randomSeconds) * time.Second)
	}
}

func usersCheck(userId string) error {
	for _, user := range USERS {
		if user == userId {
			return nil
		}
	}
	return errors.New("user is valid")
}

func getWebhookData() *bytes.Buffer {
	userId := USERS[rand.Intn(4)]

	webHookData := [2]WebhookRequest{
		{
			ID:   "PAGE_ID",
			Time: time.Now().Unix(),
			Messaging: []Messaging{
				{
					Sender:    Recipient{ID: userId},
					Recipient: Recipient{ID: "PAGE_ID"},
					Timestamp: time.Now().Unix(),
					Message: &Message{
						Mid:  "mid.1457764197618:41d102a3e1ae206a38",
						Text: "ping",
					},
				},
			},
		},
		{
			ID:   "PAGE_ID",
			Time: time.Now().Unix(),
			Messaging: []Messaging{
				{
					Sender:    Recipient{ID: userId},
					Recipient: Recipient{ID: "PAGE_ID"},
					Timestamp: time.Now().Unix(),
					Postback: &Postback{
						Mid:     "mid.1457764197618:41d102a3e1ae206a38",
						Payload: "START_PAYLOAD",
					},
				},
			},
		},
	}[rand.Intn(2)]

	jsonDataText, _ := json.Marshal(webHookData)
	return bytes.NewBuffer(jsonDataText)
}

type Messaging struct {
	Sender    Recipient `json:"sender"`
	Recipient Recipient `json:"recipient"`
	Timestamp int64     `json:"timestamp"`
	Message   *Message  `json:"message,omitempty"`
	Postback  *Postback `json:"postback,omitempty"`
}

type Message struct {
	Mid  string `json:"mid"`
	Text string `json:"text"`
}

type Postback struct {
	Mid     string `json:"mid"`
	Payload string `json:"payload"`
}

type Recipient struct {
	ID string `json:"id"`
}

type WebhookRequest struct {
	ID        string      `json:"id"`
	Time      int64       `json:"time"`
	Messaging []Messaging `json:"messaging"`
}

type Payload struct {
	Recipient Recipient `json:"recipient"`
	Message   Message   `json:"message"`
}
