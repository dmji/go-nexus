// Code generated by go-swagger; DO NOT EDIT.

package security_management_users

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

// NewDeleteSecurityUsersByUserIDParams creates a new DeleteSecurityUsersByUserIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSecurityUsersByUserIDParams() *DeleteSecurityUsersByUserIDParams {
	return &DeleteSecurityUsersByUserIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSecurityUsersByUserIDParamsWithTimeout creates a new DeleteSecurityUsersByUserIDParams object
// with the ability to set a timeout on a request.
func NewDeleteSecurityUsersByUserIDParamsWithTimeout(timeout time.Duration) *DeleteSecurityUsersByUserIDParams {
	return &DeleteSecurityUsersByUserIDParams{
		timeout: timeout,
	}
}

// NewDeleteSecurityUsersByUserIDParamsWithContext creates a new DeleteSecurityUsersByUserIDParams object
// with the ability to set a context for a request.
func NewDeleteSecurityUsersByUserIDParamsWithContext(ctx context.Context) *DeleteSecurityUsersByUserIDParams {
	return &DeleteSecurityUsersByUserIDParams{
		Context: ctx,
	}
}

// NewDeleteSecurityUsersByUserIDParamsWithHTTPClient creates a new DeleteSecurityUsersByUserIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSecurityUsersByUserIDParamsWithHTTPClient(client *http.Client) *DeleteSecurityUsersByUserIDParams {
	return &DeleteSecurityUsersByUserIDParams{
		HTTPClient: client,
	}
}

/*
DeleteSecurityUsersByUserIDParams contains all the parameters to send to the API endpoint

	for the delete security users by user Id operation.

	Typically these are written to a http.Request.
*/
type DeleteSecurityUsersByUserIDParams struct {

	/* Realm.

	   The realm the request should apply to.
	*/
	Realm *string

	/* UserID.

	   The userid the request should apply to.
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete security users by user Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSecurityUsersByUserIDParams) WithDefaults() *DeleteSecurityUsersByUserIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete security users by user Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSecurityUsersByUserIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete security users by user Id params
func (o *DeleteSecurityUsersByUserIDParams) WithTimeout(timeout time.Duration) *DeleteSecurityUsersByUserIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete security users by user Id params
func (o *DeleteSecurityUsersByUserIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete security users by user Id params
func (o *DeleteSecurityUsersByUserIDParams) WithContext(ctx context.Context) *DeleteSecurityUsersByUserIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete security users by user Id params
func (o *DeleteSecurityUsersByUserIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete security users by user Id params
func (o *DeleteSecurityUsersByUserIDParams) WithHTTPClient(client *http.Client) *DeleteSecurityUsersByUserIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete security users by user Id params
func (o *DeleteSecurityUsersByUserIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRealm adds the realm to the delete security users by user Id params
func (o *DeleteSecurityUsersByUserIDParams) WithRealm(realm *string) *DeleteSecurityUsersByUserIDParams {
	o.SetRealm(realm)
	return o
}

// SetRealm adds the realm to the delete security users by user Id params
func (o *DeleteSecurityUsersByUserIDParams) SetRealm(realm *string) {
	o.Realm = realm
}

// WithUserID adds the userID to the delete security users by user Id params
func (o *DeleteSecurityUsersByUserIDParams) WithUserID(userID string) *DeleteSecurityUsersByUserIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete security users by user Id params
func (o *DeleteSecurityUsersByUserIDParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSecurityUsersByUserIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Realm != nil {

		// query param realm
		var qrRealm string

		if o.Realm != nil {
			qrRealm = *o.Realm
		}
		qRealm := qrRealm
		if qRealm != "" {

			if err := r.SetQueryParam("realm", qRealm); err != nil {
				return err
			}
		}
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
