// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APICreateUser Api create user
//
// swagger:model ApiCreateUser
type APICreateUser struct {

	// The email address associated with the user.
	EmailAddress string `json:"emailAddress,omitempty"`

	// The first name of the user.
	FirstName string `json:"firstName,omitempty"`

	// The last name of the user.
	LastName string `json:"lastName,omitempty"`

	// The password for the new user.
	Password string `json:"password,omitempty"`

	// The roles which the user has been assigned within Nexus.
	// Unique: true
	Roles []string `json:"roles"`

	// The user's status, e.g. active or disabled.
	// Required: true
	// Enum: ["active","locked","disabled","changepassword"]
	Status *string `json:"status"`

	// The userid which is required for login. This value cannot be changed.
	UserID string `json:"userId,omitempty"`
}

// Validate validates this Api create user
func (m *APICreateUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APICreateUser) validateRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	if err := validate.UniqueItems("roles", "body", m.Roles); err != nil {
		return err
	}

	return nil
}

var apiCreateUserTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","locked","disabled","changepassword"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apiCreateUserTypeStatusPropEnum = append(apiCreateUserTypeStatusPropEnum, v)
	}
}

const (

	// APICreateUserStatusActive captures enum value "active"
	APICreateUserStatusActive string = "active"

	// APICreateUserStatusLocked captures enum value "locked"
	APICreateUserStatusLocked string = "locked"

	// APICreateUserStatusDisabled captures enum value "disabled"
	APICreateUserStatusDisabled string = "disabled"

	// APICreateUserStatusChangepassword captures enum value "changepassword"
	APICreateUserStatusChangepassword string = "changepassword"
)

// prop value enum
func (m *APICreateUser) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, apiCreateUserTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *APICreateUser) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Api create user based on context it is used
func (m *APICreateUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APICreateUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APICreateUser) UnmarshalBinary(b []byte) error {
	var res APICreateUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
