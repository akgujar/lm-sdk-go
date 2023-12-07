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

// GetDeviceDatasourceInstanceListReader is a Reader for the GetDeviceDatasourceInstanceList structure.
type GetDeviceDatasourceInstanceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceDatasourceInstanceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceDatasourceInstanceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetDeviceDatasourceInstanceListTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetDeviceDatasourceInstanceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceDatasourceInstanceListOK creates a GetDeviceDatasourceInstanceListOK with default headers values
func NewGetDeviceDatasourceInstanceListOK() *GetDeviceDatasourceInstanceListOK {
	return &GetDeviceDatasourceInstanceListOK{}
}

/* GetDeviceDatasourceInstanceListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDeviceDatasourceInstanceListOK struct {
	Payload *models.DeviceDatasourceInstancePaginationResponse
}

func (o *GetDeviceDatasourceInstanceListOK) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances][%d] getDeviceDatasourceInstanceListOK  %+v", 200, o.Payload)
}
func (o *GetDeviceDatasourceInstanceListOK) GetPayload() *models.DeviceDatasourceInstancePaginationResponse {
	return o.Payload
}

func (o *GetDeviceDatasourceInstanceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceDatasourceInstancePaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceDatasourceInstanceListTooManyRequests creates a GetDeviceDatasourceInstanceListTooManyRequests with default headers values
func NewGetDeviceDatasourceInstanceListTooManyRequests() *GetDeviceDatasourceInstanceListTooManyRequests {
	return &GetDeviceDatasourceInstanceListTooManyRequests{}
}

/* GetDeviceDatasourceInstanceListTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetDeviceDatasourceInstanceListTooManyRequests struct {

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

func (o *GetDeviceDatasourceInstanceListTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances][%d] getDeviceDatasourceInstanceListTooManyRequests ", 429)
}

func (o *GetDeviceDatasourceInstanceListTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDeviceDatasourceInstanceListDefault creates a GetDeviceDatasourceInstanceListDefault with default headers values
func NewGetDeviceDatasourceInstanceListDefault(code int) *GetDeviceDatasourceInstanceListDefault {
	return &GetDeviceDatasourceInstanceListDefault{
		_statusCode: code,
	}
}

/* GetDeviceDatasourceInstanceListDefault describes a response with status code -1, with default header values.

Error
*/
type GetDeviceDatasourceInstanceListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device datasource instance list default response
func (o *GetDeviceDatasourceInstanceListDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceDatasourceInstanceListDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances][%d] getDeviceDatasourceInstanceList default  %+v", o._statusCode, o.Payload)
}
func (o *GetDeviceDatasourceInstanceListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceDatasourceInstanceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
