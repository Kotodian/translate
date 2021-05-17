package translate

import (
	"context"
)

var (
	register = make(map[string]func() Translate)
)

type Header struct {
	MessageType int32  `json:"message_type"`
	MessageID   string `json:"message_id"`
	Action      string `json:"action"`
}

type Translate interface {
	Init(messageType int32, messageID string, action string)
	Head() Header
}

func GetHeader(call Translate) Header {
	return call.Head()
}
func InitHeader(ctx context.Context, callType int32, trans Translate) {
	trans.Init(callType, MessageIDFromCtx(ctx), ActionNameFromCtx(ctx))
}

func New(header Header) Translate {
	if header.MessageType == 2 {
		return register[header.Action+"Request"]()
	} else if header.MessageType == 3 {
		return register[header.Action+"Response"]()
	} else {
		return register[header.Action+"Error"]()
	}
}
