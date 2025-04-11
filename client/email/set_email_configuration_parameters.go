// Code generated by go-swagger; DO NOT EDIT.

package email

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

	"nexus/models"
)

// NewSetEmailConfigurationParams creates a new SetEmailConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetEmailConfigurationParams() *SetEmailConfigurationParams {
	return &SetEmailConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetEmailConfigurationParamsWithTimeout creates a new SetEmailConfigurationParams object
// with the ability to set a timeout on a request.
func NewSetEmailConfigurationParamsWithTimeout(timeout time.Duration) *SetEmailConfigurationParams {
	return &SetEmailConfigurationParams{
		timeout: timeout,
	}
}

// NewSetEmailConfigurationParamsWithContext creates a new SetEmailConfigurationParams object
// with the ability to set a context for a request.
func NewSetEmailConfigurationParamsWithContext(ctx context.Context) *SetEmailConfigurationParams {
	return &SetEmailConfigurationParams{
		Context: ctx,
	}
}

// NewSetEmailConfigurationParamsWithHTTPClient creates a new SetEmailConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetEmailConfigurationParamsWithHTTPClient(client *http.Client) *SetEmailConfigurationParams {
	return &SetEmailConfigurationParams{
		HTTPClient: client,
	}
}

/*
SetEmailConfigurationParams contains all the parameters to send to the API endpoint

	for the set email configuration operation.

	Typically these are written to a http.Request.
*/
type SetEmailConfigurationParams struct {

	// Body.
	Body *models.APIEmailConfiguration

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set email configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetEmailConfigurationParams) WithDefaults() *SetEmailConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set email configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetEmailConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set email configuration params
func (o *SetEmailConfigurationParams) WithTimeout(timeout time.Duration) *SetEmailConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set email configuration params
func (o *SetEmailConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set email configuration params
func (o *SetEmailConfigurationParams) WithContext(ctx context.Context) *SetEmailConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set email configuration params
func (o *SetEmailConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set email configuration params
func (o *SetEmailConfigurationParams) WithHTTPClient(client *http.Client) *SetEmailConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set email configuration params
func (o *SetEmailConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set email configuration params
func (o *SetEmailConfigurationParams) WithBody(body *models.APIEmailConfiguration) *SetEmailConfigurationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set email configuration params
func (o *SetEmailConfigurationParams) SetBody(body *models.APIEmailConfiguration) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SetEmailConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
