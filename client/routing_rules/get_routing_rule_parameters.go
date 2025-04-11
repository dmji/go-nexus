// Code generated by go-swagger; DO NOT EDIT.

package routing_rules

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

// NewGetRoutingRuleParams creates a new GetRoutingRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRoutingRuleParams() *GetRoutingRuleParams {
	return &GetRoutingRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoutingRuleParamsWithTimeout creates a new GetRoutingRuleParams object
// with the ability to set a timeout on a request.
func NewGetRoutingRuleParamsWithTimeout(timeout time.Duration) *GetRoutingRuleParams {
	return &GetRoutingRuleParams{
		timeout: timeout,
	}
}

// NewGetRoutingRuleParamsWithContext creates a new GetRoutingRuleParams object
// with the ability to set a context for a request.
func NewGetRoutingRuleParamsWithContext(ctx context.Context) *GetRoutingRuleParams {
	return &GetRoutingRuleParams{
		Context: ctx,
	}
}

// NewGetRoutingRuleParamsWithHTTPClient creates a new GetRoutingRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRoutingRuleParamsWithHTTPClient(client *http.Client) *GetRoutingRuleParams {
	return &GetRoutingRuleParams{
		HTTPClient: client,
	}
}

/*
GetRoutingRuleParams contains all the parameters to send to the API endpoint

	for the get routing rule operation.

	Typically these are written to a http.Request.
*/
type GetRoutingRuleParams struct {

	/* Name.

	   The name of the routing rule to get
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get routing rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoutingRuleParams) WithDefaults() *GetRoutingRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get routing rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoutingRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get routing rule params
func (o *GetRoutingRuleParams) WithTimeout(timeout time.Duration) *GetRoutingRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get routing rule params
func (o *GetRoutingRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get routing rule params
func (o *GetRoutingRuleParams) WithContext(ctx context.Context) *GetRoutingRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get routing rule params
func (o *GetRoutingRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get routing rule params
func (o *GetRoutingRuleParams) WithHTTPClient(client *http.Client) *GetRoutingRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get routing rule params
func (o *GetRoutingRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get routing rule params
func (o *GetRoutingRuleParams) WithName(name string) *GetRoutingRuleParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get routing rule params
func (o *GetRoutingRuleParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoutingRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
