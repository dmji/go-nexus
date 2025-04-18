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

// NewPostSecurityPrivilegesWildcardParams creates a new PostSecurityPrivilegesWildcardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostSecurityPrivilegesWildcardParams() *PostSecurityPrivilegesWildcardParams {
	return &PostSecurityPrivilegesWildcardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostSecurityPrivilegesWildcardParamsWithTimeout creates a new PostSecurityPrivilegesWildcardParams object
// with the ability to set a timeout on a request.
func NewPostSecurityPrivilegesWildcardParamsWithTimeout(timeout time.Duration) *PostSecurityPrivilegesWildcardParams {
	return &PostSecurityPrivilegesWildcardParams{
		timeout: timeout,
	}
}

// NewPostSecurityPrivilegesWildcardParamsWithContext creates a new PostSecurityPrivilegesWildcardParams object
// with the ability to set a context for a request.
func NewPostSecurityPrivilegesWildcardParamsWithContext(ctx context.Context) *PostSecurityPrivilegesWildcardParams {
	return &PostSecurityPrivilegesWildcardParams{
		Context: ctx,
	}
}

// NewPostSecurityPrivilegesWildcardParamsWithHTTPClient creates a new PostSecurityPrivilegesWildcardParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostSecurityPrivilegesWildcardParamsWithHTTPClient(client *http.Client) *PostSecurityPrivilegesWildcardParams {
	return &PostSecurityPrivilegesWildcardParams{
		HTTPClient: client,
	}
}

/*
PostSecurityPrivilegesWildcardParams contains all the parameters to send to the API endpoint

	for the post security privileges wildcard operation.

	Typically these are written to a http.Request.
*/
type PostSecurityPrivilegesWildcardParams struct {

	/* Body.

	   The privilege to create.
	*/
	Body *models.APIPrivilegeWildcardRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post security privileges wildcard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSecurityPrivilegesWildcardParams) WithDefaults() *PostSecurityPrivilegesWildcardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post security privileges wildcard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSecurityPrivilegesWildcardParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post security privileges wildcard params
func (o *PostSecurityPrivilegesWildcardParams) WithTimeout(timeout time.Duration) *PostSecurityPrivilegesWildcardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post security privileges wildcard params
func (o *PostSecurityPrivilegesWildcardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post security privileges wildcard params
func (o *PostSecurityPrivilegesWildcardParams) WithContext(ctx context.Context) *PostSecurityPrivilegesWildcardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post security privileges wildcard params
func (o *PostSecurityPrivilegesWildcardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post security privileges wildcard params
func (o *PostSecurityPrivilegesWildcardParams) WithHTTPClient(client *http.Client) *PostSecurityPrivilegesWildcardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post security privileges wildcard params
func (o *PostSecurityPrivilegesWildcardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post security privileges wildcard params
func (o *PostSecurityPrivilegesWildcardParams) WithBody(body *models.APIPrivilegeWildcardRequest) *PostSecurityPrivilegesWildcardParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post security privileges wildcard params
func (o *PostSecurityPrivilegesWildcardParams) SetBody(body *models.APIPrivilegeWildcardRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostSecurityPrivilegesWildcardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
