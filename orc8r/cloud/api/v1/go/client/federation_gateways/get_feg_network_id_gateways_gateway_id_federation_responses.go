// Code generated by go-swagger; DO NOT EDIT.

package federation_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// GetFegNetworkIDGatewaysGatewayIDFederationReader is a Reader for the GetFegNetworkIDGatewaysGatewayIDFederation structure.
type GetFegNetworkIDGatewaysGatewayIDFederationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFegNetworkIDGatewaysGatewayIDFederationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFegNetworkIDGatewaysGatewayIDFederationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFegNetworkIDGatewaysGatewayIDFederationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFegNetworkIDGatewaysGatewayIDFederationOK creates a GetFegNetworkIDGatewaysGatewayIDFederationOK with default headers values
func NewGetFegNetworkIDGatewaysGatewayIDFederationOK() *GetFegNetworkIDGatewaysGatewayIDFederationOK {
	return &GetFegNetworkIDGatewaysGatewayIDFederationOK{}
}

/*
GetFegNetworkIDGatewaysGatewayIDFederationOK describes a response with status code 200, with default header values.

Retrieved Gateway FeG Configs
*/
type GetFegNetworkIDGatewaysGatewayIDFederationOK struct {
	Payload *models.GatewayFederationConfigs
}

func (o *GetFegNetworkIDGatewaysGatewayIDFederationOK) Error() string {
	return fmt.Sprintf("[GET /feg/{network_id}/gateways/{gateway_id}/federation][%d] getFegNetworkIdGatewaysGatewayIdFederationOK  %+v", 200, o.Payload)
}
func (o *GetFegNetworkIDGatewaysGatewayIDFederationOK) GetPayload() *models.GatewayFederationConfigs {
	return o.Payload
}

func (o *GetFegNetworkIDGatewaysGatewayIDFederationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayFederationConfigs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFegNetworkIDGatewaysGatewayIDFederationDefault creates a GetFegNetworkIDGatewaysGatewayIDFederationDefault with default headers values
func NewGetFegNetworkIDGatewaysGatewayIDFederationDefault(code int) *GetFegNetworkIDGatewaysGatewayIDFederationDefault {
	return &GetFegNetworkIDGatewaysGatewayIDFederationDefault{
		_statusCode: code,
	}
}

/*
GetFegNetworkIDGatewaysGatewayIDFederationDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetFegNetworkIDGatewaysGatewayIDFederationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get feg network ID gateways gateway ID federation default response
func (o *GetFegNetworkIDGatewaysGatewayIDFederationDefault) Code() int {
	return o._statusCode
}

func (o *GetFegNetworkIDGatewaysGatewayIDFederationDefault) Error() string {
	return fmt.Sprintf("[GET /feg/{network_id}/gateways/{gateway_id}/federation][%d] GetFegNetworkIDGatewaysGatewayIDFederation default  %+v", o._statusCode, o.Payload)
}
func (o *GetFegNetworkIDGatewaysGatewayIDFederationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFegNetworkIDGatewaysGatewayIDFederationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
