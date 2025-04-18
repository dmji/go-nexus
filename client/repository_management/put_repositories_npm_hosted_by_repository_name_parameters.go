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

// NewPutRepositoriesNpmHostedByRepositoryNameParams creates a new PutRepositoriesNpmHostedByRepositoryNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutRepositoriesNpmHostedByRepositoryNameParams() *PutRepositoriesNpmHostedByRepositoryNameParams {
	return &PutRepositoriesNpmHostedByRepositoryNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutRepositoriesNpmHostedByRepositoryNameParamsWithTimeout creates a new PutRepositoriesNpmHostedByRepositoryNameParams object
// with the ability to set a timeout on a request.
func NewPutRepositoriesNpmHostedByRepositoryNameParamsWithTimeout(timeout time.Duration) *PutRepositoriesNpmHostedByRepositoryNameParams {
	return &PutRepositoriesNpmHostedByRepositoryNameParams{
		timeout: timeout,
	}
}

// NewPutRepositoriesNpmHostedByRepositoryNameParamsWithContext creates a new PutRepositoriesNpmHostedByRepositoryNameParams object
// with the ability to set a context for a request.
func NewPutRepositoriesNpmHostedByRepositoryNameParamsWithContext(ctx context.Context) *PutRepositoriesNpmHostedByRepositoryNameParams {
	return &PutRepositoriesNpmHostedByRepositoryNameParams{
		Context: ctx,
	}
}

// NewPutRepositoriesNpmHostedByRepositoryNameParamsWithHTTPClient creates a new PutRepositoriesNpmHostedByRepositoryNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutRepositoriesNpmHostedByRepositoryNameParamsWithHTTPClient(client *http.Client) *PutRepositoriesNpmHostedByRepositoryNameParams {
	return &PutRepositoriesNpmHostedByRepositoryNameParams{
		HTTPClient: client,
	}
}

/*
PutRepositoriesNpmHostedByRepositoryNameParams contains all the parameters to send to the API endpoint

	for the put repositories npm hosted by repository name operation.

	Typically these are written to a http.Request.
*/
type PutRepositoriesNpmHostedByRepositoryNameParams struct {

	// Body.
	Body *models.NpmHostedRepositoryAPIRequest

	/* RepositoryName.

	   Name of the repository to update
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put repositories npm hosted by repository name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRepositoriesNpmHostedByRepositoryNameParams) WithDefaults() *PutRepositoriesNpmHostedByRepositoryNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put repositories npm hosted by repository name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRepositoriesNpmHostedByRepositoryNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put repositories npm hosted by repository name params
func (o *PutRepositoriesNpmHostedByRepositoryNameParams) WithTimeout(timeout time.Duration) *PutRepositoriesNpmHostedByRepositoryNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put repositories npm hosted by repository name params
func (o *PutRepositoriesNpmHostedByRepositoryNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put repositories npm hosted by repository name params
func (o *PutRepositoriesNpmHostedByRepositoryNameParams) WithContext(ctx context.Context) *PutRepositoriesNpmHostedByRepositoryNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put repositories npm hosted by repository name params
func (o *PutRepositoriesNpmHostedByRepositoryNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put repositories npm hosted by repository name params
func (o *PutRepositoriesNpmHostedByRepositoryNameParams) WithHTTPClient(client *http.Client) *PutRepositoriesNpmHostedByRepositoryNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put repositories npm hosted by repository name params
func (o *PutRepositoriesNpmHostedByRepositoryNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put repositories npm hosted by repository name params
func (o *PutRepositoriesNpmHostedByRepositoryNameParams) WithBody(body *models.NpmHostedRepositoryAPIRequest) *PutRepositoriesNpmHostedByRepositoryNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put repositories npm hosted by repository name params
func (o *PutRepositoriesNpmHostedByRepositoryNameParams) SetBody(body *models.NpmHostedRepositoryAPIRequest) {
	o.Body = body
}

// WithRepositoryName adds the repositoryName to the put repositories npm hosted by repository name params
func (o *PutRepositoriesNpmHostedByRepositoryNameParams) WithRepositoryName(repositoryName string) *PutRepositoriesNpmHostedByRepositoryNameParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the put repositories npm hosted by repository name params
func (o *PutRepositoriesNpmHostedByRepositoryNameParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *PutRepositoriesNpmHostedByRepositoryNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
