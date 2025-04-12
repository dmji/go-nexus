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

// NewDeleteRoutingRulesByNameParams creates a new DeleteRoutingRulesByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRoutingRulesByNameParams() *DeleteRoutingRulesByNameParams {
	return &DeleteRoutingRulesByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRoutingRulesByNameParamsWithTimeout creates a new DeleteRoutingRulesByNameParams object
// with the ability to set a timeout on a request.
func NewDeleteRoutingRulesByNameParamsWithTimeout(timeout time.Duration) *DeleteRoutingRulesByNameParams {
	return &DeleteRoutingRulesByNameParams{
		timeout: timeout,
	}
}

// NewDeleteRoutingRulesByNameParamsWithContext creates a new DeleteRoutingRulesByNameParams object
// with the ability to set a context for a request.
func NewDeleteRoutingRulesByNameParamsWithContext(ctx context.Context) *DeleteRoutingRulesByNameParams {
	return &DeleteRoutingRulesByNameParams{
		Context: ctx,
	}
}

// NewDeleteRoutingRulesByNameParamsWithHTTPClient creates a new DeleteRoutingRulesByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRoutingRulesByNameParamsWithHTTPClient(client *http.Client) *DeleteRoutingRulesByNameParams {
	return &DeleteRoutingRulesByNameParams{
		HTTPClient: client,
	}
}

/*
DeleteRoutingRulesByNameParams contains all the parameters to send to the API endpoint

	for the delete routing rules by name operation.

	Typically these are written to a http.Request.
*/
type DeleteRoutingRulesByNameParams struct {

	/* Name.

	   The name of the routing rule to delete
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete routing rules by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRoutingRulesByNameParams) WithDefaults() *DeleteRoutingRulesByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete routing rules by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRoutingRulesByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete routing rules by name params
func (o *DeleteRoutingRulesByNameParams) WithTimeout(timeout time.Duration) *DeleteRoutingRulesByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete routing rules by name params
func (o *DeleteRoutingRulesByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete routing rules by name params
func (o *DeleteRoutingRulesByNameParams) WithContext(ctx context.Context) *DeleteRoutingRulesByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete routing rules by name params
func (o *DeleteRoutingRulesByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete routing rules by name params
func (o *DeleteRoutingRulesByNameParams) WithHTTPClient(client *http.Client) *DeleteRoutingRulesByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete routing rules by name params
func (o *DeleteRoutingRulesByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete routing rules by name params
func (o *DeleteRoutingRulesByNameParams) WithName(name string) *DeleteRoutingRulesByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete routing rules by name params
func (o *DeleteRoutingRulesByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRoutingRulesByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
