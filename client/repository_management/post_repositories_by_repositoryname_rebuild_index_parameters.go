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
)

// NewPostRepositoriesByRepositorynameRebuildIndexParams creates a new PostRepositoriesByRepositorynameRebuildIndexParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostRepositoriesByRepositorynameRebuildIndexParams() *PostRepositoriesByRepositorynameRebuildIndexParams {
	return &PostRepositoriesByRepositorynameRebuildIndexParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostRepositoriesByRepositorynameRebuildIndexParamsWithTimeout creates a new PostRepositoriesByRepositorynameRebuildIndexParams object
// with the ability to set a timeout on a request.
func NewPostRepositoriesByRepositorynameRebuildIndexParamsWithTimeout(timeout time.Duration) *PostRepositoriesByRepositorynameRebuildIndexParams {
	return &PostRepositoriesByRepositorynameRebuildIndexParams{
		timeout: timeout,
	}
}

// NewPostRepositoriesByRepositorynameRebuildIndexParamsWithContext creates a new PostRepositoriesByRepositorynameRebuildIndexParams object
// with the ability to set a context for a request.
func NewPostRepositoriesByRepositorynameRebuildIndexParamsWithContext(ctx context.Context) *PostRepositoriesByRepositorynameRebuildIndexParams {
	return &PostRepositoriesByRepositorynameRebuildIndexParams{
		Context: ctx,
	}
}

// NewPostRepositoriesByRepositorynameRebuildIndexParamsWithHTTPClient creates a new PostRepositoriesByRepositorynameRebuildIndexParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostRepositoriesByRepositorynameRebuildIndexParamsWithHTTPClient(client *http.Client) *PostRepositoriesByRepositorynameRebuildIndexParams {
	return &PostRepositoriesByRepositorynameRebuildIndexParams{
		HTTPClient: client,
	}
}

/*
PostRepositoriesByRepositorynameRebuildIndexParams contains all the parameters to send to the API endpoint

	for the post repositories by repositoryname rebuild index operation.

	Typically these are written to a http.Request.
*/
type PostRepositoriesByRepositorynameRebuildIndexParams struct {

	/* RepositoryName.

	   Name of the repository to rebuild index
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post repositories by repositoryname rebuild index params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRepositoriesByRepositorynameRebuildIndexParams) WithDefaults() *PostRepositoriesByRepositorynameRebuildIndexParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post repositories by repositoryname rebuild index params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRepositoriesByRepositorynameRebuildIndexParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post repositories by repositoryname rebuild index params
func (o *PostRepositoriesByRepositorynameRebuildIndexParams) WithTimeout(timeout time.Duration) *PostRepositoriesByRepositorynameRebuildIndexParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post repositories by repositoryname rebuild index params
func (o *PostRepositoriesByRepositorynameRebuildIndexParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post repositories by repositoryname rebuild index params
func (o *PostRepositoriesByRepositorynameRebuildIndexParams) WithContext(ctx context.Context) *PostRepositoriesByRepositorynameRebuildIndexParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post repositories by repositoryname rebuild index params
func (o *PostRepositoriesByRepositorynameRebuildIndexParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post repositories by repositoryname rebuild index params
func (o *PostRepositoriesByRepositorynameRebuildIndexParams) WithHTTPClient(client *http.Client) *PostRepositoriesByRepositorynameRebuildIndexParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post repositories by repositoryname rebuild index params
func (o *PostRepositoriesByRepositorynameRebuildIndexParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepositoryName adds the repositoryName to the post repositories by repositoryname rebuild index params
func (o *PostRepositoriesByRepositorynameRebuildIndexParams) WithRepositoryName(repositoryName string) *PostRepositoriesByRepositorynameRebuildIndexParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the post repositories by repositoryname rebuild index params
func (o *PostRepositoriesByRepositorynameRebuildIndexParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *PostRepositoriesByRepositorynameRebuildIndexParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param repositoryName
	if err := r.SetPathParam("repositoryName", o.RepositoryName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
