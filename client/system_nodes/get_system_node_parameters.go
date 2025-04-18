// Code generated by go-swagger; DO NOT EDIT.

package system_nodes

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

// NewGetSystemNodeParams creates a new GetSystemNodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSystemNodeParams() *GetSystemNodeParams {
	return &GetSystemNodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemNodeParamsWithTimeout creates a new GetSystemNodeParams object
// with the ability to set a timeout on a request.
func NewGetSystemNodeParamsWithTimeout(timeout time.Duration) *GetSystemNodeParams {
	return &GetSystemNodeParams{
		timeout: timeout,
	}
}

// NewGetSystemNodeParamsWithContext creates a new GetSystemNodeParams object
// with the ability to set a context for a request.
func NewGetSystemNodeParamsWithContext(ctx context.Context) *GetSystemNodeParams {
	return &GetSystemNodeParams{
		Context: ctx,
	}
}

// NewGetSystemNodeParamsWithHTTPClient creates a new GetSystemNodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSystemNodeParamsWithHTTPClient(client *http.Client) *GetSystemNodeParams {
	return &GetSystemNodeParams{
		HTTPClient: client,
	}
}

/*
GetSystemNodeParams contains all the parameters to send to the API endpoint

	for the get system node operation.

	Typically these are written to a http.Request.
*/
type GetSystemNodeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get system node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSystemNodeParams) WithDefaults() *GetSystemNodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get system node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSystemNodeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get system node params
func (o *GetSystemNodeParams) WithTimeout(timeout time.Duration) *GetSystemNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system node params
func (o *GetSystemNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system node params
func (o *GetSystemNodeParams) WithContext(ctx context.Context) *GetSystemNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system node params
func (o *GetSystemNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get system node params
func (o *GetSystemNodeParams) WithHTTPClient(client *http.Client) *GetSystemNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get system node params
func (o *GetSystemNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
