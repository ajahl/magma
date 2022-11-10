// Code generated by go-swagger; DO NOT EDIT.

package sms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// DeleteLTENetworkIDSMSSMSPkReader is a Reader for the DeleteLTENetworkIDSMSSMSPk structure.
type DeleteLTENetworkIDSMSSMSPkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLTENetworkIDSMSSMSPkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLTENetworkIDSMSSMSPkNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteLTENetworkIDSMSSMSPkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteLTENetworkIDSMSSMSPkNoContent creates a DeleteLTENetworkIDSMSSMSPkNoContent with default headers values
func NewDeleteLTENetworkIDSMSSMSPkNoContent() *DeleteLTENetworkIDSMSSMSPkNoContent {
	return &DeleteLTENetworkIDSMSSMSPkNoContent{}
}

/*
DeleteLTENetworkIDSMSSMSPkNoContent describes a response with status code 204, with default header values.

Success
*/
type DeleteLTENetworkIDSMSSMSPkNoContent struct {
}

func (o *DeleteLTENetworkIDSMSSMSPkNoContent) Error() string {
	return fmt.Sprintf("[DELETE /lte/{network_id}/sms/{sms_pk}][%d] deleteLteNetworkIdSmsSmsPkNoContent ", 204)
}

func (o *DeleteLTENetworkIDSMSSMSPkNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLTENetworkIDSMSSMSPkDefault creates a DeleteLTENetworkIDSMSSMSPkDefault with default headers values
func NewDeleteLTENetworkIDSMSSMSPkDefault(code int) *DeleteLTENetworkIDSMSSMSPkDefault {
	return &DeleteLTENetworkIDSMSSMSPkDefault{
		_statusCode: code,
	}
}

/*
DeleteLTENetworkIDSMSSMSPkDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type DeleteLTENetworkIDSMSSMSPkDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete LTE network ID SMS SMS pk default response
func (o *DeleteLTENetworkIDSMSSMSPkDefault) Code() int {
	return o._statusCode
}

func (o *DeleteLTENetworkIDSMSSMSPkDefault) Error() string {
	return fmt.Sprintf("[DELETE /lte/{network_id}/sms/{sms_pk}][%d] DeleteLTENetworkIDSMSSMSPk default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteLTENetworkIDSMSSMSPkDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteLTENetworkIDSMSSMSPkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
