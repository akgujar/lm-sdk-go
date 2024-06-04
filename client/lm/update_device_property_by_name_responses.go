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

// UpdateDevicePropertyByNameReader is a Reader for the UpdateDevicePropertyByName structure.
type UpdateDevicePropertyByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDevicePropertyByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDevicePropertyByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewUpdateDevicePropertyByNameTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateDevicePropertyByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDevicePropertyByNameOK creates a UpdateDevicePropertyByNameOK with default headers values
func NewUpdateDevicePropertyByNameOK() *UpdateDevicePropertyByNameOK {
	return &UpdateDevicePropertyByNameOK{}
}

/*
UpdateDevicePropertyByNameOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateDevicePropertyByNameOK struct {
	Payload *models.EntityProperty
}

// IsSuccess returns true when this update device property by name o k response has a 2xx status code
func (o *UpdateDevicePropertyByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update device property by name o k response has a 3xx status code
func (o *UpdateDevicePropertyByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update device property by name o k response has a 4xx status code
func (o *UpdateDevicePropertyByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update device property by name o k response has a 5xx status code
func (o *UpdateDevicePropertyByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update device property by name o k response a status code equal to that given
func (o *UpdateDevicePropertyByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update device property by name o k response
func (o *UpdateDevicePropertyByNameOK) Code() int {
	return 200
}

func (o *UpdateDevicePropertyByNameOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /device/devices/{deviceId}/properties/{name}][%d] updateDevicePropertyByNameOK %s", 200, payload)
}

func (o *UpdateDevicePropertyByNameOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /device/devices/{deviceId}/properties/{name}][%d] updateDevicePropertyByNameOK %s", 200, payload)
}

func (o *UpdateDevicePropertyByNameOK) GetPayload() *models.EntityProperty {
	return o.Payload
}

func (o *UpdateDevicePropertyByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EntityProperty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDevicePropertyByNameTooManyRequests creates a UpdateDevicePropertyByNameTooManyRequests with default headers values
func NewUpdateDevicePropertyByNameTooManyRequests() *UpdateDevicePropertyByNameTooManyRequests {
	return &UpdateDevicePropertyByNameTooManyRequests{}
}

/*
UpdateDevicePropertyByNameTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateDevicePropertyByNameTooManyRequests struct {

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

// IsSuccess returns true when this update device property by name too many requests response has a 2xx status code
func (o *UpdateDevicePropertyByNameTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update device property by name too many requests response has a 3xx status code
func (o *UpdateDevicePropertyByNameTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update device property by name too many requests response has a 4xx status code
func (o *UpdateDevicePropertyByNameTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update device property by name too many requests response has a 5xx status code
func (o *UpdateDevicePropertyByNameTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update device property by name too many requests response a status code equal to that given
func (o *UpdateDevicePropertyByNameTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the update device property by name too many requests response
func (o *UpdateDevicePropertyByNameTooManyRequests) Code() int {
	return 429
}

func (o *UpdateDevicePropertyByNameTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /device/devices/{deviceId}/properties/{name}][%d] updateDevicePropertyByNameTooManyRequests", 429)
}

func (o *UpdateDevicePropertyByNameTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /device/devices/{deviceId}/properties/{name}][%d] updateDevicePropertyByNameTooManyRequests", 429)
}

func (o *UpdateDevicePropertyByNameTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateDevicePropertyByNameDefault creates a UpdateDevicePropertyByNameDefault with default headers values
func NewUpdateDevicePropertyByNameDefault(code int) *UpdateDevicePropertyByNameDefault {
	return &UpdateDevicePropertyByNameDefault{
		_statusCode: code,
	}
}

/*
UpdateDevicePropertyByNameDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateDevicePropertyByNameDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update device property by name default response has a 2xx status code
func (o *UpdateDevicePropertyByNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update device property by name default response has a 3xx status code
func (o *UpdateDevicePropertyByNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update device property by name default response has a 4xx status code
func (o *UpdateDevicePropertyByNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update device property by name default response has a 5xx status code
func (o *UpdateDevicePropertyByNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update device property by name default response a status code equal to that given
func (o *UpdateDevicePropertyByNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update device property by name default response
func (o *UpdateDevicePropertyByNameDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDevicePropertyByNameDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /device/devices/{deviceId}/properties/{name}][%d] updateDevicePropertyByName default %s", o._statusCode, payload)
}

func (o *UpdateDevicePropertyByNameDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /device/devices/{deviceId}/properties/{name}][%d] updateDevicePropertyByName default %s", o._statusCode, payload)
}

func (o *UpdateDevicePropertyByNameDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateDevicePropertyByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
