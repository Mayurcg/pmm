// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Alert alert
//
// swagger:model Alert
type Alert struct {

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"Created,omitempty"`

	// dashboard Id
	DashboardID int64 `json:"DashboardId,omitempty"`

	// eval data
	EvalData JSON `json:"EvalData,omitempty"`

	// execution error
	ExecutionError string `json:"ExecutionError,omitempty"`

	// for
	For Duration `json:"For,omitempty"`

	// frequency
	Frequency int64 `json:"Frequency,omitempty"`

	// handler
	Handler int64 `json:"Handler,omitempty"`

	// Id
	ID int64 `json:"Id,omitempty"`

	// message
	Message string `json:"Message,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// new state date
	// Format: date-time
	NewStateDate strfmt.DateTime `json:"NewStateDate,omitempty"`

	// org Id
	OrgID int64 `json:"OrgId,omitempty"`

	// panel Id
	PanelID int64 `json:"PanelId,omitempty"`

	// settings
	Settings JSON `json:"Settings,omitempty"`

	// severity
	Severity string `json:"Severity,omitempty"`

	// silenced
	Silenced bool `json:"Silenced,omitempty"`

	// state
	State AlertStateType `json:"State,omitempty"`

	// state changes
	StateChanges int64 `json:"StateChanges,omitempty"`

	// updated
	// Format: date-time
	Updated strfmt.DateTime `json:"Updated,omitempty"`

	// version
	Version int64 `json:"Version,omitempty"`
}

// Validate validates this alert
func (m *Alert) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewStateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Alert) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("Created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Alert) validateFor(formats strfmt.Registry) error {
	if swag.IsZero(m.For) { // not required
		return nil
	}

	if err := m.For.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("For")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("For")
		}
		return err
	}

	return nil
}

func (m *Alert) validateNewStateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.NewStateDate) { // not required
		return nil
	}

	if err := validate.FormatOf("NewStateDate", "body", "date-time", m.NewStateDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Alert) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("State")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("State")
		}
		return err
	}

	return nil
}

func (m *Alert) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("Updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this alert based on the context it is used
func (m *Alert) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Alert) contextValidateFor(ctx context.Context, formats strfmt.Registry) error {

	if err := m.For.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("For")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("For")
		}
		return err
	}

	return nil
}

func (m *Alert) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := m.State.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("State")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("State")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Alert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Alert) UnmarshalBinary(b []byte) error {
	var res Alert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
