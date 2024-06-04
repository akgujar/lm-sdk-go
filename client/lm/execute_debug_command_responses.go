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

// ExecuteDebugCommandReader is a Reader for the ExecuteDebugCommand structure.
type ExecuteDebugCommandReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecuteDebugCommandReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExecuteDebugCommandOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewExecuteDebugCommandTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewExecuteDebugCommandDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExecuteDebugCommandOK creates a ExecuteDebugCommandOK with default headers values
func NewExecuteDebugCommandOK() *ExecuteDebugCommandOK {
	return &ExecuteDebugCommandOK{}
}

/*
ExecuteDebugCommandOK describes a response with status code 200, with default header values.

successful operation
*/
type ExecuteDebugCommandOK struct {
	Payload *models.Debug
}

// IsSuccess returns true when this execute debug command o k response has a 2xx status code
func (o *ExecuteDebugCommandOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this execute debug command o k response has a 3xx status code
func (o *ExecuteDebugCommandOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this execute debug command o k response has a 4xx status code
func (o *ExecuteDebugCommandOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this execute debug command o k response has a 5xx status code
func (o *ExecuteDebugCommandOK) IsServerError() bool {
	return false
}

// IsCode returns true when this execute debug command o k response a status code equal to that given
func (o *ExecuteDebugCommandOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the execute debug command o k response
func (o *ExecuteDebugCommandOK) Code() int {
	return 200
}

func (o *ExecuteDebugCommandOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /debug][%d] executeDebugCommandOK %s", 200, payload)
}

func (o *ExecuteDebugCommandOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /debug][%d] executeDebugCommandOK %s", 200, payload)
}

func (o *ExecuteDebugCommandOK) GetPayload() *models.Debug {
	return o.Payload
}

func (o *ExecuteDebugCommandOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Debug)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecuteDebugCommandTooManyRequests creates a ExecuteDebugCommandTooManyRequests with default headers values
func NewExecuteDebugCommandTooManyRequests() *ExecuteDebugCommandTooManyRequests {
	return &ExecuteDebugCommandTooManyRequests{}
}

/*
ExecuteDebugCommandTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ExecuteDebugCommandTooManyRequests struct {

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

// IsSuccess returns true when this execute debug command too many requests response has a 2xx status code
func (o *ExecuteDebugCommandTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this execute debug command too many requests response has a 3xx status code
func (o *ExecuteDebugCommandTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this execute debug command too many requests response has a 4xx status code
func (o *ExecuteDebugCommandTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this execute debug command too many requests response has a 5xx status code
func (o *ExecuteDebugCommandTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this execute debug command too many requests response a status code equal to that given
func (o *ExecuteDebugCommandTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the execute debug command too many requests response
func (o *ExecuteDebugCommandTooManyRequests) Code() int {
	return 429
}

func (o *ExecuteDebugCommandTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /debug][%d] executeDebugCommandTooManyRequests", 429)
}

func (o *ExecuteDebugCommandTooManyRequests) String() string {
	return fmt.Sprintf("[POST /debug][%d] executeDebugCommandTooManyRequests", 429)
}

func (o *ExecuteDebugCommandTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewExecuteDebugCommandDefault creates a ExecuteDebugCommandDefault with default headers values
func NewExecuteDebugCommandDefault(code int) *ExecuteDebugCommandDefault {
	return &ExecuteDebugCommandDefault{
		_statusCode: code,
	}
}

/*
ExecuteDebugCommandDefault describes a response with status code -1, with default header values.

Error
*/
type ExecuteDebugCommandDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this execute debug command default response has a 2xx status code
func (o *ExecuteDebugCommandDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this execute debug command default response has a 3xx status code
func (o *ExecuteDebugCommandDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this execute debug command default response has a 4xx status code
func (o *ExecuteDebugCommandDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this execute debug command default response has a 5xx status code
func (o *ExecuteDebugCommandDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this execute debug command default response a status code equal to that given
func (o *ExecuteDebugCommandDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the execute debug command default response
func (o *ExecuteDebugCommandDefault) Code() int {
	return o._statusCode
}

func (o *ExecuteDebugCommandDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /debug][%d] executeDebugCommand default %s", o._statusCode, payload)
}

func (o *ExecuteDebugCommandDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /debug][%d] executeDebugCommand default %s", o._statusCode, payload)
}

func (o *ExecuteDebugCommandDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExecuteDebugCommandDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
