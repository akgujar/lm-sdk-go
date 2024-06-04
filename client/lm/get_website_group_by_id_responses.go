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

// GetWebsiteGroupByIDReader is a Reader for the GetWebsiteGroupByID structure.
type GetWebsiteGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebsiteGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWebsiteGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetWebsiteGroupByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetWebsiteGroupByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWebsiteGroupByIDOK creates a GetWebsiteGroupByIDOK with default headers values
func NewGetWebsiteGroupByIDOK() *GetWebsiteGroupByIDOK {
	return &GetWebsiteGroupByIDOK{}
}

/*
GetWebsiteGroupByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWebsiteGroupByIDOK struct {
	Payload *models.WebsiteGroup
}

// IsSuccess returns true when this get website group by Id o k response has a 2xx status code
func (o *GetWebsiteGroupByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get website group by Id o k response has a 3xx status code
func (o *GetWebsiteGroupByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get website group by Id o k response has a 4xx status code
func (o *GetWebsiteGroupByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get website group by Id o k response has a 5xx status code
func (o *GetWebsiteGroupByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get website group by Id o k response a status code equal to that given
func (o *GetWebsiteGroupByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get website group by Id o k response
func (o *GetWebsiteGroupByIDOK) Code() int {
	return 200
}

func (o *GetWebsiteGroupByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /website/groups/{id}][%d] getWebsiteGroupByIdOK %s", 200, payload)
}

func (o *GetWebsiteGroupByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /website/groups/{id}][%d] getWebsiteGroupByIdOK %s", 200, payload)
}

func (o *GetWebsiteGroupByIDOK) GetPayload() *models.WebsiteGroup {
	return o.Payload
}

func (o *GetWebsiteGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebsiteGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebsiteGroupByIDTooManyRequests creates a GetWebsiteGroupByIDTooManyRequests with default headers values
func NewGetWebsiteGroupByIDTooManyRequests() *GetWebsiteGroupByIDTooManyRequests {
	return &GetWebsiteGroupByIDTooManyRequests{}
}

/*
GetWebsiteGroupByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetWebsiteGroupByIDTooManyRequests struct {

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

// IsSuccess returns true when this get website group by Id too many requests response has a 2xx status code
func (o *GetWebsiteGroupByIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get website group by Id too many requests response has a 3xx status code
func (o *GetWebsiteGroupByIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get website group by Id too many requests response has a 4xx status code
func (o *GetWebsiteGroupByIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get website group by Id too many requests response has a 5xx status code
func (o *GetWebsiteGroupByIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get website group by Id too many requests response a status code equal to that given
func (o *GetWebsiteGroupByIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get website group by Id too many requests response
func (o *GetWebsiteGroupByIDTooManyRequests) Code() int {
	return 429
}

func (o *GetWebsiteGroupByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /website/groups/{id}][%d] getWebsiteGroupByIdTooManyRequests", 429)
}

func (o *GetWebsiteGroupByIDTooManyRequests) String() string {
	return fmt.Sprintf("[GET /website/groups/{id}][%d] getWebsiteGroupByIdTooManyRequests", 429)
}

func (o *GetWebsiteGroupByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetWebsiteGroupByIDDefault creates a GetWebsiteGroupByIDDefault with default headers values
func NewGetWebsiteGroupByIDDefault(code int) *GetWebsiteGroupByIDDefault {
	return &GetWebsiteGroupByIDDefault{
		_statusCode: code,
	}
}

/*
GetWebsiteGroupByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetWebsiteGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get website group by Id default response has a 2xx status code
func (o *GetWebsiteGroupByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get website group by Id default response has a 3xx status code
func (o *GetWebsiteGroupByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get website group by Id default response has a 4xx status code
func (o *GetWebsiteGroupByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get website group by Id default response has a 5xx status code
func (o *GetWebsiteGroupByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get website group by Id default response a status code equal to that given
func (o *GetWebsiteGroupByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get website group by Id default response
func (o *GetWebsiteGroupByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetWebsiteGroupByIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /website/groups/{id}][%d] getWebsiteGroupById default %s", o._statusCode, payload)
}

func (o *GetWebsiteGroupByIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /website/groups/{id}][%d] getWebsiteGroupById default %s", o._statusCode, payload)
}

func (o *GetWebsiteGroupByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetWebsiteGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
