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

// NewGetTaskByIDParams creates a new GetTaskByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTaskByIDParams() *GetTaskByIDParams {
	return &GetTaskByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTaskByIDParamsWithTimeout creates a new GetTaskByIDParams object
// with the ability to set a timeout on a request.
func NewGetTaskByIDParamsWithTimeout(timeout time.Duration) *GetTaskByIDParams {
	return &GetTaskByIDParams{
		timeout: timeout,
	}
}

// NewGetTaskByIDParamsWithContext creates a new GetTaskByIDParams object
// with the ability to set a context for a request.
func NewGetTaskByIDParamsWithContext(ctx context.Context) *GetTaskByIDParams {
	return &GetTaskByIDParams{
		Context: ctx,
	}
}

// NewGetTaskByIDParamsWithHTTPClient creates a new GetTaskByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTaskByIDParamsWithHTTPClient(client *http.Client) *GetTaskByIDParams {
	return &GetTaskByIDParams{
		HTTPClient: client,
	}
}

/*
GetTaskByIDParams contains all the parameters to send to the API endpoint

	for the get task by Id operation.

	Typically these are written to a http.Request.
*/
type GetTaskByIDParams struct {

	/* ID.

	   Id of the task to get
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get task by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTaskByIDParams) WithDefaults() *GetTaskByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get task by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTaskByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get task by Id params
func (o *GetTaskByIDParams) WithTimeout(timeout time.Duration) *GetTaskByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get task by Id params
func (o *GetTaskByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get task by Id params
func (o *GetTaskByIDParams) WithContext(ctx context.Context) *GetTaskByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get task by Id params
func (o *GetTaskByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get task by Id params
func (o *GetTaskByIDParams) WithHTTPClient(client *http.Client) *GetTaskByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get task by Id params
func (o *GetTaskByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get task by Id params
func (o *GetTaskByIDParams) WithID(id string) *GetTaskByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get task by Id params
func (o *GetTaskByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetTaskByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
