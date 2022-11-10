// Code generated by go-swagger; DO NOT EDIT.

package rating_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// GetNetworksNetworkIDRatingGroupsRatingGroupIDReader is a Reader for the GetNetworksNetworkIDRatingGroupsRatingGroupID structure.
type GetNetworksNetworkIDRatingGroupsRatingGroupIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDRatingGroupsRatingGroupIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDRatingGroupsRatingGroupIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDRatingGroupsRatingGroupIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDRatingGroupsRatingGroupIDOK creates a GetNetworksNetworkIDRatingGroupsRatingGroupIDOK with default headers values
func NewGetNetworksNetworkIDRatingGroupsRatingGroupIDOK() *GetNetworksNetworkIDRatingGroupsRatingGroupIDOK {
	return &GetNetworksNetworkIDRatingGroupsRatingGroupIDOK{}
}

/*
GetNetworksNetworkIDRatingGroupsRatingGroupIDOK describes a response with status code 200, with default header values.

Rating group on success
*/
type GetNetworksNetworkIDRatingGroupsRatingGroupIDOK struct {
	Payload *models.RatingGroup
}

func (o *GetNetworksNetworkIDRatingGroupsRatingGroupIDOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/rating_groups/{rating_group_id}][%d] getNetworksNetworkIdRatingGroupsRatingGroupIdOK  %+v", 200, o.Payload)
}
func (o *GetNetworksNetworkIDRatingGroupsRatingGroupIDOK) GetPayload() *models.RatingGroup {
	return o.Payload
}

func (o *GetNetworksNetworkIDRatingGroupsRatingGroupIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RatingGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDRatingGroupsRatingGroupIDDefault creates a GetNetworksNetworkIDRatingGroupsRatingGroupIDDefault with default headers values
func NewGetNetworksNetworkIDRatingGroupsRatingGroupIDDefault(code int) *GetNetworksNetworkIDRatingGroupsRatingGroupIDDefault {
	return &GetNetworksNetworkIDRatingGroupsRatingGroupIDDefault{
		_statusCode: code,
	}
}

/*
GetNetworksNetworkIDRatingGroupsRatingGroupIDDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDRatingGroupsRatingGroupIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID rating groups rating group ID default response
func (o *GetNetworksNetworkIDRatingGroupsRatingGroupIDDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDRatingGroupsRatingGroupIDDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/rating_groups/{rating_group_id}][%d] GetNetworksNetworkIDRatingGroupsRatingGroupID default  %+v", o._statusCode, o.Payload)
}
func (o *GetNetworksNetworkIDRatingGroupsRatingGroupIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDRatingGroupsRatingGroupIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
