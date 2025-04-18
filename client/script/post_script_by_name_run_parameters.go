// Code generated by go-swagger; DO NOT EDIT.

package script

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

// NewPostScriptByNameRunParams creates a new PostScriptByNameRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostScriptByNameRunParams() *PostScriptByNameRunParams {
	return &PostScriptByNameRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostScriptByNameRunParamsWithTimeout creates a new PostScriptByNameRunParams object
// with the ability to set a timeout on a request.
func NewPostScriptByNameRunParamsWithTimeout(timeout time.Duration) *PostScriptByNameRunParams {
	return &PostScriptByNameRunParams{
		timeout: timeout,
	}
}

// NewPostScriptByNameRunParamsWithContext creates a new PostScriptByNameRunParams object
// with the ability to set a context for a request.
func NewPostScriptByNameRunParamsWithContext(ctx context.Context) *PostScriptByNameRunParams {
	return &PostScriptByNameRunParams{
		Context: ctx,
	}
}

// NewPostScriptByNameRunParamsWithHTTPClient creates a new PostScriptByNameRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostScriptByNameRunParamsWithHTTPClient(client *http.Client) *PostScriptByNameRunParams {
	return &PostScriptByNameRunParams{
		HTTPClient: client,
	}
}

/*
PostScriptByNameRunParams contains all the parameters to send to the API endpoint

	for the post script by name run operation.

	Typically these are written to a http.Request.
*/
type PostScriptByNameRunParams struct {

	// Body.
	Body string

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post script by name run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostScriptByNameRunParams) WithDefaults() *PostScriptByNameRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post script by name run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostScriptByNameRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post script by name run params
func (o *PostScriptByNameRunParams) WithTimeout(timeout time.Duration) *PostScriptByNameRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post script by name run params
func (o *PostScriptByNameRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post script by name run params
func (o *PostScriptByNameRunParams) WithContext(ctx context.Context) *PostScriptByNameRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post script by name run params
func (o *PostScriptByNameRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post script by name run params
func (o *PostScriptByNameRunParams) WithHTTPClient(client *http.Client) *PostScriptByNameRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post script by name run params
func (o *PostScriptByNameRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post script by name run params
func (o *PostScriptByNameRunParams) WithBody(body string) *PostScriptByNameRunParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post script by name run params
func (o *PostScriptByNameRunParams) SetBody(body string) {
	o.Body = body
}

// WithName adds the name to the post script by name run params
func (o *PostScriptByNameRunParams) WithName(name string) *PostScriptByNameRunParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the post script by name run params
func (o *PostScriptByNameRunParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PostScriptByNameRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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
