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

// GetReportGroupByIDJSONReader is a Reader for the GetReportGroupByIDJSON structure.
type GetReportGroupByIDJSONReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReportGroupByIDJSONReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReportGroupByIDJSONOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetReportGroupByIDJSONTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetReportGroupByIDJSONDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetReportGroupByIDJSONOK creates a GetReportGroupByIDJSONOK with default headers values
func NewGetReportGroupByIDJSONOK() *GetReportGroupByIDJSONOK {
	return &GetReportGroupByIDJSONOK{}
}

/*
GetReportGroupByIDJSONOK describes a response with status code 200, with default header values.

successful operation
*/
type GetReportGroupByIDJSONOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get report group by Id Json o k response has a 2xx status code
func (o *GetReportGroupByIDJSONOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get report group by Id Json o k response has a 3xx status code
func (o *GetReportGroupByIDJSONOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report group by Id Json o k response has a 4xx status code
func (o *GetReportGroupByIDJSONOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get report group by Id Json o k response has a 5xx status code
func (o *GetReportGroupByIDJSONOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get report group by Id Json o k response a status code equal to that given
func (o *GetReportGroupByIDJSONOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get report group by Id Json o k response
func (o *GetReportGroupByIDJSONOK) Code() int {
	return 200
}

func (o *GetReportGroupByIDJSONOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /report/groups/{id}?__json=][%d] getReportGroupByIdJsonOK %s", 200, payload)
}

func (o *GetReportGroupByIDJSONOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /report/groups/{id}?__json=][%d] getReportGroupByIdJsonOK %s", 200, payload)
}

func (o *GetReportGroupByIDJSONOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetReportGroupByIDJSONOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportGroupByIDJSONTooManyRequests creates a GetReportGroupByIDJSONTooManyRequests with default headers values
func NewGetReportGroupByIDJSONTooManyRequests() *GetReportGroupByIDJSONTooManyRequests {
	return &GetReportGroupByIDJSONTooManyRequests{}
}

/*
GetReportGroupByIDJSONTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetReportGroupByIDJSONTooManyRequests struct {

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

// IsSuccess returns true when this get report group by Id Json too many requests response has a 2xx status code
func (o *GetReportGroupByIDJSONTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report group by Id Json too many requests response has a 3xx status code
func (o *GetReportGroupByIDJSONTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report group by Id Json too many requests response has a 4xx status code
func (o *GetReportGroupByIDJSONTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get report group by Id Json too many requests response has a 5xx status code
func (o *GetReportGroupByIDJSONTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get report group by Id Json too many requests response a status code equal to that given
func (o *GetReportGroupByIDJSONTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get report group by Id Json too many requests response
func (o *GetReportGroupByIDJSONTooManyRequests) Code() int {
	return 429
}

func (o *GetReportGroupByIDJSONTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /report/groups/{id}?__json=][%d] getReportGroupByIdJsonTooManyRequests", 429)
}

func (o *GetReportGroupByIDJSONTooManyRequests) String() string {
	return fmt.Sprintf("[GET /report/groups/{id}?__json=][%d] getReportGroupByIdJsonTooManyRequests", 429)
}

func (o *GetReportGroupByIDJSONTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetReportGroupByIDJSONDefault creates a GetReportGroupByIDJSONDefault with default headers values
func NewGetReportGroupByIDJSONDefault(code int) *GetReportGroupByIDJSONDefault {
	return &GetReportGroupByIDJSONDefault{
		_statusCode: code,
	}
}

/*
GetReportGroupByIDJSONDefault describes a response with status code -1, with default header values.

Error
*/
type GetReportGroupByIDJSONDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get report group by Id Json default response has a 2xx status code
func (o *GetReportGroupByIDJSONDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get report group by Id Json default response has a 3xx status code
func (o *GetReportGroupByIDJSONDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get report group by Id Json default response has a 4xx status code
func (o *GetReportGroupByIDJSONDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get report group by Id Json default response has a 5xx status code
func (o *GetReportGroupByIDJSONDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get report group by Id Json default response a status code equal to that given
func (o *GetReportGroupByIDJSONDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get report group by Id Json default response
func (o *GetReportGroupByIDJSONDefault) Code() int {
	return o._statusCode
}

func (o *GetReportGroupByIDJSONDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /report/groups/{id}?__json=][%d] getReportGroupByIdJson default %s", o._statusCode, payload)
}

func (o *GetReportGroupByIDJSONDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /report/groups/{id}?__json=][%d] getReportGroupByIdJson default %s", o._statusCode, payload)
}

func (o *GetReportGroupByIDJSONDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetReportGroupByIDJSONDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}