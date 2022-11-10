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

// GetCwfNetworkIDCarrierWifiReader is a Reader for the GetCwfNetworkIDCarrierWifi structure.
type GetCwfNetworkIDCarrierWifiReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCwfNetworkIDCarrierWifiReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCwfNetworkIDCarrierWifiOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCwfNetworkIDCarrierWifiDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCwfNetworkIDCarrierWifiOK creates a GetCwfNetworkIDCarrierWifiOK with default headers values
func NewGetCwfNetworkIDCarrierWifiOK() *GetCwfNetworkIDCarrierWifiOK {
	return &GetCwfNetworkIDCarrierWifiOK{}
}

/*
GetCwfNetworkIDCarrierWifiOK describes a response with status code 200, with default header values.

Retrieve Network Carrier Wifi Configs
*/
type GetCwfNetworkIDCarrierWifiOK struct {
	Payload *models.NetworkCarrierWifiConfigs
}

func (o *GetCwfNetworkIDCarrierWifiOK) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/carrier_wifi][%d] getCwfNetworkIdCarrierWifiOK  %+v", 200, o.Payload)
}
func (o *GetCwfNetworkIDCarrierWifiOK) GetPayload() *models.NetworkCarrierWifiConfigs {
	return o.Payload
}

func (o *GetCwfNetworkIDCarrierWifiOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkCarrierWifiConfigs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCwfNetworkIDCarrierWifiDefault creates a GetCwfNetworkIDCarrierWifiDefault with default headers values
func NewGetCwfNetworkIDCarrierWifiDefault(code int) *GetCwfNetworkIDCarrierWifiDefault {
	return &GetCwfNetworkIDCarrierWifiDefault{
		_statusCode: code,
	}
}

/*
GetCwfNetworkIDCarrierWifiDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetCwfNetworkIDCarrierWifiDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cwf network ID carrier wifi default response
func (o *GetCwfNetworkIDCarrierWifiDefault) Code() int {
	return o._statusCode
}

func (o *GetCwfNetworkIDCarrierWifiDefault) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/carrier_wifi][%d] GetCwfNetworkIDCarrierWifi default  %+v", o._statusCode, o.Payload)
}
func (o *GetCwfNetworkIDCarrierWifiDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCwfNetworkIDCarrierWifiDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
