package entities

type Buttons struct {
	ButtonType string `json:"type"`
	Title      string `json:"title"`
	Payload    string `json:"payload"`
}

type Content struct {
	Text    string    `json:"text"`
	Buttons []Buttons `json:"buttons"`
}

type Sender struct {
	UserId      string  `json:"user_id"`
	MessageType string  `json:"message_type"`
	Content     Content `json:"content"`
}

func (s *Sender) SendText(content string) {
	s.MessageType = "text"
	s.Content.Text = content
}

func (s *Sender) SendButton(content string, buttons []Buttons) {
	s.MessageType = "button"
	s.Content.Text = content
	s.Content.Buttons = buttons
}
