// Code generated by go-swagger; DO NOT EDIT.

package gateways

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

// NewGetNetworksNetworkIDGatewaysGatewayIDParams creates a new GetNetworksNetworkIDGatewaysGatewayIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworksNetworkIDGatewaysGatewayIDParams() *GetNetworksNetworkIDGatewaysGatewayIDParams {
	return &GetNetworksNetworkIDGatewaysGatewayIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworksNetworkIDGatewaysGatewayIDParamsWithTimeout creates a new GetNetworksNetworkIDGatewaysGatewayIDParams object
// with the ability to set a timeout on a request.
func NewGetNetworksNetworkIDGatewaysGatewayIDParamsWithTimeout(timeout time.Duration) *GetNetworksNetworkIDGatewaysGatewayIDParams {
	return &GetNetworksNetworkIDGatewaysGatewayIDParams{
		timeout: timeout,
	}
}

// NewGetNetworksNetworkIDGatewaysGatewayIDParamsWithContext creates a new GetNetworksNetworkIDGatewaysGatewayIDParams object
// with the ability to set a context for a request.
func NewGetNetworksNetworkIDGatewaysGatewayIDParamsWithContext(ctx context.Context) *GetNetworksNetworkIDGatewaysGatewayIDParams {
	return &GetNetworksNetworkIDGatewaysGatewayIDParams{
		Context: ctx,
	}
}

// NewGetNetworksNetworkIDGatewaysGatewayIDParamsWithHTTPClient creates a new GetNetworksNetworkIDGatewaysGatewayIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworksNetworkIDGatewaysGatewayIDParamsWithHTTPClient(client *http.Client) *GetNetworksNetworkIDGatewaysGatewayIDParams {
	return &GetNetworksNetworkIDGatewaysGatewayIDParams{
		HTTPClient: client,
	}
}

/*
GetNetworksNetworkIDGatewaysGatewayIDParams contains all the parameters to send to the API endpoint

	for the get networks network ID gateways gateway ID operation.

	Typically these are written to a http.Request.
*/
type GetNetworksNetworkIDGatewaysGatewayIDParams struct {

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

// WithDefaults hydrates default values in the get networks network ID gateways gateway ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworksNetworkIDGatewaysGatewayIDParams) WithDefaults() *GetNetworksNetworkIDGatewaysGatewayIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get networks network ID gateways gateway ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworksNetworkIDGatewaysGatewayIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get networks network ID gateways gateway ID params
func (o *GetNetworksNetworkIDGatewaysGatewayIDParams) WithTimeout(timeout time.Duration) *GetNetworksNetworkIDGatewaysGatewayIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get networks network ID gateways gateway ID params
func (o *GetNetworksNetworkIDGatewaysGatewayIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get networks network ID gateways gateway ID params
func (o *GetNetworksNetworkIDGatewaysGatewayIDParams) WithContext(ctx context.Context) *GetNetworksNetworkIDGatewaysGatewayIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get networks network ID gateways gateway ID params
func (o *GetNetworksNetworkIDGatewaysGatewayIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get networks network ID gateways gateway ID params
func (o *GetNetworksNetworkIDGatewaysGatewayIDParams) WithHTTPClient(client *http.Client) *GetNetworksNetworkIDGatewaysGatewayIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get networks network ID gateways gateway ID params
func (o *GetNetworksNetworkIDGatewaysGatewayIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the get networks network ID gateways gateway ID params
func (o *GetNetworksNetworkIDGatewaysGatewayIDParams) WithGatewayID(gatewayID string) *GetNetworksNetworkIDGatewaysGatewayIDParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the get networks network ID gateways gateway ID params
func (o *GetNetworksNetworkIDGatewaysGatewayIDParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the get networks network ID gateways gateway ID params
func (o *GetNetworksNetworkIDGatewaysGatewayIDParams) WithNetworkID(networkID string) *GetNetworksNetworkIDGatewaysGatewayIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get networks network ID gateways gateway ID params
func (o *GetNetworksNetworkIDGatewaysGatewayIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworksNetworkIDGatewaysGatewayIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
