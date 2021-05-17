package translate

import "context"

type callType struct{}
type messageID struct{}
type actionName struct{}

func WithMessageID(ctx context.Context, m string) context.Context {
	return context.WithValue(ctx, messageID{}, m)
}

func MessageIDFromCtx(ctx context.Context) string {
	return ctx.Value(messageID{}).(string)
}

func WithActionName(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, actionName{}, name)
}

func ActionNameFromCtx(ctx context.Context) string {
	return ctx.Value(actionName{}).(string)
}
