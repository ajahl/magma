// Code generated by go-swagger; DO NOT EDIT.

package rating_groups

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

// NewGetNetworksNetworkIDRatingGroupsRatingGroupIDParams creates a new GetNetworksNetworkIDRatingGroupsRatingGroupIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworksNetworkIDRatingGroupsRatingGroupIDParams() *GetNetworksNetworkIDRatingGroupsRatingGroupIDParams {
	return &GetNetworksNetworkIDRatingGroupsRatingGroupIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworksNetworkIDRatingGroupsRatingGroupIDParamsWithTimeout creates a new GetNetworksNetworkIDRatingGroupsRatingGroupIDParams object
// with the ability to set a timeout on a request.
func NewGetNetworksNetworkIDRatingGroupsRatingGroupIDParamsWithTimeout(timeout time.Duration) *GetNetworksNetworkIDRatingGroupsRatingGroupIDParams {
	return &GetNetworksNetworkIDRatingGroupsRatingGroupIDParams{
		timeout: timeout,
	}
}

// NewGetNetworksNetworkIDRatingGroupsRatingGroupIDParamsWithContext creates a new GetNetworksNetworkIDRatingGroupsRatingGroupIDParams object
// with the ability to set a context for a request.
func NewGetNetworksNetworkIDRatingGroupsRatingGroupIDParamsWithContext(ctx context.Context) *GetNetworksNetworkIDRatingGroupsRatingGroupIDParams {
	return &GetNetworksNetworkIDRatingGroupsRatingGroupIDParams{
		Context: ctx,
	}
}

// NewGetNetworksNetworkIDRatingGroupsRatingGroupIDParamsWithHTTPClient creates a new GetNetworksNetworkIDRatingGroupsRatingGroupIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworksNetworkIDRatingGroupsRatingGroupIDParamsWithHTTPClient(client *http.Client) *GetNetworksNetworkIDRatingGroupsRatingGroupIDParams {
	return &GetNetworksNetworkIDRatingGroupsRatingGroupIDParams{
		HTTPClient: client,
	}
}

/*
GetNetworksNetworkIDRatingGroupsRatingGroupIDParams contains all the parameters to send to the API endpoint

	for the get networks network ID rating groups rating group ID operation.

	Typically these are written to a http.Request.
*/
type GetNetworksNetworkIDRatingGroupsRatingGroupIDParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* RatingGroupID.

	   Rating Group Id

	   Format: uint32
	*/
	RatingGroupID uint32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get networks network ID rating groups rating group ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworksNetworkIDRatingGroupsRatingGroupIDParams) WithDefaults() *GetNetworksNetworkIDRatingGroupsRatingGroupIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get networks network ID rating groups rating group ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworksNetworkIDRatingGroupsRatingGroupIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get networks network ID rating groups rating group ID params
func (o *GetNetworksNetworkIDRatingGroupsRatingGroupIDParams) WithTimeout(timeout time.Duration) *GetNetworksNetworkIDRatingGroupsRatingGroupIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get networks network ID rating groups rating group ID params
func (o *GetNetworksNetworkIDRatingGroupsRatingGroupIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get networks network ID rating groups rating group ID params
func (o *GetNetworksNetworkIDRatingGroupsRatingGroupIDParams) WithContext(ctx context.Context) *GetNetworksNetworkIDRatingGroupsRatingGroupIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get networks network ID rating groups rating group ID params
func (o *GetNetworksNetworkIDRatingGroupsRatingGroupIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get networks network ID rating groups rating group ID params
func (o *GetNetworksNetworkIDRatingGroupsRatingGroupIDParams) WithHTTPClient(client *http.Client) *GetNetworksNetworkIDRatingGroupsRatingGroupIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get networks network ID rating groups rating group ID params
func (o *GetNetworksNetworkIDRatingGroupsRatingGroupIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get networks network ID rating groups rating group ID params
func (o *GetNetworksNetworkIDRatingGroupsRatingGroupIDParams) WithNetworkID(networkID string) *GetNetworksNetworkIDRatingGroupsRatingGroupIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get networks network ID rating groups rating group ID params
func (o *GetNetworksNetworkIDRatingGroupsRatingGroupIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRatingGroupID adds the ratingGroupID to the get networks network ID rating groups rating group ID params
func (o *GetNetworksNetworkIDRatingGroupsRatingGroupIDParams) WithRatingGroupID(ratingGroupID uint32) *GetNetworksNetworkIDRatingGroupsRatingGroupIDParams {
	o.SetRatingGroupID(ratingGroupID)
	return o
}

// SetRatingGroupID adds the ratingGroupId to the get networks network ID rating groups rating group ID params
func (o *GetNetworksNetworkIDRatingGroupsRatingGroupIDParams) SetRatingGroupID(ratingGroupID uint32) {
	o.RatingGroupID = ratingGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworksNetworkIDRatingGroupsRatingGroupIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	// path param rating_group_id
	if err := r.SetPathParam("rating_group_id", swag.FormatUint32(o.RatingGroupID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
