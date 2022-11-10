// Code generated by go-swagger; DO NOT EDIT.

package upgrades

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// DeleteNetworksNetworkIDTiersTierIDImagesImageNameReader is a Reader for the DeleteNetworksNetworkIDTiersTierIDImagesImageName structure.
type DeleteNetworksNetworkIDTiersTierIDImagesImageNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworksNetworkIDTiersTierIDImagesImageNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworksNetworkIDTiersTierIDImagesImageNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteNetworksNetworkIDTiersTierIDImagesImageNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNetworksNetworkIDTiersTierIDImagesImageNameNoContent creates a DeleteNetworksNetworkIDTiersTierIDImagesImageNameNoContent with default headers values
func NewDeleteNetworksNetworkIDTiersTierIDImagesImageNameNoContent() *DeleteNetworksNetworkIDTiersTierIDImagesImageNameNoContent {
	return &DeleteNetworksNetworkIDTiersTierIDImagesImageNameNoContent{}
}

/*
DeleteNetworksNetworkIDTiersTierIDImagesImageNameNoContent describes a response with status code 204, with default header values.

Success
*/
type DeleteNetworksNetworkIDTiersTierIDImagesImageNameNoContent struct {
}

func (o *DeleteNetworksNetworkIDTiersTierIDImagesImageNameNoContent) Error() string {
	return fmt.Sprintf("[DELETE /networks/{network_id}/tiers/{tier_id}/images/{image_name}][%d] deleteNetworksNetworkIdTiersTierIdImagesImageNameNoContent ", 204)
}

func (o *DeleteNetworksNetworkIDTiersTierIDImagesImageNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNetworksNetworkIDTiersTierIDImagesImageNameDefault creates a DeleteNetworksNetworkIDTiersTierIDImagesImageNameDefault with default headers values
func NewDeleteNetworksNetworkIDTiersTierIDImagesImageNameDefault(code int) *DeleteNetworksNetworkIDTiersTierIDImagesImageNameDefault {
	return &DeleteNetworksNetworkIDTiersTierIDImagesImageNameDefault{
		_statusCode: code,
	}
}

/*
DeleteNetworksNetworkIDTiersTierIDImagesImageNameDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type DeleteNetworksNetworkIDTiersTierIDImagesImageNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete networks network ID tiers tier ID images image name default response
func (o *DeleteNetworksNetworkIDTiersTierIDImagesImageNameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNetworksNetworkIDTiersTierIDImagesImageNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /networks/{network_id}/tiers/{tier_id}/images/{image_name}][%d] DeleteNetworksNetworkIDTiersTierIDImagesImageName default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteNetworksNetworkIDTiersTierIDImagesImageNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteNetworksNetworkIDTiersTierIDImagesImageNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
