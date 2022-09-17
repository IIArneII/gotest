package xgrpc

import (
	"context"
	"io"
	"log"
)

type App struct {
	Clients map[string]Client
	// Надо добавить мьютексы
}

type Client struct {
	Uuid                           string
	StreamSendMessage              TestService_SendMessageServer
	StreamSendMessageToOtherClient TestService_SendMessageToOtherClientServer
}

type Handler struct {
	App *App
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

func (h *Handler) SendMessage(stream TestService_SendMessageServer) error {
	msg, err := stream.Recv()
	if err == io.EOF {
		log.Println("GRPC\tEOF")
		return nil
	}
	if err != nil {
		log.Println("GRPC\tError: ", err)
		return err
	}

	s := h.App.Clients[msg.Uuid]
	s.StreamSendMessage = stream
	h.App.Clients[msg.Uuid] = s

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Println("GRPC\tEOF")
			return nil
		}
		if err != nil {
			log.Println("GRPC\tError: ", err)
			return err
		}
		log.Println("GRPC\tSendMessage\tMsg: ", msg.Msg, "\tUuid: ", msg.Uuid)

		err = stream.Send(&MessageAnswer{
			Msg: "Message received: " + msg.Msg,
		})
		if err == io.EOF {
			log.Println("GRPC\tEOF")
			return nil
		}
		if err != nil {
			log.Println("GRPC\tError: ", err)
			return err
		}
	}
}

func (h *Handler) SendMessageToOtherClient(stream TestService_SendMessageToOtherClientServer) error {
	msg, err := stream.Recv()
	if err == io.EOF {
		log.Println("GRPC\tEOF")
		return nil
	}
	if err != nil {
		log.Println("GRPC\tError: ", err)
		return err
	}

	s := h.App.Clients[msg.Uuid]
	s.StreamSendMessageToOtherClient = stream
	h.App.Clients[msg.Uuid] = s

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Println("GRPC\tEOF")
			return nil
		}
		if err != nil {
			log.Println("GRPC\tError: ", err)
			return err
		}
		log.Println("GRPC\tSendMessage\tMsg: ", msg.Msg, "\tUuid: ", msg.Uuid, "\tTarget: ", msg.TargetUuid)

		if client, ok := h.App.Clients[msg.TargetUuid]; ok {
			client.StreamSendMessage.Send(&MessageAnswer{
				Msg: msg.Msg,
			})
			stream.Send(&MessageToOtherAnswer{
				Success: true,
			})
		}
		stream.Send(&MessageToOtherAnswer{
			Success: false,
		})
	}
}

func (Handler) mustEmbedUnimplementedTestServiceServer() {}

func New(app *App) *Handler {
	return &Handler{
		App: app,
	}
}
