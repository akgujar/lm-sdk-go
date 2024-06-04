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

// GetAwsExternalIDReader is a Reader for the GetAwsExternalID structure.
type GetAwsExternalIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAwsExternalIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAwsExternalIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetAwsExternalIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAwsExternalIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAwsExternalIDOK creates a GetAwsExternalIDOK with default headers values
func NewGetAwsExternalIDOK() *GetAwsExternalIDOK {
	return &GetAwsExternalIDOK{}
}

/*
GetAwsExternalIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAwsExternalIDOK struct {
	Payload *models.AwsExternalID
}

// IsSuccess returns true when this get aws external Id o k response has a 2xx status code
func (o *GetAwsExternalIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get aws external Id o k response has a 3xx status code
func (o *GetAwsExternalIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get aws external Id o k response has a 4xx status code
func (o *GetAwsExternalIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get aws external Id o k response has a 5xx status code
func (o *GetAwsExternalIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get aws external Id o k response a status code equal to that given
func (o *GetAwsExternalIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get aws external Id o k response
func (o *GetAwsExternalIDOK) Code() int {
	return 200
}

func (o *GetAwsExternalIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /aws/externalId][%d] getAwsExternalIdOK %s", 200, payload)
}

func (o *GetAwsExternalIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /aws/externalId][%d] getAwsExternalIdOK %s", 200, payload)
}

func (o *GetAwsExternalIDOK) GetPayload() *models.AwsExternalID {
	return o.Payload
}

func (o *GetAwsExternalIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AwsExternalID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAwsExternalIDTooManyRequests creates a GetAwsExternalIDTooManyRequests with default headers values
func NewGetAwsExternalIDTooManyRequests() *GetAwsExternalIDTooManyRequests {
	return &GetAwsExternalIDTooManyRequests{}
}

/*
GetAwsExternalIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetAwsExternalIDTooManyRequests struct {

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

// IsSuccess returns true when this get aws external Id too many requests response has a 2xx status code
func (o *GetAwsExternalIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get aws external Id too many requests response has a 3xx status code
func (o *GetAwsExternalIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get aws external Id too many requests response has a 4xx status code
func (o *GetAwsExternalIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get aws external Id too many requests response has a 5xx status code
func (o *GetAwsExternalIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get aws external Id too many requests response a status code equal to that given
func (o *GetAwsExternalIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get aws external Id too many requests response
func (o *GetAwsExternalIDTooManyRequests) Code() int {
	return 429
}

func (o *GetAwsExternalIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /aws/externalId][%d] getAwsExternalIdTooManyRequests", 429)
}

func (o *GetAwsExternalIDTooManyRequests) String() string {
	return fmt.Sprintf("[GET /aws/externalId][%d] getAwsExternalIdTooManyRequests", 429)
}

func (o *GetAwsExternalIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAwsExternalIDDefault creates a GetAwsExternalIDDefault with default headers values
func NewGetAwsExternalIDDefault(code int) *GetAwsExternalIDDefault {
	return &GetAwsExternalIDDefault{
		_statusCode: code,
	}
}

/*
GetAwsExternalIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetAwsExternalIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get aws external Id default response has a 2xx status code
func (o *GetAwsExternalIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get aws external Id default response has a 3xx status code
func (o *GetAwsExternalIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get aws external Id default response has a 4xx status code
func (o *GetAwsExternalIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get aws external Id default response has a 5xx status code
func (o *GetAwsExternalIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get aws external Id default response a status code equal to that given
func (o *GetAwsExternalIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get aws external Id default response
func (o *GetAwsExternalIDDefault) Code() int {
	return o._statusCode
}

func (o *GetAwsExternalIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /aws/externalId][%d] getAwsExternalId default %s", o._statusCode, payload)
}

func (o *GetAwsExternalIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /aws/externalId][%d] getAwsExternalId default %s", o._statusCode, payload)
}

func (o *GetAwsExternalIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAwsExternalIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
