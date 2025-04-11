// Code generated by go-swagger; DO NOT EDIT.

package security_management_privileges

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

// NewCreatePrivilege4Params creates a new CreatePrivilege4Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePrivilege4Params() *CreatePrivilege4Params {
	return &CreatePrivilege4Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePrivilege4ParamsWithTimeout creates a new CreatePrivilege4Params object
// with the ability to set a timeout on a request.
func NewCreatePrivilege4ParamsWithTimeout(timeout time.Duration) *CreatePrivilege4Params {
	return &CreatePrivilege4Params{
		timeout: timeout,
	}
}

// NewCreatePrivilege4ParamsWithContext creates a new CreatePrivilege4Params object
// with the ability to set a context for a request.
func NewCreatePrivilege4ParamsWithContext(ctx context.Context) *CreatePrivilege4Params {
	return &CreatePrivilege4Params{
		Context: ctx,
	}
}

// NewCreatePrivilege4ParamsWithHTTPClient creates a new CreatePrivilege4Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePrivilege4ParamsWithHTTPClient(client *http.Client) *CreatePrivilege4Params {
	return &CreatePrivilege4Params{
		HTTPClient: client,
	}
}

/*
CreatePrivilege4Params contains all the parameters to send to the API endpoint

	for the create privilege 4 operation.

	Typically these are written to a http.Request.
*/
type CreatePrivilege4Params struct {

	/* Body.

	   The privilege to create.
	*/
	Body *models.APIPrivilegeApplicationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create privilege 4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePrivilege4Params) WithDefaults() *CreatePrivilege4Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create privilege 4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePrivilege4Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create privilege 4 params
func (o *CreatePrivilege4Params) WithTimeout(timeout time.Duration) *CreatePrivilege4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create privilege 4 params
func (o *CreatePrivilege4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create privilege 4 params
func (o *CreatePrivilege4Params) WithContext(ctx context.Context) *CreatePrivilege4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create privilege 4 params
func (o *CreatePrivilege4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create privilege 4 params
func (o *CreatePrivilege4Params) WithHTTPClient(client *http.Client) *CreatePrivilege4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create privilege 4 params
func (o *CreatePrivilege4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create privilege 4 params
func (o *CreatePrivilege4Params) WithBody(body *models.APIPrivilegeApplicationRequest) *CreatePrivilege4Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create privilege 4 params
func (o *CreatePrivilege4Params) SetBody(body *models.APIPrivilegeApplicationRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePrivilege4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
