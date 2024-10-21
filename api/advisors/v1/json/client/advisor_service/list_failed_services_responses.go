// Code generated by go-swagger; DO NOT EDIT.

package advisor_service

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

// ListFailedServicesReader is a Reader for the ListFailedServices structure.
type ListFailedServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListFailedServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListFailedServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListFailedServicesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListFailedServicesOK creates a ListFailedServicesOK with default headers values
func NewListFailedServicesOK() *ListFailedServicesOK {
	return &ListFailedServicesOK{}
}

/*
ListFailedServicesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListFailedServicesOK struct {
	Payload *ListFailedServicesOKBody
}

// IsSuccess returns true when this list failed services Ok response has a 2xx status code
func (o *ListFailedServicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list failed services Ok response has a 3xx status code
func (o *ListFailedServicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list failed services Ok response has a 4xx status code
func (o *ListFailedServicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list failed services Ok response has a 5xx status code
func (o *ListFailedServicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list failed services Ok response a status code equal to that given
func (o *ListFailedServicesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list failed services Ok response
func (o *ListFailedServicesOK) Code() int {
	return 200
}

func (o *ListFailedServicesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/advisors/failedServices][%d] listFailedServicesOk %s", 200, payload)
}

func (o *ListFailedServicesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/advisors/failedServices][%d] listFailedServicesOk %s", 200, payload)
}

func (o *ListFailedServicesOK) GetPayload() *ListFailedServicesOKBody {
	return o.Payload
}

func (o *ListFailedServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ListFailedServicesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFailedServicesDefault creates a ListFailedServicesDefault with default headers values
func NewListFailedServicesDefault(code int) *ListFailedServicesDefault {
	return &ListFailedServicesDefault{
		_statusCode: code,
	}
}

/*
ListFailedServicesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListFailedServicesDefault struct {
	_statusCode int

	Payload *ListFailedServicesDefaultBody
}

// IsSuccess returns true when this list failed services default response has a 2xx status code
func (o *ListFailedServicesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list failed services default response has a 3xx status code
func (o *ListFailedServicesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list failed services default response has a 4xx status code
func (o *ListFailedServicesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list failed services default response has a 5xx status code
func (o *ListFailedServicesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list failed services default response a status code equal to that given
func (o *ListFailedServicesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list failed services default response
func (o *ListFailedServicesDefault) Code() int {
	return o._statusCode
}

func (o *ListFailedServicesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/advisors/failedServices][%d] ListFailedServices default %s", o._statusCode, payload)
}

func (o *ListFailedServicesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/advisors/failedServices][%d] ListFailedServices default %s", o._statusCode, payload)
}

func (o *ListFailedServicesDefault) GetPayload() *ListFailedServicesDefaultBody {
	return o.Payload
}

func (o *ListFailedServicesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ListFailedServicesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ListFailedServicesDefaultBody list failed services default body
swagger:model ListFailedServicesDefaultBody
*/
type ListFailedServicesDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ListFailedServicesDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this list failed services default body
func (o *ListFailedServicesDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListFailedServicesDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ListFailedServices default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListFailedServices default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list failed services default body based on the context it is used
func (o *ListFailedServicesDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListFailedServicesDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ListFailedServices default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListFailedServices default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListFailedServicesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListFailedServicesDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListFailedServicesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListFailedServicesDefaultBodyDetailsItems0 list failed services default body details items0
swagger:model ListFailedServicesDefaultBodyDetailsItems0
*/
type ListFailedServicesDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`

	// list failed services default body details items0
	ListFailedServicesDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *ListFailedServicesDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {
		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv ListFailedServicesDefaultBodyDetailsItems0

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
		o.ListFailedServicesDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o ListFailedServicesDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
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

	if len(o.ListFailedServicesDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.ListFailedServicesDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this list failed services default body details items0
func (o *ListFailedServicesDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list failed services default body details items0 based on context it is used
func (o *ListFailedServicesDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListFailedServicesDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListFailedServicesDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ListFailedServicesDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListFailedServicesOKBody list failed services OK body
swagger:model ListFailedServicesOKBody
*/
type ListFailedServicesOKBody struct {
	// result
	Result []*ListFailedServicesOKBodyResultItems0 `json:"result"`
}

// Validate validates this list failed services OK body
func (o *ListFailedServicesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListFailedServicesOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	for i := 0; i < len(o.Result); i++ {
		if swag.IsZero(o.Result[i]) { // not required
			continue
		}

		if o.Result[i] != nil {
			if err := o.Result[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listFailedServicesOk" + "." + "result" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listFailedServicesOk" + "." + "result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list failed services OK body based on the context it is used
func (o *ListFailedServicesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListFailedServicesOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Result); i++ {
		if o.Result[i] != nil {

			if swag.IsZero(o.Result[i]) { // not required
				return nil
			}

			if err := o.Result[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listFailedServicesOk" + "." + "result" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listFailedServicesOk" + "." + "result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListFailedServicesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListFailedServicesOKBody) UnmarshalBinary(b []byte) error {
	var res ListFailedServicesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListFailedServicesOKBodyResultItems0 CheckResultSummary is a summary of check results.
swagger:model ListFailedServicesOKBodyResultItems0
*/
type ListFailedServicesOKBodyResultItems0 struct {
	// service name
	ServiceName string `json:"service_name,omitempty"`

	// service id
	ServiceID string `json:"service_id,omitempty"`

	// Number of failed checks for this service with severity level "EMERGENCY".
	EmergencyCount int64 `json:"emergency_count,omitempty"`

	// Number of failed checks for this service with severity level "ALERT".
	AlertCount int64 `json:"alert_count,omitempty"`

	// Number of failed checks for this service with severity level "CRITICAL".
	CriticalCount int64 `json:"critical_count,omitempty"`

	// Number of failed checks for this service with severity level "ERROR".
	ErrorCount int64 `json:"error_count,omitempty"`

	// Number of failed checks for this service with severity level "WARNING".
	WarningCount int64 `json:"warning_count,omitempty"`

	// Number of failed checks for this service with severity level "NOTICE".
	NoticeCount int64 `json:"notice_count,omitempty"`

	// Number of failed checks for this service with severity level "INFO".
	InfoCount int64 `json:"info_count,omitempty"`

	// Number of failed checks for this service with severity level "DEBUG".
	DebugCount int64 `json:"debug_count,omitempty"`
}

// Validate validates this list failed services OK body result items0
func (o *ListFailedServicesOKBodyResultItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list failed services OK body result items0 based on context it is used
func (o *ListFailedServicesOKBodyResultItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListFailedServicesOKBodyResultItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListFailedServicesOKBodyResultItems0) UnmarshalBinary(b []byte) error {
	var res ListFailedServicesOKBodyResultItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
