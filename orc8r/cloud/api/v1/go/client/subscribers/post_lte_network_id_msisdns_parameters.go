// Code generated by go-swagger; DO NOT EDIT.

package subscribers

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

// NewPostLTENetworkIDMsisdnsParams creates a new PostLTENetworkIDMsisdnsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostLTENetworkIDMsisdnsParams() *PostLTENetworkIDMsisdnsParams {
	return &PostLTENetworkIDMsisdnsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostLTENetworkIDMsisdnsParamsWithTimeout creates a new PostLTENetworkIDMsisdnsParams object
// with the ability to set a timeout on a request.
func NewPostLTENetworkIDMsisdnsParamsWithTimeout(timeout time.Duration) *PostLTENetworkIDMsisdnsParams {
	return &PostLTENetworkIDMsisdnsParams{
		timeout: timeout,
	}
}

// NewPostLTENetworkIDMsisdnsParamsWithContext creates a new PostLTENetworkIDMsisdnsParams object
// with the ability to set a context for a request.
func NewPostLTENetworkIDMsisdnsParamsWithContext(ctx context.Context) *PostLTENetworkIDMsisdnsParams {
	return &PostLTENetworkIDMsisdnsParams{
		Context: ctx,
	}
}

// NewPostLTENetworkIDMsisdnsParamsWithHTTPClient creates a new PostLTENetworkIDMsisdnsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostLTENetworkIDMsisdnsParamsWithHTTPClient(client *http.Client) *PostLTENetworkIDMsisdnsParams {
	return &PostLTENetworkIDMsisdnsParams{
		HTTPClient: client,
	}
}

/*
PostLTENetworkIDMsisdnsParams contains all the parameters to send to the API endpoint

	for the post LTE network ID msisdns operation.

	Typically these are written to a http.Request.
*/
type PostLTENetworkIDMsisdnsParams struct {

	/* MsisdnAssignment.

	   MSISDN to associate with an existing subscriber ID
	*/
	MsisdnAssignment *models.MsisdnAssignment

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post LTE network ID msisdns params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLTENetworkIDMsisdnsParams) WithDefaults() *PostLTENetworkIDMsisdnsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post LTE network ID msisdns params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLTENetworkIDMsisdnsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post LTE network ID msisdns params
func (o *PostLTENetworkIDMsisdnsParams) WithTimeout(timeout time.Duration) *PostLTENetworkIDMsisdnsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post LTE network ID msisdns params
func (o *PostLTENetworkIDMsisdnsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post LTE network ID msisdns params
func (o *PostLTENetworkIDMsisdnsParams) WithContext(ctx context.Context) *PostLTENetworkIDMsisdnsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post LTE network ID msisdns params
func (o *PostLTENetworkIDMsisdnsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post LTE network ID msisdns params
func (o *PostLTENetworkIDMsisdnsParams) WithHTTPClient(client *http.Client) *PostLTENetworkIDMsisdnsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post LTE network ID msisdns params
func (o *PostLTENetworkIDMsisdnsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMsisdnAssignment adds the msisdnAssignment to the post LTE network ID msisdns params
func (o *PostLTENetworkIDMsisdnsParams) WithMsisdnAssignment(msisdnAssignment *models.MsisdnAssignment) *PostLTENetworkIDMsisdnsParams {
	o.SetMsisdnAssignment(msisdnAssignment)
	return o
}

// SetMsisdnAssignment adds the msisdnAssignment to the post LTE network ID msisdns params
func (o *PostLTENetworkIDMsisdnsParams) SetMsisdnAssignment(msisdnAssignment *models.MsisdnAssignment) {
	o.MsisdnAssignment = msisdnAssignment
}

// WithNetworkID adds the networkID to the post LTE network ID msisdns params
func (o *PostLTENetworkIDMsisdnsParams) WithNetworkID(networkID string) *PostLTENetworkIDMsisdnsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the post LTE network ID msisdns params
func (o *PostLTENetworkIDMsisdnsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLTENetworkIDMsisdnsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.MsisdnAssignment != nil {
		if err := r.SetBodyParam(o.MsisdnAssignment); err != nil {
			return err
		}
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
