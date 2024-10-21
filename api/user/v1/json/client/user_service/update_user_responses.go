// Code generated by go-swagger; DO NOT EDIT.

package user_service

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
)

// UpdateUserReader is a Reader for the UpdateUser structure.
type UpdateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateUserOK creates a UpdateUserOK with default headers values
func NewUpdateUserOK() *UpdateUserOK {
	return &UpdateUserOK{}
}

/*
UpdateUserOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateUserOK struct {
	Payload *UpdateUserOKBody
}

// IsSuccess returns true when this update user Ok response has a 2xx status code
func (o *UpdateUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update user Ok response has a 3xx status code
func (o *UpdateUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user Ok response has a 4xx status code
func (o *UpdateUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update user Ok response has a 5xx status code
func (o *UpdateUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update user Ok response a status code equal to that given
func (o *UpdateUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update user Ok response
func (o *UpdateUserOK) Code() int {
	return 200
}

func (o *UpdateUserOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/users/me][%d] updateUserOk %s", 200, payload)
}

func (o *UpdateUserOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/users/me][%d] updateUserOk %s", 200, payload)
}

func (o *UpdateUserOK) GetPayload() *UpdateUserOKBody {
	return o.Payload
}

func (o *UpdateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(UpdateUserOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserDefault creates a UpdateUserDefault with default headers values
func NewUpdateUserDefault(code int) *UpdateUserDefault {
	return &UpdateUserDefault{
		_statusCode: code,
	}
}

/*
UpdateUserDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateUserDefault struct {
	_statusCode int

	Payload *UpdateUserDefaultBody
}

// IsSuccess returns true when this update user default response has a 2xx status code
func (o *UpdateUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update user default response has a 3xx status code
func (o *UpdateUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update user default response has a 4xx status code
func (o *UpdateUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update user default response has a 5xx status code
func (o *UpdateUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update user default response a status code equal to that given
func (o *UpdateUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update user default response
func (o *UpdateUserDefault) Code() int {
	return o._statusCode
}

func (o *UpdateUserDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/users/me][%d] UpdateUser default %s", o._statusCode, payload)
}

func (o *UpdateUserDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/users/me][%d] UpdateUser default %s", o._statusCode, payload)
}

func (o *UpdateUserDefault) GetPayload() *UpdateUserDefaultBody {
	return o.Payload
}

func (o *UpdateUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(UpdateUserDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateUserBody update user body
swagger:model UpdateUserBody
*/
type UpdateUserBody struct {
	// Product Tour
	ProductTourCompleted *bool `json:"product_tour_completed,omitempty"`

	// Alerting Tour
	AlertingTourCompleted *bool `json:"alerting_tour_completed,omitempty"`

	// Snooze update alert for a PMM version
	SnoozedPMMVersion *string `json:"snoozed_pmm_version,omitempty"`
}

// Validate validates this update user body
func (o *UpdateUserBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update user body based on context it is used
func (o *UpdateUserBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateUserBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateUserBody) UnmarshalBinary(b []byte) error {
	var res UpdateUserBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateUserDefaultBody update user default body
swagger:model UpdateUserDefaultBody
*/
type UpdateUserDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*UpdateUserDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this update user default body
func (o *UpdateUserDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateUserDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("UpdateUser default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UpdateUser default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update user default body based on the context it is used
func (o *UpdateUserDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateUserDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("UpdateUser default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UpdateUser default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateUserDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateUserDefaultBody) UnmarshalBinary(b []byte) error {
	var res UpdateUserDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateUserDefaultBodyDetailsItems0 update user default body details items0
swagger:model UpdateUserDefaultBodyDetailsItems0
*/
type UpdateUserDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`

	// update user default body details items0
	UpdateUserDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *UpdateUserDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {
		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv UpdateUserDefaultBodyDetailsItems0

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
		o.UpdateUserDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o UpdateUserDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
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

	if len(o.UpdateUserDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.UpdateUserDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this update user default body details items0
func (o *UpdateUserDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update user default body details items0 based on context it is used
func (o *UpdateUserDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateUserDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateUserDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res UpdateUserDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateUserOKBody update user OK body
swagger:model UpdateUserOKBody
*/
type UpdateUserOKBody struct {
	// User ID
	UserID int64 `json:"user_id,omitempty"`

	// Product Tour
	ProductTourCompleted bool `json:"product_tour_completed,omitempty"`

	// Alerting Tour
	AlertingTourCompleted bool `json:"alerting_tour_completed,omitempty"`

	// Snooze update alert for a PMM version
	SnoozedPMMVersion string `json:"snoozed_pmm_version,omitempty"`
}

// Validate validates this update user OK body
func (o *UpdateUserOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update user OK body based on context it is used
func (o *UpdateUserOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateUserOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateUserOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateUserOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
