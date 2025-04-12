// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HTTPClientAttributesWithPreemptiveAuth Http client attributes with preemptive auth
//
// swagger:model HttpClientAttributesWithPreemptiveAuth
type HTTPClientAttributesWithPreemptiveAuth struct {

	// authentication
	Authentication *HTTPClientConnectionAuthenticationAttributesWithPreemptive `json:"authentication,omitempty"`

	// Whether to auto-block outbound connections if remote peer is detected as unreachable/unresponsive
	// Example: true
	// Required: true
	AutoBlock *bool `json:"autoBlock"`

	// Whether to block outbound connections on the repository
	// Required: true
	Blocked *bool `json:"blocked"`

	// connection
	Connection *HTTPClientConnectionAttributes `json:"connection,omitempty"`
}

// Validate validates this Http client attributes with preemptive auth
func (m *HTTPClientAttributesWithPreemptiveAuth) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoBlock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlocked(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HTTPClientAttributesWithPreemptiveAuth) validateAuthentication(formats strfmt.Registry) error {
	if swag.IsZero(m.Authentication) { // not required
		return nil
	}

	if m.Authentication != nil {
		if err := m.Authentication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authentication")
			}
			return err
		}
	}

	return nil
}

func (m *HTTPClientAttributesWithPreemptiveAuth) validateAutoBlock(formats strfmt.Registry) error {

	if err := validate.Required("autoBlock", "body", m.AutoBlock); err != nil {
		return err
	}

	return nil
}

func (m *HTTPClientAttributesWithPreemptiveAuth) validateBlocked(formats strfmt.Registry) error {

	if err := validate.Required("blocked", "body", m.Blocked); err != nil {
		return err
	}

	return nil
}

func (m *HTTPClientAttributesWithPreemptiveAuth) validateConnection(formats strfmt.Registry) error {
	if swag.IsZero(m.Connection) { // not required
		return nil
	}

	if m.Connection != nil {
		if err := m.Connection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Http client attributes with preemptive auth based on the context it is used
func (m *HTTPClientAttributesWithPreemptiveAuth) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthentication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HTTPClientAttributesWithPreemptiveAuth) contextValidateAuthentication(ctx context.Context, formats strfmt.Registry) error {

	if m.Authentication != nil {

		if swag.IsZero(m.Authentication) { // not required
			return nil
		}

		if err := m.Authentication.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authentication")
			}
			return err
		}
	}

	return nil
}

func (m *HTTPClientAttributesWithPreemptiveAuth) contextValidateConnection(ctx context.Context, formats strfmt.Registry) error {

	if m.Connection != nil {

		if swag.IsZero(m.Connection) { // not required
			return nil
		}

		if err := m.Connection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HTTPClientAttributesWithPreemptiveAuth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HTTPClientAttributesWithPreemptiveAuth) UnmarshalBinary(b []byte) error {
	var res HTTPClientAttributesWithPreemptiveAuth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
