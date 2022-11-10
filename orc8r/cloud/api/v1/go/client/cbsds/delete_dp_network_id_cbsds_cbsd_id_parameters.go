// Code generated by go-swagger; DO NOT EDIT.

package cbsds

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
	"github.com/go-openapi/swag"
)

// NewDeleteDpNetworkIDCbsdsCbsdIDParams creates a new DeleteDpNetworkIDCbsdsCbsdIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDpNetworkIDCbsdsCbsdIDParams() *DeleteDpNetworkIDCbsdsCbsdIDParams {
	return &DeleteDpNetworkIDCbsdsCbsdIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDpNetworkIDCbsdsCbsdIDParamsWithTimeout creates a new DeleteDpNetworkIDCbsdsCbsdIDParams object
// with the ability to set a timeout on a request.
func NewDeleteDpNetworkIDCbsdsCbsdIDParamsWithTimeout(timeout time.Duration) *DeleteDpNetworkIDCbsdsCbsdIDParams {
	return &DeleteDpNetworkIDCbsdsCbsdIDParams{
		timeout: timeout,
	}
}

// NewDeleteDpNetworkIDCbsdsCbsdIDParamsWithContext creates a new DeleteDpNetworkIDCbsdsCbsdIDParams object
// with the ability to set a context for a request.
func NewDeleteDpNetworkIDCbsdsCbsdIDParamsWithContext(ctx context.Context) *DeleteDpNetworkIDCbsdsCbsdIDParams {
	return &DeleteDpNetworkIDCbsdsCbsdIDParams{
		Context: ctx,
	}
}

// NewDeleteDpNetworkIDCbsdsCbsdIDParamsWithHTTPClient creates a new DeleteDpNetworkIDCbsdsCbsdIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDpNetworkIDCbsdsCbsdIDParamsWithHTTPClient(client *http.Client) *DeleteDpNetworkIDCbsdsCbsdIDParams {
	return &DeleteDpNetworkIDCbsdsCbsdIDParams{
		HTTPClient: client,
	}
}

/*
DeleteDpNetworkIDCbsdsCbsdIDParams contains all the parameters to send to the API endpoint

	for the delete dp network ID cbsds cbsd ID operation.

	Typically these are written to a http.Request.
*/
type DeleteDpNetworkIDCbsdsCbsdIDParams struct {

	/* CbsdID.

	   CBSD ID
	*/
	CbsdID int64

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete dp network ID cbsds cbsd ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDpNetworkIDCbsdsCbsdIDParams) WithDefaults() *DeleteDpNetworkIDCbsdsCbsdIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete dp network ID cbsds cbsd ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDpNetworkIDCbsdsCbsdIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete dp network ID cbsds cbsd ID params
func (o *DeleteDpNetworkIDCbsdsCbsdIDParams) WithTimeout(timeout time.Duration) *DeleteDpNetworkIDCbsdsCbsdIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete dp network ID cbsds cbsd ID params
func (o *DeleteDpNetworkIDCbsdsCbsdIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete dp network ID cbsds cbsd ID params
func (o *DeleteDpNetworkIDCbsdsCbsdIDParams) WithContext(ctx context.Context) *DeleteDpNetworkIDCbsdsCbsdIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete dp network ID cbsds cbsd ID params
func (o *DeleteDpNetworkIDCbsdsCbsdIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete dp network ID cbsds cbsd ID params
func (o *DeleteDpNetworkIDCbsdsCbsdIDParams) WithHTTPClient(client *http.Client) *DeleteDpNetworkIDCbsdsCbsdIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete dp network ID cbsds cbsd ID params
func (o *DeleteDpNetworkIDCbsdsCbsdIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCbsdID adds the cbsdID to the delete dp network ID cbsds cbsd ID params
func (o *DeleteDpNetworkIDCbsdsCbsdIDParams) WithCbsdID(cbsdID int64) *DeleteDpNetworkIDCbsdsCbsdIDParams {
	o.SetCbsdID(cbsdID)
	return o
}

// SetCbsdID adds the cbsdId to the delete dp network ID cbsds cbsd ID params
func (o *DeleteDpNetworkIDCbsdsCbsdIDParams) SetCbsdID(cbsdID int64) {
	o.CbsdID = cbsdID
}

// WithNetworkID adds the networkID to the delete dp network ID cbsds cbsd ID params
func (o *DeleteDpNetworkIDCbsdsCbsdIDParams) WithNetworkID(networkID string) *DeleteDpNetworkIDCbsdsCbsdIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete dp network ID cbsds cbsd ID params
func (o *DeleteDpNetworkIDCbsdsCbsdIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDpNetworkIDCbsdsCbsdIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cbsd_id
	if err := r.SetPathParam("cbsd_id", swag.FormatInt64(o.CbsdID)); err != nil {
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
