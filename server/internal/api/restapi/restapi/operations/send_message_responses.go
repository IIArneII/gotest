// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"test-server/internal/api/restapi/models"
)

// SendMessageOKCode is the HTTP code returned for type SendMessageOK
const SendMessageOKCode int = 200

/*SendMessageOK Ok

swagger:response sendMessageOK
*/
type SendMessageOK struct {
}

// NewSendMessageOK creates SendMessageOK with default headers values
func NewSendMessageOK() *SendMessageOK {

	return &SendMessageOK{}
}

// WriteResponse to the client
func (o *SendMessageOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*SendMessageDefault error

swagger:response sendMessageDefault
*/
type SendMessageDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSendMessageDefault creates SendMessageDefault with default headers values
func NewSendMessageDefault(code int) *SendMessageDefault {
	if code <= 0 {
		code = 500
	}

	return &SendMessageDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the send message default response
func (o *SendMessageDefault) WithStatusCode(code int) *SendMessageDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the send message default response
func (o *SendMessageDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the send message default response
func (o *SendMessageDefault) WithPayload(payload *models.Error) *SendMessageDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the send message default response
func (o *SendMessageDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SendMessageDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}