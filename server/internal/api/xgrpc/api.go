package xgrpc

import (
	"context"
	"fmt"
)

type Handler struct{}

func (h *Handler) VerifyClient(ctx context.Context, msg *Verify) (*VerifyAnswer, error) {
	fmt.Printf("VerifyClient: %s\n", msg.Uuid)
	return &VerifyAnswer{
		Success: true,
	}, nil
}

func (h *Handler) SendMessage(ctx context.Context, msg *Message) (*MessageAnswer, error) {
	fmt.Printf("SendMessage: %s\n", msg.Msg)
	return &MessageAnswer{
		Msg: "Message received:" + msg.Msg,
	}, nil
}

func (h *Handler) SendMessageToOtherClient(ctx context.Context, msg *MessageToOther) (*MessageToOtherAnswer, error) {
	fmt.Printf("SendMessageToOtherClient:\n\tUUID: %s\n\tMsg: %s\n", msg.Uuid, msg.Msg)
	return &MessageToOtherAnswer{
		Success: true,
	}, nil
}

func (Handler) mustEmbedUnimplementedTestServiceServer() {
	fmt.Printf("Заебал")
}

func New() *Handler {
	return &Handler{}
}
