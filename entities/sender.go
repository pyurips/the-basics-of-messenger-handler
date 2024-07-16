package entities

import (
	"bytes"
	"encoding/json"
	"net/http"
	"the_basics_of_messenger_handler/configs"
)

type Content struct {
	Text    string   `json:"text"`
	Buttons []Button `json:"buttons"`
}

type Sender struct {
	UserId      string  `json:"user_id"`
	MessageType string  `json:"message_type"`
	Content     Content `json:"content"`
}

type Recipient struct {
	ID string `json:"id"`
}

type MessageText struct {
	Text string `json:"text"`
}

type RequestText struct {
	Recipient Recipient   `json:"recipient"`
	Message   MessageText `json:"message"`
}

type Button struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Payload string `json:"payload"`
}

type Payload struct {
	TemplateType string   `json:"template_type"`
	Text         string   `json:"text"`
	Buttons      []Button `json:"buttons"`
}

type Attachment struct {
	Type    string  `json:"type"`
	Payload Payload `json:"payload"`
}

type MessageButton struct {
	Attachment Attachment `json:"attachment"`
}

type RequestButton struct {
	Recipient Recipient     `json:"recipient"`
	Message   MessageButton `json:"message"`
}

func (s *Sender) SendText() (*http.Response, error) {
	requestText := RequestText{
		Recipient: Recipient{
			ID: s.UserId,
		},
		Message: MessageText{
			Text: s.Content.Text,
		},
	}
	jsonRequestData, _ := json.Marshal(requestText)
	endpoint := configs.GetAPIEndpoint()
	response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonRequestData))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return response, nil
}

func (s *Sender) SendButton() (*http.Response, error) {
	requestButton := RequestButton{
		Recipient: Recipient{
			ID: s.UserId,
		},
		Message: MessageButton{
			Attachment: Attachment{
				Type: "template",
				Payload: Payload{
					TemplateType: "button",
					Text:         s.Content.Text,
					Buttons:      s.Content.Buttons,
				},
			},
		},
	}
	jsonData, _ := json.Marshal(requestButton)
	endpoint := configs.GetAPIEndpoint()
	response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return response, nil
}
