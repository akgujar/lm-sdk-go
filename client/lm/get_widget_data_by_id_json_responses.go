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

// GetWidgetDataByIDJSONReader is a Reader for the GetWidgetDataByIDJSON structure.
type GetWidgetDataByIDJSONReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWidgetDataByIDJSONReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWidgetDataByIDJSONOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetWidgetDataByIDJSONTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetWidgetDataByIDJSONDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWidgetDataByIDJSONOK creates a GetWidgetDataByIDJSONOK with default headers values
func NewGetWidgetDataByIDJSONOK() *GetWidgetDataByIDJSONOK {
	return &GetWidgetDataByIDJSONOK{}
}

/*
GetWidgetDataByIDJSONOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWidgetDataByIDJSONOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get widget data by Id Json o k response has a 2xx status code
func (o *GetWidgetDataByIDJSONOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get widget data by Id Json o k response has a 3xx status code
func (o *GetWidgetDataByIDJSONOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get widget data by Id Json o k response has a 4xx status code
func (o *GetWidgetDataByIDJSONOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get widget data by Id Json o k response has a 5xx status code
func (o *GetWidgetDataByIDJSONOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get widget data by Id Json o k response a status code equal to that given
func (o *GetWidgetDataByIDJSONOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get widget data by Id Json o k response
func (o *GetWidgetDataByIDJSONOK) Code() int {
	return 200
}

func (o *GetWidgetDataByIDJSONOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboard/widgets/{id}/data?__json=][%d] getWidgetDataByIdJsonOK %s", 200, payload)
}

func (o *GetWidgetDataByIDJSONOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboard/widgets/{id}/data?__json=][%d] getWidgetDataByIdJsonOK %s", 200, payload)
}

func (o *GetWidgetDataByIDJSONOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetWidgetDataByIDJSONOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWidgetDataByIDJSONTooManyRequests creates a GetWidgetDataByIDJSONTooManyRequests with default headers values
func NewGetWidgetDataByIDJSONTooManyRequests() *GetWidgetDataByIDJSONTooManyRequests {
	return &GetWidgetDataByIDJSONTooManyRequests{}
}

/*
GetWidgetDataByIDJSONTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetWidgetDataByIDJSONTooManyRequests struct {

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

// IsSuccess returns true when this get widget data by Id Json too many requests response has a 2xx status code
func (o *GetWidgetDataByIDJSONTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get widget data by Id Json too many requests response has a 3xx status code
func (o *GetWidgetDataByIDJSONTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get widget data by Id Json too many requests response has a 4xx status code
func (o *GetWidgetDataByIDJSONTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get widget data by Id Json too many requests response has a 5xx status code
func (o *GetWidgetDataByIDJSONTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get widget data by Id Json too many requests response a status code equal to that given
func (o *GetWidgetDataByIDJSONTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get widget data by Id Json too many requests response
func (o *GetWidgetDataByIDJSONTooManyRequests) Code() int {
	return 429
}

func (o *GetWidgetDataByIDJSONTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /dashboard/widgets/{id}/data?__json=][%d] getWidgetDataByIdJsonTooManyRequests", 429)
}

func (o *GetWidgetDataByIDJSONTooManyRequests) String() string {
	return fmt.Sprintf("[GET /dashboard/widgets/{id}/data?__json=][%d] getWidgetDataByIdJsonTooManyRequests", 429)
}

func (o *GetWidgetDataByIDJSONTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetWidgetDataByIDJSONDefault creates a GetWidgetDataByIDJSONDefault with default headers values
func NewGetWidgetDataByIDJSONDefault(code int) *GetWidgetDataByIDJSONDefault {
	return &GetWidgetDataByIDJSONDefault{
		_statusCode: code,
	}
}

/*
GetWidgetDataByIDJSONDefault describes a response with status code -1, with default header values.

Error
*/
type GetWidgetDataByIDJSONDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get widget data by Id Json default response has a 2xx status code
func (o *GetWidgetDataByIDJSONDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get widget data by Id Json default response has a 3xx status code
func (o *GetWidgetDataByIDJSONDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get widget data by Id Json default response has a 4xx status code
func (o *GetWidgetDataByIDJSONDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get widget data by Id Json default response has a 5xx status code
func (o *GetWidgetDataByIDJSONDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get widget data by Id Json default response a status code equal to that given
func (o *GetWidgetDataByIDJSONDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get widget data by Id Json default response
func (o *GetWidgetDataByIDJSONDefault) Code() int {
	return o._statusCode
}

func (o *GetWidgetDataByIDJSONDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboard/widgets/{id}/data?__json=][%d] getWidgetDataByIdJson default %s", o._statusCode, payload)
}

func (o *GetWidgetDataByIDJSONDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboard/widgets/{id}/data?__json=][%d] getWidgetDataByIdJson default %s", o._statusCode, payload)
}

func (o *GetWidgetDataByIDJSONDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetWidgetDataByIDJSONDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
