package translate

import "context"

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
