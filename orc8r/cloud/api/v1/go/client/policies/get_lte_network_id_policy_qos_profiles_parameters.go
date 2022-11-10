// Code generated by go-swagger; DO NOT EDIT.

package policies

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

// NewGetLTENetworkIDPolicyQosProfilesParams creates a new GetLTENetworkIDPolicyQosProfilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLTENetworkIDPolicyQosProfilesParams() *GetLTENetworkIDPolicyQosProfilesParams {
	return &GetLTENetworkIDPolicyQosProfilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLTENetworkIDPolicyQosProfilesParamsWithTimeout creates a new GetLTENetworkIDPolicyQosProfilesParams object
// with the ability to set a timeout on a request.
func NewGetLTENetworkIDPolicyQosProfilesParamsWithTimeout(timeout time.Duration) *GetLTENetworkIDPolicyQosProfilesParams {
	return &GetLTENetworkIDPolicyQosProfilesParams{
		timeout: timeout,
	}
}

// NewGetLTENetworkIDPolicyQosProfilesParamsWithContext creates a new GetLTENetworkIDPolicyQosProfilesParams object
// with the ability to set a context for a request.
func NewGetLTENetworkIDPolicyQosProfilesParamsWithContext(ctx context.Context) *GetLTENetworkIDPolicyQosProfilesParams {
	return &GetLTENetworkIDPolicyQosProfilesParams{
		Context: ctx,
	}
}

// NewGetLTENetworkIDPolicyQosProfilesParamsWithHTTPClient creates a new GetLTENetworkIDPolicyQosProfilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLTENetworkIDPolicyQosProfilesParamsWithHTTPClient(client *http.Client) *GetLTENetworkIDPolicyQosProfilesParams {
	return &GetLTENetworkIDPolicyQosProfilesParams{
		HTTPClient: client,
	}
}

/*
GetLTENetworkIDPolicyQosProfilesParams contains all the parameters to send to the API endpoint

	for the get LTE network ID policy qos profiles operation.

	Typically these are written to a http.Request.
*/
type GetLTENetworkIDPolicyQosProfilesParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get LTE network ID policy qos profiles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLTENetworkIDPolicyQosProfilesParams) WithDefaults() *GetLTENetworkIDPolicyQosProfilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get LTE network ID policy qos profiles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLTENetworkIDPolicyQosProfilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get LTE network ID policy qos profiles params
func (o *GetLTENetworkIDPolicyQosProfilesParams) WithTimeout(timeout time.Duration) *GetLTENetworkIDPolicyQosProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get LTE network ID policy qos profiles params
func (o *GetLTENetworkIDPolicyQosProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get LTE network ID policy qos profiles params
func (o *GetLTENetworkIDPolicyQosProfilesParams) WithContext(ctx context.Context) *GetLTENetworkIDPolicyQosProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get LTE network ID policy qos profiles params
func (o *GetLTENetworkIDPolicyQosProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get LTE network ID policy qos profiles params
func (o *GetLTENetworkIDPolicyQosProfilesParams) WithHTTPClient(client *http.Client) *GetLTENetworkIDPolicyQosProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get LTE network ID policy qos profiles params
func (o *GetLTENetworkIDPolicyQosProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get LTE network ID policy qos profiles params
func (o *GetLTENetworkIDPolicyQosProfilesParams) WithNetworkID(networkID string) *GetLTENetworkIDPolicyQosProfilesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get LTE network ID policy qos profiles params
func (o *GetLTENetworkIDPolicyQosProfilesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLTENetworkIDPolicyQosProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
