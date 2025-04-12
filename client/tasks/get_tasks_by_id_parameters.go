// Code generated by go-swagger; DO NOT EDIT.

package tasks

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

// NewGetTasksByIDParams creates a new GetTasksByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTasksByIDParams() *GetTasksByIDParams {
	return &GetTasksByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTasksByIDParamsWithTimeout creates a new GetTasksByIDParams object
// with the ability to set a timeout on a request.
func NewGetTasksByIDParamsWithTimeout(timeout time.Duration) *GetTasksByIDParams {
	return &GetTasksByIDParams{
		timeout: timeout,
	}
}

// NewGetTasksByIDParamsWithContext creates a new GetTasksByIDParams object
// with the ability to set a context for a request.
func NewGetTasksByIDParamsWithContext(ctx context.Context) *GetTasksByIDParams {
	return &GetTasksByIDParams{
		Context: ctx,
	}
}

// NewGetTasksByIDParamsWithHTTPClient creates a new GetTasksByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTasksByIDParamsWithHTTPClient(client *http.Client) *GetTasksByIDParams {
	return &GetTasksByIDParams{
		HTTPClient: client,
	}
}

/*
GetTasksByIDParams contains all the parameters to send to the API endpoint

	for the get tasks by Id operation.

	Typically these are written to a http.Request.
*/
type GetTasksByIDParams struct {

	/* ID.

	   Id of the task to get
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get tasks by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTasksByIDParams) WithDefaults() *GetTasksByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get tasks by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTasksByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get tasks by Id params
func (o *GetTasksByIDParams) WithTimeout(timeout time.Duration) *GetTasksByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tasks by Id params
func (o *GetTasksByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tasks by Id params
func (o *GetTasksByIDParams) WithContext(ctx context.Context) *GetTasksByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tasks by Id params
func (o *GetTasksByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tasks by Id params
func (o *GetTasksByIDParams) WithHTTPClient(client *http.Client) *GetTasksByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tasks by Id params
func (o *GetTasksByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get tasks by Id params
func (o *GetTasksByIDParams) WithID(id string) *GetTasksByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get tasks by Id params
func (o *GetTasksByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetTasksByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
