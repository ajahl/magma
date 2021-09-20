// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetNetworksNetworkIDPoliciesBaseNamesViewFullReader is a Reader for the GetNetworksNetworkIDPoliciesBaseNamesViewFull structure.
type GetNetworksNetworkIDPoliciesBaseNamesViewFullReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDPoliciesBaseNamesViewFullReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDPoliciesBaseNamesViewFullOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDPoliciesBaseNamesViewFullDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDPoliciesBaseNamesViewFullOK creates a GetNetworksNetworkIDPoliciesBaseNamesViewFullOK with default headers values
func NewGetNetworksNetworkIDPoliciesBaseNamesViewFullOK() *GetNetworksNetworkIDPoliciesBaseNamesViewFullOK {
	return &GetNetworksNetworkIDPoliciesBaseNamesViewFullOK{}
}

/*GetNetworksNetworkIDPoliciesBaseNamesViewFullOK handles this case with default header values.

Map of all base names in the network by name
*/
type GetNetworksNetworkIDPoliciesBaseNamesViewFullOK struct {
	Payload map[string]models.BaseNameRecord
}

func (o *GetNetworksNetworkIDPoliciesBaseNamesViewFullOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/policies/base_names?view=full][%d] getNetworksNetworkIdPoliciesBaseNamesViewFullOK  %+v", 200, o.Payload)
}

func (o *GetNetworksNetworkIDPoliciesBaseNamesViewFullOK) GetPayload() map[string]models.BaseNameRecord {
	return o.Payload
}

func (o *GetNetworksNetworkIDPoliciesBaseNamesViewFullOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDPoliciesBaseNamesViewFullDefault creates a GetNetworksNetworkIDPoliciesBaseNamesViewFullDefault with default headers values
func NewGetNetworksNetworkIDPoliciesBaseNamesViewFullDefault(code int) *GetNetworksNetworkIDPoliciesBaseNamesViewFullDefault {
	return &GetNetworksNetworkIDPoliciesBaseNamesViewFullDefault{
		_statusCode: code,
	}
}

/*GetNetworksNetworkIDPoliciesBaseNamesViewFullDefault handles this case with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDPoliciesBaseNamesViewFullDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID policies base names view full default response
func (o *GetNetworksNetworkIDPoliciesBaseNamesViewFullDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDPoliciesBaseNamesViewFullDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/policies/base_names?view=full][%d] GetNetworksNetworkIDPoliciesBaseNamesViewFull default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworksNetworkIDPoliciesBaseNamesViewFullDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDPoliciesBaseNamesViewFullDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}