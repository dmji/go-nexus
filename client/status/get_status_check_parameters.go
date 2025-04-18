// Code generated by go-swagger; DO NOT EDIT.

package status

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

// NewGetStatusCheckParams creates a new GetStatusCheckParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStatusCheckParams() *GetStatusCheckParams {
	return &GetStatusCheckParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStatusCheckParamsWithTimeout creates a new GetStatusCheckParams object
// with the ability to set a timeout on a request.
func NewGetStatusCheckParamsWithTimeout(timeout time.Duration) *GetStatusCheckParams {
	return &GetStatusCheckParams{
		timeout: timeout,
	}
}

// NewGetStatusCheckParamsWithContext creates a new GetStatusCheckParams object
// with the ability to set a context for a request.
func NewGetStatusCheckParamsWithContext(ctx context.Context) *GetStatusCheckParams {
	return &GetStatusCheckParams{
		Context: ctx,
	}
}

// NewGetStatusCheckParamsWithHTTPClient creates a new GetStatusCheckParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStatusCheckParamsWithHTTPClient(client *http.Client) *GetStatusCheckParams {
	return &GetStatusCheckParams{
		HTTPClient: client,
	}
}

/*
GetStatusCheckParams contains all the parameters to send to the API endpoint

	for the get status check operation.

	Typically these are written to a http.Request.
*/
type GetStatusCheckParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get status check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStatusCheckParams) WithDefaults() *GetStatusCheckParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get status check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStatusCheckParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get status check params
func (o *GetStatusCheckParams) WithTimeout(timeout time.Duration) *GetStatusCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get status check params
func (o *GetStatusCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get status check params
func (o *GetStatusCheckParams) WithContext(ctx context.Context) *GetStatusCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get status check params
func (o *GetStatusCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get status check params
func (o *GetStatusCheckParams) WithHTTPClient(client *http.Client) *GetStatusCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get status check params
func (o *GetStatusCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetStatusCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
