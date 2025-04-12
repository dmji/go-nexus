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

// NewDeleteSecurityUsersByUseridParams creates a new DeleteSecurityUsersByUseridParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSecurityUsersByUseridParams() *DeleteSecurityUsersByUseridParams {
	return &DeleteSecurityUsersByUseridParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSecurityUsersByUseridParamsWithTimeout creates a new DeleteSecurityUsersByUseridParams object
// with the ability to set a timeout on a request.
func NewDeleteSecurityUsersByUseridParamsWithTimeout(timeout time.Duration) *DeleteSecurityUsersByUseridParams {
	return &DeleteSecurityUsersByUseridParams{
		timeout: timeout,
	}
}

// NewDeleteSecurityUsersByUseridParamsWithContext creates a new DeleteSecurityUsersByUseridParams object
// with the ability to set a context for a request.
func NewDeleteSecurityUsersByUseridParamsWithContext(ctx context.Context) *DeleteSecurityUsersByUseridParams {
	return &DeleteSecurityUsersByUseridParams{
		Context: ctx,
	}
}

// NewDeleteSecurityUsersByUseridParamsWithHTTPClient creates a new DeleteSecurityUsersByUseridParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSecurityUsersByUseridParamsWithHTTPClient(client *http.Client) *DeleteSecurityUsersByUseridParams {
	return &DeleteSecurityUsersByUseridParams{
		HTTPClient: client,
	}
}

/*
DeleteSecurityUsersByUseridParams contains all the parameters to send to the API endpoint

	for the delete security users by userid operation.

	Typically these are written to a http.Request.
*/
type DeleteSecurityUsersByUseridParams struct {

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

// WithDefaults hydrates default values in the delete security users by userid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSecurityUsersByUseridParams) WithDefaults() *DeleteSecurityUsersByUseridParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete security users by userid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSecurityUsersByUseridParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete security users by userid params
func (o *DeleteSecurityUsersByUseridParams) WithTimeout(timeout time.Duration) *DeleteSecurityUsersByUseridParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete security users by userid params
func (o *DeleteSecurityUsersByUseridParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete security users by userid params
func (o *DeleteSecurityUsersByUseridParams) WithContext(ctx context.Context) *DeleteSecurityUsersByUseridParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete security users by userid params
func (o *DeleteSecurityUsersByUseridParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete security users by userid params
func (o *DeleteSecurityUsersByUseridParams) WithHTTPClient(client *http.Client) *DeleteSecurityUsersByUseridParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete security users by userid params
func (o *DeleteSecurityUsersByUseridParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRealm adds the realm to the delete security users by userid params
func (o *DeleteSecurityUsersByUseridParams) WithRealm(realm *string) *DeleteSecurityUsersByUseridParams {
	o.SetRealm(realm)
	return o
}

// SetRealm adds the realm to the delete security users by userid params
func (o *DeleteSecurityUsersByUseridParams) SetRealm(realm *string) {
	o.Realm = realm
}

// WithUserID adds the userID to the delete security users by userid params
func (o *DeleteSecurityUsersByUseridParams) WithUserID(userID string) *DeleteSecurityUsersByUseridParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete security users by userid params
func (o *DeleteSecurityUsersByUseridParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSecurityUsersByUseridParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
