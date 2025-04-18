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

// NewPutRepositoriesRubygemsGroupByRepositoryNameParams creates a new PutRepositoriesRubygemsGroupByRepositoryNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutRepositoriesRubygemsGroupByRepositoryNameParams() *PutRepositoriesRubygemsGroupByRepositoryNameParams {
	return &PutRepositoriesRubygemsGroupByRepositoryNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutRepositoriesRubygemsGroupByRepositoryNameParamsWithTimeout creates a new PutRepositoriesRubygemsGroupByRepositoryNameParams object
// with the ability to set a timeout on a request.
func NewPutRepositoriesRubygemsGroupByRepositoryNameParamsWithTimeout(timeout time.Duration) *PutRepositoriesRubygemsGroupByRepositoryNameParams {
	return &PutRepositoriesRubygemsGroupByRepositoryNameParams{
		timeout: timeout,
	}
}

// NewPutRepositoriesRubygemsGroupByRepositoryNameParamsWithContext creates a new PutRepositoriesRubygemsGroupByRepositoryNameParams object
// with the ability to set a context for a request.
func NewPutRepositoriesRubygemsGroupByRepositoryNameParamsWithContext(ctx context.Context) *PutRepositoriesRubygemsGroupByRepositoryNameParams {
	return &PutRepositoriesRubygemsGroupByRepositoryNameParams{
		Context: ctx,
	}
}

// NewPutRepositoriesRubygemsGroupByRepositoryNameParamsWithHTTPClient creates a new PutRepositoriesRubygemsGroupByRepositoryNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutRepositoriesRubygemsGroupByRepositoryNameParamsWithHTTPClient(client *http.Client) *PutRepositoriesRubygemsGroupByRepositoryNameParams {
	return &PutRepositoriesRubygemsGroupByRepositoryNameParams{
		HTTPClient: client,
	}
}

/*
PutRepositoriesRubygemsGroupByRepositoryNameParams contains all the parameters to send to the API endpoint

	for the put repositories rubygems group by repository name operation.

	Typically these are written to a http.Request.
*/
type PutRepositoriesRubygemsGroupByRepositoryNameParams struct {

	// Body.
	Body *models.RubyGemsGroupRepositoryAPIRequest

	/* RepositoryName.

	   Name of the repository to update
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put repositories rubygems group by repository name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRepositoriesRubygemsGroupByRepositoryNameParams) WithDefaults() *PutRepositoriesRubygemsGroupByRepositoryNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put repositories rubygems group by repository name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRepositoriesRubygemsGroupByRepositoryNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put repositories rubygems group by repository name params
func (o *PutRepositoriesRubygemsGroupByRepositoryNameParams) WithTimeout(timeout time.Duration) *PutRepositoriesRubygemsGroupByRepositoryNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put repositories rubygems group by repository name params
func (o *PutRepositoriesRubygemsGroupByRepositoryNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put repositories rubygems group by repository name params
func (o *PutRepositoriesRubygemsGroupByRepositoryNameParams) WithContext(ctx context.Context) *PutRepositoriesRubygemsGroupByRepositoryNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put repositories rubygems group by repository name params
func (o *PutRepositoriesRubygemsGroupByRepositoryNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put repositories rubygems group by repository name params
func (o *PutRepositoriesRubygemsGroupByRepositoryNameParams) WithHTTPClient(client *http.Client) *PutRepositoriesRubygemsGroupByRepositoryNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put repositories rubygems group by repository name params
func (o *PutRepositoriesRubygemsGroupByRepositoryNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put repositories rubygems group by repository name params
func (o *PutRepositoriesRubygemsGroupByRepositoryNameParams) WithBody(body *models.RubyGemsGroupRepositoryAPIRequest) *PutRepositoriesRubygemsGroupByRepositoryNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put repositories rubygems group by repository name params
func (o *PutRepositoriesRubygemsGroupByRepositoryNameParams) SetBody(body *models.RubyGemsGroupRepositoryAPIRequest) {
	o.Body = body
}

// WithRepositoryName adds the repositoryName to the put repositories rubygems group by repository name params
func (o *PutRepositoriesRubygemsGroupByRepositoryNameParams) WithRepositoryName(repositoryName string) *PutRepositoriesRubygemsGroupByRepositoryNameParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the put repositories rubygems group by repository name params
func (o *PutRepositoriesRubygemsGroupByRepositoryNameParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *PutRepositoriesRubygemsGroupByRepositoryNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
