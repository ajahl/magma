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
)

// NewDeleteFegNetworkIDGatewaysGatewayIDParams creates a new DeleteFegNetworkIDGatewaysGatewayIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteFegNetworkIDGatewaysGatewayIDParams() *DeleteFegNetworkIDGatewaysGatewayIDParams {
	return &DeleteFegNetworkIDGatewaysGatewayIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFegNetworkIDGatewaysGatewayIDParamsWithTimeout creates a new DeleteFegNetworkIDGatewaysGatewayIDParams object
// with the ability to set a timeout on a request.
func NewDeleteFegNetworkIDGatewaysGatewayIDParamsWithTimeout(timeout time.Duration) *DeleteFegNetworkIDGatewaysGatewayIDParams {
	return &DeleteFegNetworkIDGatewaysGatewayIDParams{
		timeout: timeout,
	}
}

// NewDeleteFegNetworkIDGatewaysGatewayIDParamsWithContext creates a new DeleteFegNetworkIDGatewaysGatewayIDParams object
// with the ability to set a context for a request.
func NewDeleteFegNetworkIDGatewaysGatewayIDParamsWithContext(ctx context.Context) *DeleteFegNetworkIDGatewaysGatewayIDParams {
	return &DeleteFegNetworkIDGatewaysGatewayIDParams{
		Context: ctx,
	}
}

// NewDeleteFegNetworkIDGatewaysGatewayIDParamsWithHTTPClient creates a new DeleteFegNetworkIDGatewaysGatewayIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteFegNetworkIDGatewaysGatewayIDParamsWithHTTPClient(client *http.Client) *DeleteFegNetworkIDGatewaysGatewayIDParams {
	return &DeleteFegNetworkIDGatewaysGatewayIDParams{
		HTTPClient: client,
	}
}

/*
DeleteFegNetworkIDGatewaysGatewayIDParams contains all the parameters to send to the API endpoint

	for the delete feg network ID gateways gateway ID operation.

	Typically these are written to a http.Request.
*/
type DeleteFegNetworkIDGatewaysGatewayIDParams struct {

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

// WithDefaults hydrates default values in the delete feg network ID gateways gateway ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteFegNetworkIDGatewaysGatewayIDParams) WithDefaults() *DeleteFegNetworkIDGatewaysGatewayIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete feg network ID gateways gateway ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteFegNetworkIDGatewaysGatewayIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete feg network ID gateways gateway ID params
func (o *DeleteFegNetworkIDGatewaysGatewayIDParams) WithTimeout(timeout time.Duration) *DeleteFegNetworkIDGatewaysGatewayIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete feg network ID gateways gateway ID params
func (o *DeleteFegNetworkIDGatewaysGatewayIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete feg network ID gateways gateway ID params
func (o *DeleteFegNetworkIDGatewaysGatewayIDParams) WithContext(ctx context.Context) *DeleteFegNetworkIDGatewaysGatewayIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete feg network ID gateways gateway ID params
func (o *DeleteFegNetworkIDGatewaysGatewayIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete feg network ID gateways gateway ID params
func (o *DeleteFegNetworkIDGatewaysGatewayIDParams) WithHTTPClient(client *http.Client) *DeleteFegNetworkIDGatewaysGatewayIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete feg network ID gateways gateway ID params
func (o *DeleteFegNetworkIDGatewaysGatewayIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the delete feg network ID gateways gateway ID params
func (o *DeleteFegNetworkIDGatewaysGatewayIDParams) WithGatewayID(gatewayID string) *DeleteFegNetworkIDGatewaysGatewayIDParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the delete feg network ID gateways gateway ID params
func (o *DeleteFegNetworkIDGatewaysGatewayIDParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the delete feg network ID gateways gateway ID params
func (o *DeleteFegNetworkIDGatewaysGatewayIDParams) WithNetworkID(networkID string) *DeleteFegNetworkIDGatewaysGatewayIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete feg network ID gateways gateway ID params
func (o *DeleteFegNetworkIDGatewaysGatewayIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFegNetworkIDGatewaysGatewayIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
