// Code generated by go-swagger; DO NOT EDIT.

package db_clusters

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
)

// NewListDBClustersParams creates a new ListDBClustersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListDBClustersParams() *ListDBClustersParams {
	return &ListDBClustersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListDBClustersParamsWithTimeout creates a new ListDBClustersParams object
// with the ability to set a timeout on a request.
func NewListDBClustersParamsWithTimeout(timeout time.Duration) *ListDBClustersParams {
	return &ListDBClustersParams{
		timeout: timeout,
	}
}

// NewListDBClustersParamsWithContext creates a new ListDBClustersParams object
// with the ability to set a context for a request.
func NewListDBClustersParamsWithContext(ctx context.Context) *ListDBClustersParams {
	return &ListDBClustersParams{
		Context: ctx,
	}
}

// NewListDBClustersParamsWithHTTPClient creates a new ListDBClustersParams object
// with the ability to set a custom HTTPClient for a request.
func NewListDBClustersParamsWithHTTPClient(client *http.Client) *ListDBClustersParams {
	return &ListDBClustersParams{
		HTTPClient: client,
	}
}

/*
ListDBClustersParams contains all the parameters to send to the API endpoint

	for the list DB clusters operation.

	Typically these are written to a http.Request.
*/
type ListDBClustersParams struct {
	// Body.
	Body ListDBClustersBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list DB clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDBClustersParams) WithDefaults() *ListDBClustersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list DB clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDBClustersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list DB clusters params
func (o *ListDBClustersParams) WithTimeout(timeout time.Duration) *ListDBClustersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list DB clusters params
func (o *ListDBClustersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list DB clusters params
func (o *ListDBClustersParams) WithContext(ctx context.Context) *ListDBClustersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list DB clusters params
func (o *ListDBClustersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list DB clusters params
func (o *ListDBClustersParams) WithHTTPClient(client *http.Client) *ListDBClustersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list DB clusters params
func (o *ListDBClustersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list DB clusters params
func (o *ListDBClustersParams) WithBody(body ListDBClustersBody) *ListDBClustersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list DB clusters params
func (o *ListDBClustersParams) SetBody(body ListDBClustersBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListDBClustersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
