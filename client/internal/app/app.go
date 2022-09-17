package app

import (
	"context"
	"errors"
	"io"
	"log"
	"test-client/internal/api/xgrpc"
)

type App struct {
	UUID                           string
	GRPC                           xgrpc.TestServiceClient
	StreamSendMessage              xgrpc.TestService_SendMessageClient
	streamSendMessageToOtherClient xgrpc.TestService_SendMessageToOtherClientClient
}

func SendMessageListener(stream xgrpc.TestService_SendMessageClient) {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Println("SendMessageListener: EOF")
			return
		}
		if err != nil {
			log.Println("SendMessageListener: ", err)
			return
		}
		log.Println("SendMessageListener: ", msg.Msg)
	}
}

func (app *App) SetAPIKey(ctx context.Context, key string) error {
	answer, err := app.GRPC.VerifyClient(ctx, &xgrpc.Verify{Uuid: key})
	if err != nil {
		return err
	}
	if !answer.Success {
		return errors.New("Error VerifyClient")
	}

	context := context.Background()
	streamSendMessage, err := app.GRPC.SendMessage(context)
	if err != nil {
		return err
	}
	streamSendMessageToOtherClient, err := app.GRPC.SendMessageToOtherClient(context)
	if err != nil {
		return err
	}

	err = streamSendMessage.Send(&xgrpc.Message{Uuid: key})
	if err == io.EOF {
		log.Println("GRPC\tEOF")
		return err
	}
	if err != nil {
		log.Println("GRPC\tError: ", err)
		return err
	}

	err = streamSendMessageToOtherClient.Send(&xgrpc.MessageToOther{Uuid: key})
	if err == io.EOF {
		log.Println("GRPC\tEOF")
		return err
	}
	if err != nil {
		log.Println("GRPC\tError: ", err)
		return err
	}

	app.UUID = key
	app.StreamSendMessage = streamSendMessage
	app.streamSendMessageToOtherClient = streamSendMessageToOtherClient
	go SendMessageListener(streamSendMessage)

	return nil
}

func (app *App) SendMessage(ctx context.Context, uuid, msg string) error {
	log.Println("Sending a message to the client ", uuid, " ", msg)
	err := app.streamSendMessageToOtherClient.Send(&xgrpc.MessageToOther{
		TargetUuid: uuid,
		Msg:        msg,
	})
	if err == io.EOF {
		log.Println("SendMessageListener: EOF")
		return err
	}
	if err != nil {
		log.Println("SendMessageListener Error: ", err)
		return err
	}
	return nil
}
