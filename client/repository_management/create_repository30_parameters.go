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

// NewCreateRepository30Params creates a new CreateRepository30Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRepository30Params() *CreateRepository30Params {
	return &CreateRepository30Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRepository30ParamsWithTimeout creates a new CreateRepository30Params object
// with the ability to set a timeout on a request.
func NewCreateRepository30ParamsWithTimeout(timeout time.Duration) *CreateRepository30Params {
	return &CreateRepository30Params{
		timeout: timeout,
	}
}

// NewCreateRepository30ParamsWithContext creates a new CreateRepository30Params object
// with the ability to set a context for a request.
func NewCreateRepository30ParamsWithContext(ctx context.Context) *CreateRepository30Params {
	return &CreateRepository30Params{
		Context: ctx,
	}
}

// NewCreateRepository30ParamsWithHTTPClient creates a new CreateRepository30Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRepository30ParamsWithHTTPClient(client *http.Client) *CreateRepository30Params {
	return &CreateRepository30Params{
		HTTPClient: client,
	}
}

/*
CreateRepository30Params contains all the parameters to send to the API endpoint

	for the create repository 30 operation.

	Typically these are written to a http.Request.
*/
type CreateRepository30Params struct {

	// Body.
	Body *models.NugetHostedRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 30 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository30Params) WithDefaults() *CreateRepository30Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 30 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository30Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 30 params
func (o *CreateRepository30Params) WithTimeout(timeout time.Duration) *CreateRepository30Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 30 params
func (o *CreateRepository30Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 30 params
func (o *CreateRepository30Params) WithContext(ctx context.Context) *CreateRepository30Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 30 params
func (o *CreateRepository30Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 30 params
func (o *CreateRepository30Params) WithHTTPClient(client *http.Client) *CreateRepository30Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 30 params
func (o *CreateRepository30Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 30 params
func (o *CreateRepository30Params) WithBody(body *models.NugetHostedRepositoryAPIRequest) *CreateRepository30Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 30 params
func (o *CreateRepository30Params) SetBody(body *models.NugetHostedRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRepository30Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
