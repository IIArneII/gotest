// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	SendMessage(params *SendMessageParams, opts ...ClientOption) (*SendMessageOK, error)

	SetAPIKey(params *SetAPIKeyParams, opts ...ClientOption) (*SetAPIKeyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  SendMessage send message API
*/
func (a *Client) SendMessage(params *SendMessageParams, opts ...ClientOption) (*SendMessageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendMessageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "sendMessage",
		Method:             "POST",
		PathPattern:        "/send-message",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SendMessageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SendMessageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SendMessageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  SetAPIKey set Api key API
*/
func (a *Client) SetAPIKey(params *SetAPIKeyParams, opts ...ClientOption) (*SetAPIKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetAPIKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "setApiKey",
		Method:             "POST",
		PathPattern:        "/set-api-key",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetAPIKeyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetAPIKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SetAPIKeyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
