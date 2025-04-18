// Code generated by go-swagger; DO NOT EDIT.

package security_management_anonymous_access

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

// NewGetSecurityAnonymousParams creates a new GetSecurityAnonymousParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSecurityAnonymousParams() *GetSecurityAnonymousParams {
	return &GetSecurityAnonymousParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSecurityAnonymousParamsWithTimeout creates a new GetSecurityAnonymousParams object
// with the ability to set a timeout on a request.
func NewGetSecurityAnonymousParamsWithTimeout(timeout time.Duration) *GetSecurityAnonymousParams {
	return &GetSecurityAnonymousParams{
		timeout: timeout,
	}
}

// NewGetSecurityAnonymousParamsWithContext creates a new GetSecurityAnonymousParams object
// with the ability to set a context for a request.
func NewGetSecurityAnonymousParamsWithContext(ctx context.Context) *GetSecurityAnonymousParams {
	return &GetSecurityAnonymousParams{
		Context: ctx,
	}
}

// NewGetSecurityAnonymousParamsWithHTTPClient creates a new GetSecurityAnonymousParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSecurityAnonymousParamsWithHTTPClient(client *http.Client) *GetSecurityAnonymousParams {
	return &GetSecurityAnonymousParams{
		HTTPClient: client,
	}
}

/*
GetSecurityAnonymousParams contains all the parameters to send to the API endpoint

	for the get security anonymous operation.

	Typically these are written to a http.Request.
*/
type GetSecurityAnonymousParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get security anonymous params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSecurityAnonymousParams) WithDefaults() *GetSecurityAnonymousParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get security anonymous params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSecurityAnonymousParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get security anonymous params
func (o *GetSecurityAnonymousParams) WithTimeout(timeout time.Duration) *GetSecurityAnonymousParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get security anonymous params
func (o *GetSecurityAnonymousParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get security anonymous params
func (o *GetSecurityAnonymousParams) WithContext(ctx context.Context) *GetSecurityAnonymousParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get security anonymous params
func (o *GetSecurityAnonymousParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get security anonymous params
func (o *GetSecurityAnonymousParams) WithHTTPClient(client *http.Client) *GetSecurityAnonymousParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get security anonymous params
func (o *GetSecurityAnonymousParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSecurityAnonymousParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
