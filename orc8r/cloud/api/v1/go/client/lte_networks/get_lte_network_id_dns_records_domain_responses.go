// Code generated by go-swagger; DO NOT EDIT.

package lte_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// GetLTENetworkIDDNSRecordsDomainReader is a Reader for the GetLTENetworkIDDNSRecordsDomain structure.
type GetLTENetworkIDDNSRecordsDomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDDNSRecordsDomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDDNSRecordsDomainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDDNSRecordsDomainDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDDNSRecordsDomainOK creates a GetLTENetworkIDDNSRecordsDomainOK with default headers values
func NewGetLTENetworkIDDNSRecordsDomainOK() *GetLTENetworkIDDNSRecordsDomainOK {
	return &GetLTENetworkIDDNSRecordsDomainOK{}
}

/*
GetLTENetworkIDDNSRecordsDomainOK describes a response with status code 200, with default header values.

DNS config record
*/
type GetLTENetworkIDDNSRecordsDomainOK struct {
	Payload *models.DNSConfigRecord
}

func (o *GetLTENetworkIDDNSRecordsDomainOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/dns/records/{domain}][%d] getLteNetworkIdDnsRecordsDomainOK  %+v", 200, o.Payload)
}
func (o *GetLTENetworkIDDNSRecordsDomainOK) GetPayload() *models.DNSConfigRecord {
	return o.Payload
}

func (o *GetLTENetworkIDDNSRecordsDomainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DNSConfigRecord)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDDNSRecordsDomainDefault creates a GetLTENetworkIDDNSRecordsDomainDefault with default headers values
func NewGetLTENetworkIDDNSRecordsDomainDefault(code int) *GetLTENetworkIDDNSRecordsDomainDefault {
	return &GetLTENetworkIDDNSRecordsDomainDefault{
		_statusCode: code,
	}
}

/*
GetLTENetworkIDDNSRecordsDomainDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetLTENetworkIDDNSRecordsDomainDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID DNS records domain default response
func (o *GetLTENetworkIDDNSRecordsDomainDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDDNSRecordsDomainDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/dns/records/{domain}][%d] GetLTENetworkIDDNSRecordsDomain default  %+v", o._statusCode, o.Payload)
}
func (o *GetLTENetworkIDDNSRecordsDomainDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDDNSRecordsDomainDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
