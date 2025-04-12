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
)

// NewDeleteSecurityPrivilegesByPrivilegenameParams creates a new DeleteSecurityPrivilegesByPrivilegenameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSecurityPrivilegesByPrivilegenameParams() *DeleteSecurityPrivilegesByPrivilegenameParams {
	return &DeleteSecurityPrivilegesByPrivilegenameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSecurityPrivilegesByPrivilegenameParamsWithTimeout creates a new DeleteSecurityPrivilegesByPrivilegenameParams object
// with the ability to set a timeout on a request.
func NewDeleteSecurityPrivilegesByPrivilegenameParamsWithTimeout(timeout time.Duration) *DeleteSecurityPrivilegesByPrivilegenameParams {
	return &DeleteSecurityPrivilegesByPrivilegenameParams{
		timeout: timeout,
	}
}

// NewDeleteSecurityPrivilegesByPrivilegenameParamsWithContext creates a new DeleteSecurityPrivilegesByPrivilegenameParams object
// with the ability to set a context for a request.
func NewDeleteSecurityPrivilegesByPrivilegenameParamsWithContext(ctx context.Context) *DeleteSecurityPrivilegesByPrivilegenameParams {
	return &DeleteSecurityPrivilegesByPrivilegenameParams{
		Context: ctx,
	}
}

// NewDeleteSecurityPrivilegesByPrivilegenameParamsWithHTTPClient creates a new DeleteSecurityPrivilegesByPrivilegenameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSecurityPrivilegesByPrivilegenameParamsWithHTTPClient(client *http.Client) *DeleteSecurityPrivilegesByPrivilegenameParams {
	return &DeleteSecurityPrivilegesByPrivilegenameParams{
		HTTPClient: client,
	}
}

/*
DeleteSecurityPrivilegesByPrivilegenameParams contains all the parameters to send to the API endpoint

	for the delete security privileges by privilegename operation.

	Typically these are written to a http.Request.
*/
type DeleteSecurityPrivilegesByPrivilegenameParams struct {

	/* PrivilegeName.

	   The name of the privilege to delete.
	*/
	PrivilegeName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete security privileges by privilegename params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSecurityPrivilegesByPrivilegenameParams) WithDefaults() *DeleteSecurityPrivilegesByPrivilegenameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete security privileges by privilegename params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSecurityPrivilegesByPrivilegenameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete security privileges by privilegename params
func (o *DeleteSecurityPrivilegesByPrivilegenameParams) WithTimeout(timeout time.Duration) *DeleteSecurityPrivilegesByPrivilegenameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete security privileges by privilegename params
func (o *DeleteSecurityPrivilegesByPrivilegenameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete security privileges by privilegename params
func (o *DeleteSecurityPrivilegesByPrivilegenameParams) WithContext(ctx context.Context) *DeleteSecurityPrivilegesByPrivilegenameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete security privileges by privilegename params
func (o *DeleteSecurityPrivilegesByPrivilegenameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete security privileges by privilegename params
func (o *DeleteSecurityPrivilegesByPrivilegenameParams) WithHTTPClient(client *http.Client) *DeleteSecurityPrivilegesByPrivilegenameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete security privileges by privilegename params
func (o *DeleteSecurityPrivilegesByPrivilegenameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPrivilegeName adds the privilegeName to the delete security privileges by privilegename params
func (o *DeleteSecurityPrivilegesByPrivilegenameParams) WithPrivilegeName(privilegeName string) *DeleteSecurityPrivilegesByPrivilegenameParams {
	o.SetPrivilegeName(privilegeName)
	return o
}

// SetPrivilegeName adds the privilegeName to the delete security privileges by privilegename params
func (o *DeleteSecurityPrivilegesByPrivilegenameParams) SetPrivilegeName(privilegeName string) {
	o.PrivilegeName = privilegeName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSecurityPrivilegesByPrivilegenameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param privilegeName
	if err := r.SetPathParam("privilegeName", o.PrivilegeName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
