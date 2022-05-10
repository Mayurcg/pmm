// Code generated by go-swagger; DO NOT EDIT.

package testing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/percona/pmm/api/grafana/gmodels"
)

// RouteEvalQueriesReader is a Reader for the RouteEvalQueries structure.
type RouteEvalQueriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RouteEvalQueriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRouteEvalQueriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRouteEvalQueriesOK creates a RouteEvalQueriesOK with default headers values
func NewRouteEvalQueriesOK() *RouteEvalQueriesOK {
	return &RouteEvalQueriesOK{}
}

/* RouteEvalQueriesOK describes a response with status code 200, with default header values.

EvalQueriesResponse
*/
type RouteEvalQueriesOK struct {
	Payload gmodels.EvalQueriesResponse
}

func (o *RouteEvalQueriesOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/eval][%d] routeEvalQueriesOK  %+v", 200, o.Payload)
}
func (o *RouteEvalQueriesOK) GetPayload() gmodels.EvalQueriesResponse {
	return o.Payload
}

func (o *RouteEvalQueriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
