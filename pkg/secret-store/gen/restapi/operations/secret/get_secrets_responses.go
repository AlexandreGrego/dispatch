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

// GetSecretsOKCode is the HTTP code returned for type GetSecretsOK
const GetSecretsOKCode int = 200

/*GetSecretsOK An array of registered secrets

swagger:response getSecretsOK
*/
type GetSecretsOK struct {

	/*
	  In: Body
	*/
	Payload models.GetSecretsOKBody `json:"body,omitempty"`
}

// NewGetSecretsOK creates GetSecretsOK with default headers values
func NewGetSecretsOK() *GetSecretsOK {
	return &GetSecretsOK{}
}

// WithPayload adds the payload to the get secrets o k response
func (o *GetSecretsOK) WithPayload(payload models.GetSecretsOKBody) *GetSecretsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get secrets o k response
func (o *GetSecretsOK) SetPayload(payload models.GetSecretsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSecretsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.GetSecretsOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetSecretsDefault Standard error

swagger:response getSecretsDefault
*/
type GetSecretsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSecretsDefault creates GetSecretsDefault with default headers values
func NewGetSecretsDefault(code int) *GetSecretsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSecretsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get secrets default response
func (o *GetSecretsDefault) WithStatusCode(code int) *GetSecretsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get secrets default response
func (o *GetSecretsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get secrets default response
func (o *GetSecretsDefault) WithPayload(payload *models.Error) *GetSecretsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get secrets default response
func (o *GetSecretsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSecretsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
