// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NotificationTestCommand notification test command
//
// swagger:model NotificationTestCommand
type NotificationTestCommand struct {

	// disable resolve message
	DisableResolveMessage bool `json:"disableResolveMessage,omitempty"`

	// frequency
	Frequency string `json:"frequency,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// secure settings
	SecureSettings map[string]string `json:"secureSettings,omitempty"`

	// send reminder
	SendReminder bool `json:"sendReminder,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// settings
	Settings JSON `json:"settings,omitempty"`
}

// Validate validates this notification test command
func (m *NotificationTestCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this notification test command based on context it is used
func (m *NotificationTestCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NotificationTestCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotificationTestCommand) UnmarshalBinary(b []byte) error {
	var res NotificationTestCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
