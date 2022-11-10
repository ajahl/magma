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

// NewGetCwfNetworkIDHaPairsHaPairIDStatusParams creates a new GetCwfNetworkIDHaPairsHaPairIDStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCwfNetworkIDHaPairsHaPairIDStatusParams() *GetCwfNetworkIDHaPairsHaPairIDStatusParams {
	return &GetCwfNetworkIDHaPairsHaPairIDStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCwfNetworkIDHaPairsHaPairIDStatusParamsWithTimeout creates a new GetCwfNetworkIDHaPairsHaPairIDStatusParams object
// with the ability to set a timeout on a request.
func NewGetCwfNetworkIDHaPairsHaPairIDStatusParamsWithTimeout(timeout time.Duration) *GetCwfNetworkIDHaPairsHaPairIDStatusParams {
	return &GetCwfNetworkIDHaPairsHaPairIDStatusParams{
		timeout: timeout,
	}
}

// NewGetCwfNetworkIDHaPairsHaPairIDStatusParamsWithContext creates a new GetCwfNetworkIDHaPairsHaPairIDStatusParams object
// with the ability to set a context for a request.
func NewGetCwfNetworkIDHaPairsHaPairIDStatusParamsWithContext(ctx context.Context) *GetCwfNetworkIDHaPairsHaPairIDStatusParams {
	return &GetCwfNetworkIDHaPairsHaPairIDStatusParams{
		Context: ctx,
	}
}

// NewGetCwfNetworkIDHaPairsHaPairIDStatusParamsWithHTTPClient creates a new GetCwfNetworkIDHaPairsHaPairIDStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCwfNetworkIDHaPairsHaPairIDStatusParamsWithHTTPClient(client *http.Client) *GetCwfNetworkIDHaPairsHaPairIDStatusParams {
	return &GetCwfNetworkIDHaPairsHaPairIDStatusParams{
		HTTPClient: client,
	}
}

/*
GetCwfNetworkIDHaPairsHaPairIDStatusParams contains all the parameters to send to the API endpoint

	for the get cwf network ID ha pairs ha pair ID status operation.

	Typically these are written to a http.Request.
*/
type GetCwfNetworkIDHaPairsHaPairIDStatusParams struct {

	/* HaPairID.

	   HA Gateway Pair ID
	*/
	HaPairID string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cwf network ID ha pairs ha pair ID status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCwfNetworkIDHaPairsHaPairIDStatusParams) WithDefaults() *GetCwfNetworkIDHaPairsHaPairIDStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cwf network ID ha pairs ha pair ID status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCwfNetworkIDHaPairsHaPairIDStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cwf network ID ha pairs ha pair ID status params
func (o *GetCwfNetworkIDHaPairsHaPairIDStatusParams) WithTimeout(timeout time.Duration) *GetCwfNetworkIDHaPairsHaPairIDStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cwf network ID ha pairs ha pair ID status params
func (o *GetCwfNetworkIDHaPairsHaPairIDStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cwf network ID ha pairs ha pair ID status params
func (o *GetCwfNetworkIDHaPairsHaPairIDStatusParams) WithContext(ctx context.Context) *GetCwfNetworkIDHaPairsHaPairIDStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cwf network ID ha pairs ha pair ID status params
func (o *GetCwfNetworkIDHaPairsHaPairIDStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cwf network ID ha pairs ha pair ID status params
func (o *GetCwfNetworkIDHaPairsHaPairIDStatusParams) WithHTTPClient(client *http.Client) *GetCwfNetworkIDHaPairsHaPairIDStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cwf network ID ha pairs ha pair ID status params
func (o *GetCwfNetworkIDHaPairsHaPairIDStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHaPairID adds the haPairID to the get cwf network ID ha pairs ha pair ID status params
func (o *GetCwfNetworkIDHaPairsHaPairIDStatusParams) WithHaPairID(haPairID string) *GetCwfNetworkIDHaPairsHaPairIDStatusParams {
	o.SetHaPairID(haPairID)
	return o
}

// SetHaPairID adds the haPairId to the get cwf network ID ha pairs ha pair ID status params
func (o *GetCwfNetworkIDHaPairsHaPairIDStatusParams) SetHaPairID(haPairID string) {
	o.HaPairID = haPairID
}

// WithNetworkID adds the networkID to the get cwf network ID ha pairs ha pair ID status params
func (o *GetCwfNetworkIDHaPairsHaPairIDStatusParams) WithNetworkID(networkID string) *GetCwfNetworkIDHaPairsHaPairIDStatusParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get cwf network ID ha pairs ha pair ID status params
func (o *GetCwfNetworkIDHaPairsHaPairIDStatusParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCwfNetworkIDHaPairsHaPairIDStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ha_pair_id
	if err := r.SetPathParam("ha_pair_id", o.HaPairID); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
