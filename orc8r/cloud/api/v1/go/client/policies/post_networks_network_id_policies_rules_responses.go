// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// PostNetworksNetworkIDPoliciesRulesReader is a Reader for the PostNetworksNetworkIDPoliciesRules structure.
type PostNetworksNetworkIDPoliciesRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNetworksNetworkIDPoliciesRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostNetworksNetworkIDPoliciesRulesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostNetworksNetworkIDPoliciesRulesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostNetworksNetworkIDPoliciesRulesCreated creates a PostNetworksNetworkIDPoliciesRulesCreated with default headers values
func NewPostNetworksNetworkIDPoliciesRulesCreated() *PostNetworksNetworkIDPoliciesRulesCreated {
	return &PostNetworksNetworkIDPoliciesRulesCreated{}
}

/*
PostNetworksNetworkIDPoliciesRulesCreated describes a response with status code 201, with default header values.

Rule id
*/
type PostNetworksNetworkIDPoliciesRulesCreated struct {
	Payload models.RuleID
}

func (o *PostNetworksNetworkIDPoliciesRulesCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{network_id}/policies/rules][%d] postNetworksNetworkIdPoliciesRulesCreated  %+v", 201, o.Payload)
}
func (o *PostNetworksNetworkIDPoliciesRulesCreated) GetPayload() models.RuleID {
	return o.Payload
}

func (o *PostNetworksNetworkIDPoliciesRulesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNetworksNetworkIDPoliciesRulesDefault creates a PostNetworksNetworkIDPoliciesRulesDefault with default headers values
func NewPostNetworksNetworkIDPoliciesRulesDefault(code int) *PostNetworksNetworkIDPoliciesRulesDefault {
	return &PostNetworksNetworkIDPoliciesRulesDefault{
		_statusCode: code,
	}
}

/*
PostNetworksNetworkIDPoliciesRulesDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PostNetworksNetworkIDPoliciesRulesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post networks network ID policies rules default response
func (o *PostNetworksNetworkIDPoliciesRulesDefault) Code() int {
	return o._statusCode
}

func (o *PostNetworksNetworkIDPoliciesRulesDefault) Error() string {
	return fmt.Sprintf("[POST /networks/{network_id}/policies/rules][%d] PostNetworksNetworkIDPoliciesRules default  %+v", o._statusCode, o.Payload)
}
func (o *PostNetworksNetworkIDPoliciesRulesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNetworksNetworkIDPoliciesRulesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
