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

// GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsReader is a Reader for the GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerials structure.
type GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsOK creates a GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsOK with default headers values
func NewGetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsOK() *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsOK {
	return &GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsOK{}
}

/*
GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsOK describes a response with status code 200, with default header values.

The SNs of all enodeBs
*/
type GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsOK struct {
	Payload models.ENODEBSerials
}

func (o *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways/{gateway_id}/connected_enodeb_serials][%d] getLteNetworkIdGatewaysGatewayIdConnectedEnodebSerialsOK  %+v", 200, o.Payload)
}
func (o *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsOK) GetPayload() models.ENODEBSerials {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsDefault creates a GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsDefault with default headers values
func NewGetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsDefault(code int) *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsDefault {
	return &GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsDefault{
		_statusCode: code,
	}
}

/*
GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID gateways gateway ID connected ENODEB serials default response
func (o *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways/{gateway_id}/connected_enodeb_serials][%d] GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerials default  %+v", o._statusCode, o.Payload)
}
func (o *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
