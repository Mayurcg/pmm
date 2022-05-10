// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PagerdutyImage PagerdutyImage is an image
//
// swagger:model PagerdutyImage
type PagerdutyImage struct {

	// alt
	Alt string `json:"alt,omitempty"`

	// href
	Href string `json:"href,omitempty"`

	// src
	Src string `json:"src,omitempty"`
}

// Validate validates this pagerduty image
func (m *PagerdutyImage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pagerduty image based on context it is used
func (m *PagerdutyImage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PagerdutyImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PagerdutyImage) UnmarshalBinary(b []byte) error {
	var res PagerdutyImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
