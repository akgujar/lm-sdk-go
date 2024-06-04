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

// GetAuditLogListReader is a Reader for the GetAuditLogList structure.
type GetAuditLogListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuditLogListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuditLogListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetAuditLogListTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAuditLogListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAuditLogListOK creates a GetAuditLogListOK with default headers values
func NewGetAuditLogListOK() *GetAuditLogListOK {
	return &GetAuditLogListOK{}
}

/*
GetAuditLogListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAuditLogListOK struct {
	Payload *models.AccessLogPaginationResponse
}

// IsSuccess returns true when this get audit log list o k response has a 2xx status code
func (o *GetAuditLogListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get audit log list o k response has a 3xx status code
func (o *GetAuditLogListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get audit log list o k response has a 4xx status code
func (o *GetAuditLogListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get audit log list o k response has a 5xx status code
func (o *GetAuditLogListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get audit log list o k response a status code equal to that given
func (o *GetAuditLogListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get audit log list o k response
func (o *GetAuditLogListOK) Code() int {
	return 200
}

func (o *GetAuditLogListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/accesslogs][%d] getAuditLogListOK %s", 200, payload)
}

func (o *GetAuditLogListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/accesslogs][%d] getAuditLogListOK %s", 200, payload)
}

func (o *GetAuditLogListOK) GetPayload() *models.AccessLogPaginationResponse {
	return o.Payload
}

func (o *GetAuditLogListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessLogPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditLogListTooManyRequests creates a GetAuditLogListTooManyRequests with default headers values
func NewGetAuditLogListTooManyRequests() *GetAuditLogListTooManyRequests {
	return &GetAuditLogListTooManyRequests{}
}

/*
GetAuditLogListTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetAuditLogListTooManyRequests struct {

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

// IsSuccess returns true when this get audit log list too many requests response has a 2xx status code
func (o *GetAuditLogListTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get audit log list too many requests response has a 3xx status code
func (o *GetAuditLogListTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get audit log list too many requests response has a 4xx status code
func (o *GetAuditLogListTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get audit log list too many requests response has a 5xx status code
func (o *GetAuditLogListTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get audit log list too many requests response a status code equal to that given
func (o *GetAuditLogListTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get audit log list too many requests response
func (o *GetAuditLogListTooManyRequests) Code() int {
	return 429
}

func (o *GetAuditLogListTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /setting/accesslogs][%d] getAuditLogListTooManyRequests", 429)
}

func (o *GetAuditLogListTooManyRequests) String() string {
	return fmt.Sprintf("[GET /setting/accesslogs][%d] getAuditLogListTooManyRequests", 429)
}

func (o *GetAuditLogListTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAuditLogListDefault creates a GetAuditLogListDefault with default headers values
func NewGetAuditLogListDefault(code int) *GetAuditLogListDefault {
	return &GetAuditLogListDefault{
		_statusCode: code,
	}
}

/*
GetAuditLogListDefault describes a response with status code -1, with default header values.

Error
*/
type GetAuditLogListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get audit log list default response has a 2xx status code
func (o *GetAuditLogListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get audit log list default response has a 3xx status code
func (o *GetAuditLogListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get audit log list default response has a 4xx status code
func (o *GetAuditLogListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get audit log list default response has a 5xx status code
func (o *GetAuditLogListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get audit log list default response a status code equal to that given
func (o *GetAuditLogListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get audit log list default response
func (o *GetAuditLogListDefault) Code() int {
	return o._statusCode
}

func (o *GetAuditLogListDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/accesslogs][%d] getAuditLogList default %s", o._statusCode, payload)
}

func (o *GetAuditLogListDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/accesslogs][%d] getAuditLogList default %s", o._statusCode, payload)
}

func (o *GetAuditLogListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAuditLogListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
