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

// Token token
//
// swagger:model Token
type Token struct {

	// account
	Account string `json:"account,omitempty"`

	// company
	Company string `json:"company,omitempty"`

	// details Url
	DetailsURL string `json:"details_url,omitempty"`

	// expires
	Expires int64 `json:"exp,omitempty"`

	// Id
	ID string `json:"jti,omitempty"`

	// included admins
	IncludedAdmins int64 `json:"included_admins,omitempty"`

	// included users
	IncludedUsers int64 `json:"included_users,omitempty"`

	// included viewers
	IncludedViewers int64 `json:"included_viewers,omitempty"`

	// issued
	Issued int64 `json:"iat,omitempty"`

	// issuer
	Issuer string `json:"iss,omitempty"`

	// license expires
	LicenseExpires int64 `json:"lexp,omitempty"`

	// license expires warn days
	LicenseExpiresWarnDays int64 `json:"lic_exp_warn_days,omitempty"`

	// license Id
	LicenseID string `json:"lid,omitempty"`

	// license issued
	LicenseIssued int64 `json:"nbf,omitempty"`

	// limit by
	LimitBy string `json:"limit_by,omitempty"`

	// max concurrent user sessions
	MaxConcurrentUserSessions int64 `json:"max_concurrent_user_sessions,omitempty"`

	// products
	Products []string `json:"prod"`

	// slug
	Slug string `json:"slug,omitempty"`

	// subject
	Subject string `json:"sub,omitempty"`

	// token expires warn days
	TokenExpiresWarnDays int64 `json:"tok_exp_warn_days,omitempty"`

	// trial
	Trial bool `json:"trial,omitempty"`

	// trial expires
	TrialExpires int64 `json:"trial_exp,omitempty"`

	// update days
	UpdateDays int64 `json:"update_days,omitempty"`

	// usage billing
	UsageBilling bool `json:"usage_billing,omitempty"`

	// status
	Status TokenStatus `json:"status,omitempty"`
}

// Validate validates this token
func (m *Token) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Token) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// ContextValidate validate this token based on the context it is used
func (m *Token) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Token) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Token) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Token) UnmarshalBinary(b []byte) error {
	var res Token
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
