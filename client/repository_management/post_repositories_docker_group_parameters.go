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

// NewPostRepositoriesDockerGroupParams creates a new PostRepositoriesDockerGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostRepositoriesDockerGroupParams() *PostRepositoriesDockerGroupParams {
	return &PostRepositoriesDockerGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostRepositoriesDockerGroupParamsWithTimeout creates a new PostRepositoriesDockerGroupParams object
// with the ability to set a timeout on a request.
func NewPostRepositoriesDockerGroupParamsWithTimeout(timeout time.Duration) *PostRepositoriesDockerGroupParams {
	return &PostRepositoriesDockerGroupParams{
		timeout: timeout,
	}
}

// NewPostRepositoriesDockerGroupParamsWithContext creates a new PostRepositoriesDockerGroupParams object
// with the ability to set a context for a request.
func NewPostRepositoriesDockerGroupParamsWithContext(ctx context.Context) *PostRepositoriesDockerGroupParams {
	return &PostRepositoriesDockerGroupParams{
		Context: ctx,
	}
}

// NewPostRepositoriesDockerGroupParamsWithHTTPClient creates a new PostRepositoriesDockerGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostRepositoriesDockerGroupParamsWithHTTPClient(client *http.Client) *PostRepositoriesDockerGroupParams {
	return &PostRepositoriesDockerGroupParams{
		HTTPClient: client,
	}
}

/*
PostRepositoriesDockerGroupParams contains all the parameters to send to the API endpoint

	for the post repositories docker group operation.

	Typically these are written to a http.Request.
*/
type PostRepositoriesDockerGroupParams struct {

	// Body.
	Body *models.DockerGroupRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post repositories docker group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRepositoriesDockerGroupParams) WithDefaults() *PostRepositoriesDockerGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post repositories docker group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRepositoriesDockerGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post repositories docker group params
func (o *PostRepositoriesDockerGroupParams) WithTimeout(timeout time.Duration) *PostRepositoriesDockerGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post repositories docker group params
func (o *PostRepositoriesDockerGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post repositories docker group params
func (o *PostRepositoriesDockerGroupParams) WithContext(ctx context.Context) *PostRepositoriesDockerGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post repositories docker group params
func (o *PostRepositoriesDockerGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post repositories docker group params
func (o *PostRepositoriesDockerGroupParams) WithHTTPClient(client *http.Client) *PostRepositoriesDockerGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post repositories docker group params
func (o *PostRepositoriesDockerGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post repositories docker group params
func (o *PostRepositoriesDockerGroupParams) WithBody(body *models.DockerGroupRepositoryAPIRequest) *PostRepositoriesDockerGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post repositories docker group params
func (o *PostRepositoriesDockerGroupParams) SetBody(body *models.DockerGroupRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRepositoriesDockerGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
