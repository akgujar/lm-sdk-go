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

// GetAlertRuleByIDJSONReader is a Reader for the GetAlertRuleByIDJSON structure.
type GetAlertRuleByIDJSONReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertRuleByIDJSONReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertRuleByIDJSONOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetAlertRuleByIDJSONTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAlertRuleByIDJSONDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAlertRuleByIDJSONOK creates a GetAlertRuleByIDJSONOK with default headers values
func NewGetAlertRuleByIDJSONOK() *GetAlertRuleByIDJSONOK {
	return &GetAlertRuleByIDJSONOK{}
}

/*
GetAlertRuleByIDJSONOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAlertRuleByIDJSONOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get alert rule by Id Json o k response has a 2xx status code
func (o *GetAlertRuleByIDJSONOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get alert rule by Id Json o k response has a 3xx status code
func (o *GetAlertRuleByIDJSONOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert rule by Id Json o k response has a 4xx status code
func (o *GetAlertRuleByIDJSONOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alert rule by Id Json o k response has a 5xx status code
func (o *GetAlertRuleByIDJSONOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert rule by Id Json o k response a status code equal to that given
func (o *GetAlertRuleByIDJSONOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get alert rule by Id Json o k response
func (o *GetAlertRuleByIDJSONOK) Code() int {
	return 200
}

func (o *GetAlertRuleByIDJSONOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/alert/rules/{id}?__json=][%d] getAlertRuleByIdJsonOK %s", 200, payload)
}

func (o *GetAlertRuleByIDJSONOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/alert/rules/{id}?__json=][%d] getAlertRuleByIdJsonOK %s", 200, payload)
}

func (o *GetAlertRuleByIDJSONOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetAlertRuleByIDJSONOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertRuleByIDJSONTooManyRequests creates a GetAlertRuleByIDJSONTooManyRequests with default headers values
func NewGetAlertRuleByIDJSONTooManyRequests() *GetAlertRuleByIDJSONTooManyRequests {
	return &GetAlertRuleByIDJSONTooManyRequests{}
}

/*
GetAlertRuleByIDJSONTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetAlertRuleByIDJSONTooManyRequests struct {

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

// IsSuccess returns true when this get alert rule by Id Json too many requests response has a 2xx status code
func (o *GetAlertRuleByIDJSONTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alert rule by Id Json too many requests response has a 3xx status code
func (o *GetAlertRuleByIDJSONTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert rule by Id Json too many requests response has a 4xx status code
func (o *GetAlertRuleByIDJSONTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alert rule by Id Json too many requests response has a 5xx status code
func (o *GetAlertRuleByIDJSONTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert rule by Id Json too many requests response a status code equal to that given
func (o *GetAlertRuleByIDJSONTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get alert rule by Id Json too many requests response
func (o *GetAlertRuleByIDJSONTooManyRequests) Code() int {
	return 429
}

func (o *GetAlertRuleByIDJSONTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /setting/alert/rules/{id}?__json=][%d] getAlertRuleByIdJsonTooManyRequests", 429)
}

func (o *GetAlertRuleByIDJSONTooManyRequests) String() string {
	return fmt.Sprintf("[GET /setting/alert/rules/{id}?__json=][%d] getAlertRuleByIdJsonTooManyRequests", 429)
}

func (o *GetAlertRuleByIDJSONTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAlertRuleByIDJSONDefault creates a GetAlertRuleByIDJSONDefault with default headers values
func NewGetAlertRuleByIDJSONDefault(code int) *GetAlertRuleByIDJSONDefault {
	return &GetAlertRuleByIDJSONDefault{
		_statusCode: code,
	}
}

/*
GetAlertRuleByIDJSONDefault describes a response with status code -1, with default header values.

Error
*/
type GetAlertRuleByIDJSONDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get alert rule by Id Json default response has a 2xx status code
func (o *GetAlertRuleByIDJSONDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get alert rule by Id Json default response has a 3xx status code
func (o *GetAlertRuleByIDJSONDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get alert rule by Id Json default response has a 4xx status code
func (o *GetAlertRuleByIDJSONDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get alert rule by Id Json default response has a 5xx status code
func (o *GetAlertRuleByIDJSONDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get alert rule by Id Json default response a status code equal to that given
func (o *GetAlertRuleByIDJSONDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get alert rule by Id Json default response
func (o *GetAlertRuleByIDJSONDefault) Code() int {
	return o._statusCode
}

func (o *GetAlertRuleByIDJSONDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/alert/rules/{id}?__json=][%d] getAlertRuleByIdJson default %s", o._statusCode, payload)
}

func (o *GetAlertRuleByIDJSONDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/alert/rules/{id}?__json=][%d] getAlertRuleByIdJson default %s", o._statusCode, payload)
}

func (o *GetAlertRuleByIDJSONDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAlertRuleByIDJSONDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}