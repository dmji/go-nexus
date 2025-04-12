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

// NewGetBlobstoresFileByNameParams creates a new GetBlobstoresFileByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBlobstoresFileByNameParams() *GetBlobstoresFileByNameParams {
	return &GetBlobstoresFileByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBlobstoresFileByNameParamsWithTimeout creates a new GetBlobstoresFileByNameParams object
// with the ability to set a timeout on a request.
func NewGetBlobstoresFileByNameParamsWithTimeout(timeout time.Duration) *GetBlobstoresFileByNameParams {
	return &GetBlobstoresFileByNameParams{
		timeout: timeout,
	}
}

// NewGetBlobstoresFileByNameParamsWithContext creates a new GetBlobstoresFileByNameParams object
// with the ability to set a context for a request.
func NewGetBlobstoresFileByNameParamsWithContext(ctx context.Context) *GetBlobstoresFileByNameParams {
	return &GetBlobstoresFileByNameParams{
		Context: ctx,
	}
}

// NewGetBlobstoresFileByNameParamsWithHTTPClient creates a new GetBlobstoresFileByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBlobstoresFileByNameParamsWithHTTPClient(client *http.Client) *GetBlobstoresFileByNameParams {
	return &GetBlobstoresFileByNameParams{
		HTTPClient: client,
	}
}

/*
GetBlobstoresFileByNameParams contains all the parameters to send to the API endpoint

	for the get blobstores file by name operation.

	Typically these are written to a http.Request.
*/
type GetBlobstoresFileByNameParams struct {

	/* Name.

	   The name of the file blob store to read
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get blobstores file by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBlobstoresFileByNameParams) WithDefaults() *GetBlobstoresFileByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get blobstores file by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBlobstoresFileByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get blobstores file by name params
func (o *GetBlobstoresFileByNameParams) WithTimeout(timeout time.Duration) *GetBlobstoresFileByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get blobstores file by name params
func (o *GetBlobstoresFileByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get blobstores file by name params
func (o *GetBlobstoresFileByNameParams) WithContext(ctx context.Context) *GetBlobstoresFileByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get blobstores file by name params
func (o *GetBlobstoresFileByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get blobstores file by name params
func (o *GetBlobstoresFileByNameParams) WithHTTPClient(client *http.Client) *GetBlobstoresFileByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get blobstores file by name params
func (o *GetBlobstoresFileByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get blobstores file by name params
func (o *GetBlobstoresFileByNameParams) WithName(name string) *GetBlobstoresFileByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get blobstores file by name params
func (o *GetBlobstoresFileByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetBlobstoresFileByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
