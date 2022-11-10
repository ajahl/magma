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

// NewGetCwfNetworkIDGatewaysGatewayIDHealthStatusParams creates a new GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCwfNetworkIDGatewaysGatewayIDHealthStatusParams() *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams {
	return &GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCwfNetworkIDGatewaysGatewayIDHealthStatusParamsWithTimeout creates a new GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams object
// with the ability to set a timeout on a request.
func NewGetCwfNetworkIDGatewaysGatewayIDHealthStatusParamsWithTimeout(timeout time.Duration) *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams {
	return &GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams{
		timeout: timeout,
	}
}

// NewGetCwfNetworkIDGatewaysGatewayIDHealthStatusParamsWithContext creates a new GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams object
// with the ability to set a context for a request.
func NewGetCwfNetworkIDGatewaysGatewayIDHealthStatusParamsWithContext(ctx context.Context) *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams {
	return &GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams{
		Context: ctx,
	}
}

// NewGetCwfNetworkIDGatewaysGatewayIDHealthStatusParamsWithHTTPClient creates a new GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCwfNetworkIDGatewaysGatewayIDHealthStatusParamsWithHTTPClient(client *http.Client) *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams {
	return &GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams{
		HTTPClient: client,
	}
}

/*
GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams contains all the parameters to send to the API endpoint

	for the get cwf network ID gateways gateway ID health status operation.

	Typically these are written to a http.Request.
*/
type GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams struct {

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

// WithDefaults hydrates default values in the get cwf network ID gateways gateway ID health status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams) WithDefaults() *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cwf network ID gateways gateway ID health status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cwf network ID gateways gateway ID health status params
func (o *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams) WithTimeout(timeout time.Duration) *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cwf network ID gateways gateway ID health status params
func (o *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cwf network ID gateways gateway ID health status params
func (o *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams) WithContext(ctx context.Context) *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cwf network ID gateways gateway ID health status params
func (o *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cwf network ID gateways gateway ID health status params
func (o *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams) WithHTTPClient(client *http.Client) *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cwf network ID gateways gateway ID health status params
func (o *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the get cwf network ID gateways gateway ID health status params
func (o *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams) WithGatewayID(gatewayID string) *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the get cwf network ID gateways gateway ID health status params
func (o *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the get cwf network ID gateways gateway ID health status params
func (o *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams) WithNetworkID(networkID string) *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get cwf network ID gateways gateway ID health status params
func (o *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
