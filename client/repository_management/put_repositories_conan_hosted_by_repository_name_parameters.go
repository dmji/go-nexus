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

// NewPutRepositoriesConanHostedByRepositoryNameParams creates a new PutRepositoriesConanHostedByRepositoryNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutRepositoriesConanHostedByRepositoryNameParams() *PutRepositoriesConanHostedByRepositoryNameParams {
	return &PutRepositoriesConanHostedByRepositoryNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutRepositoriesConanHostedByRepositoryNameParamsWithTimeout creates a new PutRepositoriesConanHostedByRepositoryNameParams object
// with the ability to set a timeout on a request.
func NewPutRepositoriesConanHostedByRepositoryNameParamsWithTimeout(timeout time.Duration) *PutRepositoriesConanHostedByRepositoryNameParams {
	return &PutRepositoriesConanHostedByRepositoryNameParams{
		timeout: timeout,
	}
}

// NewPutRepositoriesConanHostedByRepositoryNameParamsWithContext creates a new PutRepositoriesConanHostedByRepositoryNameParams object
// with the ability to set a context for a request.
func NewPutRepositoriesConanHostedByRepositoryNameParamsWithContext(ctx context.Context) *PutRepositoriesConanHostedByRepositoryNameParams {
	return &PutRepositoriesConanHostedByRepositoryNameParams{
		Context: ctx,
	}
}

// NewPutRepositoriesConanHostedByRepositoryNameParamsWithHTTPClient creates a new PutRepositoriesConanHostedByRepositoryNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutRepositoriesConanHostedByRepositoryNameParamsWithHTTPClient(client *http.Client) *PutRepositoriesConanHostedByRepositoryNameParams {
	return &PutRepositoriesConanHostedByRepositoryNameParams{
		HTTPClient: client,
	}
}

/*
PutRepositoriesConanHostedByRepositoryNameParams contains all the parameters to send to the API endpoint

	for the put repositories conan hosted by repository name operation.

	Typically these are written to a http.Request.
*/
type PutRepositoriesConanHostedByRepositoryNameParams struct {

	// Body.
	Body *models.ConanHostedRepositoryAPIRequest

	/* RepositoryName.

	   Name of the repository to update
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put repositories conan hosted by repository name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRepositoriesConanHostedByRepositoryNameParams) WithDefaults() *PutRepositoriesConanHostedByRepositoryNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put repositories conan hosted by repository name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRepositoriesConanHostedByRepositoryNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put repositories conan hosted by repository name params
func (o *PutRepositoriesConanHostedByRepositoryNameParams) WithTimeout(timeout time.Duration) *PutRepositoriesConanHostedByRepositoryNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put repositories conan hosted by repository name params
func (o *PutRepositoriesConanHostedByRepositoryNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put repositories conan hosted by repository name params
func (o *PutRepositoriesConanHostedByRepositoryNameParams) WithContext(ctx context.Context) *PutRepositoriesConanHostedByRepositoryNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put repositories conan hosted by repository name params
func (o *PutRepositoriesConanHostedByRepositoryNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put repositories conan hosted by repository name params
func (o *PutRepositoriesConanHostedByRepositoryNameParams) WithHTTPClient(client *http.Client) *PutRepositoriesConanHostedByRepositoryNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put repositories conan hosted by repository name params
func (o *PutRepositoriesConanHostedByRepositoryNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put repositories conan hosted by repository name params
func (o *PutRepositoriesConanHostedByRepositoryNameParams) WithBody(body *models.ConanHostedRepositoryAPIRequest) *PutRepositoriesConanHostedByRepositoryNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put repositories conan hosted by repository name params
func (o *PutRepositoriesConanHostedByRepositoryNameParams) SetBody(body *models.ConanHostedRepositoryAPIRequest) {
	o.Body = body
}

// WithRepositoryName adds the repositoryName to the put repositories conan hosted by repository name params
func (o *PutRepositoriesConanHostedByRepositoryNameParams) WithRepositoryName(repositoryName string) *PutRepositoriesConanHostedByRepositoryNameParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the put repositories conan hosted by repository name params
func (o *PutRepositoriesConanHostedByRepositoryNameParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *PutRepositoriesConanHostedByRepositoryNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
