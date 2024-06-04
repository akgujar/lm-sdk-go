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

// AddRoleReader is a Reader for the AddRole structure.
type AddRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewAddRoleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddRoleOK creates a AddRoleOK with default headers values
func NewAddRoleOK() *AddRoleOK {
	return &AddRoleOK{}
}

/*
AddRoleOK describes a response with status code 200, with default header values.

successful operation
*/
type AddRoleOK struct {
	Payload *models.Role
}

// IsSuccess returns true when this add role o k response has a 2xx status code
func (o *AddRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add role o k response has a 3xx status code
func (o *AddRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add role o k response has a 4xx status code
func (o *AddRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add role o k response has a 5xx status code
func (o *AddRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add role o k response a status code equal to that given
func (o *AddRoleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add role o k response
func (o *AddRoleOK) Code() int {
	return 200
}

func (o *AddRoleOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /setting/roles][%d] addRoleOK %s", 200, payload)
}

func (o *AddRoleOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /setting/roles][%d] addRoleOK %s", 200, payload)
}

func (o *AddRoleOK) GetPayload() *models.Role {
	return o.Payload
}

func (o *AddRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddRoleTooManyRequests creates a AddRoleTooManyRequests with default headers values
func NewAddRoleTooManyRequests() *AddRoleTooManyRequests {
	return &AddRoleTooManyRequests{}
}

/*
AddRoleTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AddRoleTooManyRequests struct {

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

// IsSuccess returns true when this add role too many requests response has a 2xx status code
func (o *AddRoleTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add role too many requests response has a 3xx status code
func (o *AddRoleTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add role too many requests response has a 4xx status code
func (o *AddRoleTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this add role too many requests response has a 5xx status code
func (o *AddRoleTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this add role too many requests response a status code equal to that given
func (o *AddRoleTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the add role too many requests response
func (o *AddRoleTooManyRequests) Code() int {
	return 429
}

func (o *AddRoleTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /setting/roles][%d] addRoleTooManyRequests", 429)
}

func (o *AddRoleTooManyRequests) String() string {
	return fmt.Sprintf("[POST /setting/roles][%d] addRoleTooManyRequests", 429)
}

func (o *AddRoleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddRoleDefault creates a AddRoleDefault with default headers values
func NewAddRoleDefault(code int) *AddRoleDefault {
	return &AddRoleDefault{
		_statusCode: code,
	}
}

/*
AddRoleDefault describes a response with status code -1, with default header values.

Error
*/
type AddRoleDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this add role default response has a 2xx status code
func (o *AddRoleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this add role default response has a 3xx status code
func (o *AddRoleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this add role default response has a 4xx status code
func (o *AddRoleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this add role default response has a 5xx status code
func (o *AddRoleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this add role default response a status code equal to that given
func (o *AddRoleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the add role default response
func (o *AddRoleDefault) Code() int {
	return o._statusCode
}

func (o *AddRoleDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /setting/roles][%d] addRole default %s", o._statusCode, payload)
}

func (o *AddRoleDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /setting/roles][%d] addRole default %s", o._statusCode, payload)
}

func (o *AddRoleDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
