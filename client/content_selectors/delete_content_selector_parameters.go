// Code generated by go-swagger; DO NOT EDIT.

package content_selectors

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

// NewDeleteContentSelectorParams creates a new DeleteContentSelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteContentSelectorParams() *DeleteContentSelectorParams {
	return &DeleteContentSelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteContentSelectorParamsWithTimeout creates a new DeleteContentSelectorParams object
// with the ability to set a timeout on a request.
func NewDeleteContentSelectorParamsWithTimeout(timeout time.Duration) *DeleteContentSelectorParams {
	return &DeleteContentSelectorParams{
		timeout: timeout,
	}
}

// NewDeleteContentSelectorParamsWithContext creates a new DeleteContentSelectorParams object
// with the ability to set a context for a request.
func NewDeleteContentSelectorParamsWithContext(ctx context.Context) *DeleteContentSelectorParams {
	return &DeleteContentSelectorParams{
		Context: ctx,
	}
}

// NewDeleteContentSelectorParamsWithHTTPClient creates a new DeleteContentSelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteContentSelectorParamsWithHTTPClient(client *http.Client) *DeleteContentSelectorParams {
	return &DeleteContentSelectorParams{
		HTTPClient: client,
	}
}

/*
DeleteContentSelectorParams contains all the parameters to send to the API endpoint

	for the delete content selector operation.

	Typically these are written to a http.Request.
*/
type DeleteContentSelectorParams struct {

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete content selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteContentSelectorParams) WithDefaults() *DeleteContentSelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete content selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteContentSelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete content selector params
func (o *DeleteContentSelectorParams) WithTimeout(timeout time.Duration) *DeleteContentSelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete content selector params
func (o *DeleteContentSelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete content selector params
func (o *DeleteContentSelectorParams) WithContext(ctx context.Context) *DeleteContentSelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete content selector params
func (o *DeleteContentSelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete content selector params
func (o *DeleteContentSelectorParams) WithHTTPClient(client *http.Client) *DeleteContentSelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete content selector params
func (o *DeleteContentSelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete content selector params
func (o *DeleteContentSelectorParams) WithName(name string) *DeleteContentSelectorParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete content selector params
func (o *DeleteContentSelectorParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteContentSelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
