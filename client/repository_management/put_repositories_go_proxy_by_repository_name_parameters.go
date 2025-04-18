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

// NewPutRepositoriesGoProxyByRepositoryNameParams creates a new PutRepositoriesGoProxyByRepositoryNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutRepositoriesGoProxyByRepositoryNameParams() *PutRepositoriesGoProxyByRepositoryNameParams {
	return &PutRepositoriesGoProxyByRepositoryNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutRepositoriesGoProxyByRepositoryNameParamsWithTimeout creates a new PutRepositoriesGoProxyByRepositoryNameParams object
// with the ability to set a timeout on a request.
func NewPutRepositoriesGoProxyByRepositoryNameParamsWithTimeout(timeout time.Duration) *PutRepositoriesGoProxyByRepositoryNameParams {
	return &PutRepositoriesGoProxyByRepositoryNameParams{
		timeout: timeout,
	}
}

// NewPutRepositoriesGoProxyByRepositoryNameParamsWithContext creates a new PutRepositoriesGoProxyByRepositoryNameParams object
// with the ability to set a context for a request.
func NewPutRepositoriesGoProxyByRepositoryNameParamsWithContext(ctx context.Context) *PutRepositoriesGoProxyByRepositoryNameParams {
	return &PutRepositoriesGoProxyByRepositoryNameParams{
		Context: ctx,
	}
}

// NewPutRepositoriesGoProxyByRepositoryNameParamsWithHTTPClient creates a new PutRepositoriesGoProxyByRepositoryNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutRepositoriesGoProxyByRepositoryNameParamsWithHTTPClient(client *http.Client) *PutRepositoriesGoProxyByRepositoryNameParams {
	return &PutRepositoriesGoProxyByRepositoryNameParams{
		HTTPClient: client,
	}
}

/*
PutRepositoriesGoProxyByRepositoryNameParams contains all the parameters to send to the API endpoint

	for the put repositories go proxy by repository name operation.

	Typically these are written to a http.Request.
*/
type PutRepositoriesGoProxyByRepositoryNameParams struct {

	// Body.
	Body *models.GolangProxyRepositoryAPIRequest

	/* RepositoryName.

	   Name of the repository to update
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put repositories go proxy by repository name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRepositoriesGoProxyByRepositoryNameParams) WithDefaults() *PutRepositoriesGoProxyByRepositoryNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put repositories go proxy by repository name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRepositoriesGoProxyByRepositoryNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put repositories go proxy by repository name params
func (o *PutRepositoriesGoProxyByRepositoryNameParams) WithTimeout(timeout time.Duration) *PutRepositoriesGoProxyByRepositoryNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put repositories go proxy by repository name params
func (o *PutRepositoriesGoProxyByRepositoryNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put repositories go proxy by repository name params
func (o *PutRepositoriesGoProxyByRepositoryNameParams) WithContext(ctx context.Context) *PutRepositoriesGoProxyByRepositoryNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put repositories go proxy by repository name params
func (o *PutRepositoriesGoProxyByRepositoryNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put repositories go proxy by repository name params
func (o *PutRepositoriesGoProxyByRepositoryNameParams) WithHTTPClient(client *http.Client) *PutRepositoriesGoProxyByRepositoryNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put repositories go proxy by repository name params
func (o *PutRepositoriesGoProxyByRepositoryNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put repositories go proxy by repository name params
func (o *PutRepositoriesGoProxyByRepositoryNameParams) WithBody(body *models.GolangProxyRepositoryAPIRequest) *PutRepositoriesGoProxyByRepositoryNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put repositories go proxy by repository name params
func (o *PutRepositoriesGoProxyByRepositoryNameParams) SetBody(body *models.GolangProxyRepositoryAPIRequest) {
	o.Body = body
}

// WithRepositoryName adds the repositoryName to the put repositories go proxy by repository name params
func (o *PutRepositoriesGoProxyByRepositoryNameParams) WithRepositoryName(repositoryName string) *PutRepositoriesGoProxyByRepositoryNameParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the put repositories go proxy by repository name params
func (o *PutRepositoriesGoProxyByRepositoryNameParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *PutRepositoriesGoProxyByRepositoryNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
