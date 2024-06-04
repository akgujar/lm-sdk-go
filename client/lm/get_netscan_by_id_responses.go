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

// GetNetscanByIDReader is a Reader for the GetNetscanByID structure.
type GetNetscanByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetscanByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetscanByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetNetscanByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetNetscanByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetscanByIDOK creates a GetNetscanByIDOK with default headers values
func NewGetNetscanByIDOK() *GetNetscanByIDOK {
	return &GetNetscanByIDOK{}
}

/*
GetNetscanByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetNetscanByIDOK struct {
	Payload models.Netscan
}

// IsSuccess returns true when this get netscan by Id o k response has a 2xx status code
func (o *GetNetscanByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get netscan by Id o k response has a 3xx status code
func (o *GetNetscanByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get netscan by Id o k response has a 4xx status code
func (o *GetNetscanByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get netscan by Id o k response has a 5xx status code
func (o *GetNetscanByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get netscan by Id o k response a status code equal to that given
func (o *GetNetscanByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get netscan by Id o k response
func (o *GetNetscanByIDOK) Code() int {
	return 200
}

func (o *GetNetscanByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/netscans/{id}][%d] getNetscanByIdOK %s", 200, payload)
}

func (o *GetNetscanByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/netscans/{id}][%d] getNetscanByIdOK %s", 200, payload)
}

func (o *GetNetscanByIDOK) GetPayload() models.Netscan {
	return o.Payload
}

func (o *GetNetscanByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalNetscan(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetNetscanByIDTooManyRequests creates a GetNetscanByIDTooManyRequests with default headers values
func NewGetNetscanByIDTooManyRequests() *GetNetscanByIDTooManyRequests {
	return &GetNetscanByIDTooManyRequests{}
}

/*
GetNetscanByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetNetscanByIDTooManyRequests struct {

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

// IsSuccess returns true when this get netscan by Id too many requests response has a 2xx status code
func (o *GetNetscanByIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get netscan by Id too many requests response has a 3xx status code
func (o *GetNetscanByIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get netscan by Id too many requests response has a 4xx status code
func (o *GetNetscanByIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get netscan by Id too many requests response has a 5xx status code
func (o *GetNetscanByIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get netscan by Id too many requests response a status code equal to that given
func (o *GetNetscanByIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get netscan by Id too many requests response
func (o *GetNetscanByIDTooManyRequests) Code() int {
	return 429
}

func (o *GetNetscanByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /setting/netscans/{id}][%d] getNetscanByIdTooManyRequests", 429)
}

func (o *GetNetscanByIDTooManyRequests) String() string {
	return fmt.Sprintf("[GET /setting/netscans/{id}][%d] getNetscanByIdTooManyRequests", 429)
}

func (o *GetNetscanByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetNetscanByIDDefault creates a GetNetscanByIDDefault with default headers values
func NewGetNetscanByIDDefault(code int) *GetNetscanByIDDefault {
	return &GetNetscanByIDDefault{
		_statusCode: code,
	}
}

/*
GetNetscanByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetNetscanByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get netscan by Id default response has a 2xx status code
func (o *GetNetscanByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get netscan by Id default response has a 3xx status code
func (o *GetNetscanByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get netscan by Id default response has a 4xx status code
func (o *GetNetscanByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get netscan by Id default response has a 5xx status code
func (o *GetNetscanByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get netscan by Id default response a status code equal to that given
func (o *GetNetscanByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get netscan by Id default response
func (o *GetNetscanByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetNetscanByIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/netscans/{id}][%d] getNetscanById default %s", o._statusCode, payload)
}

func (o *GetNetscanByIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/netscans/{id}][%d] getNetscanById default %s", o._statusCode, payload)
}

func (o *GetNetscanByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetNetscanByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
