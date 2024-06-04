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

// GetCollectorListReader is a Reader for the GetCollectorList structure.
type GetCollectorListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCollectorListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCollectorListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetCollectorListTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCollectorListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCollectorListOK creates a GetCollectorListOK with default headers values
func NewGetCollectorListOK() *GetCollectorListOK {
	return &GetCollectorListOK{}
}

/*
GetCollectorListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetCollectorListOK struct {
	Payload *models.CollectorPaginationResponse
}

// IsSuccess returns true when this get collector list o k response has a 2xx status code
func (o *GetCollectorListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get collector list o k response has a 3xx status code
func (o *GetCollectorListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get collector list o k response has a 4xx status code
func (o *GetCollectorListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get collector list o k response has a 5xx status code
func (o *GetCollectorListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get collector list o k response a status code equal to that given
func (o *GetCollectorListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get collector list o k response
func (o *GetCollectorListOK) Code() int {
	return 200
}

func (o *GetCollectorListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/collector/collectors][%d] getCollectorListOK %s", 200, payload)
}

func (o *GetCollectorListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/collector/collectors][%d] getCollectorListOK %s", 200, payload)
}

func (o *GetCollectorListOK) GetPayload() *models.CollectorPaginationResponse {
	return o.Payload
}

func (o *GetCollectorListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CollectorPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCollectorListTooManyRequests creates a GetCollectorListTooManyRequests with default headers values
func NewGetCollectorListTooManyRequests() *GetCollectorListTooManyRequests {
	return &GetCollectorListTooManyRequests{}
}

/*
GetCollectorListTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetCollectorListTooManyRequests struct {

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

// IsSuccess returns true when this get collector list too many requests response has a 2xx status code
func (o *GetCollectorListTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get collector list too many requests response has a 3xx status code
func (o *GetCollectorListTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get collector list too many requests response has a 4xx status code
func (o *GetCollectorListTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get collector list too many requests response has a 5xx status code
func (o *GetCollectorListTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get collector list too many requests response a status code equal to that given
func (o *GetCollectorListTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get collector list too many requests response
func (o *GetCollectorListTooManyRequests) Code() int {
	return 429
}

func (o *GetCollectorListTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /setting/collector/collectors][%d] getCollectorListTooManyRequests", 429)
}

func (o *GetCollectorListTooManyRequests) String() string {
	return fmt.Sprintf("[GET /setting/collector/collectors][%d] getCollectorListTooManyRequests", 429)
}

func (o *GetCollectorListTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCollectorListDefault creates a GetCollectorListDefault with default headers values
func NewGetCollectorListDefault(code int) *GetCollectorListDefault {
	return &GetCollectorListDefault{
		_statusCode: code,
	}
}

/*
GetCollectorListDefault describes a response with status code -1, with default header values.

Error
*/
type GetCollectorListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get collector list default response has a 2xx status code
func (o *GetCollectorListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get collector list default response has a 3xx status code
func (o *GetCollectorListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get collector list default response has a 4xx status code
func (o *GetCollectorListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get collector list default response has a 5xx status code
func (o *GetCollectorListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get collector list default response a status code equal to that given
func (o *GetCollectorListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get collector list default response
func (o *GetCollectorListDefault) Code() int {
	return o._statusCode
}

func (o *GetCollectorListDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/collector/collectors][%d] getCollectorList default %s", o._statusCode, payload)
}

func (o *GetCollectorListDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/collector/collectors][%d] getCollectorList default %s", o._statusCode, payload)
}

func (o *GetCollectorListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetCollectorListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
