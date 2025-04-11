// Code generated by go-swagger; DO NOT EDIT.

package azure_blob_store

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

	"nexus/models"
)

// NewVerifyConnectionParams creates a new VerifyConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVerifyConnectionParams() *VerifyConnectionParams {
	return &VerifyConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVerifyConnectionParamsWithTimeout creates a new VerifyConnectionParams object
// with the ability to set a timeout on a request.
func NewVerifyConnectionParamsWithTimeout(timeout time.Duration) *VerifyConnectionParams {
	return &VerifyConnectionParams{
		timeout: timeout,
	}
}

// NewVerifyConnectionParamsWithContext creates a new VerifyConnectionParams object
// with the ability to set a context for a request.
func NewVerifyConnectionParamsWithContext(ctx context.Context) *VerifyConnectionParams {
	return &VerifyConnectionParams{
		Context: ctx,
	}
}

// NewVerifyConnectionParamsWithHTTPClient creates a new VerifyConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewVerifyConnectionParamsWithHTTPClient(client *http.Client) *VerifyConnectionParams {
	return &VerifyConnectionParams{
		HTTPClient: client,
	}
}

/*
VerifyConnectionParams contains all the parameters to send to the API endpoint

	for the verify connection operation.

	Typically these are written to a http.Request.
*/
type VerifyConnectionParams struct {

	// Body.
	Body *models.AzureConnectionXO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the verify connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VerifyConnectionParams) WithDefaults() *VerifyConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the verify connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VerifyConnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the verify connection params
func (o *VerifyConnectionParams) WithTimeout(timeout time.Duration) *VerifyConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the verify connection params
func (o *VerifyConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the verify connection params
func (o *VerifyConnectionParams) WithContext(ctx context.Context) *VerifyConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the verify connection params
func (o *VerifyConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the verify connection params
func (o *VerifyConnectionParams) WithHTTPClient(client *http.Client) *VerifyConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the verify connection params
func (o *VerifyConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the verify connection params
func (o *VerifyConnectionParams) WithBody(body *models.AzureConnectionXO) *VerifyConnectionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the verify connection params
func (o *VerifyConnectionParams) SetBody(body *models.AzureConnectionXO) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *VerifyConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
