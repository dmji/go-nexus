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

	"github.com/dmji/go-nexus/models"
)

// NewPutRoutingRulesByNameParams creates a new PutRoutingRulesByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutRoutingRulesByNameParams() *PutRoutingRulesByNameParams {
	return &PutRoutingRulesByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutRoutingRulesByNameParamsWithTimeout creates a new PutRoutingRulesByNameParams object
// with the ability to set a timeout on a request.
func NewPutRoutingRulesByNameParamsWithTimeout(timeout time.Duration) *PutRoutingRulesByNameParams {
	return &PutRoutingRulesByNameParams{
		timeout: timeout,
	}
}

// NewPutRoutingRulesByNameParamsWithContext creates a new PutRoutingRulesByNameParams object
// with the ability to set a context for a request.
func NewPutRoutingRulesByNameParamsWithContext(ctx context.Context) *PutRoutingRulesByNameParams {
	return &PutRoutingRulesByNameParams{
		Context: ctx,
	}
}

// NewPutRoutingRulesByNameParamsWithHTTPClient creates a new PutRoutingRulesByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutRoutingRulesByNameParamsWithHTTPClient(client *http.Client) *PutRoutingRulesByNameParams {
	return &PutRoutingRulesByNameParams{
		HTTPClient: client,
	}
}

/*
PutRoutingRulesByNameParams contains all the parameters to send to the API endpoint

	for the put routing rules by name operation.

	Typically these are written to a http.Request.
*/
type PutRoutingRulesByNameParams struct {

	/* Body.

	   A routing rule configuration
	*/
	Body *models.RoutingRuleXO

	/* Name.

	   The name of the routing rule to update
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put routing rules by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRoutingRulesByNameParams) WithDefaults() *PutRoutingRulesByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put routing rules by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRoutingRulesByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put routing rules by name params
func (o *PutRoutingRulesByNameParams) WithTimeout(timeout time.Duration) *PutRoutingRulesByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put routing rules by name params
func (o *PutRoutingRulesByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put routing rules by name params
func (o *PutRoutingRulesByNameParams) WithContext(ctx context.Context) *PutRoutingRulesByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put routing rules by name params
func (o *PutRoutingRulesByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put routing rules by name params
func (o *PutRoutingRulesByNameParams) WithHTTPClient(client *http.Client) *PutRoutingRulesByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put routing rules by name params
func (o *PutRoutingRulesByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put routing rules by name params
func (o *PutRoutingRulesByNameParams) WithBody(body *models.RoutingRuleXO) *PutRoutingRulesByNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put routing rules by name params
func (o *PutRoutingRulesByNameParams) SetBody(body *models.RoutingRuleXO) {
	o.Body = body
}

// WithName adds the name to the put routing rules by name params
func (o *PutRoutingRulesByNameParams) WithName(name string) *PutRoutingRulesByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put routing rules by name params
func (o *PutRoutingRulesByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PutRoutingRulesByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
