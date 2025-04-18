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

// NewPutRepositoriesYumProxyByRepositoryNameParams creates a new PutRepositoriesYumProxyByRepositoryNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutRepositoriesYumProxyByRepositoryNameParams() *PutRepositoriesYumProxyByRepositoryNameParams {
	return &PutRepositoriesYumProxyByRepositoryNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutRepositoriesYumProxyByRepositoryNameParamsWithTimeout creates a new PutRepositoriesYumProxyByRepositoryNameParams object
// with the ability to set a timeout on a request.
func NewPutRepositoriesYumProxyByRepositoryNameParamsWithTimeout(timeout time.Duration) *PutRepositoriesYumProxyByRepositoryNameParams {
	return &PutRepositoriesYumProxyByRepositoryNameParams{
		timeout: timeout,
	}
}

// NewPutRepositoriesYumProxyByRepositoryNameParamsWithContext creates a new PutRepositoriesYumProxyByRepositoryNameParams object
// with the ability to set a context for a request.
func NewPutRepositoriesYumProxyByRepositoryNameParamsWithContext(ctx context.Context) *PutRepositoriesYumProxyByRepositoryNameParams {
	return &PutRepositoriesYumProxyByRepositoryNameParams{
		Context: ctx,
	}
}

// NewPutRepositoriesYumProxyByRepositoryNameParamsWithHTTPClient creates a new PutRepositoriesYumProxyByRepositoryNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutRepositoriesYumProxyByRepositoryNameParamsWithHTTPClient(client *http.Client) *PutRepositoriesYumProxyByRepositoryNameParams {
	return &PutRepositoriesYumProxyByRepositoryNameParams{
		HTTPClient: client,
	}
}

/*
PutRepositoriesYumProxyByRepositoryNameParams contains all the parameters to send to the API endpoint

	for the put repositories yum proxy by repository name operation.

	Typically these are written to a http.Request.
*/
type PutRepositoriesYumProxyByRepositoryNameParams struct {

	// Body.
	Body *models.YumProxyRepositoryAPIRequest

	/* RepositoryName.

	   Name of the repository to update
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put repositories yum proxy by repository name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRepositoriesYumProxyByRepositoryNameParams) WithDefaults() *PutRepositoriesYumProxyByRepositoryNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put repositories yum proxy by repository name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRepositoriesYumProxyByRepositoryNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put repositories yum proxy by repository name params
func (o *PutRepositoriesYumProxyByRepositoryNameParams) WithTimeout(timeout time.Duration) *PutRepositoriesYumProxyByRepositoryNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put repositories yum proxy by repository name params
func (o *PutRepositoriesYumProxyByRepositoryNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put repositories yum proxy by repository name params
func (o *PutRepositoriesYumProxyByRepositoryNameParams) WithContext(ctx context.Context) *PutRepositoriesYumProxyByRepositoryNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put repositories yum proxy by repository name params
func (o *PutRepositoriesYumProxyByRepositoryNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put repositories yum proxy by repository name params
func (o *PutRepositoriesYumProxyByRepositoryNameParams) WithHTTPClient(client *http.Client) *PutRepositoriesYumProxyByRepositoryNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put repositories yum proxy by repository name params
func (o *PutRepositoriesYumProxyByRepositoryNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put repositories yum proxy by repository name params
func (o *PutRepositoriesYumProxyByRepositoryNameParams) WithBody(body *models.YumProxyRepositoryAPIRequest) *PutRepositoriesYumProxyByRepositoryNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put repositories yum proxy by repository name params
func (o *PutRepositoriesYumProxyByRepositoryNameParams) SetBody(body *models.YumProxyRepositoryAPIRequest) {
	o.Body = body
}

// WithRepositoryName adds the repositoryName to the put repositories yum proxy by repository name params
func (o *PutRepositoriesYumProxyByRepositoryNameParams) WithRepositoryName(repositoryName string) *PutRepositoriesYumProxyByRepositoryNameParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the put repositories yum proxy by repository name params
func (o *PutRepositoriesYumProxyByRepositoryNameParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *PutRepositoriesYumProxyByRepositoryNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
