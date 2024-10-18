// Code generated by go-swagger; DO NOT EDIT.

package server_service

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

// ChangeSettingsReader is a Reader for the ChangeSettings structure.
type ChangeSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeSettingsOK creates a ChangeSettingsOK with default headers values
func NewChangeSettingsOK() *ChangeSettingsOK {
	return &ChangeSettingsOK{}
}

/*
ChangeSettingsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ChangeSettingsOK struct {
	Payload *ChangeSettingsOKBody
}

// IsSuccess returns true when this change settings Ok response has a 2xx status code
func (o *ChangeSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this change settings Ok response has a 3xx status code
func (o *ChangeSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change settings Ok response has a 4xx status code
func (o *ChangeSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this change settings Ok response has a 5xx status code
func (o *ChangeSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this change settings Ok response a status code equal to that given
func (o *ChangeSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the change settings Ok response
func (o *ChangeSettingsOK) Code() int {
	return 200
}

func (o *ChangeSettingsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/server/settings][%d] changeSettingsOk %s", 200, payload)
}

func (o *ChangeSettingsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/server/settings][%d] changeSettingsOk %s", 200, payload)
}

func (o *ChangeSettingsOK) GetPayload() *ChangeSettingsOKBody {
	return o.Payload
}

func (o *ChangeSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ChangeSettingsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeSettingsDefault creates a ChangeSettingsDefault with default headers values
func NewChangeSettingsDefault(code int) *ChangeSettingsDefault {
	return &ChangeSettingsDefault{
		_statusCode: code,
	}
}

/*
ChangeSettingsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ChangeSettingsDefault struct {
	_statusCode int

	Payload *ChangeSettingsDefaultBody
}

// IsSuccess returns true when this change settings default response has a 2xx status code
func (o *ChangeSettingsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this change settings default response has a 3xx status code
func (o *ChangeSettingsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this change settings default response has a 4xx status code
func (o *ChangeSettingsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this change settings default response has a 5xx status code
func (o *ChangeSettingsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this change settings default response a status code equal to that given
func (o *ChangeSettingsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the change settings default response
func (o *ChangeSettingsDefault) Code() int {
	return o._statusCode
}

func (o *ChangeSettingsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/server/settings][%d] ChangeSettings default %s", o._statusCode, payload)
}

func (o *ChangeSettingsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/server/settings][%d] ChangeSettings default %s", o._statusCode, payload)
}

func (o *ChangeSettingsDefault) GetPayload() *ChangeSettingsDefaultBody {
	return o.Payload
}

func (o *ChangeSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ChangeSettingsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ChangeSettingsBody change settings body
swagger:model ChangeSettingsBody
*/
type ChangeSettingsBody struct {
	// enable updates
	EnableUpdates *bool `json:"enable_updates,omitempty"`

	// enable telemetry
	EnableTelemetry *bool `json:"enable_telemetry,omitempty"`

	// A number of full days for Prometheus and QAN data retention. Should have a suffix in JSON: 2592000s, 43200m, 720h.
	DataRetention string `json:"data_retention,omitempty"`

	// ssh key
	SSHKey *string `json:"ssh_key,omitempty"`

	// Enable Advisor.
	EnableAdvisor *bool `json:"enable_advisor,omitempty"`

	// Enable Alerting.
	EnableAlerting *bool `json:"enable_alerting,omitempty"`

	// PMM Server public address.
	PMMPublicAddress *string `json:"pmm_public_address,omitempty"`

	// Enable Azure Discover.
	EnableAzurediscover *bool `json:"enable_azurediscover,omitempty"`

	// Enable Backup Management.
	EnableBackupManagement *bool `json:"enable_backup_management,omitempty"`

	// Enable Access Control
	EnableAccessControl *bool `json:"enable_access_control,omitempty"`

	// advisor run intervals
	AdvisorRunIntervals *ChangeSettingsParamsBodyAdvisorRunIntervals `json:"advisor_run_intervals,omitempty"`

	// aws partitions
	AWSPartitions *ChangeSettingsParamsBodyAWSPartitions `json:"aws_partitions,omitempty"`

	// metrics resolutions
	MetricsResolutions *ChangeSettingsParamsBodyMetricsResolutions `json:"metrics_resolutions,omitempty"`
}

