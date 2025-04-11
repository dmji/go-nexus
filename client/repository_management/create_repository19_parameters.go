// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"nexus/models"
)

// NewCreateRepository19Params creates a new CreateRepository19Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRepository19Params() *CreateRepository19Params {
	return &CreateRepository19Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRepository19ParamsWithTimeout creates a new CreateRepository19Params object
// with the ability to set a timeout on a request.
func NewCreateRepository19ParamsWithTimeout(timeout time.Duration) *CreateRepository19Params {
	return &CreateRepository19Params{
		timeout: timeout,
	}
}

// NewCreateRepository19ParamsWithContext creates a new CreateRepository19Params object
// with the ability to set a context for a request.
func NewCreateRepository19ParamsWithContext(ctx context.Context) *CreateRepository19Params {
	return &CreateRepository19Params{
		Context: ctx,
	}
}

// NewCreateRepository19ParamsWithHTTPClient creates a new CreateRepository19Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRepository19ParamsWithHTTPClient(client *http.Client) *CreateRepository19Params {
	return &CreateRepository19Params{
		HTTPClient: client,
	}
}

/*
CreateRepository19Params contains all the parameters to send to the API endpoint

	for the create repository 19 operation.

	Typically these are written to a http.Request.
*/
type CreateRepository19Params struct {

	// Body.
	Body *models.DockerProxyRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 19 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository19Params) WithDefaults() *CreateRepository19Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 19 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository19Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 19 params
func (o *CreateRepository19Params) WithTimeout(timeout time.Duration) *CreateRepository19Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 19 params
func (o *CreateRepository19Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 19 params
func (o *CreateRepository19Params) WithContext(ctx context.Context) *CreateRepository19Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 19 params
func (o *CreateRepository19Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 19 params
func (o *CreateRepository19Params) WithHTTPClient(client *http.Client) *CreateRepository19Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 19 params
func (o *CreateRepository19Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 19 params
func (o *CreateRepository19Params) WithBody(body *models.DockerProxyRepositoryAPIRequest) *CreateRepository19Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 19 params
func (o *CreateRepository19Params) SetBody(body *models.DockerProxyRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRepository19Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
