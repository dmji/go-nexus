// Code generated by go-swagger; DO NOT EDIT.

package formats

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

// NewGetFormatsUploadSpecsParams creates a new GetFormatsUploadSpecsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFormatsUploadSpecsParams() *GetFormatsUploadSpecsParams {
	return &GetFormatsUploadSpecsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFormatsUploadSpecsParamsWithTimeout creates a new GetFormatsUploadSpecsParams object
// with the ability to set a timeout on a request.
func NewGetFormatsUploadSpecsParamsWithTimeout(timeout time.Duration) *GetFormatsUploadSpecsParams {
	return &GetFormatsUploadSpecsParams{
		timeout: timeout,
	}
}

// NewGetFormatsUploadSpecsParamsWithContext creates a new GetFormatsUploadSpecsParams object
// with the ability to set a context for a request.
func NewGetFormatsUploadSpecsParamsWithContext(ctx context.Context) *GetFormatsUploadSpecsParams {
	return &GetFormatsUploadSpecsParams{
		Context: ctx,
	}
}

// NewGetFormatsUploadSpecsParamsWithHTTPClient creates a new GetFormatsUploadSpecsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFormatsUploadSpecsParamsWithHTTPClient(client *http.Client) *GetFormatsUploadSpecsParams {
	return &GetFormatsUploadSpecsParams{
		HTTPClient: client,
	}
}

/*
GetFormatsUploadSpecsParams contains all the parameters to send to the API endpoint

	for the get formats upload specs operation.

	Typically these are written to a http.Request.
*/
type GetFormatsUploadSpecsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get formats upload specs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFormatsUploadSpecsParams) WithDefaults() *GetFormatsUploadSpecsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get formats upload specs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFormatsUploadSpecsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get formats upload specs params
func (o *GetFormatsUploadSpecsParams) WithTimeout(timeout time.Duration) *GetFormatsUploadSpecsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get formats upload specs params
func (o *GetFormatsUploadSpecsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get formats upload specs params
func (o *GetFormatsUploadSpecsParams) WithContext(ctx context.Context) *GetFormatsUploadSpecsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get formats upload specs params
func (o *GetFormatsUploadSpecsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get formats upload specs params
func (o *GetFormatsUploadSpecsParams) WithHTTPClient(client *http.Client) *GetFormatsUploadSpecsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get formats upload specs params
func (o *GetFormatsUploadSpecsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetFormatsUploadSpecsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
