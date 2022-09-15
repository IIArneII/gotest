package api

import (
	"log"
	"test-server/internal/api/restapi/models"
	"test-server/internal/api/restapi/restapi"
	"test-server/internal/api/restapi/restapi/operations"
	"test-server/internal/app"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
)

type Config struct {
	Host     string
	Port     int
	BasePath string
}

type service struct {
	App *app.App
}

func (svc service) CreateAPIKey(params operations.CreateAPIKeyParams) middleware.Responder {
	log.Println("HTTP\tCreateAPIKey")
	key, err := svc.App.CreateAPIKey()

	if err != nil {
		log.Printf("HTTP\tCreateAPIKey: %s", err)
		return operations.NewCreateAPIKeyDefault(500)
	}
	return operations.NewCreateAPIKeyOK().WithPayload(&models.UUID{UUID: &key})
}

func (svc service) SendMessage(params operations.SendMessageParams) middleware.Responder {
	log.Println("HTTP\tSendMessage")
	err := svc.App.SendMessage(*params.Body.UUID, params.Body.Message)

	if err != nil {
		log.Printf("HTTP\tSendMessage: %s", err)
		return operations.NewSendMessageDefault(500)
	}
	return operations.NewSendMessageOK()
}

func NewServer(app *app.App, cfg Config) (*restapi.Server, error) {
	svc := &service{
		App: app,
	}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, err
	}

	if cfg.BasePath == "" {
		cfg.BasePath = swaggerSpec.BasePath()
	}
	swaggerSpec.Spec().BasePath = cfg.BasePath

	api := operations.NewTestServerAPI(swaggerSpec)

	api.CreateAPIKeyHandler = operations.CreateAPIKeyHandlerFunc(svc.CreateAPIKey)
	api.SendMessageHandler = operations.SendMessageHandlerFunc(svc.SendMessage)

	server := restapi.NewServer(api)
	server.Host = cfg.Host
	server.Port = cfg.Port

	return server, nil
}
