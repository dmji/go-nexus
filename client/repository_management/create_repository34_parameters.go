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

// NewCreateRepository34Params creates a new CreateRepository34Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRepository34Params() *CreateRepository34Params {
	return &CreateRepository34Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRepository34ParamsWithTimeout creates a new CreateRepository34Params object
// with the ability to set a timeout on a request.
func NewCreateRepository34ParamsWithTimeout(timeout time.Duration) *CreateRepository34Params {
	return &CreateRepository34Params{
		timeout: timeout,
	}
}

// NewCreateRepository34ParamsWithContext creates a new CreateRepository34Params object
// with the ability to set a context for a request.
func NewCreateRepository34ParamsWithContext(ctx context.Context) *CreateRepository34Params {
	return &CreateRepository34Params{
		Context: ctx,
	}
}

// NewCreateRepository34ParamsWithHTTPClient creates a new CreateRepository34Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRepository34ParamsWithHTTPClient(client *http.Client) *CreateRepository34Params {
	return &CreateRepository34Params{
		HTTPClient: client,
	}
}

/*
CreateRepository34Params contains all the parameters to send to the API endpoint

	for the create repository 34 operation.

	Typically these are written to a http.Request.
*/
type CreateRepository34Params struct {

	// Body.
	Body *models.PypiHostedRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 34 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository34Params) WithDefaults() *CreateRepository34Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 34 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository34Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 34 params
func (o *CreateRepository34Params) WithTimeout(timeout time.Duration) *CreateRepository34Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 34 params
func (o *CreateRepository34Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 34 params
func (o *CreateRepository34Params) WithContext(ctx context.Context) *CreateRepository34Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 34 params
func (o *CreateRepository34Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 34 params
func (o *CreateRepository34Params) WithHTTPClient(client *http.Client) *CreateRepository34Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 34 params
func (o *CreateRepository34Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 34 params
func (o *CreateRepository34Params) WithBody(body *models.PypiHostedRepositoryAPIRequest) *CreateRepository34Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 34 params
func (o *CreateRepository34Params) SetBody(body *models.PypiHostedRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRepository34Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
