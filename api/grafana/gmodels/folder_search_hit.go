// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FolderSearchHit folder search hit
//
// swagger:model FolderSearchHit
type FolderSearchHit struct {

	// Id
	ID int64 `json:"id,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// Uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this folder search hit
func (m *FolderSearchHit) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this folder search hit based on context it is used
func (m *FolderSearchHit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FolderSearchHit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FolderSearchHit) UnmarshalBinary(b []byte) error {
	var res FolderSearchHit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
