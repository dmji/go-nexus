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
)

// NewGetRepositoriesGoProxyByRepositoryNameParams creates a new GetRepositoriesGoProxyByRepositoryNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRepositoriesGoProxyByRepositoryNameParams() *GetRepositoriesGoProxyByRepositoryNameParams {
	return &GetRepositoriesGoProxyByRepositoryNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesGoProxyByRepositoryNameParamsWithTimeout creates a new GetRepositoriesGoProxyByRepositoryNameParams object
// with the ability to set a timeout on a request.
func NewGetRepositoriesGoProxyByRepositoryNameParamsWithTimeout(timeout time.Duration) *GetRepositoriesGoProxyByRepositoryNameParams {
	return &GetRepositoriesGoProxyByRepositoryNameParams{
		timeout: timeout,
	}
}

// NewGetRepositoriesGoProxyByRepositoryNameParamsWithContext creates a new GetRepositoriesGoProxyByRepositoryNameParams object
// with the ability to set a context for a request.
func NewGetRepositoriesGoProxyByRepositoryNameParamsWithContext(ctx context.Context) *GetRepositoriesGoProxyByRepositoryNameParams {
	return &GetRepositoriesGoProxyByRepositoryNameParams{
		Context: ctx,
	}
}

// NewGetRepositoriesGoProxyByRepositoryNameParamsWithHTTPClient creates a new GetRepositoriesGoProxyByRepositoryNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRepositoriesGoProxyByRepositoryNameParamsWithHTTPClient(client *http.Client) *GetRepositoriesGoProxyByRepositoryNameParams {
	return &GetRepositoriesGoProxyByRepositoryNameParams{
		HTTPClient: client,
	}
}

/*
GetRepositoriesGoProxyByRepositoryNameParams contains all the parameters to send to the API endpoint

	for the get repositories go proxy by repository name operation.

	Typically these are written to a http.Request.
*/
type GetRepositoriesGoProxyByRepositoryNameParams struct {

	// RepositoryName.
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get repositories go proxy by repository name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepositoriesGoProxyByRepositoryNameParams) WithDefaults() *GetRepositoriesGoProxyByRepositoryNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get repositories go proxy by repository name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepositoriesGoProxyByRepositoryNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get repositories go proxy by repository name params
func (o *GetRepositoriesGoProxyByRepositoryNameParams) WithTimeout(timeout time.Duration) *GetRepositoriesGoProxyByRepositoryNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories go proxy by repository name params
func (o *GetRepositoriesGoProxyByRepositoryNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories go proxy by repository name params
func (o *GetRepositoriesGoProxyByRepositoryNameParams) WithContext(ctx context.Context) *GetRepositoriesGoProxyByRepositoryNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories go proxy by repository name params
func (o *GetRepositoriesGoProxyByRepositoryNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories go proxy by repository name params
func (o *GetRepositoriesGoProxyByRepositoryNameParams) WithHTTPClient(client *http.Client) *GetRepositoriesGoProxyByRepositoryNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories go proxy by repository name params
func (o *GetRepositoriesGoProxyByRepositoryNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepositoryName adds the repositoryName to the get repositories go proxy by repository name params
func (o *GetRepositoriesGoProxyByRepositoryNameParams) WithRepositoryName(repositoryName string) *GetRepositoriesGoProxyByRepositoryNameParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the get repositories go proxy by repository name params
func (o *GetRepositoriesGoProxyByRepositoryNameParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesGoProxyByRepositoryNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param repositoryName
	if err := r.SetPathParam("repositoryName", o.RepositoryName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
