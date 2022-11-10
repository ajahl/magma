// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// GetNetworksNetworkIDNameReader is a Reader for the GetNetworksNetworkIDName structure.
type GetNetworksNetworkIDNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDNameOK creates a GetNetworksNetworkIDNameOK with default headers values
func NewGetNetworksNetworkIDNameOK() *GetNetworksNetworkIDNameOK {
	return &GetNetworksNetworkIDNameOK{}
}

/*
GetNetworksNetworkIDNameOK describes a response with status code 200, with default header values.

Name of the network
*/
type GetNetworksNetworkIDNameOK struct {
	Payload models.NetworkName
}

func (o *GetNetworksNetworkIDNameOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/name][%d] getNetworksNetworkIdNameOK  %+v", 200, o.Payload)
}
func (o *GetNetworksNetworkIDNameOK) GetPayload() models.NetworkName {
	return o.Payload
}

func (o *GetNetworksNetworkIDNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDNameDefault creates a GetNetworksNetworkIDNameDefault with default headers values
func NewGetNetworksNetworkIDNameDefault(code int) *GetNetworksNetworkIDNameDefault {
	return &GetNetworksNetworkIDNameDefault{
		_statusCode: code,
	}
}

/*
GetNetworksNetworkIDNameDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID name default response
func (o *GetNetworksNetworkIDNameDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDNameDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/name][%d] GetNetworksNetworkIDName default  %+v", o._statusCode, o.Payload)
}
func (o *GetNetworksNetworkIDNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
