// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserQuotaDTO user quota d t o
//
// swagger:model UserQuotaDTO
type UserQuotaDTO struct {

	// limit
	Limit int64 `json:"limit,omitempty"`

	// target
	Target string `json:"target,omitempty"`

	// used
	Used int64 `json:"used,omitempty"`

	// user Id
	UserID int64 `json:"user_id,omitempty"`
}

// Validate validates this user quota d t o
func (m *UserQuotaDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user quota d t o based on context it is used
func (m *UserQuotaDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserQuotaDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserQuotaDTO) UnmarshalBinary(b []byte) error {
	var res UserQuotaDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
