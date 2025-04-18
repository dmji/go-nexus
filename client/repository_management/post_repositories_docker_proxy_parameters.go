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

// NewPostRepositoriesDockerProxyParams creates a new PostRepositoriesDockerProxyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostRepositoriesDockerProxyParams() *PostRepositoriesDockerProxyParams {
	return &PostRepositoriesDockerProxyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostRepositoriesDockerProxyParamsWithTimeout creates a new PostRepositoriesDockerProxyParams object
// with the ability to set a timeout on a request.
func NewPostRepositoriesDockerProxyParamsWithTimeout(timeout time.Duration) *PostRepositoriesDockerProxyParams {
	return &PostRepositoriesDockerProxyParams{
		timeout: timeout,
	}
}

// NewPostRepositoriesDockerProxyParamsWithContext creates a new PostRepositoriesDockerProxyParams object
// with the ability to set a context for a request.
func NewPostRepositoriesDockerProxyParamsWithContext(ctx context.Context) *PostRepositoriesDockerProxyParams {
	return &PostRepositoriesDockerProxyParams{
		Context: ctx,
	}
}

// NewPostRepositoriesDockerProxyParamsWithHTTPClient creates a new PostRepositoriesDockerProxyParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostRepositoriesDockerProxyParamsWithHTTPClient(client *http.Client) *PostRepositoriesDockerProxyParams {
	return &PostRepositoriesDockerProxyParams{
		HTTPClient: client,
	}
}

/*
PostRepositoriesDockerProxyParams contains all the parameters to send to the API endpoint

	for the post repositories docker proxy operation.

	Typically these are written to a http.Request.
*/
type PostRepositoriesDockerProxyParams struct {

	// Body.
	Body *models.DockerProxyRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post repositories docker proxy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRepositoriesDockerProxyParams) WithDefaults() *PostRepositoriesDockerProxyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post repositories docker proxy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRepositoriesDockerProxyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post repositories docker proxy params
func (o *PostRepositoriesDockerProxyParams) WithTimeout(timeout time.Duration) *PostRepositoriesDockerProxyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post repositories docker proxy params
func (o *PostRepositoriesDockerProxyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post repositories docker proxy params
func (o *PostRepositoriesDockerProxyParams) WithContext(ctx context.Context) *PostRepositoriesDockerProxyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post repositories docker proxy params
func (o *PostRepositoriesDockerProxyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post repositories docker proxy params
func (o *PostRepositoriesDockerProxyParams) WithHTTPClient(client *http.Client) *PostRepositoriesDockerProxyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post repositories docker proxy params
func (o *PostRepositoriesDockerProxyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post repositories docker proxy params
func (o *PostRepositoriesDockerProxyParams) WithBody(body *models.DockerProxyRepositoryAPIRequest) *PostRepositoriesDockerProxyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post repositories docker proxy params
func (o *PostRepositoriesDockerProxyParams) SetBody(body *models.DockerProxyRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRepositoriesDockerProxyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
