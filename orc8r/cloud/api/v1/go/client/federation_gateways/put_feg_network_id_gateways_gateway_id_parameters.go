// Code generated by go-swagger; DO NOT EDIT.

package federation_gateways

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

	"magma/orc8r/cloud/api/v1/go/models"
)

// NewPutFegNetworkIDGatewaysGatewayIDParams creates a new PutFegNetworkIDGatewaysGatewayIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutFegNetworkIDGatewaysGatewayIDParams() *PutFegNetworkIDGatewaysGatewayIDParams {
	return &PutFegNetworkIDGatewaysGatewayIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutFegNetworkIDGatewaysGatewayIDParamsWithTimeout creates a new PutFegNetworkIDGatewaysGatewayIDParams object
// with the ability to set a timeout on a request.
func NewPutFegNetworkIDGatewaysGatewayIDParamsWithTimeout(timeout time.Duration) *PutFegNetworkIDGatewaysGatewayIDParams {
	return &PutFegNetworkIDGatewaysGatewayIDParams{
		timeout: timeout,
	}
}

// NewPutFegNetworkIDGatewaysGatewayIDParamsWithContext creates a new PutFegNetworkIDGatewaysGatewayIDParams object
// with the ability to set a context for a request.
func NewPutFegNetworkIDGatewaysGatewayIDParamsWithContext(ctx context.Context) *PutFegNetworkIDGatewaysGatewayIDParams {
	return &PutFegNetworkIDGatewaysGatewayIDParams{
		Context: ctx,
	}
}

// NewPutFegNetworkIDGatewaysGatewayIDParamsWithHTTPClient creates a new PutFegNetworkIDGatewaysGatewayIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutFegNetworkIDGatewaysGatewayIDParamsWithHTTPClient(client *http.Client) *PutFegNetworkIDGatewaysGatewayIDParams {
	return &PutFegNetworkIDGatewaysGatewayIDParams{
		HTTPClient: client,
	}
}

/*
PutFegNetworkIDGatewaysGatewayIDParams contains all the parameters to send to the API endpoint

	for the put feg network ID gateways gateway ID operation.

	Typically these are written to a http.Request.
*/
type PutFegNetworkIDGatewaysGatewayIDParams struct {

	/* Gateway.

	   Full desired configuration of the gateway
	*/
	Gateway *models.MutableFederationGateway

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

// WithDefaults hydrates default values in the put feg network ID gateways gateway ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutFegNetworkIDGatewaysGatewayIDParams) WithDefaults() *PutFegNetworkIDGatewaysGatewayIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put feg network ID gateways gateway ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutFegNetworkIDGatewaysGatewayIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put feg network ID gateways gateway ID params
func (o *PutFegNetworkIDGatewaysGatewayIDParams) WithTimeout(timeout time.Duration) *PutFegNetworkIDGatewaysGatewayIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put feg network ID gateways gateway ID params
func (o *PutFegNetworkIDGatewaysGatewayIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put feg network ID gateways gateway ID params
func (o *PutFegNetworkIDGatewaysGatewayIDParams) WithContext(ctx context.Context) *PutFegNetworkIDGatewaysGatewayIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put feg network ID gateways gateway ID params
func (o *PutFegNetworkIDGatewaysGatewayIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put feg network ID gateways gateway ID params
func (o *PutFegNetworkIDGatewaysGatewayIDParams) WithHTTPClient(client *http.Client) *PutFegNetworkIDGatewaysGatewayIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put feg network ID gateways gateway ID params
func (o *PutFegNetworkIDGatewaysGatewayIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGateway adds the gateway to the put feg network ID gateways gateway ID params
func (o *PutFegNetworkIDGatewaysGatewayIDParams) WithGateway(gateway *models.MutableFederationGateway) *PutFegNetworkIDGatewaysGatewayIDParams {
	o.SetGateway(gateway)
	return o
}

// SetGateway adds the gateway to the put feg network ID gateways gateway ID params
func (o *PutFegNetworkIDGatewaysGatewayIDParams) SetGateway(gateway *models.MutableFederationGateway) {
	o.Gateway = gateway
}

// WithGatewayID adds the gatewayID to the put feg network ID gateways gateway ID params
func (o *PutFegNetworkIDGatewaysGatewayIDParams) WithGatewayID(gatewayID string) *PutFegNetworkIDGatewaysGatewayIDParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the put feg network ID gateways gateway ID params
func (o *PutFegNetworkIDGatewaysGatewayIDParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the put feg network ID gateways gateway ID params
func (o *PutFegNetworkIDGatewaysGatewayIDParams) WithNetworkID(networkID string) *PutFegNetworkIDGatewaysGatewayIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put feg network ID gateways gateway ID params
func (o *PutFegNetworkIDGatewaysGatewayIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutFegNetworkIDGatewaysGatewayIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Gateway != nil {
		if err := r.SetBodyParam(o.Gateway); err != nil {
			return err
		}
	}

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
