// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MaliciousRiskSummaryResponse malicious risk summary response
//
// swagger:model MaliciousRiskSummaryResponse
type MaliciousRiskSummaryResponse struct {

	// count by ecosystem
	CountByEcosystem []*MaliciousRiskCountByEcosystem `json:"countByEcosystem"`

	// hds error
	HdsError bool `json:"hdsError,omitempty"`

	// quarantine enabled repository count
	QuarantineEnabledRepositoryCount int64 `json:"quarantineEnabledRepositoryCount,omitempty"`

	// total malicious risk count
	TotalMaliciousRiskCount int64 `json:"totalMaliciousRiskCount,omitempty"`

	// total proxy repository count
	TotalProxyRepositoryCount int64 `json:"totalProxyRepositoryCount,omitempty"`
}

// Validate validates this malicious risk summary response
func (m *MaliciousRiskSummaryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountByEcosystem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MaliciousRiskSummaryResponse) validateCountByEcosystem(formats strfmt.Registry) error {
	if swag.IsZero(m.CountByEcosystem) { // not required
		return nil
	}

	for i := 0; i < len(m.CountByEcosystem); i++ {
		if swag.IsZero(m.CountByEcosystem[i]) { // not required
			continue
		}

		if m.CountByEcosystem[i] != nil {
			if err := m.CountByEcosystem[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("countByEcosystem" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("countByEcosystem" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this malicious risk summary response based on the context it is used
func (m *MaliciousRiskSummaryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCountByEcosystem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MaliciousRiskSummaryResponse) contextValidateCountByEcosystem(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CountByEcosystem); i++ {

		if m.CountByEcosystem[i] != nil {

			if swag.IsZero(m.CountByEcosystem[i]) { // not required
				return nil
			}

			if err := m.CountByEcosystem[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("countByEcosystem" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("countByEcosystem" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MaliciousRiskSummaryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MaliciousRiskSummaryResponse) UnmarshalBinary(b []byte) error {
	var res MaliciousRiskSummaryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
