// Code generated by go-swagger; DO NOT EDIT.

package repository_management

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

// NewPostRepositoriesMavenGroupParams creates a new PostRepositoriesMavenGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostRepositoriesMavenGroupParams() *PostRepositoriesMavenGroupParams {
	return &PostRepositoriesMavenGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostRepositoriesMavenGroupParamsWithTimeout creates a new PostRepositoriesMavenGroupParams object
// with the ability to set a timeout on a request.
func NewPostRepositoriesMavenGroupParamsWithTimeout(timeout time.Duration) *PostRepositoriesMavenGroupParams {
	return &PostRepositoriesMavenGroupParams{
		timeout: timeout,
	}
}

// NewPostRepositoriesMavenGroupParamsWithContext creates a new PostRepositoriesMavenGroupParams object
// with the ability to set a context for a request.
func NewPostRepositoriesMavenGroupParamsWithContext(ctx context.Context) *PostRepositoriesMavenGroupParams {
	return &PostRepositoriesMavenGroupParams{
		Context: ctx,
	}
}

// NewPostRepositoriesMavenGroupParamsWithHTTPClient creates a new PostRepositoriesMavenGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostRepositoriesMavenGroupParamsWithHTTPClient(client *http.Client) *PostRepositoriesMavenGroupParams {
	return &PostRepositoriesMavenGroupParams{
		HTTPClient: client,
	}
}

/*
PostRepositoriesMavenGroupParams contains all the parameters to send to the API endpoint

	for the post repositories maven group operation.

	Typically these are written to a http.Request.
*/
type PostRepositoriesMavenGroupParams struct {

	// Body.
	Body *models.MavenGroupRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post repositories maven group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRepositoriesMavenGroupParams) WithDefaults() *PostRepositoriesMavenGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post repositories maven group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRepositoriesMavenGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post repositories maven group params
func (o *PostRepositoriesMavenGroupParams) WithTimeout(timeout time.Duration) *PostRepositoriesMavenGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post repositories maven group params
func (o *PostRepositoriesMavenGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post repositories maven group params
func (o *PostRepositoriesMavenGroupParams) WithContext(ctx context.Context) *PostRepositoriesMavenGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post repositories maven group params
func (o *PostRepositoriesMavenGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post repositories maven group params
func (o *PostRepositoriesMavenGroupParams) WithHTTPClient(client *http.Client) *PostRepositoriesMavenGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post repositories maven group params
func (o *PostRepositoriesMavenGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post repositories maven group params
func (o *PostRepositoriesMavenGroupParams) WithBody(body *models.MavenGroupRepositoryAPIRequest) *PostRepositoriesMavenGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post repositories maven group params
func (o *PostRepositoriesMavenGroupParams) SetBody(body *models.MavenGroupRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRepositoriesMavenGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
