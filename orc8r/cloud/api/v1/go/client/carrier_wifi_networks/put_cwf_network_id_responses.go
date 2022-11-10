// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// PutCwfNetworkIDReader is a Reader for the PutCwfNetworkID structure.
type PutCwfNetworkIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCwfNetworkIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutCwfNetworkIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCwfNetworkIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCwfNetworkIDNoContent creates a PutCwfNetworkIDNoContent with default headers values
func NewPutCwfNetworkIDNoContent() *PutCwfNetworkIDNoContent {
	return &PutCwfNetworkIDNoContent{}
}

/*
PutCwfNetworkIDNoContent describes a response with status code 204, with default header values.

Success
*/
type PutCwfNetworkIDNoContent struct {
}

func (o *PutCwfNetworkIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /cwf/{network_id}][%d] putCwfNetworkIdNoContent ", 204)
}

func (o *PutCwfNetworkIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCwfNetworkIDDefault creates a PutCwfNetworkIDDefault with default headers values
func NewPutCwfNetworkIDDefault(code int) *PutCwfNetworkIDDefault {
	return &PutCwfNetworkIDDefault{
		_statusCode: code,
	}
}

/*
PutCwfNetworkIDDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PutCwfNetworkIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put cwf network ID default response
func (o *PutCwfNetworkIDDefault) Code() int {
	return o._statusCode
}

func (o *PutCwfNetworkIDDefault) Error() string {
	return fmt.Sprintf("[PUT /cwf/{network_id}][%d] PutCwfNetworkID default  %+v", o._statusCode, o.Payload)
}
func (o *PutCwfNetworkIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutCwfNetworkIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
