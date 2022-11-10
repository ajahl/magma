// Code generated by go-swagger; DO NOT EDIT.

package lte_networks

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

// NewDeleteLTENetworkIDDNSRecordsDomainParams creates a new DeleteLTENetworkIDDNSRecordsDomainParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteLTENetworkIDDNSRecordsDomainParams() *DeleteLTENetworkIDDNSRecordsDomainParams {
	return &DeleteLTENetworkIDDNSRecordsDomainParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLTENetworkIDDNSRecordsDomainParamsWithTimeout creates a new DeleteLTENetworkIDDNSRecordsDomainParams object
// with the ability to set a timeout on a request.
func NewDeleteLTENetworkIDDNSRecordsDomainParamsWithTimeout(timeout time.Duration) *DeleteLTENetworkIDDNSRecordsDomainParams {
	return &DeleteLTENetworkIDDNSRecordsDomainParams{
		timeout: timeout,
	}
}

// NewDeleteLTENetworkIDDNSRecordsDomainParamsWithContext creates a new DeleteLTENetworkIDDNSRecordsDomainParams object
// with the ability to set a context for a request.
func NewDeleteLTENetworkIDDNSRecordsDomainParamsWithContext(ctx context.Context) *DeleteLTENetworkIDDNSRecordsDomainParams {
	return &DeleteLTENetworkIDDNSRecordsDomainParams{
		Context: ctx,
	}
}

// NewDeleteLTENetworkIDDNSRecordsDomainParamsWithHTTPClient creates a new DeleteLTENetworkIDDNSRecordsDomainParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteLTENetworkIDDNSRecordsDomainParamsWithHTTPClient(client *http.Client) *DeleteLTENetworkIDDNSRecordsDomainParams {
	return &DeleteLTENetworkIDDNSRecordsDomainParams{
		HTTPClient: client,
	}
}

/*
DeleteLTENetworkIDDNSRecordsDomainParams contains all the parameters to send to the API endpoint

	for the delete LTE network ID DNS records domain operation.

	Typically these are written to a http.Request.
*/
type DeleteLTENetworkIDDNSRecordsDomainParams struct {

	/* Domain.

	   DNS record domain
	*/
	Domain string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete LTE network ID DNS records domain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLTENetworkIDDNSRecordsDomainParams) WithDefaults() *DeleteLTENetworkIDDNSRecordsDomainParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete LTE network ID DNS records domain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLTENetworkIDDNSRecordsDomainParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete LTE network ID DNS records domain params
func (o *DeleteLTENetworkIDDNSRecordsDomainParams) WithTimeout(timeout time.Duration) *DeleteLTENetworkIDDNSRecordsDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete LTE network ID DNS records domain params
func (o *DeleteLTENetworkIDDNSRecordsDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete LTE network ID DNS records domain params
func (o *DeleteLTENetworkIDDNSRecordsDomainParams) WithContext(ctx context.Context) *DeleteLTENetworkIDDNSRecordsDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete LTE network ID DNS records domain params
func (o *DeleteLTENetworkIDDNSRecordsDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete LTE network ID DNS records domain params
func (o *DeleteLTENetworkIDDNSRecordsDomainParams) WithHTTPClient(client *http.Client) *DeleteLTENetworkIDDNSRecordsDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete LTE network ID DNS records domain params
func (o *DeleteLTENetworkIDDNSRecordsDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomain adds the domain to the delete LTE network ID DNS records domain params
func (o *DeleteLTENetworkIDDNSRecordsDomainParams) WithDomain(domain string) *DeleteLTENetworkIDDNSRecordsDomainParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the delete LTE network ID DNS records domain params
func (o *DeleteLTENetworkIDDNSRecordsDomainParams) SetDomain(domain string) {
	o.Domain = domain
}

// WithNetworkID adds the networkID to the delete LTE network ID DNS records domain params
func (o *DeleteLTENetworkIDDNSRecordsDomainParams) WithNetworkID(networkID string) *DeleteLTENetworkIDDNSRecordsDomainParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete LTE network ID DNS records domain params
func (o *DeleteLTENetworkIDDNSRecordsDomainParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLTENetworkIDDNSRecordsDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domain
	if err := r.SetPathParam("domain", o.Domain); err != nil {
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
