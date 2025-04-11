// Code generated by go-swagger; DO NOT EDIT.

package malicious_risk_on_disk

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

// NewGetMaliciousRiskOnDiskCountParams creates a new GetMaliciousRiskOnDiskCountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMaliciousRiskOnDiskCountParams() *GetMaliciousRiskOnDiskCountParams {
	return &GetMaliciousRiskOnDiskCountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMaliciousRiskOnDiskCountParamsWithTimeout creates a new GetMaliciousRiskOnDiskCountParams object
// with the ability to set a timeout on a request.
func NewGetMaliciousRiskOnDiskCountParamsWithTimeout(timeout time.Duration) *GetMaliciousRiskOnDiskCountParams {
	return &GetMaliciousRiskOnDiskCountParams{
		timeout: timeout,
	}
}

// NewGetMaliciousRiskOnDiskCountParamsWithContext creates a new GetMaliciousRiskOnDiskCountParams object
// with the ability to set a context for a request.
func NewGetMaliciousRiskOnDiskCountParamsWithContext(ctx context.Context) *GetMaliciousRiskOnDiskCountParams {
	return &GetMaliciousRiskOnDiskCountParams{
		Context: ctx,
	}
}

// NewGetMaliciousRiskOnDiskCountParamsWithHTTPClient creates a new GetMaliciousRiskOnDiskCountParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMaliciousRiskOnDiskCountParamsWithHTTPClient(client *http.Client) *GetMaliciousRiskOnDiskCountParams {
	return &GetMaliciousRiskOnDiskCountParams{
		HTTPClient: client,
	}
}

/*
GetMaliciousRiskOnDiskCountParams contains all the parameters to send to the API endpoint

	for the get malicious risk on disk count operation.

	Typically these are written to a http.Request.
*/
type GetMaliciousRiskOnDiskCountParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get malicious risk on disk count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMaliciousRiskOnDiskCountParams) WithDefaults() *GetMaliciousRiskOnDiskCountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get malicious risk on disk count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMaliciousRiskOnDiskCountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get malicious risk on disk count params
func (o *GetMaliciousRiskOnDiskCountParams) WithTimeout(timeout time.Duration) *GetMaliciousRiskOnDiskCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get malicious risk on disk count params
func (o *GetMaliciousRiskOnDiskCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get malicious risk on disk count params
func (o *GetMaliciousRiskOnDiskCountParams) WithContext(ctx context.Context) *GetMaliciousRiskOnDiskCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get malicious risk on disk count params
func (o *GetMaliciousRiskOnDiskCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get malicious risk on disk count params
func (o *GetMaliciousRiskOnDiskCountParams) WithHTTPClient(client *http.Client) *GetMaliciousRiskOnDiskCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get malicious risk on disk count params
func (o *GetMaliciousRiskOnDiskCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetMaliciousRiskOnDiskCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
