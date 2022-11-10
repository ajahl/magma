// Code generated by go-swagger; DO NOT EDIT.

package gateways

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

// NewGetNetworksNetworkIDGatewaysParams creates a new GetNetworksNetworkIDGatewaysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworksNetworkIDGatewaysParams() *GetNetworksNetworkIDGatewaysParams {
	return &GetNetworksNetworkIDGatewaysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworksNetworkIDGatewaysParamsWithTimeout creates a new GetNetworksNetworkIDGatewaysParams object
// with the ability to set a timeout on a request.
func NewGetNetworksNetworkIDGatewaysParamsWithTimeout(timeout time.Duration) *GetNetworksNetworkIDGatewaysParams {
	return &GetNetworksNetworkIDGatewaysParams{
		timeout: timeout,
	}
}

// NewGetNetworksNetworkIDGatewaysParamsWithContext creates a new GetNetworksNetworkIDGatewaysParams object
// with the ability to set a context for a request.
func NewGetNetworksNetworkIDGatewaysParamsWithContext(ctx context.Context) *GetNetworksNetworkIDGatewaysParams {
	return &GetNetworksNetworkIDGatewaysParams{
		Context: ctx,
	}
}

// NewGetNetworksNetworkIDGatewaysParamsWithHTTPClient creates a new GetNetworksNetworkIDGatewaysParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworksNetworkIDGatewaysParamsWithHTTPClient(client *http.Client) *GetNetworksNetworkIDGatewaysParams {
	return &GetNetworksNetworkIDGatewaysParams{
		HTTPClient: client,
	}
}

/*
GetNetworksNetworkIDGatewaysParams contains all the parameters to send to the API endpoint

	for the get networks network ID gateways operation.

	Typically these are written to a http.Request.
*/
type GetNetworksNetworkIDGatewaysParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* PageSize.

	   Maximum number of entities to return

	   Format: uint32
	*/
	PageSize *uint32

	/* PageToken.

	   Opaque page token for paginated requests
	*/
	PageToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get networks network ID gateways params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworksNetworkIDGatewaysParams) WithDefaults() *GetNetworksNetworkIDGatewaysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get networks network ID gateways params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworksNetworkIDGatewaysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get networks network ID gateways params
func (o *GetNetworksNetworkIDGatewaysParams) WithTimeout(timeout time.Duration) *GetNetworksNetworkIDGatewaysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get networks network ID gateways params
func (o *GetNetworksNetworkIDGatewaysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get networks network ID gateways params
func (o *GetNetworksNetworkIDGatewaysParams) WithContext(ctx context.Context) *GetNetworksNetworkIDGatewaysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get networks network ID gateways params
func (o *GetNetworksNetworkIDGatewaysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get networks network ID gateways params
func (o *GetNetworksNetworkIDGatewaysParams) WithHTTPClient(client *http.Client) *GetNetworksNetworkIDGatewaysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get networks network ID gateways params
func (o *GetNetworksNetworkIDGatewaysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get networks network ID gateways params
func (o *GetNetworksNetworkIDGatewaysParams) WithNetworkID(networkID string) *GetNetworksNetworkIDGatewaysParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get networks network ID gateways params
func (o *GetNetworksNetworkIDGatewaysParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithPageSize adds the pageSize to the get networks network ID gateways params
func (o *GetNetworksNetworkIDGatewaysParams) WithPageSize(pageSize *uint32) *GetNetworksNetworkIDGatewaysParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get networks network ID gateways params
func (o *GetNetworksNetworkIDGatewaysParams) SetPageSize(pageSize *uint32) {
	o.PageSize = pageSize
}

// WithPageToken adds the pageToken to the get networks network ID gateways params
func (o *GetNetworksNetworkIDGatewaysParams) WithPageToken(pageToken *string) *GetNetworksNetworkIDGatewaysParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the get networks network ID gateways params
func (o *GetNetworksNetworkIDGatewaysParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworksNetworkIDGatewaysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize uint32

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatUint32(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.PageToken != nil {

		// query param page_token
		var qrPageToken string

		if o.PageToken != nil {
			qrPageToken = *o.PageToken
		}
		qPageToken := qrPageToken
		if qPageToken != "" {

			if err := r.SetQueryParam("page_token", qPageToken); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
