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

// PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameReader is a Reader for the PostCwfNetworkIDSubscriberConfigBaseNamesBaseName structure.
type PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostCwfNetworkIDSubscriberConfigBaseNamesBaseNameCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostCwfNetworkIDSubscriberConfigBaseNamesBaseNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCwfNetworkIDSubscriberConfigBaseNamesBaseNameCreated creates a PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameCreated with default headers values
func NewPostCwfNetworkIDSubscriberConfigBaseNamesBaseNameCreated() *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameCreated {
	return &PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameCreated{}
}

/*
PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameCreated describes a response with status code 201, with default header values.

Success
*/
type PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameCreated struct {
}

func (o *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameCreated) Error() string {
	return fmt.Sprintf("[POST /cwf/{network_id}/subscriber_config/base_names/{base_name}][%d] postCwfNetworkIdSubscriberConfigBaseNamesBaseNameCreated ", 201)
}

func (o *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCwfNetworkIDSubscriberConfigBaseNamesBaseNameDefault creates a PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameDefault with default headers values
func NewPostCwfNetworkIDSubscriberConfigBaseNamesBaseNameDefault(code int) *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameDefault {
	return &PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameDefault{
		_statusCode: code,
	}
}

/*
PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post cwf network ID subscriber config base names base name default response
func (o *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameDefault) Code() int {
	return o._statusCode
}

func (o *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameDefault) Error() string {
	return fmt.Sprintf("[POST /cwf/{network_id}/subscriber_config/base_names/{base_name}][%d] PostCwfNetworkIDSubscriberConfigBaseNamesBaseName default  %+v", o._statusCode, o.Payload)
}
func (o *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
