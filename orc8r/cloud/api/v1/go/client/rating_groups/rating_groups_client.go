// Code generated by go-swagger; DO NOT EDIT.

package rating_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new rating groups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for rating groups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteNetworksNetworkIDRatingGroupsRatingGroupID deletes a rating group
*/
func (a *Client) DeleteNetworksNetworkIDRatingGroupsRatingGroupID(params *DeleteNetworksNetworkIDRatingGroupsRatingGroupIDParams) (*DeleteNetworksNetworkIDRatingGroupsRatingGroupIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNetworksNetworkIDRatingGroupsRatingGroupIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteNetworksNetworkIDRatingGroupsRatingGroupID",
		Method:             "DELETE",
		PathPattern:        "/networks/{network_id}/rating_groups/{rating_group_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteNetworksNetworkIDRatingGroupsRatingGroupIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteNetworksNetworkIDRatingGroupsRatingGroupIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteNetworksNetworkIDRatingGroupsRatingGroupIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetNetworksNetworkIDRatingGroups lists rating groups
*/
func (a *Client) GetNetworksNetworkIDRatingGroups(params *GetNetworksNetworkIDRatingGroupsParams) (*GetNetworksNetworkIDRatingGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworksNetworkIDRatingGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNetworksNetworkIDRatingGroups",
		Method:             "GET",
		PathPattern:        "/networks/{network_id}/rating_groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNetworksNetworkIDRatingGroupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworksNetworkIDRatingGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNetworksNetworkIDRatingGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetNetworksNetworkIDRatingGroupsRatingGroupID gets rating group
*/
func (a *Client) GetNetworksNetworkIDRatingGroupsRatingGroupID(params *GetNetworksNetworkIDRatingGroupsRatingGroupIDParams) (*GetNetworksNetworkIDRatingGroupsRatingGroupIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworksNetworkIDRatingGroupsRatingGroupIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNetworksNetworkIDRatingGroupsRatingGroupID",
		Method:             "GET",
		PathPattern:        "/networks/{network_id}/rating_groups/{rating_group_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNetworksNetworkIDRatingGroupsRatingGroupIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworksNetworkIDRatingGroupsRatingGroupIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNetworksNetworkIDRatingGroupsRatingGroupIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostNetworksNetworkIDRatingGroups adds a new rating group
*/
func (a *Client) PostNetworksNetworkIDRatingGroups(params *PostNetworksNetworkIDRatingGroupsParams) (*PostNetworksNetworkIDRatingGroupsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNetworksNetworkIDRatingGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostNetworksNetworkIDRatingGroups",
		Method:             "POST",
		PathPattern:        "/networks/{network_id}/rating_groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostNetworksNetworkIDRatingGroupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostNetworksNetworkIDRatingGroupsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostNetworksNetworkIDRatingGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PutNetworksNetworkIDRatingGroupsRatingGroupID modifies a rating group
*/
func (a *Client) PutNetworksNetworkIDRatingGroupsRatingGroupID(params *PutNetworksNetworkIDRatingGroupsRatingGroupIDParams) (*PutNetworksNetworkIDRatingGroupsRatingGroupIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutNetworksNetworkIDRatingGroupsRatingGroupIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutNetworksNetworkIDRatingGroupsRatingGroupID",
		Method:             "PUT",
		PathPattern:        "/networks/{network_id}/rating_groups/{rating_group_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutNetworksNetworkIDRatingGroupsRatingGroupIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutNetworksNetworkIDRatingGroupsRatingGroupIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutNetworksNetworkIDRatingGroupsRatingGroupIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}