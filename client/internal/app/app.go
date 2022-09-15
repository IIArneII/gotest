package app

import (
	"context"
	"errors"
	"test-client/internal/api/xgrpc"
)

type App struct {
	UUID string
	GRPC xgrpc.TestServiceClient
}

func (app *App) SetAPIKey(ctx context.Context, key string) error {
	answer, err := app.GRPC.VerifyClient(ctx, &xgrpc.Verify{Uuid: key})
	if err != nil {
		return err
	}
	if !answer.Success {
		return errors.New("Нихуя")
	}
	app.UUID = key
	return nil
}

func (app *App) SendMessage(ctx context.Context, uuid, msg string) error {
	answer, err := app.GRPC.SendMessageToOtherClient(ctx, &xgrpc.MessageToOther{
		Uuid: uuid,
		Msg:  msg,
	})
	if err != nil {
		return err
	}
	if !answer.Success {
		return errors.New("Нихуя")
	}
	return nil
}
