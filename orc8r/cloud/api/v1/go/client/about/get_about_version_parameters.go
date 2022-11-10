// Code generated by go-swagger; DO NOT EDIT.

package about

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

// NewGetAboutVersionParams creates a new GetAboutVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAboutVersionParams() *GetAboutVersionParams {
	return &GetAboutVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAboutVersionParamsWithTimeout creates a new GetAboutVersionParams object
// with the ability to set a timeout on a request.
func NewGetAboutVersionParamsWithTimeout(timeout time.Duration) *GetAboutVersionParams {
	return &GetAboutVersionParams{
		timeout: timeout,
	}
}

// NewGetAboutVersionParamsWithContext creates a new GetAboutVersionParams object
// with the ability to set a context for a request.
func NewGetAboutVersionParamsWithContext(ctx context.Context) *GetAboutVersionParams {
	return &GetAboutVersionParams{
		Context: ctx,
	}
}

// NewGetAboutVersionParamsWithHTTPClient creates a new GetAboutVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAboutVersionParamsWithHTTPClient(client *http.Client) *GetAboutVersionParams {
	return &GetAboutVersionParams{
		HTTPClient: client,
	}
}

/*
GetAboutVersionParams contains all the parameters to send to the API endpoint

	for the get about version operation.

	Typically these are written to a http.Request.
*/
type GetAboutVersionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get about version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAboutVersionParams) WithDefaults() *GetAboutVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get about version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAboutVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get about version params
func (o *GetAboutVersionParams) WithTimeout(timeout time.Duration) *GetAboutVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get about version params
func (o *GetAboutVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get about version params
func (o *GetAboutVersionParams) WithContext(ctx context.Context) *GetAboutVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get about version params
func (o *GetAboutVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get about version params
func (o *GetAboutVersionParams) WithHTTPClient(client *http.Client) *GetAboutVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get about version params
func (o *GetAboutVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAboutVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
