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

// GitLfsHostedRepositoryAPIRequest git lfs hosted repository Api request
//
// swagger:model GitLfsHostedRepositoryApiRequest
type GitLfsHostedRepositoryAPIRequest struct {

	// cleanup
	Cleanup *CleanupPolicyAttributes `json:"cleanup,omitempty"`

	// component
	Component *ComponentAttributes `json:"component,omitempty"`

	// A unique identifier for this repository
	// Example: internal
	// Required: true
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name *string `json:"name"`

	// Whether this repository accepts incoming requests
	// Example: true
	// Required: true
	Online *bool `json:"online"`

	// storage
	// Required: true
	Storage *HostedStorageAttributes `json:"storage"`
}

// Validate validates this git lfs hosted repository Api request
func (m *GitLfsHostedRepositoryAPIRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCleanup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitLfsHostedRepositoryAPIRequest) validateCleanup(formats strfmt.Registry) error {
	if swag.IsZero(m.Cleanup) { // not required
		return nil
	}

	if m.Cleanup != nil {
		if err := m.Cleanup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cleanup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cleanup")
			}
			return err
		}
	}

	return nil
}

func (m *GitLfsHostedRepositoryAPIRequest) validateComponent(formats strfmt.Registry) error {
	if swag.IsZero(m.Component) { // not required
		return nil
	}

	if m.Component != nil {
		if err := m.Component.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("component")
			}
			return err
		}
	}

	return nil
}

func (m *GitLfsHostedRepositoryAPIRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", *m.Name, `^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$`); err != nil {
		return err
	}

	return nil
}

func (m *GitLfsHostedRepositoryAPIRequest) validateOnline(formats strfmt.Registry) error {

	if err := validate.Required("online", "body", m.Online); err != nil {
		return err
	}

	return nil
}

func (m *GitLfsHostedRepositoryAPIRequest) validateStorage(formats strfmt.Registry) error {

	if err := validate.Required("storage", "body", m.Storage); err != nil {
		return err
	}

	if m.Storage != nil {
		if err := m.Storage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this git lfs hosted repository Api request based on the context it is used
func (m *GitLfsHostedRepositoryAPIRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCleanup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComponent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitLfsHostedRepositoryAPIRequest) contextValidateCleanup(ctx context.Context, formats strfmt.Registry) error {

	if m.Cleanup != nil {

		if swag.IsZero(m.Cleanup) { // not required
			return nil
		}

		if err := m.Cleanup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cleanup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cleanup")
			}
			return err
		}
	}

	return nil
}

func (m *GitLfsHostedRepositoryAPIRequest) contextValidateComponent(ctx context.Context, formats strfmt.Registry) error {

	if m.Component != nil {

		if swag.IsZero(m.Component) { // not required
			return nil
		}

		if err := m.Component.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("component")
			}
			return err
		}
	}

	return nil
}

func (m *GitLfsHostedRepositoryAPIRequest) contextValidateStorage(ctx context.Context, formats strfmt.Registry) error {

	if m.Storage != nil {

		if err := m.Storage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GitLfsHostedRepositoryAPIRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitLfsHostedRepositoryAPIRequest) UnmarshalBinary(b []byte) error {
	var res GitLfsHostedRepositoryAPIRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
