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

// NewGetRepositoriesDockerGroupByRepositorynameParams creates a new GetRepositoriesDockerGroupByRepositorynameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRepositoriesDockerGroupByRepositorynameParams() *GetRepositoriesDockerGroupByRepositorynameParams {
	return &GetRepositoriesDockerGroupByRepositorynameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesDockerGroupByRepositorynameParamsWithTimeout creates a new GetRepositoriesDockerGroupByRepositorynameParams object
// with the ability to set a timeout on a request.
func NewGetRepositoriesDockerGroupByRepositorynameParamsWithTimeout(timeout time.Duration) *GetRepositoriesDockerGroupByRepositorynameParams {
	return &GetRepositoriesDockerGroupByRepositorynameParams{
		timeout: timeout,
	}
}

// NewGetRepositoriesDockerGroupByRepositorynameParamsWithContext creates a new GetRepositoriesDockerGroupByRepositorynameParams object
// with the ability to set a context for a request.
func NewGetRepositoriesDockerGroupByRepositorynameParamsWithContext(ctx context.Context) *GetRepositoriesDockerGroupByRepositorynameParams {
	return &GetRepositoriesDockerGroupByRepositorynameParams{
		Context: ctx,
	}
}

// NewGetRepositoriesDockerGroupByRepositorynameParamsWithHTTPClient creates a new GetRepositoriesDockerGroupByRepositorynameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRepositoriesDockerGroupByRepositorynameParamsWithHTTPClient(client *http.Client) *GetRepositoriesDockerGroupByRepositorynameParams {
	return &GetRepositoriesDockerGroupByRepositorynameParams{
		HTTPClient: client,
	}
}

/*
GetRepositoriesDockerGroupByRepositorynameParams contains all the parameters to send to the API endpoint

	for the get repositories docker group by repositoryname operation.

	Typically these are written to a http.Request.
*/
type GetRepositoriesDockerGroupByRepositorynameParams struct {

	// RepositoryName.
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get repositories docker group by repositoryname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepositoriesDockerGroupByRepositorynameParams) WithDefaults() *GetRepositoriesDockerGroupByRepositorynameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get repositories docker group by repositoryname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepositoriesDockerGroupByRepositorynameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get repositories docker group by repositoryname params
func (o *GetRepositoriesDockerGroupByRepositorynameParams) WithTimeout(timeout time.Duration) *GetRepositoriesDockerGroupByRepositorynameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories docker group by repositoryname params
func (o *GetRepositoriesDockerGroupByRepositorynameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories docker group by repositoryname params
func (o *GetRepositoriesDockerGroupByRepositorynameParams) WithContext(ctx context.Context) *GetRepositoriesDockerGroupByRepositorynameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories docker group by repositoryname params
func (o *GetRepositoriesDockerGroupByRepositorynameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories docker group by repositoryname params
func (o *GetRepositoriesDockerGroupByRepositorynameParams) WithHTTPClient(client *http.Client) *GetRepositoriesDockerGroupByRepositorynameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories docker group by repositoryname params
func (o *GetRepositoriesDockerGroupByRepositorynameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepositoryName adds the repositoryName to the get repositories docker group by repositoryname params
func (o *GetRepositoriesDockerGroupByRepositorynameParams) WithRepositoryName(repositoryName string) *GetRepositoriesDockerGroupByRepositorynameParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the get repositories docker group by repositoryname params
func (o *GetRepositoriesDockerGroupByRepositorynameParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesDockerGroupByRepositorynameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
