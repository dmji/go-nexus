// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CargoAttributes cargo attributes
//
// swagger:model CargoAttributes
type CargoAttributes struct {

	// Indicates if this repository requires authentication overriding anonymous access.
	RequireAuthentication bool `json:"requireAuthentication,omitempty"`
}

// Validate validates this cargo attributes
func (m *CargoAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cargo attributes based on context it is used
func (m *CargoAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CargoAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CargoAttributes) UnmarshalBinary(b []byte) error {
	var res CargoAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
