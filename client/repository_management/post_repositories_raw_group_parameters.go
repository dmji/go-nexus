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

// NewPostRepositoriesRawGroupParams creates a new PostRepositoriesRawGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostRepositoriesRawGroupParams() *PostRepositoriesRawGroupParams {
	return &PostRepositoriesRawGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostRepositoriesRawGroupParamsWithTimeout creates a new PostRepositoriesRawGroupParams object
// with the ability to set a timeout on a request.
func NewPostRepositoriesRawGroupParamsWithTimeout(timeout time.Duration) *PostRepositoriesRawGroupParams {
	return &PostRepositoriesRawGroupParams{
		timeout: timeout,
	}
}

// NewPostRepositoriesRawGroupParamsWithContext creates a new PostRepositoriesRawGroupParams object
// with the ability to set a context for a request.
func NewPostRepositoriesRawGroupParamsWithContext(ctx context.Context) *PostRepositoriesRawGroupParams {
	return &PostRepositoriesRawGroupParams{
		Context: ctx,
	}
}

// NewPostRepositoriesRawGroupParamsWithHTTPClient creates a new PostRepositoriesRawGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostRepositoriesRawGroupParamsWithHTTPClient(client *http.Client) *PostRepositoriesRawGroupParams {
	return &PostRepositoriesRawGroupParams{
		HTTPClient: client,
	}
}

/*
PostRepositoriesRawGroupParams contains all the parameters to send to the API endpoint

	for the post repositories raw group operation.

	Typically these are written to a http.Request.
*/
type PostRepositoriesRawGroupParams struct {

	// Body.
	Body *models.RawGroupRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post repositories raw group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRepositoriesRawGroupParams) WithDefaults() *PostRepositoriesRawGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post repositories raw group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRepositoriesRawGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post repositories raw group params
func (o *PostRepositoriesRawGroupParams) WithTimeout(timeout time.Duration) *PostRepositoriesRawGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post repositories raw group params
func (o *PostRepositoriesRawGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post repositories raw group params
func (o *PostRepositoriesRawGroupParams) WithContext(ctx context.Context) *PostRepositoriesRawGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post repositories raw group params
func (o *PostRepositoriesRawGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post repositories raw group params
func (o *PostRepositoriesRawGroupParams) WithHTTPClient(client *http.Client) *PostRepositoriesRawGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post repositories raw group params
func (o *PostRepositoriesRawGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post repositories raw group params
func (o *PostRepositoriesRawGroupParams) WithBody(body *models.RawGroupRepositoryAPIRequest) *PostRepositoriesRawGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post repositories raw group params
func (o *PostRepositoriesRawGroupParams) SetBody(body *models.RawGroupRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRepositoriesRawGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
