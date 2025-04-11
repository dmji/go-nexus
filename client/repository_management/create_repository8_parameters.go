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

	"github.com/dmji/go-nexus/models"
)

// NewCreateRepository8Params creates a new CreateRepository8Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRepository8Params() *CreateRepository8Params {
	return &CreateRepository8Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRepository8ParamsWithTimeout creates a new CreateRepository8Params object
// with the ability to set a timeout on a request.
func NewCreateRepository8ParamsWithTimeout(timeout time.Duration) *CreateRepository8Params {
	return &CreateRepository8Params{
		timeout: timeout,
	}
}

// NewCreateRepository8ParamsWithContext creates a new CreateRepository8Params object
// with the ability to set a context for a request.
func NewCreateRepository8ParamsWithContext(ctx context.Context) *CreateRepository8Params {
	return &CreateRepository8Params{
		Context: ctx,
	}
}

// NewCreateRepository8ParamsWithHTTPClient creates a new CreateRepository8Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRepository8ParamsWithHTTPClient(client *http.Client) *CreateRepository8Params {
	return &CreateRepository8Params{
		HTTPClient: client,
	}
}

/*
CreateRepository8Params contains all the parameters to send to the API endpoint

	for the create repository 8 operation.

	Typically these are written to a http.Request.
*/
type CreateRepository8Params struct {

	// Body.
	Body *models.CargoGroupRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 8 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository8Params) WithDefaults() *CreateRepository8Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 8 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository8Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 8 params
func (o *CreateRepository8Params) WithTimeout(timeout time.Duration) *CreateRepository8Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 8 params
func (o *CreateRepository8Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 8 params
func (o *CreateRepository8Params) WithContext(ctx context.Context) *CreateRepository8Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 8 params
func (o *CreateRepository8Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 8 params
func (o *CreateRepository8Params) WithHTTPClient(client *http.Client) *CreateRepository8Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 8 params
func (o *CreateRepository8Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 8 params
func (o *CreateRepository8Params) WithBody(body *models.CargoGroupRepositoryAPIRequest) *CreateRepository8Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 8 params
func (o *CreateRepository8Params) SetBody(body *models.CargoGroupRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRepository8Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
