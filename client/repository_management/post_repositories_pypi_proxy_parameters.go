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

// NewPostRepositoriesPypiProxyParams creates a new PostRepositoriesPypiProxyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostRepositoriesPypiProxyParams() *PostRepositoriesPypiProxyParams {
	return &PostRepositoriesPypiProxyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostRepositoriesPypiProxyParamsWithTimeout creates a new PostRepositoriesPypiProxyParams object
// with the ability to set a timeout on a request.
func NewPostRepositoriesPypiProxyParamsWithTimeout(timeout time.Duration) *PostRepositoriesPypiProxyParams {
	return &PostRepositoriesPypiProxyParams{
		timeout: timeout,
	}
}

// NewPostRepositoriesPypiProxyParamsWithContext creates a new PostRepositoriesPypiProxyParams object
// with the ability to set a context for a request.
func NewPostRepositoriesPypiProxyParamsWithContext(ctx context.Context) *PostRepositoriesPypiProxyParams {
	return &PostRepositoriesPypiProxyParams{
		Context: ctx,
	}
}

// NewPostRepositoriesPypiProxyParamsWithHTTPClient creates a new PostRepositoriesPypiProxyParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostRepositoriesPypiProxyParamsWithHTTPClient(client *http.Client) *PostRepositoriesPypiProxyParams {
	return &PostRepositoriesPypiProxyParams{
		HTTPClient: client,
	}
}

/*
PostRepositoriesPypiProxyParams contains all the parameters to send to the API endpoint

	for the post repositories pypi proxy operation.

	Typically these are written to a http.Request.
*/
type PostRepositoriesPypiProxyParams struct {

	// Body.
	Body *models.PypiProxyRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post repositories pypi proxy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRepositoriesPypiProxyParams) WithDefaults() *PostRepositoriesPypiProxyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post repositories pypi proxy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRepositoriesPypiProxyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post repositories pypi proxy params
func (o *PostRepositoriesPypiProxyParams) WithTimeout(timeout time.Duration) *PostRepositoriesPypiProxyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post repositories pypi proxy params
func (o *PostRepositoriesPypiProxyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post repositories pypi proxy params
func (o *PostRepositoriesPypiProxyParams) WithContext(ctx context.Context) *PostRepositoriesPypiProxyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post repositories pypi proxy params
func (o *PostRepositoriesPypiProxyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post repositories pypi proxy params
func (o *PostRepositoriesPypiProxyParams) WithHTTPClient(client *http.Client) *PostRepositoriesPypiProxyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post repositories pypi proxy params
func (o *PostRepositoriesPypiProxyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post repositories pypi proxy params
func (o *PostRepositoriesPypiProxyParams) WithBody(body *models.PypiProxyRepositoryAPIRequest) *PostRepositoriesPypiProxyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post repositories pypi proxy params
func (o *PostRepositoriesPypiProxyParams) SetBody(body *models.PypiProxyRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRepositoriesPypiProxyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
