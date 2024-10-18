// Code generated by go-swagger; DO NOT EDIT.

package qan_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetLabelsReader is a Reader for the GetLabels structure.
type GetLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLabelsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLabelsOK creates a GetLabelsOK with default headers values
func NewGetLabelsOK() *GetLabelsOK {
	return &GetLabelsOK{}
}

/*
GetLabelsOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetLabelsOK struct {
	Payload *GetLabelsOKBody
}

// IsSuccess returns true when this get labels Ok response has a 2xx status code
func (o *GetLabelsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get labels Ok response has a 3xx status code
func (o *GetLabelsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get labels Ok response has a 4xx status code
func (o *GetLabelsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get labels Ok response has a 5xx status code
func (o *GetLabelsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get labels Ok response a status code equal to that given
func (o *GetLabelsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get labels Ok response
func (o *GetLabelsOK) Code() int {
	return 200
}

func (o *GetLabelsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/qan:getLabels][%d] getLabelsOk %s", 200, payload)
}

func (o *GetLabelsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/qan:getLabels][%d] getLabelsOk %s", 200, payload)
}

func (o *GetLabelsOK) GetPayload() *GetLabelsOKBody {
	return o.Payload
}

func (o *GetLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(GetLabelsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLabelsDefault creates a GetLabelsDefault with default headers values
func NewGetLabelsDefault(code int) *GetLabelsDefault {
	return &GetLabelsDefault{
		_statusCode: code,
	}
}

/*
GetLabelsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetLabelsDefault struct {
	_statusCode int

	Payload *GetLabelsDefaultBody
}

// IsSuccess returns true when this get labels default response has a 2xx status code
func (o *GetLabelsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get labels default response has a 3xx status code
func (o *GetLabelsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get labels default response has a 4xx status code
func (o *GetLabelsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get labels default response has a 5xx status code
func (o *GetLabelsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get labels default response a status code equal to that given
func (o *GetLabelsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get labels default response
func (o *GetLabelsDefault) Code() int {
	return o._statusCode
}

func (o *GetLabelsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/qan:getLabels][%d] GetLabels default %s", o._statusCode, payload)
}

func (o *GetLabelsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/qan:getLabels][%d] GetLabels default %s", o._statusCode, payload)
}

func (o *GetLabelsDefault) GetPayload() *GetLabelsDefaultBody {
	return o.Payload
}

func (o *GetLabelsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(GetLabelsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetLabelsBody GetLabelsRequest defines filtering of object detail's labels for specific value of
// dimension (ex.: host=hostname1 or queryid=1D410B4BE5060972.
swagger:model GetLabelsBody
*/
type GetLabelsBody struct {
	// period start from
	// Format: date-time
	PeriodStartFrom strfmt.DateTime `json:"period_start_from,omitempty"`

	// period start to
	// Format: date-time
	PeriodStartTo strfmt.DateTime `json:"period_start_to,omitempty"`

	// dimension value: ex: queryid - 1D410B4BE5060972.
	FilterBy string `json:"filter_by,omitempty"`

	// one of dimension: queryid | host ...
	GroupBy string `json:"group_by,omitempty"`
}

// Validate validates this get labels body
func (o *GetLabelsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePeriodStartFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePeriodStartTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLabelsBody) validatePeriodStartFrom(formats strfmt.Registry) error {
	if swag.IsZero(o.PeriodStartFrom) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"period_start_from", "body", "date-time", o.PeriodStartFrom.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetLabelsBody) validatePeriodStartTo(formats strfmt.Registry) error {
	if swag.IsZero(o.PeriodStartTo) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"period_start_to", "body", "date-time", o.PeriodStartTo.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get labels body based on context it is used
func (o *GetLabelsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetLabelsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLabelsBody) UnmarshalBinary(b []byte) error {
	var res GetLabelsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetLabelsDefaultBody get labels default body
swagger:model GetLabelsDefaultBody
*/
type GetLabelsDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*GetLabelsDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this get labels default body
func (o *GetLabelsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLabelsDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetLabels default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetLabels default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get labels default body based on the context it is used
func (o *GetLabelsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLabelsDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetLabels default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetLabels default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLabelsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLabelsDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetLabelsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetLabelsDefaultBodyDetailsItems0 get labels default body details items0
swagger:model GetLabelsDefaultBodyDetailsItems0
*/
type GetLabelsDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`

	// get labels default body details items0
	GetLabelsDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *GetLabelsDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {
		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv GetLabelsDefaultBodyDetailsItems0

	rcv.AtType = stage1.AtType
	*o = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "@type")
	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		o.GetLabelsDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o GetLabelsDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
	var stage1 struct {
		// at type
		AtType string `json:"@type,omitempty"`
	}

	stage1.AtType = o.AtType

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(o.GetLabelsDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.GetLabelsDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this get labels default body details items0
func (o *GetLabelsDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get labels default body details items0 based on context it is used
func (o *GetLabelsDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetLabelsDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLabelsDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetLabelsDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetLabelsOKBody GetLabelsResponse is a map of labels names as keys and labels values as a list.
swagger:model GetLabelsOKBody
*/
type GetLabelsOKBody struct {
	// labels
	Labels map[string]GetLabelsOKBodyLabelsAnon `json:"labels,omitempty"`
}

// Validate validates this get labels OK body
func (o *GetLabelsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLabelsOKBody) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(o.Labels) { // not required
		return nil
	}

	for k := range o.Labels {

		if swag.IsZero(o.Labels[k]) { // not required
			continue
		}
		if val, ok := o.Labels[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLabelsOk" + "." + "labels" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getLabelsOk" + "." + "labels" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get labels OK body based on the context it is used
func (o *GetLabelsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLabelsOKBody) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {
	for k := range o.Labels {
		if val, ok := o.Labels[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLabelsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLabelsOKBody) UnmarshalBinary(b []byte) error {
	var res GetLabelsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetLabelsOKBodyLabelsAnon ListLabelValues is list of label's values.
swagger:model GetLabelsOKBodyLabelsAnon
*/
type GetLabelsOKBodyLabelsAnon struct {
	// values
	Values []string `json:"values"`
}

// Validate validates this get labels OK body labels anon
func (o *GetLabelsOKBodyLabelsAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get labels OK body labels anon based on context it is used
func (o *GetLabelsOKBodyLabelsAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetLabelsOKBodyLabelsAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLabelsOKBodyLabelsAnon) UnmarshalBinary(b []byte) error {
	var res GetLabelsOKBodyLabelsAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
