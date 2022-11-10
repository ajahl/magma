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

	"magma/orc8r/cloud/api/v1/go/models"
)

// NewPutNetworksNetworkIDPrometheusAlertReceiverReceiverParams creates a new PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutNetworksNetworkIDPrometheusAlertReceiverReceiverParams() *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams {
	return &PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutNetworksNetworkIDPrometheusAlertReceiverReceiverParamsWithTimeout creates a new PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams object
// with the ability to set a timeout on a request.
func NewPutNetworksNetworkIDPrometheusAlertReceiverReceiverParamsWithTimeout(timeout time.Duration) *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams {
	return &PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams{
		timeout: timeout,
	}
}

// NewPutNetworksNetworkIDPrometheusAlertReceiverReceiverParamsWithContext creates a new PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams object
// with the ability to set a context for a request.
func NewPutNetworksNetworkIDPrometheusAlertReceiverReceiverParamsWithContext(ctx context.Context) *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams {
	return &PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams{
		Context: ctx,
	}
}

// NewPutNetworksNetworkIDPrometheusAlertReceiverReceiverParamsWithHTTPClient creates a new PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutNetworksNetworkIDPrometheusAlertReceiverReceiverParamsWithHTTPClient(client *http.Client) *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams {
	return &PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams{
		HTTPClient: client,
	}
}

/*
PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams contains all the parameters to send to the API endpoint

	for the put networks network ID prometheus alert receiver receiver operation.

	Typically these are written to a http.Request.
*/
type PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* Receiver.

	   Name of receiver to be updated
	*/
	Receiver string

	/* ReceiverConfig.

	   Updated alert receiver
	*/
	ReceiverConfig *models.AlertReceiverConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put networks network ID prometheus alert receiver receiver params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams) WithDefaults() *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put networks network ID prometheus alert receiver receiver params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put networks network ID prometheus alert receiver receiver params
func (o *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams) WithTimeout(timeout time.Duration) *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put networks network ID prometheus alert receiver receiver params
func (o *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put networks network ID prometheus alert receiver receiver params
func (o *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams) WithContext(ctx context.Context) *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put networks network ID prometheus alert receiver receiver params
func (o *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put networks network ID prometheus alert receiver receiver params
func (o *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams) WithHTTPClient(client *http.Client) *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put networks network ID prometheus alert receiver receiver params
func (o *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the put networks network ID prometheus alert receiver receiver params
func (o *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams) WithNetworkID(networkID string) *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put networks network ID prometheus alert receiver receiver params
func (o *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithReceiver adds the receiver to the put networks network ID prometheus alert receiver receiver params
func (o *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams) WithReceiver(receiver string) *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams {
	o.SetReceiver(receiver)
	return o
}

// SetReceiver adds the receiver to the put networks network ID prometheus alert receiver receiver params
func (o *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams) SetReceiver(receiver string) {
	o.Receiver = receiver
}

// WithReceiverConfig adds the receiverConfig to the put networks network ID prometheus alert receiver receiver params
func (o *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams) WithReceiverConfig(receiverConfig *models.AlertReceiverConfig) *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams {
	o.SetReceiverConfig(receiverConfig)
	return o
}

// SetReceiverConfig adds the receiverConfig to the put networks network ID prometheus alert receiver receiver params
func (o *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams) SetReceiverConfig(receiverConfig *models.AlertReceiverConfig) {
	o.ReceiverConfig = receiverConfig
}

// WriteToRequest writes these params to a swagger request
func (o *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	// path param receiver
	if err := r.SetPathParam("receiver", o.Receiver); err != nil {
		return err
	}
	if o.ReceiverConfig != nil {
		if err := r.SetBodyParam(o.ReceiverConfig); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
