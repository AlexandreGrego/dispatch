///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package serviceaccount

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// NewUpdateServiceAccountParams creates a new UpdateServiceAccountParams object
// with the default values initialized.
func NewUpdateServiceAccountParams() *UpdateServiceAccountParams {
	var ()
	return &UpdateServiceAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateServiceAccountParamsWithTimeout creates a new UpdateServiceAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateServiceAccountParamsWithTimeout(timeout time.Duration) *UpdateServiceAccountParams {
	var ()
	return &UpdateServiceAccountParams{

		timeout: timeout,
	}
}

// NewUpdateServiceAccountParamsWithContext creates a new UpdateServiceAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateServiceAccountParamsWithContext(ctx context.Context) *UpdateServiceAccountParams {
	var ()
	return &UpdateServiceAccountParams{

		Context: ctx,
	}
}

// NewUpdateServiceAccountParamsWithHTTPClient creates a new UpdateServiceAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateServiceAccountParamsWithHTTPClient(client *http.Client) *UpdateServiceAccountParams {
	var ()
	return &UpdateServiceAccountParams{
		HTTPClient: client,
	}
}

/*UpdateServiceAccountParams contains all the parameters to send to the API endpoint
for the update service account operation typically these are written to a http.Request
*/
type UpdateServiceAccountParams struct {

	/*Body
	  Service Account object

	*/
	Body *v1.ServiceAccount
	/*ServiceAccountName
	  Name of ServiceAccount to work on

	*/
	ServiceAccountName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update service account params
func (o *UpdateServiceAccountParams) WithTimeout(timeout time.Duration) *UpdateServiceAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update service account params
func (o *UpdateServiceAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update service account params
func (o *UpdateServiceAccountParams) WithContext(ctx context.Context) *UpdateServiceAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update service account params
func (o *UpdateServiceAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update service account params
func (o *UpdateServiceAccountParams) WithHTTPClient(client *http.Client) *UpdateServiceAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update service account params
func (o *UpdateServiceAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update service account params
func (o *UpdateServiceAccountParams) WithBody(body *v1.ServiceAccount) *UpdateServiceAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update service account params
func (o *UpdateServiceAccountParams) SetBody(body *v1.ServiceAccount) {
	o.Body = body
}

// WithServiceAccountName adds the serviceAccountName to the update service account params
func (o *UpdateServiceAccountParams) WithServiceAccountName(serviceAccountName string) *UpdateServiceAccountParams {
	o.SetServiceAccountName(serviceAccountName)
	return o
}

// SetServiceAccountName adds the serviceAccountName to the update service account params
func (o *UpdateServiceAccountParams) SetServiceAccountName(serviceAccountName string) {
	o.ServiceAccountName = serviceAccountName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateServiceAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param serviceAccountName
	if err := r.SetPathParam("serviceAccountName", o.ServiceAccountName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}