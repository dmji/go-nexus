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

// NewCreateRepository17Params creates a new CreateRepository17Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRepository17Params() *CreateRepository17Params {
	return &CreateRepository17Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRepository17ParamsWithTimeout creates a new CreateRepository17Params object
// with the ability to set a timeout on a request.
func NewCreateRepository17ParamsWithTimeout(timeout time.Duration) *CreateRepository17Params {
	return &CreateRepository17Params{
		timeout: timeout,
	}
}

// NewCreateRepository17ParamsWithContext creates a new CreateRepository17Params object
// with the ability to set a context for a request.
func NewCreateRepository17ParamsWithContext(ctx context.Context) *CreateRepository17Params {
	return &CreateRepository17Params{
		Context: ctx,
	}
}

// NewCreateRepository17ParamsWithHTTPClient creates a new CreateRepository17Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRepository17ParamsWithHTTPClient(client *http.Client) *CreateRepository17Params {
	return &CreateRepository17Params{
		HTTPClient: client,
	}
}

/*
CreateRepository17Params contains all the parameters to send to the API endpoint

	for the create repository 17 operation.

	Typically these are written to a http.Request.
*/
type CreateRepository17Params struct {

	// Body.
	Body *models.DockerGroupRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 17 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository17Params) WithDefaults() *CreateRepository17Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 17 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository17Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 17 params
func (o *CreateRepository17Params) WithTimeout(timeout time.Duration) *CreateRepository17Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 17 params
func (o *CreateRepository17Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 17 params
func (o *CreateRepository17Params) WithContext(ctx context.Context) *CreateRepository17Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 17 params
func (o *CreateRepository17Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 17 params
func (o *CreateRepository17Params) WithHTTPClient(client *http.Client) *CreateRepository17Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 17 params
func (o *CreateRepository17Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 17 params
func (o *CreateRepository17Params) WithBody(body *models.DockerGroupRepositoryAPIRequest) *CreateRepository17Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 17 params
func (o *CreateRepository17Params) SetBody(body *models.DockerGroupRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRepository17Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
