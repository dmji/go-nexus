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

// NewPostBlobstoresGroupConvertByNameByNewNameForOriginalParams creates a new PostBlobstoresGroupConvertByNameByNewNameForOriginalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostBlobstoresGroupConvertByNameByNewNameForOriginalParams() *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams {
	return &PostBlobstoresGroupConvertByNameByNewNameForOriginalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostBlobstoresGroupConvertByNameByNewNameForOriginalParamsWithTimeout creates a new PostBlobstoresGroupConvertByNameByNewNameForOriginalParams object
// with the ability to set a timeout on a request.
func NewPostBlobstoresGroupConvertByNameByNewNameForOriginalParamsWithTimeout(timeout time.Duration) *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams {
	return &PostBlobstoresGroupConvertByNameByNewNameForOriginalParams{
		timeout: timeout,
	}
}

// NewPostBlobstoresGroupConvertByNameByNewNameForOriginalParamsWithContext creates a new PostBlobstoresGroupConvertByNameByNewNameForOriginalParams object
// with the ability to set a context for a request.
func NewPostBlobstoresGroupConvertByNameByNewNameForOriginalParamsWithContext(ctx context.Context) *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams {
	return &PostBlobstoresGroupConvertByNameByNewNameForOriginalParams{
		Context: ctx,
	}
}

// NewPostBlobstoresGroupConvertByNameByNewNameForOriginalParamsWithHTTPClient creates a new PostBlobstoresGroupConvertByNameByNewNameForOriginalParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostBlobstoresGroupConvertByNameByNewNameForOriginalParamsWithHTTPClient(client *http.Client) *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams {
	return &PostBlobstoresGroupConvertByNameByNewNameForOriginalParams{
		HTTPClient: client,
	}
}

/*
PostBlobstoresGroupConvertByNameByNewNameForOriginalParams contains all the parameters to send to the API endpoint

	for the post blobstores group convert by name by new name for original operation.

	Typically these are written to a http.Request.
*/
type PostBlobstoresGroupConvertByNameByNewNameForOriginalParams struct {

	/* Name.

	   The name of the group blob store
	*/
	Name string

	/* NewNameForOriginal.

	   A new name to the original blob store
	*/
	NewNameForOriginal string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post blobstores group convert by name by new name for original params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams) WithDefaults() *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post blobstores group convert by name by new name for original params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post blobstores group convert by name by new name for original params
func (o *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams) WithTimeout(timeout time.Duration) *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post blobstores group convert by name by new name for original params
func (o *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post blobstores group convert by name by new name for original params
func (o *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams) WithContext(ctx context.Context) *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post blobstores group convert by name by new name for original params
func (o *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post blobstores group convert by name by new name for original params
func (o *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams) WithHTTPClient(client *http.Client) *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post blobstores group convert by name by new name for original params
func (o *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the post blobstores group convert by name by new name for original params
func (o *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams) WithName(name string) *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the post blobstores group convert by name by new name for original params
func (o *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams) SetName(name string) {
	o.Name = name
}

// WithNewNameForOriginal adds the newNameForOriginal to the post blobstores group convert by name by new name for original params
func (o *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams) WithNewNameForOriginal(newNameForOriginal string) *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams {
	o.SetNewNameForOriginal(newNameForOriginal)
	return o
}

// SetNewNameForOriginal adds the newNameForOriginal to the post blobstores group convert by name by new name for original params
func (o *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams) SetNewNameForOriginal(newNameForOriginal string) {
	o.NewNameForOriginal = newNameForOriginal
}

// WriteToRequest writes these params to a swagger request
func (o *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param newNameForOriginal
	if err := r.SetPathParam("newNameForOriginal", o.NewNameForOriginal); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
