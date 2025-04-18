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

// NewPostRepositoriesByRepositoryNameHealthCheckParams creates a new PostRepositoriesByRepositoryNameHealthCheckParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostRepositoriesByRepositoryNameHealthCheckParams() *PostRepositoriesByRepositoryNameHealthCheckParams {
	return &PostRepositoriesByRepositoryNameHealthCheckParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostRepositoriesByRepositoryNameHealthCheckParamsWithTimeout creates a new PostRepositoriesByRepositoryNameHealthCheckParams object
// with the ability to set a timeout on a request.
func NewPostRepositoriesByRepositoryNameHealthCheckParamsWithTimeout(timeout time.Duration) *PostRepositoriesByRepositoryNameHealthCheckParams {
	return &PostRepositoriesByRepositoryNameHealthCheckParams{
		timeout: timeout,
	}
}

// NewPostRepositoriesByRepositoryNameHealthCheckParamsWithContext creates a new PostRepositoriesByRepositoryNameHealthCheckParams object
// with the ability to set a context for a request.
func NewPostRepositoriesByRepositoryNameHealthCheckParamsWithContext(ctx context.Context) *PostRepositoriesByRepositoryNameHealthCheckParams {
	return &PostRepositoriesByRepositoryNameHealthCheckParams{
		Context: ctx,
	}
}

// NewPostRepositoriesByRepositoryNameHealthCheckParamsWithHTTPClient creates a new PostRepositoriesByRepositoryNameHealthCheckParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostRepositoriesByRepositoryNameHealthCheckParamsWithHTTPClient(client *http.Client) *PostRepositoriesByRepositoryNameHealthCheckParams {
	return &PostRepositoriesByRepositoryNameHealthCheckParams{
		HTTPClient: client,
	}
}

/*
PostRepositoriesByRepositoryNameHealthCheckParams contains all the parameters to send to the API endpoint

	for the post repositories by repository name health check operation.

	Typically these are written to a http.Request.
*/
type PostRepositoriesByRepositoryNameHealthCheckParams struct {

	/* RepositoryName.

	   Name of the repository to enable Repository Health Check for
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post repositories by repository name health check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRepositoriesByRepositoryNameHealthCheckParams) WithDefaults() *PostRepositoriesByRepositoryNameHealthCheckParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post repositories by repository name health check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRepositoriesByRepositoryNameHealthCheckParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post repositories by repository name health check params
func (o *PostRepositoriesByRepositoryNameHealthCheckParams) WithTimeout(timeout time.Duration) *PostRepositoriesByRepositoryNameHealthCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post repositories by repository name health check params
func (o *PostRepositoriesByRepositoryNameHealthCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post repositories by repository name health check params
func (o *PostRepositoriesByRepositoryNameHealthCheckParams) WithContext(ctx context.Context) *PostRepositoriesByRepositoryNameHealthCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post repositories by repository name health check params
func (o *PostRepositoriesByRepositoryNameHealthCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post repositories by repository name health check params
func (o *PostRepositoriesByRepositoryNameHealthCheckParams) WithHTTPClient(client *http.Client) *PostRepositoriesByRepositoryNameHealthCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post repositories by repository name health check params
func (o *PostRepositoriesByRepositoryNameHealthCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepositoryName adds the repositoryName to the post repositories by repository name health check params
func (o *PostRepositoriesByRepositoryNameHealthCheckParams) WithRepositoryName(repositoryName string) *PostRepositoriesByRepositoryNameHealthCheckParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the post repositories by repository name health check params
func (o *PostRepositoriesByRepositoryNameHealthCheckParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *PostRepositoriesByRepositoryNameHealthCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
