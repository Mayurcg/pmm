// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AdminUpdateUserPasswordForm admin update user password form
//
// swagger:model AdminUpdateUserPasswordForm
type AdminUpdateUserPasswordForm struct {

	// password
	Password string `json:"password,omitempty"`
}

// Validate validates this admin update user password form
func (m *AdminUpdateUserPasswordForm) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this admin update user password form based on context it is used
func (m *AdminUpdateUserPasswordForm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AdminUpdateUserPasswordForm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdminUpdateUserPasswordForm) UnmarshalBinary(b []byte) error {
	var res AdminUpdateUserPasswordForm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
