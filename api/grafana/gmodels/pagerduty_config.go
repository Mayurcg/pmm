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

// PagerdutyConfig PagerdutyConfig configures notifications via PagerDuty.
//
// swagger:model PagerdutyConfig
type PagerdutyConfig struct {

	// class
	Class string `json:"class,omitempty"`

	// client
	Client string `json:"client,omitempty"`

	// client URL
	ClientURL string `json:"client_url,omitempty"`

	// component
	Component string `json:"component,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// details
	Details map[string]string `json:"details,omitempty"`

	// group
	Group string `json:"group,omitempty"`

	// images
	Images []*PagerdutyImage `json:"images"`

	// links
	Links []*PagerdutyLink `json:"links"`

	// severity
	Severity string `json:"severity,omitempty"`

	// v send resolved
	VSendResolved bool `json:"send_resolved,omitempty"`

	// http config
	HTTPConfig *HTTPClientConfig `json:"http_config,omitempty"`

	// routing key
	RoutingKey Secret `json:"routing_key,omitempty"`

	// service key
	ServiceKey Secret `json:"service_key,omitempty"`

	// url
	URL *URL `json:"url,omitempty"`
}

// Validate validates this pagerduty config
func (m *PagerdutyConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoutingKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PagerdutyConfig) validateImages(formats strfmt.Registry) error {
	if swag.IsZero(m.Images) { // not required
		return nil
	}

	for i := 0; i < len(m.Images); i++ {
		if swag.IsZero(m.Images[i]) { // not required
			continue
		}

		if m.Images[i] != nil {
			if err := m.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PagerdutyConfig) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	for i := 0; i < len(m.Links); i++ {
		if swag.IsZero(m.Links[i]) { // not required
			continue
		}

		if m.Links[i] != nil {
			if err := m.Links[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PagerdutyConfig) validateHTTPConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPConfig) { // not required
		return nil
	}

	if m.HTTPConfig != nil {
		if err := m.HTTPConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("http_config")
			}
			return err
		}
	}

	return nil
}

func (m *PagerdutyConfig) validateRoutingKey(formats strfmt.Registry) error {
	if swag.IsZero(m.RoutingKey) { // not required
		return nil
	}

	if err := m.RoutingKey.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("routing_key")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("routing_key")
		}
		return err
	}

	return nil
}

func (m *PagerdutyConfig) validateServiceKey(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceKey) { // not required
		return nil
	}

	if err := m.ServiceKey.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("service_key")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("service_key")
		}
		return err
	}

	return nil
}

func (m *PagerdutyConfig) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if m.URL != nil {
		if err := m.URL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("url")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this pagerduty config based on the context it is used
func (m *PagerdutyConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHTTPConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoutingKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PagerdutyConfig) contextValidateImages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Images); i++ {

		if m.Images[i] != nil {
			if err := m.Images[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PagerdutyConfig) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Links); i++ {

		if m.Links[i] != nil {
			if err := m.Links[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PagerdutyConfig) contextValidateHTTPConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.HTTPConfig != nil {
		if err := m.HTTPConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("http_config")
			}
			return err
		}
	}

	return nil
}

func (m *PagerdutyConfig) contextValidateRoutingKey(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RoutingKey.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("routing_key")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("routing_key")
		}
		return err
	}

	return nil
}

func (m *PagerdutyConfig) contextValidateServiceKey(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ServiceKey.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("service_key")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("service_key")
		}
		return err
	}

	return nil
}

func (m *PagerdutyConfig) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if m.URL != nil {
		if err := m.URL.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("url")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PagerdutyConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PagerdutyConfig) UnmarshalBinary(b []byte) error {
	var res PagerdutyConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
