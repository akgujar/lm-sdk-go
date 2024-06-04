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

// GetWebsiteGroupByIDJSONReader is a Reader for the GetWebsiteGroupByIDJSON structure.
type GetWebsiteGroupByIDJSONReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebsiteGroupByIDJSONReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWebsiteGroupByIDJSONOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetWebsiteGroupByIDJSONTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetWebsiteGroupByIDJSONDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWebsiteGroupByIDJSONOK creates a GetWebsiteGroupByIDJSONOK with default headers values
func NewGetWebsiteGroupByIDJSONOK() *GetWebsiteGroupByIDJSONOK {
	return &GetWebsiteGroupByIDJSONOK{}
}

/*
GetWebsiteGroupByIDJSONOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWebsiteGroupByIDJSONOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get website group by Id Json o k response has a 2xx status code
func (o *GetWebsiteGroupByIDJSONOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get website group by Id Json o k response has a 3xx status code
func (o *GetWebsiteGroupByIDJSONOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get website group by Id Json o k response has a 4xx status code
func (o *GetWebsiteGroupByIDJSONOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get website group by Id Json o k response has a 5xx status code
func (o *GetWebsiteGroupByIDJSONOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get website group by Id Json o k response a status code equal to that given
func (o *GetWebsiteGroupByIDJSONOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get website group by Id Json o k response
func (o *GetWebsiteGroupByIDJSONOK) Code() int {
	return 200
}

func (o *GetWebsiteGroupByIDJSONOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /website/groups/{id}?__json=][%d] getWebsiteGroupByIdJsonOK %s", 200, payload)
}

func (o *GetWebsiteGroupByIDJSONOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /website/groups/{id}?__json=][%d] getWebsiteGroupByIdJsonOK %s", 200, payload)
}

func (o *GetWebsiteGroupByIDJSONOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetWebsiteGroupByIDJSONOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebsiteGroupByIDJSONTooManyRequests creates a GetWebsiteGroupByIDJSONTooManyRequests with default headers values
func NewGetWebsiteGroupByIDJSONTooManyRequests() *GetWebsiteGroupByIDJSONTooManyRequests {
	return &GetWebsiteGroupByIDJSONTooManyRequests{}
}

/*
GetWebsiteGroupByIDJSONTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetWebsiteGroupByIDJSONTooManyRequests struct {

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

// IsSuccess returns true when this get website group by Id Json too many requests response has a 2xx status code
func (o *GetWebsiteGroupByIDJSONTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get website group by Id Json too many requests response has a 3xx status code
func (o *GetWebsiteGroupByIDJSONTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get website group by Id Json too many requests response has a 4xx status code
func (o *GetWebsiteGroupByIDJSONTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get website group by Id Json too many requests response has a 5xx status code
func (o *GetWebsiteGroupByIDJSONTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get website group by Id Json too many requests response a status code equal to that given
func (o *GetWebsiteGroupByIDJSONTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get website group by Id Json too many requests response
func (o *GetWebsiteGroupByIDJSONTooManyRequests) Code() int {
	return 429
}

func (o *GetWebsiteGroupByIDJSONTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /website/groups/{id}?__json=][%d] getWebsiteGroupByIdJsonTooManyRequests", 429)
}

func (o *GetWebsiteGroupByIDJSONTooManyRequests) String() string {
	return fmt.Sprintf("[GET /website/groups/{id}?__json=][%d] getWebsiteGroupByIdJsonTooManyRequests", 429)
}

func (o *GetWebsiteGroupByIDJSONTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetWebsiteGroupByIDJSONDefault creates a GetWebsiteGroupByIDJSONDefault with default headers values
func NewGetWebsiteGroupByIDJSONDefault(code int) *GetWebsiteGroupByIDJSONDefault {
	return &GetWebsiteGroupByIDJSONDefault{
		_statusCode: code,
	}
}

/*
GetWebsiteGroupByIDJSONDefault describes a response with status code -1, with default header values.

Error
*/
type GetWebsiteGroupByIDJSONDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get website group by Id Json default response has a 2xx status code
func (o *GetWebsiteGroupByIDJSONDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get website group by Id Json default response has a 3xx status code
func (o *GetWebsiteGroupByIDJSONDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get website group by Id Json default response has a 4xx status code
func (o *GetWebsiteGroupByIDJSONDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get website group by Id Json default response has a 5xx status code
func (o *GetWebsiteGroupByIDJSONDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get website group by Id Json default response a status code equal to that given
func (o *GetWebsiteGroupByIDJSONDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get website group by Id Json default response
func (o *GetWebsiteGroupByIDJSONDefault) Code() int {
	return o._statusCode
}

func (o *GetWebsiteGroupByIDJSONDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /website/groups/{id}?__json=][%d] getWebsiteGroupByIdJson default %s", o._statusCode, payload)
}

func (o *GetWebsiteGroupByIDJSONDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /website/groups/{id}?__json=][%d] getWebsiteGroupByIdJson default %s", o._statusCode, payload)
}

func (o *GetWebsiteGroupByIDJSONDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetWebsiteGroupByIDJSONDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
