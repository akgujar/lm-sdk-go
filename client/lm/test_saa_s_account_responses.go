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

// TestSaaSAccountReader is a Reader for the TestSaaSAccount structure.
type TestSaaSAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestSaaSAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTestSaaSAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewTestSaaSAccountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewTestSaaSAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTestSaaSAccountOK creates a TestSaaSAccountOK with default headers values
func NewTestSaaSAccountOK() *TestSaaSAccountOK {
	return &TestSaaSAccountOK{}
}

/*
TestSaaSAccountOK describes a response with status code 200, with default header values.

successful operation
*/
type TestSaaSAccountOK struct {
	Payload models.RestCloudOkPermissionsV3
}

// IsSuccess returns true when this test saa s account o k response has a 2xx status code
func (o *TestSaaSAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this test saa s account o k response has a 3xx status code
func (o *TestSaaSAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this test saa s account o k response has a 4xx status code
func (o *TestSaaSAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this test saa s account o k response has a 5xx status code
func (o *TestSaaSAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this test saa s account o k response a status code equal to that given
func (o *TestSaaSAccountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the test saa s account o k response
func (o *TestSaaSAccountOK) Code() int {
	return 200
}

func (o *TestSaaSAccountOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /saas/functions/testAccount][%d] testSaaSAccountOK %s", 200, payload)
}

func (o *TestSaaSAccountOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /saas/functions/testAccount][%d] testSaaSAccountOK %s", 200, payload)
}

func (o *TestSaaSAccountOK) GetPayload() models.RestCloudOkPermissionsV3 {
	return o.Payload
}

func (o *TestSaaSAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestSaaSAccountTooManyRequests creates a TestSaaSAccountTooManyRequests with default headers values
func NewTestSaaSAccountTooManyRequests() *TestSaaSAccountTooManyRequests {
	return &TestSaaSAccountTooManyRequests{}
}

/*
TestSaaSAccountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type TestSaaSAccountTooManyRequests struct {

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

// IsSuccess returns true when this test saa s account too many requests response has a 2xx status code
func (o *TestSaaSAccountTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this test saa s account too many requests response has a 3xx status code
func (o *TestSaaSAccountTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this test saa s account too many requests response has a 4xx status code
func (o *TestSaaSAccountTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this test saa s account too many requests response has a 5xx status code
func (o *TestSaaSAccountTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this test saa s account too many requests response a status code equal to that given
func (o *TestSaaSAccountTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the test saa s account too many requests response
func (o *TestSaaSAccountTooManyRequests) Code() int {
	return 429
}

func (o *TestSaaSAccountTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /saas/functions/testAccount][%d] testSaaSAccountTooManyRequests", 429)
}

func (o *TestSaaSAccountTooManyRequests) String() string {
	return fmt.Sprintf("[POST /saas/functions/testAccount][%d] testSaaSAccountTooManyRequests", 429)
}

func (o *TestSaaSAccountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewTestSaaSAccountDefault creates a TestSaaSAccountDefault with default headers values
func NewTestSaaSAccountDefault(code int) *TestSaaSAccountDefault {
	return &TestSaaSAccountDefault{
		_statusCode: code,
	}
}

/*
TestSaaSAccountDefault describes a response with status code -1, with default header values.

Error
*/
type TestSaaSAccountDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this test saa s account default response has a 2xx status code
func (o *TestSaaSAccountDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this test saa s account default response has a 3xx status code
func (o *TestSaaSAccountDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this test saa s account default response has a 4xx status code
func (o *TestSaaSAccountDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this test saa s account default response has a 5xx status code
func (o *TestSaaSAccountDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this test saa s account default response a status code equal to that given
func (o *TestSaaSAccountDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the test saa s account default response
func (o *TestSaaSAccountDefault) Code() int {
	return o._statusCode
}

func (o *TestSaaSAccountDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /saas/functions/testAccount][%d] testSaaSAccount default %s", o._statusCode, payload)
}

func (o *TestSaaSAccountDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /saas/functions/testAccount][%d] testSaaSAccount default %s", o._statusCode, payload)
}

func (o *TestSaaSAccountDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TestSaaSAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}