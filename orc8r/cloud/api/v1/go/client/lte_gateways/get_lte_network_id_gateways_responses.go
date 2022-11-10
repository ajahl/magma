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

// GetLTENetworkIDGatewaysReader is a Reader for the GetLTENetworkIDGateways structure.
type GetLTENetworkIDGatewaysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDGatewaysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDGatewaysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDGatewaysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDGatewaysOK creates a GetLTENetworkIDGatewaysOK with default headers values
func NewGetLTENetworkIDGatewaysOK() *GetLTENetworkIDGatewaysOK {
	return &GetLTENetworkIDGatewaysOK{}
}

/*
GetLTENetworkIDGatewaysOK describes a response with status code 200, with default header values.

Map of all LTE gateways inside the network by gatewayID
*/
type GetLTENetworkIDGatewaysOK struct {
	Payload map[string]models.LTEGateway
}

func (o *GetLTENetworkIDGatewaysOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways][%d] getLteNetworkIdGatewaysOK  %+v", 200, o.Payload)
}
func (o *GetLTENetworkIDGatewaysOK) GetPayload() map[string]models.LTEGateway {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDGatewaysDefault creates a GetLTENetworkIDGatewaysDefault with default headers values
func NewGetLTENetworkIDGatewaysDefault(code int) *GetLTENetworkIDGatewaysDefault {
	return &GetLTENetworkIDGatewaysDefault{
		_statusCode: code,
	}
}

/*
GetLTENetworkIDGatewaysDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetLTENetworkIDGatewaysDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID gateways default response
func (o *GetLTENetworkIDGatewaysDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDGatewaysDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways][%d] GetLTENetworkIDGateways default  %+v", o._statusCode, o.Payload)
}
func (o *GetLTENetworkIDGatewaysDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
