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

// NewPutRepositoriesMavenGroupByRepositorynameParams creates a new PutRepositoriesMavenGroupByRepositorynameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutRepositoriesMavenGroupByRepositorynameParams() *PutRepositoriesMavenGroupByRepositorynameParams {
	return &PutRepositoriesMavenGroupByRepositorynameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutRepositoriesMavenGroupByRepositorynameParamsWithTimeout creates a new PutRepositoriesMavenGroupByRepositorynameParams object
// with the ability to set a timeout on a request.
func NewPutRepositoriesMavenGroupByRepositorynameParamsWithTimeout(timeout time.Duration) *PutRepositoriesMavenGroupByRepositorynameParams {
	return &PutRepositoriesMavenGroupByRepositorynameParams{
		timeout: timeout,
	}
}

// NewPutRepositoriesMavenGroupByRepositorynameParamsWithContext creates a new PutRepositoriesMavenGroupByRepositorynameParams object
// with the ability to set a context for a request.
func NewPutRepositoriesMavenGroupByRepositorynameParamsWithContext(ctx context.Context) *PutRepositoriesMavenGroupByRepositorynameParams {
	return &PutRepositoriesMavenGroupByRepositorynameParams{
		Context: ctx,
	}
}

// NewPutRepositoriesMavenGroupByRepositorynameParamsWithHTTPClient creates a new PutRepositoriesMavenGroupByRepositorynameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutRepositoriesMavenGroupByRepositorynameParamsWithHTTPClient(client *http.Client) *PutRepositoriesMavenGroupByRepositorynameParams {
	return &PutRepositoriesMavenGroupByRepositorynameParams{
		HTTPClient: client,
	}
}

/*
PutRepositoriesMavenGroupByRepositorynameParams contains all the parameters to send to the API endpoint

	for the put repositories maven group by repositoryname operation.

	Typically these are written to a http.Request.
*/
type PutRepositoriesMavenGroupByRepositorynameParams struct {

	// Body.
	Body *models.MavenGroupRepositoryAPIRequest

	/* RepositoryName.

	   Name of the repository to update
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put repositories maven group by repositoryname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRepositoriesMavenGroupByRepositorynameParams) WithDefaults() *PutRepositoriesMavenGroupByRepositorynameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put repositories maven group by repositoryname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRepositoriesMavenGroupByRepositorynameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put repositories maven group by repositoryname params
func (o *PutRepositoriesMavenGroupByRepositorynameParams) WithTimeout(timeout time.Duration) *PutRepositoriesMavenGroupByRepositorynameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put repositories maven group by repositoryname params
func (o *PutRepositoriesMavenGroupByRepositorynameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put repositories maven group by repositoryname params
func (o *PutRepositoriesMavenGroupByRepositorynameParams) WithContext(ctx context.Context) *PutRepositoriesMavenGroupByRepositorynameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put repositories maven group by repositoryname params
func (o *PutRepositoriesMavenGroupByRepositorynameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put repositories maven group by repositoryname params
func (o *PutRepositoriesMavenGroupByRepositorynameParams) WithHTTPClient(client *http.Client) *PutRepositoriesMavenGroupByRepositorynameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put repositories maven group by repositoryname params
func (o *PutRepositoriesMavenGroupByRepositorynameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put repositories maven group by repositoryname params
func (o *PutRepositoriesMavenGroupByRepositorynameParams) WithBody(body *models.MavenGroupRepositoryAPIRequest) *PutRepositoriesMavenGroupByRepositorynameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put repositories maven group by repositoryname params
func (o *PutRepositoriesMavenGroupByRepositorynameParams) SetBody(body *models.MavenGroupRepositoryAPIRequest) {
	o.Body = body
}

// WithRepositoryName adds the repositoryName to the put repositories maven group by repositoryname params
func (o *PutRepositoriesMavenGroupByRepositorynameParams) WithRepositoryName(repositoryName string) *PutRepositoriesMavenGroupByRepositorynameParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the put repositories maven group by repositoryname params
func (o *PutRepositoriesMavenGroupByRepositorynameParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *PutRepositoriesMavenGroupByRepositorynameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
