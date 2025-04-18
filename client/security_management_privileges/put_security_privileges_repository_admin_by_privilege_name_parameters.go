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

// NewPutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams creates a new PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams() *PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams {
	return &PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutSecurityPrivilegesRepositoryAdminByPrivilegeNameParamsWithTimeout creates a new PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams object
// with the ability to set a timeout on a request.
func NewPutSecurityPrivilegesRepositoryAdminByPrivilegeNameParamsWithTimeout(timeout time.Duration) *PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams {
	return &PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams{
		timeout: timeout,
	}
}

// NewPutSecurityPrivilegesRepositoryAdminByPrivilegeNameParamsWithContext creates a new PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams object
// with the ability to set a context for a request.
func NewPutSecurityPrivilegesRepositoryAdminByPrivilegeNameParamsWithContext(ctx context.Context) *PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams {
	return &PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams{
		Context: ctx,
	}
}

// NewPutSecurityPrivilegesRepositoryAdminByPrivilegeNameParamsWithHTTPClient creates a new PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutSecurityPrivilegesRepositoryAdminByPrivilegeNameParamsWithHTTPClient(client *http.Client) *PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams {
	return &PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams{
		HTTPClient: client,
	}
}

/*
PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams contains all the parameters to send to the API endpoint

	for the put security privileges repository admin by privilege name operation.

	Typically these are written to a http.Request.
*/
type PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams struct {

	/* Body.

	   The privilege to update.
	*/
	Body *models.APIPrivilegeRepositoryAdminRequest

	/* PrivilegeName.

	   The name of the privilege to update.
	*/
	PrivilegeName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put security privileges repository admin by privilege name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams) WithDefaults() *PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put security privileges repository admin by privilege name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put security privileges repository admin by privilege name params
func (o *PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams) WithTimeout(timeout time.Duration) *PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put security privileges repository admin by privilege name params
func (o *PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put security privileges repository admin by privilege name params
func (o *PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams) WithContext(ctx context.Context) *PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put security privileges repository admin by privilege name params
func (o *PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put security privileges repository admin by privilege name params
func (o *PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams) WithHTTPClient(client *http.Client) *PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put security privileges repository admin by privilege name params
func (o *PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put security privileges repository admin by privilege name params
func (o *PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams) WithBody(body *models.APIPrivilegeRepositoryAdminRequest) *PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put security privileges repository admin by privilege name params
func (o *PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams) SetBody(body *models.APIPrivilegeRepositoryAdminRequest) {
	o.Body = body
}

// WithPrivilegeName adds the privilegeName to the put security privileges repository admin by privilege name params
func (o *PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams) WithPrivilegeName(privilegeName string) *PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams {
	o.SetPrivilegeName(privilegeName)
	return o
}

// SetPrivilegeName adds the privilegeName to the put security privileges repository admin by privilege name params
func (o *PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams) SetPrivilegeName(privilegeName string) {
	o.PrivilegeName = privilegeName
}

// WriteToRequest writes these params to a swagger request
func (o *PutSecurityPrivilegesRepositoryAdminByPrivilegeNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param privilegeName
	if err := r.SetPathParam("privilegeName", o.PrivilegeName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
