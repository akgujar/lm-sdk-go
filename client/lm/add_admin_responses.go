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

// AddAdminReader is a Reader for the AddAdmin structure.
type AddAdminReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddAdminReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddAdminOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewAddAdminTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddAdminDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddAdminOK creates a AddAdminOK with default headers values
func NewAddAdminOK() *AddAdminOK {
	return &AddAdminOK{}
}

/*
AddAdminOK describes a response with status code 200, with default header values.

successful operation
*/
type AddAdminOK struct {
	Payload *models.Admin
}

// IsSuccess returns true when this add admin o k response has a 2xx status code
func (o *AddAdminOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add admin o k response has a 3xx status code
func (o *AddAdminOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add admin o k response has a 4xx status code
func (o *AddAdminOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add admin o k response has a 5xx status code
func (o *AddAdminOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add admin o k response a status code equal to that given
func (o *AddAdminOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add admin o k response
func (o *AddAdminOK) Code() int {
	return 200
}

func (o *AddAdminOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /setting/admins][%d] addAdminOK %s", 200, payload)
}

func (o *AddAdminOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /setting/admins][%d] addAdminOK %s", 200, payload)
}

func (o *AddAdminOK) GetPayload() *models.Admin {
	return o.Payload
}

func (o *AddAdminOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Admin)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddAdminTooManyRequests creates a AddAdminTooManyRequests with default headers values
func NewAddAdminTooManyRequests() *AddAdminTooManyRequests {
	return &AddAdminTooManyRequests{}
}

/*
AddAdminTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AddAdminTooManyRequests struct {

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

// IsSuccess returns true when this add admin too many requests response has a 2xx status code
func (o *AddAdminTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add admin too many requests response has a 3xx status code
func (o *AddAdminTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add admin too many requests response has a 4xx status code
func (o *AddAdminTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this add admin too many requests response has a 5xx status code
func (o *AddAdminTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this add admin too many requests response a status code equal to that given
func (o *AddAdminTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the add admin too many requests response
func (o *AddAdminTooManyRequests) Code() int {
	return 429
}

func (o *AddAdminTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /setting/admins][%d] addAdminTooManyRequests", 429)
}

func (o *AddAdminTooManyRequests) String() string {
	return fmt.Sprintf("[POST /setting/admins][%d] addAdminTooManyRequests", 429)
}

func (o *AddAdminTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddAdminDefault creates a AddAdminDefault with default headers values
func NewAddAdminDefault(code int) *AddAdminDefault {
	return &AddAdminDefault{
		_statusCode: code,
	}
}

/*
AddAdminDefault describes a response with status code -1, with default header values.

Error
*/
type AddAdminDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this add admin default response has a 2xx status code
func (o *AddAdminDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this add admin default response has a 3xx status code
func (o *AddAdminDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this add admin default response has a 4xx status code
func (o *AddAdminDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this add admin default response has a 5xx status code
func (o *AddAdminDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this add admin default response a status code equal to that given
func (o *AddAdminDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the add admin default response
func (o *AddAdminDefault) Code() int {
	return o._statusCode
}

func (o *AddAdminDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /setting/admins][%d] addAdmin default %s", o._statusCode, payload)
}

func (o *AddAdminDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /setting/admins][%d] addAdmin default %s", o._statusCode, payload)
}

func (o *AddAdminDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddAdminDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
