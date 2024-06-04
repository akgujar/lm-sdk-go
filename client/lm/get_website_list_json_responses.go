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

// GetWebsiteListJSONReader is a Reader for the GetWebsiteListJSON structure.
type GetWebsiteListJSONReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebsiteListJSONReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWebsiteListJSONOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetWebsiteListJSONTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetWebsiteListJSONDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWebsiteListJSONOK creates a GetWebsiteListJSONOK with default headers values
func NewGetWebsiteListJSONOK() *GetWebsiteListJSONOK {
	return &GetWebsiteListJSONOK{}
}

/*
GetWebsiteListJSONOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWebsiteListJSONOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get website list Json o k response has a 2xx status code
func (o *GetWebsiteListJSONOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get website list Json o k response has a 3xx status code
func (o *GetWebsiteListJSONOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get website list Json o k response has a 4xx status code
func (o *GetWebsiteListJSONOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get website list Json o k response has a 5xx status code
func (o *GetWebsiteListJSONOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get website list Json o k response a status code equal to that given
func (o *GetWebsiteListJSONOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get website list Json o k response
func (o *GetWebsiteListJSONOK) Code() int {
	return 200
}

func (o *GetWebsiteListJSONOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /website/websites?__json=][%d] getWebsiteListJsonOK %s", 200, payload)
}

func (o *GetWebsiteListJSONOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /website/websites?__json=][%d] getWebsiteListJsonOK %s", 200, payload)
}

func (o *GetWebsiteListJSONOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetWebsiteListJSONOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebsiteListJSONTooManyRequests creates a GetWebsiteListJSONTooManyRequests with default headers values
func NewGetWebsiteListJSONTooManyRequests() *GetWebsiteListJSONTooManyRequests {
	return &GetWebsiteListJSONTooManyRequests{}
}

/*
GetWebsiteListJSONTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetWebsiteListJSONTooManyRequests struct {

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

// IsSuccess returns true when this get website list Json too many requests response has a 2xx status code
func (o *GetWebsiteListJSONTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get website list Json too many requests response has a 3xx status code
func (o *GetWebsiteListJSONTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get website list Json too many requests response has a 4xx status code
func (o *GetWebsiteListJSONTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get website list Json too many requests response has a 5xx status code
func (o *GetWebsiteListJSONTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get website list Json too many requests response a status code equal to that given
func (o *GetWebsiteListJSONTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get website list Json too many requests response
func (o *GetWebsiteListJSONTooManyRequests) Code() int {
	return 429
}

func (o *GetWebsiteListJSONTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /website/websites?__json=][%d] getWebsiteListJsonTooManyRequests", 429)
}

func (o *GetWebsiteListJSONTooManyRequests) String() string {
	return fmt.Sprintf("[GET /website/websites?__json=][%d] getWebsiteListJsonTooManyRequests", 429)
}

func (o *GetWebsiteListJSONTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetWebsiteListJSONDefault creates a GetWebsiteListJSONDefault with default headers values
func NewGetWebsiteListJSONDefault(code int) *GetWebsiteListJSONDefault {
	return &GetWebsiteListJSONDefault{
		_statusCode: code,
	}
}

/*
GetWebsiteListJSONDefault describes a response with status code -1, with default header values.

Error
*/
type GetWebsiteListJSONDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get website list Json default response has a 2xx status code
func (o *GetWebsiteListJSONDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get website list Json default response has a 3xx status code
func (o *GetWebsiteListJSONDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get website list Json default response has a 4xx status code
func (o *GetWebsiteListJSONDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get website list Json default response has a 5xx status code
func (o *GetWebsiteListJSONDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get website list Json default response a status code equal to that given
func (o *GetWebsiteListJSONDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get website list Json default response
func (o *GetWebsiteListJSONDefault) Code() int {
	return o._statusCode
}

func (o *GetWebsiteListJSONDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /website/websites?__json=][%d] getWebsiteListJson default %s", o._statusCode, payload)
}

func (o *GetWebsiteListJSONDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /website/websites?__json=][%d] getWebsiteListJson default %s", o._statusCode, payload)
}

func (o *GetWebsiteListJSONDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetWebsiteListJSONDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}