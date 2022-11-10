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

// PostFegNetworkIDGatewaysReader is a Reader for the PostFegNetworkIDGateways structure.
type PostFegNetworkIDGatewaysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFegNetworkIDGatewaysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostFegNetworkIDGatewaysCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostFegNetworkIDGatewaysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostFegNetworkIDGatewaysCreated creates a PostFegNetworkIDGatewaysCreated with default headers values
func NewPostFegNetworkIDGatewaysCreated() *PostFegNetworkIDGatewaysCreated {
	return &PostFegNetworkIDGatewaysCreated{}
}

/*
PostFegNetworkIDGatewaysCreated describes a response with status code 201, with default header values.

Success
*/
type PostFegNetworkIDGatewaysCreated struct {
}

func (o *PostFegNetworkIDGatewaysCreated) Error() string {
	return fmt.Sprintf("[POST /feg/{network_id}/gateways][%d] postFegNetworkIdGatewaysCreated ", 201)
}

func (o *PostFegNetworkIDGatewaysCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostFegNetworkIDGatewaysDefault creates a PostFegNetworkIDGatewaysDefault with default headers values
func NewPostFegNetworkIDGatewaysDefault(code int) *PostFegNetworkIDGatewaysDefault {
	return &PostFegNetworkIDGatewaysDefault{
		_statusCode: code,
	}
}

/*
PostFegNetworkIDGatewaysDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PostFegNetworkIDGatewaysDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post feg network ID gateways default response
func (o *PostFegNetworkIDGatewaysDefault) Code() int {
	return o._statusCode
}

func (o *PostFegNetworkIDGatewaysDefault) Error() string {
	return fmt.Sprintf("[POST /feg/{network_id}/gateways][%d] PostFegNetworkIDGateways default  %+v", o._statusCode, o.Payload)
}
func (o *PostFegNetworkIDGatewaysDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostFegNetworkIDGatewaysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
