package translate

type Common struct {
	MessageType int32  `json:"message_type"`
	MessageID   string `json:"message_id"`
	Action      string `json:"action"`
}
