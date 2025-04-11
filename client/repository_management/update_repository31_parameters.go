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

// NewUpdateRepository31Params creates a new UpdateRepository31Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRepository31Params() *UpdateRepository31Params {
	return &UpdateRepository31Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRepository31ParamsWithTimeout creates a new UpdateRepository31Params object
// with the ability to set a timeout on a request.
func NewUpdateRepository31ParamsWithTimeout(timeout time.Duration) *UpdateRepository31Params {
	return &UpdateRepository31Params{
		timeout: timeout,
	}
}

// NewUpdateRepository31ParamsWithContext creates a new UpdateRepository31Params object
// with the ability to set a context for a request.
func NewUpdateRepository31ParamsWithContext(ctx context.Context) *UpdateRepository31Params {
	return &UpdateRepository31Params{
		Context: ctx,
	}
}

// NewUpdateRepository31ParamsWithHTTPClient creates a new UpdateRepository31Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRepository31ParamsWithHTTPClient(client *http.Client) *UpdateRepository31Params {
	return &UpdateRepository31Params{
		HTTPClient: client,
	}
}

/*
UpdateRepository31Params contains all the parameters to send to the API endpoint

	for the update repository 31 operation.

	Typically these are written to a http.Request.
*/
type UpdateRepository31Params struct {

	// Body.
	Body *models.NugetProxyRepositoryAPIRequest

	/* RepositoryName.

	   Name of the repository to update
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update repository 31 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRepository31Params) WithDefaults() *UpdateRepository31Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update repository 31 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRepository31Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update repository 31 params
func (o *UpdateRepository31Params) WithTimeout(timeout time.Duration) *UpdateRepository31Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update repository 31 params
func (o *UpdateRepository31Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update repository 31 params
func (o *UpdateRepository31Params) WithContext(ctx context.Context) *UpdateRepository31Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update repository 31 params
func (o *UpdateRepository31Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update repository 31 params
func (o *UpdateRepository31Params) WithHTTPClient(client *http.Client) *UpdateRepository31Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update repository 31 params
func (o *UpdateRepository31Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update repository 31 params
func (o *UpdateRepository31Params) WithBody(body *models.NugetProxyRepositoryAPIRequest) *UpdateRepository31Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update repository 31 params
func (o *UpdateRepository31Params) SetBody(body *models.NugetProxyRepositoryAPIRequest) {
	o.Body = body
}

// WithRepositoryName adds the repositoryName to the update repository 31 params
func (o *UpdateRepository31Params) WithRepositoryName(repositoryName string) *UpdateRepository31Params {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the update repository 31 params
func (o *UpdateRepository31Params) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRepository31Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
