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

// NewPostDpNetworkIDCbsdsCbsdIDRelinquishParams creates a new PostDpNetworkIDCbsdsCbsdIDRelinquishParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDpNetworkIDCbsdsCbsdIDRelinquishParams() *PostDpNetworkIDCbsdsCbsdIDRelinquishParams {
	return &PostDpNetworkIDCbsdsCbsdIDRelinquishParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDpNetworkIDCbsdsCbsdIDRelinquishParamsWithTimeout creates a new PostDpNetworkIDCbsdsCbsdIDRelinquishParams object
// with the ability to set a timeout on a request.
func NewPostDpNetworkIDCbsdsCbsdIDRelinquishParamsWithTimeout(timeout time.Duration) *PostDpNetworkIDCbsdsCbsdIDRelinquishParams {
	return &PostDpNetworkIDCbsdsCbsdIDRelinquishParams{
		timeout: timeout,
	}
}

// NewPostDpNetworkIDCbsdsCbsdIDRelinquishParamsWithContext creates a new PostDpNetworkIDCbsdsCbsdIDRelinquishParams object
// with the ability to set a context for a request.
func NewPostDpNetworkIDCbsdsCbsdIDRelinquishParamsWithContext(ctx context.Context) *PostDpNetworkIDCbsdsCbsdIDRelinquishParams {
	return &PostDpNetworkIDCbsdsCbsdIDRelinquishParams{
		Context: ctx,
	}
}

// NewPostDpNetworkIDCbsdsCbsdIDRelinquishParamsWithHTTPClient creates a new PostDpNetworkIDCbsdsCbsdIDRelinquishParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDpNetworkIDCbsdsCbsdIDRelinquishParamsWithHTTPClient(client *http.Client) *PostDpNetworkIDCbsdsCbsdIDRelinquishParams {
	return &PostDpNetworkIDCbsdsCbsdIDRelinquishParams{
		HTTPClient: client,
	}
}

/*
PostDpNetworkIDCbsdsCbsdIDRelinquishParams contains all the parameters to send to the API endpoint

	for the post dp network ID cbsds cbsd ID relinquish operation.

	Typically these are written to a http.Request.
*/
type PostDpNetworkIDCbsdsCbsdIDRelinquishParams struct {

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

// WithDefaults hydrates default values in the post dp network ID cbsds cbsd ID relinquish params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDpNetworkIDCbsdsCbsdIDRelinquishParams) WithDefaults() *PostDpNetworkIDCbsdsCbsdIDRelinquishParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post dp network ID cbsds cbsd ID relinquish params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDpNetworkIDCbsdsCbsdIDRelinquishParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post dp network ID cbsds cbsd ID relinquish params
func (o *PostDpNetworkIDCbsdsCbsdIDRelinquishParams) WithTimeout(timeout time.Duration) *PostDpNetworkIDCbsdsCbsdIDRelinquishParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post dp network ID cbsds cbsd ID relinquish params
func (o *PostDpNetworkIDCbsdsCbsdIDRelinquishParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post dp network ID cbsds cbsd ID relinquish params
func (o *PostDpNetworkIDCbsdsCbsdIDRelinquishParams) WithContext(ctx context.Context) *PostDpNetworkIDCbsdsCbsdIDRelinquishParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post dp network ID cbsds cbsd ID relinquish params
func (o *PostDpNetworkIDCbsdsCbsdIDRelinquishParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post dp network ID cbsds cbsd ID relinquish params
func (o *PostDpNetworkIDCbsdsCbsdIDRelinquishParams) WithHTTPClient(client *http.Client) *PostDpNetworkIDCbsdsCbsdIDRelinquishParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post dp network ID cbsds cbsd ID relinquish params
func (o *PostDpNetworkIDCbsdsCbsdIDRelinquishParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCbsdID adds the cbsdID to the post dp network ID cbsds cbsd ID relinquish params
func (o *PostDpNetworkIDCbsdsCbsdIDRelinquishParams) WithCbsdID(cbsdID int64) *PostDpNetworkIDCbsdsCbsdIDRelinquishParams {
	o.SetCbsdID(cbsdID)
	return o
}

// SetCbsdID adds the cbsdId to the post dp network ID cbsds cbsd ID relinquish params
func (o *PostDpNetworkIDCbsdsCbsdIDRelinquishParams) SetCbsdID(cbsdID int64) {
	o.CbsdID = cbsdID
}

// WithNetworkID adds the networkID to the post dp network ID cbsds cbsd ID relinquish params
func (o *PostDpNetworkIDCbsdsCbsdIDRelinquishParams) WithNetworkID(networkID string) *PostDpNetworkIDCbsdsCbsdIDRelinquishParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the post dp network ID cbsds cbsd ID relinquish params
func (o *PostDpNetworkIDCbsdsCbsdIDRelinquishParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PostDpNetworkIDCbsdsCbsdIDRelinquishParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
