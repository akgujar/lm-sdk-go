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

// GetAlertRuleListReader is a Reader for the GetAlertRuleList structure.
type GetAlertRuleListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertRuleListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertRuleListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetAlertRuleListTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAlertRuleListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAlertRuleListOK creates a GetAlertRuleListOK with default headers values
func NewGetAlertRuleListOK() *GetAlertRuleListOK {
	return &GetAlertRuleListOK{}
}

/*
GetAlertRuleListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAlertRuleListOK struct {
	Payload *models.AlertRulePaginationResponse
}

// IsSuccess returns true when this get alert rule list o k response has a 2xx status code
func (o *GetAlertRuleListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get alert rule list o k response has a 3xx status code
func (o *GetAlertRuleListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert rule list o k response has a 4xx status code
func (o *GetAlertRuleListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alert rule list o k response has a 5xx status code
func (o *GetAlertRuleListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert rule list o k response a status code equal to that given
func (o *GetAlertRuleListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get alert rule list o k response
func (o *GetAlertRuleListOK) Code() int {
	return 200
}

func (o *GetAlertRuleListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/alert/rules][%d] getAlertRuleListOK %s", 200, payload)
}

func (o *GetAlertRuleListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/alert/rules][%d] getAlertRuleListOK %s", 200, payload)
}

func (o *GetAlertRuleListOK) GetPayload() *models.AlertRulePaginationResponse {
	return o.Payload
}

func (o *GetAlertRuleListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertRulePaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertRuleListTooManyRequests creates a GetAlertRuleListTooManyRequests with default headers values
func NewGetAlertRuleListTooManyRequests() *GetAlertRuleListTooManyRequests {
	return &GetAlertRuleListTooManyRequests{}
}

/*
GetAlertRuleListTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetAlertRuleListTooManyRequests struct {

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

// IsSuccess returns true when this get alert rule list too many requests response has a 2xx status code
func (o *GetAlertRuleListTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alert rule list too many requests response has a 3xx status code
func (o *GetAlertRuleListTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert rule list too many requests response has a 4xx status code
func (o *GetAlertRuleListTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alert rule list too many requests response has a 5xx status code
func (o *GetAlertRuleListTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert rule list too many requests response a status code equal to that given
func (o *GetAlertRuleListTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get alert rule list too many requests response
func (o *GetAlertRuleListTooManyRequests) Code() int {
	return 429
}

func (o *GetAlertRuleListTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /setting/alert/rules][%d] getAlertRuleListTooManyRequests", 429)
}

func (o *GetAlertRuleListTooManyRequests) String() string {
	return fmt.Sprintf("[GET /setting/alert/rules][%d] getAlertRuleListTooManyRequests", 429)
}

func (o *GetAlertRuleListTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAlertRuleListDefault creates a GetAlertRuleListDefault with default headers values
func NewGetAlertRuleListDefault(code int) *GetAlertRuleListDefault {
	return &GetAlertRuleListDefault{
		_statusCode: code,
	}
}

/*
GetAlertRuleListDefault describes a response with status code -1, with default header values.

Error
*/
type GetAlertRuleListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get alert rule list default response has a 2xx status code
func (o *GetAlertRuleListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get alert rule list default response has a 3xx status code
func (o *GetAlertRuleListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get alert rule list default response has a 4xx status code
func (o *GetAlertRuleListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get alert rule list default response has a 5xx status code
func (o *GetAlertRuleListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get alert rule list default response a status code equal to that given
func (o *GetAlertRuleListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get alert rule list default response
func (o *GetAlertRuleListDefault) Code() int {
	return o._statusCode
}

func (o *GetAlertRuleListDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/alert/rules][%d] getAlertRuleList default %s", o._statusCode, payload)
}

func (o *GetAlertRuleListDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/alert/rules][%d] getAlertRuleList default %s", o._statusCode, payload)
}

func (o *GetAlertRuleListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAlertRuleListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
