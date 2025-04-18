// Code generated by go-swagger; DO NOT EDIT.

package read_only

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

// NewPostReadOnlyFreezeParams creates a new PostReadOnlyFreezeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostReadOnlyFreezeParams() *PostReadOnlyFreezeParams {
	return &PostReadOnlyFreezeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostReadOnlyFreezeParamsWithTimeout creates a new PostReadOnlyFreezeParams object
// with the ability to set a timeout on a request.
func NewPostReadOnlyFreezeParamsWithTimeout(timeout time.Duration) *PostReadOnlyFreezeParams {
	return &PostReadOnlyFreezeParams{
		timeout: timeout,
	}
}

// NewPostReadOnlyFreezeParamsWithContext creates a new PostReadOnlyFreezeParams object
// with the ability to set a context for a request.
func NewPostReadOnlyFreezeParamsWithContext(ctx context.Context) *PostReadOnlyFreezeParams {
	return &PostReadOnlyFreezeParams{
		Context: ctx,
	}
}

// NewPostReadOnlyFreezeParamsWithHTTPClient creates a new PostReadOnlyFreezeParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostReadOnlyFreezeParamsWithHTTPClient(client *http.Client) *PostReadOnlyFreezeParams {
	return &PostReadOnlyFreezeParams{
		HTTPClient: client,
	}
}

/*
PostReadOnlyFreezeParams contains all the parameters to send to the API endpoint

	for the post read only freeze operation.

	Typically these are written to a http.Request.
*/
type PostReadOnlyFreezeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post read only freeze params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostReadOnlyFreezeParams) WithDefaults() *PostReadOnlyFreezeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post read only freeze params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostReadOnlyFreezeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post read only freeze params
func (o *PostReadOnlyFreezeParams) WithTimeout(timeout time.Duration) *PostReadOnlyFreezeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post read only freeze params
func (o *PostReadOnlyFreezeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post read only freeze params
func (o *PostReadOnlyFreezeParams) WithContext(ctx context.Context) *PostReadOnlyFreezeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post read only freeze params
func (o *PostReadOnlyFreezeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post read only freeze params
func (o *PostReadOnlyFreezeParams) WithHTTPClient(client *http.Client) *PostReadOnlyFreezeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post read only freeze params
func (o *PostReadOnlyFreezeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostReadOnlyFreezeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
