package xgrpc

import (
	"context"
	"log"
	"test-server/internal/app"
)

type Handler struct {
	App *app.App
}

func (h *Handler) VerifyClient(ctx context.Context, msg *Verify) (*VerifyAnswer, error) {
	log.Println("GRPC\tVerifyClient\tUuid: ", msg.Uuid)
	if _, ok := h.App.Clients[msg.Uuid]; !ok {
		log.Println("GRPC\tVerifyClient\tFalse")
		return &VerifyAnswer{Success: false}, nil
	}
	log.Println("GRPC\tVerifyClient\tTrue")
	return &VerifyAnswer{Success: true}, nil
}

func (h *Handler) SendMessage(ctx context.Context, msg *Message) (*MessageAnswer, error) {
	log.Println("GRPC\tSendMessage\tMsg: ", msg.Msg)
	return &MessageAnswer{
		Msg: "Message received:" + msg.Msg,
	}, nil
}

func (h *Handler) SendMessageToOtherClient(ctx context.Context, msg *MessageToOther) (*MessageToOtherAnswer, error) {
	log.Println("GRPC\tSendMessageToOtherClient\tUuid: ", msg.Uuid, "\tMsg: ", msg.Msg)
	return &MessageToOtherAnswer{
		Success: true,
	}, nil
}

func (Handler) mustEmbedUnimplementedTestServiceServer() {}

func New(app *app.App) *Handler {
	return &Handler{
		App: app,
	}
}
