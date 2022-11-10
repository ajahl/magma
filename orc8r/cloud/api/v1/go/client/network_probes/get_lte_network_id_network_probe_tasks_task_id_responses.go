// Code generated by go-swagger; DO NOT EDIT.

package network_probes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// GetLTENetworkIDNetworkProbeTasksTaskIDReader is a Reader for the GetLTENetworkIDNetworkProbeTasksTaskID structure.
type GetLTENetworkIDNetworkProbeTasksTaskIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDNetworkProbeTasksTaskIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDNetworkProbeTasksTaskIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDNetworkProbeTasksTaskIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDNetworkProbeTasksTaskIDOK creates a GetLTENetworkIDNetworkProbeTasksTaskIDOK with default headers values
func NewGetLTENetworkIDNetworkProbeTasksTaskIDOK() *GetLTENetworkIDNetworkProbeTasksTaskIDOK {
	return &GetLTENetworkIDNetworkProbeTasksTaskIDOK{}
}

/*
GetLTENetworkIDNetworkProbeTasksTaskIDOK describes a response with status code 200, with default header values.

NetworkProbeTask Info
*/
type GetLTENetworkIDNetworkProbeTasksTaskIDOK struct {
	Payload *models.NetworkProbeTask
}

func (o *GetLTENetworkIDNetworkProbeTasksTaskIDOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/network_probe/tasks/{task_id}][%d] getLteNetworkIdNetworkProbeTasksTaskIdOK  %+v", 200, o.Payload)
}
func (o *GetLTENetworkIDNetworkProbeTasksTaskIDOK) GetPayload() *models.NetworkProbeTask {
	return o.Payload
}

func (o *GetLTENetworkIDNetworkProbeTasksTaskIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkProbeTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDNetworkProbeTasksTaskIDDefault creates a GetLTENetworkIDNetworkProbeTasksTaskIDDefault with default headers values
func NewGetLTENetworkIDNetworkProbeTasksTaskIDDefault(code int) *GetLTENetworkIDNetworkProbeTasksTaskIDDefault {
	return &GetLTENetworkIDNetworkProbeTasksTaskIDDefault{
		_statusCode: code,
	}
}

/*
GetLTENetworkIDNetworkProbeTasksTaskIDDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetLTENetworkIDNetworkProbeTasksTaskIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID network probe tasks task ID default response
func (o *GetLTENetworkIDNetworkProbeTasksTaskIDDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDNetworkProbeTasksTaskIDDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/network_probe/tasks/{task_id}][%d] GetLTENetworkIDNetworkProbeTasksTaskID default  %+v", o._statusCode, o.Payload)
}
func (o *GetLTENetworkIDNetworkProbeTasksTaskIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDNetworkProbeTasksTaskIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
