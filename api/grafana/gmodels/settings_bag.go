// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// SettingsBag settings bag
//
// swagger:model SettingsBag
type SettingsBag map[string]map[string]string

// Validate validates this settings bag
func (m SettingsBag) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this settings bag based on context it is used
func (m SettingsBag) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
