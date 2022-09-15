package app

import (
	uuid "github.com/satori/go.uuid"
)

type Client struct {
	uuid string
}

type App struct {
	Clients map[string]Client
}

func (app *App) CreateAPIKey() (string, error) {
	key := uuid.NewV4().String()
	app.Clients[key] = Client{uuid: key}
	return key, nil
}

func (app *App) SendMessage(uuid, msg string) error {
	return nil
}
