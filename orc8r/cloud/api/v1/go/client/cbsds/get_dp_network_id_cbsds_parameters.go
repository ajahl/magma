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

// NewGetDpNetworkIDCbsdsParams creates a new GetDpNetworkIDCbsdsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDpNetworkIDCbsdsParams() *GetDpNetworkIDCbsdsParams {
	return &GetDpNetworkIDCbsdsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDpNetworkIDCbsdsParamsWithTimeout creates a new GetDpNetworkIDCbsdsParams object
// with the ability to set a timeout on a request.
func NewGetDpNetworkIDCbsdsParamsWithTimeout(timeout time.Duration) *GetDpNetworkIDCbsdsParams {
	return &GetDpNetworkIDCbsdsParams{
		timeout: timeout,
	}
}

// NewGetDpNetworkIDCbsdsParamsWithContext creates a new GetDpNetworkIDCbsdsParams object
// with the ability to set a context for a request.
func NewGetDpNetworkIDCbsdsParamsWithContext(ctx context.Context) *GetDpNetworkIDCbsdsParams {
	return &GetDpNetworkIDCbsdsParams{
		Context: ctx,
	}
}

// NewGetDpNetworkIDCbsdsParamsWithHTTPClient creates a new GetDpNetworkIDCbsdsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDpNetworkIDCbsdsParamsWithHTTPClient(client *http.Client) *GetDpNetworkIDCbsdsParams {
	return &GetDpNetworkIDCbsdsParams{
		HTTPClient: client,
	}
}

/*
GetDpNetworkIDCbsdsParams contains all the parameters to send to the API endpoint

	for the get dp network ID cbsds operation.

	Typically these are written to a http.Request.
*/
type GetDpNetworkIDCbsdsParams struct {

	/* Limit.

	   Number of record to return
	*/
	Limit *int64

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* Offset.

	   Start index for pagination
	*/
	Offset *int64

	/* SerialNumber.

	   serial number of cbsd
	*/
	SerialNumber *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get dp network ID cbsds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDpNetworkIDCbsdsParams) WithDefaults() *GetDpNetworkIDCbsdsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get dp network ID cbsds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDpNetworkIDCbsdsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get dp network ID cbsds params
func (o *GetDpNetworkIDCbsdsParams) WithTimeout(timeout time.Duration) *GetDpNetworkIDCbsdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dp network ID cbsds params
func (o *GetDpNetworkIDCbsdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dp network ID cbsds params
func (o *GetDpNetworkIDCbsdsParams) WithContext(ctx context.Context) *GetDpNetworkIDCbsdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dp network ID cbsds params
func (o *GetDpNetworkIDCbsdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dp network ID cbsds params
func (o *GetDpNetworkIDCbsdsParams) WithHTTPClient(client *http.Client) *GetDpNetworkIDCbsdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dp network ID cbsds params
func (o *GetDpNetworkIDCbsdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get dp network ID cbsds params
func (o *GetDpNetworkIDCbsdsParams) WithLimit(limit *int64) *GetDpNetworkIDCbsdsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get dp network ID cbsds params
func (o *GetDpNetworkIDCbsdsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNetworkID adds the networkID to the get dp network ID cbsds params
func (o *GetDpNetworkIDCbsdsParams) WithNetworkID(networkID string) *GetDpNetworkIDCbsdsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get dp network ID cbsds params
func (o *GetDpNetworkIDCbsdsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithOffset adds the offset to the get dp network ID cbsds params
func (o *GetDpNetworkIDCbsdsParams) WithOffset(offset *int64) *GetDpNetworkIDCbsdsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get dp network ID cbsds params
func (o *GetDpNetworkIDCbsdsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSerialNumber adds the serialNumber to the get dp network ID cbsds params
func (o *GetDpNetworkIDCbsdsParams) WithSerialNumber(serialNumber *string) *GetDpNetworkIDCbsdsParams {
	o.SetSerialNumber(serialNumber)
	return o
}

// SetSerialNumber adds the serialNumber to the get dp network ID cbsds params
func (o *GetDpNetworkIDCbsdsParams) SetSerialNumber(serialNumber *string) {
	o.SerialNumber = serialNumber
}

// WriteToRequest writes these params to a swagger request
func (o *GetDpNetworkIDCbsdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.SerialNumber != nil {

		// query param serial_number
		var qrSerialNumber string

		if o.SerialNumber != nil {
			qrSerialNumber = *o.SerialNumber
		}
		qSerialNumber := qrSerialNumber
		if qSerialNumber != "" {

			if err := r.SetQueryParam("serial_number", qSerialNumber); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