// Validate validates this change settings body
func (o *ChangeSettingsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAdvisorRunIntervals(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAWSPartitions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMetricsResolutions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeSettingsBody) validateAdvisorRunIntervals(formats strfmt.Registry) error {
	if swag.IsZero(o.AdvisorRunIntervals) { // not required
		return nil
	}

	if o.AdvisorRunIntervals != nil {
		if err := o.AdvisorRunIntervals.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "advisor_run_intervals")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "advisor_run_intervals")
			}
			return err
		}
	}

	return nil
}

func (o *ChangeSettingsBody) validateAWSPartitions(formats strfmt.Registry) error {
	if swag.IsZero(o.AWSPartitions) { // not required
		return nil
	}

	if o.AWSPartitions != nil {
		if err := o.AWSPartitions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "aws_partitions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "aws_partitions")
			}
			return err
		}
	}

	return nil
}

func (o *ChangeSettingsBody) validateMetricsResolutions(formats strfmt.Registry) error {
	if swag.IsZero(o.MetricsResolutions) { // not required
		return nil
	}

	if o.MetricsResolutions != nil {
		if err := o.MetricsResolutions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "metrics_resolutions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "metrics_resolutions")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this change settings body based on the context it is used
func (o *ChangeSettingsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAdvisorRunIntervals(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateAWSPartitions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMetricsResolutions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeSettingsBody) contextValidateAdvisorRunIntervals(ctx context.Context, formats strfmt.Registry) error {
	if o.AdvisorRunIntervals != nil {

		if swag.IsZero(o.AdvisorRunIntervals) { // not required
			return nil
		}

		if err := o.AdvisorRunIntervals.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "advisor_run_intervals")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "advisor_run_intervals")
			}
			return err
		}
	}

	return nil
}

func (o *ChangeSettingsBody) contextValidateAWSPartitions(ctx context.Context, formats strfmt.Registry) error {
	if o.AWSPartitions != nil {

		if swag.IsZero(o.AWSPartitions) { // not required
			return nil
		}

		if err := o.AWSPartitions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "aws_partitions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "aws_partitions")
			}
			return err
		}
	}

	return nil
}

