// Code generated by go-swagger; DO NOT EDIT.

package enode_bs

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

// NewPostLTENetworkIDEnodebsParams creates a new PostLTENetworkIDEnodebsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostLTENetworkIDEnodebsParams() *PostLTENetworkIDEnodebsParams {
	return &PostLTENetworkIDEnodebsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostLTENetworkIDEnodebsParamsWithTimeout creates a new PostLTENetworkIDEnodebsParams object
// with the ability to set a timeout on a request.
func NewPostLTENetworkIDEnodebsParamsWithTimeout(timeout time.Duration) *PostLTENetworkIDEnodebsParams {
	return &PostLTENetworkIDEnodebsParams{
		timeout: timeout,
	}
}

// NewPostLTENetworkIDEnodebsParamsWithContext creates a new PostLTENetworkIDEnodebsParams object
// with the ability to set a context for a request.
func NewPostLTENetworkIDEnodebsParamsWithContext(ctx context.Context) *PostLTENetworkIDEnodebsParams {
	return &PostLTENetworkIDEnodebsParams{
		Context: ctx,
	}
}

// NewPostLTENetworkIDEnodebsParamsWithHTTPClient creates a new PostLTENetworkIDEnodebsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostLTENetworkIDEnodebsParamsWithHTTPClient(client *http.Client) *PostLTENetworkIDEnodebsParams {
	return &PostLTENetworkIDEnodebsParams{
		HTTPClient: client,
	}
}

/*
PostLTENetworkIDEnodebsParams contains all the parameters to send to the API endpoint

	for the post LTE network ID enodebs operation.

	Typically these are written to a http.Request.
*/
type PostLTENetworkIDEnodebsParams struct {

	/* ENODEB.

	   Configuration of the enodeB
	*/
	ENODEB *models.ENODEB

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post LTE network ID enodebs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLTENetworkIDEnodebsParams) WithDefaults() *PostLTENetworkIDEnodebsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post LTE network ID enodebs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLTENetworkIDEnodebsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post LTE network ID enodebs params
func (o *PostLTENetworkIDEnodebsParams) WithTimeout(timeout time.Duration) *PostLTENetworkIDEnodebsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post LTE network ID enodebs params
func (o *PostLTENetworkIDEnodebsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post LTE network ID enodebs params
func (o *PostLTENetworkIDEnodebsParams) WithContext(ctx context.Context) *PostLTENetworkIDEnodebsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post LTE network ID enodebs params
func (o *PostLTENetworkIDEnodebsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post LTE network ID enodebs params
func (o *PostLTENetworkIDEnodebsParams) WithHTTPClient(client *http.Client) *PostLTENetworkIDEnodebsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post LTE network ID enodebs params
func (o *PostLTENetworkIDEnodebsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithENODEB adds the enodeb to the post LTE network ID enodebs params
func (o *PostLTENetworkIDEnodebsParams) WithENODEB(enodeb *models.ENODEB) *PostLTENetworkIDEnodebsParams {
	o.SetENODEB(enodeb)
	return o
}

// SetENODEB adds the enodeb to the post LTE network ID enodebs params
func (o *PostLTENetworkIDEnodebsParams) SetENODEB(enodeb *models.ENODEB) {
	o.ENODEB = enodeb
}

// WithNetworkID adds the networkID to the post LTE network ID enodebs params
func (o *PostLTENetworkIDEnodebsParams) WithNetworkID(networkID string) *PostLTENetworkIDEnodebsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the post LTE network ID enodebs params
func (o *PostLTENetworkIDEnodebsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLTENetworkIDEnodebsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ENODEB != nil {
		if err := r.SetBodyParam(o.ENODEB); err != nil {
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
