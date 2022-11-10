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

// GetNetworksReader is a Reader for the GetNetworks structure.
type GetNetworksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksOK creates a GetNetworksOK with default headers values
func NewGetNetworksOK() *GetNetworksOK {
	return &GetNetworksOK{}
}

/*
GetNetworksOK describes a response with status code 200, with default header values.

List of network IDs
*/
type GetNetworksOK struct {
	Payload []string
}

func (o *GetNetworksOK) Error() string {
	return fmt.Sprintf("[GET /networks][%d] getNetworksOK  %+v", 200, o.Payload)
}
func (o *GetNetworksOK) GetPayload() []string {
	return o.Payload
}

func (o *GetNetworksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksDefault creates a GetNetworksDefault with default headers values
func NewGetNetworksDefault(code int) *GetNetworksDefault {
	return &GetNetworksDefault{
		_statusCode: code,
	}
}

/*
GetNetworksDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetNetworksDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks default response
func (o *GetNetworksDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksDefault) Error() string {
	return fmt.Sprintf("[GET /networks][%d] GetNetworks default  %+v", o._statusCode, o.Payload)
}
func (o *GetNetworksDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
