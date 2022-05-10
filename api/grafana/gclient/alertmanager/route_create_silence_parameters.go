// Code generated by go-swagger; DO NOT EDIT.

package alertmanager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/percona/pmm/api/grafana/gmodels"
)

// NewRouteCreateSilenceParams creates a new RouteCreateSilenceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRouteCreateSilenceParams() *RouteCreateSilenceParams {
	return &RouteCreateSilenceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRouteCreateSilenceParamsWithTimeout creates a new RouteCreateSilenceParams object
// with the ability to set a timeout on a request.
func NewRouteCreateSilenceParamsWithTimeout(timeout time.Duration) *RouteCreateSilenceParams {
	return &RouteCreateSilenceParams{
		timeout: timeout,
	}
}

// NewRouteCreateSilenceParamsWithContext creates a new RouteCreateSilenceParams object
// with the ability to set a context for a request.
func NewRouteCreateSilenceParamsWithContext(ctx context.Context) *RouteCreateSilenceParams {
	return &RouteCreateSilenceParams{
		Context: ctx,
	}
}

// NewRouteCreateSilenceParamsWithHTTPClient creates a new RouteCreateSilenceParams object
// with the ability to set a custom HTTPClient for a request.
func NewRouteCreateSilenceParamsWithHTTPClient(client *http.Client) *RouteCreateSilenceParams {
	return &RouteCreateSilenceParams{
		HTTPClient: client,
	}
}

/* RouteCreateSilenceParams contains all the parameters to send to the API endpoint
   for the route create silence operation.

   Typically these are written to a http.Request.
*/
type RouteCreateSilenceParams struct {

	/* Recipient.

	     Recipient should be "grafana" for requests to be handled by grafana
	and the numeric datasource id for requests to be forwarded to a datasource
	*/
	Recipient string

	// Silence.
	Silence *gmodels.PostableSilence

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the route create silence params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RouteCreateSilenceParams) WithDefaults() *RouteCreateSilenceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the route create silence params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RouteCreateSilenceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the route create silence params
func (o *RouteCreateSilenceParams) WithTimeout(timeout time.Duration) *RouteCreateSilenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the route create silence params
func (o *RouteCreateSilenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the route create silence params
func (o *RouteCreateSilenceParams) WithContext(ctx context.Context) *RouteCreateSilenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the route create silence params
func (o *RouteCreateSilenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the route create silence params
func (o *RouteCreateSilenceParams) WithHTTPClient(client *http.Client) *RouteCreateSilenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the route create silence params
func (o *RouteCreateSilenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRecipient adds the recipient to the route create silence params
func (o *RouteCreateSilenceParams) WithRecipient(recipient string) *RouteCreateSilenceParams {
	o.SetRecipient(recipient)
	return o
}

// SetRecipient adds the recipient to the route create silence params
func (o *RouteCreateSilenceParams) SetRecipient(recipient string) {
	o.Recipient = recipient
}

// WithSilence adds the silence to the route create silence params
func (o *RouteCreateSilenceParams) WithSilence(silence *gmodels.PostableSilence) *RouteCreateSilenceParams {
	o.SetSilence(silence)
	return o
}

// SetSilence adds the silence to the route create silence params
func (o *RouteCreateSilenceParams) SetSilence(silence *gmodels.PostableSilence) {
	o.Silence = silence
}

// WriteToRequest writes these params to a swagger request
func (o *RouteCreateSilenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Recipient
	if err := r.SetPathParam("Recipient", o.Recipient); err != nil {
		return err
	}
	if o.Silence != nil {
		if err := r.SetBodyParam(o.Silence); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
