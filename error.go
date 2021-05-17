package translate

type Error struct {
	Header           `json:",inline"`
	ErrorCode        string `json:"code"`
	ErrorDescription string `json:"description"`
	ErrorDetails     interface{}
}

func NewError() Translate {
	return &Error{}
}
func (d *Error) Head() Header {
	return Header{
		MessageType: d.MessageType,
		MessageID:   d.MessageID,
		Action:      d.Action,
	}
}

func (d *Error) Init(messageType int32, messageID string, action string) {
	d.MessageType = messageType
	d.MessageID = messageID
	d.Action = action
}
