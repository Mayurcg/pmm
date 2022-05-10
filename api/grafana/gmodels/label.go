// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Label Label is a key/value pair of strings.
//
// swagger:model Label
type Label struct {

	// value
	Value string `json:"Name,omitempty"`
}

// Validate validates this label
func (m *Label) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this label based on context it is used
func (m *Label) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Label) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Label) UnmarshalBinary(b []byte) error {
	var res Label
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
