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

// NewPutRepositoriesDockerHostedByRepositorynameParams creates a new PutRepositoriesDockerHostedByRepositorynameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutRepositoriesDockerHostedByRepositorynameParams() *PutRepositoriesDockerHostedByRepositorynameParams {
	return &PutRepositoriesDockerHostedByRepositorynameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutRepositoriesDockerHostedByRepositorynameParamsWithTimeout creates a new PutRepositoriesDockerHostedByRepositorynameParams object
// with the ability to set a timeout on a request.
func NewPutRepositoriesDockerHostedByRepositorynameParamsWithTimeout(timeout time.Duration) *PutRepositoriesDockerHostedByRepositorynameParams {
	return &PutRepositoriesDockerHostedByRepositorynameParams{
		timeout: timeout,
	}
}

// NewPutRepositoriesDockerHostedByRepositorynameParamsWithContext creates a new PutRepositoriesDockerHostedByRepositorynameParams object
// with the ability to set a context for a request.
func NewPutRepositoriesDockerHostedByRepositorynameParamsWithContext(ctx context.Context) *PutRepositoriesDockerHostedByRepositorynameParams {
	return &PutRepositoriesDockerHostedByRepositorynameParams{
		Context: ctx,
	}
}

// NewPutRepositoriesDockerHostedByRepositorynameParamsWithHTTPClient creates a new PutRepositoriesDockerHostedByRepositorynameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutRepositoriesDockerHostedByRepositorynameParamsWithHTTPClient(client *http.Client) *PutRepositoriesDockerHostedByRepositorynameParams {
	return &PutRepositoriesDockerHostedByRepositorynameParams{
		HTTPClient: client,
	}
}

/*
PutRepositoriesDockerHostedByRepositorynameParams contains all the parameters to send to the API endpoint

	for the put repositories docker hosted by repositoryname operation.

	Typically these are written to a http.Request.
*/
type PutRepositoriesDockerHostedByRepositorynameParams struct {

	// Body.
	Body *models.DockerHostedRepositoryAPIRequest

	/* RepositoryName.

	   Name of the repository to update
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put repositories docker hosted by repositoryname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRepositoriesDockerHostedByRepositorynameParams) WithDefaults() *PutRepositoriesDockerHostedByRepositorynameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put repositories docker hosted by repositoryname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRepositoriesDockerHostedByRepositorynameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put repositories docker hosted by repositoryname params
func (o *PutRepositoriesDockerHostedByRepositorynameParams) WithTimeout(timeout time.Duration) *PutRepositoriesDockerHostedByRepositorynameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put repositories docker hosted by repositoryname params
func (o *PutRepositoriesDockerHostedByRepositorynameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put repositories docker hosted by repositoryname params
func (o *PutRepositoriesDockerHostedByRepositorynameParams) WithContext(ctx context.Context) *PutRepositoriesDockerHostedByRepositorynameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put repositories docker hosted by repositoryname params
func (o *PutRepositoriesDockerHostedByRepositorynameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put repositories docker hosted by repositoryname params
func (o *PutRepositoriesDockerHostedByRepositorynameParams) WithHTTPClient(client *http.Client) *PutRepositoriesDockerHostedByRepositorynameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put repositories docker hosted by repositoryname params
func (o *PutRepositoriesDockerHostedByRepositorynameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put repositories docker hosted by repositoryname params
func (o *PutRepositoriesDockerHostedByRepositorynameParams) WithBody(body *models.DockerHostedRepositoryAPIRequest) *PutRepositoriesDockerHostedByRepositorynameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put repositories docker hosted by repositoryname params
func (o *PutRepositoriesDockerHostedByRepositorynameParams) SetBody(body *models.DockerHostedRepositoryAPIRequest) {
	o.Body = body
}

// WithRepositoryName adds the repositoryName to the put repositories docker hosted by repositoryname params
func (o *PutRepositoriesDockerHostedByRepositorynameParams) WithRepositoryName(repositoryName string) *PutRepositoriesDockerHostedByRepositorynameParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the put repositories docker hosted by repositoryname params
func (o *PutRepositoriesDockerHostedByRepositorynameParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *PutRepositoriesDockerHostedByRepositorynameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
