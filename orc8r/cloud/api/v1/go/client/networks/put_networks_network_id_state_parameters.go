// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// NewPutNetworksNetworkIDStateParams creates a new PutNetworksNetworkIDStateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutNetworksNetworkIDStateParams() *PutNetworksNetworkIDStateParams {
	return &PutNetworksNetworkIDStateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutNetworksNetworkIDStateParamsWithTimeout creates a new PutNetworksNetworkIDStateParams object
// with the ability to set a timeout on a request.
func NewPutNetworksNetworkIDStateParamsWithTimeout(timeout time.Duration) *PutNetworksNetworkIDStateParams {
	return &PutNetworksNetworkIDStateParams{
		timeout: timeout,
	}
}

// NewPutNetworksNetworkIDStateParamsWithContext creates a new PutNetworksNetworkIDStateParams object
// with the ability to set a context for a request.
func NewPutNetworksNetworkIDStateParamsWithContext(ctx context.Context) *PutNetworksNetworkIDStateParams {
	return &PutNetworksNetworkIDStateParams{
		Context: ctx,
	}
}

// NewPutNetworksNetworkIDStateParamsWithHTTPClient creates a new PutNetworksNetworkIDStateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutNetworksNetworkIDStateParamsWithHTTPClient(client *http.Client) *PutNetworksNetworkIDStateParams {
	return &PutNetworksNetworkIDStateParams{
		HTTPClient: client,
	}
}

/*
PutNetworksNetworkIDStateParams contains all the parameters to send to the API endpoint

	for the put networks network ID state operation.

	Typically these are written to a http.Request.
*/
type PutNetworksNetworkIDStateParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	// StateConfig.
	StateConfig *models.StateConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put networks network ID state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutNetworksNetworkIDStateParams) WithDefaults() *PutNetworksNetworkIDStateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put networks network ID state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutNetworksNetworkIDStateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put networks network ID state params
func (o *PutNetworksNetworkIDStateParams) WithTimeout(timeout time.Duration) *PutNetworksNetworkIDStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put networks network ID state params
func (o *PutNetworksNetworkIDStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put networks network ID state params
func (o *PutNetworksNetworkIDStateParams) WithContext(ctx context.Context) *PutNetworksNetworkIDStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put networks network ID state params
func (o *PutNetworksNetworkIDStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put networks network ID state params
func (o *PutNetworksNetworkIDStateParams) WithHTTPClient(client *http.Client) *PutNetworksNetworkIDStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put networks network ID state params
func (o *PutNetworksNetworkIDStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the put networks network ID state params
func (o *PutNetworksNetworkIDStateParams) WithNetworkID(networkID string) *PutNetworksNetworkIDStateParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put networks network ID state params
func (o *PutNetworksNetworkIDStateParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithStateConfig adds the stateConfig to the put networks network ID state params
func (o *PutNetworksNetworkIDStateParams) WithStateConfig(stateConfig *models.StateConfig) *PutNetworksNetworkIDStateParams {
	o.SetStateConfig(stateConfig)
	return o
}

// SetStateConfig adds the stateConfig to the put networks network ID state params
func (o *PutNetworksNetworkIDStateParams) SetStateConfig(stateConfig *models.StateConfig) {
	o.StateConfig = stateConfig
}

// WriteToRequest writes these params to a swagger request
func (o *PutNetworksNetworkIDStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}
	if o.StateConfig != nil {
		if err := r.SetBodyParam(o.StateConfig); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
