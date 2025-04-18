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

	"github.com/dmji/go-nexus/models"
)

// NewPostScriptParams creates a new PostScriptParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostScriptParams() *PostScriptParams {
	return &PostScriptParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostScriptParamsWithTimeout creates a new PostScriptParams object
// with the ability to set a timeout on a request.
func NewPostScriptParamsWithTimeout(timeout time.Duration) *PostScriptParams {
	return &PostScriptParams{
		timeout: timeout,
	}
}

// NewPostScriptParamsWithContext creates a new PostScriptParams object
// with the ability to set a context for a request.
func NewPostScriptParamsWithContext(ctx context.Context) *PostScriptParams {
	return &PostScriptParams{
		Context: ctx,
	}
}

// NewPostScriptParamsWithHTTPClient creates a new PostScriptParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostScriptParamsWithHTTPClient(client *http.Client) *PostScriptParams {
	return &PostScriptParams{
		HTTPClient: client,
	}
}

/*
PostScriptParams contains all the parameters to send to the API endpoint

	for the post script operation.

	Typically these are written to a http.Request.
*/
type PostScriptParams struct {

	// Body.
	Body *models.ScriptXO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post script params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostScriptParams) WithDefaults() *PostScriptParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post script params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostScriptParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post script params
func (o *PostScriptParams) WithTimeout(timeout time.Duration) *PostScriptParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post script params
func (o *PostScriptParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post script params
func (o *PostScriptParams) WithContext(ctx context.Context) *PostScriptParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post script params
func (o *PostScriptParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post script params
func (o *PostScriptParams) WithHTTPClient(client *http.Client) *PostScriptParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post script params
func (o *PostScriptParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post script params
func (o *PostScriptParams) WithBody(body *models.ScriptXO) *PostScriptParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post script params
func (o *PostScriptParams) SetBody(body *models.ScriptXO) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostScriptParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
