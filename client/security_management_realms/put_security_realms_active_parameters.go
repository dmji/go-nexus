// Code generated by go-swagger; DO NOT EDIT.

package security_management_realms

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

// NewPutSecurityRealmsActiveParams creates a new PutSecurityRealmsActiveParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutSecurityRealmsActiveParams() *PutSecurityRealmsActiveParams {
	return &PutSecurityRealmsActiveParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutSecurityRealmsActiveParamsWithTimeout creates a new PutSecurityRealmsActiveParams object
// with the ability to set a timeout on a request.
func NewPutSecurityRealmsActiveParamsWithTimeout(timeout time.Duration) *PutSecurityRealmsActiveParams {
	return &PutSecurityRealmsActiveParams{
		timeout: timeout,
	}
}

// NewPutSecurityRealmsActiveParamsWithContext creates a new PutSecurityRealmsActiveParams object
// with the ability to set a context for a request.
func NewPutSecurityRealmsActiveParamsWithContext(ctx context.Context) *PutSecurityRealmsActiveParams {
	return &PutSecurityRealmsActiveParams{
		Context: ctx,
	}
}

// NewPutSecurityRealmsActiveParamsWithHTTPClient creates a new PutSecurityRealmsActiveParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutSecurityRealmsActiveParamsWithHTTPClient(client *http.Client) *PutSecurityRealmsActiveParams {
	return &PutSecurityRealmsActiveParams{
		HTTPClient: client,
	}
}

/*
PutSecurityRealmsActiveParams contains all the parameters to send to the API endpoint

	for the put security realms active operation.

	Typically these are written to a http.Request.
*/
type PutSecurityRealmsActiveParams struct {

	/* Body.

	   The realm IDs
	*/
	Body []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put security realms active params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutSecurityRealmsActiveParams) WithDefaults() *PutSecurityRealmsActiveParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put security realms active params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutSecurityRealmsActiveParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put security realms active params
func (o *PutSecurityRealmsActiveParams) WithTimeout(timeout time.Duration) *PutSecurityRealmsActiveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put security realms active params
func (o *PutSecurityRealmsActiveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put security realms active params
func (o *PutSecurityRealmsActiveParams) WithContext(ctx context.Context) *PutSecurityRealmsActiveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put security realms active params
func (o *PutSecurityRealmsActiveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put security realms active params
func (o *PutSecurityRealmsActiveParams) WithHTTPClient(client *http.Client) *PutSecurityRealmsActiveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put security realms active params
func (o *PutSecurityRealmsActiveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put security realms active params
func (o *PutSecurityRealmsActiveParams) WithBody(body []string) *PutSecurityRealmsActiveParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put security realms active params
func (o *PutSecurityRealmsActiveParams) SetBody(body []string) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutSecurityRealmsActiveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
