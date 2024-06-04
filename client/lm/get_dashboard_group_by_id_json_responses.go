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

// GetDashboardGroupByIDJSONReader is a Reader for the GetDashboardGroupByIDJSON structure.
type GetDashboardGroupByIDJSONReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDashboardGroupByIDJSONReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDashboardGroupByIDJSONOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetDashboardGroupByIDJSONTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetDashboardGroupByIDJSONDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDashboardGroupByIDJSONOK creates a GetDashboardGroupByIDJSONOK with default headers values
func NewGetDashboardGroupByIDJSONOK() *GetDashboardGroupByIDJSONOK {
	return &GetDashboardGroupByIDJSONOK{}
}

/*
GetDashboardGroupByIDJSONOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDashboardGroupByIDJSONOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get dashboard group by Id Json o k response has a 2xx status code
func (o *GetDashboardGroupByIDJSONOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get dashboard group by Id Json o k response has a 3xx status code
func (o *GetDashboardGroupByIDJSONOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dashboard group by Id Json o k response has a 4xx status code
func (o *GetDashboardGroupByIDJSONOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dashboard group by Id Json o k response has a 5xx status code
func (o *GetDashboardGroupByIDJSONOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get dashboard group by Id Json o k response a status code equal to that given
func (o *GetDashboardGroupByIDJSONOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get dashboard group by Id Json o k response
func (o *GetDashboardGroupByIDJSONOK) Code() int {
	return 200
}

func (o *GetDashboardGroupByIDJSONOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboard/groups/{id}?__json=][%d] getDashboardGroupByIdJsonOK %s", 200, payload)
}

func (o *GetDashboardGroupByIDJSONOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboard/groups/{id}?__json=][%d] getDashboardGroupByIdJsonOK %s", 200, payload)
}

func (o *GetDashboardGroupByIDJSONOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetDashboardGroupByIDJSONOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardGroupByIDJSONTooManyRequests creates a GetDashboardGroupByIDJSONTooManyRequests with default headers values
func NewGetDashboardGroupByIDJSONTooManyRequests() *GetDashboardGroupByIDJSONTooManyRequests {
	return &GetDashboardGroupByIDJSONTooManyRequests{}
}

/*
GetDashboardGroupByIDJSONTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetDashboardGroupByIDJSONTooManyRequests struct {

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

// IsSuccess returns true when this get dashboard group by Id Json too many requests response has a 2xx status code
func (o *GetDashboardGroupByIDJSONTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dashboard group by Id Json too many requests response has a 3xx status code
func (o *GetDashboardGroupByIDJSONTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dashboard group by Id Json too many requests response has a 4xx status code
func (o *GetDashboardGroupByIDJSONTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dashboard group by Id Json too many requests response has a 5xx status code
func (o *GetDashboardGroupByIDJSONTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get dashboard group by Id Json too many requests response a status code equal to that given
func (o *GetDashboardGroupByIDJSONTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get dashboard group by Id Json too many requests response
func (o *GetDashboardGroupByIDJSONTooManyRequests) Code() int {
	return 429
}

func (o *GetDashboardGroupByIDJSONTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /dashboard/groups/{id}?__json=][%d] getDashboardGroupByIdJsonTooManyRequests", 429)
}

func (o *GetDashboardGroupByIDJSONTooManyRequests) String() string {
	return fmt.Sprintf("[GET /dashboard/groups/{id}?__json=][%d] getDashboardGroupByIdJsonTooManyRequests", 429)
}

func (o *GetDashboardGroupByIDJSONTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDashboardGroupByIDJSONDefault creates a GetDashboardGroupByIDJSONDefault with default headers values
func NewGetDashboardGroupByIDJSONDefault(code int) *GetDashboardGroupByIDJSONDefault {
	return &GetDashboardGroupByIDJSONDefault{
		_statusCode: code,
	}
}

/*
GetDashboardGroupByIDJSONDefault describes a response with status code -1, with default header values.

Error
*/
type GetDashboardGroupByIDJSONDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get dashboard group by Id Json default response has a 2xx status code
func (o *GetDashboardGroupByIDJSONDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get dashboard group by Id Json default response has a 3xx status code
func (o *GetDashboardGroupByIDJSONDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get dashboard group by Id Json default response has a 4xx status code
func (o *GetDashboardGroupByIDJSONDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get dashboard group by Id Json default response has a 5xx status code
func (o *GetDashboardGroupByIDJSONDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get dashboard group by Id Json default response a status code equal to that given
func (o *GetDashboardGroupByIDJSONDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get dashboard group by Id Json default response
func (o *GetDashboardGroupByIDJSONDefault) Code() int {
	return o._statusCode
}

func (o *GetDashboardGroupByIDJSONDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboard/groups/{id}?__json=][%d] getDashboardGroupByIdJson default %s", o._statusCode, payload)
}

func (o *GetDashboardGroupByIDJSONDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboard/groups/{id}?__json=][%d] getDashboardGroupByIdJson default %s", o._statusCode, payload)
}

func (o *GetDashboardGroupByIDJSONDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDashboardGroupByIDJSONDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
