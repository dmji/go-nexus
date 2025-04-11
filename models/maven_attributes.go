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

// MavenAttributes maven attributes
//
// swagger:model MavenAttributes
type MavenAttributes struct {

	// Content Disposition
	// Example: ATTACHMENT
	// Pattern: INLINE|ATTACHMENT
	// Enum: ["INLINE","ATTACHMENT"]
	ContentDisposition string `json:"contentDisposition,omitempty"`

	// Validate that all paths are maven artifact or metadata paths
	// Example: STRICT
	// Pattern: STRICT|PERMISSIVE
	// Enum: ["STRICT","PERMISSIVE"]
	LayoutPolicy string `json:"layoutPolicy,omitempty"`

	// What type of artifacts does this repository store?
	// Example: MIXED
	// Pattern: RELEASE|SNAPSHOT|MIXED
	// Enum: ["RELEASE","SNAPSHOT","MIXED"]
	VersionPolicy string `json:"versionPolicy,omitempty"`
}

// Validate validates this maven attributes
func (m *MavenAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentDisposition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLayoutPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var mavenAttributesTypeContentDispositionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INLINE","ATTACHMENT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mavenAttributesTypeContentDispositionPropEnum = append(mavenAttributesTypeContentDispositionPropEnum, v)
	}
}

const (

	// MavenAttributesContentDispositionINLINE captures enum value "INLINE"
	MavenAttributesContentDispositionINLINE string = "INLINE"

	// MavenAttributesContentDispositionATTACHMENT captures enum value "ATTACHMENT"
	MavenAttributesContentDispositionATTACHMENT string = "ATTACHMENT"
)

// prop value enum
func (m *MavenAttributes) validateContentDispositionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mavenAttributesTypeContentDispositionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MavenAttributes) validateContentDisposition(formats strfmt.Registry) error {
	if swag.IsZero(m.ContentDisposition) { // not required
		return nil
	}

	if err := validate.Pattern("contentDisposition", "body", m.ContentDisposition, `INLINE|ATTACHMENT`); err != nil {
		return err
	}

	// value enum
	if err := m.validateContentDispositionEnum("contentDisposition", "body", m.ContentDisposition); err != nil {
		return err
	}

	return nil
}

var mavenAttributesTypeLayoutPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STRICT","PERMISSIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mavenAttributesTypeLayoutPolicyPropEnum = append(mavenAttributesTypeLayoutPolicyPropEnum, v)
	}
}

const (

	// MavenAttributesLayoutPolicySTRICT captures enum value "STRICT"
	MavenAttributesLayoutPolicySTRICT string = "STRICT"

	// MavenAttributesLayoutPolicyPERMISSIVE captures enum value "PERMISSIVE"
	MavenAttributesLayoutPolicyPERMISSIVE string = "PERMISSIVE"
)

// prop value enum
func (m *MavenAttributes) validateLayoutPolicyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mavenAttributesTypeLayoutPolicyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MavenAttributes) validateLayoutPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.LayoutPolicy) { // not required
		return nil
	}

	if err := validate.Pattern("layoutPolicy", "body", m.LayoutPolicy, `STRICT|PERMISSIVE`); err != nil {
		return err
	}

	// value enum
	if err := m.validateLayoutPolicyEnum("layoutPolicy", "body", m.LayoutPolicy); err != nil {
		return err
	}

	return nil
}

var mavenAttributesTypeVersionPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RELEASE","SNAPSHOT","MIXED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mavenAttributesTypeVersionPolicyPropEnum = append(mavenAttributesTypeVersionPolicyPropEnum, v)
	}
}

const (

	// MavenAttributesVersionPolicyRELEASE captures enum value "RELEASE"
	MavenAttributesVersionPolicyRELEASE string = "RELEASE"

	// MavenAttributesVersionPolicySNAPSHOT captures enum value "SNAPSHOT"
	MavenAttributesVersionPolicySNAPSHOT string = "SNAPSHOT"

	// MavenAttributesVersionPolicyMIXED captures enum value "MIXED"
	MavenAttributesVersionPolicyMIXED string = "MIXED"
)

// prop value enum
func (m *MavenAttributes) validateVersionPolicyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mavenAttributesTypeVersionPolicyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MavenAttributes) validateVersionPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.VersionPolicy) { // not required
		return nil
	}

	if err := validate.Pattern("versionPolicy", "body", m.VersionPolicy, `RELEASE|SNAPSHOT|MIXED`); err != nil {
		return err
	}

	// value enum
	if err := m.validateVersionPolicyEnum("versionPolicy", "body", m.VersionPolicy); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this maven attributes based on context it is used
func (m *MavenAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MavenAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MavenAttributes) UnmarshalBinary(b []byte) error {
	var res MavenAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
