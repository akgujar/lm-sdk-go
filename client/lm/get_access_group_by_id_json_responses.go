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

// GetAccessGroupByIDJSONReader is a Reader for the GetAccessGroupByIDJSON structure.
type GetAccessGroupByIDJSONReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccessGroupByIDJSONReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccessGroupByIDJSONOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetAccessGroupByIDJSONTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAccessGroupByIDJSONDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAccessGroupByIDJSONOK creates a GetAccessGroupByIDJSONOK with default headers values
func NewGetAccessGroupByIDJSONOK() *GetAccessGroupByIDJSONOK {
	return &GetAccessGroupByIDJSONOK{}
}

/*
GetAccessGroupByIDJSONOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAccessGroupByIDJSONOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get access group by Id Json o k response has a 2xx status code
func (o *GetAccessGroupByIDJSONOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get access group by Id Json o k response has a 3xx status code
func (o *GetAccessGroupByIDJSONOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get access group by Id Json o k response has a 4xx status code
func (o *GetAccessGroupByIDJSONOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get access group by Id Json o k response has a 5xx status code
func (o *GetAccessGroupByIDJSONOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get access group by Id Json o k response a status code equal to that given
func (o *GetAccessGroupByIDJSONOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get access group by Id Json o k response
func (o *GetAccessGroupByIDJSONOK) Code() int {
	return 200
}

func (o *GetAccessGroupByIDJSONOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/accessgroup/{id}?__json=][%d] getAccessGroupByIdJsonOK %s", 200, payload)
}

func (o *GetAccessGroupByIDJSONOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/accessgroup/{id}?__json=][%d] getAccessGroupByIdJsonOK %s", 200, payload)
}

func (o *GetAccessGroupByIDJSONOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetAccessGroupByIDJSONOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccessGroupByIDJSONTooManyRequests creates a GetAccessGroupByIDJSONTooManyRequests with default headers values
func NewGetAccessGroupByIDJSONTooManyRequests() *GetAccessGroupByIDJSONTooManyRequests {
	return &GetAccessGroupByIDJSONTooManyRequests{}
}

/*
GetAccessGroupByIDJSONTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetAccessGroupByIDJSONTooManyRequests struct {

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

// IsSuccess returns true when this get access group by Id Json too many requests response has a 2xx status code
func (o *GetAccessGroupByIDJSONTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get access group by Id Json too many requests response has a 3xx status code
func (o *GetAccessGroupByIDJSONTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get access group by Id Json too many requests response has a 4xx status code
func (o *GetAccessGroupByIDJSONTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get access group by Id Json too many requests response has a 5xx status code
func (o *GetAccessGroupByIDJSONTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get access group by Id Json too many requests response a status code equal to that given
func (o *GetAccessGroupByIDJSONTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get access group by Id Json too many requests response
func (o *GetAccessGroupByIDJSONTooManyRequests) Code() int {
	return 429
}

func (o *GetAccessGroupByIDJSONTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /setting/accessgroup/{id}?__json=][%d] getAccessGroupByIdJsonTooManyRequests", 429)
}

func (o *GetAccessGroupByIDJSONTooManyRequests) String() string {
	return fmt.Sprintf("[GET /setting/accessgroup/{id}?__json=][%d] getAccessGroupByIdJsonTooManyRequests", 429)
}

func (o *GetAccessGroupByIDJSONTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAccessGroupByIDJSONDefault creates a GetAccessGroupByIDJSONDefault with default headers values
func NewGetAccessGroupByIDJSONDefault(code int) *GetAccessGroupByIDJSONDefault {
	return &GetAccessGroupByIDJSONDefault{
		_statusCode: code,
	}
}

/*
GetAccessGroupByIDJSONDefault describes a response with status code -1, with default header values.

Error
*/
type GetAccessGroupByIDJSONDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get access group by Id Json default response has a 2xx status code
func (o *GetAccessGroupByIDJSONDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get access group by Id Json default response has a 3xx status code
func (o *GetAccessGroupByIDJSONDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get access group by Id Json default response has a 4xx status code
func (o *GetAccessGroupByIDJSONDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get access group by Id Json default response has a 5xx status code
func (o *GetAccessGroupByIDJSONDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get access group by Id Json default response a status code equal to that given
func (o *GetAccessGroupByIDJSONDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get access group by Id Json default response
func (o *GetAccessGroupByIDJSONDefault) Code() int {
	return o._statusCode
}

func (o *GetAccessGroupByIDJSONDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/accessgroup/{id}?__json=][%d] getAccessGroupByIdJson default %s", o._statusCode, payload)
}

func (o *GetAccessGroupByIDJSONDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/accessgroup/{id}?__json=][%d] getAccessGroupByIdJson default %s", o._statusCode, payload)
}

func (o *GetAccessGroupByIDJSONDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAccessGroupByIDJSONDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}