func (o *ChangeSettingsBody) contextValidateMetricsResolutions(ctx context.Context, formats strfmt.Registry) error {
	if o.MetricsResolutions != nil {

		if swag.IsZero(o.MetricsResolutions) { // not required
			return nil
		}

		if err := o.MetricsResolutions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "metrics_resolutions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "metrics_resolutions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsBody) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeSettingsDefaultBody change settings default body
swagger:model ChangeSettingsDefaultBody
*/
type ChangeSettingsDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ChangeSettingsDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this change settings default body
func (o *ChangeSettingsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeSettingsDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ChangeSettings default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChangeSettings default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this change settings default body based on the context it is used
func (o *ChangeSettingsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeSettingsDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ChangeSettings default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChangeSettings default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeSettingsDefaultBodyDetailsItems0 `Any` contains an arbitrary serialized protocol buffer message along with a
// URL that describes the type of the serialized message.
//
// Protobuf library provides support to pack/unpack Any values in the form
// of utility functions or additional generated methods of the Any type.
//
// Example 1: Pack and unpack a message in C++.
//
//     Foo foo = ...;
//     Any any;
//     any.PackFrom(foo);
//     ...
//     if (any.UnpackTo(&foo)) {
//       ...
//     }
//
// Example 2: Pack and unpack a message in Java.
//
//     Foo foo = ...;
//     Any any = Any.pack(foo);
//     ...
//     if (any.is(Foo.class)) {
//       foo = any.unpack(Foo.class);
//     }
//     // or ...
//     if (any.isSameTypeAs(Foo.getDefaultInstance())) {
//       foo = any.unpack(Foo.getDefaultInstance());
//     }
//
//  Example 3: Pack and unpack a message in Python.
//
//     foo = Foo(...)
//     any = Any()
//     any.Pack(foo)
//     ...
//     if any.Is(Foo.DESCRIPTOR):
//       any.Unpack(foo)
//       ...
//
//  Example 4: Pack and unpack a message in Go
//
//      foo := &pb.Foo{...}
//      any, err := anypb.New(foo)
//      if err != nil {
//        ...
//      }
//      ...
//      foo := &pb.Foo{}
//      if err := any.UnmarshalTo(foo); err != nil {
//        ...
//      }
//
// The pack methods provided by protobuf library will by default use
// 'type.googleapis.com/full.type.name' as the type URL and the unpack
// methods only use the fully qualified type name after the last '/'
// in the type URL, for example "foo.bar.com/x/y.z" will yield type
// name "y.z".
//
// JSON
// ====
// The JSON representation of an `Any` value uses the regular
// representation of the deserialized, embedded message, with an
// additional field `@type` which contains the type URL. Example:
//
//     package google.profile;
//     message Person {
//       string first_name = 1;
//       string last_name = 2;
//     }
//
//     {
//       "@type": "type.googleapis.com/google.profile.Person",
//       "firstName": <string>,
//       "lastName": <string>
//     }
//
// If the embedded message type is well-known and has a custom JSON
// representation, that representation will be embedded adding a field
// `value` which holds the custom JSON in addition to the `@type`
// field. Example (for message [google.protobuf.Duration][]):
//
//     {
//       "@type": "type.googleapis.com/google.protobuf.Duration",
//       "value": "1.212s"
//     }
swagger:model ChangeSettingsDefaultBodyDetailsItems0
*/
type ChangeSettingsDefaultBodyDetailsItems0 struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	// protocol buffer message. This string must contain at least
	// one "/" character. The last segment of the URL's path must represent
	// the fully qualified name of the type (as in
	// `path/google.protobuf.Duration`). The name should be in a canonical form
	// (e.g., leading "." is not accepted).
	//
	// In practice, teams usually precompile into the binary all types that they
	// expect it to use in the context of Any. However, for URLs which use the
	// scheme `http`, `https`, or no scheme, one can optionally set up a type
	// server that maps type URLs to message definitions as follows:
	//
	// * If no scheme is provided, `https` is assumed.
	// * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//   value in binary format, or produce an error.
	// * Applications are allowed to cache lookup results based on the
	//   URL, or have them precompiled into a binary to avoid any
	//   lookup. Therefore, binary compatibility needs to be preserved
	//   on changes to types. (Use versioned type names to manage
	//   breaking changes.)
	//
	// Note: this functionality is not currently available in the official
	// protobuf release, and it is not used for type URLs beginning with
	// type.googleapis.com. As of May 2023, there are no widely used type server
	// implementations and no plans to implement one.
	//
	// Schemes other than `http`, `https` (or the empty scheme) might be
	// used with implementation specific semantics.
	AtType string `json:"@type,omitempty"`

	// change settings default body details items0
	ChangeSettingsDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *ChangeSettingsDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {
		// A URL/resource name that uniquely identifies the type of the serialized
		// protocol buffer message. This string must contain at least
		// one "/" character. The last segment of the URL's path must represent
		// the fully qualified name of the type (as in
		// `path/google.protobuf.Duration`). The name should be in a canonical form
		// (e.g., leading "." is not accepted).
		//
		// In practice, teams usually precompile into the binary all types that they
		// expect it to use in the context of Any. However, for URLs which use the
		// scheme `http`, `https`, or no scheme, one can optionally set up a type
		// server that maps type URLs to message definitions as follows:
		//
		// * If no scheme is provided, `https` is assumed.
		// * An HTTP GET on the URL must yield a [google.protobuf.Type][]
		//   value in binary format, or produce an error.
		// * Applications are allowed to cache lookup results based on the
		//   URL, or have them precompiled into a binary to avoid any
		//   lookup. Therefore, binary compatibility needs to be preserved
		//   on changes to types. (Use versioned type names to manage
		//   breaking changes.)
		//
		// Note: this functionality is not currently available in the official
		// protobuf release, and it is not used for type URLs beginning with
		// type.googleapis.com. As of May 2023, there are no widely used type server
		// implementations and no plans to implement one.
		//
		// Schemes other than `http`, `https` (or the empty scheme) might be
		// used with implementation specific semantics.
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv ChangeSettingsDefaultBodyDetailsItems0

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
		o.ChangeSettingsDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o ChangeSettingsDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
	var stage1 struct {
		// A URL/resource name that uniquely identifies the type of the serialized
		// protocol buffer message. This string must contain at least
		// one "/" character. The last segment of the URL's path must represent
		// the fully qualified name of the type (as in
		// `path/google.protobuf.Duration`). The name should be in a canonical form
		// (e.g., leading "." is not accepted).
		//
		// In practice, teams usually precompile into the binary all types that they
		// expect it to use in the context of Any. However, for URLs which use the
		// scheme `http`, `https`, or no scheme, one can optionally set up a type
		// server that maps type URLs to message definitions as follows:
		//
		// * If no scheme is provided, `https` is assumed.
		// * An HTTP GET on the URL must yield a [google.protobuf.Type][]
		//   value in binary format, or produce an error.
		// * Applications are allowed to cache lookup results based on the
		//   URL, or have them precompiled into a binary to avoid any
		//   lookup. Therefore, binary compatibility needs to be preserved
		//   on changes to types. (Use versioned type names to manage
		//   breaking changes.)
		//
		// Note: this functionality is not currently available in the official
		// protobuf release, and it is not used for type URLs beginning with
		// type.googleapis.com. As of May 2023, there are no widely used type server
		// implementations and no plans to implement one.
		//
		// Schemes other than `http`, `https` (or the empty scheme) might be
		// used with implementation specific semantics.
		AtType string `json:"@type,omitempty"`
	}

	stage1.AtType = o.AtType

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(o.ChangeSettingsDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.ChangeSettingsDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this change settings default body details items0
func (o *ChangeSettingsDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change settings default body details items0 based on context it is used
func (o *ChangeSettingsDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeSettingsOKBody change settings OK body
swagger:model ChangeSettingsOKBody
*/
type ChangeSettingsOKBody struct {
	// settings
	Settings *ChangeSettingsOKBodySettings `json:"settings,omitempty"`
}

// Validate validates this change settings OK body
func (o *ChangeSettingsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeSettingsOKBody) validateSettings(formats strfmt.Registry) error {
	if swag.IsZero(o.Settings) { // not required
		return nil
	}

	if o.Settings != nil {
		if err := o.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeSettingsOk" + "." + "settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("changeSettingsOk" + "." + "settings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this change settings OK body based on the context it is used
func (o *ChangeSettingsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeSettingsOKBody) contextValidateSettings(ctx context.Context, formats strfmt.Registry) error {
	if o.Settings != nil {

		if swag.IsZero(o.Settings) { // not required
			return nil
		}

		if err := o.Settings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeSettingsOk" + "." + "settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("changeSettingsOk" + "." + "settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsOKBody) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeSettingsOKBodySettings Settings represents PMM Server settings.
swagger:model ChangeSettingsOKBodySettings
*/
type ChangeSettingsOKBodySettings struct {
	// True if updates are enabled.
	UpdatesEnabled bool `json:"updates_enabled,omitempty"`

	// True if telemetry is enabled.
	TelemetryEnabled bool `json:"telemetry_enabled,omitempty"`

	// data retention
	DataRetention string `json:"data_retention,omitempty"`

	// ssh key
	SSHKey string `json:"ssh_key,omitempty"`

	// aws partitions
	AWSPartitions []string `json:"aws_partitions"`

	// True if Advisor is enabled.
	AdvisorEnabled bool `json:"advisor_enabled,omitempty"`

	// Percona Platform user's email, if this PMM instance is linked to the Platform.
	PlatformEmail string `json:"platform_email,omitempty"`

	// True if Alerting is enabled.
	AlertingEnabled bool `json:"alerting_enabled,omitempty"`

	// PMM Server public address.
	PMMPublicAddress string `json:"pmm_public_address,omitempty"`

	// True if Backup Management is enabled.
	BackupManagementEnabled bool `json:"backup_management_enabled,omitempty"`

	// True if Azure Discover is enabled.
	AzurediscoverEnabled bool `json:"azurediscover_enabled,omitempty"`

	// True if the PMM instance is connected to Platform
	ConnectedToPlatform bool `json:"connected_to_platform,omitempty"`

	// Includes list of collected telemetry
	TelemetrySummaries []string `json:"telemetry_summaries"`

	// True if Access Control is enabled.
	EnableAccessControl bool `json:"enable_access_control,omitempty"`

	// Default Access Control role ID for new users.
	DefaultRoleID int64 `json:"default_role_id,omitempty"`

	// advisor run intervals
	AdvisorRunIntervals *ChangeSettingsOKBodySettingsAdvisorRunIntervals `json:"advisor_run_intervals,omitempty"`

	// metrics resolutions
	MetricsResolutions *ChangeSettingsOKBodySettingsMetricsResolutions `json:"metrics_resolutions,omitempty"`
}

// Validate validates this change settings OK body settings
func (o *ChangeSettingsOKBodySettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAdvisorRunIntervals(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMetricsResolutions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeSettingsOKBodySettings) validateAdvisorRunIntervals(formats strfmt.Registry) error {
	if swag.IsZero(o.AdvisorRunIntervals) { // not required
		return nil
	}

	if o.AdvisorRunIntervals != nil {
		if err := o.AdvisorRunIntervals.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeSettingsOk" + "." + "settings" + "." + "advisor_run_intervals")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("changeSettingsOk" + "." + "settings" + "." + "advisor_run_intervals")
			}
			return err
		}
	}

	return nil
}

func (o *ChangeSettingsOKBodySettings) validateMetricsResolutions(formats strfmt.Registry) error {
	if swag.IsZero(o.MetricsResolutions) { // not required
		return nil
	}

	if o.MetricsResolutions != nil {
		if err := o.MetricsResolutions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeSettingsOk" + "." + "settings" + "." + "metrics_resolutions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("changeSettingsOk" + "." + "settings" + "." + "metrics_resolutions")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this change settings OK body settings based on the context it is used
func (o *ChangeSettingsOKBodySettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAdvisorRunIntervals(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMetricsResolutions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeSettingsOKBodySettings) contextValidateAdvisorRunIntervals(ctx context.Context, formats strfmt.Registry) error {
	if o.AdvisorRunIntervals != nil {

		if swag.IsZero(o.AdvisorRunIntervals) { // not required
			return nil
		}

		if err := o.AdvisorRunIntervals.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeSettingsOk" + "." + "settings" + "." + "advisor_run_intervals")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("changeSettingsOk" + "." + "settings" + "." + "advisor_run_intervals")
			}
			return err
		}
	}

	return nil
}

func (o *ChangeSettingsOKBodySettings) contextValidateMetricsResolutions(ctx context.Context, formats strfmt.Registry) error {
	if o.MetricsResolutions != nil {

		if swag.IsZero(o.MetricsResolutions) { // not required
			return nil
		}

		if err := o.MetricsResolutions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeSettingsOk" + "." + "settings" + "." + "metrics_resolutions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("changeSettingsOk" + "." + "settings" + "." + "metrics_resolutions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsOKBodySettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsOKBodySettings) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsOKBodySettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeSettingsOKBodySettingsAdvisorRunIntervals AdvisorRunIntervals represents intervals between each run of Advisor checks.
swagger:model ChangeSettingsOKBodySettingsAdvisorRunIntervals
*/
type ChangeSettingsOKBodySettingsAdvisorRunIntervals struct {
	// Standard check interval.
	StandardInterval string `json:"standard_interval,omitempty"`

	// Interval for rare check runs.
	RareInterval string `json:"rare_interval,omitempty"`

	// Interval for frequent check runs.
	FrequentInterval string `json:"frequent_interval,omitempty"`
}

// Validate validates this change settings OK body settings advisor run intervals
func (o *ChangeSettingsOKBodySettingsAdvisorRunIntervals) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change settings OK body settings advisor run intervals based on context it is used
func (o *ChangeSettingsOKBodySettingsAdvisorRunIntervals) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsOKBodySettingsAdvisorRunIntervals) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsOKBodySettingsAdvisorRunIntervals) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsOKBodySettingsAdvisorRunIntervals
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeSettingsOKBodySettingsMetricsResolutions MetricsResolutions represents Prometheus exporters metrics resolutions.
swagger:model ChangeSettingsOKBodySettingsMetricsResolutions
*/
type ChangeSettingsOKBodySettingsMetricsResolutions struct {
	// High resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Hr string `json:"hr,omitempty"`

	// Medium resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Mr string `json:"mr,omitempty"`

	// Low resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Lr string `json:"lr,omitempty"`
}

// Validate validates this change settings OK body settings metrics resolutions
func (o *ChangeSettingsOKBodySettingsMetricsResolutions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change settings OK body settings metrics resolutions based on context it is used
func (o *ChangeSettingsOKBodySettingsMetricsResolutions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsOKBodySettingsMetricsResolutions) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsOKBodySettingsMetricsResolutions) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsOKBodySettingsMetricsResolutions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeSettingsParamsBodyAWSPartitions A wrapper for a string array. This type allows to distinguish between an empty array and a null value.
swagger:model ChangeSettingsParamsBodyAWSPartitions
*/
type ChangeSettingsParamsBodyAWSPartitions struct {
	// values
	Values []string `json:"values"`
}

// Validate validates this change settings params body AWS partitions
func (o *ChangeSettingsParamsBodyAWSPartitions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change settings params body AWS partitions based on context it is used
func (o *ChangeSettingsParamsBodyAWSPartitions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsParamsBodyAWSPartitions) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsParamsBodyAWSPartitions) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsParamsBodyAWSPartitions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeSettingsParamsBodyAdvisorRunIntervals AdvisorRunIntervals represents intervals between each run of Advisor checks.
swagger:model ChangeSettingsParamsBodyAdvisorRunIntervals
*/
type ChangeSettingsParamsBodyAdvisorRunIntervals struct {
	// Standard check interval.
	StandardInterval string `json:"standard_interval,omitempty"`

	// Interval for rare check runs.
	RareInterval string `json:"rare_interval,omitempty"`

	// Interval for frequent check runs.
	FrequentInterval string `json:"frequent_interval,omitempty"`
}

// Validate validates this change settings params body advisor run intervals
func (o *ChangeSettingsParamsBodyAdvisorRunIntervals) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change settings params body advisor run intervals based on context it is used
func (o *ChangeSettingsParamsBodyAdvisorRunIntervals) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsParamsBodyAdvisorRunIntervals) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsParamsBodyAdvisorRunIntervals) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsParamsBodyAdvisorRunIntervals
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeSettingsParamsBodyMetricsResolutions MetricsResolutions represents Prometheus exporters metrics resolutions.
swagger:model ChangeSettingsParamsBodyMetricsResolutions
*/
type ChangeSettingsParamsBodyMetricsResolutions struct {
	// High resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Hr string `json:"hr,omitempty"`

	// Medium resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Mr string `json:"mr,omitempty"`

	// Low resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Lr string `json:"lr,omitempty"`
}

// Validate validates this change settings params body metrics resolutions
func (o *ChangeSettingsParamsBodyMetricsResolutions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change settings params body metrics resolutions based on context it is used
func (o *ChangeSettingsParamsBodyMetricsResolutions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsParamsBodyMetricsResolutions) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsParamsBodyMetricsResolutions) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsParamsBodyMetricsResolutions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
