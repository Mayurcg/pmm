// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MuteTimeInterval MuteTimeInterval represents a named set of time intervals for which a route should be muted.
//
// swagger:model MuteTimeInterval
type MuteTimeInterval struct {

	// name
	Name string `json:"name,omitempty"`

	// time intervals
	TimeIntervals []*TimeInterval `json:"time_intervals"`
}

// Validate validates this mute time interval
func (m *MuteTimeInterval) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimeIntervals(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MuteTimeInterval) validateTimeIntervals(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeIntervals) { // not required
		return nil
	}

	for i := 0; i < len(m.TimeIntervals); i++ {
		if swag.IsZero(m.TimeIntervals[i]) { // not required
			continue
		}

		if m.TimeIntervals[i] != nil {
			if err := m.TimeIntervals[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("time_intervals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("time_intervals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this mute time interval based on the context it is used
func (m *MuteTimeInterval) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTimeIntervals(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MuteTimeInterval) contextValidateTimeIntervals(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TimeIntervals); i++ {

		if m.TimeIntervals[i] != nil {
			if err := m.TimeIntervals[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("time_intervals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("time_intervals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MuteTimeInterval) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MuteTimeInterval) UnmarshalBinary(b []byte) error {
	var res MuteTimeInterval
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
