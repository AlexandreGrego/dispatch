///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-swagger; DO NOT EDIT.

package secret

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.eng.vmware.com/serverless/serverless/pkg/secret-store/gen/models"
)

// DeleteSecretNoContentCode is the HTTP code returned for type DeleteSecretNoContent
const DeleteSecretNoContentCode int = 204

/*DeleteSecretNoContent Successful deletion

swagger:response deleteSecretNoContent
*/
type DeleteSecretNoContent struct {
}

// NewDeleteSecretNoContent creates DeleteSecretNoContent with default headers values
func NewDeleteSecretNoContent() *DeleteSecretNoContent {
	return &DeleteSecretNoContent{}
}

// WriteResponse to the client
func (o *DeleteSecretNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

// DeleteSecretNotFoundCode is the HTTP code returned for type DeleteSecretNotFound
const DeleteSecretNotFoundCode int = 404

/*DeleteSecretNotFound Resource Not Found if no secret exists with the given name

swagger:response deleteSecretNotFound
*/
type DeleteSecretNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteSecretNotFound creates DeleteSecretNotFound with default headers values
func NewDeleteSecretNotFound() *DeleteSecretNotFound {
	return &DeleteSecretNotFound{}
}

// WithPayload adds the payload to the delete secret not found response
func (o *DeleteSecretNotFound) WithPayload(payload *models.Error) *DeleteSecretNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete secret not found response
func (o *DeleteSecretNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSecretNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteSecretDefault generic error

swagger:response deleteSecretDefault
*/
type DeleteSecretDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteSecretDefault creates DeleteSecretDefault with default headers values
func NewDeleteSecretDefault(code int) *DeleteSecretDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteSecretDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete secret default response
func (o *DeleteSecretDefault) WithStatusCode(code int) *DeleteSecretDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete secret default response
func (o *DeleteSecretDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete secret default response
func (o *DeleteSecretDefault) WithPayload(payload *models.Error) *DeleteSecretDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete secret default response
func (o *DeleteSecretDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSecretDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
