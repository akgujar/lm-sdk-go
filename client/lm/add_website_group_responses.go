// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/logicmonitor/lm-sdk-go/v3/models"
)

// AddWebsiteGroupReader is a Reader for the AddWebsiteGroup structure.
type AddWebsiteGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddWebsiteGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddWebsiteGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewAddWebsiteGroupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddWebsiteGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddWebsiteGroupOK creates a AddWebsiteGroupOK with default headers values
func NewAddWebsiteGroupOK() *AddWebsiteGroupOK {
	return &AddWebsiteGroupOK{}
}

/* AddWebsiteGroupOK describes a response with status code 200, with default header values.

successful operation
*/
type AddWebsiteGroupOK struct {
	Payload *models.WebsiteGroup
}

func (o *AddWebsiteGroupOK) Error() string {
	return fmt.Sprintf("[POST /website/groups][%d] addWebsiteGroupOK  %+v", 200, o.Payload)
}
func (o *AddWebsiteGroupOK) GetPayload() *models.WebsiteGroup {
	return o.Payload
}

func (o *AddWebsiteGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebsiteGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddWebsiteGroupTooManyRequests creates a AddWebsiteGroupTooManyRequests with default headers values
func NewAddWebsiteGroupTooManyRequests() *AddWebsiteGroupTooManyRequests {
	return &AddWebsiteGroupTooManyRequests{}
}

/* AddWebsiteGroupTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AddWebsiteGroupTooManyRequests struct {

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

func (o *AddWebsiteGroupTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /website/groups][%d] addWebsiteGroupTooManyRequests ", 429)
}

func (o *AddWebsiteGroupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddWebsiteGroupDefault creates a AddWebsiteGroupDefault with default headers values
func NewAddWebsiteGroupDefault(code int) *AddWebsiteGroupDefault {
	return &AddWebsiteGroupDefault{
		_statusCode: code,
	}
}

/* AddWebsiteGroupDefault describes a response with status code -1, with default header values.

Error
*/
type AddWebsiteGroupDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add website group default response
func (o *AddWebsiteGroupDefault) Code() int {
	return o._statusCode
}

func (o *AddWebsiteGroupDefault) Error() string {
	return fmt.Sprintf("[POST /website/groups][%d] addWebsiteGroup default  %+v", o._statusCode, o.Payload)
}
func (o *AddWebsiteGroupDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddWebsiteGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
