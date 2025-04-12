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

// NewGetRepositoriesRGroupByRepositorynameParams creates a new GetRepositoriesRGroupByRepositorynameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRepositoriesRGroupByRepositorynameParams() *GetRepositoriesRGroupByRepositorynameParams {
	return &GetRepositoriesRGroupByRepositorynameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesRGroupByRepositorynameParamsWithTimeout creates a new GetRepositoriesRGroupByRepositorynameParams object
// with the ability to set a timeout on a request.
func NewGetRepositoriesRGroupByRepositorynameParamsWithTimeout(timeout time.Duration) *GetRepositoriesRGroupByRepositorynameParams {
	return &GetRepositoriesRGroupByRepositorynameParams{
		timeout: timeout,
	}
}

// NewGetRepositoriesRGroupByRepositorynameParamsWithContext creates a new GetRepositoriesRGroupByRepositorynameParams object
// with the ability to set a context for a request.
func NewGetRepositoriesRGroupByRepositorynameParamsWithContext(ctx context.Context) *GetRepositoriesRGroupByRepositorynameParams {
	return &GetRepositoriesRGroupByRepositorynameParams{
		Context: ctx,
	}
}

// NewGetRepositoriesRGroupByRepositorynameParamsWithHTTPClient creates a new GetRepositoriesRGroupByRepositorynameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRepositoriesRGroupByRepositorynameParamsWithHTTPClient(client *http.Client) *GetRepositoriesRGroupByRepositorynameParams {
	return &GetRepositoriesRGroupByRepositorynameParams{
		HTTPClient: client,
	}
}

/*
GetRepositoriesRGroupByRepositorynameParams contains all the parameters to send to the API endpoint

	for the get repositories r group by repositoryname operation.

	Typically these are written to a http.Request.
*/
type GetRepositoriesRGroupByRepositorynameParams struct {

	// RepositoryName.
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get repositories r group by repositoryname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepositoriesRGroupByRepositorynameParams) WithDefaults() *GetRepositoriesRGroupByRepositorynameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get repositories r group by repositoryname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepositoriesRGroupByRepositorynameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get repositories r group by repositoryname params
func (o *GetRepositoriesRGroupByRepositorynameParams) WithTimeout(timeout time.Duration) *GetRepositoriesRGroupByRepositorynameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories r group by repositoryname params
func (o *GetRepositoriesRGroupByRepositorynameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories r group by repositoryname params
func (o *GetRepositoriesRGroupByRepositorynameParams) WithContext(ctx context.Context) *GetRepositoriesRGroupByRepositorynameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories r group by repositoryname params
func (o *GetRepositoriesRGroupByRepositorynameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories r group by repositoryname params
func (o *GetRepositoriesRGroupByRepositorynameParams) WithHTTPClient(client *http.Client) *GetRepositoriesRGroupByRepositorynameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories r group by repositoryname params
func (o *GetRepositoriesRGroupByRepositorynameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepositoryName adds the repositoryName to the get repositories r group by repositoryname params
func (o *GetRepositoriesRGroupByRepositorynameParams) WithRepositoryName(repositoryName string) *GetRepositoriesRGroupByRepositorynameParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the get repositories r group by repositoryname params
func (o *GetRepositoriesRGroupByRepositorynameParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesRGroupByRepositorynameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
