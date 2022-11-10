// Code generated by go-swagger; DO NOT EDIT.

package lte_gateways

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

// NewGetLTENetworkIDGatewaysGatewayIDNameParams creates a new GetLTENetworkIDGatewaysGatewayIDNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLTENetworkIDGatewaysGatewayIDNameParams() *GetLTENetworkIDGatewaysGatewayIDNameParams {
	return &GetLTENetworkIDGatewaysGatewayIDNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDNameParamsWithTimeout creates a new GetLTENetworkIDGatewaysGatewayIDNameParams object
// with the ability to set a timeout on a request.
func NewGetLTENetworkIDGatewaysGatewayIDNameParamsWithTimeout(timeout time.Duration) *GetLTENetworkIDGatewaysGatewayIDNameParams {
	return &GetLTENetworkIDGatewaysGatewayIDNameParams{
		timeout: timeout,
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDNameParamsWithContext creates a new GetLTENetworkIDGatewaysGatewayIDNameParams object
// with the ability to set a context for a request.
func NewGetLTENetworkIDGatewaysGatewayIDNameParamsWithContext(ctx context.Context) *GetLTENetworkIDGatewaysGatewayIDNameParams {
	return &GetLTENetworkIDGatewaysGatewayIDNameParams{
		Context: ctx,
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDNameParamsWithHTTPClient creates a new GetLTENetworkIDGatewaysGatewayIDNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLTENetworkIDGatewaysGatewayIDNameParamsWithHTTPClient(client *http.Client) *GetLTENetworkIDGatewaysGatewayIDNameParams {
	return &GetLTENetworkIDGatewaysGatewayIDNameParams{
		HTTPClient: client,
	}
}

/*
GetLTENetworkIDGatewaysGatewayIDNameParams contains all the parameters to send to the API endpoint

	for the get LTE network ID gateways gateway ID name operation.

	Typically these are written to a http.Request.
*/
type GetLTENetworkIDGatewaysGatewayIDNameParams struct {

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

// WithDefaults hydrates default values in the get LTE network ID gateways gateway ID name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLTENetworkIDGatewaysGatewayIDNameParams) WithDefaults() *GetLTENetworkIDGatewaysGatewayIDNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get LTE network ID gateways gateway ID name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLTENetworkIDGatewaysGatewayIDNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get LTE network ID gateways gateway ID name params
func (o *GetLTENetworkIDGatewaysGatewayIDNameParams) WithTimeout(timeout time.Duration) *GetLTENetworkIDGatewaysGatewayIDNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get LTE network ID gateways gateway ID name params
func (o *GetLTENetworkIDGatewaysGatewayIDNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get LTE network ID gateways gateway ID name params
func (o *GetLTENetworkIDGatewaysGatewayIDNameParams) WithContext(ctx context.Context) *GetLTENetworkIDGatewaysGatewayIDNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get LTE network ID gateways gateway ID name params
func (o *GetLTENetworkIDGatewaysGatewayIDNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get LTE network ID gateways gateway ID name params
func (o *GetLTENetworkIDGatewaysGatewayIDNameParams) WithHTTPClient(client *http.Client) *GetLTENetworkIDGatewaysGatewayIDNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get LTE network ID gateways gateway ID name params
func (o *GetLTENetworkIDGatewaysGatewayIDNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the get LTE network ID gateways gateway ID name params
func (o *GetLTENetworkIDGatewaysGatewayIDNameParams) WithGatewayID(gatewayID string) *GetLTENetworkIDGatewaysGatewayIDNameParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the get LTE network ID gateways gateway ID name params
func (o *GetLTENetworkIDGatewaysGatewayIDNameParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the get LTE network ID gateways gateway ID name params
func (o *GetLTENetworkIDGatewaysGatewayIDNameParams) WithNetworkID(networkID string) *GetLTENetworkIDGatewaysGatewayIDNameParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get LTE network ID gateways gateway ID name params
func (o *GetLTENetworkIDGatewaysGatewayIDNameParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLTENetworkIDGatewaysGatewayIDNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
