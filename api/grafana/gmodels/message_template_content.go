// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MessageTemplateContent message template content
//
// swagger:model MessageTemplateContent
type MessageTemplateContent struct {

	// template
	Template string `json:"Template,omitempty"`
}

// Validate validates this message template content
func (m *MessageTemplateContent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this message template content based on context it is used
func (m *MessageTemplateContent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MessageTemplateContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessageTemplateContent) UnmarshalBinary(b []byte) error {
	var res MessageTemplateContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
