// Code generated by go-swagger; DO NOT EDIT.

package gmodels

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

// ClusterStatus ClusterStatus cluster status
//
// swagger:model ClusterStatus
type ClusterStatus struct {

	// name
	Name string `json:"name,omitempty"`

	// peers
	Peers []*PeerStatus `json:"peers"`

	// status
	// Required: true
	// Enum: [[ready settling disabled]]
	Status *string `json:"status"`
}

// Validate validates this cluster status
func (m *ClusterStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePeers(formats); err != nil {
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

func (m *ClusterStatus) validatePeers(formats strfmt.Registry) error {
	if swag.IsZero(m.Peers) { // not required
		return nil
	}

	for i := 0; i < len(m.Peers); i++ {
		if swag.IsZero(m.Peers[i]) { // not required
			continue
		}

		if m.Peers[i] != nil {
			if err := m.Peers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("peers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("peers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var clusterStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["[ready settling disabled]"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterStatusTypeStatusPropEnum = append(clusterStatusTypeStatusPropEnum, v)
	}
}

const (

	// ClusterStatusStatusReadySettlingDisabled captures enum value "[ready settling disabled]"
	ClusterStatusStatusReadySettlingDisabled string = "[ready settling disabled]"
)

// prop value enum
func (m *ClusterStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clusterStatusTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClusterStatus) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cluster status based on the context it is used
func (m *ClusterStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePeers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterStatus) contextValidatePeers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Peers); i++ {

		if m.Peers[i] != nil {
			if err := m.Peers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("peers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("peers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterStatus) UnmarshalBinary(b []byte) error {
	var res ClusterStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
