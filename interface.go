package translate

type Header struct {
	MessageType int32  `json:"message_type"`
	MessageID   string `json:"message_id"`
	Action      string `json:"action"`
}

type Call interface {
	Init(messageType int32, messageID string, action string)
	Head() Header
}

func GetHeader(call Call) Header {
	return call.Head()
}
