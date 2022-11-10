// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tenants API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tenants API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteTenantsTenantID(params *DeleteTenantsTenantIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteTenantsTenantIDNoContent, error)

	GetTenants(params *GetTenantsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTenantsOK, error)

	GetTenantsTenantID(params *GetTenantsTenantIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTenantsTenantIDOK, error)

	GetTenantsTenantIDControlProxy(params *GetTenantsTenantIDControlProxyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTenantsTenantIDControlProxyOK, error)

	PostTenants(params *PostTenantsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostTenantsCreated, error)

	PutTenantsTenantID(params *PutTenantsTenantIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutTenantsTenantIDNoContent, error)

	PutTenantsTenantIDControlProxy(params *PutTenantsTenantIDControlProxyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutTenantsTenantIDControlProxyNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteTenantsTenantID deletes a tenant this should not be called manually tenants and organizations are updated in n m s and synced to orc8r
*/
func (a *Client) DeleteTenantsTenantID(params *DeleteTenantsTenantIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteTenantsTenantIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTenantsTenantIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteTenantsTenantID",
		Method:             "DELETE",
		PathPattern:        "/tenants/{tenant_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteTenantsTenantIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteTenantsTenantIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteTenantsTenantIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetTenants retrieves all tenants
*/
func (a *Client) GetTenants(params *GetTenantsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTenantsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTenantsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetTenants",
		Method:             "GET",
		PathPattern:        "/tenants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTenantsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTenantsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetTenantsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetTenantsTenantID retrieves tenant info by tenant ID
*/
func (a *Client) GetTenantsTenantID(params *GetTenantsTenantIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTenantsTenantIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTenantsTenantIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetTenantsTenantID",
		Method:             "GET",
		PathPattern:        "/tenants/{tenant_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTenantsTenantIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTenantsTenantIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetTenantsTenantIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetTenantsTenantIDControlProxy retrieves control proxy content by tenant ID
*/
func (a *Client) GetTenantsTenantIDControlProxy(params *GetTenantsTenantIDControlProxyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTenantsTenantIDControlProxyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTenantsTenantIDControlProxyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetTenantsTenantIDControlProxy",
		Method:             "GET",
		PathPattern:        "/tenants/{tenant_id}/control_proxy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTenantsTenantIDControlProxyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTenantsTenantIDControlProxyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetTenantsTenantIDControlProxyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostTenants creates a tenant this should not be called manually tenants and organizations are updated in n m s and synced to orc8r
*/
func (a *Client) PostTenants(params *PostTenantsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostTenantsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostTenantsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostTenants",
		Method:             "POST",
		PathPattern:        "/tenants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostTenantsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostTenantsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostTenantsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PutTenantsTenantID sets tenant info this should not be called manually tenants and organizations are updated in n m s and synced to orc8r
*/
func (a *Client) PutTenantsTenantID(params *PutTenantsTenantIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutTenantsTenantIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutTenantsTenantIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutTenantsTenantID",
		Method:             "PUT",
		PathPattern:        "/tenants/{tenant_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutTenantsTenantIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutTenantsTenantIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutTenantsTenantIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PutTenantsTenantIDControlProxy creates or update control proxy content
*/
func (a *Client) PutTenantsTenantIDControlProxy(params *PutTenantsTenantIDControlProxyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutTenantsTenantIDControlProxyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutTenantsTenantIDControlProxyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutTenantsTenantIDControlProxy",
		Method:             "PUT",
		PathPattern:        "/tenants/{tenant_id}/control_proxy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutTenantsTenantIDControlProxyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutTenantsTenantIDControlProxyNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutTenantsTenantIDControlProxyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
