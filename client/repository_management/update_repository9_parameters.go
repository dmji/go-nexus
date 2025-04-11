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

// NewUpdateRepository9Params creates a new UpdateRepository9Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRepository9Params() *UpdateRepository9Params {
	return &UpdateRepository9Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRepository9ParamsWithTimeout creates a new UpdateRepository9Params object
// with the ability to set a timeout on a request.
func NewUpdateRepository9ParamsWithTimeout(timeout time.Duration) *UpdateRepository9Params {
	return &UpdateRepository9Params{
		timeout: timeout,
	}
}

// NewUpdateRepository9ParamsWithContext creates a new UpdateRepository9Params object
// with the ability to set a context for a request.
func NewUpdateRepository9ParamsWithContext(ctx context.Context) *UpdateRepository9Params {
	return &UpdateRepository9Params{
		Context: ctx,
	}
}

// NewUpdateRepository9ParamsWithHTTPClient creates a new UpdateRepository9Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRepository9ParamsWithHTTPClient(client *http.Client) *UpdateRepository9Params {
	return &UpdateRepository9Params{
		HTTPClient: client,
	}
}

/*
UpdateRepository9Params contains all the parameters to send to the API endpoint

	for the update repository 9 operation.

	Typically these are written to a http.Request.
*/
type UpdateRepository9Params struct {

	// Body.
	Body *models.CargoHostedRepositoryAPIRequest

	/* RepositoryName.

	   Name of the repository to update
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update repository 9 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRepository9Params) WithDefaults() *UpdateRepository9Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update repository 9 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRepository9Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update repository 9 params
func (o *UpdateRepository9Params) WithTimeout(timeout time.Duration) *UpdateRepository9Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update repository 9 params
func (o *UpdateRepository9Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update repository 9 params
func (o *UpdateRepository9Params) WithContext(ctx context.Context) *UpdateRepository9Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update repository 9 params
func (o *UpdateRepository9Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update repository 9 params
func (o *UpdateRepository9Params) WithHTTPClient(client *http.Client) *UpdateRepository9Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update repository 9 params
func (o *UpdateRepository9Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update repository 9 params
func (o *UpdateRepository9Params) WithBody(body *models.CargoHostedRepositoryAPIRequest) *UpdateRepository9Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update repository 9 params
func (o *UpdateRepository9Params) SetBody(body *models.CargoHostedRepositoryAPIRequest) {
	o.Body = body
}

// WithRepositoryName adds the repositoryName to the update repository 9 params
func (o *UpdateRepository9Params) WithRepositoryName(repositoryName string) *UpdateRepository9Params {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the update repository 9 params
func (o *UpdateRepository9Params) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRepository9Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param repositoryName
	if err := r.SetPathParam("repositoryName", o.RepositoryName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
