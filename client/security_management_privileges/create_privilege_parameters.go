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

// NewCreatePrivilegeParams creates a new CreatePrivilegeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePrivilegeParams() *CreatePrivilegeParams {
	return &CreatePrivilegeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePrivilegeParamsWithTimeout creates a new CreatePrivilegeParams object
// with the ability to set a timeout on a request.
func NewCreatePrivilegeParamsWithTimeout(timeout time.Duration) *CreatePrivilegeParams {
	return &CreatePrivilegeParams{
		timeout: timeout,
	}
}

// NewCreatePrivilegeParamsWithContext creates a new CreatePrivilegeParams object
// with the ability to set a context for a request.
func NewCreatePrivilegeParamsWithContext(ctx context.Context) *CreatePrivilegeParams {
	return &CreatePrivilegeParams{
		Context: ctx,
	}
}

// NewCreatePrivilegeParamsWithHTTPClient creates a new CreatePrivilegeParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePrivilegeParamsWithHTTPClient(client *http.Client) *CreatePrivilegeParams {
	return &CreatePrivilegeParams{
		HTTPClient: client,
	}
}

/*
CreatePrivilegeParams contains all the parameters to send to the API endpoint

	for the create privilege operation.

	Typically these are written to a http.Request.
*/
type CreatePrivilegeParams struct {

	/* Body.

	   The privilege to create.
	*/
	Body *models.APIPrivilegeRepositoryContentSelectorRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create privilege params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePrivilegeParams) WithDefaults() *CreatePrivilegeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create privilege params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePrivilegeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create privilege params
func (o *CreatePrivilegeParams) WithTimeout(timeout time.Duration) *CreatePrivilegeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create privilege params
func (o *CreatePrivilegeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create privilege params
func (o *CreatePrivilegeParams) WithContext(ctx context.Context) *CreatePrivilegeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create privilege params
func (o *CreatePrivilegeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create privilege params
func (o *CreatePrivilegeParams) WithHTTPClient(client *http.Client) *CreatePrivilegeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create privilege params
func (o *CreatePrivilegeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create privilege params
func (o *CreatePrivilegeParams) WithBody(body *models.APIPrivilegeRepositoryContentSelectorRequest) *CreatePrivilegeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create privilege params
func (o *CreatePrivilegeParams) SetBody(body *models.APIPrivilegeRepositoryContentSelectorRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePrivilegeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
