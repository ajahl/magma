// Code generated by go-swagger; DO NOT EDIT.

package lte_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// GetLTENetworkIDGatewayPoolsGatewayPoolIDReader is a Reader for the GetLTENetworkIDGatewayPoolsGatewayPoolID structure.
type GetLTENetworkIDGatewayPoolsGatewayPoolIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDGatewayPoolsGatewayPoolIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDGatewayPoolsGatewayPoolIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDGatewayPoolsGatewayPoolIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDGatewayPoolsGatewayPoolIDOK creates a GetLTENetworkIDGatewayPoolsGatewayPoolIDOK with default headers values
func NewGetLTENetworkIDGatewayPoolsGatewayPoolIDOK() *GetLTENetworkIDGatewayPoolsGatewayPoolIDOK {
	return &GetLTENetworkIDGatewayPoolsGatewayPoolIDOK{}
}

/*
GetLTENetworkIDGatewayPoolsGatewayPoolIDOK describes a response with status code 200, with default header values.

HA gateway pool
*/
type GetLTENetworkIDGatewayPoolsGatewayPoolIDOK struct {
	Payload *models.CellularGatewayPool
}

func (o *GetLTENetworkIDGatewayPoolsGatewayPoolIDOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateway_pools/{gateway_pool_id}][%d] getLteNetworkIdGatewayPoolsGatewayPoolIdOK  %+v", 200, o.Payload)
}
func (o *GetLTENetworkIDGatewayPoolsGatewayPoolIDOK) GetPayload() *models.CellularGatewayPool {
	return o.Payload
}

func (o *GetLTENetworkIDGatewayPoolsGatewayPoolIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CellularGatewayPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDGatewayPoolsGatewayPoolIDDefault creates a GetLTENetworkIDGatewayPoolsGatewayPoolIDDefault with default headers values
func NewGetLTENetworkIDGatewayPoolsGatewayPoolIDDefault(code int) *GetLTENetworkIDGatewayPoolsGatewayPoolIDDefault {
	return &GetLTENetworkIDGatewayPoolsGatewayPoolIDDefault{
		_statusCode: code,
	}
}

/*
GetLTENetworkIDGatewayPoolsGatewayPoolIDDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetLTENetworkIDGatewayPoolsGatewayPoolIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID gateway pools gateway pool ID default response
func (o *GetLTENetworkIDGatewayPoolsGatewayPoolIDDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDGatewayPoolsGatewayPoolIDDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateway_pools/{gateway_pool_id}][%d] GetLTENetworkIDGatewayPoolsGatewayPoolID default  %+v", o._statusCode, o.Payload)
}
func (o *GetLTENetworkIDGatewayPoolsGatewayPoolIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDGatewayPoolsGatewayPoolIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
