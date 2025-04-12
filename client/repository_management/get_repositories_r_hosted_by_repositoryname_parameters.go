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

// NewGetRepositoriesRHostedByRepositorynameParams creates a new GetRepositoriesRHostedByRepositorynameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRepositoriesRHostedByRepositorynameParams() *GetRepositoriesRHostedByRepositorynameParams {
	return &GetRepositoriesRHostedByRepositorynameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesRHostedByRepositorynameParamsWithTimeout creates a new GetRepositoriesRHostedByRepositorynameParams object
// with the ability to set a timeout on a request.
func NewGetRepositoriesRHostedByRepositorynameParamsWithTimeout(timeout time.Duration) *GetRepositoriesRHostedByRepositorynameParams {
	return &GetRepositoriesRHostedByRepositorynameParams{
		timeout: timeout,
	}
}

// NewGetRepositoriesRHostedByRepositorynameParamsWithContext creates a new GetRepositoriesRHostedByRepositorynameParams object
// with the ability to set a context for a request.
func NewGetRepositoriesRHostedByRepositorynameParamsWithContext(ctx context.Context) *GetRepositoriesRHostedByRepositorynameParams {
	return &GetRepositoriesRHostedByRepositorynameParams{
		Context: ctx,
	}
}

// NewGetRepositoriesRHostedByRepositorynameParamsWithHTTPClient creates a new GetRepositoriesRHostedByRepositorynameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRepositoriesRHostedByRepositorynameParamsWithHTTPClient(client *http.Client) *GetRepositoriesRHostedByRepositorynameParams {
	return &GetRepositoriesRHostedByRepositorynameParams{
		HTTPClient: client,
	}
}

/*
GetRepositoriesRHostedByRepositorynameParams contains all the parameters to send to the API endpoint

	for the get repositories r hosted by repositoryname operation.

	Typically these are written to a http.Request.
*/
type GetRepositoriesRHostedByRepositorynameParams struct {

	// RepositoryName.
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get repositories r hosted by repositoryname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepositoriesRHostedByRepositorynameParams) WithDefaults() *GetRepositoriesRHostedByRepositorynameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get repositories r hosted by repositoryname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepositoriesRHostedByRepositorynameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get repositories r hosted by repositoryname params
func (o *GetRepositoriesRHostedByRepositorynameParams) WithTimeout(timeout time.Duration) *GetRepositoriesRHostedByRepositorynameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories r hosted by repositoryname params
func (o *GetRepositoriesRHostedByRepositorynameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories r hosted by repositoryname params
func (o *GetRepositoriesRHostedByRepositorynameParams) WithContext(ctx context.Context) *GetRepositoriesRHostedByRepositorynameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories r hosted by repositoryname params
func (o *GetRepositoriesRHostedByRepositorynameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories r hosted by repositoryname params
func (o *GetRepositoriesRHostedByRepositorynameParams) WithHTTPClient(client *http.Client) *GetRepositoriesRHostedByRepositorynameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories r hosted by repositoryname params
func (o *GetRepositoriesRHostedByRepositorynameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepositoryName adds the repositoryName to the get repositories r hosted by repositoryname params
func (o *GetRepositoriesRHostedByRepositorynameParams) WithRepositoryName(repositoryName string) *GetRepositoriesRHostedByRepositorynameParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the get repositories r hosted by repositoryname params
func (o *GetRepositoriesRHostedByRepositorynameParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesRHostedByRepositorynameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
