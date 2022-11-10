// Code generated by go-swagger; DO NOT EDIT.

package alerts

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

// NewGetNetworksNetworkIDPrometheusAlertReceiverParams creates a new GetNetworksNetworkIDPrometheusAlertReceiverParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworksNetworkIDPrometheusAlertReceiverParams() *GetNetworksNetworkIDPrometheusAlertReceiverParams {
	return &GetNetworksNetworkIDPrometheusAlertReceiverParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworksNetworkIDPrometheusAlertReceiverParamsWithTimeout creates a new GetNetworksNetworkIDPrometheusAlertReceiverParams object
// with the ability to set a timeout on a request.
func NewGetNetworksNetworkIDPrometheusAlertReceiverParamsWithTimeout(timeout time.Duration) *GetNetworksNetworkIDPrometheusAlertReceiverParams {
	return &GetNetworksNetworkIDPrometheusAlertReceiverParams{
		timeout: timeout,
	}
}

// NewGetNetworksNetworkIDPrometheusAlertReceiverParamsWithContext creates a new GetNetworksNetworkIDPrometheusAlertReceiverParams object
// with the ability to set a context for a request.
func NewGetNetworksNetworkIDPrometheusAlertReceiverParamsWithContext(ctx context.Context) *GetNetworksNetworkIDPrometheusAlertReceiverParams {
	return &GetNetworksNetworkIDPrometheusAlertReceiverParams{
		Context: ctx,
	}
}

// NewGetNetworksNetworkIDPrometheusAlertReceiverParamsWithHTTPClient creates a new GetNetworksNetworkIDPrometheusAlertReceiverParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworksNetworkIDPrometheusAlertReceiverParamsWithHTTPClient(client *http.Client) *GetNetworksNetworkIDPrometheusAlertReceiverParams {
	return &GetNetworksNetworkIDPrometheusAlertReceiverParams{
		HTTPClient: client,
	}
}

/*
GetNetworksNetworkIDPrometheusAlertReceiverParams contains all the parameters to send to the API endpoint

	for the get networks network ID prometheus alert receiver operation.

	Typically these are written to a http.Request.
*/
type GetNetworksNetworkIDPrometheusAlertReceiverParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get networks network ID prometheus alert receiver params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworksNetworkIDPrometheusAlertReceiverParams) WithDefaults() *GetNetworksNetworkIDPrometheusAlertReceiverParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get networks network ID prometheus alert receiver params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworksNetworkIDPrometheusAlertReceiverParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get networks network ID prometheus alert receiver params
func (o *GetNetworksNetworkIDPrometheusAlertReceiverParams) WithTimeout(timeout time.Duration) *GetNetworksNetworkIDPrometheusAlertReceiverParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get networks network ID prometheus alert receiver params
func (o *GetNetworksNetworkIDPrometheusAlertReceiverParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get networks network ID prometheus alert receiver params
func (o *GetNetworksNetworkIDPrometheusAlertReceiverParams) WithContext(ctx context.Context) *GetNetworksNetworkIDPrometheusAlertReceiverParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get networks network ID prometheus alert receiver params
func (o *GetNetworksNetworkIDPrometheusAlertReceiverParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get networks network ID prometheus alert receiver params
func (o *GetNetworksNetworkIDPrometheusAlertReceiverParams) WithHTTPClient(client *http.Client) *GetNetworksNetworkIDPrometheusAlertReceiverParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get networks network ID prometheus alert receiver params
func (o *GetNetworksNetworkIDPrometheusAlertReceiverParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get networks network ID prometheus alert receiver params
func (o *GetNetworksNetworkIDPrometheusAlertReceiverParams) WithNetworkID(networkID string) *GetNetworksNetworkIDPrometheusAlertReceiverParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get networks network ID prometheus alert receiver params
func (o *GetNetworksNetworkIDPrometheusAlertReceiverParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworksNetworkIDPrometheusAlertReceiverParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
