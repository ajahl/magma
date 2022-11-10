// Code generated by go-swagger; DO NOT EDIT.

package federation_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// GetFegNetworkIDSubscriberConfigReader is a Reader for the GetFegNetworkIDSubscriberConfig structure.
type GetFegNetworkIDSubscriberConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFegNetworkIDSubscriberConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFegNetworkIDSubscriberConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFegNetworkIDSubscriberConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFegNetworkIDSubscriberConfigOK creates a GetFegNetworkIDSubscriberConfigOK with default headers values
func NewGetFegNetworkIDSubscriberConfigOK() *GetFegNetworkIDSubscriberConfigOK {
	return &GetFegNetworkIDSubscriberConfigOK{}
}

/*
GetFegNetworkIDSubscriberConfigOK describes a response with status code 200, with default header values.

Subscriber Config
*/
type GetFegNetworkIDSubscriberConfigOK struct {
	Payload *models.NetworkSubscriberConfig
}

func (o *GetFegNetworkIDSubscriberConfigOK) Error() string {
	return fmt.Sprintf("[GET /feg/{network_id}/subscriber_config][%d] getFegNetworkIdSubscriberConfigOK  %+v", 200, o.Payload)
}
func (o *GetFegNetworkIDSubscriberConfigOK) GetPayload() *models.NetworkSubscriberConfig {
	return o.Payload
}

func (o *GetFegNetworkIDSubscriberConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkSubscriberConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFegNetworkIDSubscriberConfigDefault creates a GetFegNetworkIDSubscriberConfigDefault with default headers values
func NewGetFegNetworkIDSubscriberConfigDefault(code int) *GetFegNetworkIDSubscriberConfigDefault {
	return &GetFegNetworkIDSubscriberConfigDefault{
		_statusCode: code,
	}
}

/*
GetFegNetworkIDSubscriberConfigDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetFegNetworkIDSubscriberConfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get feg network ID subscriber config default response
func (o *GetFegNetworkIDSubscriberConfigDefault) Code() int {
	return o._statusCode
}

func (o *GetFegNetworkIDSubscriberConfigDefault) Error() string {
	return fmt.Sprintf("[GET /feg/{network_id}/subscriber_config][%d] GetFegNetworkIDSubscriberConfig default  %+v", o._statusCode, o.Payload)
}
func (o *GetFegNetworkIDSubscriberConfigDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFegNetworkIDSubscriberConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
