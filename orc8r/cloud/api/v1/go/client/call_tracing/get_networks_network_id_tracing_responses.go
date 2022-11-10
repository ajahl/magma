// Code generated by go-swagger; DO NOT EDIT.

package call_tracing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// GetNetworksNetworkIDTracingReader is a Reader for the GetNetworksNetworkIDTracing structure.
type GetNetworksNetworkIDTracingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDTracingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDTracingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDTracingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDTracingOK creates a GetNetworksNetworkIDTracingOK with default headers values
func NewGetNetworksNetworkIDTracingOK() *GetNetworksNetworkIDTracingOK {
	return &GetNetworksNetworkIDTracingOK{}
}

/*
GetNetworksNetworkIDTracingOK describes a response with status code 200, with default header values.

Success
*/
type GetNetworksNetworkIDTracingOK struct {
	Payload map[string]models.CallTrace
}

func (o *GetNetworksNetworkIDTracingOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/tracing][%d] getNetworksNetworkIdTracingOK  %+v", 200, o.Payload)
}
func (o *GetNetworksNetworkIDTracingOK) GetPayload() map[string]models.CallTrace {
	return o.Payload
}

func (o *GetNetworksNetworkIDTracingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDTracingDefault creates a GetNetworksNetworkIDTracingDefault with default headers values
func NewGetNetworksNetworkIDTracingDefault(code int) *GetNetworksNetworkIDTracingDefault {
	return &GetNetworksNetworkIDTracingDefault{
		_statusCode: code,
	}
}

/*
GetNetworksNetworkIDTracingDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDTracingDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID tracing default response
func (o *GetNetworksNetworkIDTracingDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDTracingDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/tracing][%d] GetNetworksNetworkIDTracing default  %+v", o._statusCode, o.Payload)
}
func (o *GetNetworksNetworkIDTracingDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDTracingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
