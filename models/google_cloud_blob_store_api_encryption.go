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

// GoogleCloudBlobStoreAPIEncryption google cloud blob store Api encryption
//
// swagger:model GoogleCloudBlobStoreApiEncryption
type GoogleCloudBlobStoreAPIEncryption struct {

	// CryptoKey ID for KMS encryption.
	EncryptionKey string `json:"encryptionKey,omitempty"`

	// The type of GCP server side encryption to use.
	// Enum: ["default","kmsManagedEncryption"]
	EncryptionType string `json:"encryptionType,omitempty"`
}

// Validate validates this google cloud blob store Api encryption
func (m *GoogleCloudBlobStoreAPIEncryption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEncryptionType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var googleCloudBlobStoreApiEncryptionTypeEncryptionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["default","kmsManagedEncryption"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		googleCloudBlobStoreApiEncryptionTypeEncryptionTypePropEnum = append(googleCloudBlobStoreApiEncryptionTypeEncryptionTypePropEnum, v)
	}
}

const (

	// GoogleCloudBlobStoreAPIEncryptionEncryptionTypeDefault captures enum value "default"
	GoogleCloudBlobStoreAPIEncryptionEncryptionTypeDefault string = "default"

	// GoogleCloudBlobStoreAPIEncryptionEncryptionTypeKmsManagedEncryption captures enum value "kmsManagedEncryption"
	GoogleCloudBlobStoreAPIEncryptionEncryptionTypeKmsManagedEncryption string = "kmsManagedEncryption"
)

// prop value enum
func (m *GoogleCloudBlobStoreAPIEncryption) validateEncryptionTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, googleCloudBlobStoreApiEncryptionTypeEncryptionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GoogleCloudBlobStoreAPIEncryption) validateEncryptionType(formats strfmt.Registry) error {
	if swag.IsZero(m.EncryptionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateEncryptionTypeEnum("encryptionType", "body", m.EncryptionType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this google cloud blob store Api encryption based on context it is used
func (m *GoogleCloudBlobStoreAPIEncryption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GoogleCloudBlobStoreAPIEncryption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GoogleCloudBlobStoreAPIEncryption) UnmarshalBinary(b []byte) error {
	var res GoogleCloudBlobStoreAPIEncryption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
