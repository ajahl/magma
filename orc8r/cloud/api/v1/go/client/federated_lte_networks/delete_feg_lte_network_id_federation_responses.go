// Code generated by go-swagger; DO NOT EDIT.

package federated_lte_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// DeleteFegLTENetworkIDFederationReader is a Reader for the DeleteFegLTENetworkIDFederation structure.
type DeleteFegLTENetworkIDFederationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFegLTENetworkIDFederationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteFegLTENetworkIDFederationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteFegLTENetworkIDFederationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteFegLTENetworkIDFederationNoContent creates a DeleteFegLTENetworkIDFederationNoContent with default headers values
func NewDeleteFegLTENetworkIDFederationNoContent() *DeleteFegLTENetworkIDFederationNoContent {
	return &DeleteFegLTENetworkIDFederationNoContent{}
}

/*
DeleteFegLTENetworkIDFederationNoContent describes a response with status code 204, with default header values.

Success
*/
type DeleteFegLTENetworkIDFederationNoContent struct {
}

func (o *DeleteFegLTENetworkIDFederationNoContent) Error() string {
	return fmt.Sprintf("[DELETE /feg_lte/{network_id}/federation][%d] deleteFegLteNetworkIdFederationNoContent ", 204)
}

func (o *DeleteFegLTENetworkIDFederationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFegLTENetworkIDFederationDefault creates a DeleteFegLTENetworkIDFederationDefault with default headers values
func NewDeleteFegLTENetworkIDFederationDefault(code int) *DeleteFegLTENetworkIDFederationDefault {
	return &DeleteFegLTENetworkIDFederationDefault{
		_statusCode: code,
	}
}

/*
DeleteFegLTENetworkIDFederationDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type DeleteFegLTENetworkIDFederationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete feg LTE network ID federation default response
func (o *DeleteFegLTENetworkIDFederationDefault) Code() int {
	return o._statusCode
}

func (o *DeleteFegLTENetworkIDFederationDefault) Error() string {
	return fmt.Sprintf("[DELETE /feg_lte/{network_id}/federation][%d] DeleteFegLTENetworkIDFederation default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteFegLTENetworkIDFederationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteFegLTENetworkIDFederationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
