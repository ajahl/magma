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

// GetNetworksNetworkIDPoliciesRulesRuleIDReader is a Reader for the GetNetworksNetworkIDPoliciesRulesRuleID structure.
type GetNetworksNetworkIDPoliciesRulesRuleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDPoliciesRulesRuleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDPoliciesRulesRuleIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDPoliciesRulesRuleIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDPoliciesRulesRuleIDOK creates a GetNetworksNetworkIDPoliciesRulesRuleIDOK with default headers values
func NewGetNetworksNetworkIDPoliciesRulesRuleIDOK() *GetNetworksNetworkIDPoliciesRulesRuleIDOK {
	return &GetNetworksNetworkIDPoliciesRulesRuleIDOK{}
}

/*
GetNetworksNetworkIDPoliciesRulesRuleIDOK describes a response with status code 200, with default header values.

Policy rule on success
*/
type GetNetworksNetworkIDPoliciesRulesRuleIDOK struct {
	Payload *models.PolicyRule
}

func (o *GetNetworksNetworkIDPoliciesRulesRuleIDOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/policies/rules/{rule_id}][%d] getNetworksNetworkIdPoliciesRulesRuleIdOK  %+v", 200, o.Payload)
}
func (o *GetNetworksNetworkIDPoliciesRulesRuleIDOK) GetPayload() *models.PolicyRule {
	return o.Payload
}

func (o *GetNetworksNetworkIDPoliciesRulesRuleIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDPoliciesRulesRuleIDDefault creates a GetNetworksNetworkIDPoliciesRulesRuleIDDefault with default headers values
func NewGetNetworksNetworkIDPoliciesRulesRuleIDDefault(code int) *GetNetworksNetworkIDPoliciesRulesRuleIDDefault {
	return &GetNetworksNetworkIDPoliciesRulesRuleIDDefault{
		_statusCode: code,
	}
}

/*
GetNetworksNetworkIDPoliciesRulesRuleIDDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDPoliciesRulesRuleIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID policies rules rule ID default response
func (o *GetNetworksNetworkIDPoliciesRulesRuleIDDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDPoliciesRulesRuleIDDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/policies/rules/{rule_id}][%d] GetNetworksNetworkIDPoliciesRulesRuleID default  %+v", o._statusCode, o.Payload)
}
func (o *GetNetworksNetworkIDPoliciesRulesRuleIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDPoliciesRulesRuleIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
