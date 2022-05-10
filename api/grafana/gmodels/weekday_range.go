// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WeekdayRange A WeekdayRange is an inclusive range between [0, 6] where 0 = Sunday.
//
// swagger:model WeekdayRange
type WeekdayRange struct {

	// begin
	Begin int64 `json:"Begin,omitempty"`

	// end
	End int64 `json:"End,omitempty"`
}

// Validate validates this weekday range
func (m *WeekdayRange) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this weekday range based on context it is used
func (m *WeekdayRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WeekdayRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WeekdayRange) UnmarshalBinary(b []byte) error {
	var res WeekdayRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
