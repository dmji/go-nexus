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

// NewGetRepositoriesDockerHostedByRepositorynameParams creates a new GetRepositoriesDockerHostedByRepositorynameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRepositoriesDockerHostedByRepositorynameParams() *GetRepositoriesDockerHostedByRepositorynameParams {
	return &GetRepositoriesDockerHostedByRepositorynameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesDockerHostedByRepositorynameParamsWithTimeout creates a new GetRepositoriesDockerHostedByRepositorynameParams object
// with the ability to set a timeout on a request.
func NewGetRepositoriesDockerHostedByRepositorynameParamsWithTimeout(timeout time.Duration) *GetRepositoriesDockerHostedByRepositorynameParams {
	return &GetRepositoriesDockerHostedByRepositorynameParams{
		timeout: timeout,
	}
}

// NewGetRepositoriesDockerHostedByRepositorynameParamsWithContext creates a new GetRepositoriesDockerHostedByRepositorynameParams object
// with the ability to set a context for a request.
func NewGetRepositoriesDockerHostedByRepositorynameParamsWithContext(ctx context.Context) *GetRepositoriesDockerHostedByRepositorynameParams {
	return &GetRepositoriesDockerHostedByRepositorynameParams{
		Context: ctx,
	}
}

// NewGetRepositoriesDockerHostedByRepositorynameParamsWithHTTPClient creates a new GetRepositoriesDockerHostedByRepositorynameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRepositoriesDockerHostedByRepositorynameParamsWithHTTPClient(client *http.Client) *GetRepositoriesDockerHostedByRepositorynameParams {
	return &GetRepositoriesDockerHostedByRepositorynameParams{
		HTTPClient: client,
	}
}

/*
GetRepositoriesDockerHostedByRepositorynameParams contains all the parameters to send to the API endpoint

	for the get repositories docker hosted by repositoryname operation.

	Typically these are written to a http.Request.
*/
type GetRepositoriesDockerHostedByRepositorynameParams struct {

	// RepositoryName.
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get repositories docker hosted by repositoryname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepositoriesDockerHostedByRepositorynameParams) WithDefaults() *GetRepositoriesDockerHostedByRepositorynameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get repositories docker hosted by repositoryname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepositoriesDockerHostedByRepositorynameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get repositories docker hosted by repositoryname params
func (o *GetRepositoriesDockerHostedByRepositorynameParams) WithTimeout(timeout time.Duration) *GetRepositoriesDockerHostedByRepositorynameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories docker hosted by repositoryname params
func (o *GetRepositoriesDockerHostedByRepositorynameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories docker hosted by repositoryname params
func (o *GetRepositoriesDockerHostedByRepositorynameParams) WithContext(ctx context.Context) *GetRepositoriesDockerHostedByRepositorynameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories docker hosted by repositoryname params
func (o *GetRepositoriesDockerHostedByRepositorynameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories docker hosted by repositoryname params
func (o *GetRepositoriesDockerHostedByRepositorynameParams) WithHTTPClient(client *http.Client) *GetRepositoriesDockerHostedByRepositorynameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories docker hosted by repositoryname params
func (o *GetRepositoriesDockerHostedByRepositorynameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepositoryName adds the repositoryName to the get repositories docker hosted by repositoryname params
func (o *GetRepositoriesDockerHostedByRepositorynameParams) WithRepositoryName(repositoryName string) *GetRepositoriesDockerHostedByRepositorynameParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the get repositories docker hosted by repositoryname params
func (o *GetRepositoriesDockerHostedByRepositorynameParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesDockerHostedByRepositorynameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
