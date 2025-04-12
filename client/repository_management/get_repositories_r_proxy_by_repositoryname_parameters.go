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

// NewGetRepositoriesRProxyByRepositorynameParams creates a new GetRepositoriesRProxyByRepositorynameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRepositoriesRProxyByRepositorynameParams() *GetRepositoriesRProxyByRepositorynameParams {
	return &GetRepositoriesRProxyByRepositorynameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesRProxyByRepositorynameParamsWithTimeout creates a new GetRepositoriesRProxyByRepositorynameParams object
// with the ability to set a timeout on a request.
func NewGetRepositoriesRProxyByRepositorynameParamsWithTimeout(timeout time.Duration) *GetRepositoriesRProxyByRepositorynameParams {
	return &GetRepositoriesRProxyByRepositorynameParams{
		timeout: timeout,
	}
}

// NewGetRepositoriesRProxyByRepositorynameParamsWithContext creates a new GetRepositoriesRProxyByRepositorynameParams object
// with the ability to set a context for a request.
func NewGetRepositoriesRProxyByRepositorynameParamsWithContext(ctx context.Context) *GetRepositoriesRProxyByRepositorynameParams {
	return &GetRepositoriesRProxyByRepositorynameParams{
		Context: ctx,
	}
}

// NewGetRepositoriesRProxyByRepositorynameParamsWithHTTPClient creates a new GetRepositoriesRProxyByRepositorynameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRepositoriesRProxyByRepositorynameParamsWithHTTPClient(client *http.Client) *GetRepositoriesRProxyByRepositorynameParams {
	return &GetRepositoriesRProxyByRepositorynameParams{
		HTTPClient: client,
	}
}

/*
GetRepositoriesRProxyByRepositorynameParams contains all the parameters to send to the API endpoint

	for the get repositories r proxy by repositoryname operation.

	Typically these are written to a http.Request.
*/
type GetRepositoriesRProxyByRepositorynameParams struct {

	// RepositoryName.
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get repositories r proxy by repositoryname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepositoriesRProxyByRepositorynameParams) WithDefaults() *GetRepositoriesRProxyByRepositorynameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get repositories r proxy by repositoryname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepositoriesRProxyByRepositorynameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get repositories r proxy by repositoryname params
func (o *GetRepositoriesRProxyByRepositorynameParams) WithTimeout(timeout time.Duration) *GetRepositoriesRProxyByRepositorynameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories r proxy by repositoryname params
func (o *GetRepositoriesRProxyByRepositorynameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories r proxy by repositoryname params
func (o *GetRepositoriesRProxyByRepositorynameParams) WithContext(ctx context.Context) *GetRepositoriesRProxyByRepositorynameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories r proxy by repositoryname params
func (o *GetRepositoriesRProxyByRepositorynameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories r proxy by repositoryname params
func (o *GetRepositoriesRProxyByRepositorynameParams) WithHTTPClient(client *http.Client) *GetRepositoriesRProxyByRepositorynameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories r proxy by repositoryname params
func (o *GetRepositoriesRProxyByRepositorynameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepositoryName adds the repositoryName to the get repositories r proxy by repositoryname params
func (o *GetRepositoriesRProxyByRepositorynameParams) WithRepositoryName(repositoryName string) *GetRepositoriesRProxyByRepositorynameParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the get repositories r proxy by repositoryname params
func (o *GetRepositoriesRProxyByRepositorynameParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesRProxyByRepositorynameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
