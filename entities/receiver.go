package entities

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

type Receiver struct {
	ID        string      `json:"id"`
	Time      int64       `json:"time"`
	Messaging []Messaging `json:"messaging"`
}
