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

	"github.com/dmji/go-nexus/models"
)

// NewCreatePrivilege3Params creates a new CreatePrivilege3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePrivilege3Params() *CreatePrivilege3Params {
	return &CreatePrivilege3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePrivilege3ParamsWithTimeout creates a new CreatePrivilege3Params object
// with the ability to set a timeout on a request.
func NewCreatePrivilege3ParamsWithTimeout(timeout time.Duration) *CreatePrivilege3Params {
	return &CreatePrivilege3Params{
		timeout: timeout,
	}
}

// NewCreatePrivilege3ParamsWithContext creates a new CreatePrivilege3Params object
// with the ability to set a context for a request.
func NewCreatePrivilege3ParamsWithContext(ctx context.Context) *CreatePrivilege3Params {
	return &CreatePrivilege3Params{
		Context: ctx,
	}
}

// NewCreatePrivilege3ParamsWithHTTPClient creates a new CreatePrivilege3Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePrivilege3ParamsWithHTTPClient(client *http.Client) *CreatePrivilege3Params {
	return &CreatePrivilege3Params{
		HTTPClient: client,
	}
}

/*
CreatePrivilege3Params contains all the parameters to send to the API endpoint

	for the create privilege 3 operation.

	Typically these are written to a http.Request.
*/
type CreatePrivilege3Params struct {

	/* Body.

	   The privilege to create.
	*/
	Body *models.APIPrivilegeWildcardRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create privilege 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePrivilege3Params) WithDefaults() *CreatePrivilege3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create privilege 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePrivilege3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create privilege 3 params
func (o *CreatePrivilege3Params) WithTimeout(timeout time.Duration) *CreatePrivilege3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create privilege 3 params
func (o *CreatePrivilege3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create privilege 3 params
func (o *CreatePrivilege3Params) WithContext(ctx context.Context) *CreatePrivilege3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create privilege 3 params
func (o *CreatePrivilege3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create privilege 3 params
func (o *CreatePrivilege3Params) WithHTTPClient(client *http.Client) *CreatePrivilege3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create privilege 3 params
func (o *CreatePrivilege3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create privilege 3 params
func (o *CreatePrivilege3Params) WithBody(body *models.APIPrivilegeWildcardRequest) *CreatePrivilege3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create privilege 3 params
func (o *CreatePrivilege3Params) SetBody(body *models.APIPrivilegeWildcardRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePrivilege3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
