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

// NewPostSecurityPrivilegesRepositoryAdminParams creates a new PostSecurityPrivilegesRepositoryAdminParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostSecurityPrivilegesRepositoryAdminParams() *PostSecurityPrivilegesRepositoryAdminParams {
	return &PostSecurityPrivilegesRepositoryAdminParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostSecurityPrivilegesRepositoryAdminParamsWithTimeout creates a new PostSecurityPrivilegesRepositoryAdminParams object
// with the ability to set a timeout on a request.
func NewPostSecurityPrivilegesRepositoryAdminParamsWithTimeout(timeout time.Duration) *PostSecurityPrivilegesRepositoryAdminParams {
	return &PostSecurityPrivilegesRepositoryAdminParams{
		timeout: timeout,
	}
}

// NewPostSecurityPrivilegesRepositoryAdminParamsWithContext creates a new PostSecurityPrivilegesRepositoryAdminParams object
// with the ability to set a context for a request.
func NewPostSecurityPrivilegesRepositoryAdminParamsWithContext(ctx context.Context) *PostSecurityPrivilegesRepositoryAdminParams {
	return &PostSecurityPrivilegesRepositoryAdminParams{
		Context: ctx,
	}
}

// NewPostSecurityPrivilegesRepositoryAdminParamsWithHTTPClient creates a new PostSecurityPrivilegesRepositoryAdminParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostSecurityPrivilegesRepositoryAdminParamsWithHTTPClient(client *http.Client) *PostSecurityPrivilegesRepositoryAdminParams {
	return &PostSecurityPrivilegesRepositoryAdminParams{
		HTTPClient: client,
	}
}

/*
PostSecurityPrivilegesRepositoryAdminParams contains all the parameters to send to the API endpoint

	for the post security privileges repository admin operation.

	Typically these are written to a http.Request.
*/
type PostSecurityPrivilegesRepositoryAdminParams struct {

	/* Body.

	   The privilege to create.
	*/
	Body *models.APIPrivilegeRepositoryAdminRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post security privileges repository admin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSecurityPrivilegesRepositoryAdminParams) WithDefaults() *PostSecurityPrivilegesRepositoryAdminParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post security privileges repository admin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSecurityPrivilegesRepositoryAdminParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post security privileges repository admin params
func (o *PostSecurityPrivilegesRepositoryAdminParams) WithTimeout(timeout time.Duration) *PostSecurityPrivilegesRepositoryAdminParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post security privileges repository admin params
func (o *PostSecurityPrivilegesRepositoryAdminParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post security privileges repository admin params
func (o *PostSecurityPrivilegesRepositoryAdminParams) WithContext(ctx context.Context) *PostSecurityPrivilegesRepositoryAdminParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post security privileges repository admin params
func (o *PostSecurityPrivilegesRepositoryAdminParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post security privileges repository admin params
func (o *PostSecurityPrivilegesRepositoryAdminParams) WithHTTPClient(client *http.Client) *PostSecurityPrivilegesRepositoryAdminParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post security privileges repository admin params
func (o *PostSecurityPrivilegesRepositoryAdminParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post security privileges repository admin params
func (o *PostSecurityPrivilegesRepositoryAdminParams) WithBody(body *models.APIPrivilegeRepositoryAdminRequest) *PostSecurityPrivilegesRepositoryAdminParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post security privileges repository admin params
func (o *PostSecurityPrivilegesRepositoryAdminParams) SetBody(body *models.APIPrivilegeRepositoryAdminRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostSecurityPrivilegesRepositoryAdminParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
