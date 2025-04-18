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

// NewGetReadOnlyParams creates a new GetReadOnlyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetReadOnlyParams() *GetReadOnlyParams {
	return &GetReadOnlyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetReadOnlyParamsWithTimeout creates a new GetReadOnlyParams object
// with the ability to set a timeout on a request.
func NewGetReadOnlyParamsWithTimeout(timeout time.Duration) *GetReadOnlyParams {
	return &GetReadOnlyParams{
		timeout: timeout,
	}
}

// NewGetReadOnlyParamsWithContext creates a new GetReadOnlyParams object
// with the ability to set a context for a request.
func NewGetReadOnlyParamsWithContext(ctx context.Context) *GetReadOnlyParams {
	return &GetReadOnlyParams{
		Context: ctx,
	}
}

// NewGetReadOnlyParamsWithHTTPClient creates a new GetReadOnlyParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetReadOnlyParamsWithHTTPClient(client *http.Client) *GetReadOnlyParams {
	return &GetReadOnlyParams{
		HTTPClient: client,
	}
}

/*
GetReadOnlyParams contains all the parameters to send to the API endpoint

	for the get read only operation.

	Typically these are written to a http.Request.
*/
type GetReadOnlyParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get read only params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReadOnlyParams) WithDefaults() *GetReadOnlyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get read only params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReadOnlyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get read only params
func (o *GetReadOnlyParams) WithTimeout(timeout time.Duration) *GetReadOnlyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get read only params
func (o *GetReadOnlyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get read only params
func (o *GetReadOnlyParams) WithContext(ctx context.Context) *GetReadOnlyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get read only params
func (o *GetReadOnlyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get read only params
func (o *GetReadOnlyParams) WithHTTPClient(client *http.Client) *GetReadOnlyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get read only params
func (o *GetReadOnlyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetReadOnlyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
