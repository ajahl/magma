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

// DeleteNetworksNetworkIDReader is a Reader for the DeleteNetworksNetworkID structure.
type DeleteNetworksNetworkIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworksNetworkIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworksNetworkIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteNetworksNetworkIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNetworksNetworkIDNoContent creates a DeleteNetworksNetworkIDNoContent with default headers values
func NewDeleteNetworksNetworkIDNoContent() *DeleteNetworksNetworkIDNoContent {
	return &DeleteNetworksNetworkIDNoContent{}
}

/*
DeleteNetworksNetworkIDNoContent describes a response with status code 204, with default header values.

Success
*/
type DeleteNetworksNetworkIDNoContent struct {
}

func (o *DeleteNetworksNetworkIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /networks/{network_id}][%d] deleteNetworksNetworkIdNoContent ", 204)
}

func (o *DeleteNetworksNetworkIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNetworksNetworkIDDefault creates a DeleteNetworksNetworkIDDefault with default headers values
func NewDeleteNetworksNetworkIDDefault(code int) *DeleteNetworksNetworkIDDefault {
	return &DeleteNetworksNetworkIDDefault{
		_statusCode: code,
	}
}

/*
DeleteNetworksNetworkIDDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type DeleteNetworksNetworkIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete networks network ID default response
func (o *DeleteNetworksNetworkIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNetworksNetworkIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /networks/{network_id}][%d] DeleteNetworksNetworkID default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteNetworksNetworkIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteNetworksNetworkIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
