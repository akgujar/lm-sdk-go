// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/logicmonitor/lm-sdk-go/v3/models"
)

// GetDeviceConfigSourceConfigListReader is a Reader for the GetDeviceConfigSourceConfigList structure.
type GetDeviceConfigSourceConfigListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceConfigSourceConfigListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceConfigSourceConfigListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetDeviceConfigSourceConfigListTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetDeviceConfigSourceConfigListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceConfigSourceConfigListOK creates a GetDeviceConfigSourceConfigListOK with default headers values
func NewGetDeviceConfigSourceConfigListOK() *GetDeviceConfigSourceConfigListOK {
	return &GetDeviceConfigSourceConfigListOK{}
}

/*
GetDeviceConfigSourceConfigListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDeviceConfigSourceConfigListOK struct {
	Payload *models.DeviceDatasourceInstanceConfigPaginationResponse
}

// IsSuccess returns true when this get device config source config list o k response has a 2xx status code
func (o *GetDeviceConfigSourceConfigListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get device config source config list o k response has a 3xx status code
func (o *GetDeviceConfigSourceConfigListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device config source config list o k response has a 4xx status code
func (o *GetDeviceConfigSourceConfigListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device config source config list o k response has a 5xx status code
func (o *GetDeviceConfigSourceConfigListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get device config source config list o k response a status code equal to that given
func (o *GetDeviceConfigSourceConfigListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get device config source config list o k response
func (o *GetDeviceConfigSourceConfigListOK) Code() int {
	return 200
}

func (o *GetDeviceConfigSourceConfigListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/config][%d] getDeviceConfigSourceConfigListOK %s", 200, payload)
}

func (o *GetDeviceConfigSourceConfigListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/config][%d] getDeviceConfigSourceConfigListOK %s", 200, payload)
}

func (o *GetDeviceConfigSourceConfigListOK) GetPayload() *models.DeviceDatasourceInstanceConfigPaginationResponse {
	return o.Payload
}

func (o *GetDeviceConfigSourceConfigListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceDatasourceInstanceConfigPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceConfigSourceConfigListTooManyRequests creates a GetDeviceConfigSourceConfigListTooManyRequests with default headers values
func NewGetDeviceConfigSourceConfigListTooManyRequests() *GetDeviceConfigSourceConfigListTooManyRequests {
	return &GetDeviceConfigSourceConfigListTooManyRequests{}
}

/*
GetDeviceConfigSourceConfigListTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetDeviceConfigSourceConfigListTooManyRequests struct {

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

// IsSuccess returns true when this get device config source config list too many requests response has a 2xx status code
func (o *GetDeviceConfigSourceConfigListTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get device config source config list too many requests response has a 3xx status code
func (o *GetDeviceConfigSourceConfigListTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device config source config list too many requests response has a 4xx status code
func (o *GetDeviceConfigSourceConfigListTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get device config source config list too many requests response has a 5xx status code
func (o *GetDeviceConfigSourceConfigListTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get device config source config list too many requests response a status code equal to that given
func (o *GetDeviceConfigSourceConfigListTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get device config source config list too many requests response
func (o *GetDeviceConfigSourceConfigListTooManyRequests) Code() int {
	return 429
}

func (o *GetDeviceConfigSourceConfigListTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/config][%d] getDeviceConfigSourceConfigListTooManyRequests", 429)
}

func (o *GetDeviceConfigSourceConfigListTooManyRequests) String() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/config][%d] getDeviceConfigSourceConfigListTooManyRequests", 429)
}

func (o *GetDeviceConfigSourceConfigListTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDeviceConfigSourceConfigListDefault creates a GetDeviceConfigSourceConfigListDefault with default headers values
func NewGetDeviceConfigSourceConfigListDefault(code int) *GetDeviceConfigSourceConfigListDefault {
	return &GetDeviceConfigSourceConfigListDefault{
		_statusCode: code,
	}
}

/*
GetDeviceConfigSourceConfigListDefault describes a response with status code -1, with default header values.

Error
*/
type GetDeviceConfigSourceConfigListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get device config source config list default response has a 2xx status code
func (o *GetDeviceConfigSourceConfigListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get device config source config list default response has a 3xx status code
func (o *GetDeviceConfigSourceConfigListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get device config source config list default response has a 4xx status code
func (o *GetDeviceConfigSourceConfigListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get device config source config list default response has a 5xx status code
func (o *GetDeviceConfigSourceConfigListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get device config source config list default response a status code equal to that given
func (o *GetDeviceConfigSourceConfigListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get device config source config list default response
func (o *GetDeviceConfigSourceConfigListDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceConfigSourceConfigListDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/config][%d] getDeviceConfigSourceConfigList default %s", o._statusCode, payload)
}

func (o *GetDeviceConfigSourceConfigListDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/config][%d] getDeviceConfigSourceConfigList default %s", o._statusCode, payload)
}

func (o *GetDeviceConfigSourceConfigListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceConfigSourceConfigListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
