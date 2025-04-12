// Code generated by go-swagger; DO NOT EDIT.

package blob_store

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

// NewDeleteBlobstoresByNameParams creates a new DeleteBlobstoresByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBlobstoresByNameParams() *DeleteBlobstoresByNameParams {
	return &DeleteBlobstoresByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBlobstoresByNameParamsWithTimeout creates a new DeleteBlobstoresByNameParams object
// with the ability to set a timeout on a request.
func NewDeleteBlobstoresByNameParamsWithTimeout(timeout time.Duration) *DeleteBlobstoresByNameParams {
	return &DeleteBlobstoresByNameParams{
		timeout: timeout,
	}
}

// NewDeleteBlobstoresByNameParamsWithContext creates a new DeleteBlobstoresByNameParams object
// with the ability to set a context for a request.
func NewDeleteBlobstoresByNameParamsWithContext(ctx context.Context) *DeleteBlobstoresByNameParams {
	return &DeleteBlobstoresByNameParams{
		Context: ctx,
	}
}

// NewDeleteBlobstoresByNameParamsWithHTTPClient creates a new DeleteBlobstoresByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBlobstoresByNameParamsWithHTTPClient(client *http.Client) *DeleteBlobstoresByNameParams {
	return &DeleteBlobstoresByNameParams{
		HTTPClient: client,
	}
}

/*
DeleteBlobstoresByNameParams contains all the parameters to send to the API endpoint

	for the delete blobstores by name operation.

	Typically these are written to a http.Request.
*/
type DeleteBlobstoresByNameParams struct {

	/* Name.

	   The name of the blob store to delete
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete blobstores by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBlobstoresByNameParams) WithDefaults() *DeleteBlobstoresByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete blobstores by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBlobstoresByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete blobstores by name params
func (o *DeleteBlobstoresByNameParams) WithTimeout(timeout time.Duration) *DeleteBlobstoresByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete blobstores by name params
func (o *DeleteBlobstoresByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete blobstores by name params
func (o *DeleteBlobstoresByNameParams) WithContext(ctx context.Context) *DeleteBlobstoresByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete blobstores by name params
func (o *DeleteBlobstoresByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete blobstores by name params
func (o *DeleteBlobstoresByNameParams) WithHTTPClient(client *http.Client) *DeleteBlobstoresByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete blobstores by name params
func (o *DeleteBlobstoresByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete blobstores by name params
func (o *DeleteBlobstoresByNameParams) WithName(name string) *DeleteBlobstoresByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete blobstores by name params
func (o *DeleteBlobstoresByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBlobstoresByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
