// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReportOptionsDTO report options d t o
//
// swagger:model ReportOptionsDTO
type ReportOptionsDTO struct {

	// layout
	Layout string `json:"layout,omitempty"`

	// orientation
	Orientation string `json:"orientation,omitempty"`

	// time range
	TimeRange *TimeRangeDTO `json:"timeRange,omitempty"`
}

// Validate validates this report options d t o
func (m *ReportOptionsDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimeRange(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportOptionsDTO) validateTimeRange(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeRange) { // not required
		return nil
	}

	if m.TimeRange != nil {
		if err := m.TimeRange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeRange")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timeRange")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this report options d t o based on the context it is used
func (m *ReportOptionsDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTimeRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportOptionsDTO) contextValidateTimeRange(ctx context.Context, formats strfmt.Registry) error {

	if m.TimeRange != nil {
		if err := m.TimeRange.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeRange")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timeRange")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportOptionsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportOptionsDTO) UnmarshalBinary(b []byte) error {
	var res ReportOptionsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
