package api

import (
	"log"
	"test-client/internal/api/restapi/restapi"
	"test-client/internal/api/restapi/restapi/operations"
	"test-client/internal/app"

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

func (svc service) SetAPIKey(params operations.SetAPIKeyParams) middleware.Responder {
	params.HTTPRequest.Context()
	log.Println("HTTP\tSetAPIKey")
	err := svc.App.SetAPIKey(params.HTTPRequest.Context(), *params.Body.UUID)

	if err != nil {
		log.Printf("HTTP\tSetAPIKey: %s", err)
		return operations.NewSetAPIKeyDefault(500)
	}
	return operations.NewSetAPIKeyOK()
}

func (svc service) SendMessage(params operations.SendMessageParams) middleware.Responder {
	log.Println("HTTP\tSendMessage")
	err := svc.App.SendMessage(params.HTTPRequest.Context(), params.Body.UUID, *params.Body.Message)

	if err != nil {
		log.Printf("HTTP\tSendMessage: %s", err)
		return operations.NewSendMessageDefault(500)
	}
	return operations.NewSendMessageOK()
}

func NewServer(app app.App, cfg Config) (*restapi.Server, error) {
	svc := &service{
		App: &app,
	}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, err
	}

	if cfg.BasePath == "" {
		cfg.BasePath = swaggerSpec.BasePath()
	}
	swaggerSpec.Spec().BasePath = cfg.BasePath

	api := operations.NewTestClientAPI(swaggerSpec)

	api.SetAPIKeyHandler = operations.SetAPIKeyHandlerFunc(svc.SetAPIKey)
	api.SendMessageHandler = operations.SendMessageHandlerFunc(svc.SendMessage)

	server := restapi.NewServer(api)
	server.Host = cfg.Host
	server.Port = cfg.Port

	return server, nil
}
