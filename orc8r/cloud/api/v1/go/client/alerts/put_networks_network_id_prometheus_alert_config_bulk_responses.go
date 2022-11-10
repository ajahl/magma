// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// PutNetworksNetworkIDPrometheusAlertConfigBulkReader is a Reader for the PutNetworksNetworkIDPrometheusAlertConfigBulk structure.
type PutNetworksNetworkIDPrometheusAlertConfigBulkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNetworksNetworkIDPrometheusAlertConfigBulkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutNetworksNetworkIDPrometheusAlertConfigBulkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutNetworksNetworkIDPrometheusAlertConfigBulkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutNetworksNetworkIDPrometheusAlertConfigBulkOK creates a PutNetworksNetworkIDPrometheusAlertConfigBulkOK with default headers values
func NewPutNetworksNetworkIDPrometheusAlertConfigBulkOK() *PutNetworksNetworkIDPrometheusAlertConfigBulkOK {
	return &PutNetworksNetworkIDPrometheusAlertConfigBulkOK{}
}

/*
PutNetworksNetworkIDPrometheusAlertConfigBulkOK describes a response with status code 200, with default header values.

Success
*/
type PutNetworksNetworkIDPrometheusAlertConfigBulkOK struct {
	Payload *models.AlertBulkUploadResponse
}

func (o *PutNetworksNetworkIDPrometheusAlertConfigBulkOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/prometheus/alert_config/bulk][%d] putNetworksNetworkIdPrometheusAlertConfigBulkOK  %+v", 200, o.Payload)
}
func (o *PutNetworksNetworkIDPrometheusAlertConfigBulkOK) GetPayload() *models.AlertBulkUploadResponse {
	return o.Payload
}

func (o *PutNetworksNetworkIDPrometheusAlertConfigBulkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertBulkUploadResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNetworksNetworkIDPrometheusAlertConfigBulkDefault creates a PutNetworksNetworkIDPrometheusAlertConfigBulkDefault with default headers values
func NewPutNetworksNetworkIDPrometheusAlertConfigBulkDefault(code int) *PutNetworksNetworkIDPrometheusAlertConfigBulkDefault {
	return &PutNetworksNetworkIDPrometheusAlertConfigBulkDefault{
		_statusCode: code,
	}
}

/*
PutNetworksNetworkIDPrometheusAlertConfigBulkDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PutNetworksNetworkIDPrometheusAlertConfigBulkDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put networks network ID prometheus alert config bulk default response
func (o *PutNetworksNetworkIDPrometheusAlertConfigBulkDefault) Code() int {
	return o._statusCode
}

func (o *PutNetworksNetworkIDPrometheusAlertConfigBulkDefault) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/prometheus/alert_config/bulk][%d] PutNetworksNetworkIDPrometheusAlertConfigBulk default  %+v", o._statusCode, o.Payload)
}
func (o *PutNetworksNetworkIDPrometheusAlertConfigBulkDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNetworksNetworkIDPrometheusAlertConfigBulkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
