// Code generated by go-swagger; DO NOT EDIT.

package lte_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// GetLTENetworkIDGatewaysGatewayIDTierReader is a Reader for the GetLTENetworkIDGatewaysGatewayIDTier structure.
type GetLTENetworkIDGatewaysGatewayIDTierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDGatewaysGatewayIDTierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDGatewaysGatewayIDTierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDGatewaysGatewayIDTierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDTierOK creates a GetLTENetworkIDGatewaysGatewayIDTierOK with default headers values
func NewGetLTENetworkIDGatewaysGatewayIDTierOK() *GetLTENetworkIDGatewaysGatewayIDTierOK {
	return &GetLTENetworkIDGatewaysGatewayIDTierOK{}
}

/*
GetLTENetworkIDGatewaysGatewayIDTierOK describes a response with status code 200, with default header values.

The ID of the upgrade tier
*/
type GetLTENetworkIDGatewaysGatewayIDTierOK struct {
	Payload models.TierID
}

func (o *GetLTENetworkIDGatewaysGatewayIDTierOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways/{gateway_id}/tier][%d] getLteNetworkIdGatewaysGatewayIdTierOK  %+v", 200, o.Payload)
}
func (o *GetLTENetworkIDGatewaysGatewayIDTierOK) GetPayload() models.TierID {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysGatewayIDTierOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDGatewaysGatewayIDTierDefault creates a GetLTENetworkIDGatewaysGatewayIDTierDefault with default headers values
func NewGetLTENetworkIDGatewaysGatewayIDTierDefault(code int) *GetLTENetworkIDGatewaysGatewayIDTierDefault {
	return &GetLTENetworkIDGatewaysGatewayIDTierDefault{
		_statusCode: code,
	}
}

/*
GetLTENetworkIDGatewaysGatewayIDTierDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetLTENetworkIDGatewaysGatewayIDTierDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID gateways gateway ID tier default response
func (o *GetLTENetworkIDGatewaysGatewayIDTierDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDGatewaysGatewayIDTierDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways/{gateway_id}/tier][%d] GetLTENetworkIDGatewaysGatewayIDTier default  %+v", o._statusCode, o.Payload)
}
func (o *GetLTENetworkIDGatewaysGatewayIDTierDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysGatewayIDTierDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
