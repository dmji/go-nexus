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

// NewGetRepositoriesRGroupByRepositoryNameParams creates a new GetRepositoriesRGroupByRepositoryNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRepositoriesRGroupByRepositoryNameParams() *GetRepositoriesRGroupByRepositoryNameParams {
	return &GetRepositoriesRGroupByRepositoryNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesRGroupByRepositoryNameParamsWithTimeout creates a new GetRepositoriesRGroupByRepositoryNameParams object
// with the ability to set a timeout on a request.
func NewGetRepositoriesRGroupByRepositoryNameParamsWithTimeout(timeout time.Duration) *GetRepositoriesRGroupByRepositoryNameParams {
	return &GetRepositoriesRGroupByRepositoryNameParams{
		timeout: timeout,
	}
}

// NewGetRepositoriesRGroupByRepositoryNameParamsWithContext creates a new GetRepositoriesRGroupByRepositoryNameParams object
// with the ability to set a context for a request.
func NewGetRepositoriesRGroupByRepositoryNameParamsWithContext(ctx context.Context) *GetRepositoriesRGroupByRepositoryNameParams {
	return &GetRepositoriesRGroupByRepositoryNameParams{
		Context: ctx,
	}
}

// NewGetRepositoriesRGroupByRepositoryNameParamsWithHTTPClient creates a new GetRepositoriesRGroupByRepositoryNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRepositoriesRGroupByRepositoryNameParamsWithHTTPClient(client *http.Client) *GetRepositoriesRGroupByRepositoryNameParams {
	return &GetRepositoriesRGroupByRepositoryNameParams{
		HTTPClient: client,
	}
}

/*
GetRepositoriesRGroupByRepositoryNameParams contains all the parameters to send to the API endpoint

	for the get repositories r group by repository name operation.

	Typically these are written to a http.Request.
*/
type GetRepositoriesRGroupByRepositoryNameParams struct {

	// RepositoryName.
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get repositories r group by repository name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepositoriesRGroupByRepositoryNameParams) WithDefaults() *GetRepositoriesRGroupByRepositoryNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get repositories r group by repository name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepositoriesRGroupByRepositoryNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get repositories r group by repository name params
func (o *GetRepositoriesRGroupByRepositoryNameParams) WithTimeout(timeout time.Duration) *GetRepositoriesRGroupByRepositoryNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories r group by repository name params
func (o *GetRepositoriesRGroupByRepositoryNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories r group by repository name params
func (o *GetRepositoriesRGroupByRepositoryNameParams) WithContext(ctx context.Context) *GetRepositoriesRGroupByRepositoryNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories r group by repository name params
func (o *GetRepositoriesRGroupByRepositoryNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories r group by repository name params
func (o *GetRepositoriesRGroupByRepositoryNameParams) WithHTTPClient(client *http.Client) *GetRepositoriesRGroupByRepositoryNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories r group by repository name params
func (o *GetRepositoriesRGroupByRepositoryNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepositoryName adds the repositoryName to the get repositories r group by repository name params
func (o *GetRepositoriesRGroupByRepositoryNameParams) WithRepositoryName(repositoryName string) *GetRepositoriesRGroupByRepositoryNameParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the get repositories r group by repository name params
func (o *GetRepositoriesRGroupByRepositoryNameParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesRGroupByRepositoryNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
