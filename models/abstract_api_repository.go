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

// AbstractAPIRepository abstract Api repository
//
// swagger:model AbstractApiRepository
type AbstractAPIRepository struct {

	// Component format held in this repository
	// Example: npm
	Format string `json:"format,omitempty"`

	// A unique identifier for this repository
	// Example: internal
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name,omitempty"`

	// Whether this repository accepts incoming requests
	// Example: true
	// Required: true
	Online *bool `json:"online"`

	// Controls if deployments of and updates to artifacts are allowed
	// Example: hosted
	// Enum: ["hosted","proxy","group"]
	Type string `json:"type,omitempty"`

	// URL to the repository
	URL string `json:"url,omitempty"`
}

// Validate validates this abstract Api repository
func (m *AbstractAPIRepository) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AbstractAPIRepository) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", m.Name, `^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$`); err != nil {
		return err
	}

	return nil
}

func (m *AbstractAPIRepository) validateOnline(formats strfmt.Registry) error {

	if err := validate.Required("online", "body", m.Online); err != nil {
		return err
	}

	return nil
}

var abstractApiRepositoryTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["hosted","proxy","group"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		abstractApiRepositoryTypeTypePropEnum = append(abstractApiRepositoryTypeTypePropEnum, v)
	}
}

const (

	// AbstractAPIRepositoryTypeHosted captures enum value "hosted"
	AbstractAPIRepositoryTypeHosted string = "hosted"

	// AbstractAPIRepositoryTypeProxy captures enum value "proxy"
	AbstractAPIRepositoryTypeProxy string = "proxy"

	// AbstractAPIRepositoryTypeGroup captures enum value "group"
	AbstractAPIRepositoryTypeGroup string = "group"
)

// prop value enum
func (m *AbstractAPIRepository) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, abstractApiRepositoryTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AbstractAPIRepository) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this abstract Api repository based on context it is used
func (m *AbstractAPIRepository) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AbstractAPIRepository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AbstractAPIRepository) UnmarshalBinary(b []byte) error {
	var res AbstractAPIRepository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
