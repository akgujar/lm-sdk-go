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

// GetDeviceDatasourceInstanceAlertSettingListOfDSIReader is a Reader for the GetDeviceDatasourceInstanceAlertSettingListOfDSI structure.
type GetDeviceDatasourceInstanceAlertSettingListOfDSIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceDatasourceInstanceAlertSettingListOfDSIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceDatasourceInstanceAlertSettingListOfDSIOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDeviceDatasourceInstanceAlertSettingListOfDSIDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceDatasourceInstanceAlertSettingListOfDSIOK creates a GetDeviceDatasourceInstanceAlertSettingListOfDSIOK with default headers values
func NewGetDeviceDatasourceInstanceAlertSettingListOfDSIOK() *GetDeviceDatasourceInstanceAlertSettingListOfDSIOK {
	return &GetDeviceDatasourceInstanceAlertSettingListOfDSIOK{}
}

/* GetDeviceDatasourceInstanceAlertSettingListOfDSIOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDeviceDatasourceInstanceAlertSettingListOfDSIOK struct {
	Payload *models.DeviceDataSourceInstanceAlertSettingPaginationResponse
}

func (o *GetDeviceDatasourceInstanceAlertSettingListOfDSIOK) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/alertsettings][%d] getDeviceDatasourceInstanceAlertSettingListOfDSIOK  %+v", 200, o.Payload)
}
func (o *GetDeviceDatasourceInstanceAlertSettingListOfDSIOK) GetPayload() *models.DeviceDataSourceInstanceAlertSettingPaginationResponse {
	return o.Payload
}

func (o *GetDeviceDatasourceInstanceAlertSettingListOfDSIOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceDataSourceInstanceAlertSettingPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceDatasourceInstanceAlertSettingListOfDSIDefault creates a GetDeviceDatasourceInstanceAlertSettingListOfDSIDefault with default headers values
func NewGetDeviceDatasourceInstanceAlertSettingListOfDSIDefault(code int) *GetDeviceDatasourceInstanceAlertSettingListOfDSIDefault {
	return &GetDeviceDatasourceInstanceAlertSettingListOfDSIDefault{
		_statusCode: code,
	}
}

/* GetDeviceDatasourceInstanceAlertSettingListOfDSIDefault describes a response with status code -1, with default header values.

Error
*/
type GetDeviceDatasourceInstanceAlertSettingListOfDSIDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device datasource instance alert setting list of d s i default response
func (o *GetDeviceDatasourceInstanceAlertSettingListOfDSIDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceDatasourceInstanceAlertSettingListOfDSIDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/alertsettings][%d] getDeviceDatasourceInstanceAlertSettingListOfDSI default  %+v", o._statusCode, o.Payload)
}
func (o *GetDeviceDatasourceInstanceAlertSettingListOfDSIDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceDatasourceInstanceAlertSettingListOfDSIDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}