// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// GetTenantsReader is a Reader for the GetTenants structure.
type GetTenantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTenantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTenantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetTenantsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTenantsOK creates a GetTenantsOK with default headers values
func NewGetTenantsOK() *GetTenantsOK {
	return &GetTenantsOK{}
}

/*
GetTenantsOK describes a response with status code 200, with default header values.

List of tenants
*/
type GetTenantsOK struct {
	Payload []*models.Tenant
}

func (o *GetTenantsOK) Error() string {
	return fmt.Sprintf("[GET /tenants][%d] getTenantsOK  %+v", 200, o.Payload)
}
func (o *GetTenantsOK) GetPayload() []*models.Tenant {
	return o.Payload
}

func (o *GetTenantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTenantsDefault creates a GetTenantsDefault with default headers values
func NewGetTenantsDefault(code int) *GetTenantsDefault {
	return &GetTenantsDefault{
		_statusCode: code,
	}
}

/*
GetTenantsDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetTenantsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get tenants default response
func (o *GetTenantsDefault) Code() int {
	return o._statusCode
}

func (o *GetTenantsDefault) Error() string {
	return fmt.Sprintf("[GET /tenants][%d] GetTenants default  %+v", o._statusCode, o.Payload)
}
func (o *GetTenantsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTenantsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
