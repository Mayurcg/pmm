// Code generated by go-swagger; DO NOT EDIT.

package alertmanager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/percona/pmm/api/grafana/gmodels"
)

// RouteGetAMStatusReader is a Reader for the RouteGetAMStatus structure.
type RouteGetAMStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RouteGetAMStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRouteGetAMStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRouteGetAMStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRouteGetAMStatusOK creates a RouteGetAMStatusOK with default headers values
func NewRouteGetAMStatusOK() *RouteGetAMStatusOK {
	return &RouteGetAMStatusOK{}
}

/* RouteGetAMStatusOK describes a response with status code 200, with default header values.

GettableStatus
*/
type RouteGetAMStatusOK struct {
	Payload *gmodels.GettableStatus
}

func (o *RouteGetAMStatusOK) Error() string {
	return fmt.Sprintf("[GET /api/alertmanager/{Recipient}/api/v2/status][%d] routeGetAMStatusOK  %+v", 200, o.Payload)
}
func (o *RouteGetAMStatusOK) GetPayload() *gmodels.GettableStatus {
	return o.Payload
}

func (o *RouteGetAMStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(gmodels.GettableStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRouteGetAMStatusBadRequest creates a RouteGetAMStatusBadRequest with default headers values
func NewRouteGetAMStatusBadRequest() *RouteGetAMStatusBadRequest {
	return &RouteGetAMStatusBadRequest{}
}

/* RouteGetAMStatusBadRequest describes a response with status code 400, with default header values.

ValidationError
*/
type RouteGetAMStatusBadRequest struct {
	Payload *gmodels.ValidationError
}

func (o *RouteGetAMStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/alertmanager/{Recipient}/api/v2/status][%d] routeGetAMStatusBadRequest  %+v", 400, o.Payload)
}
func (o *RouteGetAMStatusBadRequest) GetPayload() *gmodels.ValidationError {
	return o.Payload
}

func (o *RouteGetAMStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(gmodels.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
