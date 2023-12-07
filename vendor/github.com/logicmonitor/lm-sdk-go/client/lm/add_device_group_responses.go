// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/logicmonitor/lm-sdk-go/models"
)

// AddDeviceGroupReader is a Reader for the AddDeviceGroup structure.
type AddDeviceGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddDeviceGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddDeviceGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddDeviceGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddDeviceGroupOK creates a AddDeviceGroupOK with default headers values
func NewAddDeviceGroupOK() *AddDeviceGroupOK {
	return &AddDeviceGroupOK{}
}

/* AddDeviceGroupOK describes a response with status code 200, with default header values.

successful operation
*/
type AddDeviceGroupOK struct {
	Payload *models.DeviceGroup
}

func (o *AddDeviceGroupOK) Error() string {
	return fmt.Sprintf("[POST /device/groups][%d] addDeviceGroupOK  %+v", 200, o.Payload)
}
func (o *AddDeviceGroupOK) GetPayload() *models.DeviceGroup {
	return o.Payload
}

func (o *AddDeviceGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDeviceGroupDefault creates a AddDeviceGroupDefault with default headers values
func NewAddDeviceGroupDefault(code int) *AddDeviceGroupDefault {
	return &AddDeviceGroupDefault{
		_statusCode: code,
	}
}

/* AddDeviceGroupDefault describes a response with status code -1, with default header values.

Error
*/
type AddDeviceGroupDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add device group default response
func (o *AddDeviceGroupDefault) Code() int {
	return o._statusCode
}

func (o *AddDeviceGroupDefault) Error() string {
	return fmt.Sprintf("[POST /device/groups][%d] addDeviceGroup default  %+v", o._statusCode, o.Payload)
}
func (o *AddDeviceGroupDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddDeviceGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
