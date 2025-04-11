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

// BlobStoreAPISoftQuota blob store Api soft quota
//
// swagger:model BlobStoreApiSoftQuota
type BlobStoreAPISoftQuota struct {

	// The limit in MB.
	// Minimum: 0
	Limit *int64 `json:"limit,omitempty"`

	// The type to use such as spaceRemainingQuota, or spaceUsedQuota
	// Enum: ["spaceRemainingQuota","spaceUsedQuota"]
	Type string `json:"type,omitempty"`
}

// Validate validates this blob store Api soft quota
func (m *BlobStoreAPISoftQuota) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimit(formats); err != nil {
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

func (m *BlobStoreAPISoftQuota) validateLimit(formats strfmt.Registry) error {
	if swag.IsZero(m.Limit) { // not required
		return nil
	}

	if err := validate.MinimumInt("limit", "body", *m.Limit, 0, false); err != nil {
		return err
	}

	return nil
}

var blobStoreApiSoftQuotaTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["spaceRemainingQuota","spaceUsedQuota"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		blobStoreApiSoftQuotaTypeTypePropEnum = append(blobStoreApiSoftQuotaTypeTypePropEnum, v)
	}
}

const (

	// BlobStoreAPISoftQuotaTypeSpaceRemainingQuota captures enum value "spaceRemainingQuota"
	BlobStoreAPISoftQuotaTypeSpaceRemainingQuota string = "spaceRemainingQuota"

	// BlobStoreAPISoftQuotaTypeSpaceUsedQuota captures enum value "spaceUsedQuota"
	BlobStoreAPISoftQuotaTypeSpaceUsedQuota string = "spaceUsedQuota"
)

// prop value enum
func (m *BlobStoreAPISoftQuota) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, blobStoreApiSoftQuotaTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BlobStoreAPISoftQuota) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this blob store Api soft quota based on context it is used
func (m *BlobStoreAPISoftQuota) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BlobStoreAPISoftQuota) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlobStoreAPISoftQuota) UnmarshalBinary(b []byte) error {
	var res BlobStoreAPISoftQuota
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
