// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APIPrivilegeRepositoryViewRequest Api privilege repository view request
//
// swagger:model ApiPrivilegeRepositoryViewRequest
type APIPrivilegeRepositoryViewRequest struct {

	// A collection of actions to associate with the privilege, using BREAD syntax (browse,read,edit,add,delete,all) as well as 'run' for script privileges.
	Actions []string `json:"actions"`

	// description
	Description string `json:"description,omitempty"`

	// The repository format (i.e 'nuget', 'npm') this privilege will grant access to (or * for all).
	Format string `json:"format,omitempty"`

	// The name of the privilege.  This value cannot be changed.
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name,omitempty"`

	// The name of the repository this privilege will grant access to (or * for all).
	Repository string `json:"repository,omitempty"`
}

// Validate validates this Api privilege repository view request
func (m *APIPrivilegeRepositoryViewRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var apiPrivilegeRepositoryViewRequestActionsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["READ","BROWSE","EDIT","ADD","DELETE","RUN","START","STOP","ASSOCIATE","DISASSOCIATE","ALL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apiPrivilegeRepositoryViewRequestActionsItemsEnum = append(apiPrivilegeRepositoryViewRequestActionsItemsEnum, v)
	}
}

func (m *APIPrivilegeRepositoryViewRequest) validateActionsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, apiPrivilegeRepositoryViewRequestActionsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *APIPrivilegeRepositoryViewRequest) validateActions(formats strfmt.Registry) error {
	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	for i := 0; i < len(m.Actions); i++ {

		// value enum
		if err := m.validateActionsItemsEnum("actions"+"."+strconv.Itoa(i), "body", m.Actions[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *APIPrivilegeRepositoryViewRequest) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", m.Name, `^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Api privilege repository view request based on context it is used
func (m *APIPrivilegeRepositoryViewRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIPrivilegeRepositoryViewRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIPrivilegeRepositoryViewRequest) UnmarshalBinary(b []byte) error {
	var res APIPrivilegeRepositoryViewRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
