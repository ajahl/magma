// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_gateways

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

// NewGetCwfNetworkIDGatewaysGatewayIDParams creates a new GetCwfNetworkIDGatewaysGatewayIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCwfNetworkIDGatewaysGatewayIDParams() *GetCwfNetworkIDGatewaysGatewayIDParams {
	return &GetCwfNetworkIDGatewaysGatewayIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCwfNetworkIDGatewaysGatewayIDParamsWithTimeout creates a new GetCwfNetworkIDGatewaysGatewayIDParams object
// with the ability to set a timeout on a request.
func NewGetCwfNetworkIDGatewaysGatewayIDParamsWithTimeout(timeout time.Duration) *GetCwfNetworkIDGatewaysGatewayIDParams {
	return &GetCwfNetworkIDGatewaysGatewayIDParams{
		timeout: timeout,
	}
}

// NewGetCwfNetworkIDGatewaysGatewayIDParamsWithContext creates a new GetCwfNetworkIDGatewaysGatewayIDParams object
// with the ability to set a context for a request.
func NewGetCwfNetworkIDGatewaysGatewayIDParamsWithContext(ctx context.Context) *GetCwfNetworkIDGatewaysGatewayIDParams {
	return &GetCwfNetworkIDGatewaysGatewayIDParams{
		Context: ctx,
	}
}

// NewGetCwfNetworkIDGatewaysGatewayIDParamsWithHTTPClient creates a new GetCwfNetworkIDGatewaysGatewayIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCwfNetworkIDGatewaysGatewayIDParamsWithHTTPClient(client *http.Client) *GetCwfNetworkIDGatewaysGatewayIDParams {
	return &GetCwfNetworkIDGatewaysGatewayIDParams{
		HTTPClient: client,
	}
}

/*
GetCwfNetworkIDGatewaysGatewayIDParams contains all the parameters to send to the API endpoint

	for the get cwf network ID gateways gateway ID operation.

	Typically these are written to a http.Request.
*/
type GetCwfNetworkIDGatewaysGatewayIDParams struct {

	/* GatewayID.

	   Gateway ID
	*/
	GatewayID string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cwf network ID gateways gateway ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCwfNetworkIDGatewaysGatewayIDParams) WithDefaults() *GetCwfNetworkIDGatewaysGatewayIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cwf network ID gateways gateway ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCwfNetworkIDGatewaysGatewayIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cwf network ID gateways gateway ID params
func (o *GetCwfNetworkIDGatewaysGatewayIDParams) WithTimeout(timeout time.Duration) *GetCwfNetworkIDGatewaysGatewayIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cwf network ID gateways gateway ID params
func (o *GetCwfNetworkIDGatewaysGatewayIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cwf network ID gateways gateway ID params
func (o *GetCwfNetworkIDGatewaysGatewayIDParams) WithContext(ctx context.Context) *GetCwfNetworkIDGatewaysGatewayIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cwf network ID gateways gateway ID params
func (o *GetCwfNetworkIDGatewaysGatewayIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cwf network ID gateways gateway ID params
func (o *GetCwfNetworkIDGatewaysGatewayIDParams) WithHTTPClient(client *http.Client) *GetCwfNetworkIDGatewaysGatewayIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cwf network ID gateways gateway ID params
func (o *GetCwfNetworkIDGatewaysGatewayIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the get cwf network ID gateways gateway ID params
func (o *GetCwfNetworkIDGatewaysGatewayIDParams) WithGatewayID(gatewayID string) *GetCwfNetworkIDGatewaysGatewayIDParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the get cwf network ID gateways gateway ID params
func (o *GetCwfNetworkIDGatewaysGatewayIDParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the get cwf network ID gateways gateway ID params
func (o *GetCwfNetworkIDGatewaysGatewayIDParams) WithNetworkID(networkID string) *GetCwfNetworkIDGatewaysGatewayIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get cwf network ID gateways gateway ID params
func (o *GetCwfNetworkIDGatewaysGatewayIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCwfNetworkIDGatewaysGatewayIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gateway_id
	if err := r.SetPathParam("gateway_id", o.GatewayID); err != nil {
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
