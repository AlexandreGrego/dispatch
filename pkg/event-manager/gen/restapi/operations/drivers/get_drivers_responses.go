///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package drivers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// GetDriversOKCode is the HTTP code returned for type GetDriversOK
const GetDriversOKCode int = 200

/*GetDriversOK Successful operation

swagger:response getDriversOK
*/
type GetDriversOK struct {

	/*
	  In: Body
	*/
	Payload []*v1.EventDriver `json:"body,omitempty"`
}

// NewGetDriversOK creates GetDriversOK with default headers values
func NewGetDriversOK() *GetDriversOK {

	return &GetDriversOK{}
}

// WithPayload adds the payload to the get drivers o k response
func (o *GetDriversOK) WithPayload(payload []*v1.EventDriver) *GetDriversOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get drivers o k response
func (o *GetDriversOK) SetPayload(payload []*v1.EventDriver) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDriversOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*v1.EventDriver, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetDriversUnauthorizedCode is the HTTP code returned for type GetDriversUnauthorized
const GetDriversUnauthorizedCode int = 401

/*GetDriversUnauthorized Unauthorized Request

swagger:response getDriversUnauthorized
*/
type GetDriversUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetDriversUnauthorized creates GetDriversUnauthorized with default headers values
func NewGetDriversUnauthorized() *GetDriversUnauthorized {

	return &GetDriversUnauthorized{}
}

// WithPayload adds the payload to the get drivers unauthorized response
func (o *GetDriversUnauthorized) WithPayload(payload *v1.Error) *GetDriversUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get drivers unauthorized response
func (o *GetDriversUnauthorized) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDriversUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDriversForbiddenCode is the HTTP code returned for type GetDriversForbidden
const GetDriversForbiddenCode int = 403

/*GetDriversForbidden access to this resource is forbidden

swagger:response getDriversForbidden
*/
type GetDriversForbidden struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetDriversForbidden creates GetDriversForbidden with default headers values
func NewGetDriversForbidden() *GetDriversForbidden {

	return &GetDriversForbidden{}
}

// WithPayload adds the payload to the get drivers forbidden response
func (o *GetDriversForbidden) WithPayload(payload *v1.Error) *GetDriversForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get drivers forbidden response
func (o *GetDriversForbidden) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDriversForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetDriversDefault Unknown error

swagger:response getDriversDefault
*/
type GetDriversDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetDriversDefault creates GetDriversDefault with default headers values
func NewGetDriversDefault(code int) *GetDriversDefault {
	if code <= 0 {
		code = 500
	}

	return &GetDriversDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get drivers default response
func (o *GetDriversDefault) WithStatusCode(code int) *GetDriversDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get drivers default response
func (o *GetDriversDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get drivers default response
func (o *GetDriversDefault) WithPayload(payload *v1.Error) *GetDriversDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get drivers default response
func (o *GetDriversDefault) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDriversDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
