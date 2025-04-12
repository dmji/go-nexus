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

// NewPostRepositoriesYumHostedParams creates a new PostRepositoriesYumHostedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostRepositoriesYumHostedParams() *PostRepositoriesYumHostedParams {
	return &PostRepositoriesYumHostedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostRepositoriesYumHostedParamsWithTimeout creates a new PostRepositoriesYumHostedParams object
// with the ability to set a timeout on a request.
func NewPostRepositoriesYumHostedParamsWithTimeout(timeout time.Duration) *PostRepositoriesYumHostedParams {
	return &PostRepositoriesYumHostedParams{
		timeout: timeout,
	}
}

// NewPostRepositoriesYumHostedParamsWithContext creates a new PostRepositoriesYumHostedParams object
// with the ability to set a context for a request.
func NewPostRepositoriesYumHostedParamsWithContext(ctx context.Context) *PostRepositoriesYumHostedParams {
	return &PostRepositoriesYumHostedParams{
		Context: ctx,
	}
}

// NewPostRepositoriesYumHostedParamsWithHTTPClient creates a new PostRepositoriesYumHostedParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostRepositoriesYumHostedParamsWithHTTPClient(client *http.Client) *PostRepositoriesYumHostedParams {
	return &PostRepositoriesYumHostedParams{
		HTTPClient: client,
	}
}

/*
PostRepositoriesYumHostedParams contains all the parameters to send to the API endpoint

	for the post repositories yum hosted operation.

	Typically these are written to a http.Request.
*/
type PostRepositoriesYumHostedParams struct {

	// Body.
	Body *models.YumHostedRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post repositories yum hosted params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRepositoriesYumHostedParams) WithDefaults() *PostRepositoriesYumHostedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post repositories yum hosted params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRepositoriesYumHostedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post repositories yum hosted params
func (o *PostRepositoriesYumHostedParams) WithTimeout(timeout time.Duration) *PostRepositoriesYumHostedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post repositories yum hosted params
func (o *PostRepositoriesYumHostedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post repositories yum hosted params
func (o *PostRepositoriesYumHostedParams) WithContext(ctx context.Context) *PostRepositoriesYumHostedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post repositories yum hosted params
func (o *PostRepositoriesYumHostedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post repositories yum hosted params
func (o *PostRepositoriesYumHostedParams) WithHTTPClient(client *http.Client) *PostRepositoriesYumHostedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post repositories yum hosted params
func (o *PostRepositoriesYumHostedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post repositories yum hosted params
func (o *PostRepositoriesYumHostedParams) WithBody(body *models.YumHostedRepositoryAPIRequest) *PostRepositoriesYumHostedParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post repositories yum hosted params
func (o *PostRepositoriesYumHostedParams) SetBody(body *models.YumHostedRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRepositoriesYumHostedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
