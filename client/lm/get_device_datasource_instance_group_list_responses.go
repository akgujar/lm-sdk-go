// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/logicmonitor/lm-sdk-go/models"
)

// GetDeviceDatasourceInstanceGroupListReader is a Reader for the GetDeviceDatasourceInstanceGroupList structure.
type GetDeviceDatasourceInstanceGroupListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceDatasourceInstanceGroupListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceDatasourceInstanceGroupListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetDeviceDatasourceInstanceGroupListTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetDeviceDatasourceInstanceGroupListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceDatasourceInstanceGroupListOK creates a GetDeviceDatasourceInstanceGroupListOK with default headers values
func NewGetDeviceDatasourceInstanceGroupListOK() *GetDeviceDatasourceInstanceGroupListOK {
	return &GetDeviceDatasourceInstanceGroupListOK{}
}

/* GetDeviceDatasourceInstanceGroupListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDeviceDatasourceInstanceGroupListOK struct {
	Payload *models.DeviceDatasourceInstanceGroupPaginationResponse
}

func (o *GetDeviceDatasourceInstanceGroupListOK) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{deviceDsId}/groups][%d] getDeviceDatasourceInstanceGroupListOK  %+v", 200, o.Payload)
}
func (o *GetDeviceDatasourceInstanceGroupListOK) GetPayload() *models.DeviceDatasourceInstanceGroupPaginationResponse {
	return o.Payload
}

func (o *GetDeviceDatasourceInstanceGroupListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceDatasourceInstanceGroupPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceDatasourceInstanceGroupListTooManyRequests creates a GetDeviceDatasourceInstanceGroupListTooManyRequests with default headers values
func NewGetDeviceDatasourceInstanceGroupListTooManyRequests() *GetDeviceDatasourceInstanceGroupListTooManyRequests {
	return &GetDeviceDatasourceInstanceGroupListTooManyRequests{}
}

/* GetDeviceDatasourceInstanceGroupListTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetDeviceDatasourceInstanceGroupListTooManyRequests struct {

	/* Request limit per X-Rate-Limit-Window
	 */
	XRateLimitLimit int64

	/* The number of requests left for the time window
	 */
	XRateLimitRemaining int64

	/* The rolling time window length with the unit of second
	 */
	XRateLimitWindow int64
}

func (o *GetDeviceDatasourceInstanceGroupListTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{deviceDsId}/groups][%d] getDeviceDatasourceInstanceGroupListTooManyRequests ", 429)
}

func (o *GetDeviceDatasourceInstanceGroupListTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-rate-limit-limit
	hdrXRateLimitLimit := response.GetHeader("x-rate-limit-limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("x-rate-limit-limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header x-rate-limit-remaining
	hdrXRateLimitRemaining := response.GetHeader("x-rate-limit-remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("x-rate-limit-remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header x-rate-limit-window
	hdrXRateLimitWindow := response.GetHeader("x-rate-limit-window")

	if hdrXRateLimitWindow != "" {
		valxRateLimitWindow, err := swag.ConvertInt64(hdrXRateLimitWindow)
		if err != nil {
			return errors.InvalidType("x-rate-limit-window", "header", "int64", hdrXRateLimitWindow)
		}
		o.XRateLimitWindow = valxRateLimitWindow
	}

	return nil
}

// NewGetDeviceDatasourceInstanceGroupListDefault creates a GetDeviceDatasourceInstanceGroupListDefault with default headers values
func NewGetDeviceDatasourceInstanceGroupListDefault(code int) *GetDeviceDatasourceInstanceGroupListDefault {
	return &GetDeviceDatasourceInstanceGroupListDefault{
		_statusCode: code,
	}
}

/* GetDeviceDatasourceInstanceGroupListDefault describes a response with status code -1, with default header values.

Error
*/
type GetDeviceDatasourceInstanceGroupListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device datasource instance group list default response
func (o *GetDeviceDatasourceInstanceGroupListDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceDatasourceInstanceGroupListDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{deviceDsId}/groups][%d] getDeviceDatasourceInstanceGroupList default  %+v", o._statusCode, o.Payload)
}
func (o *GetDeviceDatasourceInstanceGroupListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceDatasourceInstanceGroupListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
