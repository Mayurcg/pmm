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

// ErrorResponseBody error response body
//
// swagger:model ErrorResponseBody
type ErrorResponseBody struct {

	// Error An optional detailed description of the actual error. Only included if running in developer mode.
	Error string `json:"error,omitempty"`

	// a human readable version of the error
	// Required: true
	Message *string `json:"message"`

	// Status An optional status to denote the cause of the error.
	//
	// For example, a 412 Precondition Failed error may include additional information of why that error happened.
	Status string `json:"status,omitempty"`
}

// Validate validates this error response body
func (m *ErrorResponseBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ErrorResponseBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this error response body based on context it is used
func (m *ErrorResponseBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorResponseBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorResponseBody) UnmarshalBinary(b []byte) error {
	var res ErrorResponseBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
