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

// PostNetworksNetworkIDDNSRecordsDomainReader is a Reader for the PostNetworksNetworkIDDNSRecordsDomain structure.
type PostNetworksNetworkIDDNSRecordsDomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNetworksNetworkIDDNSRecordsDomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostNetworksNetworkIDDNSRecordsDomainCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostNetworksNetworkIDDNSRecordsDomainDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostNetworksNetworkIDDNSRecordsDomainCreated creates a PostNetworksNetworkIDDNSRecordsDomainCreated with default headers values
func NewPostNetworksNetworkIDDNSRecordsDomainCreated() *PostNetworksNetworkIDDNSRecordsDomainCreated {
	return &PostNetworksNetworkIDDNSRecordsDomainCreated{}
}

/*
PostNetworksNetworkIDDNSRecordsDomainCreated describes a response with status code 201, with default header values.

Success
*/
type PostNetworksNetworkIDDNSRecordsDomainCreated struct {
}

func (o *PostNetworksNetworkIDDNSRecordsDomainCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{network_id}/dns/records/{domain}][%d] postNetworksNetworkIdDnsRecordsDomainCreated ", 201)
}

func (o *PostNetworksNetworkIDDNSRecordsDomainCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostNetworksNetworkIDDNSRecordsDomainDefault creates a PostNetworksNetworkIDDNSRecordsDomainDefault with default headers values
func NewPostNetworksNetworkIDDNSRecordsDomainDefault(code int) *PostNetworksNetworkIDDNSRecordsDomainDefault {
	return &PostNetworksNetworkIDDNSRecordsDomainDefault{
		_statusCode: code,
	}
}

/*
PostNetworksNetworkIDDNSRecordsDomainDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PostNetworksNetworkIDDNSRecordsDomainDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post networks network ID DNS records domain default response
func (o *PostNetworksNetworkIDDNSRecordsDomainDefault) Code() int {
	return o._statusCode
}

func (o *PostNetworksNetworkIDDNSRecordsDomainDefault) Error() string {
	return fmt.Sprintf("[POST /networks/{network_id}/dns/records/{domain}][%d] PostNetworksNetworkIDDNSRecordsDomain default  %+v", o._statusCode, o.Payload)
}
func (o *PostNetworksNetworkIDDNSRecordsDomainDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNetworksNetworkIDDNSRecordsDomainDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
