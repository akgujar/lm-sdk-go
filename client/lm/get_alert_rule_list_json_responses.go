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

// GetAlertRuleListJSONReader is a Reader for the GetAlertRuleListJSON structure.
type GetAlertRuleListJSONReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertRuleListJSONReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertRuleListJSONOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetAlertRuleListJSONTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAlertRuleListJSONDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAlertRuleListJSONOK creates a GetAlertRuleListJSONOK with default headers values
func NewGetAlertRuleListJSONOK() *GetAlertRuleListJSONOK {
	return &GetAlertRuleListJSONOK{}
}

/*
GetAlertRuleListJSONOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAlertRuleListJSONOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get alert rule list Json o k response has a 2xx status code
func (o *GetAlertRuleListJSONOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get alert rule list Json o k response has a 3xx status code
func (o *GetAlertRuleListJSONOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert rule list Json o k response has a 4xx status code
func (o *GetAlertRuleListJSONOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alert rule list Json o k response has a 5xx status code
func (o *GetAlertRuleListJSONOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert rule list Json o k response a status code equal to that given
func (o *GetAlertRuleListJSONOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get alert rule list Json o k response
func (o *GetAlertRuleListJSONOK) Code() int {
	return 200
}

func (o *GetAlertRuleListJSONOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/alert/rules?__json=][%d] getAlertRuleListJsonOK %s", 200, payload)
}

func (o *GetAlertRuleListJSONOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/alert/rules?__json=][%d] getAlertRuleListJsonOK %s", 200, payload)
}

func (o *GetAlertRuleListJSONOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetAlertRuleListJSONOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertRuleListJSONTooManyRequests creates a GetAlertRuleListJSONTooManyRequests with default headers values
func NewGetAlertRuleListJSONTooManyRequests() *GetAlertRuleListJSONTooManyRequests {
	return &GetAlertRuleListJSONTooManyRequests{}
}

/*
GetAlertRuleListJSONTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetAlertRuleListJSONTooManyRequests struct {

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

// IsSuccess returns true when this get alert rule list Json too many requests response has a 2xx status code
func (o *GetAlertRuleListJSONTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alert rule list Json too many requests response has a 3xx status code
func (o *GetAlertRuleListJSONTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert rule list Json too many requests response has a 4xx status code
func (o *GetAlertRuleListJSONTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alert rule list Json too many requests response has a 5xx status code
func (o *GetAlertRuleListJSONTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert rule list Json too many requests response a status code equal to that given
func (o *GetAlertRuleListJSONTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get alert rule list Json too many requests response
func (o *GetAlertRuleListJSONTooManyRequests) Code() int {
	return 429
}

func (o *GetAlertRuleListJSONTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /setting/alert/rules?__json=][%d] getAlertRuleListJsonTooManyRequests", 429)
}

func (o *GetAlertRuleListJSONTooManyRequests) String() string {
	return fmt.Sprintf("[GET /setting/alert/rules?__json=][%d] getAlertRuleListJsonTooManyRequests", 429)
}

func (o *GetAlertRuleListJSONTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAlertRuleListJSONDefault creates a GetAlertRuleListJSONDefault with default headers values
func NewGetAlertRuleListJSONDefault(code int) *GetAlertRuleListJSONDefault {
	return &GetAlertRuleListJSONDefault{
		_statusCode: code,
	}
}

/*
GetAlertRuleListJSONDefault describes a response with status code -1, with default header values.

Error
*/
type GetAlertRuleListJSONDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get alert rule list Json default response has a 2xx status code
func (o *GetAlertRuleListJSONDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get alert rule list Json default response has a 3xx status code
func (o *GetAlertRuleListJSONDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get alert rule list Json default response has a 4xx status code
func (o *GetAlertRuleListJSONDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get alert rule list Json default response has a 5xx status code
func (o *GetAlertRuleListJSONDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get alert rule list Json default response a status code equal to that given
func (o *GetAlertRuleListJSONDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get alert rule list Json default response
func (o *GetAlertRuleListJSONDefault) Code() int {
	return o._statusCode
}

func (o *GetAlertRuleListJSONDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/alert/rules?__json=][%d] getAlertRuleListJson default %s", o._statusCode, payload)
}

func (o *GetAlertRuleListJSONDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/alert/rules?__json=][%d] getAlertRuleListJson default %s", o._statusCode, payload)
}

func (o *GetAlertRuleListJSONDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAlertRuleListJSONDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
