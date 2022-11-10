// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_networks

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

// NewGetCwfNetworkIDDescriptionParams creates a new GetCwfNetworkIDDescriptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCwfNetworkIDDescriptionParams() *GetCwfNetworkIDDescriptionParams {
	return &GetCwfNetworkIDDescriptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCwfNetworkIDDescriptionParamsWithTimeout creates a new GetCwfNetworkIDDescriptionParams object
// with the ability to set a timeout on a request.
func NewGetCwfNetworkIDDescriptionParamsWithTimeout(timeout time.Duration) *GetCwfNetworkIDDescriptionParams {
	return &GetCwfNetworkIDDescriptionParams{
		timeout: timeout,
	}
}

// NewGetCwfNetworkIDDescriptionParamsWithContext creates a new GetCwfNetworkIDDescriptionParams object
// with the ability to set a context for a request.
func NewGetCwfNetworkIDDescriptionParamsWithContext(ctx context.Context) *GetCwfNetworkIDDescriptionParams {
	return &GetCwfNetworkIDDescriptionParams{
		Context: ctx,
	}
}

// NewGetCwfNetworkIDDescriptionParamsWithHTTPClient creates a new GetCwfNetworkIDDescriptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCwfNetworkIDDescriptionParamsWithHTTPClient(client *http.Client) *GetCwfNetworkIDDescriptionParams {
	return &GetCwfNetworkIDDescriptionParams{
		HTTPClient: client,
	}
}

/*
GetCwfNetworkIDDescriptionParams contains all the parameters to send to the API endpoint

	for the get cwf network ID description operation.

	Typically these are written to a http.Request.
*/
type GetCwfNetworkIDDescriptionParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cwf network ID description params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCwfNetworkIDDescriptionParams) WithDefaults() *GetCwfNetworkIDDescriptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cwf network ID description params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCwfNetworkIDDescriptionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cwf network ID description params
func (o *GetCwfNetworkIDDescriptionParams) WithTimeout(timeout time.Duration) *GetCwfNetworkIDDescriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cwf network ID description params
func (o *GetCwfNetworkIDDescriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cwf network ID description params
func (o *GetCwfNetworkIDDescriptionParams) WithContext(ctx context.Context) *GetCwfNetworkIDDescriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cwf network ID description params
func (o *GetCwfNetworkIDDescriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cwf network ID description params
func (o *GetCwfNetworkIDDescriptionParams) WithHTTPClient(client *http.Client) *GetCwfNetworkIDDescriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cwf network ID description params
func (o *GetCwfNetworkIDDescriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get cwf network ID description params
func (o *GetCwfNetworkIDDescriptionParams) WithNetworkID(networkID string) *GetCwfNetworkIDDescriptionParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get cwf network ID description params
func (o *GetCwfNetworkIDDescriptionParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCwfNetworkIDDescriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
