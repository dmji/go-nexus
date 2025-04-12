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

// NewPostRepositoriesCondaProxyParams creates a new PostRepositoriesCondaProxyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostRepositoriesCondaProxyParams() *PostRepositoriesCondaProxyParams {
	return &PostRepositoriesCondaProxyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostRepositoriesCondaProxyParamsWithTimeout creates a new PostRepositoriesCondaProxyParams object
// with the ability to set a timeout on a request.
func NewPostRepositoriesCondaProxyParamsWithTimeout(timeout time.Duration) *PostRepositoriesCondaProxyParams {
	return &PostRepositoriesCondaProxyParams{
		timeout: timeout,
	}
}

// NewPostRepositoriesCondaProxyParamsWithContext creates a new PostRepositoriesCondaProxyParams object
// with the ability to set a context for a request.
func NewPostRepositoriesCondaProxyParamsWithContext(ctx context.Context) *PostRepositoriesCondaProxyParams {
	return &PostRepositoriesCondaProxyParams{
		Context: ctx,
	}
}

// NewPostRepositoriesCondaProxyParamsWithHTTPClient creates a new PostRepositoriesCondaProxyParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostRepositoriesCondaProxyParamsWithHTTPClient(client *http.Client) *PostRepositoriesCondaProxyParams {
	return &PostRepositoriesCondaProxyParams{
		HTTPClient: client,
	}
}

/*
PostRepositoriesCondaProxyParams contains all the parameters to send to the API endpoint

	for the post repositories conda proxy operation.

	Typically these are written to a http.Request.
*/
type PostRepositoriesCondaProxyParams struct {

	// Body.
	Body *models.CondaProxyRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post repositories conda proxy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRepositoriesCondaProxyParams) WithDefaults() *PostRepositoriesCondaProxyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post repositories conda proxy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRepositoriesCondaProxyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post repositories conda proxy params
func (o *PostRepositoriesCondaProxyParams) WithTimeout(timeout time.Duration) *PostRepositoriesCondaProxyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post repositories conda proxy params
func (o *PostRepositoriesCondaProxyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post repositories conda proxy params
func (o *PostRepositoriesCondaProxyParams) WithContext(ctx context.Context) *PostRepositoriesCondaProxyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post repositories conda proxy params
func (o *PostRepositoriesCondaProxyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post repositories conda proxy params
func (o *PostRepositoriesCondaProxyParams) WithHTTPClient(client *http.Client) *PostRepositoriesCondaProxyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post repositories conda proxy params
func (o *PostRepositoriesCondaProxyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post repositories conda proxy params
func (o *PostRepositoriesCondaProxyParams) WithBody(body *models.CondaProxyRepositoryAPIRequest) *PostRepositoriesCondaProxyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post repositories conda proxy params
func (o *PostRepositoriesCondaProxyParams) SetBody(body *models.CondaProxyRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRepositoriesCondaProxyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
