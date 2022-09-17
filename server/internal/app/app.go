package app

import (
	"errors"
	"test-server/internal/api/xgrpc"

	uuid "github.com/satori/go.uuid"
)

type App struct {
	Clients map[string]xgrpc.Client
}

func (app *App) CreateAPIKey() (string, error) {
	key := uuid.NewV4().String()
	app.Clients[key] = xgrpc.Client{Uuid: key}
	return key, nil
}

func (app *App) SendMessage(uuid, msg string) error {
	if client, ok := app.Clients[uuid]; ok {
		client.StreamSendMessage.Send(&xgrpc.MessageAnswer{
			Msg: msg,
		})
		return nil
	}
	return errors.New("Not found")
}
