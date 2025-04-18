// Code generated by go-swagger; DO NOT EDIT.

package manage_sonatype_repository_firewall_configuration

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

// NewPostIqVerifyConnectionParams creates a new PostIqVerifyConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostIqVerifyConnectionParams() *PostIqVerifyConnectionParams {
	return &PostIqVerifyConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostIqVerifyConnectionParamsWithTimeout creates a new PostIqVerifyConnectionParams object
// with the ability to set a timeout on a request.
func NewPostIqVerifyConnectionParamsWithTimeout(timeout time.Duration) *PostIqVerifyConnectionParams {
	return &PostIqVerifyConnectionParams{
		timeout: timeout,
	}
}

// NewPostIqVerifyConnectionParamsWithContext creates a new PostIqVerifyConnectionParams object
// with the ability to set a context for a request.
func NewPostIqVerifyConnectionParamsWithContext(ctx context.Context) *PostIqVerifyConnectionParams {
	return &PostIqVerifyConnectionParams{
		Context: ctx,
	}
}

// NewPostIqVerifyConnectionParamsWithHTTPClient creates a new PostIqVerifyConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostIqVerifyConnectionParamsWithHTTPClient(client *http.Client) *PostIqVerifyConnectionParams {
	return &PostIqVerifyConnectionParams{
		HTTPClient: client,
	}
}

/*
PostIqVerifyConnectionParams contains all the parameters to send to the API endpoint

	for the post iq verify connection operation.

	Typically these are written to a http.Request.
*/
type PostIqVerifyConnectionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post iq verify connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostIqVerifyConnectionParams) WithDefaults() *PostIqVerifyConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post iq verify connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostIqVerifyConnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post iq verify connection params
func (o *PostIqVerifyConnectionParams) WithTimeout(timeout time.Duration) *PostIqVerifyConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post iq verify connection params
func (o *PostIqVerifyConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post iq verify connection params
func (o *PostIqVerifyConnectionParams) WithContext(ctx context.Context) *PostIqVerifyConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post iq verify connection params
func (o *PostIqVerifyConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post iq verify connection params
func (o *PostIqVerifyConnectionParams) WithHTTPClient(client *http.Client) *PostIqVerifyConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post iq verify connection params
func (o *PostIqVerifyConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostIqVerifyConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
