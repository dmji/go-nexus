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

// NewGetRepositoriesGoProxyByRepositorynameParams creates a new GetRepositoriesGoProxyByRepositorynameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRepositoriesGoProxyByRepositorynameParams() *GetRepositoriesGoProxyByRepositorynameParams {
	return &GetRepositoriesGoProxyByRepositorynameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesGoProxyByRepositorynameParamsWithTimeout creates a new GetRepositoriesGoProxyByRepositorynameParams object
// with the ability to set a timeout on a request.
func NewGetRepositoriesGoProxyByRepositorynameParamsWithTimeout(timeout time.Duration) *GetRepositoriesGoProxyByRepositorynameParams {
	return &GetRepositoriesGoProxyByRepositorynameParams{
		timeout: timeout,
	}
}

// NewGetRepositoriesGoProxyByRepositorynameParamsWithContext creates a new GetRepositoriesGoProxyByRepositorynameParams object
// with the ability to set a context for a request.
func NewGetRepositoriesGoProxyByRepositorynameParamsWithContext(ctx context.Context) *GetRepositoriesGoProxyByRepositorynameParams {
	return &GetRepositoriesGoProxyByRepositorynameParams{
		Context: ctx,
	}
}

// NewGetRepositoriesGoProxyByRepositorynameParamsWithHTTPClient creates a new GetRepositoriesGoProxyByRepositorynameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRepositoriesGoProxyByRepositorynameParamsWithHTTPClient(client *http.Client) *GetRepositoriesGoProxyByRepositorynameParams {
	return &GetRepositoriesGoProxyByRepositorynameParams{
		HTTPClient: client,
	}
}

/*
GetRepositoriesGoProxyByRepositorynameParams contains all the parameters to send to the API endpoint

	for the get repositories go proxy by repositoryname operation.

	Typically these are written to a http.Request.
*/
type GetRepositoriesGoProxyByRepositorynameParams struct {

	// RepositoryName.
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get repositories go proxy by repositoryname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepositoriesGoProxyByRepositorynameParams) WithDefaults() *GetRepositoriesGoProxyByRepositorynameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get repositories go proxy by repositoryname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepositoriesGoProxyByRepositorynameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get repositories go proxy by repositoryname params
func (o *GetRepositoriesGoProxyByRepositorynameParams) WithTimeout(timeout time.Duration) *GetRepositoriesGoProxyByRepositorynameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories go proxy by repositoryname params
func (o *GetRepositoriesGoProxyByRepositorynameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories go proxy by repositoryname params
func (o *GetRepositoriesGoProxyByRepositorynameParams) WithContext(ctx context.Context) *GetRepositoriesGoProxyByRepositorynameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories go proxy by repositoryname params
func (o *GetRepositoriesGoProxyByRepositorynameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories go proxy by repositoryname params
func (o *GetRepositoriesGoProxyByRepositorynameParams) WithHTTPClient(client *http.Client) *GetRepositoriesGoProxyByRepositorynameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories go proxy by repositoryname params
func (o *GetRepositoriesGoProxyByRepositorynameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepositoryName adds the repositoryName to the get repositories go proxy by repositoryname params
func (o *GetRepositoriesGoProxyByRepositorynameParams) WithRepositoryName(repositoryName string) *GetRepositoriesGoProxyByRepositorynameParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the get repositories go proxy by repositoryname params
func (o *GetRepositoriesGoProxyByRepositorynameParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesGoProxyByRepositorynameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
