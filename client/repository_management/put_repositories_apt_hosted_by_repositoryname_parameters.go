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

// NewPutRepositoriesAptHostedByRepositorynameParams creates a new PutRepositoriesAptHostedByRepositorynameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutRepositoriesAptHostedByRepositorynameParams() *PutRepositoriesAptHostedByRepositorynameParams {
	return &PutRepositoriesAptHostedByRepositorynameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutRepositoriesAptHostedByRepositorynameParamsWithTimeout creates a new PutRepositoriesAptHostedByRepositorynameParams object
// with the ability to set a timeout on a request.
func NewPutRepositoriesAptHostedByRepositorynameParamsWithTimeout(timeout time.Duration) *PutRepositoriesAptHostedByRepositorynameParams {
	return &PutRepositoriesAptHostedByRepositorynameParams{
		timeout: timeout,
	}
}

// NewPutRepositoriesAptHostedByRepositorynameParamsWithContext creates a new PutRepositoriesAptHostedByRepositorynameParams object
// with the ability to set a context for a request.
func NewPutRepositoriesAptHostedByRepositorynameParamsWithContext(ctx context.Context) *PutRepositoriesAptHostedByRepositorynameParams {
	return &PutRepositoriesAptHostedByRepositorynameParams{
		Context: ctx,
	}
}

// NewPutRepositoriesAptHostedByRepositorynameParamsWithHTTPClient creates a new PutRepositoriesAptHostedByRepositorynameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutRepositoriesAptHostedByRepositorynameParamsWithHTTPClient(client *http.Client) *PutRepositoriesAptHostedByRepositorynameParams {
	return &PutRepositoriesAptHostedByRepositorynameParams{
		HTTPClient: client,
	}
}

/*
PutRepositoriesAptHostedByRepositorynameParams contains all the parameters to send to the API endpoint

	for the put repositories apt hosted by repositoryname operation.

	Typically these are written to a http.Request.
*/
type PutRepositoriesAptHostedByRepositorynameParams struct {

	// Body.
	Body *models.AptHostedRepositoryAPIRequest

	/* RepositoryName.

	   Name of the repository to update
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put repositories apt hosted by repositoryname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRepositoriesAptHostedByRepositorynameParams) WithDefaults() *PutRepositoriesAptHostedByRepositorynameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put repositories apt hosted by repositoryname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRepositoriesAptHostedByRepositorynameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put repositories apt hosted by repositoryname params
func (o *PutRepositoriesAptHostedByRepositorynameParams) WithTimeout(timeout time.Duration) *PutRepositoriesAptHostedByRepositorynameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put repositories apt hosted by repositoryname params
func (o *PutRepositoriesAptHostedByRepositorynameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put repositories apt hosted by repositoryname params
func (o *PutRepositoriesAptHostedByRepositorynameParams) WithContext(ctx context.Context) *PutRepositoriesAptHostedByRepositorynameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put repositories apt hosted by repositoryname params
func (o *PutRepositoriesAptHostedByRepositorynameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put repositories apt hosted by repositoryname params
func (o *PutRepositoriesAptHostedByRepositorynameParams) WithHTTPClient(client *http.Client) *PutRepositoriesAptHostedByRepositorynameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put repositories apt hosted by repositoryname params
func (o *PutRepositoriesAptHostedByRepositorynameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put repositories apt hosted by repositoryname params
func (o *PutRepositoriesAptHostedByRepositorynameParams) WithBody(body *models.AptHostedRepositoryAPIRequest) *PutRepositoriesAptHostedByRepositorynameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put repositories apt hosted by repositoryname params
func (o *PutRepositoriesAptHostedByRepositorynameParams) SetBody(body *models.AptHostedRepositoryAPIRequest) {
	o.Body = body
}

// WithRepositoryName adds the repositoryName to the put repositories apt hosted by repositoryname params
func (o *PutRepositoriesAptHostedByRepositorynameParams) WithRepositoryName(repositoryName string) *PutRepositoriesAptHostedByRepositorynameParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the put repositories apt hosted by repositoryname params
func (o *PutRepositoriesAptHostedByRepositorynameParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *PutRepositoriesAptHostedByRepositorynameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
