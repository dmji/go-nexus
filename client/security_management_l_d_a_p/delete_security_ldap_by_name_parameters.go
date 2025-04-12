// Code generated by go-swagger; DO NOT EDIT.

package security_management_l_d_a_p

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

// NewDeleteSecurityLdapByNameParams creates a new DeleteSecurityLdapByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSecurityLdapByNameParams() *DeleteSecurityLdapByNameParams {
	return &DeleteSecurityLdapByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSecurityLdapByNameParamsWithTimeout creates a new DeleteSecurityLdapByNameParams object
// with the ability to set a timeout on a request.
func NewDeleteSecurityLdapByNameParamsWithTimeout(timeout time.Duration) *DeleteSecurityLdapByNameParams {
	return &DeleteSecurityLdapByNameParams{
		timeout: timeout,
	}
}

// NewDeleteSecurityLdapByNameParamsWithContext creates a new DeleteSecurityLdapByNameParams object
// with the ability to set a context for a request.
func NewDeleteSecurityLdapByNameParamsWithContext(ctx context.Context) *DeleteSecurityLdapByNameParams {
	return &DeleteSecurityLdapByNameParams{
		Context: ctx,
	}
}

// NewDeleteSecurityLdapByNameParamsWithHTTPClient creates a new DeleteSecurityLdapByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSecurityLdapByNameParamsWithHTTPClient(client *http.Client) *DeleteSecurityLdapByNameParams {
	return &DeleteSecurityLdapByNameParams{
		HTTPClient: client,
	}
}

/*
DeleteSecurityLdapByNameParams contains all the parameters to send to the API endpoint

	for the delete security ldap by name operation.

	Typically these are written to a http.Request.
*/
type DeleteSecurityLdapByNameParams struct {

	/* Name.

	   Name of the LDAP server to delete
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete security ldap by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSecurityLdapByNameParams) WithDefaults() *DeleteSecurityLdapByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete security ldap by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSecurityLdapByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete security ldap by name params
func (o *DeleteSecurityLdapByNameParams) WithTimeout(timeout time.Duration) *DeleteSecurityLdapByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete security ldap by name params
func (o *DeleteSecurityLdapByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete security ldap by name params
func (o *DeleteSecurityLdapByNameParams) WithContext(ctx context.Context) *DeleteSecurityLdapByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete security ldap by name params
func (o *DeleteSecurityLdapByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete security ldap by name params
func (o *DeleteSecurityLdapByNameParams) WithHTTPClient(client *http.Client) *DeleteSecurityLdapByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete security ldap by name params
func (o *DeleteSecurityLdapByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete security ldap by name params
func (o *DeleteSecurityLdapByNameParams) WithName(name string) *DeleteSecurityLdapByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete security ldap by name params
func (o *DeleteSecurityLdapByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSecurityLdapByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
