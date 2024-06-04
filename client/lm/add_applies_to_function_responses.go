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

// AddAppliesToFunctionReader is a Reader for the AddAppliesToFunction structure.
type AddAppliesToFunctionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddAppliesToFunctionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddAppliesToFunctionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewAddAppliesToFunctionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddAppliesToFunctionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddAppliesToFunctionOK creates a AddAppliesToFunctionOK with default headers values
func NewAddAppliesToFunctionOK() *AddAppliesToFunctionOK {
	return &AddAppliesToFunctionOK{}
}

/*
AddAppliesToFunctionOK describes a response with status code 200, with default header values.

successful operation
*/
type AddAppliesToFunctionOK struct {
	Payload *models.AppliesToFunction
}

// IsSuccess returns true when this add applies to function o k response has a 2xx status code
func (o *AddAppliesToFunctionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add applies to function o k response has a 3xx status code
func (o *AddAppliesToFunctionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add applies to function o k response has a 4xx status code
func (o *AddAppliesToFunctionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add applies to function o k response has a 5xx status code
func (o *AddAppliesToFunctionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add applies to function o k response a status code equal to that given
func (o *AddAppliesToFunctionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add applies to function o k response
func (o *AddAppliesToFunctionOK) Code() int {
	return 200
}

func (o *AddAppliesToFunctionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /setting/functions][%d] addAppliesToFunctionOK %s", 200, payload)
}

func (o *AddAppliesToFunctionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /setting/functions][%d] addAppliesToFunctionOK %s", 200, payload)
}

func (o *AddAppliesToFunctionOK) GetPayload() *models.AppliesToFunction {
	return o.Payload
}

func (o *AddAppliesToFunctionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppliesToFunction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddAppliesToFunctionTooManyRequests creates a AddAppliesToFunctionTooManyRequests with default headers values
func NewAddAppliesToFunctionTooManyRequests() *AddAppliesToFunctionTooManyRequests {
	return &AddAppliesToFunctionTooManyRequests{}
}

/*
AddAppliesToFunctionTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AddAppliesToFunctionTooManyRequests struct {

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

// IsSuccess returns true when this add applies to function too many requests response has a 2xx status code
func (o *AddAppliesToFunctionTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add applies to function too many requests response has a 3xx status code
func (o *AddAppliesToFunctionTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add applies to function too many requests response has a 4xx status code
func (o *AddAppliesToFunctionTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this add applies to function too many requests response has a 5xx status code
func (o *AddAppliesToFunctionTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this add applies to function too many requests response a status code equal to that given
func (o *AddAppliesToFunctionTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the add applies to function too many requests response
func (o *AddAppliesToFunctionTooManyRequests) Code() int {
	return 429
}

func (o *AddAppliesToFunctionTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /setting/functions][%d] addAppliesToFunctionTooManyRequests", 429)
}

func (o *AddAppliesToFunctionTooManyRequests) String() string {
	return fmt.Sprintf("[POST /setting/functions][%d] addAppliesToFunctionTooManyRequests", 429)
}

func (o *AddAppliesToFunctionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddAppliesToFunctionDefault creates a AddAppliesToFunctionDefault with default headers values
func NewAddAppliesToFunctionDefault(code int) *AddAppliesToFunctionDefault {
	return &AddAppliesToFunctionDefault{
		_statusCode: code,
	}
}

/*
AddAppliesToFunctionDefault describes a response with status code -1, with default header values.

Error
*/
type AddAppliesToFunctionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this add applies to function default response has a 2xx status code
func (o *AddAppliesToFunctionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this add applies to function default response has a 3xx status code
func (o *AddAppliesToFunctionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this add applies to function default response has a 4xx status code
func (o *AddAppliesToFunctionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this add applies to function default response has a 5xx status code
func (o *AddAppliesToFunctionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this add applies to function default response a status code equal to that given
func (o *AddAppliesToFunctionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the add applies to function default response
func (o *AddAppliesToFunctionDefault) Code() int {
	return o._statusCode
}

func (o *AddAppliesToFunctionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /setting/functions][%d] addAppliesToFunction default %s", o._statusCode, payload)
}

func (o *AddAppliesToFunctionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /setting/functions][%d] addAppliesToFunction default %s", o._statusCode, payload)
}

func (o *AddAppliesToFunctionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddAppliesToFunctionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
