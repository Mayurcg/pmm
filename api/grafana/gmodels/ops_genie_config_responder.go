// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpsGenieConfigResponder ops genie config responder
//
// swagger:model OpsGenieConfigResponder
type OpsGenieConfigResponder struct {

	// One of those 3 should be filled.
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// team, user, escalation, schedule etc.
	Type string `json:"type,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this ops genie config responder
func (m *OpsGenieConfigResponder) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ops genie config responder based on context it is used
func (m *OpsGenieConfigResponder) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpsGenieConfigResponder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpsGenieConfigResponder) UnmarshalBinary(b []byte) error {
	var res OpsGenieConfigResponder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
