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

// DockerProxyAPIRepository docker proxy Api repository
//
// swagger:model DockerProxyApiRepository
type DockerProxyAPIRepository struct {

	// cleanup
	Cleanup *CleanupPolicyAttributes `json:"cleanup,omitempty"`

	// docker
	// Required: true
	Docker *DockerAttributes `json:"docker"`

	// docker proxy
	// Required: true
	DockerProxy *DockerProxyAttributes `json:"dockerProxy"`

	// http client
	// Required: true
	HTTPClient *HTTPClientAttributes `json:"httpClient"`

	// A unique identifier for this repository
	// Example: internal
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name,omitempty"`

	// negative cache
	// Required: true
	NegativeCache *NegativeCacheAttributes `json:"negativeCache"`

	// Whether this repository accepts incoming requests
	// Example: true
	// Required: true
	Online *bool `json:"online"`

	// proxy
	// Required: true
	Proxy *ProxyAttributes `json:"proxy"`

	// replication
	Replication *ReplicationAttributes `json:"replication,omitempty"`

	// The name of the routing rule assigned to this repository
	RoutingRuleName string `json:"routingRuleName,omitempty"`

	// storage
	// Required: true
	Storage *StorageAttributes `json:"storage"`
}

// Validate validates this docker proxy Api repository
func (m *DockerProxyAPIRepository) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCleanup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDocker(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDockerProxy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPClient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNegativeCache(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplication(formats); err != nil {
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

func (m *DockerProxyAPIRepository) validateCleanup(formats strfmt.Registry) error {
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

func (m *DockerProxyAPIRepository) validateDocker(formats strfmt.Registry) error {

	if err := validate.Required("docker", "body", m.Docker); err != nil {
		return err
	}

	if m.Docker != nil {
		if err := m.Docker.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("docker")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("docker")
			}
			return err
		}
	}

	return nil
}

func (m *DockerProxyAPIRepository) validateDockerProxy(formats strfmt.Registry) error {

	if err := validate.Required("dockerProxy", "body", m.DockerProxy); err != nil {
		return err
	}

	if m.DockerProxy != nil {
		if err := m.DockerProxy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dockerProxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dockerProxy")
			}
			return err
		}
	}

	return nil
}

func (m *DockerProxyAPIRepository) validateHTTPClient(formats strfmt.Registry) error {

	if err := validate.Required("httpClient", "body", m.HTTPClient); err != nil {
		return err
	}

	if m.HTTPClient != nil {
		if err := m.HTTPClient.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("httpClient")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("httpClient")
			}
			return err
		}
	}

	return nil
}

func (m *DockerProxyAPIRepository) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", m.Name, `^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$`); err != nil {
		return err
	}

	return nil
}

func (m *DockerProxyAPIRepository) validateNegativeCache(formats strfmt.Registry) error {

	if err := validate.Required("negativeCache", "body", m.NegativeCache); err != nil {
		return err
	}

	if m.NegativeCache != nil {
		if err := m.NegativeCache.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("negativeCache")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("negativeCache")
			}
			return err
		}
	}

	return nil
}

func (m *DockerProxyAPIRepository) validateOnline(formats strfmt.Registry) error {

	if err := validate.Required("online", "body", m.Online); err != nil {
		return err
	}

	return nil
}

func (m *DockerProxyAPIRepository) validateProxy(formats strfmt.Registry) error {

	if err := validate.Required("proxy", "body", m.Proxy); err != nil {
		return err
	}

	if m.Proxy != nil {
		if err := m.Proxy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proxy")
			}
			return err
		}
	}

	return nil
}

func (m *DockerProxyAPIRepository) validateReplication(formats strfmt.Registry) error {
	if swag.IsZero(m.Replication) { // not required
		return nil
	}

	if m.Replication != nil {
		if err := m.Replication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("replication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("replication")
			}
			return err
		}
	}

	return nil
}

func (m *DockerProxyAPIRepository) validateStorage(formats strfmt.Registry) error {

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

// ContextValidate validate this docker proxy Api repository based on the context it is used
func (m *DockerProxyAPIRepository) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCleanup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDocker(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDockerProxy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHTTPClient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNegativeCache(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProxy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReplication(ctx, formats); err != nil {
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

func (m *DockerProxyAPIRepository) contextValidateCleanup(ctx context.Context, formats strfmt.Registry) error {

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

func (m *DockerProxyAPIRepository) contextValidateDocker(ctx context.Context, formats strfmt.Registry) error {

	if m.Docker != nil {

		if err := m.Docker.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("docker")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("docker")
			}
			return err
		}
	}

	return nil
}

func (m *DockerProxyAPIRepository) contextValidateDockerProxy(ctx context.Context, formats strfmt.Registry) error {

	if m.DockerProxy != nil {

		if err := m.DockerProxy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dockerProxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dockerProxy")
			}
			return err
		}
	}

	return nil
}

func (m *DockerProxyAPIRepository) contextValidateHTTPClient(ctx context.Context, formats strfmt.Registry) error {

	if m.HTTPClient != nil {

		if err := m.HTTPClient.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("httpClient")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("httpClient")
			}
			return err
		}
	}

	return nil
}

func (m *DockerProxyAPIRepository) contextValidateNegativeCache(ctx context.Context, formats strfmt.Registry) error {

	if m.NegativeCache != nil {

		if err := m.NegativeCache.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("negativeCache")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("negativeCache")
			}
			return err
		}
	}

	return nil
}

func (m *DockerProxyAPIRepository) contextValidateProxy(ctx context.Context, formats strfmt.Registry) error {

	if m.Proxy != nil {

		if err := m.Proxy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proxy")
			}
			return err
		}
	}

	return nil
}

func (m *DockerProxyAPIRepository) contextValidateReplication(ctx context.Context, formats strfmt.Registry) error {

	if m.Replication != nil {

		if swag.IsZero(m.Replication) { // not required
			return nil
		}

		if err := m.Replication.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("replication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("replication")
			}
			return err
		}
	}

	return nil
}

func (m *DockerProxyAPIRepository) contextValidateStorage(ctx context.Context, formats strfmt.Registry) error {

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
func (m *DockerProxyAPIRepository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DockerProxyAPIRepository) UnmarshalBinary(b []byte) error {
	var res DockerProxyAPIRepository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
