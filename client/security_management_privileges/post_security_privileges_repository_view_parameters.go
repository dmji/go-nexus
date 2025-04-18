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

// NewPostSecurityPrivilegesRepositoryViewParams creates a new PostSecurityPrivilegesRepositoryViewParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostSecurityPrivilegesRepositoryViewParams() *PostSecurityPrivilegesRepositoryViewParams {
	return &PostSecurityPrivilegesRepositoryViewParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostSecurityPrivilegesRepositoryViewParamsWithTimeout creates a new PostSecurityPrivilegesRepositoryViewParams object
// with the ability to set a timeout on a request.
func NewPostSecurityPrivilegesRepositoryViewParamsWithTimeout(timeout time.Duration) *PostSecurityPrivilegesRepositoryViewParams {
	return &PostSecurityPrivilegesRepositoryViewParams{
		timeout: timeout,
	}
}

// NewPostSecurityPrivilegesRepositoryViewParamsWithContext creates a new PostSecurityPrivilegesRepositoryViewParams object
// with the ability to set a context for a request.
func NewPostSecurityPrivilegesRepositoryViewParamsWithContext(ctx context.Context) *PostSecurityPrivilegesRepositoryViewParams {
	return &PostSecurityPrivilegesRepositoryViewParams{
		Context: ctx,
	}
}

// NewPostSecurityPrivilegesRepositoryViewParamsWithHTTPClient creates a new PostSecurityPrivilegesRepositoryViewParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostSecurityPrivilegesRepositoryViewParamsWithHTTPClient(client *http.Client) *PostSecurityPrivilegesRepositoryViewParams {
	return &PostSecurityPrivilegesRepositoryViewParams{
		HTTPClient: client,
	}
}

/*
PostSecurityPrivilegesRepositoryViewParams contains all the parameters to send to the API endpoint

	for the post security privileges repository view operation.

	Typically these are written to a http.Request.
*/
type PostSecurityPrivilegesRepositoryViewParams struct {

	/* Body.

	   The privilege to create.
	*/
	Body *models.APIPrivilegeRepositoryViewRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post security privileges repository view params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSecurityPrivilegesRepositoryViewParams) WithDefaults() *PostSecurityPrivilegesRepositoryViewParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post security privileges repository view params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSecurityPrivilegesRepositoryViewParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post security privileges repository view params
func (o *PostSecurityPrivilegesRepositoryViewParams) WithTimeout(timeout time.Duration) *PostSecurityPrivilegesRepositoryViewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post security privileges repository view params
func (o *PostSecurityPrivilegesRepositoryViewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post security privileges repository view params
func (o *PostSecurityPrivilegesRepositoryViewParams) WithContext(ctx context.Context) *PostSecurityPrivilegesRepositoryViewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post security privileges repository view params
func (o *PostSecurityPrivilegesRepositoryViewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post security privileges repository view params
func (o *PostSecurityPrivilegesRepositoryViewParams) WithHTTPClient(client *http.Client) *PostSecurityPrivilegesRepositoryViewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post security privileges repository view params
func (o *PostSecurityPrivilegesRepositoryViewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post security privileges repository view params
func (o *PostSecurityPrivilegesRepositoryViewParams) WithBody(body *models.APIPrivilegeRepositoryViewRequest) *PostSecurityPrivilegesRepositoryViewParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post security privileges repository view params
func (o *PostSecurityPrivilegesRepositoryViewParams) SetBody(body *models.APIPrivilegeRepositoryViewRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostSecurityPrivilegesRepositoryViewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
