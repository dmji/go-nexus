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

// NewPutRepositoriesNpmProxyByRepositorynameParams creates a new PutRepositoriesNpmProxyByRepositorynameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutRepositoriesNpmProxyByRepositorynameParams() *PutRepositoriesNpmProxyByRepositorynameParams {
	return &PutRepositoriesNpmProxyByRepositorynameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutRepositoriesNpmProxyByRepositorynameParamsWithTimeout creates a new PutRepositoriesNpmProxyByRepositorynameParams object
// with the ability to set a timeout on a request.
func NewPutRepositoriesNpmProxyByRepositorynameParamsWithTimeout(timeout time.Duration) *PutRepositoriesNpmProxyByRepositorynameParams {
	return &PutRepositoriesNpmProxyByRepositorynameParams{
		timeout: timeout,
	}
}

// NewPutRepositoriesNpmProxyByRepositorynameParamsWithContext creates a new PutRepositoriesNpmProxyByRepositorynameParams object
// with the ability to set a context for a request.
func NewPutRepositoriesNpmProxyByRepositorynameParamsWithContext(ctx context.Context) *PutRepositoriesNpmProxyByRepositorynameParams {
	return &PutRepositoriesNpmProxyByRepositorynameParams{
		Context: ctx,
	}
}

// NewPutRepositoriesNpmProxyByRepositorynameParamsWithHTTPClient creates a new PutRepositoriesNpmProxyByRepositorynameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutRepositoriesNpmProxyByRepositorynameParamsWithHTTPClient(client *http.Client) *PutRepositoriesNpmProxyByRepositorynameParams {
	return &PutRepositoriesNpmProxyByRepositorynameParams{
		HTTPClient: client,
	}
}

/*
PutRepositoriesNpmProxyByRepositorynameParams contains all the parameters to send to the API endpoint

	for the put repositories npm proxy by repositoryname operation.

	Typically these are written to a http.Request.
*/
type PutRepositoriesNpmProxyByRepositorynameParams struct {

	// Body.
	Body *models.NpmProxyRepositoryAPIRequest

	/* RepositoryName.

	   Name of the repository to update
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put repositories npm proxy by repositoryname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRepositoriesNpmProxyByRepositorynameParams) WithDefaults() *PutRepositoriesNpmProxyByRepositorynameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put repositories npm proxy by repositoryname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRepositoriesNpmProxyByRepositorynameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put repositories npm proxy by repositoryname params
func (o *PutRepositoriesNpmProxyByRepositorynameParams) WithTimeout(timeout time.Duration) *PutRepositoriesNpmProxyByRepositorynameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put repositories npm proxy by repositoryname params
func (o *PutRepositoriesNpmProxyByRepositorynameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put repositories npm proxy by repositoryname params
func (o *PutRepositoriesNpmProxyByRepositorynameParams) WithContext(ctx context.Context) *PutRepositoriesNpmProxyByRepositorynameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put repositories npm proxy by repositoryname params
func (o *PutRepositoriesNpmProxyByRepositorynameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put repositories npm proxy by repositoryname params
func (o *PutRepositoriesNpmProxyByRepositorynameParams) WithHTTPClient(client *http.Client) *PutRepositoriesNpmProxyByRepositorynameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put repositories npm proxy by repositoryname params
func (o *PutRepositoriesNpmProxyByRepositorynameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put repositories npm proxy by repositoryname params
func (o *PutRepositoriesNpmProxyByRepositorynameParams) WithBody(body *models.NpmProxyRepositoryAPIRequest) *PutRepositoriesNpmProxyByRepositorynameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put repositories npm proxy by repositoryname params
func (o *PutRepositoriesNpmProxyByRepositorynameParams) SetBody(body *models.NpmProxyRepositoryAPIRequest) {
	o.Body = body
}

// WithRepositoryName adds the repositoryName to the put repositories npm proxy by repositoryname params
func (o *PutRepositoriesNpmProxyByRepositorynameParams) WithRepositoryName(repositoryName string) *PutRepositoriesNpmProxyByRepositorynameParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the put repositories npm proxy by repositoryname params
func (o *PutRepositoriesNpmProxyByRepositorynameParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *PutRepositoriesNpmProxyByRepositorynameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
