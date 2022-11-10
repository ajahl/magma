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

	"magma/orc8r/cloud/api/v1/go/models"
)

// NewPutNetworksNetworkIDRatingGroupsRatingGroupIDParams creates a new PutNetworksNetworkIDRatingGroupsRatingGroupIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutNetworksNetworkIDRatingGroupsRatingGroupIDParams() *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams {
	return &PutNetworksNetworkIDRatingGroupsRatingGroupIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutNetworksNetworkIDRatingGroupsRatingGroupIDParamsWithTimeout creates a new PutNetworksNetworkIDRatingGroupsRatingGroupIDParams object
// with the ability to set a timeout on a request.
func NewPutNetworksNetworkIDRatingGroupsRatingGroupIDParamsWithTimeout(timeout time.Duration) *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams {
	return &PutNetworksNetworkIDRatingGroupsRatingGroupIDParams{
		timeout: timeout,
	}
}

// NewPutNetworksNetworkIDRatingGroupsRatingGroupIDParamsWithContext creates a new PutNetworksNetworkIDRatingGroupsRatingGroupIDParams object
// with the ability to set a context for a request.
func NewPutNetworksNetworkIDRatingGroupsRatingGroupIDParamsWithContext(ctx context.Context) *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams {
	return &PutNetworksNetworkIDRatingGroupsRatingGroupIDParams{
		Context: ctx,
	}
}

// NewPutNetworksNetworkIDRatingGroupsRatingGroupIDParamsWithHTTPClient creates a new PutNetworksNetworkIDRatingGroupsRatingGroupIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutNetworksNetworkIDRatingGroupsRatingGroupIDParamsWithHTTPClient(client *http.Client) *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams {
	return &PutNetworksNetworkIDRatingGroupsRatingGroupIDParams{
		HTTPClient: client,
	}
}

/*
PutNetworksNetworkIDRatingGroupsRatingGroupIDParams contains all the parameters to send to the API endpoint

	for the put networks network ID rating groups rating group ID operation.

	Typically these are written to a http.Request.
*/
type PutNetworksNetworkIDRatingGroupsRatingGroupIDParams struct {

	/* RatingGroup.

	   Rating group
	*/
	RatingGroup *models.MutableRatingGroup

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

// WithDefaults hydrates default values in the put networks network ID rating groups rating group ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams) WithDefaults() *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put networks network ID rating groups rating group ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put networks network ID rating groups rating group ID params
func (o *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams) WithTimeout(timeout time.Duration) *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put networks network ID rating groups rating group ID params
func (o *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put networks network ID rating groups rating group ID params
func (o *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams) WithContext(ctx context.Context) *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put networks network ID rating groups rating group ID params
func (o *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put networks network ID rating groups rating group ID params
func (o *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams) WithHTTPClient(client *http.Client) *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put networks network ID rating groups rating group ID params
func (o *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRatingGroup adds the ratingGroup to the put networks network ID rating groups rating group ID params
func (o *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams) WithRatingGroup(ratingGroup *models.MutableRatingGroup) *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams {
	o.SetRatingGroup(ratingGroup)
	return o
}

// SetRatingGroup adds the ratingGroup to the put networks network ID rating groups rating group ID params
func (o *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams) SetRatingGroup(ratingGroup *models.MutableRatingGroup) {
	o.RatingGroup = ratingGroup
}

// WithNetworkID adds the networkID to the put networks network ID rating groups rating group ID params
func (o *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams) WithNetworkID(networkID string) *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put networks network ID rating groups rating group ID params
func (o *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRatingGroupID adds the ratingGroupID to the put networks network ID rating groups rating group ID params
func (o *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams) WithRatingGroupID(ratingGroupID uint32) *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams {
	o.SetRatingGroupID(ratingGroupID)
	return o
}

// SetRatingGroupID adds the ratingGroupId to the put networks network ID rating groups rating group ID params
func (o *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams) SetRatingGroupID(ratingGroupID uint32) {
	o.RatingGroupID = ratingGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.RatingGroup != nil {
		if err := r.SetBodyParam(o.RatingGroup); err != nil {
			return err
		}
	}

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
