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

// AddWidgetReader is a Reader for the AddWidget structure.
type AddWidgetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddWidgetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddWidgetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewAddWidgetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddWidgetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddWidgetOK creates a AddWidgetOK with default headers values
func NewAddWidgetOK() *AddWidgetOK {
	return &AddWidgetOK{}
}

/*
AddWidgetOK describes a response with status code 200, with default header values.

successful operation
*/
type AddWidgetOK struct {
	Payload models.Widget
}

// IsSuccess returns true when this add widget o k response has a 2xx status code
func (o *AddWidgetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add widget o k response has a 3xx status code
func (o *AddWidgetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add widget o k response has a 4xx status code
func (o *AddWidgetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add widget o k response has a 5xx status code
func (o *AddWidgetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add widget o k response a status code equal to that given
func (o *AddWidgetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add widget o k response
func (o *AddWidgetOK) Code() int {
	return 200
}

func (o *AddWidgetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dashboard/widgets][%d] addWidgetOK %s", 200, payload)
}

func (o *AddWidgetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dashboard/widgets][%d] addWidgetOK %s", 200, payload)
}

func (o *AddWidgetOK) GetPayload() models.Widget {
	return o.Payload
}

func (o *AddWidgetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalWidget(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewAddWidgetTooManyRequests creates a AddWidgetTooManyRequests with default headers values
func NewAddWidgetTooManyRequests() *AddWidgetTooManyRequests {
	return &AddWidgetTooManyRequests{}
}

/*
AddWidgetTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AddWidgetTooManyRequests struct {

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

// IsSuccess returns true when this add widget too many requests response has a 2xx status code
func (o *AddWidgetTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add widget too many requests response has a 3xx status code
func (o *AddWidgetTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add widget too many requests response has a 4xx status code
func (o *AddWidgetTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this add widget too many requests response has a 5xx status code
func (o *AddWidgetTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this add widget too many requests response a status code equal to that given
func (o *AddWidgetTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the add widget too many requests response
func (o *AddWidgetTooManyRequests) Code() int {
	return 429
}

func (o *AddWidgetTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /dashboard/widgets][%d] addWidgetTooManyRequests", 429)
}

func (o *AddWidgetTooManyRequests) String() string {
	return fmt.Sprintf("[POST /dashboard/widgets][%d] addWidgetTooManyRequests", 429)
}

func (o *AddWidgetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddWidgetDefault creates a AddWidgetDefault with default headers values
func NewAddWidgetDefault(code int) *AddWidgetDefault {
	return &AddWidgetDefault{
		_statusCode: code,
	}
}

/*
AddWidgetDefault describes a response with status code -1, with default header values.

Error
*/
type AddWidgetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this add widget default response has a 2xx status code
func (o *AddWidgetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this add widget default response has a 3xx status code
func (o *AddWidgetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this add widget default response has a 4xx status code
func (o *AddWidgetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this add widget default response has a 5xx status code
func (o *AddWidgetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this add widget default response a status code equal to that given
func (o *AddWidgetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the add widget default response
func (o *AddWidgetDefault) Code() int {
	return o._statusCode
}

func (o *AddWidgetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dashboard/widgets][%d] addWidget default %s", o._statusCode, payload)
}

func (o *AddWidgetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dashboard/widgets][%d] addWidget default %s", o._statusCode, payload)
}

func (o *AddWidgetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddWidgetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
