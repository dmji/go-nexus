// Code generated by go-swagger; DO NOT EDIT.

package support

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

// NewSupportzipParams creates a new SupportzipParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSupportzipParams() *SupportzipParams {
	return &SupportzipParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSupportzipParamsWithTimeout creates a new SupportzipParams object
// with the ability to set a timeout on a request.
func NewSupportzipParamsWithTimeout(timeout time.Duration) *SupportzipParams {
	return &SupportzipParams{
		timeout: timeout,
	}
}

// NewSupportzipParamsWithContext creates a new SupportzipParams object
// with the ability to set a context for a request.
func NewSupportzipParamsWithContext(ctx context.Context) *SupportzipParams {
	return &SupportzipParams{
		Context: ctx,
	}
}

// NewSupportzipParamsWithHTTPClient creates a new SupportzipParams object
// with the ability to set a custom HTTPClient for a request.
func NewSupportzipParamsWithHTTPClient(client *http.Client) *SupportzipParams {
	return &SupportzipParams{
		HTTPClient: client,
	}
}

/*
SupportzipParams contains all the parameters to send to the API endpoint

	for the supportzip operation.

	Typically these are written to a http.Request.
*/
type SupportzipParams struct {

	// Body.
	Body *models.SupportZipGeneratorRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the supportzip params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SupportzipParams) WithDefaults() *SupportzipParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the supportzip params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SupportzipParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the supportzip params
func (o *SupportzipParams) WithTimeout(timeout time.Duration) *SupportzipParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the supportzip params
func (o *SupportzipParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the supportzip params
func (o *SupportzipParams) WithContext(ctx context.Context) *SupportzipParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the supportzip params
func (o *SupportzipParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the supportzip params
func (o *SupportzipParams) WithHTTPClient(client *http.Client) *SupportzipParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the supportzip params
func (o *SupportzipParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the supportzip params
func (o *SupportzipParams) WithBody(body *models.SupportZipGeneratorRequest) *SupportzipParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the supportzip params
func (o *SupportzipParams) SetBody(body *models.SupportZipGeneratorRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SupportzipParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
