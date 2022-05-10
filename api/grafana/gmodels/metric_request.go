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

// MetricRequest metric request
//
// swagger:model MetricRequest
type MetricRequest struct {

	// debug
	Debug bool `json:"debug,omitempty"`

	// From Start time in epoch timestamps in milliseconds or relative using Grafana time units.
	// Example: now-1h
	// Required: true
	From *string `json:"from"`

	// queries.refId – Specifies an identifier of the query. Is optional and default to “A”.
	// queries.datasourceId – Specifies the data source to be queried. Each query in the request must have an unique datasourceId.
	// queries.maxDataPoints - Species maximum amount of data points that dashboard panel can render. Is optional and default to 100.
	// queries.intervalMs - Specifies the time interval in milliseconds of time series. Is optional and defaults to 1000.
	// Example: [{"datasource":{"uid":"PD8C576611E62080A"},"format":"table","intervalMs":86400000,"maxDataPoints":1092,"rawSql":"SELECT 1 as valueOne, 2 as valueTwo","refId":"A"}]
	// Required: true
	Queries []JSON `json:"queries"`

	// To End time in epoch timestamps in milliseconds or relative using Grafana time units.
	// Example: now
	// Required: true
	To *string `json:"to"`
}

// Validate validates this metric request
func (m *MetricRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricRequest) validateFrom(formats strfmt.Registry) error {

	if err := validate.Required("from", "body", m.From); err != nil {
		return err
	}

	return nil
}

func (m *MetricRequest) validateQueries(formats strfmt.Registry) error {

	if err := validate.Required("queries", "body", m.Queries); err != nil {
		return err
	}

	return nil
}

func (m *MetricRequest) validateTo(formats strfmt.Registry) error {

	if err := validate.Required("to", "body", m.To); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this metric request based on context it is used
func (m *MetricRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MetricRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricRequest) UnmarshalBinary(b []byte) error {
	var res MetricRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
