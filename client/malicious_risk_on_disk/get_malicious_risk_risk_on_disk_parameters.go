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

// NewGetMaliciousRiskRiskOnDiskParams creates a new GetMaliciousRiskRiskOnDiskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMaliciousRiskRiskOnDiskParams() *GetMaliciousRiskRiskOnDiskParams {
	return &GetMaliciousRiskRiskOnDiskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMaliciousRiskRiskOnDiskParamsWithTimeout creates a new GetMaliciousRiskRiskOnDiskParams object
// with the ability to set a timeout on a request.
func NewGetMaliciousRiskRiskOnDiskParamsWithTimeout(timeout time.Duration) *GetMaliciousRiskRiskOnDiskParams {
	return &GetMaliciousRiskRiskOnDiskParams{
		timeout: timeout,
	}
}

// NewGetMaliciousRiskRiskOnDiskParamsWithContext creates a new GetMaliciousRiskRiskOnDiskParams object
// with the ability to set a context for a request.
func NewGetMaliciousRiskRiskOnDiskParamsWithContext(ctx context.Context) *GetMaliciousRiskRiskOnDiskParams {
	return &GetMaliciousRiskRiskOnDiskParams{
		Context: ctx,
	}
}

// NewGetMaliciousRiskRiskOnDiskParamsWithHTTPClient creates a new GetMaliciousRiskRiskOnDiskParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMaliciousRiskRiskOnDiskParamsWithHTTPClient(client *http.Client) *GetMaliciousRiskRiskOnDiskParams {
	return &GetMaliciousRiskRiskOnDiskParams{
		HTTPClient: client,
	}
}

/*
GetMaliciousRiskRiskOnDiskParams contains all the parameters to send to the API endpoint

	for the get malicious risk risk on disk operation.

	Typically these are written to a http.Request.
*/
type GetMaliciousRiskRiskOnDiskParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get malicious risk risk on disk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMaliciousRiskRiskOnDiskParams) WithDefaults() *GetMaliciousRiskRiskOnDiskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get malicious risk risk on disk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMaliciousRiskRiskOnDiskParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get malicious risk risk on disk params
func (o *GetMaliciousRiskRiskOnDiskParams) WithTimeout(timeout time.Duration) *GetMaliciousRiskRiskOnDiskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get malicious risk risk on disk params
func (o *GetMaliciousRiskRiskOnDiskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get malicious risk risk on disk params
func (o *GetMaliciousRiskRiskOnDiskParams) WithContext(ctx context.Context) *GetMaliciousRiskRiskOnDiskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get malicious risk risk on disk params
func (o *GetMaliciousRiskRiskOnDiskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get malicious risk risk on disk params
func (o *GetMaliciousRiskRiskOnDiskParams) WithHTTPClient(client *http.Client) *GetMaliciousRiskRiskOnDiskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get malicious risk risk on disk params
func (o *GetMaliciousRiskRiskOnDiskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetMaliciousRiskRiskOnDiskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
